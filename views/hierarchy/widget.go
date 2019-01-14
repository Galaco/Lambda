package hierarchy

import (
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/inkyblackness/imgui-go"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Widget struct {
	allTab allEntities
	currentViewMode int
	currentGroupMode int
}

func (mod *Widget) Initialize() {
	event.Singleton().Subscribe(events.TypeEntityCreated, mod.newEntityCreated)
}

func (mod *Widget) Render(window *glfw.Window) {
	_, h := window.GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 18})
	imgui.SetNextWindowSize(imgui.Vec2{X: 320, Y: float32(h - 18)})
	if imgui.BeginV("Hierarchy", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus|imgui.WindowFlagsMenuBar) {
		mod.renderMenuBar()
		mod.allTab.Render()
		imgui.End()
	}
}

func (mod *Widget) Update() {

}

func (mod *Widget) renderMenuBar() {
	if imgui.BeginMenuBar() {
		if imgui.BeginMenu("View") {
			if imgui.MenuItem("All") {
				if mod.currentViewMode != entityFilterNone {
					mod.currentViewMode = entityFilterNone
					mod.allTab.Filter(entityFilterNone)
				}
				// View all
			}
			if imgui.MenuItem("Point entities") {
				if mod.currentViewMode != entityFilterPointOnly {
					mod.currentViewMode = entityFilterPointOnly
					mod.allTab.Filter(entityFilterPointOnly)
				}
				// View collapsible by classname
			}
			if imgui.MenuItem("Brush entities") {
				if mod.currentViewMode != entityFilterBrushOnly {
					mod.currentViewMode = entityFilterBrushOnly
					mod.allTab.Filter(entityFilterBrushOnly)
				}
				// View collapsible by classname
			}
			if imgui.MenuItem("Props") {
				if mod.currentViewMode != entityFilterPropOnly {
					mod.currentViewMode = entityFilterPropOnly
					mod.allTab.Filter(entityFilterPropOnly)
				}
				// View collapsible by classname
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Group") {
			if imgui.MenuItem("None") {
				mod.currentGroupMode = 0
				// No grouping
			}
			if imgui.MenuItem("By Classname") {
				mod.currentGroupMode = 1
				// View collapsible by classname
			}
			imgui.EndMenu()
		}
		imgui.EndMenuBar()
	}
}

func (mod *Widget) newEntityCreated(received event.IEvent) {
	ent := received.(*events.EntityCreated).Target()
	mod.allTab.AddEntity(
		ent.IntForKey("id"),
		ent.ValueForKey("classname"),
		ent.ValueForKey("targetname"))
}

func NewWidget() *Widget {
	return &Widget{}
}
