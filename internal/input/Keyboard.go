package input

import "github.com/vulkan-go/glfw/v3.3/glfw"

type Keyboard struct {
	keyMap map[Key]bool
}

func (keyboard *Keyboard) GlfwKeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		keyboard.keyMap[Key(key)] = true
	case glfw.Release:
		keyboard.keyMap[Key(key)] = false
	}
}

func (keyboard *Keyboard) IsKeyDown(key Key) bool {
	return keyboard.keyMap[key]
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		keyMap: map[Key]bool{},
	}
}
