package controllers

type menuController struct {
}

func (controller *menuController) RegisterEventListeners() {
}

func NewMenuController() *menuController {
	return &menuController{
	}
}