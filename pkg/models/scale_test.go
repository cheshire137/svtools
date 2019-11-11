package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScale(t *testing.T) {
	xmlScale := `<scale><X>9.099997</X><Y>5.99999666</Y></scale>`
	var scale Scale

	err := xml.Unmarshal([]byte(xmlScale), &scale)

	require.NoError(t, err)
	require.Equal(t, float32(9.099997), scale.X)
	require.Equal(t, float32(5.99999666), scale.Y)
}
