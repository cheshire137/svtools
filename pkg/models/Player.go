package models

import "encoding/xml"

type Player struct {
	XMLName                     xml.Name `xml:"player"`
	Name                        string
	FarmName                    string
	FavoriteThing               string
	HorseName                   string
	CatPerson                   bool
	IsEmoting                   bool `xml:"isEmoting"`
	IsEmoting2                  bool `xml:"IsEmoting"`
	CurrentEmote                int
	IsCharging                  bool
	WillDestroyObjectsUnderfoot bool
	IsGlowing                   bool
	ColoredBorder               bool
	Flip                        bool
	DrawOnTop                   bool
	FaceTowardFarmer            bool
	FaceAwayFromFarmer          bool
	IgnoreMovementAnimation     bool
	TimeBeforeAIMovementAgain   int
	GlowingTransparency         float32
	GlowRate                    float32
	Position                    Position
	Speed                       int
	FacingDirection             int
	Scale                       int
	Professions                 []int
	ExperiencePoints            []int
	Items                       []Item
	DialogueQuestionsAnswered   []int
	CookingRecipes              []Recipe
	CraftingRecipes             []Recipe
	EventsSeen                  []int
	SecretNotesSeen             []int
	SongsHeard                  []string
	Achievements                []int
	SpecialItems                []int
	SpecialBigCraftables        []int
	MailReceived                []string
	Stats                       PlayerStats
	SlotCanHost                 bool
	MostRecentBed               Position
	Shirt                       int
	Hair                        int
	Skin                        int
	Shoes                       int
	Accessory                   int
	FacialHair                  int
	HairstyleColor              Color
	PantsColor                  Color
	EyeColor                    Color `xml:"newEyeColor"`
	DivorceTonight              bool
	WoodPieces                  int
	StonePieces                 int
	CopperPieces                int
	IronPieces                  int
	CoalPieces                  int
	GoldPieces                  int
	IridiumPieces               int
	QuartzPieces                int
	CaveChoice                  int
	Feed                        int
	FarmingLevel                int
	MiningLevel                 int
	CombatLevel                 int
	ForagingLevel               int
	FishingLevel                int
	LuckLevel                   int
	NewSkillPointsToSpend       int
	AddedFarmingLevel           int
	AddedMiningLevel            int
	AddedCombatLevel            int
	AddedForagingLevel          int
	AddedFishingLevel           int
	AddedLuckLevel              int
	ID                          int64 `xml:"UniqueMultiplayerID"`
}
