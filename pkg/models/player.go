package models

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	Professions                 Professions `xml:"professions"`
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
	Inventory                   PlayerItems `xml:"items"`
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
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s, %s Farm", p.Name, p.FarmName))
	sb.WriteString("\n\nSkills:\n  ")
	sb.WriteString(fmt.Sprintf("Farming: %d / Mining: %d / Combat: %d / Foraging: %d / Fishing: %d",
		p.FarmingLevel, p.MiningLevel, p.CombatLevel, p.ForagingLevel, p.FishingLevel))
	sb.WriteString("\n\nInventory:\n  ")
	sb.WriteString(p.Inventory.String())
	sb.WriteString("\n\nAppearance:\n")
	sb.WriteString(fmt.Sprintf("  hair #%d %s, skin #%d, facial hair #%d, eyes %s", p.Hair,
		p.HairstyleColor.String(), p.Skin, p.FacialHair, p.EyeColor.String()))
	sb.WriteString("\n\nClothing:\n")
	sb.WriteString(fmt.Sprintf("  shirt #%d, pants %s, shoes #%d, accessory #%d", p.Shirt,
		p.PantsColor.String(), p.Shoes, p.Accessory))
	sb.WriteString("\n\nStats:\n  ")
	sb.WriteString(p.Stats.String())
	sb.WriteString("\n\nProfessions:\n  ")
	sb.WriteString(p.Professions.String())
	return sb.String()
}
