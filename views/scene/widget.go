package scene

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/graphics"
	"github.com/galaco/Lambda/input"
	"github.com/galaco/Lambda/renderer"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/ui/imgui-layouts"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	imgui_layouts.Panel

	dispatcher      *event.Dispatcher
	keyboard 		*input.Keyboard
	filesystem 		*filesystem.FileSystem
	graphicsAdapter graphics.Adapter

	window   *renderer.RenderWindow
	renderer *renderer.Renderer

	width, height int

	scene  *Scene
	camera *entity.Camera

	isActive bool
}

func (widget *Widget) Initialize() {
	widget.window = renderer.NewRenderWindow(widget.graphicsAdapter, widget.width, widget.height)
	widget.dispatcher.Subscribe(events.TypeNewSolidCreated, widget.newSolidCreated)
	widget.dispatcher.Subscribe(events.TypeNewCameraCreated, widget.newCameraCreated)
	widget.dispatcher.Subscribe(events.TypeCameraChanged, widget.cameraChanged)
	widget.dispatcher.Subscribe(events.TypeSceneClosed, widget.sceneClosed)

	widget.DisplayProperties.HasTitleBar = true
	widget.DisplayProperties.HasMenuBar = false

	widget.renderer.BindShader(renderer.LoadShader())
}

func (widget *Widget) RenderScene(ctx *context.Context) {
	dirtyComposition := widget.scene.frameCompositor.IsOutdated()
	if dirtyComposition {
		widget.scene.RecomposeScene(widget.filesystem)
	}

	widget.renderer.StartFrame()
	widget.renderer.BindCamera(widget.scene.ActiveCamera())
	widget.window.Bind()
	widget.renderer.DrawComposition(widget.scene.frameComposed, widget.scene.Composition(), widget.scene.CompositionMaterials())
	widget.graphicsAdapter.Error()
	widget.window.Unbind()
}

func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
	widgetWidth := int(w - (2 * 320))
	widgetHeight := int(h - 48) // / 2

	if widgetWidth != widget.width || widgetHeight != widget.height {
		widget.width = widgetWidth
		widget.height = widgetHeight
		widget.window.SetSize(widget.width, widget.height)
	}
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(320), Y: 48})
	imgui.SetNextWindowSize(imgui.Vec2{X: float32(widgetWidth), Y: float32(widgetHeight)})

	imgui.PushStyleColor(imgui.StyleColorChildBg, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{X: 0, Y: 0})
	if widget.Start() {

		imgui.SetCursorPos(imgui.Vec2{
			X: 0, //float32(widget.width / 2),
			Y: 0, //float32(widget.height / 2),
		})
		widget.graphicsAdapter.Viewport(0, 0, int32(widget.width), int32(widget.height))
		imgui.ImageV(imgui.TextureID(widget.window.BufferId()), imgui.Vec2{
			X: float32(widget.width),
			Y: float32(widget.height)},
			imgui.Vec2{},
			imgui.Vec2{1, 1},
			imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1}, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		widget.graphicsAdapter.Viewport(0, 0, int32(w), int32(h))
		widget.End()
	}
	imgui.PopStyleVar()
	imgui.PopStyleColor()
}

func (widget *Widget) Update(dt float64) {
	if widget.keyboard.IsKeyDown(input.KeyW) {
		widget.scene.ActiveCamera().Forwards(dt * 0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyA) {
		widget.scene.ActiveCamera().Left(dt * 0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyS) {
		widget.scene.ActiveCamera().Backwards(dt * 0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyD) {
		widget.scene.ActiveCamera().Right(dt * 0.1)
	}


	if widget.keyboard.IsKeyDown(input.KeyUp) {
		widget.scene.ActiveCamera().Rotate(0, 0, -float32(dt)*0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyLeft) {
		widget.scene.ActiveCamera().Rotate(float32(dt)*0.1, 0, 0)
	}
	if widget.keyboard.IsKeyDown(input.KeyDown) {
		widget.scene.ActiveCamera().Rotate(0, 0, float32(dt)*0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyRight) {
		widget.scene.ActiveCamera().Rotate(-float32(dt)*0.1, 0, 0)
	}

	widget.scene.ActiveCamera().Update(1000 / 60)
	//widget.scene.ActiveCamera().Rotate(float32(dt)*0.1, 0, 0)
}

func (widget *Widget) newSolidCreated(received event.IEvent) {
	widget.scene.AddSolid(received.(*events.NewSolidCreated).Target())
}

func (widget *Widget) newCameraCreated(received event.IEvent) {
	widget.scene.AddCamera(received.(*events.NewCameraCreated).Target())
}

func (widget *Widget) cameraChanged(received event.IEvent) {
	widget.scene.ChangeCamera(received.(*events.CameraChanged).Target())
}

func (widget *Widget) sceneClosed(received event.IEvent) {
	widget.scene.Close()
	widget.scene = NewScene()
	widget.window.Close()
	widget.window = renderer.NewRenderWindow(widget.graphicsAdapter, widget.width, widget.height)
}

func (widget *Widget) Close() {
	widget.scene.Close()
	widget.window.Close()
}

func NewWidget(dispatcher *event.Dispatcher, filesystem *filesystem.FileSystem, keyboard *input.Keyboard, graphicsAdapter graphics.Adapter) *Widget {
	return &Widget{
		dispatcher:      dispatcher,
		keyboard:        keyboard,
		filesystem:      filesystem,
		graphicsAdapter: graphicsAdapter,
		width:           1024,
		height:          768,
		scene:           NewScene(),
		renderer:        renderer.NewRenderer(graphicsAdapter),
	}
}
