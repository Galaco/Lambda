package mesh

import (
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
)

// IMesh Generic Mesh interface
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IMesh interface {
	// AddVertex
	AddVertex(...float32)
	// AddNormal
	AddNormal(...float32)
	// AddUV
	AddUV(...float32)
	// AddLightmapCoordinate
	AddLightmapCoordinate(...float32)
	// GenerateTangents
	GenerateTangents()

	// Vertices
	Vertices() []float32
	// Normals
	Normals() []float32
	// UVs
	UVs() []float32
	// Tangents
	Tangents() []float32
	// LightmapCoordinates
	LightmapCoordinates() []float32

	// Material
	Material() material.IMaterial
	// SetMaterial
	SetMaterial(material.IMaterial)
	// Lightmap
	Lightmap() texture.ITexture
	// SetLightmap
	SetLightmap(texture.ITexture)
}
