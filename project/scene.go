package project

import (
	"github.com/galaco/Lambda/valve/world"
	"github.com/galaco/source-tools-common/entity"
)

type Scene struct {
	world    *world.World
	entities *entity.List
}

func (s *Scene) World() *world.World {
	return s.world
}

func (s *Scene) SetWorld(objects *world.World) {
	s.world = objects
}

func (s *Scene) Entities() *entity.List {
	return s.entities
}

func (s *Scene) SetEntities(entities *entity.List) {
	s.entities = entities
}

func NewScene() *Scene {
	return &Scene{}
}
