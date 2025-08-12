package humanize

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
)

type groupingStyle int

const (
	groupWestern groupingStyle = iota
	groupIndian
	groupNone
)

type CurrencyFormat struct {
	Symbol        string
	Name          string
	DecimalPlaces int
	GroupingStyle groupingStyle
}

var currencyFormats = map[string]CurrencyFormat{
	"USD": {"$", "US Dollar", 2, groupWestern},
	"EUR": {"€", "Euro", 2, groupWestern},
	"GBP": {"£", "British Pound", 2, groupWestern},
	"JPY": {"¥", "Japanese Yen", 0, groupWestern},
	"CNY": {"¥", "Chinese Yuan", 2, groupWestern},
	"CHF": {"₣", "Swiss Franc", 2, groupWestern},
	"CAD": {"C$", "Canadian Dollar", 2, groupWestern},
	"AUD": {"A$", "Australian Dollar", 2, groupWestern},
	"NZD": {"NZ$", "New Zealand Dollar", 2, groupWestern},
	"SEK": {"kr", "Swedish Krona", 2, groupWestern},
	"NOK": {"kr", "Norwegian Krone", 2, groupWestern},
	"DKK": {"kr", "Danish Krone", 2, groupWestern},
	"INR": {"₹", "Indian Rupee", 2, groupIndian},
	"KRW": {"₩", "South Korean Won", 0, groupWestern},
	"SGD": {"S$", "Singapore Dollar", 2, groupWestern},
	"HKD": {"HK$", "Hong Kong Dollar", 2, groupWestern},
	"THB": {"฿", "Thai Baht", 2, groupWestern},
	"MYR": {"RM", "Malaysian Ringgit", 2, groupWestern},
	"IDR": {"Rp", "Indonesian Rupiah", 0, groupWestern},
	"PHP": {"₱", "Philippine Peso", 2, groupWestern},
	"VND": {"₫", "Vietnamese Dong", 0, groupWestern},
	"TWD": {"NT$", "Taiwan Dollar", 2, groupWestern},
	"AED": {"د.إ", "UAE Dirham", 2, groupWestern},
	"SAR": {"﷼", "Saudi Riyal", 2, groupWestern},
	"ILS": {"₪", "Israeli Shekel", 2, groupWestern},
	"EGP": {"£", "Egyptian Pound", 2, groupWestern},
	"ZAR": {"R", "South African Rand", 2, groupWestern},
	"NGN": {"₦", "Nigerian Naira", 2, groupWestern},
	"KES": {"KSh", "Kenyan Shilling", 2, groupWestern},
	"TRY": {"₺", "Turkish Lira", 2, groupWestern},
	"MXN": {"$", "Mexican Peso", 2, groupWestern},
	"BRL": {"R$", "Brazilian Real", 2, groupWestern},
	"ARS": {"$", "Argentine Peso", 2, groupWestern},
	"CLP": {"$", "Chilean Peso", 0, groupWestern},
	"COP": {"$", "Colombian Peso", 2, groupWestern},
	"PEN": {"S/", "Peruvian Sol", 2, groupWestern},
	"PLN": {"zł", "Polish Złoty", 2, groupWestern},
	"CZK": {"Kč", "Czech Koruna", 2, groupWestern},
	"HUF": {"Ft", "Hungarian Forint", 0, groupWestern},
	"RON": {"lei", "Romanian Leu", 2, groupWestern},
	"UAH": {"₴", "Ukrainian Hryvnia", 2, groupWestern},
	"BTC": {"₿", "Bitcoin", 8, groupWestern},
	"ETH": {"Ξ", "Ethereum", 18, groupWestern},
	"XAU": {"oz", "Gold Ounce", 4, groupWestern},
	"XAG": {"oz", "Silver Ounce", 4, groupWestern},
}

// Currency formats an amount with the symbol and grouping rules for the given currency code.
func Currency(amount interface{}, code string) string {
	format, ok := currencyFormats[strings.ToUpper(code)]
	if !ok {
		format = CurrencyFormat{Symbol: strings.ToUpper(code), Name: "Unknown Currency", DecimalPlaces: 2, GroupingStyle: groupWestern}
	}
	return formatCurrency(amount, format)
}

// CurrencyWithName formats an amount and appends the currency name.
func CurrencyWithName(amount interface{}, code string) string {
	format, ok := currencyFormats[strings.ToUpper(code)]
	if !ok {
		return Currency(amount, code)
	}
	return fmt.Sprintf("%s (%s)", formatCurrency(amount, format), format.Name)
}

// GetCurrencyInfo returns the format definition for a currency code.
func GetCurrencyInfo(code string) (CurrencyFormat, bool) {
	format, ok := currencyFormats[strings.ToUpper(code)]
	return format, ok
}

// SupportedCurrencies returns all available currency codes.
func SupportedCurrencies() []string {
	codes := make([]string, 0, len(currencyFormats))
	for code := range currencyFormats {
		codes = append(codes, code)
	}
	return codes
}

// IsSupported checks if a currency code is supported.
func IsSupported(code string) bool {
	_, ok := currencyFormats[strings.ToUpper(code)]
	return ok
}

// --- internal helpers ---

func formatCurrency(amount interface{}, format CurrencyFormat) string {
	var v big.Float
	switch val := amount.(type) {
	case int, int8, int16, int32, int64:
		v.SetInt64(toInt64(val))
	case uint, uint8, uint16, uint32, uint64:
		v.SetUint64(toUint64(val))
	case float32:
		v.SetFloat64(float64(val))
	case float64:
		v.SetFloat64(val)
	case *big.Float:
		v.Copy(val)
	case *big.Int:
		v.SetInt(val)
	case string:
		if _, ok := v.SetString(val); !ok {
			return fmt.Sprintf("%s %s", format.Symbol, val)
		}
	default:
		return fmt.Sprintf("%s %v", format.Symbol, amount)
	}

	negative := v.Sign() < 0
	if negative {
		v.Abs(&v)
	}

	formatted := v.Text('f', format.DecimalPlaces)
	parts := strings.Split(formatted, ".")
	intPart := parts[0]
	decPart := ""
	if len(parts) > 1 && parts[1] != "" {
		decPart = "." + parts[1]
	}

	switch format.GroupingStyle {
	case groupIndian:
		intPart = formatIndian(intPart)
	case groupWestern:
		intPart = formatWestern(intPart)
	}

	result := intPart + decPart
	if negative {
		return fmt.Sprintf("%s -%s", format.Symbol, result)
	}
	return fmt.Sprintf("%s %s", format.Symbol, result)
}

func formatWestern(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func formatIndian(s string) string {
	if len(s) <= 3 {
		return s
	}
	last3 := s[len(s)-3:]
	remaining := s[:len(s)-3]
	var parts []string
	for len(remaining) > 2 {
		parts = append([]string{remaining[len(remaining)-2:]}, parts...)
		remaining = remaining[:len(remaining)-2]
	}
	if remaining != "" {
		parts = append([]string{remaining}, parts...)
	}
	if len(parts) > 0 {
		return strings.Join(parts, ",") + "," + last3
	}
	return last3
}

func toInt64(v interface{}) int64 {
	switch t := v.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	}
	return 0
}

func toUint64(v interface{}) uint64 {
	switch t := v.(type) {
	case uint:
		return uint64(t)
	case uint8:
		return uint64(t)
	case uint16:
		return uint64(t)
	case uint32:
		return uint64(t)
	case uint64:
		return t
	}
	return 0
}
