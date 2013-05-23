package humanize

import (
	"testing"
)

func TestCommas(t *testing.T) {
	assert(t, "0", Comma(0), "0")
	assert(t, "10", Comma(10), "10")
	assert(t, "100", Comma(100), "100")
	assert(t, "1,000", Comma(1000), "1,000")
	assert(t, "10,000", Comma(10000), "10,000")
	assert(t, "100,000", Comma(100000), "100,000")
	assert(t, "10,000,000", Comma(10000000), "10,000,000")
	assert(t, "10,100,000", Comma(10100000), "10,100,000")
	assert(t, "10,010,000", Comma(10010000), "10,010,000")
	assert(t, "10,001,000", Comma(10001000), "10,001,000")
	assert(t, "123,456,789", Comma(123456789), "123,456,789")
	assert(t, "-123,456,789", Comma(-123456789), "-123,456,789")
	assert(t, "-10,100,000", Comma(-10100000), "-10,100,000")
	assert(t, "-10,010,000", Comma(-10010000), "-10,010,000")
	assert(t, "-10,001,000", Comma(-10001000), "-10,001,000")
	assert(t, "-10,000,000", Comma(-10000000), "-10,000,000")
	assert(t, "-100,000", Comma(-100000), "-100,000")
	assert(t, "-10,000", Comma(-10000), "-10,000")
	assert(t, "-1,000", Comma(-1000), "-1,000")
	assert(t, "-100", Comma(-100), "-100")
	assert(t, "-10", Comma(-10), "-10")
}

func BenchmarkCommas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma(1000000000)
		Comma(1234567890)
	}
}
