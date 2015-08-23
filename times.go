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

// Magnitude stores one magnitude and the output format for this magnitude of relative time.
type Magnitude struct {
	d      time.Duration
	format string
	divby  time.Duration
}

// NewMagnitude returns a Magnitude object.
//
// d is the max number value of relative time for this magnitude.
// divby is the divisor to turn input number value into expected unit.
// format is the expected output format string.
//
// Also refer to RelTimeMagnitudes for examples.
func NewMagnitude(d time.Duration, format string, divby time.Duration) Magnitude {
	return Magnitude{
		d:      d,
		format: format,
		divby:  divby,
	}
}

var defaultMagnitudes = []Magnitude{
	NewMagnitude(time.Second, "now", time.Second),
	NewMagnitude(2*time.Second, "1 second %s", time.Second),
	NewMagnitude(Minute*time.Second, "%d seconds %s", time.Second),
	NewMagnitude(2*Minute*time.Second, "1 minute %s", time.Second),
	NewMagnitude(Hour*time.Second, "%d minutes %s", Minute*time.Second),
	NewMagnitude(2*Hour*time.Second, "1 hour %s", time.Second),
	NewMagnitude(Day*time.Second, "%d hours %s", Hour*time.Second),
	NewMagnitude(2*Day*time.Second, "1 day %s", time.Second),
	NewMagnitude(Week*time.Second, "%d days %s", Day*time.Second),
	NewMagnitude(2*Week*time.Second, "1 week %s", time.Second),
	NewMagnitude(Month*time.Second, "%d weeks %s", Week*time.Second),
	NewMagnitude(2*Month*time.Second, "1 month %s", time.Second),
	NewMagnitude(Year*time.Second, "%d months %s", Month*time.Second),
	NewMagnitude(18*Month*time.Second, "1 year %s", time.Second),
	NewMagnitude(2*Year*time.Second, "2 years %s", time.Second),
	NewMagnitude(LongTime*time.Second, "%d years %s", Year*time.Second),
	NewMagnitude(math.MaxInt64, "a long while %s", time.Second),
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
//		NewMagnitude(time.Second, "now", time.Second),
//		NewMagnitude(60*time.Second, "%d seconds %s", time.Second),
//		NewMagnitude(120*time.Second,"a minute %s", time.Second),
//		NewMagnitude(360*time.Second, "%d minutes %s", 60*time.Second),
// }
// albl: earlier
// blbl: later
//
// b - a                     output
//  -130*time.Second         2 minutes later
//   0                       now
//   30*time.Second          30 seconds earlier
//   80*time.Second          a minute  earlier
//   340*time.Second         5 minutes earlier
//   400*time.Second         undefined
func RelTimeMagnitudes(a, b time.Time, albl, blbl string, magnitudes []Magnitude) string {
	lbl := albl
	diff := b.Sub(a)

	after := a.After(b)
	if after {
		lbl = blbl
		diff = a.Sub(b)
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].d >= diff
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
