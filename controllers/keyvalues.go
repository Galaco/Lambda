package controllers

import (
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/mvc/event"
	"github.com/galaco/Lambda/services/persistence/scene"
	"strconv"
)

type keyValuesController struct {
}

func (controller *keyValuesController) RegisterEventListeners() {
	event.Singleton().Subscribe(events.TypeSceneNodeSelected, controller.listenEntityNodeSelected)
}

func (controller *keyValuesController) listenEntityNodeSelected(action event.IEvent) {
	e := action.(*events.SceneNodeSelected)
	ent := scene.Singleton().Entities().FindByKeyValue("id", strconv.Itoa(e.Id))
	event.Singleton().Dispatch(events.NewEntitySelected(ent))
}

func NewKeyValuesController() *keyValuesController {
	return &keyValuesController{}
}
