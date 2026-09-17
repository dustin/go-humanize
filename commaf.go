//go:build go1.6
// +build go1.6

package humanize

import (
	"bytes"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// BigCommaf produces a string form of the given big.Float in base 10
// with commas after every three orders of magnitude.
func BigCommaf(v *big.Float) string {
	buf := &bytes.Buffer{}
	if v.Sign() < 0 {
		buf.Write([]byte{'-'})
		v.Abs(v)
	}

	comma := []byte{','}

	parts := strings.Split(v.Text('f', -1), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

// FormatFloatComma formats a float64 number with thousands separators and
// the specified number of decimal places. It aligns with the behavior
// of strconv.FormatFloat with 'f' format.
//
// Key differences from CommafWithDigits:
//   - FormatFloatComma uses strconv.FormatFloat, which rounds the decimal
//     part to the specified number of digits (half-to-even rounding)
//   - CommafWithDigits simply truncates the decimal part without rounding
//
// Special floating-point values (NaN, +Inf, -Inf) are returned as-is without
// thousands separator processing.
//
// - Integer part gets comma separators every three digits
// - Decimal part retains exactly 'digits' places (with rounding)
// - When digits = 0, no decimal point is output
// - Negative numbers have the minus sign at the front
//
// e.g. FormatFloatComma(1234.5678, 2) -> "1,234.57" (rounded)
//
//	CommafWithDigits(1234.5678, 2)    -> "1,234.56" (truncated)
//
//	FormatFloatComma(1234567.89, 2) -> "1,234,567.89"
//	FormatFloatComma(-1234.5, 0)     -> "-1,234"
//	FormatFloatComma(math.NaN(), 2)   -> "NaN"
//	FormatFloatComma(math.Inf(1), 2)  -> "+Inf"
//	FormatFloatComma(math.Inf(-1), 2) -> "-Inf"
func FormatFloatComma(num float64, digits int) string {
	if math.IsNaN(num) {
		return "NaN"
	}
	if math.IsInf(num, 1) {
		return "+Inf"
	}
	if math.IsInf(num, -1) {
		return "-Inf"
	}

	buf := &bytes.Buffer{}
	if num < 0 {
		buf.Write([]byte{'-'})
		num = 0 - num
	}

	formatted := strconv.FormatFloat(num, 'f', digits, 64)
	parts := strings.Split(formatted, ".")

	comma := []byte{','}

	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}

	return buf.String()
}
