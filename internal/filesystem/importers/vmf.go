package importers

import (
	"errors"
	"github.com/galaco/Lambda/pkg/valve"
	"github.com/galaco/Lambda/pkg/valve/world"
	"github.com/galaco/source-tools-common/entity"
	"github.com/galaco/vmf"
	"os"
	"strconv"
)

// Imports a Vmf file and returns a Vmf Model
type VmfImporter struct {
}

// Public loader function to open and import a vmf file
// Will error out if the file is malformed or cannot be opened
func (importer *VmfImporter) LoadVmf(filepath string) (*valve.Vmf, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := vmf.NewReader(file)
	importable, err := reader.Read()

	if err != nil {
		return nil, err
	}

	// Create models for different vmf properties
	versionInfo, err := importer.loadVersionInfo(&importable.VersionInfo)
	if err != nil || versionInfo == nil {
		return nil, err
	}
	visGroups, err := importer.loadVisGroups(&importable.VisGroup)
	if err != nil || visGroups == nil {
		return nil, err
	}
	worldspawn, err := importer.loadWorld(&importable.World)
	if err != nil || worldspawn == nil {
		return nil, err
	}

	cameras, err := importer.loadCameras(&importable.Cameras)
	if err != nil || cameras == nil {
		return nil, err
	}

	entities := importer.loadEntities(&importable.Entities)

	return valve.NewVmf(versionInfo, visGroups, worldspawn, entities, cameras), nil
}

// loadVersionInfo creates a VersionInfo model
// from the versioninfo vmf block
func (importer *VmfImporter) loadVersionInfo(root *vmf.Node) (*valve.VersionInfo, error) {
	if root == nil {
		return nil, errors.New("missing versioninfo")
	}
	editorVersion, err := strconv.ParseInt(root.GetProperty("editorversion"), 10, 64)
	if err != nil {
		return nil, err
	}
	editorBuild, err := strconv.ParseInt(root.GetProperty("editorbuild"), 10, 64)
	if err != nil {
		return nil, err
	}
	mapVersion, err := strconv.ParseInt(root.GetProperty("mapversion"), 10, 64)
	if err != nil {
		return nil, err
	}
	formatVersion, err := strconv.ParseInt(root.GetProperty("formatversion"), 10, 64)
	if err != nil {
		return nil, err
	}
	prefab := false
	if root.GetProperty("prefab") != "0" {
		prefab = true
	}

	return valve.NewVersionInfo(int(editorVersion), int(editorBuild), int(mapVersion), int(formatVersion), prefab), nil
}

// loadVisgroups loads all visgroup information from the
// visgroups block of a vmf
func (importer *VmfImporter) loadVisGroups(root *vmf.Node) (*valve.VisGroups, error) {
	return &valve.VisGroups{}, nil
}

func (importer *VmfImporter) loadWorld(root *vmf.Node) (*world.World, error) {
	solidNodes := root.GetChildrenByKey("solid")
	worldSpawn := entity.FromVmfNode(root)

	solids := make([]world.Solid, len(solidNodes))
	for idx, solidNode := range solidNodes {
		solid, err := importer.loadSolid(&solidNode)
		if err != nil {
			return nil, err
		}
		solids[idx] = *solid
	}

	return world.NewWorld(&worldSpawn, solids), nil
}

// loadSolid takes a vmf node tree that represents a solid and turns
// it into a properly defind model structure for the solid with
// proper type definitions.
func (importer *VmfImporter) loadSolid(node *vmf.Node) (*world.Solid, error) {
	id, err := strconv.ParseInt(node.GetProperty("id"), 10, 64)
	if err != nil {
		return world.NewSolid(-1, nil, nil), err
	}
	sideNodes := node.GetChildrenByKey("side")
	// Create sides for solid
	sides := make([]world.Side, len(sideNodes))
	for idx, sideNode := range sideNodes {
		var id int64
		var plane world.Plane
		var material string
		var u, v world.UVTransform
		var rotation, lmScale float64
		var smoothing bool

		id, err := strconv.ParseInt(sideNode.GetProperty("id"), 10, 64)
		if err != nil {
			return nil, err
		}
		plane = *world.NewPlaneFromString(sideNode.GetProperty("plane"))

		material = sideNode.GetProperty("material")

		u = *world.NewUVTransformFromString(sideNode.GetProperty("uaxis"))
		v = *world.NewUVTransformFromString(sideNode.GetProperty("vaxis"))

		rotation, err = strconv.ParseFloat(sideNode.GetProperty("rotation"), 64)
		if err != nil {
			return nil, err
		}
		lmScale, err = strconv.ParseFloat(sideNode.GetProperty("lightmapscale"), 64)
		if err != nil {
			return nil, err
		}
		smoothing, err = strconv.ParseBool(sideNode.GetProperty("smoothing_groups"))
		if err != nil {
			return nil, err
		}

		sides[idx] = *world.NewSide(int(id), plane, material, u, v, rotation, lmScale, smoothing)
	}

	return world.NewSolid(int(id), sides, nil), nil
}

// loadEntities creates models from the entity data block
// from a vmf
func (importer *VmfImporter) loadEntities(node *vmf.Node) *entity.List {
	entities := entity.FromVmfNodeTree(*node)

	return &entities
}

// loadCameras creates cameras from the vmf camera list
func (importer *VmfImporter) loadCameras(node *vmf.Node) (*valve.Cameras, error) {
	activeCamProp := node.GetProperty("activecamera")
	activeCamIdx, _ := strconv.ParseInt(activeCamProp, 10, 32)

	cameras := make([]valve.Camera, 0)

	cameraProps := node.GetChildrenByKey("camera")
	for _, camProp := range cameraProps {
		pos := camProp.GetProperty("position")
		look := camProp.GetProperty("look")

		cameras = append(cameras, *valve.NewCamera(world.NewVec3FromString(pos), world.NewVec3FromString(look)))
	}

	return valve.NewCameras(int(activeCamIdx), cameras), nil
}

func NewVmfImporter() *VmfImporter {
	return &VmfImporter{}
}
