package main

import (
	"github.com/galaco/Lambda/actions"
	"github.com/galaco/Lambda/lib"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/widgets/menu"
	"github.com/galaco/Lambda/widgets/scenegraph"
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

	app := lib.NewApplication()
	app.RegisterModule(menu.NewWidget())
	app.RegisterModule(scenegraph.NewWidget())

	event.Singleton().Initialize()

	event.Singleton().Listen(actions.TypeWindowClosed, func(action event.IAction) {
		window.Destroy()
	})

	for !window.ShouldClose() {
		glfw.PollEvents()
		impl.NewFrame()

		app.Render()


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
