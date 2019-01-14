package hierarchy

import (
	"fmt"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
	nodes []Item
}

func (mod *Widget) Initialize() {
	event.Singleton().Subscribe(events.TypeEntityCreated, mod.newEntityCreated)
}

func (mod *Widget) Render(window *glfw.Window) {
	_, h := window.GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 16})
	imgui.SetNextWindowSize(imgui.Vec2{X: 320, Y: float32(h - 16)})
	if imgui.BeginV("Hierarchy", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		imgui.BeginChild("Scrolling")
		for _, row := range mod.nodes {
			row.Render()
		}
		imgui.EndChild()
		imgui.End()
	}
}

func (mod *Widget) Update() {

}

func (mod *Widget) newEntityCreated(received event.IEvent) {
	ent := received.(*events.EntityCreated).Target()
	mod.nodes = append(mod.nodes, NewItem(
		ent.IntForKey("id"),
		fmt.Sprintf("%d %s: %s", ent.IntForKey("id"), ent.ValueForKey("classname"), ent.ValueForKey("targetname"))))
}

func NewWidget() *Widget {
	return &Widget{}
}
