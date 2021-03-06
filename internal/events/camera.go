package events

import (
	"github.com/galaco/Lambda/internal/model/valve"
)

const (
	// TypeNewCameraCreated event type
	TypeNewCameraCreated = "NewCameraCreated"
	// TypeCameraChanged event type
	TypeCameraChanged = "CameraChanged"
)

type NewCameraCreated struct {
	camera *valve.Camera
}

func (act *NewCameraCreated) Type() string {
	return TypeNewCameraCreated
}

func (act *NewCameraCreated) Target() *valve.Camera {
	return act.camera
}

func NewNewCameraCreated(selected *valve.Camera) *NewCameraCreated {
	return &NewCameraCreated{
		camera: selected,
	}
}

type CameraChanged struct {
	camera *valve.Camera
}

func (act *CameraChanged) Type() string {
	return TypeCameraChanged
}

func (act *CameraChanged) Target() *valve.Camera {
	return act.camera
}

func NewCameraChanged(selected *valve.Camera) *CameraChanged {
	return &CameraChanged{
		camera: selected,
	}
}
