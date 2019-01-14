package lib

import "github.com/vulkan-go/glfw/v3.3/glfw"

type application struct {
	controllers []IController
	views       []IView
	models      []IModel
}

// AddController adds a new controller to the
// current application.
func (app *application) AddController(controller IController) {
	controller.RegisterEventListeners()
	app.controllers = append(app.controllers, controller)
}

func (app *application) AddView(view IView) {
	view.Initialize()
	app.views = append(app.views, view)
}

func (app *application) AddModel(model IModel) {
	app.models = append(app.models, model)
}

// Render draws the widgets registered with the application.
func (app *application) Render(window *glfw.Window) {
	for _, view := range app.views {
		view.Render(window)
	}
}

// Update ensures that views and controllers can update
// themselves automatically.
func (app *application) Update() {
}

// NewApplication returns a new application.
func NewApplication() *application {
	return &application{}
}
