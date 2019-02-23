package model

import "github.com/galaco/Lambda/pkg/valve"

type Project struct {
	Filename string
	Vmf      *valve.Vmf
}

func NewProject() *Project {
	return &Project{
		Vmf: &valve.Vmf{},
	}
}
