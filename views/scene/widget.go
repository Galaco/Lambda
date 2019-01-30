package scene

import (
	"github.com/galaco/Lambda/views/scene/renderer"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
	window        *renderer.Window
	width, height int
}

func (widget *Widget) Initialize() {
	widget.window = renderer.NewWindow(widget.width, widget.height)
}

func (widget *Widget) Render(window *glfw.Window) {
	w, h := window.GetSize()
	widgetWidth := int(w - (2 * 320))
	widgetHeight := int(h - 48)
	if widgetWidth != widget.width || widgetHeight != widget.height {
		widget.width = widgetWidth
		widget.height = widgetHeight
		widget.window.SetSize(widget.width, widget.height)
	}
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(320), Y: 48})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(widgetWidth), Y: float32(widgetHeight)})
	if imgui.BeginV("Scene", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		widget.window.DrawFrame()
		imgui.End()
	}
}

func (widget *Widget) Update() {

}

func (widget *Widget) Destroy() {

}

func NewWidget() *Widget {
	return &Widget{
		width:  1024,
		height: 768,
	}
}
