package models

import "encoding/xml"

type PlayerItem struct {
	XMLName                 xml.Name `xml:"Item"`
	Type                    string   `xml:"xsi:type,attr"`
	Category                int      `xml:"category"`
	Name                    string   `xml:"name"`
	SpecificName            string   `xml:"Name"`
	DisplayName             string   `xml:"DisplayName"`
	BaseName                string   `xml:"BaseName"`
	SpecialItem             bool     `xml:"specialItem"`
	HasBeenInInventory      bool     `xml:"hasBeenInInventory"`
	SpecialVariable         int      `xml:"SpecialVariable"`
	Stack                   int      `xml:"Stack"`
	InitialParentTileIndex  int      `xml:"initialParentTileIndex"`
	InitialParentTileIndex2 int      `xml:"InitialParentTileIndex"`
	CurrentParentTileIndex  int      `xml:"currentParentTileIndex"`
	IndexOfMenuItemView     int      `xml:"indexOfMenuItemView"`
	IndexOfMenuItemView2    int      `xml:"IndexOfMenuItemView"`
	Stackable               bool     `xml:"stackable"`
	Stackable2              bool     `xml:"Stackable"`
	InstantUse              bool     `xml:"instantUse"`
	InstantUse2             bool     `xml:"InstantUse"`
	UpgradeLevel            int      `xml:"upgradeLevel"`
	NumAttachmentSlots      int      `xml:"numAttachmentSlots"`
	WaterCanMax             int      `xml:"waterCanMax"`
	WaterLeft               int      `xml:"WaterLeft"`
}
