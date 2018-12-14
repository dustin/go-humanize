package humanize

import (
	"testing"
)

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
		// Bug #42
		{"1,005.03 MB", 1005030000},
		// Large testing, breaks when too much larger than
		// this.
		{"12.5 EB", uint64(12.5 * float64(EByte))},
		{"12.5 E", uint64(12.5 * float64(EByte))},
		{"12.5 EiB", uint64(12.5 * float64(EiByte))},
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
	got, err = ParseBytes("")
	if err == nil {
		t.Errorf("Expected error parsing nothing")
	}
	got, err = ParseBytes("16 EiB")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
}

func TestBytes(t *testing.T) {
	testList{
		{"bytes(0)", Bytes(0), "0 B"},
		{"bytes(1)", Bytes(1), "1 B"},
		{"bytes(803)", Bytes(803), "803 B"},
		{"bytes(999)", Bytes(999), "999 B"},

		{"bytes(1024)", Bytes(1024), "1.0 kB"},
		{"bytes(9999)", Bytes(9999), "10 kB"},
		{"bytes(1MB - 1)", Bytes(MByte - Byte), "1000 kB"},

		{"bytes(1MB)", Bytes(1024 * 1024), "1.0 MB"},
		{"bytes(1GB - 1K)", Bytes(GByte - KByte), "1000 MB"},

		{"bytes(1GB)", Bytes(GByte), "1.0 GB"},
		{"bytes(1TB - 1M)", Bytes(TByte - MByte), "1000 GB"},
		{"bytes(10MB)", Bytes(9999 * 1000), "10 MB"},

		{"bytes(1TB)", Bytes(TByte), "1.0 TB"},
		{"bytes(1PB - 1T)", Bytes(PByte - TByte), "999 TB"},

		{"bytes(1PB)", Bytes(PByte), "1.0 PB"},
		{"bytes(1PB - 1T)", Bytes(EByte - PByte), "999 PB"},

		{"bytes(1EB)", Bytes(EByte), "1.0 EB"},
		{"bytes(92160871366656)", Bytes(92160871366656), "92 TB"},
		// Overflows.
		// {"bytes(1EB - 1P)", Bytes((KByte*EByte)-PByte), "1023EB"},

		{"bytes(0)", IBytes(0), "0 B"},
		{"bytes(1)", IBytes(1), "1 B"},
		{"bytes(803)", IBytes(803), "803 B"},
		{"bytes(1023)", IBytes(1023), "1023 B"},

		{"bytes(1024)", IBytes(1024), "1.0 KiB"},
		{"bytes(1MB - 1)", IBytes(MiByte - IByte), "1024 KiB"},

		{"bytes(1MB)", IBytes(1024 * 1024), "1.0 MiB"},
		{"bytes(1GB - 1K)", IBytes(GiByte - KiByte), "1024 MiB"},

		{"bytes(1GB)", IBytes(GiByte), "1.0 GiB"},
		{"bytes(1TB - 1M)", IBytes(TiByte - MiByte), "1024 GiB"},

		{"bytes(1TB)", IBytes(TiByte), "1.0 TiB"},
		{"bytes(1PB - 1T)", IBytes(PiByte - TiByte), "1023 TiB"},

		{"bytes(1PB)", IBytes(PiByte), "1.0 PiB"},
		{"bytes(1PB - 1T)", IBytes(EiByte - PiByte), "1023 PiB"},

		{"bytes(1EiB)", IBytes(EiByte), "1.0 EiB"},
		// Overflows.
		// {"bytes(1EB - 1P)", IBytes((KIByte*EIByte)-PiByte), "1023EB"},

		{"bytes(5.5GiB)", IBytes(5.5 * GiByte), "5.5 GiB"},

		{"bytes(92160871366656)", IBytes(92160871366656), "84 TiB"},
	}.validate(t)
}

