package humanize

import (
	"strconv"
	"strings"
	"testing"
)

func TestFormatIntegerExact(t *testing.T) {
	tests := []struct {
		value   string
		grouped string
	}{
		{"0", "0"},
		{"1", "1"},
		{"-1", "-1"},
		{"12345", "12,345"},
		{"-12345", "-12,345"},
		{"2147483647", "2,147,483,647"},
		{"-2147483648", "-2,147,483,648"},
		{"4503599627370497", "4,503,599,627,370,497"},
		{"-4503599627370497", "-4,503,599,627,370,497"},
		{"9007199254740991", "9,007,199,254,740,991"},
		{"9007199254740993", "9,007,199,254,740,993"},
		{"-9007199254740993", "-9,007,199,254,740,993"},
		{"9223372036854775807", "9,223,372,036,854,775,807"},
		{"-9223372036854775808", "-9,223,372,036,854,775,808"},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			n, err := strconv.Atoi(tt.value)
			if err != nil {
				if strconv.IntSize == 32 {
					t.Skip("value requires a 64-bit int")
				}
				t.Fatal(err)
			}
			positive := ""
			if n > 0 {
				positive = "+"
			}
			formats := []struct {
				format string
				want   string
			}{
				{"#.", tt.value},
				{"#", tt.value + ".000000000"},
				{"", tt.grouped + ".00"},
				{"#,###.", tt.grouped},
				{"+#,###.##", positive + tt.grouped + ".00"},
				{"#.###,######", strings.ReplaceAll(tt.grouped, ",", ".") + ",000000"},
				{"#\u202f###,##", strings.ReplaceAll(tt.grouped, ",", "\u202f") + ",00"},
				{"#,##", tt.value + ",00"},
			}
			for _, format := range formats {
				t.Run(format.format, func(t *testing.T) {
					if got := FormatInteger(format.format, n); got != format.want {
						t.Fatalf("got %q, want %q", got, format.want)
					}
				})
			}
		})
	}
}

func TestFormatIntegerInvalidFormat(t *testing.T) {
	for _, format := range []string{"-", "0.01", "0.0000000000"} {
		t.Run(format, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected invalid format to panic")
				}
			}()
			FormatInteger(format, 12345)
		})
	}
}
