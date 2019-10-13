package list

import (
	"testing"
)

func TestList_AddRow(t *testing.T) {
	list := List{}

	id := 0
	className := "foo"
	list.AddRow(id, className, nil)

	if list.nodes[0].label != className+"##0" {
		t.Error("returned row does not matched added row")
	}

	id = 1
	className = "foo2"
	list.AddRow(id, className, nil)

	if list.nodes[1].label != className+"##1" {
		t.Error("returned row does not matched added row")
	}
}

func TestList_Filter(t *testing.T) {
	list := List{}

	ids := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	classNames := [8]string{
		"zero",
		"one",
		"prop_two",
		"three",
		"four",
		"prop_five",
		"six",
		"seven",
	}

	for i := 0; i < 8; i++ {
		list.AddRow(ids[i], classNames[i], nil)
	}

	filtered := list.Filter(EntityFilterPropOnly)

	if filtered.rows[0].label != classNames[2]+"##2" ||
		filtered.rows[1].label != classNames[5]+"##5" {
		t.Error("expected row did not pass filter")
	}
}
