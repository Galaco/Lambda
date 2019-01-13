package controllers

import (
	"github.com/galaco/Lambda/core/persistance/scene"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/views/keyvalues"
	"strconv"
)

type keyValuesController struct {
	entityPropertiesView *keyvalues.Widget
}

func (controller *keyValuesController) RegisterEventListeners() {
	event.Singleton().Subscribe(events.TypeSceneNodeSelected, controller.listenEntityNodeSelected)
}

func (controller *keyValuesController) Render() {
	controller.entityPropertiesView.Render()
}

func (controller *keyValuesController) listenEntityNodeSelected(action event.IEvent) {
	e := action.(*events.SceneNodeSelected)
	ent := scene.Singleton().Entities().FindByKeyValue("id", strconv.Itoa(e.Id))
	controller.entityPropertiesView.SetActiveEntity(ent)
}



func NewKeyValuesController() *keyValuesController {
	return &keyValuesController{
		entityPropertiesView: keyvalues.NewWidget(),
	}
}