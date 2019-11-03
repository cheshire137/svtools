package models

import "encoding/xml"

type PlayerItems struct {
	XMLName xml.Name     `xml:"items"`
	Items   []PlayerItem `xml:"Item"`
}
