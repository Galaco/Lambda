package renderer

import (
	"github.com/galaco/Lambda/graphics"
)

type Window struct {
	adapter     graphics.Adapter
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
	win.adapter.ClearColor(1, 0, 1, 0)
	win.adapter.ClearAll()

	win.frameBuffer.Unbind()
	win.adapter.ClearColor(0, 0, 0, 0)
}

func (win *Window) SetSize(width int, height int) {
	win.width = width
	win.height = height
	win.frameBuffer.Destroy()
	win.frameBuffer = NewFbo(win.adapter, width, height)
}

func NewWindow(adapter graphics.Adapter, width int, height int) *Window {
	return &Window{
		adapter:     adapter,
		width:       width,
		height:      height,
		frameBuffer: NewFbo(adapter, width, height),
	}
}
