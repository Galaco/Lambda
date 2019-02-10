package renderer

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda/graphics"
)

type RenderWindow struct {
	adapter     graphics.Adapter
	width       int
	height      int
	frameBuffer *fbo
	renderer *Renderer
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

	win.adapter.Viewport(0, 0, int32(win.width), int32(win.height))
	win.frameBuffer.Bind()

	for _,solid := range scene.RenderableSolids {
		win.renderer.DrawSolid(solid)
	}

	win.frameBuffer.Unbind()

	scene.activeCamera.Update(1000.0/60)
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
		renderer: newRenderer(),
	}


	r.renderer.BindShader(loadShader())
	r.frameBuffer = NewFbo(adapter, width, height)

	return r
}
