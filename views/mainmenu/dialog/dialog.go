package dialog

import "github.com/inkyblackness/imgui-go"

type Dialog struct {
	isOpen bool
	name string
}

func (d *Dialog) IsOpen() bool {
	return d.isOpen
}

func (d *Dialog) Open() {
	d.isOpen = true
}

func (d *Dialog) close() {
	imgui.CloseCurrentPopup()
	d.isOpen = false
}
