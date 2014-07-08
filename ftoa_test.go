package humanize

import (
	"testing"
)

func TestFtoa(t *testing.T) {
	testList{
		{"200", Ftoa(200), "200"},
		{"2", Ftoa(2), "2"},
		{"2.2", Ftoa(2.2), "2.2"},
		{"2.02", Ftoa(2.02), "2.02"},
		{"200.02", Ftoa(200.02), "200.02"},
	}.validate(t)
}
