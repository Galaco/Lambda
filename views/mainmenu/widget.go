package mainmenu

import (
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/mvc/event"
	"github.com/galaco/Lambda/views/mainmenu/dialog"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
}

func (mod *Widget) Initialize() {

}

func (mod *Widget) Render(window *glfw.Window) {
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("File") {
			if imgui.MenuItemV("New..", "Ctrl+N", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItemV("Open..", "Ctrl+O", false, true) {
				/* Do stuff */
				// This needs to dispatch an event that will actually call load elsewhere
				if filename := openFile(); filename != "" {
					event.Singleton().Dispatch(events.NewOpenScene(filename))
				}
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

func (mod *Widget) Update() {

}

func (mod *Widget) Destroy() {

}

func NewWidget() *Widget {
	return &Widget{}
}

func openFile() string {
	filename,err := dialog.FileOpen()
	if err != nil {
		filename = "./ze_bioshock_v6_4.vmf"
		return ""
	}
	return filename
}
