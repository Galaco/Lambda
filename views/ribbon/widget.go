package ribbon

import (
	"github.com/galaco/Lambda/internal/ui/context"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
}

func (mod *Widget) Initialize() {
}

func (mod *Widget) Render(ctx *context.Context) {
	w, _ := ctx.Window().GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 16})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(w), Y: 16})
	if imgui.BeginV("Ribbon", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus|imgui.WindowFlagsNoTitleBar) {

		imgui.End()
	}
}

func NewWidget() *Widget {
	return &Widget{}
}
