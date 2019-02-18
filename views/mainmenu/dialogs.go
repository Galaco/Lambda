package mainmenu

import (
	"github.com/galaco/Lambda/views/mainmenu/dialog"
	"io/ioutil"
)

const dialogWidth = 480
const dialogHeight = 640

func openFile() string {
	filename, err := dialog.FileOpen()
	if err != nil {
		filename = "./ze_bioshock_v6_4.vmf"
		return ""
	}
	return filename
}

func saveFile(filename string, data string) (err error) {
	// Saving a new file
	if filename == "" {
		filename, err = dialog.FileSave()
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(filename, []byte(data), 755)
}
