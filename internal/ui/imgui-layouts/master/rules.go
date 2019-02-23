package master

type Rulable interface {
	Resolve(pane *Panel)
}
