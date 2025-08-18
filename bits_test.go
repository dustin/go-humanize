package humanize

import (
	"testing"
)

func TestBitParsing(t *testing.T) {
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
		{"42 Mb", 42000000},
		{"42 Mib", 44040192},
		{"42 mb", 42000000},
		{"42 mib", 44040192},
		{"42 MIb", 44040192},
		{"42.5MB", 42500000},
		{"42.5MiB", 44564480},
		{"42.5 Mb", 42500000},
		{"42.5 Mib", 44564480},
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
		{"1,005.03 Mb", 1005030000},
		// Large testing, breaks when too much larger than
		// this.
		{"12.5 Eb", uint64(12.5 * float64(EBit))},
		{"12.5 E", uint64(12.5 * float64(EBit))},
		{"12.5 Eib", uint64(12.5 * float64(EiBit))},
	}

	for _, p := range tests {
		got, err := ParseBits(p.in)
		if err != nil {
			t.Errorf("Couldn't parse %v: %v", p.in, err)
		}
		if got != p.exp {
			t.Errorf("Expected %v for %v, got %v",
				p.exp, p.in, got)
		}
	}
}

func TestBitErrors(t *testing.T) {
	got, err := ParseBits("84 Jb")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
	got, err = ParseBits("")
	if err == nil {
		t.Errorf("Expected error parsing nothing")
	}
	got, err = ParseBits("16 Eib")
	if err == nil {
		t.Errorf("Expected error, got %v", got)
	}
}

func TestBits(t *testing.T) {
	testList{
		{"bits(0)", Bits(0), "0 b"},
		{"bits(1)", Bits(1), "1 b"},
		{"bits(803)", Bits(803), "803 b"},
		{"bits(999)", Bits(999), "999 b"},

		{"bits(1024)", Bits(1024), "1.0 kb"},
		{"bits(9999)", Bits(9999), "10 kb"},
		{"bits(1Mb - 1)", Bits(MBit - Bit), "1000 kb"},

		{"bits(1Mb)", Bits(1024 * 1024), "1.0 Mb"},
		{"bits(1Gb - 1K)", Bits(GBit - KBit), "1000 Mb"},

		{"bits(1Gb)", Bits(GBit), "1.0 Gb"},
		{"bits(1Tb - 1M)", Bits(TBit - MBit), "1000 Gb"},
		{"bits(10Mb)", Bits(9999 * 1000), "10 Mb"},

		{"bits(1Tb)", Bits(TBit), "1.0 Tb"},
		{"bits(1Pb - 1T)", Bits(PBit - TBit), "999 Tb"},

		{"bits(1Pb)", Bits(PBit), "1.0 Pb"},
		{"bits(1Pb - 1T)", Bits(EBit - PBit), "999 Pb"},

		{"bits(1Eb)", Bits(EBit), "1.0 Eb"},
		// Overflows.
		// {"bits(1EB - 1P)", Bits((KBit*EBit)-PBit), "1023EB"},

		{"bits(0)", IBits(0), "0 b"},
		{"bits(1)", IBits(1), "1 b"},
		{"bits(803)", IBits(803), "803 b"},
		{"bits(1023)", IBits(1023), "1023 b"},

		{"bits(1024)", IBits(1024), "1.0 Kib"},
		{"bits(1Mb - 1)", IBits(MiBit - IBit), "1024 Kib"},

		{"bits(1Mb)", IBits(1024 * 1024), "1.0 Mib"},
		{"bits(1Gb - 1K)", IBits(GiBit - KiBit), "1024 Mib"},

		{"bits(1Gb)", IBits(GiBit), "1.0 Gib"},
		{"bits(1Tb - 1M)", IBits(TiBit - MiBit), "1024 Gib"},

		{"bits(1Tb)", IBits(TiBit), "1.0 Tib"},
		{"bits(1Pb - 1T)", IBits(PiBit - TiBit), "1023 Tib"},

		{"bits(1Pb)", IBits(PiBit), "1.0 Pib"},
		{"bits(1Pb - 1T)", IBits(EiBit - PiBit), "1023 Pib"},

		{"bits(1Eib)", IBits(EiBit), "1.0 Eib"},
		// Overflows.
		// {"bits(1EB - 1P)", IBits((KIBit*EIBit)-PiBit), "1023EB"},

		{"bits(5.5Gib)", IBits(5.5 * GiBit), "5.5 Gib"},

		{"bits(5.5Gb)", Bits(5.5 * GBit), "5.5 Gb"},
	}.validate(t)
}

func BenchmarkParseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseBits("16.5 Gb")
	}
}

func BenchmarkBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bits(16.5 * GBit)
	}
}
