package game

import (
	"github.com/galaco/Lambda/pkg/lambda-core/game/entity/common"
)

// CounterstrikeSource
type CounterstrikeSource struct{}

// RegisterEntityClasses loads all Game entity classes into the engine.
func (target *CounterstrikeSource) RegisterEntityClasses() {
	loader.RegisterClass(&common.PropDoorRotating{})
	loader.RegisterClass(&common.PropDynamic{})
	loader.RegisterClass(&common.PropDynamicOrnament{})
	loader.RegisterClass(&common.PropDynamicOverride{})
	loader.RegisterClass(&common.PropPhysics{})
	loader.RegisterClass(&common.PropPhysicsMultiplayer{})
	loader.RegisterClass(&common.PropPhysicsOverride{})
	loader.RegisterClass(&common.PropRagdoll{})
}
