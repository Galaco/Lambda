package context

import (
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"runtime"
)

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

	window, err := glfw.CreateWindow(1920, 1080, "Lambda", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	return window
}
