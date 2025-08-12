package humanize

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCurrency(t *testing.T) {
	tests := []struct {
		name   string
		amount interface{}
		code   string
		want   string
	}{
		// Basic Western formatting
		{"USD basic", 60000000, "USD", "$ 60,000,000.00"},
		{"GBP basic", 60000000, "GBP", "£ 60,000,000.00"},
		{"EUR basic", 1234567.89, "EUR", "€ 1,234,567.89"},
		{"CAD basic", 12345, "CAD", "C$ 12,345.00"},

		// Zero decimal currencies
		{"JPY no decimals", 60000000, "JPY", "¥ 60,000,000"},
		{"KRW no decimals", 1000000, "KRW", "₩ 1,000,000"},
		{"IDR no decimals", 15000000, "IDR", "Rp 15,000,000"},

		// Indian numbering system
		{"INR small", 12345, "INR", "₹ 12,345.00"},
		{"INR medium", 60000000, "INR", "₹ 6,00,00,000.00"},
		{"INR large", 987654321, "INR", "₹ 98,76,54,321.00"},
		{"INR very large", 1234567890, "INR", "₹ 1,23,45,67,890.00"},

		// Negative values
		{"USD negative", -1234567, "USD", "$ -1,234,567.00"},
		{"INR negative", -987654321, "INR", "₹ -98,76,54,321.00"},
		{"JPY negative", -500000, "JPY", "¥ -500,000"},

		// Decimal handling
		{"USD with decimals", 1234567.89, "USD", "$ 1,234,567.89"},
		{"GBP with decimals", 999.99, "GBP", "£ 999.99"},
		{"CHF with decimals", 12.34, "CHF", "₣ 12.34"},

		// Cryptocurrency (high precision)
		{"BTC", 1.23456789, "BTC", "₿ 1.23456789"},
		{"ETH", 1.5, "ETH", "Ξ 1.500000000000000000"}, // Simple value that float64 can represent exactly

		// Different numeric types
		{"int8", int8(123), "USD", "$ 123.00"},
		{"int16", int16(12345), "USD", "$ 12,345.00"},
		{"int32", int32(1234567), "USD", "$ 1,234,567.00"},
		{"uint", uint(12345), "USD", "$ 12,345.00"},
		{"uint64", uint64(123456789), "USD", "$ 123,456,789.00"},
		{"float32", float32(123.45), "USD", "$ 123.45"},

		// Edge cases
		{"zero", 0, "USD", "$ 0.00"},
		{"small decimal", 0.01, "USD", "$ 0.01"},
		{"large number", 999999999999, "USD", "$ 999,999,999,999.00"},

		// Fallback for unknown currency
		{"unknown currency", 12345, "XYZ", "XYZ 12,345.00"},
		{"empty code", 12345, "", " 12,345.00"},

		// Case insensitive
		{"lowercase code", 12345, "usd", "$ 12,345.00"},
		{"mixed case code", 12345, "GbP", "£ 12,345.00"},

		// Various currencies
		{"CNY", 100000, "CNY", "¥ 100,000.00"},
		{"SGD", 75000, "SGD", "S$ 75,000.00"},
		{"AUD", 50000, "AUD", "A$ 50,000.00"},
		{"BRL", 200000, "BRL", "R$ 200,000.00"},
		{"UAH", 3000000, "UAH", "₴ 3,000,000.00"},
		{"TRY", 850000, "TRY", "₺ 850,000.00"},
		{"ZAR", 180000, "ZAR", "R 180,000.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Currency(tt.amount, tt.code)
			if got != tt.want {
				t.Errorf("Currency(%v, %q) = %q; want %q", tt.amount, tt.code, got, tt.want)
			}
		})
	}
}

