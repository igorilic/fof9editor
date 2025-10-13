package models

import (
	"testing"
)

func TestTeamGetDisplayName(t *testing.T) {
	team := &Team{
		TeamName: "Dallas",
		NickName: "Cowboys",
	}

	expected := "Dallas Cowboys"
	actual := team.GetDisplayName()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestTeamGetPrimaryColor(t *testing.T) {
	team := &Team{
		PrimaryRed:   255,
		PrimaryGreen: 0,
		PrimaryBlue:  0,
	}

	c := team.GetPrimaryColor()
	r, g, b, a := c.RGBA()

	// RGBA() returns uint32 values 0-65535, so we need to scale down
	if r>>8 != 255 || g>>8 != 0 || b>>8 != 0 || a>>8 != 255 {
		t.Errorf("Expected RGBA(255,0,0,255), got RGBA(%d,%d,%d,%d)", r>>8, g>>8, b>>8, a>>8)
	}
}

func TestTeamGetSecondaryColor(t *testing.T) {
	team := &Team{
		SecondaryRed:   0,
		SecondaryGreen: 0,
		SecondaryBlue:  255,
	}

	c := team.GetSecondaryColor()
	r, g, b, a := c.RGBA()

	// RGBA() returns uint32 values 0-65535, so we need to scale down
	if r>>8 != 0 || g>>8 != 0 || b>>8 != 255 || a>>8 != 255 {
		t.Errorf("Expected RGBA(0,0,255,255), got RGBA(%d,%d,%d,%d)", r>>8, g>>8, b>>8, a>>8)
	}
}

func TestTeamStructFields(t *testing.T) {
	team := &Team{
		Year:           2024,
		TeamID:         1,
		TeamName:       "Arizona",
		NickName:       "Cardinals",
		Abbreviation:   "ARI",
		Conference:     2,
		Division:       4,
		City:           1110,
		PrimaryRed:     151,
		PrimaryGreen:   35,
		PrimaryBlue:    63,
		SecondaryRed:   255,
		SecondaryGreen: 182,
		SecondaryBlue:  18,
		Roof:           RoofDome,
		Turf:           TurfGrass,
		Built:          1958,
		Capacity:       71706,
		Luxury:         68,
		Condition:      2,
		Attendance:     54,
		Support:        0,
		Plan:           1,
		Completed:      2006,
		Future:         993,
	}

	if team.TeamID != 1 {
		t.Errorf("Expected TeamID 1, got %d", team.TeamID)
	}

	if team.Abbreviation != "ARI" {
		t.Errorf("Expected Abbreviation ARI, got %s", team.Abbreviation)
	}

	if team.Capacity != 71706 {
		t.Errorf("Expected Capacity 71706, got %d", team.Capacity)
	}

	if team.Roof != RoofDome {
		t.Errorf("Expected Roof %d, got %d", RoofDome, team.Roof)
	}
}

func TestTeamConstants(t *testing.T) {
	// Verify roof constants
	if RoofOutdoor != 0 {
		t.Errorf("Expected RoofOutdoor = 0, got %d", RoofOutdoor)
	}
	if RoofDome != 1 {
		t.Errorf("Expected RoofDome = 1, got %d", RoofDome)
	}
	if RoofRetractable != 2 {
		t.Errorf("Expected RoofRetractable = 2, got %d", RoofRetractable)
	}

	// Verify turf constants
	if TurfGrass != 0 {
		t.Errorf("Expected TurfGrass = 0, got %d", TurfGrass)
	}
	if TurfArtificial != 1 {
		t.Errorf("Expected TurfArtificial = 1, got %d", TurfArtificial)
	}
	if TurfHybrid != 2 {
		t.Errorf("Expected TurfHybrid = 2, got %d", TurfHybrid)
	}
}
