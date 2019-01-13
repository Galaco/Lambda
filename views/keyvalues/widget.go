package keyvalues

import (
	"github.com/galaco/source-tools-common/entity"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	selectedEntity *entity.Entity
	keyValueViews []keyvalue
}

func (mod *Widget) Initialize() {

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

func (mod *Widget) SetActiveEntity(selected *entity.Entity) {
	mod.selectedEntity = selected
	mod.keyValueViews = make([]keyvalue, 0)

	kv := mod.selectedEntity.EPairs
	for kv != nil {
		mod.keyValueViews = append(mod.keyValueViews, newKeyValue(kv.Key, kv.Value, false))
		kv = kv.Next
	}
}

func NewWidget() *Widget {
	return &Widget{}
}