package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPosition(t *testing.T) {
	xmlPosition := "<Location><X>50</X><Y>100</Y></Location>"
	var position Position

	err := xml.Unmarshal([]byte(xmlPosition), &position)

	require.NoError(t, err)
	require.Equal(t, 50, position.X)
	require.Equal(t, 100, position.Y)
}
