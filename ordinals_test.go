package humanize

import (
	"testing"
)

func TestOrdinals(t *testing.T) {
	assert(t, "0", Ordinal(0), "0th")
	assert(t, "1", Ordinal(1), "1st")
	assert(t, "2", Ordinal(2), "2nd")
	assert(t, "3", Ordinal(3), "3rd")
	assert(t, "4", Ordinal(4), "4th")
	assert(t, "10", Ordinal(10), "10th")
	assert(t, "11", Ordinal(11), "11th")
	assert(t, "12", Ordinal(12), "12th")
	assert(t, "13", Ordinal(13), "13th")
	assert(t, "101", Ordinal(101), "101st")
	assert(t, "102", Ordinal(102), "102nd")
	assert(t, "103", Ordinal(103), "103rd")
}
