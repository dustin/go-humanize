package humanize

import "strconv"

// Ordinal gives you the input number in a rank/ordinal format.
//
// e.g. Ordinal(3) -> 3rd
func Ordinal(x int) string {
	return strconv.Itoa(x) + calculateSuffix(x)
}

// OrdinalComma gives you the input number in a rank/ordinal format,
// with commas after every three orders of magnitude.
//
// e.g. OrdinalComma(834143) -> 834,143rd
func OrdinalComma(x int) string {
	return Comma(int64(x)) + calculateSuffix(x)
}

func calculateSuffix(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return suffix
}
