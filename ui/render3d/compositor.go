package render3d

import (
	"github.com/galaco/Lambda-Core/core/mesh"
	"github.com/galaco/Lambda-Core/core/model"
)

// Compositor is a struct that provides a mechanism to compose 1 or more models into a single renderable set of data,
// indexed by material.
// This is super handy for reducing draw calls down a bunch.
// A resultant composition should result in a single set of vertex data + 1 pair of index offset+length info per material
// referenced by all models composed.
type Compositor struct {
	models []*model.Model
}

// AddModel adds a new model to be composed.
func (compositor *Compositor) AddModel(m *model.Model) {
	compositor.models = append(compositor.models, m)
}

// ComposeScene builds a sceneComposition mesh for rendering
func (compositor *Compositor) ComposeScene() *composition {
	texMappings := map[string][]mesh.IMesh{}

	// Step 1. Map meshes into contiguous groups by texture
	for _,mod := range compositor.models {
		for _,m := range mod.GetMeshes() {
			if _, ok := texMappings[m.GetMaterial().GetFilePath()]; !ok {
				texMappings[m.GetMaterial().GetFilePath()] = make([]mesh.IMesh, 0)
			}

			texMappings[m.GetMaterial().GetFilePath()] = append(texMappings[m.GetMaterial().GetFilePath()], m)
		}
	}

	// Step 2. Construct a single vertex object composition ordered by material
	sceneComposition := NewComposition()
	vertCount := 0
	for key,texMesh := range texMappings {
		// @TODO verify if this is the vertex offset of the actual array offset (vertexOffset * 3)
		matVertOffset := vertCount
		matVertCount := 0
		for _,sMesh := range texMesh {
			sceneComposition.AddVertex(sMesh.Vertices()...)
			sceneComposition.AddNormal(sMesh.Normals()...)
			sceneComposition.AddUV(sMesh.UVs()...)

			vertCount += len(sMesh.Vertices()) / 3
		}

		sceneComposition.GenerateTangents()
		sceneComposition.AddMesh(NewCompositionMesh(key, matVertOffset, matVertCount - matVertOffset))
	}

	// Step 3. Generate indices from composed materials
	sceneComposition.Compose()

	return sceneComposition
}