func TestCurrencyWithBigNumbers(t *testing.T) {
	tests := []struct {
		name   string
		amount interface{}
		code   string
		want   string
	}{
		{
			"big.Float",
			func() *big.Float {
				f := new(big.Float)
				f.SetString("123456789.123456789")
				return f
			}(),
			"USD",
			"$ 123,456,789.12",
		},
		{
			"big.Int",
			func() *big.Int {
				i := new(big.Int)
				i.SetString("9876543210987654321", 10)
				return i
			}(),
			"USD",
			"$ 9,876,543,210,987,654,321.00",
		},
		{
			"string number",
			"12345678901234567890",
			"BTC",
			"₿ 12,345,678,901,234,567,890.00000000",
		},
		{
			"invalid string",
			"not-a-number",
			"USD",
			"$ not-a-number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Currency(tt.amount, tt.code)
			if got != tt.want {
				t.Errorf("Currency(%v, %q) = %q; want %q", tt.amount, tt.code, got, tt.want)
			}
		})
	}
}

func TestCurrencyWithName(t *testing.T) {
	tests := []struct {
		name   string
		amount interface{}
		code   string
		want   string
	}{
		{"USD with name", 12345, "USD", "$ 12,345.00 (US Dollar)"},
		{"EUR with name", 98765, "EUR", "€ 98,765.00 (Euro)"},
		{"JPY with name", 1000000, "JPY", "¥ 1,000,000 (Japanese Yen)"},
		{"INR with name", 500000, "INR", "₹ 5,00,000.00 (Indian Rupee)"},
		{"unknown currency", 12345, "XYZ", "XYZ 12,345.00"}, // fallback to regular Currency
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CurrencyWithName(tt.amount, tt.code)
			if got != tt.want {
				t.Errorf("CurrencyWithName(%v, %q) = %q; want %q", tt.amount, tt.code, got, tt.want)
			}
		})
	}
}

func TestFormatWestern(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"single digit", "5", "5"},
		{"three digits", "123", "123"},
		{"four digits", "1234", "1,234"},
		{"seven digits", "1234567", "1,234,567"},
		{"ten digits", "1234567890", "1,234,567,890"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatWestern(tt.s)
			if got != tt.want {
				t.Errorf("formatWestern(%q) = %q; want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestFormatIndian(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"single digit", "5", "5"},
		{"three digits", "123", "123"},
		{"four digits", "1234", "1,234"},
		{"five digits", "12345", "12,345"},
		{"six digits", "123456", "1,23,456"},
		{"seven digits", "1234567", "12,34,567"},
		{"eight digits", "12345678", "1,23,45,678"},
		{"nine digits", "123456789", "12,34,56,789"},
		{"ten digits", "1234567890", "1,23,45,67,890"},
		{"very large", "123456789012", "1,23,45,67,89,012"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatIndian(tt.s)
			if got != tt.want {
				t.Errorf("formatIndian(%q) = %q; want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestGetCurrencyInfo(t *testing.T) {
	tests := []struct {
		name       string
		code       string
		wantExists bool
		wantSymbol string
		wantName   string
	}{
		{"USD exists", "USD", true, "$", "US Dollar"},
		{"eur lowercase", "eur", true, "€", "Euro"},
		{"JPY exists", "JPY", true, "¥", "Japanese Yen"},
		{"INR exists", "INR", true, "₹", "Indian Rupee"},
		{"unknown currency", "XYZ", false, "", ""},
		{"empty code", "", false, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			format, exists := GetCurrencyInfo(tt.code)
			if exists != tt.wantExists {
				t.Errorf("GetCurrencyInfo(%q) exists = %v; want %v", tt.code, exists, tt.wantExists)
			}
			if exists && format.Symbol != tt.wantSymbol {
				t.Errorf("GetCurrencyInfo(%q) symbol = %q; want %q", tt.code, format.Symbol, tt.wantSymbol)
			}
			if exists && format.Name != tt.wantName {
				t.Errorf("GetCurrencyInfo(%q) name = %q; want %q", tt.code, format.Name, tt.wantName)
			}
		})
	}
}

func TestIsSupported(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{"USD supported", "USD", true},
		{"eur lowercase", "eur", true},
		{"JPY supported", "JPY", true},
		{"unknown currency", "XYZ", false},
		{"empty code", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSupported(tt.code)
			if got != tt.want {
				t.Errorf("IsSupported(%q) = %v; want %v", tt.code, got, tt.want)
			}
		})
	}
}

func TestSupportedCurrencies(t *testing.T) {
	currencies := SupportedCurrencies()

	// Check that we have a reasonable number of currencies
	if len(currencies) < 40 {
		t.Errorf("SupportedCurrencies() returned %d currencies; want at least 40", len(currencies))
	}

	// Check that major currencies are included
	majorCurrencies := []string{"USD", "EUR", "GBP", "JPY", "CNY", "INR"}
	currencySet := make(map[string]bool)
	for _, c := range currencies {
		currencySet[c] = true
	}

	for _, major := range majorCurrencies {
		if !currencySet[major] {
			t.Errorf("SupportedCurrencies() missing major currency: %s", major)
		}
	}

	// Check for duplicates
	if len(currencies) != len(currencySet) {
		t.Error("SupportedCurrencies() returned duplicate currencies")
	}
}

func TestGroupingStyles(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		code   string
		want   string
	}{
		{"Western style", 1234567, "USD", "$ 1,234,567.00"},
		{"Indian style", 1234567, "INR", "₹ 12,34,567.00"},
		// Test a currency with no grouping if we had one
		// {"No grouping", 1234567, "NONE", "NONE 1234567.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Currency(tt.amount, tt.code)
			if got != tt.want {
				t.Errorf("Currency(%v, %q) = %q; want %q", tt.amount, tt.code, got, tt.want)
			}
		})
	}
}

