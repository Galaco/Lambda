package events

import "github.com/galaco/source-tools-common/entity"

const TypeEntityCreated = "NewEntityCreated"

type EntityCreated struct {
	entity *entity.Entity
}

func (act *EntityCreated) Type() string {
	return TypeEntityCreated
}

func (act *EntityCreated) Target() *entity.Entity {
	return act.entity
}

func NewEntityCreated(selected *entity.Entity) *EntityCreated {
	return &EntityCreated{
		entity: selected,
	}
}
