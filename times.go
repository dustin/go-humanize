package humanize

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// Seconds-based time units
const (
	Minute   = 60
	Hour     = 60 * Minute
	Day      = 24 * Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 37 * Year
)

// Time formats a time into a relative string.
//
// Time(someT) -> "3 weeks ago"
func Time(then time.Time) string {
	return RelTime(then, time.Now(), "ago", "from now")
}

// Seconds represents the elapsed time between two instants as an int64 second count.
type Seconds int64

// Magnitude stores one magnitude and the output format for this magnitude of relative time.
type Magnitude struct {
	d      int64
	format string
	divby  int64
}

// NewMagnitude returns a Magnitude object.
//
// d is the max number value of relative time for this magnitude.
// divby is the divisor to turn input number value into expected unit.
// format is the expected output format string.
//
// Also refer to RelTimeMagnitudes for examples.
func NewMagnitude(d Seconds, format string, divby Seconds) Magnitude {
	return Magnitude{
		d:      int64(d),
		format: format,
		divby:  int64(divby),
	}
}

var defaultMagnitudes = []Magnitude{
	NewMagnitude(1, "now", 1),
	NewMagnitude(2, "1 second %s", 1),
	NewMagnitude(Minute, "%d seconds %s", 1),
	NewMagnitude(2*Minute, "1 minute %s", 1),
	NewMagnitude(Hour, "%d minutes %s", Minute),
	NewMagnitude(2*Hour, "1 hour %s", 1),
	NewMagnitude(Day, "%d hours %s", Hour),
	NewMagnitude(2*Day, "1 day %s", 1),
	NewMagnitude(Week, "%d days %s", Day),
	NewMagnitude(2*Week, "1 week %s", 1),
	NewMagnitude(Month, "%d weeks %s", Week),
	NewMagnitude(2*Month, "1 month %s", 1),
	NewMagnitude(Year, "%d months %s", Month),
	NewMagnitude(18*Month, "1 year %s", 1),
	NewMagnitude(2*Year, "2 years %s", 1),
	NewMagnitude(LongTime, "%d years %s", Year),
	NewMagnitude(math.MaxInt64, "a long while %s", 1),
}

// RelTime formats a time into a relative string.
//
// It takes two times and two labels.  In addition to the generic time
// delta string (e.g. 5 minutes), the labels are used applied so that
// the label corresponding to the smaller time is applied.
//
// RelTime(timeInPast, timeInFuture, "earlier", "later") -> "3 weeks earlier"
func RelTime(a, b time.Time, albl, blbl string) string {
	return RelTimeMagnitudes(a, b, albl, blbl, defaultMagnitudes)
}

// RelTimeMagnitudes accepts a magnitudes parameter to allow custom defined units and output format.
//
// example:
// magitudes:
// {
//		NewMagnitude(1, "now", 1),
//		NewMagnitude(60, "%d seconds %s", 1),
//		NewMagnitude(120,"a minute %s", 1),
//		NewMagnitude(360, "%d minutes %s", 60),
// }
// albl: earlier
// blbl: later
//
// b - a         output
//  -130         2 minutes later
//   0           now
//   30          30 seconds earlier
//   80          a minute  earlier
//   340         5 minutes earlier
//   400         undefined
func RelTimeMagnitudes(a, b time.Time, albl, blbl string, magnitudes []Magnitude) string {
	lbl := albl
	diff := b.Unix() - a.Unix()

	after := a.After(b)
	if after {
		lbl = blbl
		diff = a.Unix() - b.Unix()
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].d > diff
	})

	if n >= len(magnitudes) {
		return "undefined"
	}

	mag := magnitudes[n]
	args := []interface{}{}
	escaped := false
	for _, ch := range mag.format {
		if escaped {
			switch ch {
			case 's':
				args = append(args, lbl)
			case 'd':
				args = append(args, diff/mag.divby)
			}
			escaped = false
		} else {
			escaped = ch == '%'
		}
	}
	return fmt.Sprintf(mag.format, args...)
}
