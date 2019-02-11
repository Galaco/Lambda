package preferences

import "github.com/inkyblackness/imgui-go"

type PageAppearance struct {
	themeOptions []string
	themeLabels []string
}

func (d *PageAppearance) Render() {
	imgui.Text("Theme")
		if imgui.BeginCombo("Theme", d.themeLabels[0]) {
			for _,label := range d.themeLabels {
				imgui.Selectable(label)
			}
			imgui.EndCombo()
		}
}

func NewPageAppearance() *PageAppearance {
	page := &PageAppearance{}

	page.themeOptions = []string{
		"0",
		"1",
	}

	page.themeLabels = []string{
		"Light",
		"Dark",
	}

	return page
}