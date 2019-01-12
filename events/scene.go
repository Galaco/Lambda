package events

import "github.com/galaco/Lambda/core/models"

const TypeNewScene = "NewScene"
const TypeOpenScene = "OpenScene"
const TypeOpenSceneFailed = "OpenSceneFailed"

type NewScene struct {
	scene *models.Vmf
}

func (act *NewScene) Type() string {
	return TypeNewScene
}

func NewNewScene(scene *models.Vmf) *NewScene {
	return &NewScene{
		scene: scene,
	}
}


type OpenScene struct {
	filepath string
}

func (act *OpenScene) Type() string {
	return TypeOpenScene
}

func (act *OpenScene) Path() string {
	return act.filepath
}

func NewOpenScene(filepath string) *OpenScene {
	return &OpenScene{
		filepath: filepath,
	}
}


type OpenSceneFailed struct {
}

func (act *OpenSceneFailed) Type() string {
	return TypeOpenSceneFailed
}

func NewOpenSceneFailed() *OpenSceneFailed {
	return &OpenSceneFailed{}
}