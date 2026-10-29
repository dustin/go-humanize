package humanize

import "strconv"

// Ordinal gives you the input number in a rank/ordinal format.
//
// Ordinal(3) -> 3rd
func Ordinal(x int) string {
	mod10 := x % 10
	if mod10 < 0 {
		mod10 = -mod10
	}
	mod100 := x % 100
	if mod100 < 0 {
		mod100 = -mod100
	}
	suffix := "th"
	switch mod10 {
	case 1:
		if mod100 != 11 {
			suffix = "st"
		}
	case 2:
		if mod100 != 12 {
			suffix = "nd"
		}
	case 3:
		if mod100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
