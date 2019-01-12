package world

import "github.com/galaco/source-tools-common/entity"

type World struct {
	keyvalues *entity.Entity
	solids             []Solid
}

func (world *World) AddSolid(solid *Solid) error {
	// @TODO
	// Assign an id to a solid that is unique, and check
	world.solids = append(world.solids, *solid)

	return nil
}

func NewWorld(entityKvs *entity.Entity, solids []Solid) *World {
	return &World{
		keyvalues: entityKvs,
		solids: solids,
	}
}
