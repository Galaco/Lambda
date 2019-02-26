package list

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"log"
)

// row is a single row in a list
type row struct {
	Id    int
	label string

	onClick func(id int)
	selected bool
}

// render renders the row as a imgui::Selectable.
// It also dispatches a notification when the row is selected.
func (item *row) render() bool {
	if imgui.SelectableV(item.label, item.selected, 0, imgui.Vec2{}) {
		item.selected = true
		item.onClick(item.Id)
		log.Println(fmt.Sprintf("%d selected", item.Id))

		return true
	}
	return false
}

// newRow returns a new row
func newRow(id int, label string, onClick func(id int)) row {
	format := "%d %s"

	return row{
		Id:      id,
		label:   fmt.Sprintf(format, id, label),
		onClick: onClick,
	}
}
