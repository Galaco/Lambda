package dialog

import (
	"github.com/galaco/Lambda/internal/event"
	"github.com/galaco/Lambda/internal/events"
	"github.com/galaco/Lambda/internal/model"
	"github.com/galaco/Lambda/internal/ui/imgui-layouts/columns"
	"github.com/galaco/Lambda/views/mainmenu/dialog/preferences"
	"github.com/inkyblackness/imgui-go"
)

type Preferences struct {
	Dialog
	dispatcher *event.Dispatcher

	twoPanel *columns.View
	sidebar  *preferences.Sidebar

	pages       map[string]preferences.IPage
	currentPage preferences.IPage
}

func (d *Preferences) Render(width, height int) {
	if !d.IsOpen() {
		return
	}

	imgui.PushStyleVarVec2(imgui.StyleVarWindowMinSize, imgui.Vec2{X: float32(width), Y: float32(height)})
	imgui.OpenPopup(d.name)
	if imgui.BeginPopupModal(d.name) {
		d.twoPanel.Render(width, height)

		imgui.EndPopup()
	}
	imgui.PopStyleVar()

	d.currentPage = d.pages[d.sidebar.CurrentTab()]
}

func (d *Preferences) renderTab() {
	d.currentPage.Render()

	if imgui.Button("Save") {
		evt := events.NewPreferencesUpdated()
		evt.Appearance.Theme = d.pages["appearance"].(*preferences.PageAppearance).CurrentOption
		d.dispatcher.Dispatch(evt)
	}
	imgui.SameLine()

	if imgui.Button("Cancel") {
		d.close()
	}
}

func NewPreferences(dispatch *event.Dispatcher, model *model.Preferences) *Preferences {
	dialog := &Preferences{
		Dialog: Dialog{
			name: "Preferences",
		},
		dispatcher: dispatch,

		sidebar:  preferences.NewNavbar(),
		twoPanel: columns.NewColumns(2),
		pages:    map[string]preferences.IPage{},
	}

	dialog.pages["general"] = &preferences.PageGeneral{}
	dialog.pages["appearance"] = preferences.NewPageAppearance(model.Appearance.Theme)

	dialog.currentPage = dialog.pages[dialog.sidebar.CurrentTab()]

	_ = dialog.twoPanel.SetColumnContents(0, dialog.sidebar.Render, columns.NewColumnWidth(120, false))
	_ = dialog.twoPanel.SetColumnContents(1, dialog.renderTab, nil)

	return dialog
}
