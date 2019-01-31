package mainmenu

import (
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/project"
	"github.com/galaco/Lambda/filesystem/importers"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/views/mainmenu/dialog"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	dispatcher *event.Dispatcher
	importer *importers.VmfImporter
	model *project.Model
}

func (widget *Widget) Initialize() {

}

func (widget *Widget) Render(ctx *context.Context) {
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("File") {
			if imgui.MenuItemV("New..", "Ctrl+N", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItemV("Open..", "Ctrl+O", false, true) {
				/* Do stuff */
				// This needs to dispatch an event that will actually call load elsewhere
				if filename := openFile(); filename != "" {
					widget.loadVmf(filename)
					//widget.dispatcher.Dispatch(events.NewOpenScene(filename))
				}
			}
			if imgui.MenuItemV("Save", "Ctrl+S", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItemV("Close", "Ctrl+W", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItem("Exit") {
				widget.dispatcher.Dispatch(events.NewWindowClosed())
				/* Do stuff */
			}
			imgui.EndMenu()
		}
		imgui.EndMainMenuBar()
	}
}

func (widget *Widget) Update() {

}

func (widget *Widget) Destroy() {

}

func (widget *Widget) loadVmf(filename string) {
	widget.dispatcher.Dispatch(events.NewOpenScene(filename))

	sceneModel, err := widget.importer.LoadVmf(filename)
	if err != nil {
		widget.dispatcher.Dispatch(events.NewOpenSceneFailed())
		return
	}
	widget.model.Scene().SetWorld(sceneModel.Worldspawn())
	widget.model.Scene().SetEntities(sceneModel.Entities())

	for i := 0; i < widget.model.Scene().Entities().Length(); i++ {
		widget.dispatcher.Dispatch(events.NewEntityCreated(widget.model.Scene().Entities().Get(i)))
	}
}

func NewWidget(dispatcher *event.Dispatcher, importer *importers.VmfImporter, model *project.Model) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		importer:importer,
		model:model,
	}
}

func openFile() string {
	filename, err := dialog.FileOpen()
	if err != nil {
		filename = "./ze_bioshock_v6_4.vmf"
		return ""
	}
	return filename
}
