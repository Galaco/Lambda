package renderer

import (
	"fmt"
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda/graphics"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type fbo struct {
	adapter 		   graphics.Adapter
	framebuffer        uint32
	framebufferTexture uint32
	width              int
	height             int
}

func (win *fbo) Resize(width int, height int) {
	win.width = width
	win.height = height

	logger.Notice(fmt.Sprintf("%d, %d", width, height))

	win.Bind()

	if win.framebufferTexture != 0 {

		win.adapter.DeleteTextures(1, &win.framebufferTexture)
	}

	win.adapter.LambdaCreateTexture2D(&win.framebufferTexture, int32(win.width), int32(win.height), nil)

	win.adapter.LambdaBindTexture2D(win.framebufferTexture)
	if glError := gl.GetError(); glError != gl.NO_ERROR {
		logger.Error("error: %d\n", glError)
	}

	win.adapter.LambdaBindTexture2DToFramebuffer(win.framebufferTexture)

	if glError := gl.GetError(); glError != gl.NO_ERROR {
		logger.Error("error: %d\n", glError)
	}


	win.Unbind()
}

func (win *fbo) Bind() {
	win.adapter.LambdaBindFramebuffer(win.framebuffer)
}

func (win *fbo) Unbind() {
	win.adapter.LambdaBindFramebuffer(0)
}

func (win *fbo) Destroy() {
	win.adapter.DeleteFramebuffers(1, &win.framebuffer)
}

func NewFbo(adapter graphics.Adapter, width int, height int) *fbo {
	f := &fbo{
		adapter: adapter,
	}
	f.adapter.CreateFramebuffers(1, &f.framebuffer)
	f.Resize(width, height)
	return f
}
