// ABOUTME: This file defines the Coach data structure for FOF9 custom leagues
// ABOUTME: It includes coach attributes, position types, and coaching styles
package models

// Coach position constants
const (
	PositionHeadCoach                 = 0
	PositionOffensiveCoordinator      = 1
	PositionDefensiveCoordinator      = 2
	PositionSpecialTeamsCoordinator   = 3
	PositionStrengthConditioning      = 4
)

// Coach represents a football coach with all attributes
type Coach struct {
	// Basic Info
	LastName  string `csv:"LASTNAME"`
	FirstName string `csv:"FIRSTNAME"`

	// Birth Info
	BirthMonth int `csv:"BIRTHMONTH"`
	BirthDay   int `csv:"BIRTHDAY"`
	BirthYear  int `csv:"BIRTHYEAR"`

	// Birth Location (text field not used by game, but helpful for editing)
	BirthCity   string `csv:"BIRTHCITY"`
	BirthCityID int    `csv:"CITYID"`

	// College (text field not used by game, but helpful for editing)
	College   string `csv:"COLLEGE"`
	CollegeID int    `csv:"COLLEGEID"`

	// Position and Team
	Team          int `csv:"TEAM"`
	Position      int `csv:"POSITION"`      // 0=Head Coach, 1=OC, 2=DC, 3=ST, 4=S&C
	PositionGroup int `csv:"POSITIONGROUP"` // Position-specific role

	// Coaching Styles
	OffensiveStyle int `csv:"OFFENSIVESTYLE"` // 0-6
	DefensiveStyle int `csv:"DEFENSIVESTYLE"` // 0-4

	// Compensation
	PayScale int `csv:"PAYSCALE"` // In units of $10,000
}

// GetDisplayName returns the coach's full name
func (c *Coach) GetDisplayName() string {
	return c.FirstName + " " + c.LastName
}

// GetPositionName returns the human-readable position name
func (c *Coach) GetPositionName() string {
	switch c.Position {
	case PositionHeadCoach:
		return "Head Coach"
	case PositionOffensiveCoordinator:
		return "Offensive Coordinator"
	case PositionDefensiveCoordinator:
		return "Defensive Coordinator"
	case PositionSpecialTeamsCoordinator:
		return "Special Teams Coordinator"
	case PositionStrengthConditioning:
		return "Strength & Conditioning"
	default:
		return "Unknown"
	}
}
