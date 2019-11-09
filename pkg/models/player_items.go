package models

import (
	"encoding/xml"
	"strings"
)

type PlayerItems struct {
	XMLName xml.Name     `xml:"items"`
	Items   []PlayerItem `xml:"Item"`
}

func (pi *PlayerItems) Total() int {
	total := 0
	for _, item := range pi.Items {
		if !item.IsNil {
			total = total + 1
		}
	}
	return total
}

func (pi *PlayerItems) String() string {
	itemStrings := []string{}
	for _, item := range pi.Items {
		if !item.IsNil {
			itemStrings = append(itemStrings, item.String())
		}
	}
	return strings.Join(itemStrings, ", ")
}
