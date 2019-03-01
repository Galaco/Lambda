package config

import (
	"encoding/json"
	"github.com/galaco/Lambda/internal/model"
	"io/ioutil"
)

// Load attempts to open and unmarshall
// json configuration
func Load(path string) (*model.Preferences, error) {
	data, err := ioutil.ReadFile(path)
	config := new(model.Preferences)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}

	return config, nil
}
