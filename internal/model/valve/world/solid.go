package world

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
)

type Solid struct {
	Id     int
	Sides  []Side
	Editor *Editor
}

type Side struct {
	Id              int
	Plane           Plane
	Material        string
	UAxis           UVTransform
	VAxis           UVTransform
	Rotation        float64
	LightmapScale   float64
	SmoothingGroups bool
}

type UVTransform struct {
	Transform mgl64.Vec4
	Scale     float64
}

type Editor struct {
	color             mgl64.Vec3
	visgroupShown     bool
	visGroupAutoShown bool

	logicalPos mgl64.Vec2 // only exists on brush entities?
}

type Plane [3]mgl64.Vec3

func NewSolid(id int, sides []Side, editor *Editor) *Solid {
	return &Solid{
		Id:     id,
		Sides:  sides,
		Editor: editor,
	}
}

func NewSide(id int, plane Plane, material string, uAxis UVTransform, vAxis UVTransform, rotation float64, lightmapScale float64, smoothingGroups bool) *Side {
	return &Side{
		Id:              id,
		Plane:           plane,
		Material:        material,
		UAxis:           uAxis,
		VAxis:           vAxis,
		Rotation:        rotation,
		LightmapScale:   lightmapScale,
		SmoothingGroups: smoothingGroups,
	}
}

// NewEditor
func NewEditor(color mgl64.Vec3, visgroupShown bool, visgroupAutoShown bool) *Editor {
	return &Editor{
		color:             color,
		visgroupShown:     visgroupShown,
		visGroupAutoShown: visgroupAutoShown,
		logicalPos: mgl64.Vec2{},
	}
}

// NewPlane
func NewPlane(a mgl64.Vec3, b mgl64.Vec3, c mgl64.Vec3) *Plane {
	p := Plane([3]mgl64.Vec3{a, b, c})
	return &p
}

// NewPlaneFromString
func NewPlaneFromString(marshalled string) *Plane {
	var v1, v2, v3 = float64(0), float64(0), float64(0)
	var v4, v5, v6 = float64(0), float64(0), float64(0)
	var v7, v8, v9 = float64(0), float64(0), float64(0)
	fmt.Sscanf(marshalled, "(%f %f %f) (%f %f %f) (%f %f %f)", &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9)

	return NewPlane(
		mgl64.Vec3{v1, v2, v3},
		mgl64.Vec3{v4, v5, v6},
		mgl64.Vec3{v7, v8, v9})
}

// NewUVTransform
func NewUVTransform(transform mgl64.Vec4, scale float64) *UVTransform {
	return &UVTransform{
		Transform: transform,
		Scale:     scale,
	}
}

// NewUVTransformFromString
func NewUVTransformFromString(marshalled string) *UVTransform {
	var v1, v2, v3, v4 = float64(0), float64(0), float64(0), float64(0)
	var scale = float64(0)
	fmt.Sscanf(marshalled, "[%f %f %f %f] %f", &v1, &v2, &v3, &v4, &scale)
	return NewUVTransform(mgl64.Vec4{v1, v2, v3, v4}, scale)
}
