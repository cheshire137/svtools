package models

import (
	"encoding/xml"
	"strings"
)

type SpecificMonstersKilled struct {
	XMLName xml.Name `xml:"specificMonstersKilled"`
	Items   []Item   `xml:"item"`
}

func (m *SpecificMonstersKilled) Counts() map[string]int {
	counts := make(map[string]int, len(m.Items))
	for _, item := range m.Items {
		counts[item.Key.String] = item.Value.Int
	}
	return counts
}

func (m *SpecificMonstersKilled) String() string {
	var sb strings.Builder
	for _, item := range m.Items {
		sb.WriteString(item.String())
	}
	return sb.String()
}
