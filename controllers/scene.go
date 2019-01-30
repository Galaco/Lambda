package controllers

import (
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/mvc/event"
	"github.com/galaco/Lambda/services/importers"
	"github.com/galaco/Lambda/services/persistence/scene"
)

type sceneController struct {
}

func (controller *sceneController) RegisterEventListeners() {
	event.Singleton().Subscribe(events.TypeOpenScene, controller.listenerOpenScene)
}

func (controller *sceneController) listenerOpenScene(action event.IEvent) {
	filename := action.(*events.OpenScene).Path()

	importer := importers.VmfImporter{}
	sceneModel, err := importer.LoadVmf(filename)
	if err != nil {
		event.Singleton().Dispatch(events.NewOpenSceneFailed())
		return
	}
	scene.Singleton().SetWorld(sceneModel.Worldspawn())
	scene.Singleton().SetEntities(sceneModel.Entities())

	for i := 0; i < scene.Singleton().Entities().Length(); i++ {
		event.Singleton().Dispatch(events.NewEntityCreated(scene.Singleton().Entities().Get(i)))
	}
}

func NewSceneController() *sceneController {
	return &sceneController{}
}
