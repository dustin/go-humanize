package humanize

import (
	"fmt"
	"math"
)

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

func log1024(n float64) float64 {
	return math.Log(n) / math.Log(1024)
}

func Bytes(s uint64) string {
	if s == 0 {
		return "0B"
	}
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	e := math.Floor(log1024(float64(s)))
	suffix := sizes[int(e)]
	return fmt.Sprintf("%.0f%s", float64(s)/math.Pow(1024, math.Floor(e)), suffix)
}
