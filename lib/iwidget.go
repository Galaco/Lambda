package lib

// IWidget is essentially a view/subset of a view,
// View being a view like in any MVC framework.
type IWidget interface {
	// Initialize prepares any one-time setup
	Initialize()
	// Render renders the current view
	Render()
	// Update updates the current view
	Update()
}
