package models

import "encoding/xml"

type SaveGame struct {
	XMLName                xml.Name `xml:"SaveGame"`
	Player                 Player
	Locations              []Location
	CurrentSeason          string
	SamBandName            string
	ElliottBookName        string
	DayOfMonth             int
	Year                   int
	FarmerWallpaper        int
	FarmerFloor            string `xml:"FarmerFloor"`
	CurrentWallpaper       int
	CurrentFloor           int
	CurrentSongIndex       int
	IncubatingEgg          Position
	ChanceToRainTomorrow   float32
	DailyLuck              float32
	ID                     int64 `xml:"uniqueIDForThisGame"`
	WeddingToday           bool
	IsRaining              bool
	IsDebrisWeather        bool
	ShippingTax            bool
	BloomDay               bool
	IsLightning            bool
	IsSnowing              bool
	ShouldSpawnMonsters    bool
	Has13Update            bool `xml:"hasApplied1_3_UpdateChanges"`
	MusicVolume            float32
	SoundVolume            float32
	CropsOfTheWeek         []int
	DishOfTheDay           DishOfTheDay
	LatestID               int64
	Options                Options
	LowestMineLevelReached int `xml:"mine_lowestLevelReached"`
	MinecartHighScore      int
	WeatherForTomorrow     int
	FarmType               int `xml:"whichFarm"`
}
