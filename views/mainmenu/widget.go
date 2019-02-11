package mainmenu

import (
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/filesystem/exporters"
	"github.com/galaco/Lambda/filesystem/importers"
	"github.com/galaco/Lambda/project"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/views/mainmenu/dialog"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	dispatcher *event.Dispatcher
	importer   *importers.VmfImporter
	exporter *exporters.VmfExporter
	model      *project.Model

	isProjectLoaded bool

	dialogPreferences *dialog.Preferences
}

// Initialize sets up widget specific properties.
func (widget *Widget) Initialize() {

}

// Render draws the main menu bar across the top of the screen
func (widget *Widget) Render(ctx *context.Context) {
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("File") {
			if imgui.MenuItemV("New..", "Ctrl+N", false, true) {
				/* Do stuff */
			}
			if imgui.MenuItemV("Open..", "Ctrl+O", false, true) {
				if filename := openFile(); filename != "" {
					widget.loadVmf(filename)
				}
			}
			if imgui.MenuItemV("Save", "Ctrl+S", false, widget.isProjectLoaded) {
				data,err := widget.exporter.Export(widget.model.Vmf)
				if err == nil {
					err = saveFile(widget.model.Filename, data)
					if err != nil {
						logger.Error(err)
					}
				} else {
					logger.Error(err)
				}
			}
			if imgui.MenuItemV("Save As", "", false, widget.isProjectLoaded) {
				data,err := widget.exporter.Export(widget.model.Vmf)
				if err == nil {
					err = saveFile("", data)
					if err != nil {
						logger.Error(err)
					}
				} else {
					logger.Error(err)
				}
			}
			if imgui.MenuItemV("Close", "Ctrl+W", false, widget.isProjectLoaded) {
				widget.dispatcher.Dispatch(events.NewSceneClosed())
				widget.isProjectLoaded = false
			}
			if imgui.MenuItem("Exit") {
				widget.dispatcher.Dispatch(events.NewWindowClosed())
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Edit") {
			if imgui.MenuItem("Preferences") {
				widget.dialogPreferences.Open()
			}
			imgui.EndMenu()
		}
		imgui.EndMainMenuBar()
	}

	if widget.dialogPreferences.IsOpen() {
		widget.dialogPreferences.Render(dialogWidth, dialogHeight)
	}
}

func (widget *Widget) loadVmf(filename string) {
	widget.dispatcher.Dispatch(events.NewOpenScene(filename))

	sceneModel, err := widget.importer.LoadVmf(filename)
	if err != nil {
		widget.dispatcher.Dispatch(events.NewOpenSceneFailed())
		return
	}
	widget.model.Vmf = sceneModel
	widget.model.Filename = filename

	for i := 0; i < widget.model.Vmf.Entities().Length(); i++ {
		widget.dispatcher.Dispatch(events.NewEntityCreated(widget.model.Vmf.Entities().Get(i)))
	}

	for i := 0; i < len(widget.model.Vmf.Worldspawn().Solids); i++ {
		widget.dispatcher.Dispatch(events.NewNewSolidCreated(&widget.model.Vmf.Worldspawn().Solids[i]))
	}

	for i := 0; i < len(widget.model.Vmf.Cameras().CameraList); i++ {
		widget.dispatcher.Dispatch(events.NewNewCameraCreated(&widget.model.Vmf.Cameras().CameraList[i]))
	}
	if widget.model.Vmf.Cameras().ActiveCamera != -1 {
		widget.dispatcher.Dispatch(events.NewCameraChanged(&widget.model.Vmf.Cameras().CameraList[widget.model.Vmf.Cameras().ActiveCamera]))
	}

	widget.isProjectLoaded = true
}

func NewWidget(dispatcher *event.Dispatcher, model *project.Model, importer *importers.VmfImporter, exporter *exporters.VmfExporter) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		importer:   importer,
		model:      model,
		exporter:   exporter,
		dialogPreferences: dialog.NewPreferences(),
	}
}