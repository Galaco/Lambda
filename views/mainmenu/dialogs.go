package mainmenu

import (
	"github.com/galaco/Lambda/views/mainmenu/dialog"
	"io/ioutil"
)

const dialogWidth = 480
const dialogHeight = 640

func openFile() (string, error) {
	return dialog.FileOpen()
}

func saveFile(filename string, data string) (err error) {
	// Saving a new file
	if filename == "" {
		filename, err = dialog.FileSave()
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(filename, []byte(data), 0755)
}
