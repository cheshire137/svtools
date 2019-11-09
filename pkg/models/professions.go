package models

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Professions struct {
	XMLName xml.Name `xml:"professions"`
	Values  []int    `xml:"int"`
}

func (p *Professions) String() string {
	strValues := make([]string, len(p.Values))
	for i, value := range p.Values {
		strValues[i] = fmt.Sprintf("%d", value)
	}
	return strings.Join(strValues, ", ")
}
