package model

import "github.com/galaco/Lambda/internal/model/preferences"

type Preferences struct {
	General preferences.General `json:"general"`
	Appearance preferences.Appearance `json:"appearance"`
	Viewport preferences.Viewport `json:"viewport"`
}