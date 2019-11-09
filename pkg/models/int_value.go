package models

import "encoding/xml"

type IntValue struct {
	XMLName xml.Name `xml:"value"`
	Int     int      `xml:"int"`
}
