package render3d

import (
	"github.com/galaco/Lambda/pkg/lambda-core/mesh"
)

type Composition struct {
	mesh.Mesh
	materialCompositions []*compositionMesh

	indices []uint32
}

// Compose constructs the indices information for the current state of the Composition
func (comp *Composition) Compose() {
	comp.indices = make([]uint32, 0)
	for _, materialComposition := range comp.materialCompositions {
		materialComposition.GenerateIndicesList()
		comp.indices = append(comp.indices, materialComposition.indices...)
	}
}

// MaterialMeshes returns composed material information
func (comp *Composition) MaterialMeshes() []*compositionMesh {
	return comp.materialCompositions
}

// Indices returns the indices of this compositions faces
func (comp *Composition) Indices() []uint32 {
	return comp.indices
}

// AddMesh
func (comp *Composition) AddMesh(mat *compositionMesh) {
	comp.materialCompositions = append(comp.materialCompositions, mat)
}

// NewComposition returns a new Composition.
func NewComposition() *Composition {
	return &Composition{}
}

type compositionMesh struct {
	texturePath string
	offset      int
	length      int

	indices []uint32
}

func (texMesh *compositionMesh) Material() string {
	return texMesh.texturePath
}

// Indices returns all indices for vertices that use this material
func (texMesh *compositionMesh) Indices() []uint32 {
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

// GenerateIndicesList generates the indices list from offset and length of Composition vertex data.
func (texMesh *compositionMesh) GenerateIndicesList() {
	indices := make([]uint32, 0)
	for i := texMesh.offset; i < texMesh.offset+texMesh.length; i++ {
		indices = append(indices, uint32(i))
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
