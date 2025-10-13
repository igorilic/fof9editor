// ABOUTME: This file defines the Team data structure for FOF9 custom leagues
// ABOUTME: It includes team identity, stadium info, colors, and future stadium plans
package models

import "image/color"

// Team roof type constants
const (
	RoofOutdoor     = 0
	RoofDome        = 1
	RoofRetractable = 2
)

// Team turf type constants
const (
	TurfGrass      = 0
	TurfArtificial = 1
	TurfHybrid     = 2
)

// Team represents a football team with all attributes
type Team struct {
	// Team Identity
	Year         int    `csv:"YEAR"`
	TeamID       int    `csv:"TEAMID"`
	TeamName     string `csv:"TEAMNAME"`
	NickName     string `csv:"NICKNAME"`
	Abbreviation string `csv:"ABBREVIATION"`

	// League Structure
	Conference int `csv:"CONFERENCE"`
	Division   int `csv:"DIVISION"`

	// Location
	City int `csv:"CITY"` // References cities.csv

	// Team Colors (RGB values 0-255)
	PrimaryRed      int `csv:"PRIMARYRED"`
	PrimaryGreen    int `csv:"PRIMARYGREEN"`
	PrimaryBlue     int `csv:"PRIMARYBLUE"`
	SecondaryRed    int `csv:"SECONDARYRED"`
	SecondaryGreen  int `csv:"SECONDARYGREEN"`
	SecondaryBlue   int `csv:"SECONDARYBLUE"`

	// Stadium Info
	Roof      int `csv:"ROOF"`      // 0=outdoor, 1=dome, 2=retractable
	Turf      int `csv:"TURF"`      // 0=grass, 1=artificial, 2=hybrid
	Built     int `csv:"BUILT"`     // Year built
	Capacity  int `csv:"CAPACITY"`  // Stadium capacity
	Luxury    int `csv:"LUXURY"`    // Luxury boxes
	Condition int `csv:"CONDITION"` // Stadium condition (1-10)

	// Financial Data
	Attendance int `csv:"ATTENDANCE"` // Average attendance
	Support    int `csv:"SUPPORT"`    // Fan support level

	// Future Stadium Plans
	Plan              int    `csv:"PLAN"`              // 0=no plans, 1=plans active
	Completed         int    `csv:"COMPLETED"`         // Year future stadium completed
	Future            int    `csv:"FUTURE"`            // Future stadium active
	FutureName        string `csv:"FUTURENAME"`        // Future stadium name
	FutureAbbr        string `csv:"FUTUREABBR"`        // Future abbreviation
	FutureRoof        int    `csv:"FUTUREROOF"`        // Future roof type
	FutureTurf        int    `csv:"FUTURETURF"`        // Future turf type
	FutureCap         int    `csv:"FUTURECAP"`         // Future capacity
	FutureLuxury      int    `csv:"FUTURELUXURY"`      // Future luxury boxes
	TeamContribution  int    `csv:"TEAMCONTRIBUTION"`  // Team's financial contribution
}

// GetDisplayName returns the team's full name
func (t *Team) GetDisplayName() string {
	return t.TeamName + " " + t.NickName
}

// GetPrimaryColor returns the primary color as color.Color
func (t *Team) GetPrimaryColor() color.Color {
	return color.RGBA{
		R: uint8(t.PrimaryRed),
		G: uint8(t.PrimaryGreen),
		B: uint8(t.PrimaryBlue),
		A: 255,
	}
}

// GetSecondaryColor returns the secondary color as color.Color
func (t *Team) GetSecondaryColor() color.Color {
	return color.RGBA{
		R: uint8(t.SecondaryRed),
		G: uint8(t.SecondaryGreen),
		B: uint8(t.SecondaryBlue),
		A: 255,
	}
}
