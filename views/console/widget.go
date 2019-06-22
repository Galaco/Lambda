package console

import (
	lambdaFS "github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda/internal/event"
	"github.com/galaco/Lambda/internal/graphics"
	"github.com/galaco/Lambda/internal/model"
	"github.com/galaco/Lambda/internal/ui"
	"github.com/galaco/Lambda/internal/ui/context"
	"github.com/galaco/Lambda/internal/ui/imgui-layouts"
	"github.com/galaco/Lambda/internal/ui/imgui-layouts/master/rule"
	"github.com/inkyblackness/imgui-go"
)

type Widget struct {
	masterPanel     *imgui_layouts.Panel
	dispatcher      *event.Dispatcher
	graphicsAdapter graphics.Adapter
	fileSystem      *lambdaFS.FileSystem
	model           *model.Model
}

func (widget *Widget) Initialize() {

}

func (widget *Widget) Render(ctx *context.Context) {
	w, h := ctx.Window().GetSize()
	if widget.masterPanel.Start("Console", w, h) {
		for _, msg := range widget.model.Logs.GetLogs(model.LogTypeApplication) {
			imgui.BeginChild("ConsoleScrolling")
			imgui.Text(msg)
			imgui.EndChild()
		}
		widget.masterPanel.End()
	}
}

func (widget *Widget) selectedEntityChanged(received event.Dispatchable) {
}

func NewWidget(dispatcher *event.Dispatcher, fileSystem *lambdaFS.FileSystem, model *model.Model) *Widget {
	return &Widget{
		dispatcher: dispatcher,
		fileSystem: fileSystem,
		model:      model,
		masterPanel: imgui_layouts.NewPanel().
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampBottom, 0)).
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampLeft, int(320.0*ui.DPIScale()))).
			WithDisplayRule(rule.NewRuleClampToEdge(rule.ClampRight, int(320.0*ui.DPIScale()))).
			WithDisplayRule(rule.NewRuleFixedHeight(320, false, 0)),
	}
}
