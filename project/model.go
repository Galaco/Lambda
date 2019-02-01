package project

type Model struct {
	scene *Scene
}

func (mod *Model) Scene() *Scene {
	return mod.scene
}

func NewModel() *Model {
	return &Model{
		scene: NewScene(),
	}
}
