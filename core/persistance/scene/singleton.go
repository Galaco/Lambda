package scene

import (
	"github.com/galaco/Lambda/core/models/world"
	"github.com/galaco/source-tools-common/entity"
)

var singleton scene

type scene struct {
	world *world.World
	entities *entity.List
}

func (s *scene) World() *world.World {
	return s.world
}

func (s *scene) SetWorld(objects *world.World) {
	s.world = objects
}

func (s *scene) Entities() *entity.List {
	return s.entities
}

func (s *scene) SetEntities(entities *entity.List) {
	s.entities = entities
}

func Singleton() *scene {
	return &singleton
}