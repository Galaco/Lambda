package keyvalues

import (
	"github.com/galaco/KeyValues"
	"github.com/galaco/filesystem"
)

// ReadKeyValues loads a keyvalues file.
// Its just a simple wrapper that combines the KeyValues library and
// the filesystem module.
func ReadKeyValues(filePath string, fs *filesystem.FileSystem) (*keyvalues.KeyValue, error) {
	stream, err := fs.GetFile(filePath)
	if err != nil {
		return nil, err
	}

	reader := keyvalues.NewReader(stream)
	kvs, err := reader.Read()

	return &kvs, err
}
