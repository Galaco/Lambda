package hierarchy

import (
	"fmt"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/inkyblackness/imgui-go"
	"log"
)

type Item struct {
	Id int
	Label string
}

func (item *Item) Render() {
	if imgui.Selectable(item.Label) {
		event.Singleton().Dispatch(events.NewSceneNodeSelected(item.Id))
		log.Println(fmt.Sprintf("%d selected", item.Id))
	}
}

func NewItem(id int, label string) Item {
	return Item{
		Id: id,
		Label: label,
	}
}