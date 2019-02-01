package opengl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type OpenGL struct {
}

func (ogl *OpenGL) Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func (ogl *OpenGL) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (ogl *OpenGL) Clear(mask uint32) {
	gl.Clear(mask)
}
