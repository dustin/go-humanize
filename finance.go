package humanize

import (
	"fmt"
)

func Finance(f float64) string {
	switch n := f; {
	case n >= 1_000_000_000_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("$%sQ", s)

	case n >= 1_000_000_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("$%sT", s)

	case n >= 1_000_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("$%sB", s)

	case n >= 1_000_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("$%sM", s)

	case n >= 1_000:
		s := fmt.Sprintf("%.f", f)
		s = insertAt(1, '.', []rune(s[:4]))
		return fmt.Sprintf("$%sK", s)

	case n < 1_000:
		return fmt.Sprintf("$%.f", f)

	default:
		return "NaN"
	}
}

func insertAt(index int, elem rune, slice []rune) string {
	copy(slice[index+1:], slice[index:])
	slice[index] = elem
	return string(slice)
}
