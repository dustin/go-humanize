package humanize

import (
	"testing"
)

func assert(t *testing.T, name string, got interface{}, expected interface{}) {
	if got != expected {
		t.Errorf("Expected %#v for %s, got %#v", expected, name, got)
	}
}

func TestByteParsing(t *testing.T) {
	tests := []struct {
		in  string
		exp uint64
	}{
		{"42", 42},
		{"42MB", 42000000},
		{"42MiB", 44040192},
		{"42mb", 42000000},
		{"42mib", 44040192},
		{"42MIB", 44040192},
		{"42 MB", 42000000},
		{"42 MiB", 44040192},
		{"42 mb", 42000000},
		{"42 mib", 44040192},
		{"42 MIB", 44040192},
		{"42.5MB", 42500000},
		{"42.5MiB", 44564480},
		{"42.5 MB", 42500000},
		{"42.5 MiB", 44564480},
		// No need to say B
		{"42M", 42000000},
		{"42Mi", 44040192},
		{"42m", 42000000},
		{"42mi", 44040192},
		{"42MI", 44040192},
		{"42 M", 42000000},
		{"42 Mi", 44040192},
		{"42 m", 42000000},
		{"42 mi", 44040192},
		{"42 MI", 44040192},
		{"42.5M", 42500000},
		{"42.5Mi", 44564480},
		{"42.5 M", 42500000},
		{"42.5 Mi", 44564480},
		// Large testing, breaks when too much larger than
		// this.
		{"12.5 EB", uint64(12.5 * float64(EByte))},
		{"12.5 E", uint64(12.5 * float64(EByte))},
	}

	for _, p := range tests {
		got, err := ParseBytes(p.in)
		if err != nil {
			t.Errorf("Couldn't parse %v: %v", p.in, err)
		}
		if got != p.exp {
			t.Errorf("Expected %v for %v, got %v",
				p.exp, p.in, got)
		}
	}
}

func TestByteErrors(t *testing.T) {
	got, err := ParseBytes("84 JB")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
	// TODO: Figure out how to induce failure in the float parser.
}

func TestBytes(t *testing.T) {
	assert(t, "bytes(0)", Bytes(0), "0B")
	assert(t, "bytes(1)", Bytes(1), "1B")
	assert(t, "bytes(803)", Bytes(803), "803B")
	assert(t, "bytes(999)", Bytes(999), "999B")
}

func TestK(t *testing.T) {
	assert(t, "bytes(1024)", Bytes(1024), "1.0KB")
	assert(t, "bytes(1MB - 1)", Bytes(MByte-Byte), "1000KB")
}

func TestM(t *testing.T) {
	assert(t, "bytes(1MB)", Bytes(1024*1024), "1.0MB")
	assert(t, "bytes(1GB - 1K)", Bytes(GByte-KByte), "1000MB")
}

func TestG(t *testing.T) {
	assert(t, "bytes(1GB)", Bytes(GByte), "1.0GB")
	assert(t, "bytes(1TB - 1M)", Bytes(TByte-MByte), "1000GB")
}

func TestT(t *testing.T) {
	assert(t, "bytes(1TB)", Bytes(TByte), "1.0TB")
	assert(t, "bytes(1PB - 1T)", Bytes(PByte-TByte), "999TB")
}

func TestP(t *testing.T) {
	assert(t, "bytes(1PB)", Bytes(PByte), "1.0PB")
	assert(t, "bytes(1PB - 1T)", Bytes(EByte-PByte), "999PB")
}

func TestE(t *testing.T) {
	assert(t, "bytes(1EB)", Bytes(EByte), "1.0EB")
	// Overflows.
	// assert(t, "bytes(1EB - 1P)", Bytes((KByte*EByte)-PByte), "1023EB")
}

func TestIIBytes(t *testing.T) {
	assert(t, "bytes(0)", IBytes(0), "0B")
	assert(t, "bytes(1)", IBytes(1), "1B")
	assert(t, "bytes(803)", IBytes(803), "803B")
	assert(t, "bytes(1023)", IBytes(1023), "1023B")
}

func TestIK(t *testing.T) {
	assert(t, "bytes(1024)", IBytes(1024), "1.0KiB")
	assert(t, "bytes(1MB - 1)", IBytes(MiByte-IByte), "1024KiB")
}

func TestIM(t *testing.T) {
	assert(t, "bytes(1MB)", IBytes(1024*1024), "1.0MiB")
	assert(t, "bytes(1GB - 1K)", IBytes(GiByte-KiByte), "1024MiB")
}

func TestIG(t *testing.T) {
	assert(t, "bytes(1GB)", IBytes(GiByte), "1.0GiB")
	assert(t, "bytes(1TB - 1M)", IBytes(TiByte-MiByte), "1024GiB")
}

func TestIT(t *testing.T) {
	assert(t, "bytes(1TB)", IBytes(TiByte), "1.0TiB")
	assert(t, "bytes(1PB - 1T)", IBytes(PiByte-TiByte), "1023TiB")
}

func TestIP(t *testing.T) {
	assert(t, "bytes(1PB)", IBytes(PiByte), "1.0PiB")
	assert(t, "bytes(1PB - 1T)", IBytes(EiByte-PiByte), "1023PiB")
}

func TestIE(t *testing.T) {
	assert(t, "bytes(1EiB)", IBytes(EiByte), "1.0EiB")
	// Overflows.
	// assert(t, "bytes(1EB - 1P)", IBytes((KIByte*EIByte)-PiByte), "1023EB")
}

func TestIHalf(t *testing.T) {
	assert(t, "bytes(5.5GiB)", IBytes(5.5*GiByte), "5.5GiB")
}

func TestHalf(t *testing.T) {
	assert(t, "bytes(5.5GB)", Bytes(5.5*GByte), "5.5GB")
}
