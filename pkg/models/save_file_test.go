package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveFile_LoadsFromXmlPath(t *testing.T) {
	saveFile := NewSaveFile("fixtures/Anne_130176584")
	err := saveFile.Load()
	require.NoError(t, err)
}

func TestSaveFile_CanSaveAsXmlFile(t *testing.T) {
	// Load a Stardew Valley save
	saveFile := NewSaveFile("fixtures/Anne_130176584")
	err := saveFile.Load()
	require.NoError(t, err)

	// Write a copy of the save file elsewhere on disk
	dupePath := "tmp-can-save-duplicate-test.xml"
	err = saveFile.Save(dupePath, false)
	require.NoError(t, err)
	defer func() {
		os.Remove(dupePath)
	}()

	// Should be able to load that copy without error
	dupeFile := NewSaveFile(dupePath)
	err = dupeFile.Load()
	require.NoError(t, err)
}
