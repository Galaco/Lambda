package renderer

import (
	"github.com/galaco/Lambda/internal/graphics"
)

type fbo struct {
	adapter       graphics.Adapter
	framebuffer   uint32
	colourTexture uint32
	depthTexture  uint32
	width         int
	height        int
}

// Resize resizes this framebuffer object
func (win *fbo) Resize(width int, height int) {
	win.width = width
	win.height = height

	win.Bind()

	if win.colourTexture != 0 {
		win.adapter.DeleteTextures(1, &win.colourTexture)
		win.adapter.DeleteRenderBuffer(1, &win.depthTexture)
	}

	win.depthTexture = win.adapter.LambdaCreateRenderbufferStorageDepth(int32(win.width), int32(win.height))

	win.adapter.LambdaCreateTextureStorage2D(&win.colourTexture, int32(win.width), int32(win.height))
	win.adapter.LambdaBindTexture2D(win.colourTexture)
	win.adapter.LambdaBindTexture2DToFramebuffer(win.colourTexture)
	win.adapter.LambdaBindDepthBufferToFramebuffer(win.depthTexture)
	win.adapter.LambdaDrawBuffers()
	win.adapter.ClearColor(0, 0, 0, 0)
	win.adapter.ClearAll()
	win.adapter.LambdaBindTexture2D(0)

	win.Unbind()
}

// Bind this framebuffer
func (win *fbo) Bind() {
	win.adapter.LambdaBindFramebuffer(win.framebuffer)
}

// Unbind unbind this framebuffer
func (win *fbo) Unbind() {
	win.adapter.LambdaBindFramebuffer(0)
}

// Destroy deletes and cleans up this framebuffer
func (win *fbo) Destroy() {
	win.adapter.DeleteFramebuffers(1, &win.framebuffer)
}

// NewFbo returns a new framebuffer
func NewFbo(adapter graphics.Adapter, width int, height int) *fbo {
	f := &fbo{
		adapter: adapter,
	}
	f.adapter.CreateFramebuffers(1, &f.framebuffer)
	f.Resize(width, height)
	return f
}
