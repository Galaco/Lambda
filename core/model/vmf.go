package model

import (
	"github.com/galaco/Lambda/core/model/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Vmf struct {
	versionInfo  VersionInfo
	visGroups    VisGroups
	viewSettings ViewSettings
	world        world.World
	cameras      Cameras
	cordon       Cordon
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
