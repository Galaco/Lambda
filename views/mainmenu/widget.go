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
	"io/ioutil"
)

type Widget struct {
	dispatcher *event.Dispatcher
	importer   *importers.VmfImporter
	exporter *exporters.VmfExporter
	model      *project.Model
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
			if imgui.MenuItemV("Save", "Ctrl+S", false, true) {
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
			if imgui.MenuItemV("Save As", "", false, true) {
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
}

func NewWidget(dispatcher *event.Dispatcher, model *project.Model, importer *importers.VmfImporter, exporter *exporters.VmfExporter) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		importer:   importer,
		model:      model,
		exporter:   exporter,
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

func saveFile(filename string, data string) (err error) {
	// Saving a new file
	if filename == "" {
		filename,err = dialog.FileSave()
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(filename, []byte(data), 755)
}
