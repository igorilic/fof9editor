// ABOUTME: This file defines the Player data structure for FOF9 custom leagues
// ABOUTME: It includes all player attributes, physical data, career info, and skill ratings
package models

// Player represents a football player with all attributes
type Player struct {
	// Basic Info
	PlayerID  int    `csv:"PLAYERID"`
	LastName  string `csv:"LASTNAME"`
	FirstName string `csv:"FIRSTNAME"`

	// Team and Position
	Team        int `csv:"TEAM"`
	PositionKey int `csv:"POSITION_KEY"`
	Uniform     int `csv:"UNIFORM"`

	// Physical Attributes
	Height    int `csv:"HEIGHT"`
	Weight    int `csv:"WEIGHT"`
	HandSize  int `csv:"HANDSIZE"`
	ArmLength int `csv:"ARMLENGTH"`

	// Birth Info
	BirthMonth int `csv:"BIRTHMONTH"`
	BirthDay   int `csv:"BIRTHDAY"`
	BirthYear  int `csv:"BIRTHYEAR"`

	// Birth Location (text fields not used by game, but helpful for editing)
	BirthCity   string `csv:"BIRTHCITY"`
	BirthCityID int    `csv:"CITYID"`

	// College (text field not used by game, but helpful for editing)
	College   string `csv:"COLLEGE"`
	CollegeID int    `csv:"COLLEGEID"`

	// Draft History
	YearEntry         int `csv:"YEARENTRY"`
	RoundDrafted      int `csv:"ROUNDDRAFTED"`
	SelectionDrafted  int `csv:"SELECTIONDRAFTED"`
	Supplemental      int `csv:"SUPPLEMENTAL"`
	OriginalTeam      int `csv:"ORIGINALTEAM"`

	// Career Stats
	Experience        int `csv:"EXPERIENCE"`
	YearSigned        int `csv:"YEARSIGNED"`
	PlayPercentage    int `csv:"PLAYPERCENTAGE"`
	HallOfFamePoints  int `csv:"HALLOFFAMEPOINTS"`

	// Contract
	SalaryYears int `csv:"SALARYYEARS"`
	SalaryYear1 int `csv:"SALARYYEAR1"`
	BonusYear1  int `csv:"BONUSYEAR1"`
	SalaryYear2 int `csv:"SALARYYEAR2"`
	BonusYear2  int `csv:"BONUSYEAR2"`
	SalaryYear3 int `csv:"SALARYYEAR3"`
	BonusYear3  int `csv:"BONUSYEAR3"`
	SalaryYear4 int `csv:"SALARYYEAR4"`
	BonusYear4  int `csv:"BONUSYEAR4"`
	SalaryYear5 int `csv:"SALARYYEAR5"`
	BonusYear5  int `csv:"BONUSYEAR5"`

	// Overall Rating
	OverallRating int `csv:"OVERALLRATING"`

	// Skill Attributes (all default to -1 for auto-generate)
	SkillSpeed               int `csv:"SKILL_SPEED"`
	SkillPower               int `csv:"SKILL_POWER"`
	HoleRecognition          int `csv:"HOLE_RECOGNITION"`
	Elusiveness              int `csv:"ELUSIVENESS"`
	BlitzPickup              int `csv:"BLITZ_PICKUP"`
	CatchHands               int `csv:"CATCH_HANDS"`
	AdjustToBall             int `csv:"ADJUST_TO_BALL"`
	RouteRunning             int `csv:"ROUTE_RUNNING"`
	CatchInTraffic           int `csv:"CATCH_IN_TRAFFIC"`
	DefeatBlockers           int `csv:"DEFEAT_BLOCKERS"`
	SecureHandling           int `csv:"SECURE_HANDLING"`
	RunBlockTechnique        int `csv:"RUN_BLOCK_TECHNIQUE"`
	PassBlockTechnique       int `csv:"PASS_BLOCK_TECHNIQUE"`
	BlockingStrength         int `csv:"BLOCKING_STRENGTH"`
	SchemeAcquisition        int `csv:"SCHEME_ACQUISITION"`
	PuntDistance             int `csv:"PUNT_DISTANCE"`
	PuntHangTime             int `csv:"PUNT_HANG_TIME"`
	PuntDirectional          int `csv:"PUNT_DIRECTIONAL"`
	KickoffHangTime          int `csv:"KICKOFF_HANG_TIME"`
	FieldGoalAccuracy        int `csv:"FIELD_GOAL_ACCURACY"`
	FieldGoalDistance        int `csv:"FIELD_GOAL_DISTANCE"`
	RunDefense               int `csv:"RUN_DEFENSE"`
	PassRushTechnique        int `csv:"PASS_RUSH_TECHNIQUE"`
	PassRushStrength         int `csv:"PASS_RUSH_STRENGTH"`
	PassDefenseMan           int `csv:"PASS_DEFENSE_MAN"`
	PassDefensePhysical      int `csv:"PASS_DEFENSE_PHYSICAL"`
	PassDefenseZone          int `csv:"PASS_DEFENSE_ZONE"`
	PassDefenseHands         int `csv:"PASS_DEFENSE_HANDS"`
	DefensiveDiagnosis       int `csv:"DEFENSIVE_DIAGNOSIS"`
	SpecialTeams             int `csv:"SPECIAL_TEAMS"`
	PuntReturns              int `csv:"PUNT_RETURNS"`
	KickReturns              int `csv:"KICK_RETURNS"`
	LongSnapping             int `csv:"LONG_SNAPPING"`
	KickHolding              int `csv:"KICK_HOLDING"`
	Endurance                int `csv:"ENDURANCE"`

	// Base Year (determines draft class)
	BaseYear int `csv:"BASE_YEAR"`
}

// GetDisplayName returns the player's full name
func (p *Player) GetDisplayName() string {
	return p.FirstName + " " + p.LastName
}
