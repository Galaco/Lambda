package renderer

import (
	"github.com/galaco/Lambda/internal/graphics"
)

type RenderWindow struct {
	adapter     graphics.Adapter
	width       int
	height      int
	frameBuffer *fbo
}

func (win *RenderWindow) BufferId() uint32 {
	return win.frameBuffer.framebufferTexture
}

func (win *RenderWindow) Bind() {
	win.adapter.Viewport(0, 0, int32(win.width), int32(win.height))
	win.frameBuffer.Bind()

	win.adapter.ClearAll()
}

func (win *RenderWindow) Unbind() {
	win.frameBuffer.Unbind()
}

func (win *RenderWindow) SetSize(width int, height int) {
	win.width = width
	win.height = height
	win.frameBuffer.Destroy()
	win.frameBuffer = NewFbo(win.adapter, width, height)
}

func (win *RenderWindow) Close() {
	win.frameBuffer.Destroy()
}

func NewRenderWindow(adapter graphics.Adapter, width int, height int) *RenderWindow {
	r := &RenderWindow{
		adapter:     adapter,
		width:       width,
		height:      height,
		frameBuffer: nil,
	}
	r.frameBuffer = NewFbo(adapter, width, height)

	return r
}
