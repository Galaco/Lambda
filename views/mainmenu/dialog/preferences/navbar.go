package preferences

import "github.com/inkyblackness/imgui-go"

type Sidebar struct {
	options []string
}

func (nav *Sidebar) Render() {
	for _,option := range nav.options {
		imgui.Selectable(option)
	}
}

func NewNavbar() *Sidebar {
	return &Sidebar{
		options: []string{
			"General",
			"Appearance",
		},
	}
}