package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Window struct {
	width       int
	height      int
	frameBuffer *fbo
}

func (win *Window) BufferId() uint32 {
	return win.frameBuffer.framebufferTexture
}

func (win *Window) DrawFrame() {
	win.frameBuffer.Bind()

	//RENDER
	gl.ClearColor(1, 0, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	win.frameBuffer.Unbind()
	gl.ClearColor(0, 0, 0, 0)
}

func (win *Window) SetSize(width int, height int) {
	win.width = width
	win.height = height
	win.frameBuffer.Destroy()
	win.frameBuffer = NewFbo(width, height)
}

func NewWindow(width int, height int) *Window {
	return &Window{
		width:       width,
		height:      height,
		frameBuffer: NewFbo(width, height),
	}
}
