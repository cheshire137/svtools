package models

import "encoding/xml"

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
	AverageBedtime         int      `xml:"averageBedtime"`
	TimesFished            int      `xml:"timesFished"`
	FishCaught             int      `xml:"fishCaught"`
	BouldersCracked        int      `xml:"bouldersCracked"`
	StumpsChopped          int      `xml:"stumpsChopped"`
	StepsTaken             int      `xml:"stepsTaken"`
	TotalMonstersKilled    int      `xml:"monstersKilled"`
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
