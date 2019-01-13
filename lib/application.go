package lib

type application struct {
	controllers []IController
}

// AddController adds a new controller to the
// current application.
func (app *application) AddController(controller IController) {
	controller.RegisterEventListeners()
	app.controllers = append(app.controllers, controller)
}


// Render draws the widgets registered with the application.
func (app *application) Render() {
	for _, controller := range app.controllers {
		controller.Render()
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
