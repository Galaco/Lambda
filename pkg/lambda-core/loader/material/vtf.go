package material

import (
	filesystem2 "github.com/galaco/Lambda/pkg/lambda-core/filesystem"
	"github.com/galaco/Lambda/pkg/lambda-core/resource"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"github.com/galaco/vtf"
	"github.com/golang-source-engine/filesystem"
	"strings"
)

// LoadVtfFromFilesystem
func LoadVtfFromFilesystem(fs *filesystem.FileSystem, filePath string) (texture.ITexture, error) {
	if filePath == "" {
		return resource.Manager().Texture(resource.Manager().ErrorTextureName()).(texture.ITexture), nil
	}
	filePath = filesystem2.BasePathMaterial + filesystem.NormalisePath(filePath)
	if !strings.HasSuffix(filePath, filesystem2.ExtensionVtf) {
		filePath = filePath + filesystem2.ExtensionVtf
	}
	if resource.Manager().HasTexture(filePath) {
		return resource.Manager().Texture(filePath).(texture.ITexture), nil
	}
	mat, err := readVtf(filePath, fs)
	if err != nil {
		return resource.Manager().Texture(resource.Manager().ErrorTextureName()).(texture.ITexture), err
	}
	return mat, nil
}

// readVtf
func readVtf(path string, fs *filesystem.FileSystem) (texture.ITexture, error) {
	ResourceManager := resource.Manager()
	stream, err := fs.GetFile(path)
	if err != nil {
		return nil, err
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	read, err := vtf.ReadFromStream(stream)
	if err != nil {
		return nil, err
	}
	// Store filesystem containing raw data in memory
	ResourceManager.AddTexture(
		texture.NewTexture2D(
			path,
			read,
			int(read.Header().Width),
			int(read.Header().Height)))

	// Finally generate the gpu buffer for the material
	return ResourceManager.Texture(path).(texture.ITexture), nil
}
