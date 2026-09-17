package humanize

import (
	"math"
	"math/big"
	"testing"
)

func TestCommas(t *testing.T) {
	testList{
		{"0", Comma(0), "0"},
		{"10", Comma(10), "10"},
		{"100", Comma(100), "100"},
		{"1,000", Comma(1000), "1,000"},
		{"10,000", Comma(10000), "10,000"},
		{"100,000", Comma(100000), "100,000"},
		{"10,000,000", Comma(10000000), "10,000,000"},
		{"10,100,000", Comma(10100000), "10,100,000"},
		{"10,010,000", Comma(10010000), "10,010,000"},
		{"10,001,000", Comma(10001000), "10,001,000"},
		{"123,456,789", Comma(123456789), "123,456,789"},
		{"maxint", Comma(9.223372e+18), "9,223,372,000,000,000,000"},
		{"math.maxint", Comma(math.MaxInt64), "9,223,372,036,854,775,807"},
		{"math.minint", Comma(math.MinInt64), "-9,223,372,036,854,775,808"},
		{"minint", Comma(-9.223372e+18), "-9,223,372,000,000,000,000"},
		{"-123,456,789", Comma(-123456789), "-123,456,789"},
		{"-10,100,000", Comma(-10100000), "-10,100,000"},
		{"-10,010,000", Comma(-10010000), "-10,010,000"},
		{"-10,001,000", Comma(-10001000), "-10,001,000"},
		{"-10,000,000", Comma(-10000000), "-10,000,000"},
		{"-100,000", Comma(-100000), "-100,000"},
		{"-10,000", Comma(-10000), "-10,000"},
		{"-1,000", Comma(-1000), "-1,000"},
		{"-100", Comma(-100), "-100"},
		{"-10", Comma(-10), "-10"},
	}.validate(t)
}

func TestCommafWithDigits(t *testing.T) {
	testList{
		{"1.23, 0", CommafWithDigits(1.23, 0), "1"},
		{"1.23, 1", CommafWithDigits(1.23, 1), "1.2"},
		{"1.23, 2", CommafWithDigits(1.23, 2), "1.23"},
		{"1.23, 3", CommafWithDigits(1.23, 3), "1.23"},
	}.validate(t)
}

func TestCommafs(t *testing.T) {
	testList{
		{"0", Commaf(0), "0"},
		{"10.11", Commaf(10.11), "10.11"},
		{"100", Commaf(100), "100"},
		{"1,000", Commaf(1000), "1,000"},
		{"10,000", Commaf(10000), "10,000"},
		{"100,000", Commaf(100000), "100,000"},
		{"834,142.32", Commaf(834142.32), "834,142.32"},
		{"10,000,000", Commaf(10000000), "10,000,000"},
		{"10,100,000", Commaf(10100000), "10,100,000"},
		{"10,010,000", Commaf(10010000), "10,010,000"},
		{"10,001,000", Commaf(10001000), "10,001,000"},
		{"123,456,789", Commaf(123456789), "123,456,789"},
		{"maxf64", Commaf(math.MaxFloat64), "179,769,313,486,231,570,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000,000"},
		{"minf64", Commaf(math.SmallestNonzeroFloat64), "0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005"},
		{"-123,456,789", Commaf(-123456789), "-123,456,789"},
		{"-10,100,000", Commaf(-10100000), "-10,100,000"},
		{"-10,010,000", Commaf(-10010000), "-10,010,000"},
		{"-10,001,000", Commaf(-10001000), "-10,001,000"},
		{"-10,000,000", Commaf(-10000000), "-10,000,000"},
		{"-100,000", Commaf(-100000), "-100,000"},
		{"-10,000", Commaf(-10000), "-10,000"},
		{"-1,000", Commaf(-1000), "-1,000"},
		{"-100.11", Commaf(-100.11), "-100.11"},
		{"-10", Commaf(-10), "-10"},
	}.validate(t)
}

func BenchmarkCommas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma(1234567890)
	}
}

func BenchmarkCommaf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Commaf(1234567890.83584)
	}
}

func BenchmarkBigCommas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigComma(big.NewInt(1234567890))
	}
}

func bigComma(i int64) string {
	return BigComma(big.NewInt(i))
}

func TestBigCommas(t *testing.T) {
	testList{
		{"0", bigComma(0), "0"},
		{"10", bigComma(10), "10"},
		{"100", bigComma(100), "100"},
		{"1,000", bigComma(1000), "1,000"},
		{"10,000", bigComma(10000), "10,000"},
		{"100,000", bigComma(100000), "100,000"},
		{"10,000,000", bigComma(10000000), "10,000,000"},
		{"10,100,000", bigComma(10100000), "10,100,000"},
		{"10,010,000", bigComma(10010000), "10,010,000"},
		{"10,001,000", bigComma(10001000), "10,001,000"},
		{"123,456,789", bigComma(123456789), "123,456,789"},
		{"maxint", bigComma(9.223372e+18), "9,223,372,000,000,000,000"},
		{"minint", bigComma(-9.223372e+18), "-9,223,372,000,000,000,000"},
		{"-123,456,789", bigComma(-123456789), "-123,456,789"},
		{"-10,100,000", bigComma(-10100000), "-10,100,000"},
		{"-10,010,000", bigComma(-10010000), "-10,010,000"},
		{"-10,001,000", bigComma(-10001000), "-10,001,000"},
		{"-10,000,000", bigComma(-10000000), "-10,000,000"},
		{"-100,000", bigComma(-100000), "-100,000"},
		{"-10,000", bigComma(-10000), "-10,000"},
		{"-1,000", bigComma(-1000), "-1,000"},
		{"-100", bigComma(-100), "-100"},
		{"-10", bigComma(-10), "-10"},
	}.validate(t)
}

