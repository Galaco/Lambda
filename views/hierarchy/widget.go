package hierarchy

import (
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/ui/imgui-layouts"
	"github.com/galaco/Lambda/ui/imgui-layouts/master/rule"
	"github.com/galaco/Lambda/views/hierarchy/list"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	masterPanel *imgui_layouts.Panel

	dispatcher *event.Dispatcher

	list      entityList
	solidList entityList
}

func (widget *Widget) Initialize() {
	widget.dispatcher.Subscribe(events.TypeEntityCreated, widget.newEntityCreated)
	widget.dispatcher.Subscribe(events.TypeNewSolidCreated, widget.newSolidCreated)
	widget.dispatcher.Subscribe(events.TypeSceneClosed, widget.sceneClosed)
}

func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
	if widget.masterPanel.Start("Hierarchy", w, h) {
		widget.renderMenuBar()
		widget.list.getFiltered().Render()
		widget.solidList.getFiltered().Render()

		widget.masterPanel.End()
	}
	//_, h := ctx.Window().GetSize()
	//imgui.SetNextWindowPos(imgui.Vec2{X: 0, Y: 48})
	//imgui.SetNextWindowSize(imgui.Vec2{X: 320, Y: float32(h - 368)})
	//if imgui.BeginV("Hierarchy", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus|imgui.WindowFlagsMenuBar) {
	//	widget.renderMenuBar()
	//	widget.list.getFiltered().Render()
	//	widget.solidList.getFiltered().Render()
	//	imgui.End()
	//}
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
		ent.ValueForKey("classname")+" "+ent.ValueForKey("targetname"),
		func(id int) {
			widget.dispatcher.Dispatch(events.NewEntityNodeSelected(id))
		})
}

func (widget *Widget) newSolidCreated(received event.IEvent) {
	ent := received.(*events.NewSolidCreated).Target()
	widget.solidList.addEntity(
		ent.Id,
		"Solid",
		func(id int) {
			widget.dispatcher.Dispatch(events.NewSolidNodeSelected(id))
		})
}

func (widget *Widget) sceneClosed(received event.IEvent) {
	widget.list = entityList{}
	widget.solidList = entityList{}
}

func NewWidget(dispatcher *event.Dispatcher) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		masterPanel: imgui_layouts.NewPanel().
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampTop, 48)).
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampLeft, 0)).
			WithDisplayRule(rule.NewRuleFixedWidth(320)),
	}
}
