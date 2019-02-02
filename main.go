package main

import (
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/filesystem"
	"github.com/galaco/Lambda/filesystem/importers"
	"github.com/galaco/Lambda/graphics/opengl"
	"github.com/galaco/Lambda/project"
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
	app.Model = project.NewModel()
	app.GraphicsAdapter = &opengl.OpenGL{}

	uiContext := app.InitializeUIContext()
	app.InitializeGUITheme()
	app.InitializeViews()

	// Subscribe to window closing event
	windowShouldClose := false
	app.EventDispatcher.Subscribe(events.TypeWindowClosed, func(action event.IEvent) {
		windowShouldClose = true
	})

	for !uiContext.Window().ShouldClose() && !windowShouldClose {
		glfw.PollEvents()
		app.Render()

		displayWidth, displayHeight := uiContext.Window().GetFramebufferSize()
		app.GraphicsAdapter.Viewport(0, 0, int32(displayWidth), int32(displayHeight))
		app.GraphicsAdapter.ClearColor(0, 0, 0, 0)
		app.GraphicsAdapter.ClearAll()

		imgui.Render()
		uiContext.Imgui().Render(imgui.RenderedDrawData())

		uiContext.Window().SwapBuffers()
		<-time.After(time.Millisecond * 25)
	}
}
