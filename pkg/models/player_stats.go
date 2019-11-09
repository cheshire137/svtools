package models

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type PlayerStats struct {
	XMLName                xml.Name `xml:"stats"`
	SeedsSown              int      `xml:"seedsSown"`
	ItemsShipped           int      `xml:"itemsShipped"`
	ItemsCooked            int      `xml:"itemsCooked"`
	ItemsCrafted           int      `xml:"itemsCrafted"`
	ChickenEggsLayed       int      `xml:"chickenEggsLayed"`
	DuckEggsLayed          int      `xml:"duckEggsLayed"`
	CowMilkProduced        int      `xml:"cowMilkProduced"`
	GoatMilkProduced       int      `xml:"goatMilkProduced"`
	RabbitWoolProduced     int      `xml:"rabbitWoolProduced"`
	SheepWoolProduced      int      `xml:"sheepWoolProduced"`
	CheeseMade             int      `xml:"cheeseMade"`
	GoatCheeseMade         int      `xml:"goatCheeseMade"`
	TrufflesFound          int      `xml:"trufflesFound"`
	StoneGathered          int      `xml:"stoneGathered"`
	RocksCrushed           int      `xml:"rocksCrushed"`
	DirtHoed               int      `xml:"dirtHoed"`
	GiftsGiven             int      `xml:"giftsGiven"`
	TimesUnconscious       int      `xml:"timesUnconscious"`
	TimesUnconscious2      int      `xml:"TimesUnconscious"`
	AverageBedtime         int      `xml:"averageBedtime"`
	TimesFished            int      `xml:"timesFished"`
	FishCaught             int      `xml:"fishCaught"`
	BouldersCracked        int      `xml:"bouldersCracked"`
	BouldersCracked2       int      `xml:"BouldersCracked"`
	StumpsChopped          int      `xml:"stumpsChopped"`
	StepsTaken             int      `xml:"stepsTaken"`
	TotalMonstersKilled    int      `xml:"monstersKilled"`
	TotalMonstersKilled2   int      `xml:"MonstersKilled"`
	MonstersKilled         SpecificMonstersKilled
	DiamondsFound          int `xml:"diamondsFound"`
	PrismaticShardsFound   int `xml:"prismaticShardsFound"`
	OtherPreciousGemsFound int `xml:"otherPreciousGemsFound"`
	CaveCarrotsFound       int `xml:"caveCarrotsFound"`
	CopperFound            int `xml:"copperFound"`
	IronFound              int `xml:"ironFound"`
	CoalFound              int `xml:"coalFound"`
	CoinsFound             int `xml:"coinsFound"`
	GoldFound              int `xml:"goldFound"`
	IridiumFound           int `xml:"iridiumFound"`
	BarsSmelted            int `xml:"barsSmelted"`
	BeveragesMade          int `xml:"beveragesMade"`
	PreservesMade          int `xml:"preservesMade"`
	PiecesOfTrashRecycled  int `xml:"piecesOfTrashRecycled"`
	MysticStonesCrushed    int `xml:"mysticStonesCrushed"`
	DaysPlayed             int `xml:"daysPlayed"`
	WeedsEliminated        int `xml:"weedsEliminated"`
	SticksChopped          int `xml:"sticksChopped"`
	NotesFound             int `xml:"notesFound"`
	QuestsCompleted        int `xml:"questsCompleted"`
	StarLevelCropsShipped  int `xml:"starLevelCropsShipped"`
	CropsShipped           int `xml:"cropsShipped"`
	ItemsForaged           int `xml:"itemsForaged"`
	SlimesKilled           int `xml:"slimesKilled"`
	GeodesCracked          int `xml:"geodesCracked"`
	GoodFriends            int `xml:"goodFriends"`
}

func (s *PlayerStats) Counts() map[string]int {
	counts := map[string]int{}
	counts["Seeds sown"] = s.SeedsSown
	counts["Items shipped"] = s.ItemsShipped
	counts["Items cooked"] = s.ItemsCooked
	counts["Items crafted"] = s.ItemsCrafted
	counts["Chicken eggs layed"] = s.ChickenEggsLayed
	counts["Duck eggs layed"] = s.DuckEggsLayed
	counts["Cow milk produced"] = s.CowMilkProduced
	counts["Goat milk produced"] = s.GoatMilkProduced
	counts["Rabbit wool produced"] = s.RabbitWoolProduced
	counts["Sheep wool produced"] = s.SheepWoolProduced
	counts["Cheese made"] = s.CheeseMade
	counts["Goat cheese made"] = s.GoatCheeseMade
	counts["Truffles found"] = s.TrufflesFound
	counts["Stone gathered"] = s.StoneGathered
	counts["Rocks crushed"] = s.RocksCrushed
	counts["Dirt hoed"] = s.DirtHoed
	counts["Gifts given"] = s.GiftsGiven
	counts["Times unconscious"] = s.TimesUnconscious
	if s.TimesUnconscious == 0 {
		counts["Times unconscious"] = s.TimesUnconscious2
	}
	counts["Average bedtime"] = s.AverageBedtime
	counts["Times fished"] = s.TimesFished
	counts["Fish caught"] = s.FishCaught
	counts["Boulders cracked"] = s.BouldersCracked
	if s.BouldersCracked == 0 {
		counts["Boulders cracked"] = s.BouldersCracked2
	}
	counts["Stumps chopped"] = s.StumpsChopped
	counts["Steps taken"] = s.StepsTaken
	counts["Total monsters killed"] = s.TotalMonstersKilled
	if s.TotalMonstersKilled == 0 {
		counts["Total monsters killed"] = s.TotalMonstersKilled2
	}
	counts["Diamonds found"] = s.DiamondsFound
	counts["Prismatic shards found"] = s.PrismaticShardsFound
	counts["Other precious gems found"] = s.OtherPreciousGemsFound
	counts["Cave carrots found"] = s.CaveCarrotsFound
	counts["Copper found"] = s.CopperFound
	counts["Iron found"] = s.IronFound
	counts["Coal found"] = s.CoalFound
	counts["Coins found"] = s.CoinsFound
	counts["Gold found"] = s.GoldFound
	counts["Iridium found"] = s.IridiumFound
	counts["Bars smelted"] = s.BarsSmelted
	counts["Beverages made"] = s.BeveragesMade
	counts["Preserves made"] = s.PreservesMade
	counts["Pieces of trash recycled"] = s.PiecesOfTrashRecycled
	counts["Mystic stones crushed"] = s.MysticStonesCrushed
	counts["Days played"] = s.DaysPlayed
	counts["Weeds eliminated"] = s.WeedsEliminated
	counts["Sticks chopped"] = s.SticksChopped
	counts["Notes found"] = s.NotesFound
	counts["Quests completed"] = s.QuestsCompleted
	counts["Star-level crops shipped"] = s.StarLevelCropsShipped
	counts["Crops shipped"] = s.CropsShipped
	counts["Items foraged"] = s.ItemsForaged
	counts["Slimes killed"] = s.SlimesKilled
	counts["Geodes cracked"] = s.GeodesCracked
	counts["Good friends"] = s.GoodFriends
	return counts
}

func (s *PlayerStats) String() string {
	counts := s.Counts()
	pairs := []string{}
	for stat, value := range counts {
		pairs = append(pairs, fmt.Sprintf("%s: %d", stat, value))
	}
	return strings.Join(pairs, ", ")
}
