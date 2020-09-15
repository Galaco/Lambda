package event

import (
	"reflect"
	"testing"
)

func TestGetEventManager(t *testing.T) {
	if reflect.TypeOf(Manager()) != reflect.TypeOf(&manager{}) || Manager() == nil {
		t.Error("Unexpected value for event manager")
	}
}

func TestManager_Dispatch(t *testing.T) {

}

func TestManager_Listen(t *testing.T) {

}

func TestManager_Unlisten(t *testing.T) {

}
