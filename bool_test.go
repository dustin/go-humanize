package humanize

import "testing"

func TestBool(t *testing.T) {
	testList{
		{"yes", Bool(true), "yes"},
		{"no", Bool(false), "no"},
	}.validate(t)
}
