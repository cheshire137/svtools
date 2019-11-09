package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColor(t *testing.T) {
	xmlColor := "<tint><R>1</R><G>2</G><B>3</B><A>4</A><PackedValue>4294967295</PackedValue></tint>"
	var color Color

	err := xml.Unmarshal([]byte(xmlColor), &color)

	require.NoError(t, err)
	require.Equal(t, 1, color.R)
	require.Equal(t, 2, color.G)
	require.Equal(t, 3, color.B)
	require.Equal(t, 4, color.A)
	require.Equal(t, int64(4294967295), color.PackedValue)
}
