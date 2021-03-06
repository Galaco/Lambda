package valve

import (
	"github.com/galaco/Lambda/internal/model/valve/world"
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
	EditorVersion int
	EditorBuild   int
	MapVersion    int
	FormatVersion int
	Prefab        bool
}

func NewVersionInfo(version int, build int, mapRevision int, format int, isPrefab bool) *VersionInfo {
	return &VersionInfo{
		EditorVersion: version,
		EditorBuild:   build,
		MapVersion:    mapRevision,
		FormatVersion: format,
		Prefab:        isPrefab,
	}
}

type VisGroups struct {
}

type ViewSettings struct {
	SnapToGrid      bool
	ShowGrid        bool
	ShowLogicalGrid bool
	GridSpacing     int
	Show3DGrid      bool
}

type Cameras struct {
	ActiveCamera int
	CameraList   []Camera
}

func NewCameras(activeCameraIndex int, cameras []Camera) *Cameras {
	return &Cameras{
		ActiveCamera: activeCameraIndex,
		CameraList:   cameras,
	}
}

type Camera struct {
	Position mgl64.Vec3
	Look     mgl64.Vec3
}

func NewCamera(position mgl64.Vec3, look mgl64.Vec3) *Camera {
	return &Camera{
		Position: position,
		Look:     look,
	}
}

type Cordon struct {
	mins   mgl64.Vec3
	maxs   mgl64.Vec3
	active bool
}

func NewVmf(version *VersionInfo, visgroups *VisGroups, worldSpawn *world.World, entities *entity.List, cameras *Cameras) *Vmf {
	return &Vmf{
		versionInfo: *version,
		visGroups:   *visgroups,
		world:       *worldSpawn,
		entities:    *entities,
		cameras:     *cameras,
	}
}
