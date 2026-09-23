//go:build go1.18
// +build go1.18

package humanize_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/dustin/go-humanize"
)

func FuzzBigCommaf(f *testing.F) {
	f.Add(0.0)
	f.Add(-1.0)
	f.Add(10.11)
	f.Add(-10.11)
	f.Add(1234567.5)
	f.Add(-1234567.5)
	f.Add(math.MaxFloat64)
	f.Add(-math.MaxFloat64)
	f.Add(math.SmallestNonzeroFloat64)
	f.Add(-math.SmallestNonzeroFloat64)
	f.Add(math.Inf(1))
	f.Add(math.Inf(-1))

	f.Fuzz(func(t *testing.T, v float64) {
		if math.IsNaN(v) {
			t.Skip("big.Float cannot represent NaN")
		}

		in := big.NewFloat(v)
		want := new(big.Float).Copy(in)

		got := humanize.BigCommaf(in)

		// BigCommaf must not modify its argument.
		if in.Cmp(want) != 0 {
			t.Fatalf("BigCommaf(%v) modified its argument: got %v, want %v", v, in, want)
		}

		// Calling it again on the same (unmodified) value should produce
		// the same result.
		if got2 := humanize.BigCommaf(in); got2 != got {
			t.Fatalf("BigCommaf(%v) not idempotent: first %q, second %q", v, got, got2)
		}
	})
}
