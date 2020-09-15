package vgui

// Panel
type Panel struct {
	X      float64
	Y      float64
	Width  float64
	Height float64

	enabled      bool
	proportional bool

	children []*Panel
	elements []Element
}

// Proportional
func (panel *Panel) Proportional() bool {
	return panel.proportional
}

// SetProportional
func (panel *Panel) SetProportional(proportional bool) {
	panel.proportional = proportional
}

// Enabled
func (panel *Panel) Enabled() bool {
	return panel.enabled
}

// SetEnabled
func (panel *Panel) SetEnabled(enabled bool) {
	panel.enabled = enabled
}

// Children returns all panels that are a child of this panel
func (panel *Panel) Children() []*Panel {
	return panel.children
}

// AddChild adds a new child to this panel
func (panel *Panel) AddChild(child *Panel) {
	panel.children = append(panel.children, child)
}

func (panel *Panel) AddElement(element Element) {
	panel.elements = append(panel.elements, element)
}

// Draw
func (panel *Panel) Draw() {
	if !panel.enabled {
		return
	}
	for _, p := range panel.elements {
		p.Draw()
	}
	for _, child := range panel.Children() {
		child.Draw()
	}
}

// Resize
func (panel *Panel) Resize(parentWidth, parentHeight float64) {
	for _, p := range panel.elements {
		p.Resize(parentWidth, parentHeight)
	}
	for _, child := range panel.Children() {
		child.Resize(parentWidth, parentHeight)
	}
}

// NewChildPanel
func (panel *Panel) NewChildPanel(X, Y, Width, Height float64, Enabled bool) *Panel {
	p := &Panel{
		X:       X,
		Y:       Y,
		Width:   Width,
		Height:  Height,
		enabled: Enabled,
	}

	panel.children = append(panel.children, p)

	return p
}
