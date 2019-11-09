package models

import "encoding/xml"

type StringKey struct {
	XMLName xml.Name `xml:"key"`
	String  string   `xml:"string"`
}
