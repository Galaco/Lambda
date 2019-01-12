package lib

type application struct {
	controllers []IController
	widgets []IWidget
}

// AddController adds a new controller to the
// current application.
func (app *application) AddController(controller IController) {
	controller.RegisterEventListeners()
	app.controllers = append(app.controllers, controller)
}

// AddWidget adds a widget to the application.
func (app *application) AddWidget(mod IWidget) {
	mod.Initialize()
	app.widgets = append(app.widgets, mod)
}

// Render draws the widgets registered with the application.
func (app *application) Render() {
	for _, controller := range app.controllers {
		controller.Render()
	}
	for _, mod := range app.widgets {
		mod.Render()
	}
}

// Update ensures that views and controllers can update
// themselves automatically.
func (app *application) Update() {
	for _, mod := range app.widgets {
		mod.Update()
	}
}

// NewApplication returns a new application.
func NewApplication() *application {
	return &application{}
}
