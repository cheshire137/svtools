package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveFile_LoadsFromXmlPath(t *testing.T) {
	saveFile := NewSaveFile("fixtures/Anne_130176584")
	err := saveFile.Load()
	require.NoError(t, err)
}
