package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoundingBox(t *testing.T) {
	xmlBoundingBox := `<boundingBox><X>0</X><Y>0</Y><Width>64</Width><Height>64</Height><Location><X>0</X><Y>0</Y></Location></boundingBox>`
	var boundingBox BoundingBox

	err := xml.Unmarshal([]byte(xmlBoundingBox), &boundingBox)

	require.NoError(t, err)
	require.Equal(t, 0, boundingBox.X)
	require.Equal(t, 0, boundingBox.Y)
	require.Equal(t, 64, boundingBox.Width)
	require.Equal(t, 64, boundingBox.Height)
	require.NotNil(t, boundingBox.Location)
	require.Equal(t, 0, boundingBox.Location.X)
	require.Equal(t, 0, boundingBox.Location.Y)
}
