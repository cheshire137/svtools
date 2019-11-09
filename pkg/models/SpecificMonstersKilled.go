package models

import (
	"encoding/xml"
	"strings"
)

type SpecificMonstersKilled struct {
	XMLName xml.Name `xml:"specificMonstersKilled"`
	Items   []Item   `xml:"item"`
}

func (m *SpecificMonstersKilled) String() string {
	var sb strings.Builder
	for _, item := range m.Items {
		sb.WriteString(item.String())
	}
	return sb.String()
}
