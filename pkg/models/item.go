package models

import (
	"encoding/xml"
	"fmt"
)

type Item struct {
	XMLName xml.Name `xml:"item"`
	Key     StringKey
	Value   IntValue
}

func (i *Item) String() string {
	return fmt.Sprintf("%s: %d", i.Key.String, i.Value.Int)
}
