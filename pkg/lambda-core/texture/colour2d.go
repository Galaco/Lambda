package texture

import (
	"github.com/galaco/vtf/format"
)

// Colour2D is a material defined by raw/computed colour data,
// rather than loaded vtf data
type Colour2D struct {
	Texture2D
	rawColourData []uint8
}

// Format returns colour format
func (error *Colour2D) Format() uint32 {
	return uint32(format.RGB888)
}

// PixelDataForFrame returns raw colour data for specific animation
// frame
func (error *Colour2D) PixelDataForFrame(frame int) []byte {
	return error.rawColourData
}

// Thumbnail return a low resolution version of the image
func (error *Colour2D) Thumbnail() []byte {
	return append(error.rawColourData, append(error.rawColourData, append(error.rawColourData, error.rawColourData...)...)...)
}

// NewError returns new Error material
func NewError(name string) *Colour2D {
	mat := Colour2D{}

	mat.width = 8
	mat.height = 8
	mat.filePath = name

	// This generates purple & black chequers.
	mat.rawColourData = []uint8{
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
	}

	return &mat
}
