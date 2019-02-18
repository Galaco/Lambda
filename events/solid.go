package events

import (
	"github.com/galaco/Lambda/valve/world"
)

const TypeNewSolidCreated = "NewSolidCreated"

type NewSolidCreated struct {
	solid *world.Solid
}

func (act *NewSolidCreated) Type() string {
	return TypeNewSolidCreated
}

func (act *NewSolidCreated) Target() *world.Solid {
	return act.solid
}

func NewNewSolidCreated(selected *world.Solid) *NewSolidCreated {
	return &NewSolidCreated{
		solid: selected,
	}
}
