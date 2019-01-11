package world

type World struct {
	id                 int
	mapVersion         int
	classname          string
	detailMaterial     string
	detailVbsp         string
	maxPropScreenWidth int
	skyName            string
	targetname         string
	solids             []Solid
}

func (world *World) AddSolid(solid *Solid) error {
	// @TODO
	// Assign an id to a solid that is unique, and check
	world.solids = append(world.solids, *solid)

	return nil
}
