package imgui_layouts

import (
	"github.com/galaco/Lambda/internal/ui/imgui-layouts/master"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/inkyblackness/imgui-go"
)

type Panel struct {
	panel *master.Panel

	previousWidth, previousHeight int
	sizeChangeCallbacks           []func(int, int)

	offset, size mgl32.Vec2
}

func (panel *Panel) WithDisplayRule(rule master.Rulable) *Panel {
	panel.panel.AddRule(rule)
	return panel
}

func (panel *Panel) Size() mgl32.Vec2 {
	return panel.size
}

func (panel *Panel) OnChangeSize(callback func(int, int)) {
	panel.sizeChangeCallbacks = append(panel.sizeChangeCallbacks, callback)
}

func (panel *Panel) Start(label string, width, height int) bool {
	if panel.previousWidth != width || panel.previousHeight != height {
		for _, cb := range panel.sizeChangeCallbacks {
			cb(width, height)
		}
		panel.offset, panel.size = panel.panel.Resolve(width, height)
	}
	panel.previousWidth = width
	panel.previousHeight = height

	imgui.SetNextWindowPos(imgui.Vec2{X: panel.offset[0], Y: panel.offset[1]})
	imgui.SetNextWindowSize(imgui.Vec2{X: panel.size[0], Y: panel.size[1]})

	return imgui.BeginV(label, nil, imgui.WindowFlagsNoResize|
		imgui.WindowFlagsNoMove|
		imgui.WindowFlagsNoBringToFrontOnFocus|
		imgui.WindowFlagsNoScrollbar|
		imgui.WindowFlagsNoScrollWithMouse|
		imgui.WindowFlagsNoNav|
		imgui.WindowFlagsNoInputs)
}

func (panel *Panel) End() {
	imgui.End()
}

func NewPanel() *Panel {
	return &Panel{
		panel: master.NewPanel(),
	}
}
