package entity

import (
	"reflect"
	"testing"
)

func TestBase_Classname(t *testing.T) {
	sut := &Base{}
	if sut.Classname() != "generic" {
		t.Errorf("unexpected classname. Received %s, but expected %s", sut.Classname(), "generic")
	}
}

func TestBase_New(t *testing.T) {
	sut := &Base{}

	actual := sut.New().(*Base)

	if reflect.TypeOf(actual) != reflect.TypeOf(sut) {
		t.Errorf("Expected: %d, but got: %s", reflect.TypeOf(sut), reflect.TypeOf(actual))
	}
}

func TestBase_KeyValues(t *testing.T) {
	t.Skip()
}

func TestBase_SetKeyValues(t *testing.T) {
	t.Skip()
}

func TestBase_Transform(t *testing.T) {
	t.Skip()
}
