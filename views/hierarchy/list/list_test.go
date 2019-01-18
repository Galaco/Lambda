package list

import "testing"

func TestList_AddRow(t *testing.T) {
	list := List{}

	id := 0
	className := "foo"
	targetName := "bar"
	list.AddRow(id, className, targetName)

	if list.nodes[0].TargetName != targetName {
		t.Error("returned row does not matched added row")
	}

	id = 1
	className = "foo2"
	targetName = ""
	list.AddRow(id, className, targetName)

	if list.nodes[1].Classname != className {
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
	targetNames := [8]string{
		"Tzero",
		"Tone",
		"Ttwo",
		"Tthree",
		"Tfour",
		"Tfive",
		"Tsix",
		"Tseven",
	}

	for i := 0; i < 8; i++ {
		list.AddRow(ids[i], classNames[i], targetNames[i])
	}

	filtered := list.Filter(EntityFilterPropOnly)

	if filtered.rows[0].TargetName != targetNames[2] ||
		filtered.rows[1].TargetName != targetNames[5] {
		t.Error("expected row did not pass filter")
	}
}
