package renderer

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda/graphics"
	"github.com/galaco/Lambda/renderer/render3d"
	"github.com/galaco/gosigl"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	adapter graphics.Adapter

	shader *gosigl.Context
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

	gl.UniformMatrix4fv(renderer.uniforms["projection"], 1, false, &proj[0])
	gl.UniformMatrix4fv(renderer.uniforms["view"], 1, false, &view[0])
	gl.UniformMatrix4fv(renderer.uniforms["model"], 1, false, &model[0])
}

func (renderer *Renderer) DrawComposition(composition *render3d.Composition, mesh *gosigl.VertexObject) {
	//for _,solid := range scene.RenderableSolids {
	//	win.renderer.DrawSolid(solid)
	//	win.adapter.Error()
	//}
	if mesh == nil {
		return
	}
	gosigl.BindMesh(mesh)
	gl.DrawArrays(gl.TRIANGLES, 0, int32( len(composition.Vertices()) / 3))
	renderer.adapter.Error()
	//for _,matObj := range composition.MaterialMeshes() {
	//	gl.DrawRangeElements(
	//		gl.TRIANGLES,
	//		uint32(matObj.Offset()),
	//		uint32(matObj.Offset() + matObj.Length()),
	//		matObj.Length(), gl.UNSIGNED_INT,
	//		gl.Ptr(composition.Indices()))
	//	renderer.adapter.Error()
	//}
	//
}

func (renderer *Renderer) DrawSolid(model []*gosigl.VertexObject) {
	for _,vobj := range model {
		gosigl.BindMesh(vobj)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
	}
}

func NewRenderer(adapter graphics.Adapter) *Renderer {
	return &Renderer{
		adapter: adapter,
		uniforms: make(map[string]int32),
	}
}