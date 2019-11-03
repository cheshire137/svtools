package xml

import (
	"fmt"
	"os"
)

// SaveFile represents a Stardew Valley save file.
type SaveFile struct {
	Path string
}

// Load opens a save file for editing.
func (save *SaveFile) Load() error {
	xmlFile, err := os.Open(save.Path)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	fmt.Println("opened", save.Path)

	return nil
}
