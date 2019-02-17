package scene

import (
	"github.com/galaco/Lambda-Core/core/entity"
	lambdaModel "github.com/galaco/Lambda-Core/core/model"
	"github.com/galaco/Lambda/renderer/conversion"
	"github.com/galaco/Lambda/renderer/render3d"
	"github.com/galaco/Lambda/valve"
	"github.com/galaco/Lambda/valve/world"
	"github.com/galaco/gosigl"
	"github.com/go-gl/mathgl/mgl32"
)

type Scene struct {
	Solids map[int]*world.Solid
	SolidMeshes map[int]*lambdaModel.Model
	//RenderableSolids map[int][]*gosigl.VertexObject

	cameras map[*valve.Camera]*entity.Camera
	activeCamera *entity.Camera

	frameCompositor *render3d.Compositor
	frameComposed *render3d.Composition
	frameMesh *gosigl.VertexObject
}

func (scene *Scene) ActiveCamera() *entity.Camera {
	return scene.activeCamera
}

func (scene *Scene) ComposedMesh() *gosigl.VertexObject {
	if !scene.frameCompositor.IsOutdated() {
		return scene.frameMesh
	}

	if scene.frameMesh != nil {
		gosigl.DeleteMesh(scene.frameMesh)
	}

	scene.frameComposed = scene.frameCompositor.ComposeScene()
	sceneMesh := gosigl.NewMesh(scene.frameComposed.Vertices())
	gosigl.CreateVertexAttributeArrayBuffer(sceneMesh, scene.frameComposed.UVs(), 2)
	gosigl.CreateVertexAttributeElementArrayBuffer(sceneMesh, scene.frameComposed.Indices(), 1)
	gosigl.FinishMesh()

	scene.frameMesh = sceneMesh

	return scene.frameMesh
}

func (scene *Scene) AddSolid(solid *world.Solid) {
	scene.Solids[solid.Id] = solid

	model := conversion.SolidToModel(solid)
	scene.SolidMeshes[solid.Id] = model

	//scene.RenderableSolids[solid.Id] = make([]*gosigl.VertexObject, 0)

	for idx := range model.GetMeshes() {
		//vobj := gosigl.NewMesh(mesh.Vertices())
		//gosigl.CreateVertexAttributeArrayBuffer(vobj, mesh.UVs(), 2)
		//gosigl.FinishMesh()

		//scene.RenderableSolids[solid.Id] = append(scene.RenderableSolids[solid.Id], vobj)

		scene.frameCompositor.AddMesh(model.GetMeshes()[idx])
	}
}

func (scene *Scene) AddCamera(camera *valve.Camera) {
	scene.cameras[camera] = entity.NewCamera(90, 1024/768)
	scene.cameras[camera].Transform().Position = mgl32.Vec3{float32(camera.Position.X()), float32(camera.Position.Y()), float32(camera.Position.Z())}
	scene.cameras[camera].Transform().Rotation = mgl32.Vec3{float32(camera.Look.X()), float32(camera.Look.Y()), float32(camera.Look.Z())}

	if scene.activeCamera == nil {
		scene.activeCamera = scene.cameras[camera]
	}
}

func (scene *Scene) ChangeCamera(camera *valve.Camera) {
	if scene.cameras[camera] != nil {
		scene.activeCamera = scene.cameras[camera]
	}
}

func (scene *Scene) Close() {
	gosigl.DeleteMesh(scene.frameMesh)
}

func NewScene() *Scene {
	return &Scene{
		Solids: map[int]*world.Solid{},
		SolidMeshes: map[int]*lambdaModel.Model{},
		//RenderableSolids: map[int][]*gosigl.VertexObject{},
		cameras: map[*valve.Camera]*entity.Camera{},
		activeCamera: entity.NewCamera(90, 1024/768),
		frameCompositor: &render3d.Compositor{},
	}
}

