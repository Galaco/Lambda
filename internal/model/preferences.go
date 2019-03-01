package model

type Preferences struct {
	General struct {
		GameDirectory string `json:"gameDirectory"`
	} `json:"general"`
	Appearance struct {
		Theme         int     `json:"theme"`
		WindowScaling float32 `json:"windowScaling"`
	} `json:"appearance"`
	ViewPort struct {
		FOV int `json:"fov"`
	} `json:"viewport"`
}
