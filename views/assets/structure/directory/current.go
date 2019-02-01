package directory

import (
	"github.com/galaco/Lambda/ui/imgui-layouts/directory"
	"log"
	"strings"
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
		dirName := strings.Split(f, "/")
		thumbs[idx] = *directory.NewThumbnail(dirName[len(dirName)-1], false, func() {
			log.Println(dirName[len(dirName)-1])
		})
	}

	return &Directory{
		viewDir: directory.NewDirectory(thumbs),
	}
}
