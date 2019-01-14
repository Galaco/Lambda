package models

import (
	"github.com/galaco/Lambda/models/world"
	"github.com/galaco/source-tools-common/entity"
	"github.com/go-gl/mathgl/mgl64"
)

type Vmf struct {
	versionInfo  VersionInfo
	visGroups    VisGroups
	viewSettings ViewSettings
	world        world.World
	entities     entity.List
	cameras      Cameras
	cordon       Cordon
}

func (vmf *Vmf) VersionInfo() *VersionInfo {
	return &vmf.versionInfo
}

func (vmf *Vmf) Visgroups() *VisGroups {
	return &vmf.visGroups
}

func (vmf *Vmf) ViewSettings() *ViewSettings {
	return &vmf.viewSettings
}

func (vmf *Vmf) Worldspawn() *world.World {
	return &vmf.world
}

func (vmf *Vmf) Entities() *entity.List {
	return &vmf.entities
}

func (vmf *Vmf) Cameras() *Cameras {
	return &vmf.cameras
}

func (vmf *Vmf) Cordons() *Cordon {
	return &vmf.cordon
}

type VersionInfo struct {
	editorVersion int
	editorBuild   int
	mapVersion    int
	formatVersion int
	prefab        bool
}

func NewVersionInfo(version int, build int, mapRevision int, format int, isPrefab bool) *VersionInfo {
	return &VersionInfo{
		editorVersion: version,
		editorBuild:   build,
		mapVersion:    mapRevision,
		formatVersion: format,
		prefab:        isPrefab,
	}
}

type VisGroups struct {
}

type ViewSettings struct {
	snapToGrid      bool
	showGrid        bool
	showLogicalGrid bool
	gridSpacing     int
	show3DGrid      bool
}

type Cameras struct {
	activeCamera int
	cameraList   []Camera
}

type Camera struct {
	position mgl64.Vec3
	look     mgl64.Vec3
}

type Cordon struct {
	mins   mgl64.Vec3
	maxs   mgl64.Vec3
	active bool
}

func NewVmf(version *VersionInfo, visgroups *VisGroups, worldSpawn *world.World, entities *entity.List) *Vmf {
	return &Vmf{
		versionInfo: *version,
		visGroups:   *visgroups,
		world:       *worldSpawn,
		entities:    *entities,
	}
}
