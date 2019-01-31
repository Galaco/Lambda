package main

import (
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda/event"
	"github.com/galaco/Lambda/project"
	"github.com/galaco/Lambda/filesystem/importers"
	"github.com/galaco/Lambda/ui/context"
	"github.com/galaco/Lambda/views/assets"
	"github.com/galaco/Lambda/views/hierarchy"
	"github.com/galaco/Lambda/views/mainmenu"
	"github.com/galaco/Lambda/views/properties"
	"github.com/galaco/Lambda/views/ribbon"
	"github.com/galaco/Lambda/views/scene"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type Application struct {
	uiContext *context.Context

	// Tools
	FileSystem *filesystem.FileSystem
	EventDispatcher *event.Dispatcher
	VmfImporter *importers.VmfImporter

	// Model
	Model *project.Model

	//Views
	assetsView *assets.Widget
	hierarchyView *hierarchy.Widget
	mainMenuView *mainmenu.Widget
	entityPropertiesView *properties.Widget
	toolRibbonView *ribbon.Widget
	scenePreviewView *scene.Widget
}

func (app *Application) InitializeUIContext() *context.Context{
	// Window & OpenGL
	app.uiContext = context.NewContext()

	return app.uiContext
}

func (app *Application) Render() {
	app.assetsView.Render(app.uiContext)
	app.hierarchyView.Render(app.uiContext)
	app.mainMenuView.Render(app.uiContext)
	app.entityPropertiesView.Render(app.uiContext)
	app.toolRibbonView.Render(app.uiContext)
	//app.scenePreviewView.Render(app.uiContext)
}

func (app *Application) InitializeViews() {
	app.assetsView = assets.NewWidget(app.EventDispatcher, app.FileSystem)
	app.hierarchyView = hierarchy.NewWidget(app.EventDispatcher)
	app.mainMenuView = mainmenu.NewWidget(app.EventDispatcher, app.VmfImporter, app.Model)
	app.entityPropertiesView = properties.NewWidget(app.EventDispatcher, app.Model)
	app.toolRibbonView = ribbon.NewWidget()
	//app.scenePreviewView = scene.NewWidget()

	app.assetsView.Initialize()
	app.hierarchyView.Initialize()
	app.mainMenuView.Initialize()
	app.entityPropertiesView.Initialize()
	app.toolRibbonView.Initialize()
	//app.scenePreviewView.Initialize()
}

func (app *Application) InitializeGUITheme() {
	applyImguiStyles()
}

func (app *Application) Close() {
	defer glfw.Terminate()
	defer app.uiContext.Close()
}