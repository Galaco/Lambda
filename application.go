package main

import (
	"github.com/galaco/Lambda/internal/event"
	"github.com/galaco/Lambda/internal/filesystem"
	"github.com/galaco/Lambda/internal/filesystem/exporters"
	"github.com/galaco/Lambda/internal/filesystem/importers"
	"github.com/galaco/Lambda/internal/graphics"
	"github.com/galaco/Lambda/internal/input"
	"github.com/galaco/Lambda/internal/model"
	"github.com/galaco/Lambda/internal/ui"
	"github.com/galaco/Lambda/internal/ui/context"
	"github.com/galaco/Lambda/views/assets"
	"github.com/galaco/Lambda/views/console"
	"github.com/galaco/Lambda/views/hierarchy"
	"github.com/galaco/Lambda/views/mainmenu"
	"github.com/galaco/Lambda/views/properties"
	"github.com/galaco/Lambda/views/scene"
	"github.com/vulkan-go/glfw/v3.3/glfw"
)

// Application is the main Lambda application
type Application struct {
	uiContext       *context.Context
	GraphicsAdapter graphics.Adapter

	// Tools
	FileSystem      filesystem.FileSystem
	EventDispatcher *event.Dispatcher
	VmfImporter     *importers.VmfImporter
	VmfExporter     *exporters.VmfExporter

	// Model
	Model *model.Model

	// Control
	Keyboard *input.Keyboard

	//Views
	assetsView           *assets.Widget
	hierarchyView        *hierarchy.Widget
	mainMenuView         *mainmenu.Widget
	entityPropertiesView *properties.Widget
	scenePreviewView     *scene.Widget
	consoleView          *console.Widget
}

// InitializeUIContext prepares the graphics and visual context.
func (app *Application) InitializeUIContext() *context.Context {
	// Window & OpenGL
	app.uiContext = context.NewContext(app.GraphicsAdapter)

	return app.uiContext
}

// Render renders the loaded widgets
func (app *Application) Render() {
	app.scenePreviewView.RenderScene(app.uiContext)

	app.uiContext.Imgui().NewFrame()
	//app.assetsView.Render(app.uiContext)
	app.hierarchyView.Render(app.uiContext)
	app.mainMenuView.Render(app.uiContext)
	app.entityPropertiesView.Render(app.uiContext)
	app.scenePreviewView.Render(app.uiContext)
	app.consoleView.Render(app.uiContext)
}

// Update processes dispatched events.
func (app *Application) Update() {
	app.EventDispatcher.Process()

	app.scenePreviewView.Update(1000.0 / 60)
}

// InitializeViews loads and prepares application views/widgets
func (app *Application) InitializeViews() {
	app.assetsView = assets.NewWidget(app.EventDispatcher, app.FileSystem)
	app.hierarchyView = hierarchy.NewWidget(app.EventDispatcher)
	app.mainMenuView = mainmenu.NewWidget(app.EventDispatcher, app.Model, app.VmfImporter, app.VmfExporter)
	app.entityPropertiesView = properties.NewWidget(app.EventDispatcher, app.Model)
	app.scenePreviewView = scene.NewWidget(app.EventDispatcher, app.FileSystem, app.Keyboard, app.GraphicsAdapter)
	app.consoleView = console.NewWidget(app.EventDispatcher, app.FileSystem, app.Model)

	app.assetsView.Initialize()
	app.hierarchyView.Initialize()
	app.mainMenuView.Initialize()
	app.entityPropertiesView.Initialize()
	app.scenePreviewView.Initialize()
	app.consoleView.Initialize()
}

// InitializeGUITheme set the initial imgui layout and colour scheme
func (app *Application) InitializeGUITheme() {
	ui.ApplyImguiStyles(app.Model.Preferences.Appearance.Theme)
}

// Close shuts doen the application
func (app *Application) Close() {
	defer glfw.Terminate()
	defer app.uiContext.Close()
}
