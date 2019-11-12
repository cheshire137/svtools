package models

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// SaveFile represents a Stardew Valley save file.
type SaveFile struct {
	Path string
	Data *SaveGame
}

// NewSaveFile returns a SaveFile for the Stardew Valley save at the given file path.
func NewSaveFile(path string) *SaveFile {
	return &SaveFile{Path: path}
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
	var saveGame SaveGame
	err = xml.Unmarshal(bytes, &saveGame)
	if err != nil {
		return err
	}
	s.Data = &saveGame
	return nil
}

// ToXML returns a byte array of the save file as XML.
func (s *SaveFile) ToXML(indent bool) ([]byte, error) {
	saveData := s.Data
	if indent {
		prefix := ""
		return xml.MarshalIndent(saveData, prefix, "  ")
	}
	return xml.Marshal(saveData)
}

// Save the save file to the specified path as an XML file.
func (s *SaveFile) Save(path string, indent bool) error {
	bytes, err := s.ToXML(indent)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0644)
}
