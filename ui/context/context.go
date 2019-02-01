package context

import (
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Context struct {
	window       *glfw.Window
	imguiContext *imgui.Context
	imguiBind    *imguiGlfw3
}

func (ctx *Context) Window() *glfw.Window {
	return ctx.window
}

func (ctx *Context) Imgui() *imguiGlfw3 {
	return ctx.imguiBind
}

func (ctx *Context) Close() {
	defer ctx.window.Destroy()
	defer ctx.imguiContext.Destroy()
	defer ctx.imguiBind.Shutdown()
}

func NewContext() *Context {
	window := initGlfw()
	initOpenGL()
	ctx := &Context{
		window:       window,
		imguiContext: imgui.CreateContext(nil),
		imguiBind:    imguiGlfw3Init(window),
	}

	return ctx
}
