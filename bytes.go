package humanize

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"strconv"
	"strings"
	"unicode"
)

// IEC Sizes.
// kibis of bits
const (
	Byte = 1 << (iota * 10)
	KiByte
	MiByte
	GiByte
	TiByte
	PiByte
	EiByte
)

// SI Sizes.
const (
	IByte = 1
	KByte = IByte * 1000
	MByte = KByte * 1000
	GByte = MByte * 1000
	TByte = GByte * 1000
	PByte = TByte * 1000
	EByte = PByte * 1000
)

var bytesSizeTable = map[string]uint64{
	"b":   Byte,
	"kib": KiByte,
	"kb":  KByte,
	"mib": MiByte,
	"mb":  MByte,
	"gib": GiByte,
	"gb":  GByte,
	"tib": TiByte,
	"tb":  TByte,
	"pib": PiByte,
	"pb":  PByte,
	"eib": EiByte,
	"eb":  EByte,
	// Without suffix
	"":   Byte,
	"ki": KiByte,
	"k":  KByte,
	"mi": MiByte,
	"m":  MByte,
	"gi": GiByte,
	"g":  GByte,
	"ti": TiByte,
	"t":  TByte,
	"pi": PiByte,
	"p":  PByte,
	"ei": EiByte,
	"e":  EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func countDigits(n int64) int {
	digits := 0
	for n != 0 {
		n /= 10
		digits += 1
	}
	return digits
}

func countDigitsU(n uint64) int {
	digits := 0
	for n != 0 {
		n /= 10
		digits++
	}
	return digits
}

// pow10u holds 10^0 through 10^19, the largest power of ten that fits
// in a uint64.
var pow10u = [...]uint64{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000,
	1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000,
	100000000000000, 1000000000000000, 10000000000000000,
	100000000000000000, 1000000000000000000, 10000000000000000000,
}

// scaledRound returns s/pow rounded half-up to d decimal places,
// scaled by 10^d, in exact 128-bit integer arithmetic. It reports
// false when the result does not fit this path (d above 19, or the
// scaled quotient above 64 bits), and the caller falls back to
// big.Rat. s * 10^d is at most 128 bits for the digit counts used
// here, and the half-up test compares r against pow-r so 2*r never
// overflows.
func scaledRound(s, pow uint64, d int) (uint64, bool) {
	if d > 19 {
		return 0, false
	}
	hi, lo := bits.Mul64(s, pow10u[d])
	if hi >= pow { // quotient would not fit in 64 bits
		return 0, false
	}
	q, r := bits.Div64(hi, lo, pow)
	if r >= pow-r { // 2r >= pow without overflow
		if q == math.MaxUint64 {
			return 0, false
		}
		q++
	}
	return q, true
}

// formatScaled renders a value previously rounded with scaledRound at
// the same number of decimal places.
func formatScaled(q uint64, digits int) string {
	if digits == 0 {
		return strconv.FormatUint(q, 10)
	}
	scale := pow10u[digits]
	fracPart := strconv.FormatUint(q%scale, 10)
	for len(fracPart) < digits {
		fracPart = "0" + fracPart
	}
	return strconv.FormatUint(q/scale, 10) + "." + fracPart
}

// scaleBytes rounds s/pow half-up to the precision implied by
// minDigits, returning the scaled value (units of 10^-digits), the
// decimal count, and the unit index. It mirrors humanateBytesRat
// exactly, including the carry re-round and the promotion to the next
// unit, but without any allocation. It reports false when the exact
// result does not fit the 128-bit integer path, and the caller falls
// back to big.Rat.
func scaleBytes(s, pow, ib uint64, e, minDigits, nunits int) (uint64, int, int, bool) {
	intDigits := countDigitsU(s / pow)
	digits := max(minDigits-intDigits, 0)
	q, ok := scaledRound(s, pow, digits)
	if !ok {
		return 0, 0, e, false
	}

	// Rounding may have carried into an additional integer digit (e.g.
	// 9.996 -> 10.0), which means fewer decimal places should be shown
	// than originally computed.
	if newIntDigits := countDigitsU(q / pow10u[digits]); newIntDigits > intDigits {
		digits = max(digits-(newIntDigits-intDigits), 0)
		q, ok = scaledRound(s, pow, digits)
		if !ok {
			return 0, 0, e, false
		}
	}

	// A rounded value may now fit the next unit (1000 kB -> 1.0 MB).
	// The promoted value is exactly 1 in the next unit, since the
	// unrounded value is below ib and rounding half-up can reach at
	// most ib.
	if e+1 < nunits && q/pow10u[digits] >= ib {
		e++
		digits = max(minDigits-1, 0)
		q = pow10u[digits]
	}
	return q, digits, e, true
}

// ratFloor returns the greatest integer not greater than r. The values
// handled here are non-negative and well within int64 range.
func ratFloor(r *big.Rat) int64 {
	return new(big.Int).Quo(r.Num(), r.Denom()).Int64()
}

// roundRat rounds r to the given number of decimal places, rounding
// halves up, in exact arithmetic.
func roundRat(r *big.Rat, digits int) *big.Rat {
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil)
	scaled := new(big.Rat).Mul(r, new(big.Rat).SetInt(scale))
	scaled.Add(scaled, big.NewRat(1, 2))
	n := new(big.Int).Quo(scaled.Num(), scaled.Denom())
	return new(big.Rat).SetFrac(n, scale)
}

