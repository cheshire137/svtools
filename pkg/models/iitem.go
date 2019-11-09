package models

import (
	"encoding/xml"
	"fmt"
)

type Item struct {
	XMLName xml.Name `xml:"item"`
	Key     keyNode
	Value   valueNode
}

func (i *Item) String() string {
	return fmt.Sprintf("%s: %d", i.Key.String, i.Value.Int)
}

type keyNode struct {
	XMLName xml.Name `xml:"key"`
	String  string   `xml:"string"`
}

type valueNode struct {
	XMLName xml.Name `xml:"value"`
	Int     int      `xml:"int"`
}
