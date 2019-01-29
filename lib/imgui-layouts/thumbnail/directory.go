package thumbnail

import (
	"github.com/inkyblackness/imgui-go"
)

const thumbnailSizeDefault = 4

type Directory struct {
	thumbnails []Thumbnail
	thumbnailSize int

	displayColumns [][]*Thumbnail
}

func (dir *Directory) Render() {
	imgui.BeginChild("Scrolling")
	imgui.BeginColumns("Files", dir.thumbnailSize)
	for i := 0; i < dir.thumbnailSize; i++ {
		for _,thumb := range dir.displayColumns[i] {
			thumb.Render()
		}
		imgui.NextColumn()
	}
	imgui.EndChild()
}

func (dir *Directory) SetThumbnailSize(size int) {
	if size < 1 || size > 32 {
		return
	}
	dir.thumbnailSize = size
	dir.reorderThumbnails()
}

func (dir *Directory) reorderThumbnails() {
	cols := make([][]*Thumbnail, dir.thumbnailSize)
	for i :=0; i< dir.thumbnailSize; i++ {
		cols[i] = []*Thumbnail{}
	}
	for idx := range dir.thumbnails {
		for i := dir.thumbnailSize; i >= 1; i-- {
			if idx % i == 0 {
				cols[i-1] = append(cols[i-1], &dir.thumbnails[idx])
				break
			}
		}
	}

	dir.displayColumns = cols
}

func NewDirectory(thumbnails []Thumbnail) *Directory {
	dir := &Directory{
		thumbnails:thumbnails,
		thumbnailSize: thumbnailSizeDefault,
	}
	dir.reorderThumbnails()
	return dir
}