package list

import (
	"fmt"
	"github.com/galaco/lambda-core/lib/util"
	"github.com/inkyblackness/imgui-go"
)

// row is a single row in a list
type row struct {
	Id    int
	label string

	onClick  func(id int)
	selected bool
}

// render renders the row as a imgui::Selectable.
// It also dispatches a notification when the row is selected.
func (item *row) render() bool {
	if imgui.SelectableV(item.label, item.selected, 0, imgui.Vec2{}) {
		item.selected = true
		item.onClick(item.Id)
		util.Logger().Notice(fmt.Sprintf("%d selected", item.Id))

		return true
	}
	return false
}

// newRow returns a new row
func newRow(id int, label string, onClick func(id int)) row {
	format := "%s##%d"

	return row{
		Id:      id,
		label:   fmt.Sprintf(format, label, id),
		onClick: onClick,
	}
}
