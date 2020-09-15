package entity

import (
	entity3 "github.com/galaco/Lambda/pkg/lambda-core/entity"
	"github.com/galaco/Lambda/pkg/lambda-core/loader/entity/classmap"
	"github.com/galaco/source-tools-common/entity"
	"github.com/galaco/vmf"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang-source-engine/filesystem"
	"strings"
)

// ParseEntities Parse Base block.
// Vmf lib is actually capable of doing this;
// contents are loaded into Vmf.Unclassified
func ParseEntities(data string) (vmf.Vmf, error) {
	stringReader := strings.NewReader(data)
	reader := vmf.NewReader(stringReader)

	return reader.Read()
}

// CreateEntity creates a new entity with common properties
// e.g. origin and angles
func CreateEntity(ent *entity.Entity, fs *filesystem.FileSystem) entity3.IEntity {
	localEdict := loader.New(ent.ValueForKey("classname"))
	if localEdict == nil {
		localEdict = entity3.NewGenericEntity(ent)
	} else {
		localEdict.SetKeyValues(ent)
	}

	origin := ent.VectorForKey("origin")
	localEdict.Transform().Position = mgl32.Vec3{origin.X(), origin.Y(), origin.Z()}
	angles := ent.VectorForKey("angles")
	localEdict.Transform().Rotation = mgl32.Vec3{angles.X(), angles.Y(), angles.Z()}

	AssignProperties(localEdict, fs)

	return localEdict
}

// AssignProperties assigns type specific properties.
// @TODO This is probably going to grow massively as more common types get implemented.
// It should probably be refactored.
func AssignProperties(ent entity3.IEntity, fs *filesystem.FileSystem) {
	if DoesEntityReferenceStudioModel(ent) {
		AssignStudioModelToEntity(ent, fs)
	}
}
