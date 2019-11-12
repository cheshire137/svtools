package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObject(t *testing.T) {
	xmlObject := `<Object><parentSheetIndex>128</parentSheetIndex><category>-9</category><name>Mushroom Box</name><specialItem>false</specialItem><hasBeenInInventory>false</hasBeenInInventory><SpecialVariable>0</SpecialVariable><DisplayName>Mushroom Box</DisplayName><Name>Mushroom Box</Name><Stack>1</Stack><isRecipe>false</isRecipe><bigCraftable>true</bigCraftable><type>Crafting</type><tileLocation><X>4</X><Y>5</Y></tileLocation><owner>0</owner><canBeSetDown>true</canBeSetDown><canBeGrabbed>true</canBeGrabbed><isHoedirt>false</isHoedirt><isSpawnedObject>false</isSpawnedObject><questItem>false</questItem><questId>0</questId><isOn>true</isOn><fragility>2</fragility><price>0</price><edibility>-300</edibility><stack>1</stack><quality>0</quality><setOutdoors>true</setOutdoors><setIndoors>true</setIndoors><readyForHarvest>true</readyForHarvest><showNextIndex>true</showNextIndex><flipped>false</flipped><hasBeenPickedUpByFarmer>false</hasBeenPickedUpByFarmer><isLamp>false</isLamp><heldObject><parentSheetIndex>281</parentSheetIndex><category>-81</category><name>Chanterelle</name><specialItem>false</specialItem><hasBeenInInventory>false</hasBeenInInventory><SpecialVariable>0</SpecialVariable><DisplayName>Chanterelle</DisplayName><Name>Chanterelle</Name><Stack>1</Stack><isRecipe>false</isRecipe><bigCraftable>false</bigCraftable><type>Basic</type><tileLocation><X>0</X><Y>0</Y></tileLocation><owner>0</owner><canBeSetDown>true</canBeSetDown><canBeGrabbed>true</canBeGrabbed><isHoedirt>false</isHoedirt><isSpawnedObject>false</isSpawnedObject><questItem>false</questItem><questId>0</questId><isOn>true</isOn><fragility>0</fragility><price>160</price><edibility>30</edibility><stack>1</stack><quality>0</quality><setOutdoors>false</setOutdoors><setIndoors>false</setIndoors><readyForHarvest>false</readyForHarvest><showNextIndex>false</showNextIndex><flipped>true</flipped><hasBeenPickedUpByFarmer>false</hasBeenPickedUpByFarmer><isLamp>false</isLamp><minutesUntilReady>0</minutesUntilReady><boundingBox><X>0</X><Y>0</Y><Width>64</Width><Height>64</Height><Location><X>0</X><Y>0</Y></Location></boundingBox><scale><X>0</X><Y>0</Y></scale><uses>0</uses><preservedParentSheetIndex>0</preservedParentSheetIndex></heldObject><minutesUntilReady>0</minutesUntilReady><boundingBox><X>256</X><Y>320</Y><Width>64</Width><Height>64</Height><Location><X>256</X><Y>320</Y></Location></boundingBox><scale><X>9.099997</X><Y>5.99999666</Y></scale><uses>0</uses><preservedParentSheetIndex>0</preservedParentSheetIndex></Object>`
	var object Object

	err := xml.Unmarshal([]byte(xmlObject), &object)

	require.NoError(t, err)
	require.Equal(t, 128, object.ParentSheetIndex)
	require.Equal(t, -9, object.Category)
	require.Equal(t, "Mushroom Box", object.Name)
	require.False(t, object.SpecialItem)
	require.False(t, object.HasBeenInInventory)
	require.Equal(t, 0, object.SpecialVariable)
	require.Equal(t, "Mushroom Box", object.DisplayName)
	require.Equal(t, "Mushroom Box", object.Name)
	require.Equal(t, 1, object.Stack)
	require.False(t, object.IsRecipe)
	require.True(t, object.BigCraftable)
	require.Equal(t, "Crafting", object.Type)
	require.Equal(t, 0, object.Owner)
	require.True(t, object.CanBeSetDown)
	require.True(t, object.CanBeGrabbed)
	require.False(t, object.IsHoedirt)
	require.False(t, object.IsSpawnedObject)
	require.False(t, object.QuestItem)
	require.Equal(t, 0, object.QuestID)
	require.True(t, object.IsOn)
	require.Equal(t, 2, object.Fragility)
	require.Equal(t, 0, object.Price)
	require.Equal(t, -300, object.Edibility)
	require.Equal(t, 1, object.Stack)
	require.Equal(t, 0, object.Quality)
	require.True(t, object.SetOutdoors)
	require.True(t, object.SetIndoors)
	require.True(t, object.IsReadyForHarvest)
	require.True(t, object.ShowNextIndex)
	require.False(t, object.IsFlipped)
	require.False(t, object.HasBeenPickedUpByFarmer)
	require.False(t, object.IsLamp)
	require.Equal(t, 0, object.MinutesUntilReady)
	require.Equal(t, 0, object.Uses)
	require.Equal(t, 0, object.PreservedParentSheetIndex)
}
