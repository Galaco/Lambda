package events

import (
	"reflect"
	"testing"
)

func TestNewNewScene(t *testing.T) {
	sut := NewNewScene(nil)
	if reflect.TypeOf(sut) == reflect.TypeOf(NewScene{}) {
		t.Error("unexpected type returned")
	}
}

func TestOpenScene_Path(t *testing.T) {
	sut := NewOpenScene("foo/bar")
	if sut.Path() != "foo/bar" {
		t.Error("unexpected filepath")
	}
}

func TestNewOpenScene(t *testing.T) {
	sut := NewOpenScene("")
	if reflect.TypeOf(sut) == reflect.TypeOf(OpenScene{}) {
		t.Error("unexpected type returned")
	}
}

func TestNewOpenSceneFailed(t *testing.T) {
	sut := NewOpenSceneFailed()
	if reflect.TypeOf(sut) == reflect.TypeOf(OpenSceneFailed{}) {
		t.Error("unexpected type returned")
	}
}

func TestNewScene_Model(t *testing.T) {

}

func TestNewSceneNodeSelected(t *testing.T) {
	sut := NewEntitySelected(nil)
	if reflect.TypeOf(sut) == reflect.TypeOf(EntitySelected{}) {
		t.Error("unexpected type returned")
	}
}

func TestNewScene_Type(t *testing.T) {
	sut := NewNewScene(nil)
	if sut.Type() != TypeNewScene {
		t.Error("unexpected event type for event")
	}
}

func TestOpenScene_Type(t *testing.T) {
	sut := NewOpenScene("")
	if sut.Type() != TypeOpenScene {
		t.Error("unexpected event type for event")
	}
}

func TestOpenSceneFailed_Type(t *testing.T) {
	sut := NewOpenSceneFailed()
	if sut.Type() != TypeOpenSceneFailed {
		t.Error("unexpected event type for event")
	}
}

func TestSceneNodeSelected_Type(t *testing.T) {
	sut := NewSceneNodeSelected(0)
	if sut.Type() != TypeSceneNodeSelected {
		t.Error("unexpected event type for event")
	}
}