func TestBytesCustomFloor(t *testing.T) {
	testList{
		{"bytes(0) with precision(2) ", BytesCustomFloor(0, 2), "0 B"},
		{"bytes(1) with precision(2)", BytesCustomFloor(1, 2), "1 B"},
		{"bytes(803) with precision(2)", BytesCustomFloor(803, 2), "803 B"},
		{"bytes(999) with precision(2)", BytesCustomFloor(999, 2), "999 B"},

		{"bytes(1) with precision(2)", BytesCustomFloor(1, 2), "1 B"},
		{"bytes(803) with precision(2)", BytesCustomFloor(803, 2), "803 B"},
		{"bytes(999) with precision(2)", BytesCustomFloor(999, 2), "999 B"},

		{"bytes(1024) with precision(2)", BytesCustomFloor(1024, 2), "1.02 kB"},
		{"bytes(9999) with precision(2)", BytesCustomFloor(9999, 2), "9.99 kB"},
		{"bytes(1MB - 1) with precision(2)", BytesCustomFloor(MByte-Byte, 2), "999.99 kB"},

		{"bytes(1MB) with precision(2)", BytesCustomFloor(1024*1024, 2), "1.04 MB"},
		{"bytes(1GB - 1K) with precision(2)", BytesCustomFloor(GByte-KByte, 2), "999.99 MB"},

		{"bytes(1GB) with precision(2)", BytesCustomFloor(GByte, 2), "1 GB"},
		{"bytes(1TB - 1M) with precision(2)", BytesCustomFloor(TByte-MByte, 2), "999.99 GB"},
		{"bytes(10MB) with precision(2)", BytesCustomFloor(9999*1000, 2), "9.99 MB"},

		{"bytes(1TB) with precision(2)", BytesCustomFloor(TByte, 2), "1 TB"},
		{"bytes(1PB - 1T) with precision(2)", BytesCustomFloor(PByte-TByte, 2), "999 TB"},

		{"bytes(1PB) with precision(2)", BytesCustomFloor(PByte, 2), "1 PB"},
		{"bytes(1PB - 1T) with precision(2)", BytesCustomFloor(EByte-PByte, 2), "999 PB"},

		{"bytes(1EB) with precision(2)", BytesCustomFloor(EByte, 2), "1 EB"},

		{"bytes(92160871366656) with precision(2)", BytesCustomFloor(92160871366656, 2), "92.16 TB"},
		{"bytes(92160871366656) with precision(10)", BytesCustomFloor(92160871366656, 10), "92.1608713666 TB"},
		{"bytes(92160871366656) with precision(3)", BytesCustomFloor(92160866656, 3), "92.16 GB"},
		{"bytes(92160871366656) with precision(2)", BytesCustomFloor(92160866656, 2), "92.16 GB"},
		{"bytes(92160871366656) with precision(0)", BytesCustomFloor(102160871366656, 0), "102 TB"},
		{"bytes(92160871366656) with precision(20)", BytesCustomFloor(102160871366656, 20), "102.16087136665599643948 TB"},

		{"bytes(0) with precision(2)", IBytesCustomFloor(0, 2), "0 B"},
		{"bytes(1) with precision(2)", IBytesCustomFloor(1, 2), "1 B"},
		{"bytes(803) with precision(2)", IBytesCustomFloor(803, 2), "803 B"},
		{"bytes(1023) with precision(2)", IBytesCustomFloor(1023, 2), "1023 B"},

		{"bytes(1024) with precision(2)", IBytesCustomFloor(1024, 2), "1 KiB"},
		{"bytes(1MB - 1) with precision(2)", IBytesCustomFloor(MiByte-IByte, 2), "1023.99 KiB"},

		{"bytes(1MB) with precision(2)", IBytesCustomFloor(1024*1024, 2), "1 MiB"},
		{"bytes(1GB - 1K) with precision(2)", IBytesCustomFloor(GiByte-KiByte, 2), "1023.99 MiB"},

		{"bytes(1GB) with precision(2)", IBytesCustomFloor(GiByte, 2), "1 GiB"},
		{"bytes(1TB - 1M) with precision(2)", IBytesCustomFloor(TiByte-MiByte, 2), "1023.99 GiB"},

		{"bytes(1TB) with precision(2)", IBytesCustomFloor(TiByte, 2), "1 TiB"},
		{"bytes(1PB - 1T) with precision(2)", IBytesCustomFloor(PiByte-TiByte, 2), "1023 TiB"},

		{"bytes(1PB) with precision(2)", IBytesCustomFloor(PiByte, 2), "1 PiB"},
		{"bytes(1PB - 1T) with precision(2)", IBytesCustomFloor(EiByte-PiByte, 2), "1023 PiB"},

		{"bytes(1EiB) with precision(1)", IBytesCustomFloor(EiByte, 2), "1 EiB"},

		{"bytes(5.5GiB) with precision(3)", IBytesCustomFloor(5.5*GiByte, 3), "5.5 GiB"},

		{"bytes(92160871366656) with precision(2)", IBytesCustomFloor(92160871366656, 2), "83.81 TiB"},
		{"bytes(92160871366656) with precision(10)", IBytesCustomFloor(92160871366656, 10), "83.8198242187 TiB"},
		{"bytes(92160871366656) with precision(3)", IBytesCustomFloor(92160866656, 3), "85.831 GiB"},
		{"bytes(92160871366656) with precision(2)", IBytesCustomFloor(92160866656, 2), "85.83 GiB"},
		{"bytes(92160871366656) with precision(0)", IBytesCustomFloor(102160871366656, 0), "92 TiB"},
		{"bytes(92160871366656) with precision(20)", IBytesCustomFloor(102160871366656, 20), "92.91477123647928237915 TiB"},
	}.validate(t)
}
func TestBytesCustomCeil(t *testing.T) {
	testList{
		{"Ceil :bytes(0) with precision(2) ", BytesCustomCeil(0, 2), "0 B"},
		{"Ceil :bytes(1) with precision(2)", BytesCustomCeil(1, 2), "1 B"},
		{"Ceil :bytes(803) with precision(2)", BytesCustomCeil(803, 2), "803 B"},
		{"Ceil :bytes(999) with precision(2)", BytesCustomCeil(999, 2), "999 B"},
		{"Ceil :bytes(1) with precision(2)", BytesCustomCeil(1, 2), "1 B"},
		{"Ceil :bytes(803) with precision(2)", BytesCustomCeil(803, 2), "803 B"},
		{"Ceil :bytes(999) with precision(2)", BytesCustomCeil(999, 2), "999 B"},
		{"Ceil :bytes(1024) with precision(2)", BytesCustomCeil(1024, 2), "1.03 kB"},
		{"Ceil :bytes(9999) with precision(2)", BytesCustomCeil(9999, 2), "10 kB"},
		{"Ceil :bytes(1MB - 1) with precision(2)", BytesCustomCeil(MByte-Byte, 2), "1000 kB"},
		{"Ceil :bytes(1MB) with precision(2)", BytesCustomCeil(1024*1024, 2), "1.05 MB"},
		{"Ceil :bytes(1GB - 1K) with precision(2)", BytesCustomCeil(GByte-KByte, 2), "1000 MB"},
		{"Ceil :bytes(1GB) with precision(2)", BytesCustomCeil(GByte, 2), "1 GB"},
		{"Ceil :bytes(1TB - 1M) with precision(2)", BytesCustomCeil(TByte-MByte, 2), "1000 GB"},
		{"Ceil :bytes(10MB) with precision(2)", BytesCustomCeil(9999*1000, 2), "10 MB"},
		{"Ceil :bytes(1TB) with precision(2)", BytesCustomCeil(TByte, 2), "1 TB"},
		{"Ceil :bytes(1PB - 1T) with precision(2)", BytesCustomCeil(PByte-TByte, 2), "999 TB"},
		{"Ceil :bytes(1PB) with precision(2)", BytesCustomCeil(PByte, 2), "1 PB"},
		{"Ceil :bytes(1PB - 1T) with precision(2)", BytesCustomCeil(EByte-PByte, 2), "999 PB"},
		{"Ceil :bytes(1EB) with precision(2)", BytesCustomCeil(EByte, 2), "1 EB"},
		{"Ceil :bytes(92160871366656) with precision(2)", BytesCustomCeil(92160871366656, 2), "92.17 TB"},
		{"Ceil :bytes(92160871366656) with precision(10)", BytesCustomCeil(92160871366656, 10), "92.1608713667 TB"},
		{"Ceil :bytes(92160871366656) with precision(3)", BytesCustomCeil(92160866656, 3), "92.161 GB"},
		{"Ceil :bytes(92160871366656) with precision(2)", BytesCustomCeil(92160866656, 2), "92.17 GB"},
		{"Ceil :bytes(92160871366656) with precision(0)", BytesCustomCeil(102160871366656, 0), "103 TB"},
		{"Ceil :bytes(92160871366656) with precision(20)", BytesCustomCeil(102160871366656, 20), "102.16087136665599643948 TB"},
		{"Ceil :bytes(0) with precision(2)", IBytesCustomCeil(0, 2), "0 B"},
		{"Ceil :bytes(1) with precision(2)", IBytesCustomCeil(1, 2), "1 B"},
		{"Ceil :bytes(803) with precision(2)", IBytesCustomCeil(803, 2), "803 B"},
		{"Ceil :bytes(1023) with precision(2)", IBytesCustomCeil(1023, 2), "1023 B"},
		{"Ceil :bytes(1024) with precision(2)", IBytesCustomCeil(1024, 2), "1 KiB"},
		{"Ceil :bytes(1MB - 1) with precision(2)", IBytesCustomCeil(MiByte-IByte, 2), "1024 KiB"},
		{"Ceil :bytes(1MB) with precision(2)", IBytesCustomCeil(1024*1024, 2), "1 MiB"},
		{"Ceil :bytes(1GB - 1K) with precision(2)", IBytesCustomCeil(GiByte-KiByte, 2), "1024 MiB"},
		{"Ceil :bytes(1GB) with precision(2)", IBytesCustomCeil(GiByte, 2), "1 GiB"},
		{"Ceil :bytes(1TB - 1M) with precision(2)", IBytesCustomCeil(TiByte-MiByte, 2), "1024 GiB"},
		{"Ceil :bytes(1TB) with precision(2)", IBytesCustomCeil(TiByte, 2), "1 TiB"},
		{"Ceil :bytes(1PB - 1T) with precision(2)", IBytesCustomCeil(PiByte-TiByte, 2), "1023 TiB"},
		{"Ceil :bytes(1PB) with precision(2)", IBytesCustomCeil(PiByte, 2), "1 PiB"},
		{"Ceil :bytes(1PB - 1T) with precision(2)", IBytesCustomCeil(EiByte-PiByte, 2), "1023 PiB"},
		{"Ceil :bytes(1EiB) with precision(1)", IBytesCustomCeil(EiByte, 2), "1 EiB"},
		{"Ceil :bytes(5.5GiB) with precision(3)", IBytesCustomCeil(5.5*GiByte, 3), "5.5 GiB"},
		{"Ceil :bytes(92160871366656) with precision(2)", IBytesCustomCeil(92160871366656, 2), "83.82 TiB"},
		{"Ceil :bytes(92160871366656) with precision(10)", IBytesCustomCeil(92160871366656, 10), "83.8198242188 TiB"},
		{"Ceil :bytes(92160871366656) with precision(3)", IBytesCustomCeil(92160866656, 3), "85.832 GiB"},
		{"Ceil :bytes(92160871366656) with precision(2)", IBytesCustomCeil(92160866656, 2), "85.84 GiB"},
		{"Ceil :bytes(92160871366656) with precision(0)", IBytesCustomCeil(102160871366656, 0), "93 TiB"},
		{"Ceil :bytes(92160871366656) with precision(20)", IBytesCustomCeil(102160871366656, 20), "92.91477123647928237915 TiB"},
	}.validate(t)
}

func BenchmarkParseBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseBytes("16.5 GB")
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(16.5 * GByte)
	}
}