func TestVeryBigCommas(t *testing.T) {
	tests := []struct{ in, exp string }{
		{
			"84889279597249724975972597249849757294578485",
			"84,889,279,597,249,724,975,972,597,249,849,757,294,578,485",
		},
		{
			"-84889279597249724975972597249849757294578485",
			"-84,889,279,597,249,724,975,972,597,249,849,757,294,578,485",
		},
	}
	for _, test := range tests {
		n, _ := (&big.Int{}).SetString(test.in, 10)
		got := BigComma(n)
		if test.exp != got {
			t.Errorf("Expected %q, got %q", test.exp, got)
		}
	}
}

func TestHumanizeBigIntMutation(t *testing.T) {
	value := big.NewInt(1000000)
	value = value.Mul(value, value)
	expected := BigComma(value)
	actual := BigComma(value)
	if expected != actual {
		t.Log(expected, " != ", actual)
		t.Fail()
	}
}

func TestParseComma(t *testing.T) {
	tests := []struct {
		in  string
		exp int64
	}{
		{"0", 0},
		{"10", 10},
		{"1,000", 1000},
		{"123,456,789", 123456789},
		{"-10", -10},
		{"-1,000", -1000},
		{"-123,456,789", -123456789},
		{"-9,223,372,036,854,775,808", math.MinInt64},
		{"9,223,372,036,854,775,807", math.MaxInt64},
	}

	for _, p := range tests {
		got, err := ParseComma(p.in)
		if err != nil {
			t.Errorf("Couldn't parse %v: %v", p.in, err)
		}
		if got != p.exp {
			t.Errorf("Expected %v for %v, got %v", p.exp, p.in, got)
		}
	}

	roundTripTests := []int64{
		0,
		123456789,
		-123456789,
		math.MaxInt64,
		math.MinInt64,
	}

	for _, n := range roundTripTests {
		got, err := ParseComma(Comma(n))
		if err != nil {
			t.Errorf("Round-trip failed for %v: %v", n, err)
		}
		if got != n {
			t.Errorf("Round-trip failed: expected %v, got %v", n, got)
		}
	}
}

func TestParseCommaErrors(t *testing.T) {
	errorTests := []string{
		"",
		"abc",
		"123a456",
		"1.5",
	}

	for _, s := range errorTests {
		got, err := ParseComma(s)
		if err == nil {
			t.Errorf("Expected error for %q, got %v", s, got)
		}
		if got != 0 {
			t.Errorf("Expected 0 on error for %q, got %v", s, got)
		}
	}
}

func TestParseCommaf(t *testing.T) {
	tests := []struct {
		in  string
		exp float64
	}{
		{"0", 0},
		{"10", 10},
		{"1,000", 1000},
		{"123,456,789", 123456789},
		{"10.11", 10.11},
		{"834,142.32", 834142.32},
		{"-10", -10},
		{"-1,000", -1000},
		{"-100.11", -100.11},
		{"-123,456,789.123", -123456789.123},
	}

	for _, p := range tests {
		got, err := ParseCommaf(p.in)
		if err != nil {
			t.Errorf("Couldn't parse %v: %v", p.in, err)
		}
		if math.Abs(got-p.exp) > math.Abs(p.exp)*1e-9 {
			t.Errorf("Expected %v for %v, got %v", p.exp, p.in, got)
		}
	}

	roundTripTests := []float64{
		0,
		123456789,
		123456789.123,
		-123456789,
		-123456789.123,
	}

	for _, f := range roundTripTests {
		got, err := ParseCommaf(Commaf(f))
		if err != nil {
			t.Errorf("Round-trip failed for %v: %v", f, err)
		}
		if math.Abs(got-f) > math.Abs(f)*1e-9 {
			t.Errorf("Round-trip failed: expected %v, got %v", f, got)
		}
	}
}

func TestParseCommafErrors(t *testing.T) {
	errorTests := []string{
		"",
		"abc",
		"123a456",
		"1.2.3",
	}

	for _, s := range errorTests {
		got, err := ParseCommaf(s)
		if err == nil {
			t.Errorf("Expected error for %q, got %v", s, got)
		}
		if got != 0 {
			t.Errorf("Expected 0 on error for %q, got %v", s, got)
		}
	}
}
