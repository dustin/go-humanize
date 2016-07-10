package humanize

import (
	"math"
	"testing"
	"time"
)

func TestPast(t *testing.T) {
	now := time.Now().Unix()
	testList{
		{"now", Time(time.Unix(now, 0)), "now"},
		{"1 second ago", Time(time.Unix(now-1, 0)), "1 second ago"},
		{"12 seconds ago", Time(time.Unix(now-12, 0)), "12 seconds ago"},
		{"30 seconds ago", Time(time.Unix(now-30, 0)), "30 seconds ago"},
		{"45 seconds ago", Time(time.Unix(now-45, 0)), "45 seconds ago"},
		{"1 minute ago", Time(time.Unix(now-63, 0)), "1 minute ago"},
		{"15 minutes ago", Time(time.Unix(now-15*Minute, 0)), "15 minutes ago"},
		{"1 hour ago", Time(time.Unix(now-63*Minute, 0)), "1 hour ago"},
		{"2 hours ago", Time(time.Unix(now-2*Hour, 0)), "2 hours ago"},
		{"21 hours ago", Time(time.Unix(now-21*Hour, 0)), "21 hours ago"},
		{"1 day ago", Time(time.Unix(now-26*Hour, 0)), "1 day ago"},
		{"2 days ago", Time(time.Unix(now-49*Hour, 0)), "2 days ago"},
		{"3 days ago", Time(time.Unix(now-3*Day, 0)), "3 days ago"},
		{"1 week ago (1)", Time(time.Unix(now-7*Day, 0)), "1 week ago"},
		{"1 week ago (2)", Time(time.Unix(now-12*Day, 0)), "1 week ago"},
		{"2 weeks ago", Time(time.Unix(now-15*Day, 0)), "2 weeks ago"},
		{"1 month ago", Time(time.Unix(now-39*Day, 0)), "1 month ago"},
		{"3 months ago", Time(time.Unix(now-99*Day, 0)), "3 months ago"},
		{"1 year ago (1)", Time(time.Unix(now-365*Day, 0)), "1 year ago"},
		{"1 year ago (1)", Time(time.Unix(now-400*Day, 0)), "1 year ago"},
		{"2 years ago (1)", Time(time.Unix(now-548*Day, 0)), "2 years ago"},
		{"2 years ago (2)", Time(time.Unix(now-725*Day, 0)), "2 years ago"},
		{"2 years ago (3)", Time(time.Unix(now-800*Day, 0)), "2 years ago"},
		{"3 years ago", Time(time.Unix(now-3*Year, 0)), "3 years ago"},
		{"long ago", Time(time.Unix(now-LongTime, 0)), "a long while ago"},
	}.validate(t)
}

func TestFuture(t *testing.T) {
	now := time.Now().Unix()
	// add 1 second offset for test time to balance decimal fraction of time.Now()
	offset := int64(time.Second)
	testList{
		{"now", Time(time.Unix(now, 0)), "now"},
		{"1 second from now", Time(time.Unix(now+1, offset)), "1 second from now"},
		{"12 seconds from now", Time(time.Unix(now+12, offset)), "12 seconds from now"},
		{"30 seconds from now", Time(time.Unix(now+30, offset)), "30 seconds from now"},
		{"45 seconds from now", Time(time.Unix(now+45, offset)), "45 seconds from now"},
		{"15 minutes from now", Time(time.Unix(now+15*Minute, offset)), "15 minutes from now"},
		{"2 hours from now", Time(time.Unix(now+2*Hour, offset)), "2 hours from now"},
		{"21 hours from now", Time(time.Unix(now+21*Hour, offset)), "21 hours from now"},
		{"1 day from now", Time(time.Unix(now+26*Hour, offset)), "1 day from now"},
		{"2 days from now", Time(time.Unix(now+49*Hour, offset)), "2 days from now"},
		{"3 days from now", Time(time.Unix(now+3*Day, offset)), "3 days from now"},
		{"1 week from now (1)", Time(time.Unix(now+7*Day, offset)), "1 week from now"},
		{"1 week from now (2)", Time(time.Unix(now+12*Day, offset)), "1 week from now"},
		{"2 weeks from now", Time(time.Unix(now+15*Day, offset)), "2 weeks from now"},
		{"1 month from now", Time(time.Unix(now+30*Day, offset)), "1 month from now"},
		{"1 year from now", Time(time.Unix(now+365*Day, offset)), "1 year from now"},
		{"2 years from now", Time(time.Unix(now+2*Year, offset)), "2 years from now"},
		{"a while from now", Time(time.Unix(now+LongTime, offset)), "a long while from now"},
	}.validate(t)
}

func TestRange(t *testing.T) {
	start := time.Time{}
	end := time.Unix(math.MaxInt64, math.MaxInt64)
	x := RelTime(start, end, "ago", "from now")
	if x != "a long while from now" {
		t.Errorf("Expected a long while from now, got %q", x)
	}
}

func TestRelTimeMagnitudes(t *testing.T) {
	magnitudes := []Magnitude{
		NewMagnitude(1*time.Second, "now", time.Second),
		NewMagnitude(2*time.Second, "1s %s", time.Second),
		NewMagnitude(Minute*time.Second, "s %s", time.Second),
		NewMagnitude(2*Minute*time.Second, "1m %s", time.Second),
		NewMagnitude(Hour*time.Second, "%dm %s", Minute*time.Second),
		NewMagnitude(2*Hour*time.Second, "1h %s", time.Second),
		NewMagnitude(Day*time.Second, "%dh %s", Hour*time.Second),
		NewMagnitude(2*Day*time.Second, "1D %s", time.Second),
		NewMagnitude(Month*time.Second, "%dD %s", Day*time.Second),
		NewMagnitude(2*Month*time.Second, "1M %s", time.Second),
		NewMagnitude(Year*time.Second, "%dM %s", Month*time.Second),
		NewMagnitude(18*Month*time.Second, "1Y %s", time.Second),
		NewMagnitude(2*Year*time.Second, "2Y %s", time.Second),
	}
	now := time.Now().Unix()
	timeNow := time.Unix(now, 0)
	testList{
		{"now", RelTimeMagnitudes(time.Unix(now, 0), timeNow, "ago", "later", magnitudes), "now"},
		{"1 second from now", RelTimeMagnitudes(time.Unix(now+1, 1), timeNow, "ago", "later", magnitudes), "1s later"},
		// Unit week has been removed from magnitudes
		{"1 week ago", RelTimeMagnitudes(time.Unix(now-12*Day, 0), timeNow, "ago", "", magnitudes), "12D ago"},
		{"3 months ago", RelTimeMagnitudes(time.Unix(now-99*Day, 0), timeNow, "ago", "later", magnitudes), "3M ago"},
		{"1 year ago", RelTimeMagnitudes(time.Unix(now-365*Day, 0), timeNow, "", "later", magnitudes), "1Y "},
		{"out of defined magnitudes", RelTimeMagnitudes(time.Unix(now+LongTime, 0), timeNow, "ago", "later", magnitudes), "undefined"},
	}.validate(t)
}
