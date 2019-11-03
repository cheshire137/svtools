package models

type Color struct {
	R           int   `xml:"R"`
	G           int   `xml:"G"`
	B           int   `xml:"B"`
	A           int   `xml:"A"`
	PackedValue int64 `xml:"PackedValue"`
}
