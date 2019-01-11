package events

const TypeWindowResized = "WindowResized"
const TypeWindowClosed = "WindowClosed"

type windowResized struct {
	actionType string
	x, y       int
}

func (act *windowResized) Type() string {
	return TypeWindowResized
}

func NewWindowResized(x int, y int) *windowResized {
	return &windowResized{
		x: x,
		y: y,
	}
}

type windowClosed struct {
	actionType string
	x, y       int
}

func (act *windowClosed) Type() string {
	return TypeWindowClosed
}

func NewWindowClosed() *windowClosed {
	return &windowClosed{}
}
