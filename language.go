package humanize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Local type for the ordinals and times.
type Local string

// Ruleset for accessing rules
type Ruleset struct {
	Mags magnitudes `json:"magnitudes"`
	Inds indicators `json:"indicators"`
	Ords [][]string `json:"ordinals"`
}

type magnitudes struct {
	Now    string
	Second string
	Minute string
	Hour   string
	Day    string
	Week   string
	Month  string
	Year   string

	Seconds string
	Minutes string
	Hours   string
	Days    string
	Weeks   string
	Months  string
	Years   string

	Longtime string
}

type indicator struct {
	Word string
	Fix  string
}

type indicators struct {
	Before indicator
	Later  indicator
}

// Local for constant language values
const (
	English       Local = "en_US"
	Turkish       Local = "tr_TR"
	Uninitialized Local = ""
)

var active = Uninitialized
var ruleset = Ruleset{}

// GetLanguage of the humanizing option.
func GetLanguage() Local {
	return active
}

// GetRuleset returns current ruleset option
func GetRuleset() Ruleset {
	return ruleset
}

// SetLanguage of the humanizing option.
func SetLanguage(l Local) {
	active = l
	parseRuleset(l)
	UpdateMagnitudes()
}

func parseRuleset(l Local) {
	fmt.Println("Reading", "locals/"+string(l)+".json")
	f, err := ioutil.ReadFile("locals/" + string(l) + ".json")
	if err == nil {
		ruleset = Ruleset{}
		err := json.Unmarshal(f, &ruleset)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error ! Can not read file.")
	}
}
