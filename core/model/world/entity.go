package world

type Entity struct {
	id int

	// Entity keyvalues

	// only for brush entities
	solids []Solid
}
