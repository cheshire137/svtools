package models

import "encoding/xml"

type Scale struct {
	XMLName xml.Name `xml:"scale"`
	X       float32  `xml:"X"`
	Y       float32  `xml:"Y"`
}
