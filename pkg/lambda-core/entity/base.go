package entity

import (
	entity2 "github.com/galaco/source-tools-common/entity"
)

// Base is an object in the game world.
// By itself entity is nothing more than an identifiable object located at a point in space
type Base struct {
	keyValues *entity2.Entity
	transform Transform

	handle EntityId
}

// SetKeyValues set this entity's keyvalues
func (entity *Base) SetKeyValues(keyValues *entity2.Entity) {
	entity.keyValues = keyValues
}

// KeyValues returns this entitys keyvalues
func (entity *Base) KeyValues() *entity2.Entity {
	return entity.keyValues
}

// Classname gets this entitiy's classname
func (entity *Base) Classname() string {
	if entity.keyValues == nil {
		return "generic"
	}
	return entity.keyValues.ValueForKey("classname")
}

// Transform Returns this entity's transform component
func (entity *Base) Transform() *Transform {
	return &entity.transform
}

// New entity
func (entity *Base) New() IEntity {
	return &Base{}
}
