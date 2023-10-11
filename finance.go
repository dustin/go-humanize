package humanize

import (
	"fmt"
)

var (
	FinanceSign = "$"
)

func Finance(f float64) string {
	switch n := f; {
	case n >= 1_000_000_000_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("%s%sQ", FinanceSign, s)

	case n >= 1_000_000_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("%s%sT", FinanceSign, s)

	case n >= 1_000_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("%s%sB", FinanceSign, s)

	case n >= 1_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("%s%sM", FinanceSign, s)

	case n >= 1_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("%s%sK", FinanceSign, s)

	case n < 1_000:
		return fmt.Sprintf("%s%.f", FinanceSign, f)

	default:
		return "NaN"
	}
}

func insertAt(index int, elem rune, slice []rune) string {
	copy(slice[index+1:], slice[index:])
	slice[index] = elem
	return string(slice)
}
