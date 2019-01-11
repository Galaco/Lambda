package importers

import (
	"errors"
	"github.com/galaco/Lambda/core/model"
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