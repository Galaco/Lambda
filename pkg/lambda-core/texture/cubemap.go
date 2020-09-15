package texture

// Cubemap is a 6-sided edgeless texture that can be mapped to a cube,
// Used mainly for pre-computed reflections
type Cubemap struct {
	Texture2D
	Faces []ITexture
}

// Width Get material width.
// Must have exactly 6 faces, and all faces are assumed the same size
func (material *Cubemap) Width() int {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Width()
}

// Height Get material height.
// Must have exactly 6 faces, and all faces are assumed the same size
func (material *Cubemap) Height() int {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Height()
}

// Format get material format
// Same format for all faces assumed
func (material *Cubemap) Format() uint32 {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Format()
}

// NewCubemap returns a new cubemap material
func NewCubemap(materials []ITexture) *Cubemap {
	return &Cubemap{
		Faces: materials,
	}
}
