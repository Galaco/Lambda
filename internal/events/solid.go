package events

import (
	"github.com/galaco/Lambda/internal/model/valve/world"
)

const (
	// TypeNewSolidCreated event type
	TypeNewSolidCreated = "NewSolidCreated"
)

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
