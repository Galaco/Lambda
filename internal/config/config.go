package config

import (
	"encoding/json"
	"io/ioutil"
)

// Project configuration properties
// Engine needs to know where to locate its game data
type Config struct {
	GameDirectory string
}

// @TODO Implement something nicer than this scoped variable
var config Config

// Get returns (kind-of) static config object
func Singleton() *Config {
	return &config
}

// Load attempts to open and unmarshall
// json configuration
func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}
