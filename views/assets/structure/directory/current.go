package directory

import (
	"github.com/galaco/Lambda/lib/imgui-layouts/directory"
	"log"
)

type Directory struct {
	viewDir *directory.Directory
}

func (dir *Directory) Render() {
	dir.viewDir.Render()
}

func NewDirectory(fileList []string) *Directory {
	thumbs := make([]directory.Thumbnail, len(fileList))
	for idx, f := range fileList {
		thumbs[idx] = *directory.NewThumbnail(f, false, func() {
			log.Println(f)
		})
	}

	return &Directory{
		viewDir: directory.NewDirectory(thumbs),
	}
}