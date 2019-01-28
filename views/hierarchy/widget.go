package hierarchy

import (
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/Lambda/views/hierarchy/list"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
	list entityList
}

func (widget *Widget) Initialize() {
	event.Singleton().Subscribe(events.TypeEntityCreated, widget.newEntityCreated)
}

func (widget *Widget) Render(window *glfw.Window) {
	_, h := window.GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 48})
	imgui.SetNextWindowSize(imgui.Vec2{X: 320, Y: float32(h - 48)})
	if imgui.BeginV("Hierarchy", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus|imgui.WindowFlagsMenuBar) {
		widget.renderMenuBar()
		widget.list.getFiltered().Render()
		imgui.End()
	}
}

func (widget *Widget) Update() {
}

func (widget *Widget) Destroy() {

}

func (widget *Widget) renderMenuBar() {
	if imgui.BeginMenuBar() {
		if imgui.BeginMenu("View") {
			if imgui.MenuItem("All") {
				widget.list.setFilterMode(list.EntityFilterNone)
				// View all
			}
			if imgui.MenuItem("Point entities") {
				widget.list.setFilterMode(list.EntityFilterPointOnly)
				// View collapsible by classname
			}
			if imgui.MenuItem("Brush entities") {
				widget.list.setFilterMode(list.EntityFilterBrushOnly)
				// View collapsible by classname
			}
			if imgui.MenuItem("Props") {
				widget.list.setFilterMode(list.EntityFilterPropOnly)
				// View collapsible by classname
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Group") {
			if imgui.MenuItem("None") {
				// No grouping
			}
			if imgui.MenuItem("By Classname") {
				// View collapsible by classname
			}
			imgui.EndMenu()
		}
		imgui.EndMenuBar()
	}
}

func (widget *Widget) newEntityCreated(received event.IEvent) {
	ent := received.(*events.EntityCreated).Target()
	widget.list.addEntity(
		ent.IntForKey("id"),
		ent.ValueForKey("classname"),
		ent.ValueForKey("targetname"))
}

func NewWidget() *Widget {
	return &Widget{}
}
