package models

import (
	"encoding/xml"
	"fmt"
)

type Player struct {
	XMLName                     xml.Name    `xml:"player"`
	Name                        string      `xml:"name"`
	FarmName                    string      `xml:"farmName"`
	IsEmoting                   bool        `xml:"isEmoting"`
	IsEmoting2                  bool        `xml:"IsEmoting"`
	FavoriteThing               string      `xml:"favoriteThing"`
	HorseName                   string      `xml:"horseName"`
	CatPerson                   bool        `xml:"catPerson"`
	CurrentEmote                int         `xml:"currentEmote"`
	IsCharging                  bool        `xml:"isCharging"`
	WillDestroyObjectsUnderfoot bool        `xml:"willDestroyObjectsUnderfoot"`
	IsGlowing                   bool        `xml:"isGlowing"`
	ColoredBorder               bool        `xml:"coloredBorder"`
	Flip                        bool        `xml:"flip"`
	DrawOnTop                   bool        `xml:"drawOnTop"`
	FaceTowardFarmer            bool        `xml:"faceTowardFarmer"`
	FaceAwayFromFarmer          bool        `xml:"faceAwayFromFarmer"`
	IgnoreMovementAnimation     bool        `xml:"ignoreMovementAnimation"`
	TimeBeforeAIMovementAgain   int         `xml:"timeBeforeAIMovementAgain"`
	GlowingTransparency         float32     `xml:"glowingTransparency"`
	GlowRate                    float32     `xml:"glowRate"`
	Speed                       int         `xml:"speed"`
	FacingDirection             int         `xml:"facingDirection"`
	Scale                       int         `xml:"scale"`
	Professions                 []int       `xml:"professions"`
	ExperiencePoints            []int       `xml:"experiencePoints"`
	DialogueQuestionsAnswered   []int       `xml:"dialogueQuestionsAnswered"`
	EventsSeen                  []int       `xml:"eventsSeen"`
	SecretNotesSeen             []int       `xml:"secretNotesSeen"`
	SongsHeard                  []string    `xml:"songsHeard"`
	Achievements                []int       `xml:"achievements"`
	SpecialItems                []int       `xml:"specialItems"`
	SpecialBigCraftables        []int       `xml:"specialBigCraftables"`
	MailReceived                []string    `xml:"mailReceived"`
	SlotCanHost                 bool        `xml:"slotCanHost"`
	Shirt                       int         `xml:"shirt"`
	Hair                        int         `xml:"hair"`
	Skin                        int         `xml:"skin"`
	Shoes                       int         `xml:"shoes"`
	Accessory                   int         `xml:"accessory"`
	FacialHair                  int         `xml:"facialHair"`
	DivorceTonight              bool        `xml:"divorceTonight"`
	WoodPieces                  int         `xml:"woodPieces"`
	StonePieces                 int         `xml:"stonePieces"`
	CopperPieces                int         `xml:"copperPieces"`
	IronPieces                  int         `xml:"ironPieces"`
	CoalPieces                  int         `xml:"coalPieces"`
	GoldPieces                  int         `xml:"goldPieces"`
	IridiumPieces               int         `xml:"iridiumPieces"`
	QuartzPieces                int         `xml:"quartzPieces"`
	CaveChoice                  int         `xml:"caveChoice"`
	Feed                        int         `xml:"feed"`
	FarmingLevel                int         `xml:"farmingLevel"`
	MiningLevel                 int         `xml:"miningLevel"`
	CombatLevel                 int         `xml:"combatLevel"`
	ForagingLevel               int         `xml:"foragingLevel"`
	FishingLevel                int         `xml:"fishingLevel"`
	LuckLevel                   int         `xml:"luckLevel"`
	NewSkillPointsToSpend       int         `xml:"newSkillPointsToSpend"`
	AddedFarmingLevel           int         `xml:"addedFarmingLevel"`
	AddedMiningLevel            int         `xml:"addedMiningLevel"`
	AddedCombatLevel            int         `xml:"addedCombatLevel"`
	AddedForagingLevel          int         `xml:"addedForagingLevel"`
	AddedFishingLevel           int         `xml:"addedFishingLevel"`
	AddedLuckLevel              int         `xml:"addedLuckLevel"`
	Items                       PlayerItems `xml:"items"`
	Position                    Position
	CookingRecipes              []Recipe
	CraftingRecipes             []Recipe
	Stats                       PlayerStats
	MostRecentBed               Position
	HairstyleColor              Color
	PantsColor                  Color
	EyeColor                    Color `xml:"newEyeColor"`
	ID                          int64 `xml:"UniqueMultiplayerID"`
}

func (p *Player) String() string {
	return fmt.Sprintf("%s - %s Farm", p.Name, p.FarmName)
}
