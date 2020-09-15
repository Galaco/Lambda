package model

import (
	"github.com/galaco/Lambda/pkg/lambda-core/mesh"
	"github.com/go-gl/mathgl/mgl32"
)

// ClusterLeaf represents a single cluster that contains the contents of
// all the leafs that are contained within it
type ClusterLeaf struct {
	Id          int16
	Faces       []mesh.Face
	StaticProps []*StaticProp
	DispFaces   []int
	Mins, Maxs  mgl32.Vec3
	Origin      mgl32.Vec3
}
