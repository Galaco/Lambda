package world

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
)

func NewVec3FromString(marshalled string) mgl64.Vec3 {
	var x, y, z float64
	fmt.Sscanf(marshalled, "[%f %f %f]", &x, &y, &z)

	return mgl64.Vec3{x, y, z}
}