// formatRat renders a value previously rounded with roundRat at the
// same number of decimal places.
func formatRat(r *big.Rat, digits int) string {
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil)
	n := new(big.Int).Mul(r.Num(), scale)
	n.Quo(n, r.Denom())
	if digits == 0 {
		return n.String()
	}
	intPart := new(big.Int).Quo(n, scale)
	fracPart := new(big.Int).Mod(n, scale).String()
	for len(fracPart) < digits {
		fracPart = "0" + fracPart
	}
	return intPart.String() + "." + fracPart
}

func humanateBytes(s uint64, base float64, minDigits int, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}

	// base is always an integral value (1000 or 1024).
	ib := uint64(base)
	pow := uint64(1)
	e := 0
	for e+1 < len(sizes) && pow <= s/ib {
		pow *= ib
		e++
	}

	// The scaling and rounding are done in exact 128-bit integer
	// arithmetic where that fits (see scaleBytes); the big.Rat path
	// below remains only as a fallback for digit counts too large for
	// it. Converting s to float64 loses precision on large inputs, and
	// a scaled value that sits exactly on a rounding boundary (e.g.
	// 1.005 MB) is stored slightly below it in binary floating point,
	// so float arithmetic rounds it down instead of up (see #182).
	if q, digits, ne, ok := scaleBytes(s, pow, ib, e, minDigits, len(sizes)); ok {
		return formatScaled(q, digits) + " " + sizes[ne]
	}
	return humanateBytesRat(s, ib, pow, e, minDigits, sizes)
}

