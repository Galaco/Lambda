package assets

import (
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/views/assets/structure"
	"github.com/galaco/Lambda/views/assets/structure/directory"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
	directoryList *structure.Tree
	currentDirectory *directory.Directory
}

func (mod *Widget) Initialize() {
}

func (mod *Widget) Render(window *glfw.Window) {
	w, h := window.GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(320), Y: float32(h / 2)})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(w - 640), Y: float32(h / 2)})
	if imgui.BeginV("Assets", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		mod.directoryList.Render()
		if mod.currentDirectory != nil {
			mod.currentDirectory.Render()
		}
		imgui.End()
	}
}

func (mod *Widget) Update() {

}

func (mod *Widget) selectedEntityChanged(received event.IEvent) {
}

func NewWidget() *Widget {
	return &Widget{
		directoryList: structure.NewTree(),
	}
}
