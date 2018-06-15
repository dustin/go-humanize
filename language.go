package humanize

// Local type for the ordinals and times.
type Local string

// Local for constant language values
const (
	English Local = "en_US"
	Turkish Local = "tr_TR"
)

var active = English

// GetLanguage of the humanizing option.
func GetLanguage() Local {
	return active
}

// SetLanguage of the humanizing option.
func SetLanguage(l Local) {
	active = l
}

func ParseRuleset(l Local) {

}
