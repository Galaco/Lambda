package dialog

import (
	"github.com/galaco/Lambda/ui/imgui-layouts/columns"
	"github.com/galaco/Lambda/views/mainmenu/dialog/preferences"
	"github.com/inkyblackness/imgui-go"
)

type Preferences struct {
	Dialog

	twoPanel *columns.View
	sidebar *preferences.Sidebar

	pages []preferences.IPage
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
}

func (d *Preferences) renderTab() {
	d.currentPage.Render()
}

func NewPreferences() *Preferences{
	dialog := &Preferences{
		Dialog: Dialog{
			name: "Preferences",
		},
		sidebar: preferences.NewNavbar(),
		twoPanel: columns.NewColumns(2),
	}

	dialog.pages = append(dialog.pages, &preferences.PageGeneral{})
	dialog.currentPage = dialog.pages[0]

	dialog.twoPanel.SetColumnContents(0, dialog.sidebar.Render, columns.NewColumnWidth(100, false))
	dialog.twoPanel.SetColumnContents(1, dialog.renderTab, nil)

	return dialog
}