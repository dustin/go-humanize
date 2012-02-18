package humanize

import (
	"testing"
)

func TestReverse(t *testing.T) {
	assert(t, "", reverse(""), "")
	assert(t, "1", reverse("1"), "1")
	assert(t, "12", reverse("12"), "21")
	assert(t, "123", reverse("123"), "321")
	assert(t, "1234", reverse("1234"), "4321")
}

func TestCommas(t *testing.T) {
	assert(t, "0", Comma(0), "0")
	assert(t, "10", Comma(10), "10")
	assert(t, "100", Comma(100), "100")
	assert(t, "1,000", Comma(1000), "1,000")
	assert(t, "10,000", Comma(10000), "10,000")
	assert(t, "10,000,000", Comma(10000000), "10,000,000")
	assert(t, "-10,000,000", Comma(-10000000), "-10,000,000")
	assert(t, "-10,000", Comma(-10000), "-10,000")
	assert(t, "-1,000", Comma(-1000), "-1,000")
	assert(t, "-100", Comma(-100), "-100")
	assert(t, "-10", Comma(-10), "-10")
}
