package render3d

import (
	"github.com/galaco/Lambda-Core/core/mesh"
)

// Compositor is a struct that provides a mechanism to compose 1 or more models into a single renderable set of data,
// indexed by material.
// This is super handy for reducing draw calls down a bunch.
// A resultant Composition should result in a single set of vertex data + 1 pair of index offset+length info per material
// referenced by all models composed.
type Compositor struct {
	meshes []mesh.IMesh

	isOutdated bool
}

// AddModel adds a new model to be composed.
func (compositor *Compositor) AddMesh(m mesh.IMesh) {
	compositor.meshes = append(compositor.meshes, m)
	compositor.isOutdated = true
}

func (compositor *Compositor) IsOutdated() bool {
	return compositor.isOutdated
}

// ComposeScene builds a sceneComposition mesh for rendering
func (compositor *Compositor) ComposeScene() *Composition {
	compositor.isOutdated = false
	texMappings := map[string][]mesh.IMesh{}

	// Step 1. Map meshes into contiguous groups by texture
	for _,m := range compositor.meshes {
		if _, ok := texMappings[m.GetMaterial().GetFilePath()]; !ok {
			texMappings[m.GetMaterial().GetFilePath()] = make([]mesh.IMesh, 0)
		}

		texMappings[m.GetMaterial().GetFilePath()] = append(texMappings[m.GetMaterial().GetFilePath()], m)
	}

	// Step 2. Construct a single vertex object Composition ordered by material
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

			matVertCount += len(sMesh.Vertices()) / 3
		}

		sceneComposition.GenerateTangents()
		sceneComposition.AddMesh(NewCompositionMesh(key, matVertOffset, matVertCount))
		vertCount += matVertCount
	}

	// Step 3. Generate indices from composed materials
	sceneComposition.Compose()

	return sceneComposition
}