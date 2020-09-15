package entity

import (
	"reflect"
	"testing"
)

func TestNewGenericEntity(t *testing.T) {
	sut := NewGenericEntity(nil)

	if reflect.TypeOf(sut) != reflect.TypeOf(&GenericEntity{}) {
		t.Errorf("Expected: %s, but received: %s", reflect.TypeOf(&GenericEntity{}), reflect.TypeOf(sut))
	}
}
