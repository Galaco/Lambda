package events

import "github.com/galaco/Lambda/models"

const TypeNewScene = "NewScene"
const TypeOpenScene = "OpenScene"
const TypeOpenSceneFailed = "OpenSceneFailed"
const TypeSceneNodeSelected = "SceneNodeSelected"

// NewScene action includes a payload of the new scene
// Note that NewScene is also a loaded scene, as there is
// no actual distinction between an empty and a populated scene
// for the purpose of a whole scene load.
type NewScene struct {
	scene *models.Vmf
}

func (act *NewScene) Type() string {
	return TypeNewScene
}

func (act *NewScene) Model() *models.Vmf {
	return act.scene
}

func NewNewScene(scene *models.Vmf) *NewScene {
	return &NewScene{
		scene: scene,
	}
}

// OpenScene action includes the filepath of a requested scene
// to open.
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

// OpenSceneFailed denotes that there was an attempt
// to load a scene, but that some issue occurred.
type OpenSceneFailed struct {
}

func (act *OpenSceneFailed) Type() string {
	return TypeOpenSceneFailed
}

func NewOpenSceneFailed() *OpenSceneFailed {
	return &OpenSceneFailed{}
}

// SceneNodeSelected provides the unique id of the selected
// object in the scene.
type SceneNodeSelected struct {
	Id int
}

func (act *SceneNodeSelected) Type() string {
	return TypeSceneNodeSelected
}

func NewSceneNodeSelected(id int) *SceneNodeSelected {
	return &SceneNodeSelected{
		Id: id,
	}
}
