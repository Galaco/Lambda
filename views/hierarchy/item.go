package hierarchy

import (
	"fmt"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/inkyblackness/imgui-go"
	"log"
)

type Item struct {
	Id    int
	Classname string
	TargetName string
	label string
}

func (item *Item) Render() {
	if imgui.Selectable(item.label) {
		event.Singleton().Dispatch(events.NewSceneNodeSelected(item.Id))
		log.Println(fmt.Sprintf("%d selected", item.Id))
	}
}

func NewItem(id int, classname string, targetname string) Item {
	format := "%d %s"
	if targetname != "" {
		format += " : %s"
	} else {
		format += "%s"
	}

	return Item{
		Id:    id,
		Classname: classname,
		TargetName: targetname,
		label: fmt.Sprintf(format, id, classname, targetname),
	}


}
