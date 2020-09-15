package util

import "testing"

func TestRemoveDuplicatesFromList(t *testing.T) {
	samples := []string{
		"foo",
		"bar",
		"foo",
		"bar",
		"baz",
		"foo",
		"bar",
		"bat",
		"bat",
	}
	expected := []string{
		"foo",
		"bar",
		"baz",
		"bat",
	}
	actual := RemoveDuplicatesFromList(samples)
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error("unexpected entry in list after duplicate removal")
		}
	}
}
