package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cheshire137/svtools/pkg/models"
)

// SaveFile represents a Stardew Valley save file.
type SaveFile struct {
	Path string
	Data *models.SaveGame
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
	var saveGame models.SaveGame
	err = xml.Unmarshal(bytes, &saveGame)
	if err != nil {
		return err
	}
	save.Data = &saveGame
	return nil
}
