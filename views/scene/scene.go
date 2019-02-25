package scene

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda-Core/core/loader/material"
	material2 "github.com/galaco/Lambda-Core/core/material"
	lambdaModel "github.com/galaco/Lambda-Core/core/model"
	"github.com/galaco/Lambda-Core/core/resource"
	"github.com/galaco/Lambda/internal/renderer/conversion"
	"github.com/galaco/Lambda/internal/renderer/render3d"
	"github.com/galaco/Lambda/internal/valve"
	"github.com/galaco/Lambda/internal/valve/world"
	"github.com/galaco/gosigl"
	"github.com/go-gl/mathgl/mgl32"
)

type Scene struct {
	Solids      map[int]*world.Solid
	SolidMeshes map[int]*lambdaModel.Model

	cameras      map[*valve.Camera]*entity.Camera
	activeCamera *entity.Camera

	frameCompositor *render3d.Compositor
	frameComposed   *render3d.Composition
	frameMesh       *gosigl.VertexObject
	frameMaterials  map[string]gosigl.TextureBindingId
}

func (scene *Scene) ActiveCamera() *entity.Camera {
	return scene.activeCamera
}

func (scene *Scene) Composition() *gosigl.VertexObject {
	return scene.frameMesh
}

func (scene *Scene) CompositionMaterials() map[string]gosigl.TextureBindingId {
	return scene.frameMaterials
}

func (scene *Scene) RecomposeScene(fs *filesystem.FileSystem) *gosigl.VertexObject {
	if scene.frameMesh != nil {
		gosigl.DeleteMesh(scene.frameMesh)
	}

	scene.frameComposed = scene.frameCompositor.ComposeScene()
	sceneMesh := gosigl.NewMesh(scene.frameComposed.Vertices())
	gosigl.CreateVertexAttributeArrayBuffer(sceneMesh, scene.frameComposed.Normals(), 3)
	gosigl.CreateVertexAttributeArrayBuffer(sceneMesh, scene.frameComposed.UVs(), 2)
	gosigl.CreateVertexAttributeArrayBuffer(sceneMesh, scene.frameComposed.Tangents(), 3)
	gosigl.FinishMesh()

	scene.frameMesh = sceneMesh

	// @TODO This exists only to prove materials work, and that UVs are correct!
	// Ensure materials are ready
	material.LoadErrorMaterial()
	for _, matComp := range scene.frameComposed.MaterialMeshes() {
		if _, ok := scene.frameMaterials[matComp.Material()]; ok {
			continue
		}
		baseMat := material.LoadSingleMaterial(matComp.Material()+".vmt", fs)
		if baseMat == nil {
			baseMat = resource.Manager().GetMaterial(resource.Manager().ErrorTextureName())
		}
		mat := baseMat.(*material2.Material)

		scene.frameMaterials[matComp.Material()] = gosigl.CreateTexture2D(
			gosigl.TextureSlot(0),
			mat.Textures.Albedo.Width(),
			mat.Textures.Albedo.Height(),
			mat.Textures.Albedo.PixelDataForFrame(0),
			gosigl.PixelFormat(glTextureFormatFromVtfFormat(mat.Textures.Albedo.Format())),
			false)
	}

	return scene.frameMesh
}

func (scene *Scene) AddSolid(solid *world.Solid) {
	scene.Solids[solid.Id] = solid

	model := conversion.SolidToModel(solid)
	scene.SolidMeshes[solid.Id] = model

	for idx := range model.GetMeshes() {
		scene.frameCompositor.AddMesh(model.GetMeshes()[idx])
	}
}

func (scene *Scene) AddCamera(camera *valve.Camera) {
	scene.cameras[camera] = entity.NewCamera(70, 1024/768)
	scene.cameras[camera].Transform().Position = mgl32.Vec3{float32(camera.Position.X()), float32(camera.Position.Y()), float32(camera.Position.Z())}
	scene.cameras[camera].Transform().Rotation = mgl32.Vec3{
		mgl32.DegToRad(float32(camera.Look.X())),
		mgl32.DegToRad(float32(camera.Look.Y())),
		mgl32.DegToRad(float32(camera.Look.Z()))}

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

	for _, m := range scene.frameMaterials {
		gosigl.DeleteTextures(m)
	}
}

func NewScene() *Scene {
	return &Scene{
		Solids:          map[int]*world.Solid{},
		SolidMeshes:     map[int]*lambdaModel.Model{},
		cameras:         map[*valve.Camera]*entity.Camera{},
		activeCamera:    entity.NewCamera(70, 1024/768),
		frameCompositor: &render3d.Compositor{},
		frameMaterials:  map[string]gosigl.TextureBindingId{},
	}
}

// getGLTextureFormat swap vtf format to openGL format
func glTextureFormatFromVtfFormat(vtfFormat uint32) gosigl.PixelFormat {
	switch vtfFormat {
	case 0:
		return gosigl.RGBA
	case 2:
		return gosigl.RGB
	case 3:
		return gosigl.BGR
	case 12:
		return gosigl.BGRA
	case 13:
		return gosigl.DXT1
	case 14:
		return gosigl.DXT3
	case 15:
		return gosigl.DXT5
	default:
		return gosigl.RGB
	}
}
