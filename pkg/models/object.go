package models

import "encoding/xml"

type Object struct {
	XMLName                   xml.Name    `xml:"Object"`
	ParentSheetIndex          int         `xml:"parentSheetIndex"`
	Category                  int         `xml:"category"`
	Name                      string      `xml:"name"`
	SpecialItem               bool        `xml:"specialItem"`
	HasBeenInInventory        bool        `xml:"hasBeenInInventory"`
	SpecialVariable           int         `xml:"SpecialVariable"`
	DisplayName               string      `xml:"DisplayName"`
	Name2                     string      `xml:"Name"`
	Stack                     int         `xml:"Stack"`
	IsRecipe                  bool        `xml:"isRecipe"`
	BigCraftable              bool        `xml:"bigCraftable"`
	Type                      string      `xml:"type"`
	TileLocation              Position    `xml:"tileLocation"`
	Owner                     int         `xml:"owner"`
	CanBeSetDown              bool        `xml:"canBeSetDown"`
	CanBeGrabbed              bool        `xml:"canBeGrabbed"`
	IsHoedirt                 bool        `xml:"isHoedirt"`
	IsSpawnedObject           bool        `xml:"isSpawnedObject"`
	QuestItem                 bool        `xml:"questItem"`
	QuestID                   int         `xml:"questId"`
	IsOn                      bool        `xml:"isOn"`
	Fragility                 int         `xml:"fragility"`
	Price                     int         `xml:"price"`
	Edibility                 int         `xml:"edibility"`
	Stack2                    int         `xml:"stack"`
	Quality                   int         `xml:"quality"`
	SetOutdoors               bool        `xml:"setOutdoors"`
	SetIndoors                bool        `xml:"setIndoors"`
	IsReadyForHarvest         bool        `xml:"readyForHarvest"`
	ShowNextIndex             bool        `xml:"showNextIndex"`
	IsFlipped                 bool        `xml:"flipped"`
	HasBeenPickedUpByFarmer   bool        `xml:"hasBeenPickedUpByFarmer"`
	IsLamp                    bool        `xml:"isLamp"`
	HeldObject                HeldObject  `xml:"heldObject"`
	MinutesUntilReady         int         `xml:"minutesUntilReady"`
	BoundingBox               BoundingBox `xml:"boundingBox"`
	Scale                     Scale       `xml:"scale"`
	Uses                      int         `xml:"uses"`
	PreservedParentSheetIndex int         `xml:"preservedParentSheetIndex"`
}
