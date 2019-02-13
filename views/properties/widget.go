package properties

import (
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/events"
	"github.com/galaco/Lambda/model"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/ui/imgui-layouts/keyvalues"
	"github.com/galaco/source-tools-common/entity"
	"github.com/inkyblackness/imgui-go"
	"log"
	"strconv"
)

type Widget struct {
	dispatcher *event.Dispatcher
	model      *model.Model

	keyValueView *keyvalues.View
	selectedEntity *entity.Entity
	//keyValueViews  []keyValue
}

func (widget *Widget) Initialize() {
	widget.dispatcher.Subscribe(events.TypeEntityNodeSelected, widget.selectedEntityChanged)
	widget.dispatcher.Subscribe(events.TypeSolidNodeSelected, widget.selectedSolidChanged)
	widget.dispatcher.Subscribe(events.TypeSceneClosed, widget.sceneClosed)
}

func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(w - 320), Y: 48})
	imgui.SetNextWindowSize(imgui.Vec2{X: 320, Y: float32(h - 48)})
	if imgui.BeginV("Properties", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoBringToFrontOnFocus) {
		imgui.BeginChild("Scrolling")
		widget.keyValueView.Render()
		//for _, kv := range widget.keyValueViews {
		//	kv.Render()
		//}

		imgui.EndChild()
		imgui.End()
	}
}

func (widget *Widget) selectedEntityChanged(received event.IEvent) {
	widget.keyValueView = keyvalues.NewKeyValues()
	evt := received.(*events.EntityNodeSelected)
	widget.selectedEntity = widget.model.Project.Vmf.Entities().FindByKeyValue("id", strconv.Itoa(evt.Id))

	kv := widget.selectedEntity.EPairs
	for kv != nil {
		widget.keyValueView.AddKeyValue(keyvalues.NewKeyValue(kv.Key, kv.Value, func(k, v string) {
			log.Println(k + " " + v)
		}))
		kv = kv.Next
	}
}

func (widget *Widget) selectedSolidChanged(received event.IEvent) {
	widget.keyValueView = keyvalues.NewKeyValues()
	evt := received.(*events.SolidNodeSelected)
	widget.keyValueView.AddKeyValue(keyvalues.NewKeyValue("solid id", strconv.FormatInt(int64(evt.Id), 10), func(k, v string) {
		log.Println(k + " " + v)
	}))
}

func (widget *Widget) sceneClosed(received event.IEvent) {
	widget.keyValueView = keyvalues.NewKeyValues()
}

func NewWidget(dispatcher *event.Dispatcher, model *model.Model) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		model:      model,
		keyValueView: keyvalues.NewKeyValues(),
	}
}
