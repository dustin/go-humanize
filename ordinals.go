package humanize

import (
	"strconv"
	"strings"
)

// Ordinal gives you the input number in a rank/ordinal format.
//
// Ordinal(3) -> 3rd
func Ordinal(x int) (out string) {
	ValidateLanguage()

	ordinals := GetRuleset().Ords

	for _, rule := range ordinals {
		out = applyRule(x, out, rule)
	}

	out = strconv.Itoa(x) + out

	return out
}

func applyRule(x int, in string, rule []string) (out string) {
	subRules := strings.SplitAfter(rule[0], "%")
	out = in
	for _, s := range subRules {
		if s == "%" {
			continue
		}
		if s == "." {
			out = rule[1]
		} else {
			if ruleMatches(x, s) {
				out = rule[1]
			}
		}
	}

	return out
}

func ruleMatches(input int, rule string) bool {
	r := strings.Split(rule, " ")
	m1, _ := strconv.Atoi(r[0])
	m2, _ := strconv.Atoi(r[1])

	if input%m1 == m2 {
		return true
	}

	return false
}

/*
suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix

*/
