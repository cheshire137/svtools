package models

import (
	"encoding/xml"
)

type BoundingBox struct {
	XMLName  xml.Name `xml:"boundingBox"`
	X        int      `xml:"X"`
	Y        int      `xml:"Y"`
	Width    int      `xml:"Width"`
	Height   int      `xml:"Height"`
	Location Position `xml:"Location"`
}
