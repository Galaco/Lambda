package material

import "github.com/galaco/Lambda/pkg/lambda-core/texture"

// Material
type Material struct {
	filePath string
	// ShaderName
	ShaderName string
	// Textures
	Textures struct {
		// Albedo
		Albedo texture.ITexture
		// Normal
		Normal texture.ITexture
	}
	// BaseTextureName
	BaseTextureName string
	// BumpMapName
	BumpMapName string
	// Properties
	Properties struct {
	}
}

// Width returns this materials width. Albedo is used to
// determine material width where possible
func (mat *Material) Width() int {
	return mat.Textures.Albedo.Width()
}

// Height returns this materials height. Albedo is used to
// determine material height where possible
func (mat *Material) Height() int {
	return mat.Textures.Albedo.Height()
}

// FilePath returns this materials location in whatever
// filesystem it was found
func (mat *Material) FilePath() string {
	return mat.filePath
}

func NewMaterial(filePath string) *Material {
	return &Material{
		filePath: filePath,
	}
}
