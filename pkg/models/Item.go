package models

import "encoding/xml"

type Item struct {
	XMLName xml.Name `xml:"item"`
	Key     keyNode
	Value   valueNode
}

type keyNode struct {
	XMLName xml.Name `xml:"key"`
	String  string   `xml:"string"`
}

type valueNode struct {
	XMLName xml.Name `xml:"value"`
	Int     int      `xml:"int"`
}
