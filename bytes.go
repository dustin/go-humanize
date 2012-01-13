package humanize

import (
	"fmt"
	"math"
)

// IEC Sizes.
// kibis of bits
const (
	Byte   = 1
	KiByte = Byte * 1024
	MiByte = KiByte * 1024
	GiByte = MiByte * 1024
	TiByte = GiByte * 1024
	PiByte = TiByte * 1024
	EiByte = PiByte * 1024
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

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s == 0 {
		return "0B"
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	return fmt.Sprintf("%.0f%s", float64(s)/math.Pow(base, math.Floor(e)), suffix)

}

// String up an SI size.
// Bytes(82854982) -> 83MB
func Bytes(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1000, sizes)
}

// String an IEC size.
// IBytes(82854982) -> 79MiB
func IBytes(s uint64) string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return humanateBytes(uint64(s), 1024, sizes)
}
