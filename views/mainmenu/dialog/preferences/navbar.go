package preferences

import (
	"github.com/inkyblackness/imgui-go"
	"strings"
)

type Sidebar struct {
	options []string

	currentPage string
}

func (nav *Sidebar) Render() {
	for _, option := range nav.options {
		selected := false
		if strings.ToLower(option) == nav.currentPage {
			selected = true
		}
		if imgui.SelectableV(option, selected, 0, imgui.Vec2{}) {
			nav.currentPage = strings.ToLower(option)
		}
	}
}

func (nav *Sidebar) CurrentTab() string {
	return nav.currentPage
}

func NewNavbar() *Sidebar {
	return &Sidebar{
		options: []string{
			"General",
			"Appearance",
		},
		currentPage: "general",
	}
}
