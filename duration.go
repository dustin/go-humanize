package humanize

import (
	"strconv"
	"strings"
	"time"
)

type format struct {
	key  string
	unit time.Duration
	name string
}

// Duration produces a formated duration as a string
// it supports the following format flags:
// %y - year
// %m - month
// %w - week
// %H - hour
// %M - minute
// %S - second
func Duration(format string, d time.Duration) string {
	if format == "" {
		format = "%y %m %w %d %H %M %S"
	}
	process := func(f string, unit time.Duration, name string) {
		if strings.Contains(format, f) {
			segment := int(d / unit)
			format = strings.Replace(format, f, pluralize(segment, name), -1)
			d %= unit
		}
	}

	process("%y", 365*Day, "year")
	process("%m", Month, "month")
	process("%w", Week, "week")
	process("%d", Day, "day")
	process("%H", time.Hour, "hour")
	process("%M", time.Minute, "minute")
	process("%S", time.Second, "second")

	// cleanup spaces
	format = strings.Trim(format, " ")
	return strings.Replace(format, "  ", " ", -1)
}

func pluralize(i int, s string) string {
	if i == 0 {
		return ""
	}
	s = strconv.Itoa(i) + " " + s

	if i > 1 {
		s += "s"
	}
	return s
}