// humanateBytesRat is the big.Rat fallback for humanateBytes, used
// only when the exact result does not fit the 128-bit integer path
// (see scaleBytes). It follows the same steps: determine the precision
// from the unrounded value's integer-digit count, round exactly once,
// re-round with fewer digits when rounding carries into an additional
// integer digit, and promote to the next unit when the rounded value
// reaches it.
func humanateBytesRat(s, ib, pow uint64, e, minDigits int, sizes []string) string {
	val := new(big.Rat).SetFrac(new(big.Int).SetUint64(s), new(big.Int).SetUint64(pow))

	// Determine how many decimal places to keep based on the unrounded
	// value's integer-digit count, then round exactly once to that
	// precision. Rounding first to a fixed precision and then again when
	// formatting (double rounding) can flip the result at a boundary due
	// to floating point error (see #103).
	intDigits := countDigits(ratFloor(val))
	digits := max(minDigits-intDigits, 0)
	rounded := roundRat(val, digits)

	// Rounding may have carried into an additional integer digit (e.g.
	// 9.996 -> 10.0), which means fewer decimal places should be shown
	// than originally computed.
	if newIntDigits := countDigits(ratFloor(rounded)); newIntDigits > intDigits {
		digits = max(digits-(newIntDigits-intDigits), 0)
		rounded = roundRat(val, digits)
	}

	// A rounded value may now fit the next unit (1000 kB -> 1.0 MB).
	// Recompute the precision for that unit without rounding the input again.
	if rounded.Cmp(new(big.Rat).SetInt64(int64(ib))) >= 0 && e+1 < len(sizes) {
		e++
		val = new(big.Rat).Quo(rounded, new(big.Rat).SetInt64(int64(ib)))
		digits = max(minDigits-countDigits(ratFloor(val)), 0)
		rounded = roundRat(val, digits)
	}

	return fmt.Sprintf("%s %s", formatRat(rounded, digits), sizes[e])
}

// Bytes produces a human-readable representation of an SI size.
//
// See also: ParseBytes.
//
// Bytes(82854982) -> 83 MB
func Bytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, 2, sizes)
}

// BytesN produces a human-readable representation of an SI size.
// n specifies the total number of digits to output, including the decimal part.
// If n is less than or equal to the number of digits in the integer part, the decimal part will be omitted.
//
// See also: ParseBytes.
//
// BytesN(82854982, 3) -> 82.9 MB
// BytesN(82854982, 4) -> 82.85 MB
func BytesN(s uint64, n int) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, n, sizes)
}

// IBytes produces a human-readable representation of an IEC size.
//
// See also: ParseBytes.
//
// IBytes(82854982) -> 79 MiB
func IBytes(s uint64) string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return humanateBytes(s, 1024, 2, sizes)
}

// IBytesN produces a human-readable representation of an IEC size.
// n specifies the total number of digits to output, including the decimal part.
// If n is less than or equal to the number of digits in the integer part, the decimal part will be omitted.
//
// See also: ParseBytes.
//
// IBytesN(82854982, 4)  -> 79.02 MiB
// IBytesN(123456789, 3) -> 118 MiB
// IBytesN(123456789, 6) -> 117.738 MiB
func IBytesN(s uint64, n int) string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return humanateBytes(s, 1024, n, sizes)
}

// ParseBytes parses a string representation of bytes into the number
// of bytes it represents.
//
// See Also: Bytes, IBytes.
//
// ParseBytes("42 MB") -> 42000000, nil
// ParseBytes("42 mib") -> 44040192, nil
func ParseBytes(s string) (uint64, error) {
	lastDigit := 0
	hasComma := false
	for _, r := range s {
		if !(unicode.IsDigit(r) || r == '.' || r == ',') {
			break
		}
		if r == ',' {
			hasComma = true
		}
		lastDigit++
	}

	num := s[:lastDigit]
	if hasComma {
		num = strings.Replace(num, ",", "", -1)
	}

	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}

	extra := strings.ToLower(strings.TrimSpace(s[lastDigit:]))
	if m, ok := bytesSizeTable[extra]; ok {
		// Parse whole-number inputs exactly. A float64 cannot represent
		// integers above 2^53, and the multiplication below would also reject
		// otherwise-valid values near math.MaxUint64, so route integers through
		// uint64 arithmetic with an overflow check.
		if !strings.ContainsRune(num, '.') {
			if v, perr := strconv.ParseUint(num, 10, 64); perr == nil {
				if hi, lo := bits.Mul64(v, m); hi == 0 {
					return lo, nil
				}
				return 0, fmt.Errorf("too large: %v", s)
			}
		}
		f *= float64(m)
		if f >= math.MaxUint64 {
			return 0, fmt.Errorf("too large: %v", s)
		}
		return uint64(f), nil
	}

	return 0, fmt.Errorf("unhandled size name: %v", extra)
}
