package events

import "github.com/galaco/source-tools-common/entity"

const TypeEntitySelected = "EntitySelected"
const TypeEntityCreated = "NewEntityCreated"

type EntitySelected struct {
	entity *entity.Entity
}

func (act *EntitySelected) Type() string {
	return TypeEntitySelected
}

func (act *EntitySelected) Target() *entity.Entity {
	return act.entity
}

func NewEntitySelected(selected *entity.Entity) *EntitySelected {
	return &EntitySelected{
		entity: selected,
	}
}

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
