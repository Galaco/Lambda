package scenegraph

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
)

type widget struct {
}

func (mod *widget) Initialize() {

}

func (mod *widget) Render() {
	imgui.SetNextWindowPos(imgui.Vec2{X:0, Y:48})
	if imgui.BeginV("Objects", nil, imgui.WindowFlagsNoResize | imgui.WindowFlagsNoMove | imgui.WindowFlagsNoBringToFrontOnFocus) {
		imgui.BeginChild("Scrolling")
			for i := 0; i < 200; i++ {
				imgui.Text(fmt.Sprintf("%d: Some text", i))
			}
		imgui.EndChild()
		imgui.End()
	}
}

func (mod *widget) Update() {

}

func NewWidget() *widget {
	return &widget{}
}
