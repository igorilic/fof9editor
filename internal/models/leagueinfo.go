// ABOUTME: This file defines the LeagueInfo data structure for FOF9 custom leagues
// ABOUTME: It manages league configuration including schedule, salary cap, and salary minimums
package models

import "strings"

// LeagueInfo represents the league configuration settings
type LeagueInfo struct {
	ScheduleID string `csv:"SCHEDULEID"` // Format: "x_y_z" (teams_divisions_games)
	BaseYear   int    `csv:"BASE_YEAR"`  // League starting year (1900-2199)
	SalaryCap  int    `csv:"SALARYCAP"`  // In units of $100,000

	// Salary Minimums (in units of $10,000)
	Minimum   int `csv:"MINIMUM"`   // Rookie minimum
	Salary1   int `csv:"SALARY1"`   // 1 year experience
	Salary2   int `csv:"SALARY2"`   // 2 years experience
	Salary3   int `csv:"SALARY3"`   // 3 years experience
	Salary45  int `csv:"SALARY45"`  // 4-5 years experience
	Salary789 int `csv:"SALARY789"` // 7-9 years experience
	Salary10  int `csv:"SALARY10"`  // 10+ years experience
}

// NewDefaultLeagueInfo creates a LeagueInfo with default values
func NewDefaultLeagueInfo(baseYear int) *LeagueInfo {
	return &LeagueInfo{
		ScheduleID: "32_8_17", // 32 teams, 8 divisions, 17 games
		BaseYear:   baseYear,
		SalaryCap:  2000,  // $200M
		Minimum:    70,    // $700k
		Salary1:    85,    // $850k
		Salary2:    100,   // $1M
		Salary3:    115,   // $1.15M
		Salary45:   130,   // $1.3M
		Salary789:  150,   // $1.5M
		Salary10:   180,   // $1.8M
	}
}

// GetSalaryMinimum returns the appropriate salary minimum based on experience
func (l *LeagueInfo) GetSalaryMinimum(experience int) int {
	switch {
	case experience == 0:
		return l.Minimum
	case experience == 1:
		return l.Salary1
	case experience == 2:
		return l.Salary2
	case experience == 3:
		return l.Salary3
	case experience >= 4 && experience <= 5:
		return l.Salary45
	case experience >= 7 && experience <= 9:
		return l.Salary789
	case experience >= 10:
		return l.Salary10
	default:
		return l.Minimum
	}
}

// ValidateScheduleID checks if the ScheduleID has the correct format "x_y_z"
func (l *LeagueInfo) ValidateScheduleID() bool {
	parts := strings.Split(l.ScheduleID, "_")
	return len(parts) == 3
}
