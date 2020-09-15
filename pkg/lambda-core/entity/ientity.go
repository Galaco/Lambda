package entity

import (
	"github.com/galaco/source-tools-common/entity"
)

// IEntity Base interface
// All game entities need to implement this
type IEntity interface {
	// KeyValues
	KeyValues() *entity.Entity
	// SetKeyValues
	SetKeyValues(*entity.Entity)
	// Classname
	Classname() string
	// Transform
	Transform() *Transform
	// New
	New() IEntity
}

// IClassname all valid game entities should have a classname,
// but there may be temporary non-game entities that have classnames
type IClassname interface {
	// Classname
	Classname() string
}
