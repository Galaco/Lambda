package vgui

type Button struct {
	text string
}

func (btn *Button) Draw() {

}

func (btn *Button) Resize(parentWidth, parentHeight float64) {

}

func NewButton(label string) *Button {
	return &Button{
		text: label,
	}
}
