package models

import (
	"encoding/xml"
	"strings"
)

type PlayerItems struct {
	XMLName xml.Name     `xml:"items"`
	Items   []PlayerItem `xml:"Item"`
}

func (items *PlayerItems) String() string {
	itemStrings := []string{}
	for _, item := range items.Items {
		if !item.IsNil {
			itemStrings = append(itemStrings, item.String())
		}
	}
	return strings.Join(itemStrings, ", ")
}
