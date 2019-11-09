package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerItem(t *testing.T) {
	xmlPlayerItem := `<Item xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="FishingRod"><category>-99</category><name>Fishing Rod</name><specialItem>false</specialItem><hasBeenInInventory>true</hasBeenInInventory><SpecialVariable>0</SpecialVariable><DisplayName>Fiberglass Rod</DisplayName><Name>Fiberglass Rod</Name><Stack>1</Stack><initialParentTileIndex>10</initialParentTileIndex><currentParentTileIndex>10</currentParentTileIndex><indexOfMenuItemView>10</indexOfMenuItemView><stackable>false</stackable><instantUse>false</instantUse><upgradeLevel>2</upgradeLevel><numAttachmentSlots>2</numAttachmentSlots><attachments><Object xsi:nil="true" /><Object xsi:nil="true" /></attachments><BaseName>Fishing Rod</BaseName><InitialParentTileIndex>10</InitialParentTileIndex><IndexOfMenuItemView>10</IndexOfMenuItemView><InstantUse>false</InstantUse><Stackable>false</Stackable></Item>`
	var item PlayerItem

	err := xml.Unmarshal([]byte(xmlPlayerItem), &item)

	require.NoError(t, err)
	require.False(t, item.IsNil)
	require.Equal(t, "FishingRod", item.Type)
	require.Equal(t, -99, item.Category)
	require.Equal(t, "Fishing Rod", item.Name)
	require.Equal(t, false, item.SpecialItem)
	require.Equal(t, true, item.HasBeenInInventory)
	require.Equal(t, 0, item.SpecialVariable)
	require.Equal(t, "Fiberglass Rod", item.DisplayName)
	require.Equal(t, "Fiberglass Rod", item.SpecificName)
	require.Equal(t, 1, item.Stack)
	require.Equal(t, 10, item.InitialParentTileIndex)
	require.Equal(t, 10, item.CurrentParentTileIndex)
	require.Equal(t, 10, item.IndexOfMenuItemView)
	require.Equal(t, false, item.Stackable)
	require.Equal(t, false, item.InstantUse)
	require.Equal(t, 2, item.UpgradeLevel)
	require.Equal(t, 2, item.NumAttachmentSlots)
	require.Equal(t, "Fishing Rod", item.BaseName)
	require.Equal(t, 10, item.InitialParentTileIndex)
	require.Equal(t, 10, item.IndexOfMenuItemView)
	require.Equal(t, false, item.InstantUse)
	require.Equal(t, false, item.Stackable)
}
