package humanize

import "testing"

func TestFtoaWithDigitsLargeLimit(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	for _, digits := range []int{6, 7, maxInt - 1, maxInt} {
		for _, num := range []float64{1.25, -1.25, 1234.25} {
			if got, want := FtoaWithDigits(num, digits), Ftoa(num); got != want {
				t.Errorf("FtoaWithDigits(%v, %d) = %q, want %q", num, digits, got, want)
			}
		}
	}
}
