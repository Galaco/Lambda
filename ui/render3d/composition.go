package render3d

import "github.com/galaco/Lambda-Core/core/mesh"


type composition struct {
	mesh.IMesh
	materialCompositions []*compositionMesh

	indices []int32
}

// Compose constructs the indices information for the current state of the composition
func (comp *composition) Compose() {
	comp.indices = make([]int32, 0)
	for _,materialComposition := range comp.materialCompositions {
		materialComposition.GenerateIndicesList()
		comp.indices = append(comp.indices, materialComposition.indices...)
	}
}

// MaterialMeshes returns composed material information
func (comp *composition) MaterialMeshes() []*compositionMesh {
	return comp.materialCompositions
}

// Indices returns the indices of this compositions faces
func (comp *composition) Indices() []int32 {
	return comp.indices
}

// AddMesh
func (comp *composition) AddMesh(mat *compositionMesh) {
	comp.materialCompositions = append(comp.materialCompositions, mat)
}

// NewComposition returns a new composition.
func NewComposition() *composition {
	return &composition{}
}


type compositionMesh struct {
	texturePath string
	offset      int
	length      int

	indices []int32
}

// Indices returns all indices for vertices that use this material
func (texMesh *compositionMesh) Indices() []int32 {
	return texMesh.indices
}

// Indices returns the Offset for vertices that use this material
func (texMesh *compositionMesh) Offset() int32 {
	return int32(texMesh.offset)
}

// Indices returns the number for vertices that use this material
func (texMesh *compositionMesh) Length() int32 {
	return int32(texMesh.length)
}

// GenerateIndicesList generates the indices list from offset and length of composition vertex data.
func (texMesh *compositionMesh) GenerateIndicesList() {
	indices := make([]int32, 0)
	for i := texMesh.offset; i < texMesh.offset + texMesh.length; i++ {
		indices = append(indices, int32(i))
	}

	texMesh.indices = indices
}

// NewCompositionMesh returns a new compositionMesh
func NewCompositionMesh(texName string, offset int, length int) *compositionMesh {
	return &compositionMesh{
		texturePath: texName,
		length:      length,
		offset:      offset,
	}
}
