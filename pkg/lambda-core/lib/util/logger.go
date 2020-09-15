package util

import (
	"github.com/galaco/loggy"
)

var logger *loggy.Loggy

func Logger() *loggy.Loggy {
	if logger == nil {
		logger = loggy.NewLoggy()
	}

	return logger
}
