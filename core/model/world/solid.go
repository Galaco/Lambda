package world

import "github.com/go-gl/mathgl/mgl64"

type Solid struct {
	id     int
	sides  []Side
	editor Editor
}

type Side struct {
	id              int
	plane           Plane
	material        string
	uAxis           UVTransform
	vAxis           UVTransform
	rotation        float64
	lightmapScale   float64
	smoothingGroups bool
}

type UVTransform struct {
	transform mgl64.Vec4
	scale     float64
}

type Editor struct {
	color             mgl64.Vec3
	visgroupShown     bool
	visGroupAutoShown bool

	logicalPos mgl64.Vec2 // only exists on brush entities?
}

type Plane [3]mgl64.Vec3
