package renderer

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda/graphics"
	"github.com/galaco/gosigl"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type RenderWindow struct {
	adapter     graphics.Adapter
	width       int
	height      int
	frameBuffer *fbo
	renderer *Renderer

	// temp
	verts []float32
	vbo, vao uint32

	m *gosigl.VertexObject
}

func (win *RenderWindow) BufferId() uint32 {
	return win.frameBuffer.framebufferTexture
}

func (win *RenderWindow) StartFrame(camera *entity.Camera) {
	win.renderer.StartFrame()
	win.renderer.BindCamera(camera)
}

func (win *RenderWindow) DrawFrame(scene *Scene) {
	win.StartFrame(scene.activeCamera)

	win.frameBuffer.Bind()

	gl.BindVertexArray(win.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	for _,solid := range scene.RenderableSolids {
		win.renderer.DrawSolid(solid)
	}

	win.frameBuffer.Unbind()
}

func (win *RenderWindow) SetSize(width int, height int) {
	win.width = width
	win.height = height
	win.frameBuffer.Destroy()
	win.frameBuffer = NewFbo(win.adapter, width, height)
}

func (win *RenderWindow) prepTriangle() {
	win.verts = []float32{-1, -1, 0.0, 1, -1, 0.0, 0.0, 1, 0.0}

	gl.GenVertexArrays(1, &win.vao)
	gl.GenBuffers(1, &win.vbo)
	gl.BindVertexArray(win.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, win.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(win.verts) * 4, gl.Ptr(win.verts), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3 * 4, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)
}

func NewRenderWindow(adapter graphics.Adapter, width int, height int) *RenderWindow {
	r := &RenderWindow{
		adapter:     adapter,
		width:       width,
		height:      height,
		frameBuffer: nil,
		renderer: newRenderer(),
	}


	r.renderer.BindShader(loadShader())
	r.frameBuffer = NewFbo(adapter, width, height)

	r.prepTriangle()

	return r
}
