package model

type Model struct {
	Project     *Project
	Preferences *Preferences
}

func NewModel() *Model {
	return &Model{}
}
