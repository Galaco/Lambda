package events

import (
	"reflect"
	"testing"
)

func TestNewWindowClosed(t *testing.T) {
	sut := NewWindowClosed()
	if reflect.TypeOf(sut) == reflect.TypeOf(windowClosed{}) {
		t.Error("unexpected type returned")
	}
}

func TestWindowClosed_Type(t *testing.T) {
	sut := NewWindowClosed()
	if sut.Type() != TypeWindowClosed {
		t.Error("unexpected event type for event")
	}
}
