package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItem(t *testing.T) {
	xmlItem := "<item><key><string>Green Slime</string></key><value><int>153</int></value></item>"
	var item Item

	err := xml.Unmarshal([]byte(xmlItem), &item)

	require.NoError(t, err)
	require.Equal(t, "Green Slime", item.Key.String)
	require.Equal(t, 153, item.Value.Int)
}
