package importers

import (
	"errors"
	"github.com/galaco/source-tools-common/entity"
	"github.com/galaco/Lambda/core/model"
	"github.com/galaco/Lambda/core/model/world"
	"github.com/galaco/vmf"
	"os"
	"strconv"
)

type VmfImporter struct {
}

func (importer *VmfImporter) LoadVmf(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := vmf.NewReader(file)
	importable, err := reader.Read()

	if err != nil {
		return err
	}

	versionInfo, err := importer.loadVersionInfo(&importable.VersionInfo)
	if err != nil || versionInfo == nil {
		return err
	}
	visGroups,err := importer.loadVisGroups(&importable.VisGroup)
	if err != nil || visGroups == nil {
		return err
	}
	worldspawn,err := importer.loadWorld(&importable.World)
	if err != nil || worldspawn == nil {
		return err
	}


	return nil
}

func (importer *VmfImporter) loadVersionInfo(root *vmf.Node) (*model.VersionInfo, error) {
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

	return model.NewVersionInfo(int(editorVersion), int(editorBuild), int(mapVersion), int(formatVersion), prefab), nil
}

func (importer *VmfImporter) loadVisGroups(root *vmf.Node) (*model.VisGroups, error){
	return &model.VisGroups{}, nil
}

func (importer *VmfImporter) loadWorld(root *vmf.Node) (*world.World, error){
	solidNodes := root.GetChildrenByKey("solid")
	worldSpawn := entity.FromVmfNode(root)

	solids := make([]world.Solid, len(solidNodes))
	for idx,solidNode := range solidNodes {
		solid,err := importer.loadSolid(&solidNode)
		if err != nil {
			return nil,err
		}
		solids[idx] = *solid
	}

	return world.NewWorld(&worldSpawn, solids), nil
}

// loadSolid takes a vmf node tree that represents a solid and turns
// it into a properly defind model structure for the solid with
// proper type definitions.
func (importer *VmfImporter) loadSolid(node *vmf.Node) (*world.Solid, error) {
	id,err := strconv.ParseInt(node.GetProperty("id"), 10, 64)
	if err != nil {
		return world.NewSolid(-1, nil, nil),err
	}
	sideNodes := node.GetChildrenByKey("side")
	sides := make([]world.Side, len(sideNodes))
	for idx, sideNode := range sideNodes {
		var id int64
		var plane world.Plane
		var material string
		var u, v world.UVTransform
		var rotation, lmScale float64
		var smoothing bool

		id,err := strconv.ParseInt(sideNode.GetProperty("id"), 10, 64)
		if err != nil {
			return nil, err
		}
		plane = *world.NewPlaneFromString(sideNode.GetProperty("plane"))

		material = sideNode.GetProperty("material")

		u = *world.NewUVTransformFromString(sideNode.GetProperty("uaxis"))
		v = *world.NewUVTransformFromString(sideNode.GetProperty("vaxis"))

		rotation,err = strconv.ParseFloat(sideNode.GetProperty("rotation"), 64)
		if err != nil {
			return nil, err
		}
		lmScale,err = strconv.ParseFloat(sideNode.GetProperty("lightmapscale"), 64)
		if err != nil {
			return nil, err
		}
		smoothing,err = strconv.ParseBool(sideNode.GetProperty("smoothing_groups"))
		if err != nil {
			return nil, err
		}

		sides[idx] = *world.NewSide(int(id), plane, material, u, v, rotation, lmScale, smoothing)
	}

	return world.NewSolid(int(id), sides, nil), nil
}