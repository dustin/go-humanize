package humanize

import (
	"strconv"
	"strings"
)

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

func stripTrailingDigits(s string, digits int) string {
	if i := strings.Index(s, "."); i >= 0 {
		if digits <= 0 {
			return s[:i]
		}
		i++
		if i+digits >= len(s) {
			return s
		}
		return s[:i+digits]
	}
	return s
}

func clampTrailingDigits(s string, digits int) string {
	i := strings.Index(s, ".")

	// short circuit. if there is no decimal separator, add one,
	// and the appropriate number of zeros
	if i == -1 {
		// second short circuit - if they request <= 0 digits, just
		// return the string, as it already doesn't have any digits.
		if digits <= 0 {
			return s
		}
		return s + "." + strings.Repeat("0", digits)
	}

	// third short circuit - if they request <= 0 digits, remove
	// all digits and dot
	if digits <= 0 {
		return s[:i]
	}

	currentDigitLength := len(s)-i-1
	if digits <= currentDigitLength {
		return stripTrailingDigits(s,digits)
	}

	return s + strings.Repeat("0",digits-currentDigitLength)
}

// Ftoa converts a float to a string with no trailing zeros.
func Ftoa(num float64) string {
	return stripTrailingZeros(strconv.FormatFloat(num, 'f', 6, 64))
}

// FtoaWithDigits converts a float to a string but limits the resulting string
// to the given number of decimal places, and no trailing zeros.
func FtoaWithDigits(num float64, digits int) string {
	return stripTrailingZeros(stripTrailingDigits(strconv.FormatFloat(num, 'f', 6, 64), digits))
}
