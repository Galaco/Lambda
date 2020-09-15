package model

import "github.com/galaco/Lambda/pkg/lambda-core/mesh"

// Bsp is a specialised model that represents an entire bsp map
// It is represented by a single mesh and a series of visiblity structures
// that dictate what can and can't be seen from a given point
type Bsp struct {
	internalMesh mesh.IMesh

	defaultClusterLeaf  ClusterLeaf
	clusterLeafs        []ClusterLeaf
	visibleClusterLeafs []*ClusterLeaf
}

// Mesh returns Bsp Mesh
func (bsp *Bsp) Mesh() mesh.IMesh {
	return bsp.internalMesh
}

// ClusterLeafs returns all ClusterLeafs
func (bsp *Bsp) ClusterLeafs() []ClusterLeaf {
	return bsp.clusterLeafs
}

// VisibleClusterLeafs returns clusterleafs that are known to
// be visible. This is not calculated here, only stored for faster reference
func (bsp *Bsp) VisibleClusterLeafs() []*ClusterLeaf {
	return bsp.visibleClusterLeafs
}

// SetClusterLeafs set the computed cluster leafs for a Bsp
func (bsp *Bsp) SetClusterLeafs(clusterLeafs []ClusterLeaf) {
	bsp.clusterLeafs = clusterLeafs
}

// SetVisibleClusters update the visible ClusterLeafs
func (bsp *Bsp) SetVisibleClusters(clusterLeafs []*ClusterLeaf) {
	bsp.visibleClusterLeafs = clusterLeafs
}

// SetDefaultCluster
func (bsp *Bsp) SetDefaultCluster(dispFaces ClusterLeaf) {
	bsp.defaultClusterLeaf = dispFaces
}

// DefaultCluster
func (bsp *Bsp) DefaultCluster() *ClusterLeaf {
	return &bsp.defaultClusterLeaf
}

// NewBsp returns a new bsp
func NewBsp(refMesh *mesh.Mesh) *Bsp {
	return &Bsp{
		internalMesh: refMesh,
	}
}
