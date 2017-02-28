package humanize

import (
	"fmt"
	"time"
)

func Duration(d time.Duration) string {
	switch {
	case d == 0:
		return "0 nanoseconds"
	case d == 1:
		return "1 nanosecond"
	case d < time.Microsecond:
		return fmt.Sprintf("%d nanoseconds", d.Nanoseconds())

	case d == time.Microsecond:
		return "1 microsecond"
	case d < time.Millisecond:
		return fmt.Sprintf("%d microseconds", d/time.Microsecond)

	case d == time.Millisecond:
		return "1 millisecond"
	case d < time.Second:
		return fmt.Sprintf("%d milliseconds", d/time.Millisecond)

	case d == time.Second:
		return "1 second"
	case d < time.Minute:
		return fmt.Sprintf("%d seconds", d/time.Second)

	case d == time.Minute:
		return "1 minute"
	case d < time.Hour:
		return fmt.Sprintf("%d minutes", d/time.Minute)

	case d == time.Hour:
		return "1 hour"
	default:
		return fmt.Sprintf("%d hours", d/time.Hour)
	}
}
