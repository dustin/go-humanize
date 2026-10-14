package humanize

import "strconv"

// Ordinal gives you the input number in a rank/ordinal format.
//
// Ordinal(3) -> 3rd
// Ordinal(-1) -> -1st
func Ordinal(x int) string {
	// Suffix selection must use the absolute value because Go's % for
	// negative dividends yields a negative remainder (e.g. -1%10 == -1),
	// which would never match the positive cases below.
	n := x
	if n < 0 {
		n = -n
	}

	suffix := "th"
	switch n % 10 {
	case 1:
		if n%100 != 11 {
			suffix = "st"
		}
	case 2:
		if n%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if n%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
