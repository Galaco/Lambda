package conversion

import (
	"fmt"
	lambdaMesh "github.com/galaco/Lambda-Core/core/mesh"
	lambdaModel "github.com/galaco/Lambda-Core/core/model"
	"github.com/galaco/Lambda/valve/world"
	"github.com/galaco/gosigl"
	"math"
)

func SolidToModel(solid *world.Solid) *lambdaModel.Model{
	meshes := make([]lambdaMesh.IMesh, 0)

	for _,side := range solid.Sides {
		meshes = append(meshes, SideToMesh(&side))
	}

	return lambdaModel.NewModel(fmt.Sprintf("solid_%d", solid.Id), meshes...)
}

func SideToMesh(side *world.Side) lambdaMesh.IMesh {
	mesh := lambdaMesh.NewMesh()

	// Vertices
	verts := make([]float32, 0)
	{
		// a plane represents 3 vertices- bottom-left, top-left and top-right
		// Triangle 1
		verts = append(verts, float32(side.Plane[0].X()), float32(side.Plane[0].Y()), float32(side.Plane[0].Z()))
		verts = append(verts, float32(side.Plane[1].X()), float32(side.Plane[1].Y()), float32(side.Plane[1].Z()))
		verts = append(verts, float32(side.Plane[2].X()), float32(side.Plane[2].Y()), float32(side.Plane[2].Z()))
		// Triangle 2
		verts = append(verts, float32(side.Plane[0].X()), float32(side.Plane[0].Y()), float32(side.Plane[0].Z()))
		verts = append(verts, float32(side.Plane[2].X()), float32(side.Plane[2].Y()), float32(side.Plane[2].Z()))

		// Compute new vertex
		vert4 := side.Plane[2].Sub(side.Plane[1].Sub(side.Plane[0]))
		verts = append(verts, float32(vert4.X()), float32(vert4.Y()), 0)



		// @TODO
		// This is for debugging! It flattens the solid onto the X/Y plane
		for idx, v := range verts {
			if (idx + 1) % 3 == 0 {
				verts[idx] = 0
				continue
			}
			perc := math.Abs(float64((float32(v) / 16384.0))) - 0.5
			verts[idx] = float32(perc)
		}

		mesh.AddVertex(verts...)
		gosigl.FinishMesh()
	}

	// Material


	// Texture coordinates
	{
		for i := 0; i < len(verts); i+=3 {
			mesh.AddUV(uvForVertex(verts[i:i+3], &side.UAxis, &side.VAxis, 32, 32)...)
		}
	}

	return mesh
}

func uvForVertex(vertex []float32, u *world.UVTransform, v *world.UVTransform, width int, height int) (uvs []float32) {
	cu := (float32(u.Transform[0]) * vertex[0]) +
		(float32(u.Transform[1]) * vertex[1]) +
		(float32(u.Transform[2]) * vertex[2]) +
		float32(u.Scale) / float32(width)

	cv := (float32(v.Transform[0]) * vertex[0]) +
		(float32(v.Transform[1]) * vertex[1]) +
		(float32(v.Transform[2]) * vertex[2]) +
		float32(v.Scale) / float32(height)

	return []float32{cu, cv}
}