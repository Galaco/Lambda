package controllers

import (
	"github.com/galaco/Lambda/core/importers"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"log"
)

type sceneController struct {

}

func (controller *sceneController) RegisterEventListeners() {
	event.Singleton().Subscribe(events.TypeOpenScene, controller.listenerOpenScene)
}

func (controller *sceneController) listenerOpenScene(action event.IEvent) {
	filename := action.(*events.OpenScene).Path()

	importer := importers.VmfImporter{}
	scene,err := importer.LoadVmf(filename)
	if err != nil {
		event.Singleton().Dispatch(events.NewOpenSceneFailed())
		return
	}
	event.Singleton().Dispatch(events.NewNewScene(scene))
	log.Println("opened scene " + filename)
}

func NewSceneController() *sceneController {
	return &sceneController{}
}