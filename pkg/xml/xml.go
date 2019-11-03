package xml

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// SaveFile represents a Stardew Valley save file.
type SaveFile struct {
	Path string
}

// Load opens a save file for editing.
func (save *SaveFile) Load() error {
	info, err := os.Stat(save.Path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New(save.Path + " is a directory")
	}
	xmlFile, err := os.Open(save.Path)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	fmt.Println("opened", save.Path)
	bytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}
	return nil
}
