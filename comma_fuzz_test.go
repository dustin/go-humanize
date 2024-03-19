//go:build go1.18
// +build go1.18

package humanize_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/dustin/go-humanize"
)

func FuzzComma(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(-1))
	f.Add(int64(10))
	f.Add(int64(100))
	f.Add(int64(1000))
	f.Add(int64(-1000))
	f.Add(int64(math.MaxInt64))
	f.Add(int64(math.MaxInt64) - 1)
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MinInt64) + 1)

	f.Fuzz(func(t *testing.T, v int64) {
		got := humanize.Comma(v)
		gotNoCommas := strings.ReplaceAll(got, ",", "")
		expected := strconv.FormatInt(v, 10)
		if gotNoCommas != expected {
			t.Fatalf("%d: got %q, expected %q", v, got, expected)
		}

		if v < 0 {
			if got[0] != '-' {
				t.Fatalf("%d: got: %q", v, got)
			}
			// Remove sign
			got = got[1:]
		}
		// Check that commas are located every 3 digits
		l := len(got)
		for i := l - 1; i >= 0; i-- {
			var ok bool
			if (l-1-i)%4 == 3 {
				ok = got[i] == ','
			} else {
				ok = got[i] >= '0' && got[i] <= '9'
			}
			if !ok {
				t.Log(l - 1 - i)
				t.Log((l - 1 - i) % 4)
				t.Fatalf("%d: got: %q", v, got)
			}
		}
	})
}
