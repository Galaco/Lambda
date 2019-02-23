package events

const TypeWindowClosed = "WindowClosed"

type windowClosed struct {
	x, y int
}

func (act *windowClosed) Type() string {
	return TypeWindowClosed
}

func NewWindowClosed() *windowClosed {
	return &windowClosed{}
}
