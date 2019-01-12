package controllers

import (
	"fmt"
	"github.com/galaco/Lambda/core/importers"
	"github.com/galaco/Lambda/core/models/world"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/views/scenegraph"
	"github.com/galaco/source-tools-common/entity"
)

type sceneController struct {
	nodeListView *scenegraph.Widget

	world *world.World
	entities *entity.List
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
	scene := dispatched.(*events.NewScene).Model()
	controller.world = scene.Worldspawn()
	controller.entities = scene.Entities()

	for i := 0; i < controller.entities.Length(); i++ {
		ent := controller.entities.Get(i)
		controller.nodeListView.AddNode(
			ent.IntForKey("id"),
			fmt.Sprintf("%s: %s", ent.ValueForKey("classname"), ent.ValueForKey("targetname")))
	}
}

func NewSceneController() *sceneController {
	return &sceneController{
		nodeListView: scenegraph.NewWidget(),
	}
}