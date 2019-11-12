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

	// Read the original contents
	originalContents, err := saveFile.Read()
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

	// Should match the original contents of the save file
	dupeContents, err := dupeFile.Read()
	require.NoError(t, err)

	saveTestArtifact(originalContents, "test-output/original-save.xml", t)
	saveTestArtifact(dupeContents, "test-output/duplicate-save.xml", t)

	require.Equal(t, originalContents, dupeContents)
}

func saveTestArtifact(contents string, path string, t *testing.T) {
	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		require.NoError(t, err)
	}
	file, err := os.Create(path)
	require.NoError(t, err)
	defer file.Close()
	file.WriteString(contents)
	file.Sync()
}
