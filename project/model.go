package project

import "github.com/galaco/Lambda/valve"

type Model struct {
	Filename string
	Vmf *valve.Vmf
}

func NewModel() *Model {
	return &Model{
		Vmf: &valve.Vmf{},
	}
}
