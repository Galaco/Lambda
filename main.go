package main

import (
	"github.com/galaco/Lambda/controllers"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/views/assets"
	"github.com/galaco/Lambda/views/hierarchy"
	"github.com/galaco/Lambda/views/mainmenu"
	"github.com/galaco/Lambda/views/properties"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"runtime"
	"time"
)

func main() {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)

	window, err := glfw.CreateWindow(1280, 720, "Lambda", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	context := imgui.CreateContext(nil)
	defer context.Destroy()

	impl := imguiGlfw3Init(window)
	defer impl.Shutdown()

	var clearColor imgui.Vec4

	event.Singleton().Initialize()

	app := lib.NewApplication()
	app.AddController(controllers.NewMenuController())
	app.AddController(controllers.NewSceneController())
	app.AddController(controllers.NewKeyValuesController())
	app.AddView(mainmenu.NewWidget())
	app.AddView(hierarchy.NewWidget())
	app.AddView(properties.NewWidget())
	app.AddView(assets.NewWidget())

	windowShouldClose := false
	event.Singleton().Subscribe(events.TypeWindowClosed, func(action event.IEvent) {
		windowShouldClose = true
	})

	for !window.ShouldClose() && !windowShouldClose {
		glfw.PollEvents()
		impl.NewFrame()

		app.Render(window)

		displayWidth, displayHeight := window.GetFramebufferSize()
		gl.Viewport(0, 0, int32(displayWidth), int32(displayHeight))
		gl.ClearColor(clearColor.X, clearColor.Y, clearColor.Z, clearColor.W)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		impl.Render(imgui.RenderedDrawData())

		window.SwapBuffers()
		app.Update()
		<-time.After(time.Millisecond * 25)
	}
}
