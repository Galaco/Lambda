package preferences

type General struct {
	GameDirectory string `json:"gameDirectory"`
}

type Appearance struct {
	Theme         int     `json:"theme"`
	WindowScaling float32 `json:"windowScaling"`
}

type Viewport struct {
	FOV int `json:"fov"`
}
