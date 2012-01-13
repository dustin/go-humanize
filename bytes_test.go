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
	assert(t, "bytes(0)", Bytes(0).String(), "0B")
	assert(t, "bytes(1)", Bytes(1).String(), "1B")
	assert(t, "bytes(803)", Bytes(803).String(), "803B")
	assert(t, "bytes(999)", Bytes(999).String(), "999B")
}

func TestK(t *testing.T) {
	assert(t, "bytes(1024)", Bytes(1024).String(), "1KB")
	assert(t, "bytes(1MB - 1)", Bytes(MByte-Byte).String(), "1000KB")
}

func TestM(t *testing.T) {
	assert(t, "bytes(1MB)", Bytes(1024*1024).String(), "1MB")
	assert(t, "bytes(1GB - 1K)", Bytes(GByte-KByte).String(), "1000MB")
}

func TestG(t *testing.T) {
	assert(t, "bytes(1GB)", Bytes(GByte).String(), "1GB")
	assert(t, "bytes(1TB - 1M)", Bytes(TByte-MByte).String(), "1000GB")
}

func TestT(t *testing.T) {
	assert(t, "bytes(1TB)", Bytes(TByte).String(), "1TB")
	assert(t, "bytes(1PB - 1T)", Bytes(PByte-TByte).String(), "999TB")
}

func TestP(t *testing.T) {
	assert(t, "bytes(1PB)", Bytes(PByte).String(), "1PB")
	assert(t, "bytes(1PB - 1T)", Bytes(EByte-PByte).String(), "999PB")
}

func TestE(t *testing.T) {
	assert(t, "bytes(1EB)", Bytes(EByte).String(), "1EB")
	// Overflows.
	// assert(t, "bytes(1EB - 1P)", Bytes((KByte*EByte)-PByte), "1023EB")
}

func TestIIBytes(t *testing.T) {
	assert(t, "bytes(0)", IBytes(0).String(), "0B")
	assert(t, "bytes(1)", IBytes(1).String(), "1B")
	assert(t, "bytes(803)", IBytes(803).String(), "803B")
	assert(t, "bytes(1023)", IBytes(1023).String(), "1023B")
}

func TestIK(t *testing.T) {
	assert(t, "bytes(1024)", IBytes(1024).String(), "1KiB")
	assert(t, "bytes(1MB - 1)", IBytes(MiByte-IByte).String(), "1024KiB")
}

func TestIM(t *testing.T) {
	assert(t, "bytes(1MB)", IBytes(1024*1024).String(), "1MiB")
	assert(t, "bytes(1GB - 1K)", IBytes(GiByte-KiByte).String(), "1024MiB")
}

func TestIG(t *testing.T) {
	assert(t, "bytes(1GB)", IBytes(GiByte).String(), "1GiB")
	assert(t, "bytes(1TB - 1M)", IBytes(TiByte-MiByte).String(), "1024GiB")
}

func TestIT(t *testing.T) {
	assert(t, "bytes(1TB)", IBytes(TiByte).String(), "1TiB")
	assert(t, "bytes(1PB - 1T)", IBytes(PiByte-TiByte).String(), "1023TiB")
}

func TestIP(t *testing.T) {
	assert(t, "bytes(1PB)", IBytes(PiByte).String(), "1PiB")
	assert(t, "bytes(1PB - 1T)", IBytes(EiByte-PiByte).String(), "1023PiB")
}

func TestIE(t *testing.T) {
	assert(t, "bytes(1EiB)", IBytes(EiByte).String(), "1EiB")
	// Overflows.
	// assert(t, "bytes(1EB - 1P)", IBytes((KIByte*EIByte)-PiByte), "1023EB")
}
