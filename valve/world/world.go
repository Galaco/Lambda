package world

import "github.com/galaco/source-tools-common/entity"

type World struct {
	Keyvalues *entity.Entity
	Solids    []Solid
}

func (world *World) AddSolid(solid *Solid) error {
	// @TODO
	// Assign an id to a solid that is unique, and check
	world.Solids = append(world.Solids, *solid)

	return nil
}

func NewWorld(entityKvs *entity.Entity, solids []Solid) *World {
	return &World{
		Keyvalues: entityKvs,
		Solids:    solids,
	}
}
