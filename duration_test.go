package humanize

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	type test struct {
		format   string
		d        time.Duration
		expected string
	}
	cases := map[string]test{
		"default": {
			format:   "%y %m %w %d %H %M %S",
			d:        400*Day + 12*time.Hour + 17*time.Second,
			expected: "1 year 1 month 5 days 12 hours 17 seconds",
		},
		"years": {
			format:   "%y",
			d:        800 * Day,
			expected: "2 years",
		},
		"year + week": {
			format:   "%y %w",
			d:        364 * Day,
			expected: "52 weeks",
		},
		"month + hours": {
			format:   "%w %H",
			d:        15 * Day,
			expected: "2 weeks 24 hours",
		},
		"seconds only": {
			format:   "%S",
			d:        Day,
			expected: "86400 seconds",
		},
	}
	for name, v := range cases {
		result := Duration(v.format, v.d)
		if result != v.expected {
			t.Errorf("FAIL:%s\n\tExpected:%s\n\tActual:%s", name, v.expected, result)
		}
	}
}
