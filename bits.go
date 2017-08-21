package humanize

import (
	"fmt"
	"math"
)

// IEC Sizes.
// kibis of bits
const (
	Bit = 1 << (iota * 10)
	KiBit
	MiBit
	GiBit
	TiBit
	PiBit
	EiBit
)

// SI Sizes.
const (
	IBit = 1
	KBit = IBit * 1000
	MBit = KBit * 1000
	GBit = MBit * 1000
	TBit = GBit * 1000
	PBit = TBit * 1000
	EBit = PBit * 1000
)

var bitsSizeTable = map[string]uint64{
	"b":   Bit,
	"kib": KiBit,
	"kb":  KBit,
	"mib": MiBit,
	"mb":  MBit,
	"gib": GiBit,
	"gb":  GBit,
	"tib": TiBit,
	"tb":  TBit,
	"pib": PiBit,
	"pb":  PBit,
	"eib": EiBit,
	"eb":  EBit,
	// Without suffix
	"":   Bit,
	"ki": KiBit,
	"k":  KBit,
	"mi": MiBit,
	"m":  MBit,
	"gi": GiBit,
	"g":  GBit,
	"ti": TiBit,
	"t":  TBit,
	"pi": PiBit,
	"p":  PBit,
	"ei": EiBit,
	"e":  EBit,
}

func humanateBits(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d b", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}

	return fmt.Sprintf(f, val, suffix)
}

// Bits produces a human readable representation of an SI size.
//
// See also: ParseBits.
//
// Bits(82854982) -> 83 Mb
func Bits(s uint64) string {
	sizes := []string{"b", "kb", "Mb", "Gb", "Tb", "Pb", "Eb"}
	return humanateBits(s, 1000, sizes)
}

// IBits produces a human readable representation of an IEC size.
//
// See also: ParseBits.
//
// IBits(82854982) -> 79 Mib
func IBits(s uint64) string {
	sizes := []string{"b", "Kib", "Mib", "Gib", "Tib", "Pib", "Eib"}
	return humanateBits(s, 1024, sizes)
}

// ParseBits parses a string representation of bits into the number
// of bits it represents.
//
// See Also: Bits, IBits, ParseBytes.
//
// ParseBits("42 Mb") -> 42000000, nil
// ParseBits("42 mib") -> 44040192, nil
func ParseBits(s string) (uint64, error) {
	return ParseBytes(s)
}
