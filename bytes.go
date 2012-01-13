package humanize

import (
	"fmt"
	"math"
)

const (
	Byte   = 1
	KiByte = Byte * 1024
	MiByte = KiByte * 1024
	GiByte = MiByte * 1024
	TiByte = GiByte * 1024
	PiByte = TiByte * 1024
	EiByte = PiByte * 1024
)

const (
	IByte = 1
	KByte = IByte * 1000
	MByte = KByte * 1000
	GByte = MByte * 1000
	TByte = GByte * 1000
	PByte = TByte * 1000
	EByte = PByte * 1000
)

type Bytes uint64

type IBytes uint64

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

func (s Bytes) String() string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1000, sizes)
}

func (s IBytes) String() string {
	sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return humanateBytes(uint64(s), 1024, sizes)
}
