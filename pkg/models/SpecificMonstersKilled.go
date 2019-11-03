package models

import "encoding/xml"

type SpecificMonstersKilled struct {
	XMLName xml.Name `xml:"specificMonstersKilled"`
	Items   []Item   `xml:"item"`
}
