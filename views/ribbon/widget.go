package ribbon

import (
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
}

func (mod *Widget) Initialize() {
}

func (mod *Widget) Render(window *glfw.Window) {
	w,_ := window.GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 16})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(w), Y: 16})
	if imgui.BeginV("Ribbon", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus|imgui.WindowFlagsNoTitleBar) {

		imgui.End()
	}
}

func (mod *Widget) Update() {

}

func (mod *Widget) Destroy() {

}

func NewWidget() *Widget {
	return &Widget{}
}
