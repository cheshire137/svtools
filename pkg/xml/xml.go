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
func (s *SaveFile) Load() error {
	info, err := os.Stat(s.Path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New(s.Path + " is a directory")
	}
	xmlFile, err := os.Open(s.Path)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	fmt.Println("opened", s.Path)
	bytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}
	var saveGame models.SaveGame
	err = xml.Unmarshal(bytes, &saveGame)
	if err != nil {
		return err
	}
	s.Data = &saveGame
	return nil
}

func (s *SaveFile) Save(path string) error {
	prefix := ""
	indent := "  "
	bytes, err := xml.MarshalIndent(s, prefix, indent)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0644)
}
