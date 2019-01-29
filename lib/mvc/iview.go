package mvc

import "github.com/vulkan-go/glfw/v3.3/glfw"

// IWidget is essentially a view/subset of a view,
// View being a view like in any MVC framework.
type IView interface {
	// Initialize prepares any one-time setup
	Initialize()
	// Render renders the current view
	Render(*glfw.Window)
	// Update updates the current view
	Update()
	// Correctly destroy any resources (most likely gpu)
	Destroy()
}
