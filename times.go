package humanize

import (
	"fmt"
	"math"
	"sort"
	"strings"
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

// AccurateTime formats a time into a relative string with a reasonable
// level of accuracy.
//
// AccurateTime(someT) -> "3 weeks, 2 days, 11 hours ago"
func AccurateTime(then time.Time) string {
	return RelTimeMagnitudes(then, time.Now(), "ago", "from now", 2)
}

var magnitudes = []struct {
	d        int64
	format   string
	doFormat bool
	length   int64
}{
	//{1, "now", false, 1},
	{2, "1 second", false, 1},
	{Minute, "%d seconds", true, 1},
	{2 * Minute, "1 minute", false, Minute},
	{Hour, "%d minutes", true, Minute},
	{2 * Hour, "1 hour", false, Hour},
	{Day, "%d hours", true, Hour},
	{2 * Day, "1 day", false, Day},
	{Week, "%d days", true, Day},
	{2 * Week, "1 week", false, Week},
	{Month, "%d weeks", true, Week},
	{2 * Month, "1 month", false, Month},
	{Year, "%d months", true, Month},
	{2 * Year, "1 year", false, Year},
	{LongTime, "%d years", true, Year},
	{math.MaxInt64, "a long while", false, 1},
}

// RelTimeMagnitudes formats a time into a relative string.
//
// It takes two times and two labels and a magnitude.  In addition to
// the generic time delta string (e.g. 5 minutes), the labels are
// used applied so that the label corresponding to the smaller time is
// applied.
//
// The magnitude determines how accurate the time should be, a magnitude
// of 2 will return strings like "2 years, 8 months ago", 3 will return
// "2 years, 8 months, 1 week ago", etc.
//
// RelTimeMagnitude(timeInPast, timeInFuture, "earlier", "later", 2) -> "3 weeks, 2 days earlier"
func RelTimeMagnitudes(a, b time.Time, albl, blbl string, numMagnitudes int) string {
	lbl := albl
	diff := b.Unix() - a.Unix()

	after := a.After(b)
	if after {
		lbl = blbl
		diff = a.Unix() - b.Unix()
	}

	if diff == 0 {
		return "now"
	}

	strs := []string{}
	for i := 0; i < numMagnitudes; i++ {
		n := sort.Search(len(magnitudes), func(idx int) bool {
			return magnitudes[idx].d > diff
		})

		if magnitudes[n].doFormat {
			strs = append(strs, fmt.Sprintf(magnitudes[n].format, diff/magnitudes[n].length))
		} else {
			strs = append(strs, magnitudes[n].format)
		}

		diff -= magnitudes[n].length * (diff / magnitudes[n].length)
		if diff <= 0 {
			break
		}
	}

	return fmt.Sprintf("%s %s", strings.Join(strs, ", "), lbl)
}

// RelTime formats a time into a relative string.
//
// It takes two times and two labels.  In addition to the generic time
// delta string (e.g. 5 minutes), the labels are used applied so that
// the label corresponding to the smaller time is applied.
//
// RelTime(timeInPast, timeInFuture, "earlier", "later") -> "3 weeks earlier"
func RelTime(a, b time.Time, albl, blbl string) string {
	return RelTimeMagnitudes(a, b, albl, blbl, 1)
}
