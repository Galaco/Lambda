package renderer

import (
	"github.com/galaco/Lambda/internal/graphics"
)

// RenderWindow provides a broader wrapper for a framebuffer.
type RenderWindow struct {
	adapter     graphics.Adapter
	width       int
	height      int
	frameBuffer *fbo
}

// BufferId returns the RenderWindows bound FBO id.
func (win *RenderWindow) BufferId() uint32 {
	return win.frameBuffer.colourTexture
}

// Bind binds this RenderWindow(s fbo)
func (win *RenderWindow) Bind() {
	win.adapter.Viewport(0, 0, int32(win.width), int32(win.height))
	win.frameBuffer.Bind()

	win.adapter.ClearAll()
}

// Unbind this RenderWindow(s fbo)
func (win *RenderWindow) Unbind() {
	win.frameBuffer.Unbind()
}

// SetSize resizes the bound fbo
func (win *RenderWindow) SetSize(width int, height int) {
	win.width = width
	win.height = height
	win.frameBuffer.Destroy()
	win.frameBuffer = NewFbo(win.adapter, width, height)
}

// Close cleans up and destroys this RenderWindow.
func (win *RenderWindow) Close() {
	win.frameBuffer.Destroy()
}

// NewRenderWindow returns a new RenderWindow
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
