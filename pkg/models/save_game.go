package models

import "encoding/xml"

type SaveGame struct {
	XMLName                xml.Name `xml:"SaveGame"`
	Player                 Player
	Locations              []Location
	CurrentSeason          string `xml:"currentSeason"`
	SamBandName            string `xml:"samBandName"`
	ElliottBookName        string `xml:"elliottBookName"`
	DayOfMonth             int    `xml:"dayOfMonth"`
	Year                   int    `xml:"year"`
	FarmerWallpaper        int    `xml:"farmerWallpaper"`
	FarmerFloor            string `xml:"FarmerFloor"`
	CurrentWallpaper       int    `xml:"currentWallpaper"`
	CurrentFloor           int    `xml:"currentFloor"`
	ID                     int64  `xml:"uniqueIDForThisGame"`
	Has13Update            bool   `xml:"hasApplied1_3_UpdateChanges"`
	IncubatingEgg          Position
	CurrentSongIndex       int     `xml:"currentSongIndex"`
	ChanceToRainTomorrow   float32 `xml:"chanceToRainTomorrow"`
	DailyLuck              float32 `xml:"dailyLuck"`
	WeddingToday           bool    `xml:"weddingToday"`
	IsRaining              bool    `xml:"isRaining"`
	IsDebrisWeather        bool    `xml:"isDebrisWeather"`
	ShippingTax            bool    `xml:"shippingTax"`
	BloomDay               bool    `xml:"bloomDay"`
	IsLightning            bool    `xml:"isLightning"`
	IsSnowing              bool    `xml:"isSnowing"`
	ShouldSpawnMonsters    bool    `xml:"shouldSpawnMonsters"`
	MusicVolume            float32 `xml:"musicVolume"`
	SoundVolume            float32 `xml:"soundVolume"`
	CropsOfTheWeek         []int   `xml:"cropsOfTheWeek"`
	LatestID               int64   `xml:"latestID"`
	MinecartHighScore      int     `xml:"minecartHighScore"`
	WeatherForTomorrow     int     `xml:"weatherForTomorrow"`
	DishOfTheDay           DishOfTheDay
	Options                Options
	LowestMineLevelReached int `xml:"mine_lowestLevelReached"`
	FarmType               int `xml:"whichFarm"`
}

func (s *SaveGame) String() string {
	return s.Player.String()
}
