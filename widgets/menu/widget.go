package menu

import (
	"github.com/galaco/Lambda/core/importers"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/inkyblackness/imgui-go"
)

type widget struct {
}

func (mod *widget) Initialize() {

}

func (mod *widget) Render() {
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("File") {
			if imgui.MenuItemV("New..", "Ctrl+N", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItemV("Open..", "Ctrl+O", false, true) {
				importer := importers.VmfImporter{}
				importer.LoadVmf("./ze_angel_beats.vmf")
				/* Do stuff */
			}
			if imgui.MenuItemV("Save", "Ctrl+S", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItemV("Close", "Ctrl+W", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItem("Exit") {
				event.Singleton().Dispatch(events.NewWindowClosed())
				/* Do stuff */
			}
			imgui.EndMenu()
		}
		imgui.EndMainMenuBar()
	}
}

func (mod *widget) Update() {

}

func NewWidget() *widget {
	return &widget{}
}
