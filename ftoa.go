package humanize

import "fmt"

func stripTrailingZeros(s string) string {
	offset := len(s) - 1
	for offset > 0 {
		if s[offset] == '.' {
			offset--
			break
		}
		if s[offset] != '0' {
			break
		}
		offset--
	}
	return s[:offset+1]
}

// Ftoa converts a float to a string with no trailing zeros.
func Ftoa(num float64, offset int) string {
	f := fmt.Sprintf("%%.%df", offset)
	return stripTrailingZeros(fmt.Sprintf(f, num))
}
