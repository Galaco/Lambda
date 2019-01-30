package assets

import (
	"github.com/galaco/Lambda/lib/imgui-layouts/columns"
	"github.com/galaco/Lambda/lib/mvc/event"
	"github.com/galaco/Lambda/services/filesystem"
	"github.com/galaco/Lambda/views/assets/structure"
	"github.com/galaco/Lambda/views/assets/structure/directory"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"log"
)

type Widget struct {
	twoPanel 		*columns.View
	directoryList    *structure.Tree
	currentDirectory *directory.Directory
}

func (widget *Widget) Initialize() {
	widget.currentDirectory = directory.NewDirectory(filesystem.Singleton().EnumerateResourcePaths())

	widget.twoPanel = columns.NewColumns(2)
	err := widget.twoPanel.SetColumnContents(0, widget.directoryList.Render, columns.NewColumnWidth(100, false))
	if err != nil {
		log.Println(err)
	}
	err = widget.twoPanel.SetColumnContents(1, widget.currentDirectory.Render, nil)
	if err != nil {
		log.Println(err)
	}
}

func (widget *Widget) Render(window *glfw.Window) {
	w, h := window.GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(320), Y: float32(h / 2)})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(w - 640), Y: float32(h / 2)})
	if imgui.BeginV("Assets", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		widget.twoPanel.Render(w - 640, h / 2)
		imgui.End()
	}
}

func (widget *Widget) Update() {

}

func (widget *Widget) Destroy() {

}

func (widget *Widget) selectedEntityChanged(received event.IEvent) {
}

func NewWidget() *Widget {
	return &Widget{
		directoryList:    structure.NewTree(),
		currentDirectory: directory.NewDirectory([]string{}),
	}
}
