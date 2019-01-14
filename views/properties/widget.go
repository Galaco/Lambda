package properties

import (
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/lib/event"
	"github.com/galaco/source-tools-common/entity"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	selectedEntity *entity.Entity
	keyValueViews []keyValue
}

func (mod *Widget) Initialize() {
	event.Singleton().Subscribe(events.TypeEntitySelected, mod.selectedEntityChanged)
}

func (mod *Widget) Render() {
	if imgui.BeginV("Properties", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		imgui.BeginChild("Scrolling")
		for _,kv := range mod.keyValueViews {
			kv.Render()
		}

		imgui.EndChild()
		imgui.End()
	}
}

func (mod *Widget) Update() {

}

func (mod *Widget) selectedEntityChanged(received event.IEvent) {
	evt := received.(*events.EntitySelected)
	mod.selectedEntity = evt.Target()
	mod.keyValueViews = make([]keyValue, 0)

	kv := mod.selectedEntity.EPairs
	for kv != nil {
		mod.keyValueViews = append(mod.keyValueViews, newKeyValue(kv.Key, kv.Value, false))
		kv = kv.Next
	}
}

func NewWidget() *Widget {
	return &Widget{}
}