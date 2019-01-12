package scenegraph

import (
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	nodes []Item
}

func (mod *Widget) Initialize() {

}

func (mod *Widget) Render() {
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 48})
	if imgui.BeginV("Objects", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		imgui.BeginChild("Scrolling")
		for _,row := range mod.nodes {
			row.Render()
		}
		imgui.EndChild()
		imgui.End()
	}
}

func (mod *Widget) Update() {

}

func (mod *Widget) AddNode(id int, title string) {
	mod.nodes = append(mod.nodes, NewItem(id, title))
}

func NewWidget() *Widget {
	return &Widget{}
}
