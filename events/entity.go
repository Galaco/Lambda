package events

import "github.com/galaco/source-tools-common/entity"

const TypeEntitySelected = "EntityPropertiesList"

type EntitySelected struct {
	entity *entity.Entity
}

func (act *EntitySelected) Type() string {
	return TypeOpenScene
}

func (act *EntitySelected) Target() *entity.Entity {
	return act.entity
}

func NewEntitySelected(selected *entity.Entity) *EntitySelected {
	return &EntitySelected{
		entity: selected,
	}
}