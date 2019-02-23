package events

const TypePreferencesUpdated = "TypePreferencesUpdated"

type PreferencesUpdated struct {
	General    struct{}
	Appearance struct {
		Theme int
	}
}

func (act *PreferencesUpdated) Type() string {
	return TypePreferencesUpdated
}

func NewPreferencesUpdated() *PreferencesUpdated {
	return &PreferencesUpdated{}
}
