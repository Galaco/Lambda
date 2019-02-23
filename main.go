package main

import (
	"github.com/galaco/Lambda/internal/event"
	"github.com/galaco/Lambda/internal/events"
	"github.com/galaco/Lambda/internal/filesystem"
	"github.com/galaco/Lambda/internal/filesystem/importers"
	"github.com/galaco/Lambda/internal/graphics/opengl"
	"github.com/galaco/Lambda/internal/input"
	"github.com/galaco/Lambda/internal/model"
	"github.com/galaco/Lambda/internal/ui"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"time"
)

func main() {
	app := Application{}
	defer app.Close()

	app.FileSystem = filesystem.Init()
	app.EventDispatcher = event.NewDispatcher()
	app.VmfImporter = importers.NewVmfImporter()
	app.Model = model.NewModel()
	app.GraphicsAdapter = &opengl.OpenGL{}
	app.Keyboard = input.NewKeyboard()

	uiContext := app.InitializeUIContext()
	uiContext.Window().SetKeyCallback(app.Keyboard.GlfwKeyCallback)
	app.InitializeGUITheme()
	app.InitializeViews()

	// Subscribe to window closing event
	windowShouldClose := false
	app.EventDispatcher.Subscribe(events.TypeWindowClosed, func(action event.IEvent) {
		windowShouldClose = true
	})
	app.EventDispatcher.Subscribe(events.TypePreferencesUpdated, func(action event.IEvent) {
		ui.ApplyImguiStyles(action.(*events.PreferencesUpdated).Appearance.Theme)
	})

	for !uiContext.Window().ShouldClose() && !windowShouldClose {
		glfw.PollEvents()
		app.Render()

		displayWidth, displayHeight := uiContext.Window().GetFramebufferSize()
		app.GraphicsAdapter.Viewport(0, 0, int32(displayWidth), int32(displayHeight))

		imgui.Render()
		uiContext.Imgui().Render(imgui.RenderedDrawData())

		uiContext.Window().SwapBuffers()
		app.GraphicsAdapter.ClearColor(0, 0, 0, 0)
		app.GraphicsAdapter.ClearAll()

		app.Update()
		<-time.After(time.Millisecond * 25)
	}
}
