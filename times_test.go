package humanize

import (
	"testing"
	"time"
)

func checkTime(t *testing.T, expected, got string) {
	if got != expected {
		t.Fatalf("Expected %s, got %s", expected, got)
	}
}

func TestPast(t *testing.T) {

	expected := []string{
		"now",
		"12 seconds ago",
		"30 seconds ago",
		"45 seconds ago",
		"15 minutes ago",
		"2 hours ago",
		"21 hours ago",
		"1 day ago",
		"2 days ago",
		"3 days ago",
		"1 week ago",
		"1 week ago",
		"2 weeks ago",
		"1 month ago",
		"1 year ago",
	}

	i := 0
	now := time.Now().Unix()

	checkTime(t, expected[i], Humanize(time.Unix(now, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-12, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-30, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-45, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-15*Minute, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-2*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-21*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-26*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-49*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-3*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-7*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-12*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-15*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-39*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now-365*Day, 0)))
}

func TestFuture(t *testing.T) {

	expected := []string{
		"now",
		"12 seconds from now",
		"30 seconds from now",
		"45 seconds from now",
		"15 minutes from now",
		"2 hours from now",
		"21 hours from now",
		"1 day from now",
		"2 days from now",
		"3 days from now",
		"1 week from now",
		"1 week from now",
		"2 weeks from now",
		"1 month from now",
		"1 year from now",
	}

	i := 0
	now := time.Now().Unix()

	checkTime(t, expected[i], Humanize(time.Unix(now, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+12, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+30, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+45, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+15*Minute, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+2*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+21*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+26*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+49*Hour, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+3*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+7*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+12*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+15*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+39*Day, 0)))
	i++
	checkTime(t, expected[i], Humanize(time.Unix(now+365*Day, 0)))
}
