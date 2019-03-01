package config

import (
	"encoding/json"
	"github.com/galaco/Lambda/internal/model"
	"io/ioutil"
)

type Config struct {
	Preferences model.Preferences `json:"preferences"`
}

// Load attempts to open and unmarshall
// json configuration
func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	config := new(Config)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}

	return config, nil
}
