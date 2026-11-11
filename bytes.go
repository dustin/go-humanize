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

// pow10u returns 10**n as a big.Int. n is required to be non-negative.
func pow10u(n int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
}

// roundHalfUp returns round(s/unit, digits) * 10**digits as an integer,
// using exact rational arithmetic so float64 rounding cannot drop a half.
func roundHalfUp(s, unit uint64, digits int) *big.Int {
	num := new(big.Int).SetUint64(s)
	num.Mul(num, pow10u(digits))
	den := new(big.Int).SetUint64(unit)
	half := new(big.Int).Rsh(new(big.Int).Set(den), 1)
	num.Add(num, half)
	return num.Div(num, den)
}

// formatScaled renders q/10**digits with exactly `digits` decimal places.
func formatScaled(q *big.Int, digits int) string {
	s := q.String()
	if digits == 0 {
		return s
	}
	for len(s) <= digits {
		s = "0" + s
	}
	return s[:len(s)-digits] + "." + s[len(s)-digits:]
}

func integerDigits(q *big.Int, digits int) int {
	n := len(q.String()) - digits
	if n < 1 {
		return 1
	}
	return n
}

func humanateBytes(s uint64, base float64, minDigits int, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	baseU := uint64(base)

	// Choose the unit with integer arithmetic so values above 2**53
	// (where float64 cannot represent every integer) stay exact.
	e := 0
	unit := uint64(1)
	for e+1 < len(sizes) {
		if unit > ^uint64(0)/baseU {
			break
		}
		next := unit * baseU
		if s < next {
			break
		}
		unit = next
		e++
	}

	for {
		intDigits := countDigits(int64(s / unit))
		if intDigits == 0 {
			intDigits = 1
		}
		digits := max(minDigits-intDigits, 0)
		q := roundHalfUp(s, unit, digits)
		if newIntDigits := integerDigits(q, digits); newIntDigits > intDigits {
			digits = max(digits-(newIntDigits-intDigits), 0)
			q = roundHalfUp(s, unit, digits)
		}
		// Promote when rounding reaches the next unit (1000 kB -> 1.00 MB).
		if e+1 < len(sizes) {
			threshold := new(big.Int).SetUint64(baseU)
			threshold.Mul(threshold, pow10u(digits))
			if q.Cmp(threshold) >= 0 {
				if unit > ^uint64(0)/baseU {
					return fmt.Sprintf("%s %s", formatScaled(q, digits), sizes[e])
				}
				e++
				unit *= baseU
				continue
			}
		}
		return fmt.Sprintf("%s %s", formatScaled(q, digits), sizes[e])
	}
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
