package assets

import (
	lambdaFS "github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/ui/imgui-layouts/columns"
	"github.com/galaco/Lambda/views/assets/structure"
	"github.com/galaco/Lambda/views/assets/structure/directory"
	"github.com/inkyblackness/imgui-go"
	"log"
)

type Widget struct {
	dispatcher *event.Dispatcher
	fileSystem *lambdaFS.FileSystem

	twoPanel         *columns.View
	directoryList    *structure.Tree
	currentDirectory *directory.Directory
}

func (widget *Widget) Initialize() {
	widget.currentDirectory = directory.NewDirectory(widget.fileSystem.EnumerateResourcePaths())

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

func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(320), Y: float32(h / 2)})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(w - 640), Y: float32(h / 2)})
	if imgui.BeginV("Assets", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		widget.twoPanel.Render(w-640, h/2)
		imgui.End()
	}
}

func (widget *Widget) selectedEntityChanged(received event.IEvent) {
}

func NewWidget(dispatcher *event.Dispatcher, fileSystem *lambdaFS.FileSystem) *Widget {
	return &Widget{
		dispatcher:       dispatcher,
		fileSystem:       fileSystem,
		directoryList:    structure.NewTree(),
		currentDirectory: directory.NewDirectory([]string{}),
	}
}
