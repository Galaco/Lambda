package exporters

import (
	"github.com/galaco/Lambda/valve"
	"github.com/galaco/Lambda/valve/world"
	"strings"
)

type VmfExporter struct {
}

func (exporter *VmfExporter) Export(model *valve.Vmf) (export string, err error) {
	model.VersionInfo()
	versionInfo, err := exporter.exportVersionInfo(model.VersionInfo())
	if err != nil {
		return export, err
	}
	visGroups, err := exporter.exportVisGroups(model.Visgroups())
	if err != nil {
		return export, err
	}
	viewSettings, err := exporter.exportViewSettings(model.ViewSettings())
	if err != nil {
		return export, err
	}
	worldspawn, err := exporter.exportWorldspawn(model.Worldspawn())
	if err != nil {
		return export, err
	}

	export = versionInfo + visGroups + viewSettings + worldspawn

	return export, err
}

func (exporter *VmfExporter) exportVersionInfo(model *valve.VersionInfo) (export string, err error) {
	export = templateVmfVersionInfo

	export = strings.Replace(export, "{editorversion}", intToString(model.EditorVersion), 1)
	export = strings.Replace(export, "{editorbuild}", intToString(model.EditorBuild), 1)
	export = strings.Replace(export, "{mapversion}", intToString(model.MapVersion), 1)
	export = strings.Replace(export, "{formatversion}", intToString(model.FormatVersion), 1)
	export = strings.Replace(export, "{prefab}", boolToString(model.Prefab), 1)

	return export, err
}

func (exporter *VmfExporter) exportVisGroups(model *valve.VisGroups) (export string, err error) {
	export = templateVmfVisgroups
	return export, err
}

func (exporter *VmfExporter) exportViewSettings(model *valve.ViewSettings) (export string, err error) {
	export = templateVmfViewSettings

	export = strings.Replace(export, "{bSnapToGrid}", boolToString(model.SnapToGrid), 1)
	export = strings.Replace(export, "{bShowGrid}", boolToString(model.ShowGrid), 1)
	export = strings.Replace(export, "{bShowLogicalGrid}", boolToString(model.ShowLogicalGrid), 1)
	export = strings.Replace(export, "{nGridSpacing}", intToString(model.GridSpacing), 1)
	export = strings.Replace(export, "{bShow3DGrid}", boolToString(model.Show3DGrid), 1)

	return export, err
}

func (exporter *VmfExporter) exportWorldspawn(model *world.World) (export string, err error) {
	export = templateVmfWorld

	kv := model.Keyvalues.EPairs
	for kv != nil {
		if kv.Key == "solid" {
			kv = kv.Next
			continue
		}
		export += templateVmfKeyValue
		export = strings.Replace(export, "{key}", kv.Key, 1)
		export = strings.Replace(export, "{value}", kv.Value, 1)

		kv = kv.Next
	}

	export += `}
`
	return export, err
}
