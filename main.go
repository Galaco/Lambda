package main

import (
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/project"
	"github.com/galaco/Lambda/filesystem"
	"github.com/galaco/Lambda/filesystem/importers"
	"github.com/go-gl/gl/v4.1-core/gl"
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
		uiContext.Imgui().NewFrame()

		app.Render()

		displayWidth, displayHeight := uiContext.Window().GetFramebufferSize()
		gl.Viewport(0, 0, int32(displayWidth), int32(displayHeight))
		gl.ClearColor(0, 0, 0, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		uiContext.Imgui().Render(imgui.RenderedDrawData())

		uiContext.Window().SwapBuffers()
		//app2.Update()
		<-time.After(time.Millisecond * 25)
	}
}
