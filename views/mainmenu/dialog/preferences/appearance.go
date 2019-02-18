package preferences

import (
	"github.com/inkyblackness/imgui-go"
)

type PageAppearance struct {
	themeLabels []string

	CurrentOption int
}

func (d *PageAppearance) Render() {
	imgui.Text("Theme")
	if imgui.BeginCombo("Theme", d.themeLabels[d.CurrentOption]) {
		for idx, label := range d.themeLabels {
			if imgui.Selectable(label) {
				d.CurrentOption = idx
			}
		}
		imgui.EndCombo()
	}
}

func NewPageAppearance() *PageAppearance {
	page := &PageAppearance{
		CurrentOption: 0,
	}

	page.themeLabels = []string{
		"Light",
		"Dark",
	}

	return page
}
