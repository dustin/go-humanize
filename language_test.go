package humanize

import (
	"fmt"
	"time"
)

func EsxampleTurkish() {
	SetLanguage(Turkish)
	fmt.Println(Time(time.Now()))
	SetLanguage(English)
	// Output: şimdi
}
