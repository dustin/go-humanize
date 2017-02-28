package humanize

import (
	"testing"
	"time"
)

func TestDurations(t *testing.T) {
	tests := []struct {
		In       time.Duration
		Expected string
	}{
		{time.Duration(0), "0 nanoseconds"},
		{time.Nanosecond, "1 nanosecond"},
		{999 * time.Nanosecond, "999 nanoseconds"},
		{time.Microsecond, "1 microsecond"},
		{999 * time.Microsecond, "999 microseconds"},
		{time.Millisecond, "1 millisecond"},
		{999 * time.Millisecond, "999 milliseconds"},
		{time.Second, "1 second"},
		{59 * time.Second, "59 seconds"},
		{time.Minute, "1 minute"},
		{59 * time.Minute, "59 minutes"},
		{time.Hour, "1 hour"},
		{999 * time.Hour, "999 hours"},
	}

	for _, test := range tests {
		if out := Duration(test.In); out != test.Expected {
			t.Errorf("%s: got %s, want %s", test.In, out, test.Expected)
		}
	}
}
