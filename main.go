package main

import (
	"github.com/galaco/Lambda/controllers"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/mvc"
	"github.com/galaco/Lambda/lib/mvc/event"
	"github.com/galaco/Lambda/services/filesystem"
	"github.com/galaco/Lambda/views/assets"
	"github.com/galaco/Lambda/views/hierarchy"
	"github.com/galaco/Lambda/views/mainmenu"
	"github.com/galaco/Lambda/views/properties"
	"github.com/galaco/Lambda/views/ribbon"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"log"
	"runtime"
	"time"
)

func main() {
	// Window & OpenGL
	window := initGlfw()
	defer glfw.Terminate()
	defer window.Destroy()
	initOpenGL()
	context := imgui.CreateContext(nil)
	applyImguiStyles()
	defer context.Destroy()
	impl := imguiGlfw3Init(window)
	defer impl.Shutdown()

	// Begin event manager
	filesystem.Init()
	event.Singleton().Initialize()

	// Define application
	app := mvc.NewApplication()
	defer app.Close()
	app.AddController(controllers.NewMenuController())
	app.AddController(controllers.NewSceneController())
	app.AddController(controllers.NewKeyValuesController())
	app.AddView(mainmenu.NewWidget())
	app.AddView(hierarchy.NewWidget())
	app.AddView(properties.NewWidget())
	app.AddView(ribbon.NewWidget())
	//app.AddView(scene.NewWidget())
	app.AddView(assets.NewWidget())

	// Subscribe to window closing event
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
		gl.ClearColor(0, 0, 0, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		impl.Render(imgui.RenderedDrawData())

		window.SwapBuffers()
		app.Update()
		<-time.After(time.Millisecond * 25)
	}
}

// initOpenGL initializes OpenGL and returns an initialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func initGlfw() *glfw.Window {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)

	window, err := glfw.CreateWindow(1280, 720, "Lambda", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	return window
}
