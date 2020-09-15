package material

import (
	"github.com/galaco/Lambda/pkg/lambda-core/lib/stringtable"
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/resource"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/golang-source-engine/filesystem"
	stringtableLib "github.com/golang-source-engine/stringtable"
)

// LoadMaterials is the base bsp material loader function.
// All bsp materials should be loaded by this function.
// Note that this covers bsp referenced materials only, model & entity
// materials are loaded mostly ad-hoc.
func LoadMaterials(fs *filesystem.FileSystem, stringData *lumps.TexDataStringData, stringTable *lumps.TexDataStringTable, texInfos *[]texinfo.TexInfo) *stringtableLib.StringTable {
	materialStringTable := stringtable.NewTable(stringData, stringTable)
	LoadErrorMaterial()

	for _,path := range stringtable.SortUnique(materialStringTable, texInfos) {
		_,_ = LoadMaterialFromFilesystem(fs, path)
	}

	return materialStringTable
}

func LoadMaterialFromFilesystem(fs *filesystem.FileSystem, filePath string) (material.IMaterial, error) {
	if resource.Manager().HasMaterial(filePath) {
		return resource.Manager().Material(filePath), nil
	}
	props, err := LoadVmtFromFilesystem(fs, filePath)
	if err != nil {
		return resource.Manager().Material(resource.Manager().ErrorTextureName()), err
	}
	mat := material.NewMaterial(filePath)
	mat.BaseTextureName = props.BaseTexture
	mat.BumpMapName = props.Bumpmap

	if len(mat.BaseTextureName) > 0 {
		if resource.Manager().HasTexture(mat.BaseTextureName) {
			mat.Textures.Albedo = resource.Manager().Texture(mat.BaseTextureName)
		} else {
			albedo, err := LoadVtfFromFilesystem(fs, mat.BaseTextureName)
			if err != nil {
				mat.Textures.Albedo = resource.Manager().Texture(resource.Manager().ErrorTextureName())
				return mat, err
			}
			resource.Manager().AddTexture(albedo)
			mat.Textures.Albedo = albedo
		}
	} else {
		mat.Textures.Albedo = resource.Manager().Texture(resource.Manager().ErrorTextureName())
	}
	resource.Manager().AddMaterial(mat)

	return mat, nil
}

// LoadErrorMaterial ensures that the error material has been loaded
func LoadErrorMaterial() {
	name := resource.Manager().ErrorTextureName()

	if resource.Manager().HasMaterial(name) {
		return
	}

	// Ensure that error texture is available
	resource.Manager().AddTexture(texture.NewError(name))
	errorMat := material.NewMaterial(name)
	errorMat.Textures.Albedo = resource.Manager().Texture(name).(texture.ITexture)
	resource.Manager().AddMaterial(errorMat)
}
