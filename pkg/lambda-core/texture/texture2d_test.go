package texture

import "testing"

func TestTexture2D_FilePath(t *testing.T) {
	tex := NewError("error.vtf")
	if tex.FilePath() != "error.vtf" {
		t.Error("incorrect filepath for texture")
	}
}

func TestTexture2D_Height(t *testing.T) {
	tex := NewError("error.vtf")

	if tex.Height() != 8 {
		t.Error("unexpected height")
	}
}

func TestTexture2D_Width(t *testing.T) {
	tex := NewError("error.vtf")

	if tex.Width() != 8 {
		t.Error("unexpected width")
	}
}
