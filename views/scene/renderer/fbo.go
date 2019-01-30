package renderer

import (
	"github.com/go-gl/gl/all-core/gl"
)

type fbo struct {
	framebuffer        uint32
	framebufferTexture uint32
	width              int
	height             int
}

func (win *fbo) Resize(width int, height int) {
	win.width = width
	win.height = height

	win.Bind()

	if win.framebufferTexture != 0 {
		gl.DeleteTextures(1, &win.framebufferTexture)
	}

	gl.GenTextures(1, &win.framebufferTexture)
	gl.BindTexture(gl.TEXTURE_2D, win.framebufferTexture)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(win.width), int32(win.height), 0, gl.RGB, gl.UNSIGNED_BYTE, nil)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, win.framebufferTexture, 0)

	win.Unbind()
}

func (win *fbo) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, win.framebuffer)
}

func (win *fbo) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (win *fbo) Destroy() {
	gl.DeleteFramebuffers(1, &win.framebuffer)
}

func NewFbo(width int, height int) *fbo {
	f := &fbo{}
	gl.CreateFramebuffers(1, &f.framebuffer)
	f.Resize(width, height)
	return f
}
