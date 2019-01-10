package lib


type application struct {
	modules []IWidget
}

func (app *application) RegisterModule(mod IWidget) {
	mod.Initialize()
	app.modules = append(app.modules, mod)
}

func (app *application) Render() {
	for _, mod := range app.modules {
		mod.Render()
	}
}

func (app *application) Update() {
	for _, mod := range app.modules {
		mod.Update()
	}
}

func NewApplication() *application {
	return &application{}
}