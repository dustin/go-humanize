package humanize

import (
	"fmt"
	"regexp"
)

var trailingZerosRegex = regexp.MustCompile(`\.?0+$`)

// Ftoa converts a float to a string with no trailing zeros.
func Ftoa(num float64) string {
	str := fmt.Sprintf("%f", num)
	return trailingZerosRegex.ReplaceAllString(str, "")
}
