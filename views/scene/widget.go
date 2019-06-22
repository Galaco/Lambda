package scene

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda/internal/event"
	"github.com/galaco/Lambda/internal/events"
	"github.com/galaco/Lambda/internal/graphics"
	"github.com/galaco/Lambda/internal/input"
	"github.com/galaco/Lambda/internal/renderer"
	"github.com/galaco/Lambda/internal/ui"
	"github.com/galaco/Lambda/internal/ui/context"
	"github.com/galaco/Lambda/internal/ui/imgui-layouts"
	"github.com/galaco/Lambda/internal/ui/imgui-layouts/master/rule"
	"github.com/inkyblackness/imgui-go"
)

// Widget
type Widget struct {
	masterPanel *imgui_layouts.Panel

	dispatcher      *event.Dispatcher
	keyboard        *input.Keyboard
	filesystem      *filesystem.FileSystem
	graphicsAdapter graphics.Adapter

	window   *renderer.RenderWindow
	renderer *renderer.Renderer

	width, height int

	scene  *Scene
	camera *entity.Camera
}

// Initialize
func (widget *Widget) Initialize() {
	widget.window = renderer.NewRenderWindow(widget.graphicsAdapter, widget.width, widget.height)
	widget.dispatcher.Subscribe(events.TypeNewSolidCreated, widget.newSolidCreated)
	widget.dispatcher.Subscribe(events.TypeNewCameraCreated, widget.newCameraCreated)
	widget.dispatcher.Subscribe(events.TypeCameraChanged, widget.cameraChanged)
	widget.dispatcher.Subscribe(events.TypeSceneClosed, widget.sceneClosed)

	widget.renderer.BindShader(widget.graphicsAdapter.LambdaLoadSimpleShader("assets/shaders/UnlitGeneric"))

	widget.masterPanel.OnChangeSize(func(width, height int) {
		widget.window.SetSize(width, height)
	})
}

// RenderScene
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

// Render
func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
	imgui.PushStyleColor(imgui.StyleColorChildBg, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 1})
	if widget.masterPanel.Start("Scene", w, h) {
		imgui.SetCursorPos(imgui.Vec2{X: 0, Y: 0})
		imgui.ImageV(imgui.TextureID(widget.window.BufferId()),
			imgui.Vec2{X: widget.masterPanel.Size().X(), Y: widget.masterPanel.Size().Y()},
			imgui.Vec2{X: 0, Y: 1},
			imgui.Vec2{X: 1, Y: 0},
			imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1}, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0})
		widget.masterPanel.End()
	}
	imgui.PopStyleColor()
}

// Update
func (widget *Widget) Update(dt float64) {
	if widget.keyboard.IsKeyDown(input.KeyW) {
		widget.scene.ActiveCamera().Forwards(dt * 0.05)
	}
	if widget.keyboard.IsKeyDown(input.KeyA) {
		widget.scene.ActiveCamera().Left(dt * 0.05)
	}
	if widget.keyboard.IsKeyDown(input.KeyS) {
		widget.scene.ActiveCamera().Backwards(dt * 0.05)
	}
	if widget.keyboard.IsKeyDown(input.KeyD) {
		widget.scene.ActiveCamera().Right(dt * 0.05)
	}

	if widget.keyboard.IsKeyDown(input.KeyUp) {
		widget.scene.ActiveCamera().Rotate(0, 0, float32(dt)*0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyLeft) {
		widget.scene.ActiveCamera().Rotate(float32(dt)*0.1, 0, 0)
	}
	if widget.keyboard.IsKeyDown(input.KeyDown) {
		widget.scene.ActiveCamera().Rotate(0, 0, -float32(dt)*0.1)
	}
	if widget.keyboard.IsKeyDown(input.KeyRight) {
		widget.scene.ActiveCamera().Rotate(-float32(dt)*0.1, 0, 0)
	}

	widget.scene.ActiveCamera().Update(1000 / 60)
}

func (widget *Widget) newSolidCreated(received event.Dispatchable) {
	widget.scene.AddSolid(received.(*events.NewSolidCreated).Target())
}

func (widget *Widget) newCameraCreated(received event.Dispatchable) {
	widget.scene.AddCamera(received.(*events.NewCameraCreated).Target())
}

func (widget *Widget) cameraChanged(received event.Dispatchable) {
	widget.scene.ChangeCamera(received.(*events.CameraChanged).Target())
}

func (widget *Widget) sceneClosed(received event.Dispatchable) {
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
		masterPanel: imgui_layouts.NewPanel().
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampTop, int(24.0*ui.DPIScale()))).
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampLeft, int(320.0*ui.DPIScale()))).
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampRight, int(320.0*ui.DPIScale()))).
			WithDisplayRule(rule.NewRuleFixedHeight(100, true, 320+int(24.0*ui.DPIScale()))),
	}
}