func TestDecimalPrecision(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		code   string
		want   string
	}{
		{"USD 2 decimals", 123.456, "USD", "$ 123.46"},
		{"JPY 0 decimals", 123.456, "JPY", "¥ 123"},
		{"BTC 8 decimals", 1.12345678, "BTC", "₿ 1.12345678"},
		{"ETH 18 decimals", 2.5, "ETH", "Ξ 2.500000000000000000"}, // Simple value for exact representation
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Currency(tt.amount, tt.code)
			if got != tt.want {
				t.Errorf("Currency(%v, %q) = %q; want %q", tt.amount, tt.code, got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkCurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Currency(1234567.89, "USD")
	}
}

func BenchmarkCurrencyIndian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Currency(1234567890, "INR")
	}
}

func BenchmarkFormatWestern(b *testing.B) {
	s := "1234567890"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatWestern(s)
	}
}

func BenchmarkFormatIndian(b *testing.B) {
	s := "1234567890"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatIndian(s)
	}
}

// Test helper functions
func TestCurrencyFormats(t *testing.T) {
	// Test that all currency formats are valid
	for code, format := range currencyFormats {
		if format.Symbol == "" {
			t.Errorf("Currency %s has empty symbol", code)
		}
		if format.Name == "" {
			t.Errorf("Currency %s has empty name", code)
		}
		if format.DecimalPlaces < 0 || format.DecimalPlaces > 18 {
			t.Errorf("Currency %s has invalid decimal places: %d", code, format.DecimalPlaces)
		}
		if format.GroupingStyle < 0 || format.GroupingStyle > 2 {
			t.Errorf("Currency %s has invalid grouping style: %d", code, format.GroupingStyle)
		}
	}
}

// Example usage tests
func ExampleCurrency() {
	// Basic usage
	result := Currency(1234567.89, "USD")
	fmt.Println(result)
	// Output: $ 1,234,567.89
}

func ExampleCurrency_indian() {
	// Indian rupee formatting
	result := Currency(12345678, "INR")
	fmt.Println(result)
	// Output: ₹ 1,23,45,678.00
}

func ExampleCurrencyWithName() {
	// Currency with full name
	result := CurrencyWithName(50000, "EUR")
	fmt.Println(result)
	// Output: € 50,000.00 (Euro)
}
