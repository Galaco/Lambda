package thumbcache

import "testing"

func TestHashDirectory(t *testing.T) {
	dir := "C:/Project/ika/mount/materials/ika"

	expected := "1cfb0f84c90906439a2641a814cc6075"
	result := hashDirectory(dir)
	if expected != result {
		t.Error(result)
	}
}
