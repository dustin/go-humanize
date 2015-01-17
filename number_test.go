package humanize

import (
	"math"
	"testing"
)

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		name      string
		format    string
		num       float64
		formatted string
	}{
		{"default", "", 12345.6789, "12,345.68"},
		{"#", "#", 12345.6789, "12345.678900000"},
		{"#.", "#.", 12345.6789, "12346"},
		{"#,#", "#,#", 12345.6789, "12345,7"},
		{"#,##", "#,##", 12345.6789, "12345,68"},
		{"#,###", "#,###", 12345.6789, "12345,679"},
		{"#,###.", "#,###.", 12345.6789, "12,346"},
		{"#,###.##", "#,###.##", 12345.6789, "12,345.68"},
		{"#,###.###", "#,###.###", 12345.6789, "12,345.679"},
		{"#,###.####", "#,###.####", 12345.6789, "12,345.6789"},
		{"#.###,######", "#.###,######", 12345.6789, "12.345,678900"},
		{"#\u202f###,##", "#\u202f###,##", 12345.6789, "12 345,68"},

		// special cases
		{"NaN", "", math.NaN(), "NaN"},
		{"+Inf", "", math.Inf(1), "Infinity"},
		{"-Inf", "", math.Inf(-1), "-Infinity"},
	}

	for _, test := range tests {
		got := FormatFloat(test.format, test.num)
		if got != test.formatted {
			t.Errorf("On %v (%v, %v), got %v, wanted %v",
				test.name, test.format, test.num, got, test.formatted)
		}
	}
	//TODO: FormatFloat may panic with a bad format. How do we test this?
}
