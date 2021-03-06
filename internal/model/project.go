package model

import "github.com/galaco/Lambda/internal/model/valve"

type Project struct {
	Filename string
	Vmf      *valve.Vmf
}

func NewProject() *Project {
	return &Project{
		Vmf: &valve.Vmf{},
	}
}
