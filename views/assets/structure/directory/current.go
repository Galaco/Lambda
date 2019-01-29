package directory

import (
	"github.com/galaco/Lambda/lib/imgui-layouts/thumbnail"
	"log"
)

type Directory struct {
	viewDir *thumbnail.Directory
}

func (dir *Directory) Render() {
	dir.viewDir.Render()
}

func NewDirectory(fileList []string) *Directory {
	thumbs := make([]thumbnail.Thumbnail, len(fileList))
	for idx, f := range fileList {
		thumbs[idx] = *thumbnail.NewThumbnail(f, false, func() {
			log.Println(f)
		})
	}

	return &Directory{
		viewDir: thumbnail.NewDirectory(thumbs),
	}
}