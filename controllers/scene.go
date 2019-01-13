package controllers

import (
	"fmt"
	"github.com/galaco/Lambda/core/importers"
	"github.com/galaco/Lambda/core/persistance/scene"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/views/scenegraph"
)

type sceneController struct {
	nodeListView *scenegraph.Widget
}

func (controller *sceneController) RegisterEventListeners() {
	event.Singleton().Subscribe(events.TypeOpenScene, controller.listenerOpenScene)
	event.Singleton().Subscribe(events.TypeNewScene, controller.listenerNewScene)
}

func (controller *sceneController) Render() {
	controller.nodeListView.Render()
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
}

func (controller *sceneController) listenerNewScene(dispatched event.IEvent) {
	sceneModel := dispatched.(*events.NewScene).Model()
	scene.Singleton().SetWorld(sceneModel.Worldspawn())
	scene.Singleton().SetEntities(sceneModel.Entities())

	for i := 0; i < scene.Singleton().Entities().Length(); i++ {
		ent := scene.Singleton().Entities().Get(i)
		controller.nodeListView.AddNode(
			ent.IntForKey("id"),
			fmt.Sprintf("%d %s: %s", ent.IntForKey("id"), ent.ValueForKey("classname"), ent.ValueForKey("targetname")))
	}
}

func NewSceneController() *sceneController {
	return &sceneController{
		nodeListView: scenegraph.NewWidget(),
	}
}