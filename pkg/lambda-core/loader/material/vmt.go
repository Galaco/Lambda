package material

import (
	"github.com/galaco/vmt"
)

func LoadVmtFromFilesystem(fs VirtualFilesystem, filePath string) (*vmt.Properties, error) {
	mat,err := vmt.FromFilesystem(filePath, fs, vmt.NewProperties())
	if err != nil {
		return nil, err
	}

	return mat.(*vmt.Properties), nil
}