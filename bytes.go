package humanize

import (
	"fmt"
	"math"
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

var (
	nameSizes  = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	iNameSizes = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
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

func precisionSprint(val float64, precision int) string {
	_, frac := math.Modf(val)
	if frac == 0 {
		return fmt.Sprintf("%.0f", val)
	}
	f := fmt.Sprintf("%%.%df", precision)
	res := fmt.Sprintf(f, val)
	return strings.TrimRight(res, "0")
}

func simpleSprint(val float64) string {
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}
	return fmt.Sprintf(f, val)
}

func humanateBytes(s uint64, base float64, sizes []string, precision int, specRoundFunc func (float64) float64) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]

	//var addlNumber float64
	//if !floorFlag {
	//	addlNumber = 0.5
	//}

	val := specRoundFunc(float64(s)/math.Pow(base, e)*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
	var roundVal string
	if precision > 1 {
		roundVal = precisionSprint(val, precision)
	} else {
		roundVal = simpleSprint(val)
	}

	return fmt.Sprintf("%s %s", roundVal, suffix)
}

// Bytes produces a human readable representation of an SI size.
//
// See also: ParseBytes.
//
// Bytes(82854982) -> 83 MB
func Bytes(s uint64) string {
	return humanateBytes(s, 1000, nameSizes, 1, math.Round)
}

// BytesCustomFloor allow to set precision and to get less or equal value with rounding
//
// BytesCustomFloor(92160871366656, 2) -> 92.16 TB
func BytesCustomFloor(s uint64, precision int) string {
	return humanateBytes(s, 1000, nameSizes, precision, math.Floor)
}

// BytesCustomCeil allow to set precision and to get more or equal value with rounding
//
// BytesCustomCeil(92160871366656, 2) -> 92.17 TB
func BytesCustomCeil(s uint64, precision int) string {
	return humanateBytes(s, 1000, nameSizes, precision, math.Ceil)
}

// IBytes produces a human readable representation of an IEC size.
//
// See also: ParseBytes.
//
// IBytes(82854982) -> 79 MiB
func IBytes(s uint64) string {
	return humanateBytes(s, 1024, iNameSizes, 1, math.Round)
}

// IBytesCustomFloor allow to set precision and to get less or equal value with rounding
//
// IBytesCustomFloor(92160871366656, 2) -> 83.81 TiB
// IBytesCustomFloor(92160871366656, 3) -> 83.8198242187
func IBytesCustomFloor(s uint64, precision int) string {
	return humanateBytes(s, 1024, iNameSizes, precision, math.Floor)
}

// IBytesCustomCeil allow to set precision and to get more or equal value with rounding
//
// IBytesCustomCeil(92160871366656, 2) -> 83.82 TiB
func IBytesCustomCeil(s uint64, precision int) string {
	return humanateBytes(s, 1024, iNameSizes, precision, math.Ceil)
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
		f *= float64(m)
		if f >= math.MaxUint64 {
			return 0, fmt.Errorf("too large: %v", s)
		}
		return uint64(f), nil
	}

	return 0, fmt.Errorf("unhandled size name: %v", extra)
}
