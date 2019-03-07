package mainmenu

import (
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda/internal/event"
	"github.com/galaco/Lambda/internal/events"
	"github.com/galaco/Lambda/internal/filesystem/exporters"
	"github.com/galaco/Lambda/internal/filesystem/importers"
	"github.com/galaco/Lambda/internal/model"
	"github.com/galaco/Lambda/internal/ui/context"
	"github.com/galaco/Lambda/views/mainmenu/dialog"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	dispatcher *event.Dispatcher
	importer   *importers.VmfImporter
	exporter   *exporters.VmfExporter
	model      *model.Model

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
				data, err := widget.exporter.Export(widget.model.Project.Vmf)
				if err == nil {
					err = saveFile(widget.model.Project.Filename, data)
					if err != nil {
						logger.Error(err)
					}
				} else {
					logger.Error(err)
				}
			}
			if imgui.MenuItemV("Save As", "", false, widget.isProjectLoaded) {
				data, err := widget.exporter.Export(widget.model.Project.Vmf)
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
	project := model.NewProject()
	project.Vmf = sceneModel
	project.Filename = filename

	widget.model.Project = project

	for i := 0; i < widget.model.Project.Vmf.Entities().Length(); i++ {
		widget.dispatcher.Dispatch(events.NewEntityCreated(project.Vmf.Entities().Get(i)))
	}

	for i := 0; i < len(project.Vmf.Worldspawn().Solids); i++ {
		widget.dispatcher.Dispatch(events.NewNewSolidCreated(&project.Vmf.Worldspawn().Solids[i]))
	}

	for i := 0; i < len(project.Vmf.Cameras().CameraList); i++ {
		widget.dispatcher.Dispatch(events.NewNewCameraCreated(&project.Vmf.Cameras().CameraList[i]))
	}
	if project.Vmf.Cameras().ActiveCamera != -1 {
		widget.dispatcher.Dispatch(events.NewCameraChanged(&project.Vmf.Cameras().CameraList[project.Vmf.Cameras().ActiveCamera]))
	}

	widget.isProjectLoaded = true
}

func NewWidget(dispatcher *event.Dispatcher, model *model.Model, importer *importers.VmfImporter, exporter *exporters.VmfExporter) *Widget {
	return &Widget{
		dispatcher:        dispatcher,
		importer:          importer,
		model:             model,
		exporter:          exporter,
		dialogPreferences: dialog.NewPreferences(dispatcher, model.Preferences),
	}
}
