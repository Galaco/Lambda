package list

import "testing"

func TestFilteredList_Render(t *testing.T) {
	t.Skip("cannot reliably test rendering with imgui bindings")
}

func TestNewFilteredList(t *testing.T) {
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

	filtered = list.Filter(EntityFilterNone)

	if len(filtered.rows) != len(ids) {
		t.Error("unexpect row count returned from filter")
	}
}
