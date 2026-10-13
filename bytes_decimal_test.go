package humanize

import "testing"

func TestParseBytesExactDecimals(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  uint64
	}{
		{"1.001 kB", 1001},
		{"1.005 kB", 1005},
		{"9007199254740993.0 B", 9007199254740993},
		{"18446744073709551615.0 B", 18446744073709551615},
		{"9,007,199,254,740,993.0", 9007199254740993},
		{".5 KiB", 512},
		{"1.999 B", 1},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, err := ParseBytes(tc.input)
			if err != nil || got != tc.want {
				t.Fatalf("got %d, %v; want %d", got, err, tc.want)
			}
		})
	}
	for _, input := range []string{"18446744073709551616.0 B", "16.0 EiB", "1.2.3 MB"} {
		if _, err := ParseBytes(input); err == nil {
			t.Errorf("accepted invalid or overflowing size %q", input)
		}
	}
}
