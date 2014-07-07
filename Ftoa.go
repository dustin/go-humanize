package humanize

import (
	"fmt"
	"regexp"
)

var trailingZerosRegex = regexp.MustCompile(`\.?0+$`)

func Ftoa(num float64) string {
	str := fmt.Sprintf("%f", num)
	return trailingZerosRegex.ReplaceAllString(str, "")
}
