package models

import (
	"encoding/xml"
)

type HeldObject struct {
	XMLName                   xml.Name    `xml:"heldObject"`
	BoundingBox               BoundingBox `xml:"boundingBox"`
	Scale                     Position    `xml:"scale"`
	ParentSheetIndex          int         `xml:"parentSheetIndex"`
	Category                  int         `xml:"category"`
	Name                      string      `xml:"name"`
	Name2                     string      `xml:"Name"`
	SpecialItem               bool        `xml:"specialItem"`
	HasBeenInInventory        bool        `xml:"hasBeenInInventory"`
	SpecialVariable           int         `xml:"SpecialVariable"`
	DisplayName               string      `xml:"DisplayName"`
	Stack                     int         `xml:"Stack"`
	Stack2                    int         `xml:"stack"`
	IsRecipe                  bool        `xml:"isRecipe"`
	BigCraftable              bool        `xml:"bigCraftable"`
	Type                      string      `xml:"type"`
	TileLocation              Position    `xml:"tileLocation"`
	Owner                     int         `xml:"owner"`
	CanBeSetDown              bool        `xml:"canBeSetDown"`
	CanBeGrabbed              bool        `xml:"canBeGrabbed"`
	IsHoeDirt                 bool        `xml:"isHoedirt"`
	IsSpawnedObject           bool        `xml:"isSpawnedObject"`
	QuestItem                 bool        `xml:"questItem"`
	QuestID                   int         `xml:"questId"`
	IsOn                      bool        `xml:"isOn"`
	Fragility                 int         `xml:"fragility"`
	Price                     int         `xml:"price"`
	Edibility                 int         `xml:"edibility"`
	Quality                   int         `xml:"quality"`
	SetOutdoors               bool        `xml:"setOutdoors"`
	SetIndoors                bool        `xml:"setIndoors"`
	IsReadyForHarvest         bool        `xml:"readyForHarvest"`
	ShowNextIndex             bool        `xml:"showNextIndex"`
	IsFlipped                 bool        `xml:"flipped"`
	HasBeenPickedUpByFarmer   bool        `xml:"hasBeenPickedUpByFarmer"`
	IsLamp                    bool        `xml:"isLamp"`
	MinutesUntilReady         int         `xml:"minutesUntilReady"`
	Uses                      int         `xml:"uses"`
	PreservedParentSheetIndex int         `xml:"preservedParentSheetIndex"`
}
