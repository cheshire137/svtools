package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecipe(t *testing.T) {
	xmlRecipe := "<item><key><string>Stir Fry</string></key><value><int>1</int></value></item>"
	var recipe Recipe

	err := xml.Unmarshal([]byte(xmlRecipe), &recipe)

	require.NoError(t, err)
	require.Equal(t, "Stir Fry", recipe.Key.String)
	require.Equal(t, 1, recipe.Value.Int)
}
