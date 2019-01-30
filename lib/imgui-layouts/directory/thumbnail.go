package directory

import (
	"github.com/inkyblackness/imgui-go"
)

type Thumbnail struct {
	label   string
	preview bool
	onClick func()
}

func (thumb *Thumbnail) Render() {
	if imgui.Selectable(thumb.label) {
		thumb.onClick()
	}
}

func NewThumbnail(label string, preview bool, onClick func()) *Thumbnail {
	return &Thumbnail{
		label:   label,
		preview: preview,
		onClick: onClick,
	}
}
