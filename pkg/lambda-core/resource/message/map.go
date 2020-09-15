package message

import (
	"github.com/galaco/Lambda/pkg/lambda-core/event"
	"github.com/galaco/Lambda/pkg/lambda-core/model"
)

const (
	// TypeMapLoaded
	TypeMapLoaded = event.MessageType("MapLoaded")
	// TypeMapUnloaded
	TypeMapUnloaded = event.MessageType("MapUnloaded")
)

// MapLoaded
type MapLoaded struct {
	event.Message
	// Resource
	Resource *model.Bsp
}

// Type
func (message *MapLoaded) Type() event.MessageType {
	return TypeMapLoaded
}

// MapUnloaded
type MapUnloaded struct {
	event.Message
	// Resource
	Resource *model.Bsp
}

// Type
func (message *MapUnloaded) Type() event.MessageType {
	return TypeMapUnloaded
}

// LoadedMap
func LoadedMap(world *model.Bsp) event.IMessage {
	return &MapLoaded{
		Resource: world,
	}
}

// UnloadedMap
func UnloadedMap(world *model.Bsp) event.IMessage {
	return &MapUnloaded{
		Resource: world,
	}
}
