package dialog

import (
	"github.com/sqweek/dialog"
)

func FileOpen() (string, error) {
	filename, err := dialog.File().Filter("Vmf map file", "vmf").Load()
	if err != nil {
		dialog.Message("%s", "Failed to open file").Error()
		return "", err
	}

	return filename, nil
}

func FileSave() (string, error) {
	return dialog.File().Save()
}
