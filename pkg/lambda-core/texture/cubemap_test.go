package texture

import (
	"testing"
)

func TestNewCubemap(t *testing.T) {
	tex := NewError("error.vtf")

	cb := NewCubemap([]ITexture{tex, tex, tex, tex, tex, tex})
	if cb == nil {
		t.Error("failed to create cubemap from textures")
	}
}

func TestCubemap_Format(t *testing.T) {
	tex := NewError("error.vtf")

	cb := NewCubemap([]ITexture{tex, tex, tex, tex, tex, tex})
	if cb.Width() != tex.Width() {
		t.Error("unexpected cubemap width")
	}
	if tex.Format() != cb.Format() {
		t.Error("unexpected error colour data format")
	}

	cb2 := NewCubemap([]ITexture{})
	if cb2.Format() != 0 {
		t.Error("unexpected cubemap format")
	}
}

func TestCubemap_Height(t *testing.T) {
	tex := NewError("error.vtf")

	cb := NewCubemap([]ITexture{tex, tex, tex, tex, tex, tex})
	if cb.Height() != tex.Height() {
		t.Error("unexpected cubemap height")
	}

	cb2 := NewCubemap([]ITexture{})
	if cb2.Height() != 0 {
		t.Error("unexpected cubemap height")
	}
}

func TestCubemap_Width(t *testing.T) {
	tex := NewError("error.vtf")

	cb := NewCubemap([]ITexture{tex, tex, tex, tex, tex, tex})
	if cb.Width() != tex.Width() {
		t.Error("unexpected cubemap width")
	}

	cb2 := NewCubemap([]ITexture{})
	if cb2.Width() != 0 {
		t.Error("unexpected cubemap width")
	}
}
