package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeldObject(t *testing.T) {
	xmlHeldObject := `<heldObject><parentSheetIndex>281</parentSheetIndex><category>-81</category><name>Chanterelle</name><specialItem>false</specialItem><hasBeenInInventory>false</hasBeenInInventory><SpecialVariable>0</SpecialVariable><DisplayName>Chanterelle</DisplayName><Name>Chanterelle</Name><Stack>1</Stack><isRecipe>false</isRecipe><bigCraftable>false</bigCraftable><type>Basic</type><tileLocation><X>0</X><Y>0</Y></tileLocation><owner>0</owner><canBeSetDown>true</canBeSetDown><canBeGrabbed>true</canBeGrabbed><isHoedirt>false</isHoedirt><isSpawnedObject>false</isSpawnedObject><questItem>false</questItem><questId>0</questId><isOn>true</isOn><fragility>0</fragility><price>160</price><edibility>30</edibility><stack>1</stack><quality>0</quality><setOutdoors>false</setOutdoors><setIndoors>false</setIndoors><readyForHarvest>false</readyForHarvest><showNextIndex>false</showNextIndex><flipped>true</flipped><hasBeenPickedUpByFarmer>false</hasBeenPickedUpByFarmer><isLamp>false</isLamp><minutesUntilReady>0</minutesUntilReady><boundingBox><X>0</X><Y>0</Y><Width>64</Width><Height>64</Height><Location><X>0</X><Y>0</Y></Location></boundingBox><scale><X>0</X><Y>0</Y></scale><uses>0</uses><preservedParentSheetIndex>0</preservedParentSheetIndex></heldObject>`
	var heldObject HeldObject

	err := xml.Unmarshal([]byte(xmlHeldObject), &heldObject)

	require.NoError(t, err)
	require.Equal(t, 281, heldObject.ParentSheetIndex)
	require.Equal(t, -81, heldObject.Category)
	require.Equal(t, "Chanterelle", heldObject.Name)
	require.False(t, heldObject.SpecialItem)
	require.False(t, heldObject.HasBeenInInventory)
	require.Equal(t, 0, heldObject.SpecialVariable)
	require.Equal(t, "Chanterelle", heldObject.DisplayName)
	require.Equal(t, "Chanterelle", heldObject.Name)
	require.Equal(t, 1, heldObject.Stack)
	require.False(t, heldObject.IsRecipe)
	require.False(t, heldObject.BigCraftable)
	require.Equal(t, "Basic", heldObject.Type)
	require.NotNil(t, heldObject.TileLocation)
	require.Equal(t, 0, heldObject.TileLocation.X)
	require.Equal(t, 0, heldObject.TileLocation.Y)
	require.Equal(t, 0, heldObject.Owner)
	require.True(t, heldObject.CanBeSetDown)
	require.True(t, heldObject.CanBeGrabbed)
	require.False(t, heldObject.IsHoeDirt)
	require.False(t, heldObject.IsSpawnedObject)
	require.False(t, heldObject.QuestItem)
	require.Equal(t, 0, heldObject.QuestID)
	require.True(t, heldObject.IsOn)
	require.Equal(t, 0, heldObject.Fragility)
	require.Equal(t, 160, heldObject.Price)
	require.Equal(t, 30, heldObject.Edibility)
	require.Equal(t, 1, heldObject.Stack)
	require.Equal(t, 0, heldObject.Quality)
	require.False(t, heldObject.SetOutdoors)
	require.False(t, heldObject.SetIndoors)
	require.False(t, heldObject.IsReadyForHarvest)
	require.False(t, heldObject.ShowNextIndex)
	require.True(t, heldObject.IsFlipped)
	require.False(t, heldObject.HasBeenPickedUpByFarmer)
	require.False(t, heldObject.IsLamp)
	require.Equal(t, 0, heldObject.MinutesUntilReady)
	require.NotNil(t, heldObject.BoundingBox)
	require.Equal(t, 0, heldObject.BoundingBox.X)
	require.Equal(t, 0, heldObject.BoundingBox.Y)
	require.Equal(t, 64, heldObject.BoundingBox.Width)
	require.Equal(t, 64, heldObject.BoundingBox.Height)
	require.NotNil(t, heldObject.BoundingBox.Location)
	require.Equal(t, 0, heldObject.BoundingBox.Location.X)
	require.Equal(t, 0, heldObject.BoundingBox.Location.Y)
	require.Equal(t, 0, heldObject.Uses)
	require.Equal(t, 0, heldObject.PreservedParentSheetIndex)
}
