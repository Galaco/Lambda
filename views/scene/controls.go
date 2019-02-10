package scene

import "github.com/inkyblackness/imgui-go"

type Controls struct {
	isContextCurrent bool

	Actions struct {
		Forward bool
		Backwards bool
		Left bool
		Right bool
		RotX int
		RotY int
	}
}

func (controls *Controls) Enable() {
	controls.isContextCurrent = true
}

func (controls *Controls) Disable() {
	controls.isContextCurrent = false
}

func (controls *Controls) Update() {
	if !controls.isContextCurrent {
		return
	}

	if imgui.IsKeyPressed(imgui.KeyUpArrow) {
		controls.Actions.Forward = true
	} else {
		controls.Actions.Forward = false
	}
	if imgui.IsKeyPressed(imgui.KeyLeftArrow) {
		controls.Actions.Left = true
	} else {
		controls.Actions.Left = false
	}
	if imgui.IsKeyPressed(imgui.KeyDownArrow) {
		controls.Actions.Backwards = true
	} else {
		controls.Actions.Backwards = false
	}
	if imgui.IsKeyPressed(imgui.KeyRightArrow) {
		controls.Actions.Right = true
	} else {
		controls.Actions.Right = false
	}
}

func newControls() *Controls {
	return &Controls{}
}