package vgui

type Element interface {
	Draw()
	Resize(parentWidth, parentHeight float64)
}
