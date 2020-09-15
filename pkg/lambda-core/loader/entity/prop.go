package entity

import (
	"github.com/galaco/Lambda/pkg/lambda-core/entity"
	entity2 "github.com/galaco/Lambda/pkg/lambda-core/game/entity"
	"github.com/galaco/Lambda/pkg/lambda-core/loader/prop"
	"github.com/galaco/Lambda/pkg/lambda-core/resource"
	"github.com/golang-source-engine/filesystem"
	"strings"
)

// DoesEntityReferenceStudioModel tests if an entity is
// tied to a model (normally prop_* classnames, but not exclusively)
func DoesEntityReferenceStudioModel(ent entity.IEntity) bool {
	return strings.HasSuffix(ent.KeyValues().ValueForKey("model"), ".mdl")
}

// AssignStudioModelToEntity sets a renderable entity's model
func AssignStudioModelToEntity(entity entity.IEntity, fs *filesystem.FileSystem) {
	modelName := entity.KeyValues().ValueForKey("model")
	if !resource.Manager().HasModel(modelName) {
		m, _ := prop.LoadProp(modelName, fs)
		entity.(entity2.IProp).SetModel(m)
	} else {
		entity.(entity2.IProp).SetModel(resource.Manager().Model(modelName))
	}
}
