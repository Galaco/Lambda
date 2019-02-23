package renderer

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda/internal/graphics"
	"github.com/galaco/Lambda/internal/renderer/render3d"
	"github.com/galaco/gosigl"
)

type Renderer struct {
	adapter graphics.Adapter

	shader   *gosigl.Context
	uniforms map[string]int32
}

func (renderer *Renderer) BindShader(shader *gosigl.Context) {
	renderer.shader = shader

	renderer.shader.UseProgram()

	renderer.uniforms["model"] = renderer.shader.GetUniform("model")
	renderer.uniforms["view"] = renderer.shader.GetUniform("view")
	renderer.uniforms["projection"] = renderer.shader.GetUniform("projection")
}

func (renderer *Renderer) StartFrame() {
	renderer.adapter.EnableBlend()
	renderer.adapter.EnableDepthTest()
	renderer.adapter.EnableCullFaceBack()

	renderer.shader.UseProgram()
}

func (renderer *Renderer) BindCamera(cam *entity.Camera) {
	model := cam.ModelMatrix()
	view := cam.ViewMatrix()
	proj := cam.ProjectionMatrix()

	renderer.adapter.SendUniformMat4(renderer.uniforms["projection"], &proj[0])
	renderer.adapter.SendUniformMat4(renderer.uniforms["view"], &view[0])
	renderer.adapter.SendUniformMat4(renderer.uniforms["model"], &model[0])
}

func (renderer *Renderer) DrawComposition(composition *render3d.Composition, mesh *gosigl.VertexObject, materials map[string]gosigl.TextureBindingId) {
	if mesh == nil {
		return
	}
	gosigl.BindMesh(mesh)
	renderer.adapter.Error()
	for _, matObj := range composition.MaterialMeshes() {
		if _, ok := materials[matObj.Material()]; ok {
			gosigl.BindTexture2D(gosigl.TextureSlot(0), materials[matObj.Material()])
		}

		renderer.adapter.DrawTriangleArray(matObj.Offset(), matObj.Length())
		renderer.adapter.Error()
	}
}

func NewRenderer(adapter graphics.Adapter) *Renderer {
	return &Renderer{
		adapter:  adapter,
		uniforms: make(map[string]int32),
	}
}
