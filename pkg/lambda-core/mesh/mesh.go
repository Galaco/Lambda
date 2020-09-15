package mesh

import (
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/mesh/util"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
)

// Mesh
type Mesh struct {
	vertices            []float32
	normals             []float32
	uvs                 []float32
	tangents            []float32
	lightmapCoordinates []float32

	material material.IMaterial
	lightmap texture.ITexture
}

// AddVertex
func (mesh *Mesh) AddVertex(vertex ...float32) {
	mesh.vertices = append(mesh.vertices, vertex...)
}

// AddNormal
func (mesh *Mesh) AddNormal(normal ...float32) {
	mesh.normals = append(mesh.normals, normal...)
}

// AddUV
func (mesh *Mesh) AddUV(uv ...float32) {
	mesh.uvs = append(mesh.uvs, uv...)
}

// AddTangent
func (mesh *Mesh) AddTangent(tangent ...float32) {
	mesh.tangents = append(mesh.tangents, tangent...)
}

// AddLightmapCoordinate
func (mesh *Mesh) AddLightmapCoordinate(uv ...float32) {
	mesh.lightmapCoordinates = append(mesh.lightmapCoordinates, uv...)
}

// Vertices
func (mesh *Mesh) Vertices() []float32 {
	return mesh.vertices
}

// Normals
func (mesh *Mesh) Normals() []float32 {
	return mesh.normals
}

// UVs
func (mesh *Mesh) UVs() []float32 {
	return mesh.uvs
}

// Tangents
func (mesh *Mesh) Tangents() []float32 {
	return mesh.tangents
}

// LightmapCoordinates
func (mesh *Mesh) LightmapCoordinates() []float32 {
	// use standard uvs if there is no lightmap. Not ideal,
	// but there MUST be UVs, but they'll be ignored anyway if there is no
	// lightmap
	if len(mesh.lightmapCoordinates) == 0 {
		return mesh.UVs()
	}
	return mesh.lightmapCoordinates
}

// Material
func (mesh *Mesh) Material() material.IMaterial {
	return mesh.material
}

// SetMaterial
func (mesh *Mesh) SetMaterial(mat material.IMaterial) {
	mesh.material = mat
}

// Lightmap
func (mesh *Mesh) Lightmap() texture.ITexture {
	return mesh.lightmap
}

//SetLightmap
func (mesh *Mesh) SetLightmap(mat texture.ITexture) {
	mesh.lightmap = mat
}

// GenerateTangents
func (mesh *Mesh) GenerateTangents() {
	mesh.tangents = util.GenerateTangents(mesh.Vertices(), mesh.Normals(), mesh.UVs())
}

// NewMesh
func NewMesh() *Mesh {
	return &Mesh{}
}
