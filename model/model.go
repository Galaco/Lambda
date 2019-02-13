package model

type Model struct {
	Scene *Scene
	Project *Project
}

func NewModel() *Model {
	return &Model{
		Scene: NewScene(),
		Project: NewProject(),
	}
}
