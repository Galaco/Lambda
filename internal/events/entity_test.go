package events

import (
	"github.com/galaco/source-tools-common/entity"
	"reflect"
	"testing"
)

func TestEntityCreated_Target(t *testing.T) {
	e := &entity.Entity{}
	sut := NewEntityCreated(e)

	if sut.Target() != e {
		t.Error("event payload mismatch")
	}
}

func TestEntityCreated_Type(t *testing.T) {
	sut := NewEntityCreated(nil)
	if sut.Type() != TypeEntityCreated {
		t.Error("unexpected event type for event")
	}
}

func TestEntitySelected_Target(t *testing.T) {
	e := &entity.Entity{}
	sut := NewEntitySelected(e)

	if sut.Target() != e {
		t.Error("event payload mismatch")
	}
}

func TestEntitySelected_Type(t *testing.T) {
	sut := NewEntitySelected(nil)
	if sut.Type() != TypeEntitySelected {
		t.Error("unexpected event type for event")
	}
}

func TestNewEntityCreated(t *testing.T) {
	sut := NewEntityCreated(nil)
	if reflect.TypeOf(sut) == reflect.TypeOf(EntityCreated{}) {
		t.Error("unexpected type returned")
	}
}

func TestNewEntitySelected(t *testing.T) {
	sut := NewEntitySelected(nil)
	if reflect.TypeOf(sut) == reflect.TypeOf(EntitySelected{}) {
		t.Error("unexpected type returned")
	}
}
