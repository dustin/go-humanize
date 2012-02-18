package humanize

import (
	"fmt"
	"strings"
)

func reverse(in string) string {
	if len(in) < 2 {
		return in
	}
	bytes := []byte(in)
	j := 0
	for i := len(bytes) - 1; i >= len(bytes)/2; i-- {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		j++
	}
	return string(bytes)
}

// Place commas after every three orders of magnitude.
func Comma(v int64) string {
	sign := ""
	if v < 0 {
		sign = "-"
		v = 0 - v
	}
	s := reverse(fmt.Sprintf("%v", v))
	parts := []string{}
	for len(s) > 0 {
		l := len(s)
		if l > 3 {
			l = 3
		}
		parts = append(parts, s[0:l])
		s = s[l:]
	}
	return sign + reverse(strings.Join(parts, ","))
}
