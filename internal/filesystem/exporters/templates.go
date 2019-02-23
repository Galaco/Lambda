package exporters

const templateVmfVersionInfo = `versioninfo
{
	"editorversion" "{editorversion}"
	"editorbuild" "{editorbuild}"
	"mapversion" "{mapversion}"
	"formatversion" "{formatversion}"
	"prefab" "{prefab}"
}
`

const templateVmfVisgroups = `visgroups
{
}
`

const templateVmfViewSettings = `viewsettings
{
	"bSnapToGrid" "{bSnapToGrid}"
	"bShowGrid" "{bShowGrid}"
	"bShowLogicalGrid" "{bShowLogicalGrid}"
	"nGridSpacing" "{nGridSpacing}"
	"bShow3DGrid" "{bShow3DGrid}"
}
`

const templateVmfWorld = `world
{
`

const templateVmfKeyValue = `	"{key}" "{value}"
`
