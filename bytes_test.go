package humanize

import (
	"testing"
)

func assert(t *testing.T, name string, got interface{}, expected interface{}) {
	if got != expected {
		t.Fatalf("Expected %#v for %s, got %#v", expected, name, got)
	}
}

func TestBytes(t *testing.T) {
	assert(t, "bytes(0)", Bytes(0), "0B")
	assert(t, "bytes(1)", Bytes(1), "1B")
	assert(t, "bytes(803)", Bytes(803), "803B")
	assert(t, "bytes(1023)", Bytes(1023), "1023B")
}

func TestK(t *testing.T) {
	assert(t, "bytes(1024)", Bytes(1024), "1KB")
	assert(t, "bytes(1MB - 1)", Bytes(MByte-Byte), "1024KB")
}

func TestM(t *testing.T) {
	assert(t, "bytes(1MB)", Bytes(1024*1024), "1MB")
	assert(t, "bytes(1GB - 1K)", Bytes(GByte-KByte), "1024MB")
}

func TestG(t *testing.T) {
	assert(t, "bytes(1GB)", Bytes(GByte), "1GB")
	assert(t, "bytes(1TB - 1M)", Bytes(TByte-MByte), "1024GB")
}

func TestT(t *testing.T) {
	assert(t, "bytes(1TB)", Bytes(TByte), "1TB")
	assert(t, "bytes(1PB - 1T)", Bytes(PByte-TByte), "1023TB")
}

func TestP(t *testing.T) {
	assert(t, "bytes(1PB)", Bytes(PByte), "1PB")
	assert(t, "bytes(1PB - 1T)", Bytes(EByte-PByte), "1023PB")
}

func TestE(t *testing.T) {
	assert(t, "bytes(1EB)", Bytes(EByte), "1EB")
	// Overflows.
	// assert(t, "bytes(1EB - 1P)", Bytes((KByte*EByte)-PByte), "1023EB")
}
