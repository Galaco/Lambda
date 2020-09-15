package common

import (
	entity2 "github.com/galaco/Lambda/pkg/lambda-core/entity"
	"github.com/galaco/Lambda/pkg/lambda-core/game/entity"
)

type PropDoorRotating struct {
	entity2.Base
	entity.PropBase
}

func (entity *PropDoorRotating) New() entity2.IEntity {
	return &PropDoorRotating{}
}

func (entity PropDoorRotating) Classname() string {
	return "prop_door_rotating"
}
