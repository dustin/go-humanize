package humanize

// Bool converts the gotype boolean into a human readable string.
//
// true becomes "yes" and false becomes "no".
func Bool(boolean bool) string {
	if boolean {
		return "yes"
	}

	return "no"
}
