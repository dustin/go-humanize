package humanize

import (
	"fmt"
	"math"
)

var siPrefixTable = map[float64]string{
	-24: "y", // yocto
	-21: "z", // zepto
	-18: "a", // atto
	-15: "f", // femto
	-12: "p", // pico
	-9:  "n", // nano
	-6:  "µ", // micro
	-3:  "m", // milli
	0:   "",
	3:   "k", // kilo
	6:   "M", // mega
	9:   "G", // giga
	12:  "T", // tera
	15:  "P", // peta
	18:  "E", // exa
	21:  "Z", // zetta
	24:  "Y", // yotta
}

// siNumber holds computed Value and Prefix.
// can be used with fmt.Sprintf for arbitrary formatting.
type siNumber struct {
	Value  float64
	Prefix string
}

// NewSI returns a {Value, Prefix} struct for further formatting
func newSI(input float64) siNumber {
	if input == 0 {
		return siNumber{0, ""}
	}
	exponent := math.Floor(logn(input, 10))
	exponent = math.Floor(exponent/3) * 3

	value := input / math.Pow(10, exponent)

	// Handle special case where value is exactly 1000.0
	// Should return 1M instead of 1000k
	if value == 1000.0 {
		exponent += 3
		value = input / math.Pow(10, exponent)
	}

	prefix := siPrefixTable[exponent]
	return siNumber{value, prefix}
}

// SI returns a string with default formatting
// Uses Ftoa to format float value, removing trailing zeros
// e.g. SI(1000000, B) -> 1MB
// e.g. SI(2.2345e-12, "F") -> 2.2345pF
func SI(input float64, unit string) string {
	si := newSI(input)
	return fmt.Sprintf("%s%s%s", Ftoa(si.Value), si.Prefix, unit)
}
