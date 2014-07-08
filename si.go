package humanize

import "math"

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

// ComputeSI finds the most appropriate SI prefix for the given number
// and returns the prefix along with the value adjusted to be within
// that prefix.  e.g. 2.2345e-12 -> (2.2345, "p")
func ComputeSI(input float64) (float64, string) {
	if input == 0 {
		return 0, ""
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
	return value, prefix
}

// SI returns a string with default formatting
// Uses Ftoa to format float value, removing trailing zeros
// e.g. SI(1000000, B) -> 1MB
// e.g. SI(2.2345e-12, "F") -> 2.2345pF
func SI(input float64, unit string) string {
	value, prefix := ComputeSI(input)
	return Ftoa(value) + prefix + unit
}
