package models

import "fmt"

type Color struct {
	R           int   `xml:"R"`
	G           int   `xml:"G"`
	B           int   `xml:"B"`
	A           int   `xml:"A"`
	PackedValue int64 `xml:"PackedValue"`
}

func (c *Color) String() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %d)", c.R, c.G, c.B, c.A)
}
