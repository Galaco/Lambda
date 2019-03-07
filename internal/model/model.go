package model

type Model struct {
	Project     *Project
	Preferences *Preferences
	Logs        *Log
}

func NewModel() *Model {
	return &Model{
		Logs: &Log{},
	}
}
