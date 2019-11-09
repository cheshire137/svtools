package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpecificMonstersKilled(t *testing.T) {
	xmlSpecificMonstersKilled := "<specificMonstersKilled><item><key><string>Green Slime</string></key><value><int>153</int></value></item><item><key><string>Bug</string></key><value><int>73</int></value></item><item><key><string>Duggy</string></key><value><int>9</int></value></item><item><key><string>Fly</string></key><value><int>45</int></value></item><item><key><string>Rock Crab</string></key><value><int>10</int></value></item><item><key><string>Grub</string></key><value><int>103</int></value></item><item><key><string>Stone Golem</string></key><value><int>33</int></value></item><item><key><string>Bat</string></key><value><int>35</int></value></item><item><key><string>Frost Bat</string></key><value><int>111</int></value></item><item><key><string>Frost Jelly</string></key><value><int>200</int></value></item><item><key><string>Dust Spirit</string></key><value><int>167</int></value></item><item><key><string>Ghost</string></key><value><int>25</int></value></item><item><key><string>Skeleton</string></key><value><int>45</int></value></item><item><key><string>Shadow Brute</string></key><value><int>43</int></value></item><item><key><string>Shadow Shaman</string></key><value><int>31</int></value></item><item><key><string>Lava Bat</string></key><value><int>36</int></value></item><item><key><string>Sludge</string></key><value><int>282</int></value></item><item><key><string>Lava Crab</string></key><value><int>10</int></value></item><item><key><string>Metal Head</string></key><value><int>26</int></value></item><item><key><string>Squid Kid</string></key><value><int>14</int></value></item><item><key><string>Big Slime</string></key><value><int>53</int></value></item><item><key><string>Mummy</string></key><value><int>10</int></value></item><item><key><string>Serpent</string></key><value><int>116</int></value></item><item><key><string>Iridium Crab</string></key><value><int>18</int></value></item><item><key><string>Carbon Ghost</string></key><value><int>20</int></value></item><item><key><string>Iridium Bat</string></key><value><int>64</int></value></item></specificMonstersKilled>"
	var specMonKill SpecificMonstersKilled

	err := xml.Unmarshal([]byte(xmlSpecificMonstersKilled), &specMonKill)

	require.NoError(t, err)
	counts := specMonKill.Counts()
	require.Equal(t, 153, counts["Green Slime"])
	require.Equal(t, 73, counts["Bug"])
	require.Equal(t, 9, counts["Duggy"])
	require.Equal(t, 45, counts["Fly"])
	require.Equal(t, 10, counts["Rock Crab"])
	require.Equal(t, 103, counts["Grub"])
	require.Equal(t, 33, counts["Stone Golem"])
	require.Equal(t, 35, counts["Bat"])
	require.Equal(t, 111, counts["Frost Bat"])
	require.Equal(t, 200, counts["Frost Jelly"])
	require.Equal(t, 167, counts["Dust Spirit"])
	require.Equal(t, 25, counts["Ghost"])
	require.Equal(t, 45, counts["Skeleton"])
	require.Equal(t, 43, counts["Shadow Brute"])
	require.Equal(t, 31, counts["Shadow Shaman"])
	require.Equal(t, 36, counts["Lava Bat"])
	require.Equal(t, 282, counts["Sludge"])
	require.Equal(t, 10, counts["Lava Crab"])
	require.Equal(t, 26, counts["Metal Head"])
	require.Equal(t, 14, counts["Squid Kid"])
	require.Equal(t, 53, counts["Big Slime"])
	require.Equal(t, 10, counts["Mummy"])
	require.Equal(t, 116, counts["Serpent"])
	require.Equal(t, 18, counts["Iridium Crab"])
	require.Equal(t, 20, counts["Carbon Ghost"])
	require.Equal(t, 64, counts["Iridium Bat"])
}
