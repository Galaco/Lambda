package loader

import (
	material2 "github.com/galaco/Lambda/pkg/lambda-core/loader/material"
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/mesh/primitive"
	"github.com/galaco/Lambda/pkg/lambda-core/model"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"github.com/golang-source-engine/filesystem"
)

const skyboxRootDir = "skybox/"

// LoadSky loads the skymaterial cubemap.
// The materialname is normally obtained from the worldspawn entity
func LoadSky(materialName string, fs *filesystem.FileSystem) *model.Model {
	sky := model.NewModel(materialName)

	mats := make([]material.IMaterial, 6)

	mats[0],_ = material2.LoadMaterialFromFilesystem(fs, skyboxRootDir+materialName+"up.vmt")
	mats[1],_ = material2.LoadMaterialFromFilesystem(fs, skyboxRootDir+materialName+"dn.vmt")
	mats[2],_ = material2.LoadMaterialFromFilesystem(fs, skyboxRootDir+materialName+"lf.vmt")
	mats[3],_ = material2.LoadMaterialFromFilesystem(fs, skyboxRootDir+materialName+"rt.vmt")
	mats[4],_ = material2.LoadMaterialFromFilesystem(fs, skyboxRootDir+materialName+"ft.vmt")
	mats[5],_ = material2.LoadMaterialFromFilesystem(fs, skyboxRootDir+materialName+"bk.vmt")

	texs := make([]texture.ITexture, 6)
	for i := 0; i < 6; i++ {
		texs[i] = mats[i].(*material.Material).Textures.Albedo
	}

	sky.AddMesh(primitive.NewCube())

	sky.Meshes()[0].SetMaterial(texture.NewCubemap(texs))

	return sky
}
