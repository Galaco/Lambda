package renderer

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/gosigl"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	shader *gosigl.Context
	uniforms map[string]int32
}

func (renderer *Renderer) BindShader(shader *gosigl.Context) {
	renderer.shader = shader

	renderer.uniforms["model"] = renderer.shader.GetUniform("model")
	renderer.uniforms["view"] = renderer.shader.GetUniform("view")
	renderer.uniforms["projection"] = renderer.shader.GetUniform("projection")

	renderer.shader.UseProgram()
}

func (renderer *Renderer) StartFrame() {
	renderer.shader.UseProgram()
}

func (renderer *Renderer) BindCamera(cam *entity.Camera) {
	view := cam.ViewMatrix()
	proj := cam.ProjectionMatrix()

	gl.UniformMatrix4fv(renderer.uniforms["projection"], 1, false, &proj[0])
	gl.UniformMatrix4fv(renderer.uniforms["view"], 1, false, &view[0])
}

func (renderer *Renderer) DrawSolid(model []*gosigl.VertexObject) {
	for _,vobj := range model {
		gosigl.BindMesh(vobj)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
	}
}

func newRenderer() *Renderer {
	return &Renderer{
		uniforms: make(map[string]int32),
	}
}