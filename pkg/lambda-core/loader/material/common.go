package material

import "io"

type VirtualFilesystem interface {
	GetFile(string) (io.Reader, error)
}
