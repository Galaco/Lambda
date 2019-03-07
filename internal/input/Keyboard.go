package input

import "github.com/vulkan-go/glfw/v3.3/glfw"

// Keyboard provides the current state of the keyboard
type Keyboard struct {
	keyMap map[Key]bool
}

// GlfwKeyCallback is the glfw library callback function for handling
// keyboard input
func (keyboard *Keyboard) GlfwKeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		keyboard.keyMap[Key(key)] = true
	case glfw.Release:
		keyboard.keyMap[Key(key)] = false
	}
}

// IsKeyDown returns whether a particular key is pressed.
func (keyboard *Keyboard) IsKeyDown(key Key) bool {
	return keyboard.keyMap[key]
}

// NewKeyboard returns a new keyboard.
func NewKeyboard() *Keyboard {
	return &Keyboard{
		keyMap: map[Key]bool{},
	}
}
