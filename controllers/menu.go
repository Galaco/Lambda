package controllers

import (
	"github.com/galaco/Lambda/views/menu"
)

type menuController struct {
	mainmenuView *menu.Widget
}

func (controller *menuController) RegisterEventListeners() {
}

func (controller *menuController) Render() {
	controller.mainmenuView.Render()
}

func NewMenuController() *menuController {
	return &menuController{
		mainmenuView: menu.NewWidget(),
	}
}