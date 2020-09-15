package mesh

import (
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
)

// Face
type Face struct {
	offset   int32
	length   int32
	material material.IMaterial
	lightmap *texture.Lightmap
}

// Offset
func (face *Face) Offset() int32 {
	return face.offset
}

// Length
func (face *Face) Length() int32 {
	return face.length
}

// IsLightmapped
func (face *Face) IsLightmapped() bool {
	return face.Lightmap() != nil
}

// AddMaterial
func (face *Face) AddMaterial(mat material.IMaterial) {
	face.material = mat
}

// AddLightmap
func (face *Face) AddLightmap(lightmap *texture.Lightmap) {
	face.lightmap = lightmap
}

// Material
func (face *Face) Material() material.IMaterial {
	return face.material
}

// Lightmap
func (face *Face) Lightmap() *texture.Lightmap {
	return face.lightmap
}

// NewFace
func NewFace(offset int32, length int32, mat texture.ITexture, lightmap *texture.Lightmap) Face {
	return Face{
		offset:   offset,
		length:   length,
		material: mat,
		lightmap: lightmap,
	}
}
