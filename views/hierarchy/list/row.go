package list

import (
	"fmt"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/inkyblackness/imgui-go"
	"log"
)

// row is a single row in a list
type row struct {
	Id         int
	Classname  string
	TargetName string
	label      string
}

// render renders the row as a imgui::Selectable.
// It also dispatches a notification when the row is selected.
func (item *row) render() {
	if imgui.Selectable(item.label) {
		event.Singleton().Dispatch(events.NewSceneNodeSelected(item.Id))
		log.Println(fmt.Sprintf("%d selected", item.Id))
	}
}

// newRow returns a new row
func newRow(id int, classname string, targetname string) row {
	format := "%d %s"
	if targetname != "" {
		format += " : %s"
	} else {
		format += "%s"
	}

	return row{
		Id:         id,
		Classname:  classname,
		TargetName: targetname,
		label:      fmt.Sprintf(format, id, classname, targetname),
	}
}
