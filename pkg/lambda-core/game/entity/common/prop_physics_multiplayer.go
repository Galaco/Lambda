package common

import (
	"github.com/galaco/Lambda/pkg/lambda-core/entity"
	entity2 "github.com/galaco/Lambda/pkg/lambda-core/game/entity"
)

// PropPhysicsMultiplayer
type PropPhysicsMultiplayer struct {
	entity.Base
	entity2.PropBase
}

// New
func (entity *PropPhysicsMultiplayer) New() entity.IEntity {
	return &PropPhysicsMultiplayer{}
}

// Classname
func (entity PropPhysicsMultiplayer) Classname() string {
	return "prop_physics_multiplayer"
}
