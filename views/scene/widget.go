package scene

import (
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/graphics"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/valve/world"
	"github.com/galaco/Lambda/views/scene/renderer"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	dispatcher *event.Dispatcher
	graphicsAdapter graphics.Adapter

	window        *renderer.Window
	solids 		  []*world.Solid
	width, height int
}

func (widget *Widget) Initialize() {
	widget.window = renderer.NewWindow(widget.graphicsAdapter, widget.width, widget.height)
	widget.dispatcher.Subscribe(events.TypeNewSolidCreated, widget.newSolidCreated)
}

func (widget *Widget) RenderScene(ctx *context.Context) {
	widget.window.DrawFrame()
}

func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
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
		imgui.SetCursorPos(imgui.Vec2{
			X: float32(320) + float32(widget.width / 2),
			Y: float32(48) + float32(widget.height / 2),
		})
		imgui.ImageV(imgui.TextureID(widget.window.BufferId()),
			imgui.Vec2{},
			imgui.Vec2{X: 0, Y: 1},
			imgui.Vec2{X: 1, Y: 0},
			imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1},
			imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})

		imgui.End()
	}
}

func (widget *Widget) newSolidCreated(received event.IEvent) {
	widget.solids = append(widget.solids, received.(*events.NewSolidCreated).Target())
}

func NewWidget(dispatcher *event.Dispatcher, graphicsAdapter graphics.Adapter) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		graphicsAdapter:  graphicsAdapter,
		width:  1024,
		height: 768,
		solids: make([]*world.Solid, 0),
	}
}
