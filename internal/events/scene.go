package events

import "github.com/galaco/Lambda/internal/model/valve"

const (
	// TypeNewScene event type
	TypeNewScene = "NewScene"
	// TypeOpenScene event type
	TypeOpenScene = "OpenScene"
	// TypeOpenSceneFailed event type
	TypeOpenSceneFailed = "OpenSceneFailed"
	// TypeEntityNodeSelected event type
	TypeEntityNodeSelected = "EntityNodeSelected"
	// TypeSolidNodeSelected event type
	TypeSolidNodeSelected = "SolidNodeSelected"
	// TypeSceneClosed event type
	TypeSceneClosed = "SceneClosed"
)

// NewScene action includes a payload of the new scene
// Note that NewScene is also a loaded scene, as there is
// no actual distinction between an empty and a populated scene
// for the purpose of a whole scene load.
type NewScene struct {
	scene *valve.Vmf
}

func (act *NewScene) Type() string {
	return TypeNewScene
}

func (act *NewScene) Model() *valve.Vmf {
	return act.scene
}

func NewNewScene(scene *valve.Vmf) *NewScene {
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

// EntityNodeSelected provides the unique id of the selected
// object in the scene.
type EntityNodeSelected struct {
	Id int
}

func (act *EntityNodeSelected) Type() string {
	return TypeEntityNodeSelected
}

func NewEntityNodeSelected(id int) *EntityNodeSelected {
	return &EntityNodeSelected{
		Id: id,
	}
}

type SolidNodeSelected struct {
	Id int
}

func (act *SolidNodeSelected) Type() string {
	return TypeSolidNodeSelected
}

func NewSolidNodeSelected(id int) *SolidNodeSelected {
	return &SolidNodeSelected{
		Id: id,
	}
}

type SceneClosed struct {
}

func (act *SceneClosed) Type() string {
	return TypeSceneClosed
}

func NewSceneClosed() *SceneClosed {
	return &SceneClosed{}
}
