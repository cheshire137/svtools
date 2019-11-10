package models

import "encoding/xml"

type GameLocation struct {
	XMLName                     xml.Name `xml:"GameLocation"`
	Type                        string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Name                        string   `xml:"name"`
	WaterColor                  Color    `xml:"waterColor"`
	IsFarm                      bool     `xml:"isFarm"`
	IsOutdoors                  bool     `xml:"isOutdoors"`
	IsGreenhouse                bool     `xml:"isGreenhouse"`
	IsStructure                 bool     `xml:"isStructure"`
	IgnoreDebrisWeather         bool     `xml:"ignoreDebrisWeather"`
	IgnoreOutdoorLighting       bool     `xml:"ignoreOutdoorLighting"`
	IgnoreLights                bool     `xml:"ignoreLights"`
	TreatAsOutdoors             bool     `xml:"treatAsOutdoors"`
	NumberOfSpawnedObjectsOnMap int      `xml:"numberOfSpawnedObjectsOnMap"`
}
