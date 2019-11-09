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
	var sb strings.Builder
	for _, item := range items.Items {
		sb.WriteString(item.String())
	}
	return sb.String()
}
