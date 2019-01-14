package lib

// IController is a standard Controller interface you'd
// find in any MVC style application.
type IController interface {
	// RegisterEventListeners registers a controllers event handlers.
	RegisterEventListeners()
}
