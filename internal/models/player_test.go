package models

import (
	"testing"
)

func TestPlayerGetDisplayName(t *testing.T) {
	player := &Player{
		FirstName: "John",
		LastName:  "Doe",
	}

	expected := "John Doe"
	actual := player.GetDisplayName()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestPlayerStructFields(t *testing.T) {
	// Test that we can create a player with all fields
	player := &Player{
		PlayerID:     1000,
		FirstName:    "Test",
		LastName:     "Player",
		Team:         1,
		PositionKey:  2, // Running Back
		Uniform:      25,
		Height:       72,
		Weight:       215,
		HandSize:     93,
		ArmLength:    320,
		BirthMonth:   1,
		BirthDay:     15,
		BirthYear:    2002,
		BirthCity:    "New York_NY",
		BirthCityID:  1,
		College:      "Alabama",
		CollegeID:    1,
		YearEntry:    2024,
		RoundDrafted: 1,
		SelectionDrafted: 15,
		Supplemental: 0,
		OriginalTeam: 1,
		Experience:   0,
		YearSigned:   2024,
		PlayPercentage: 0,
		HallOfFamePoints: 0,
		SalaryYears:  4,
		SalaryYear1:  1000,
		BonusYear1:   500,
		OverallRating: 7,
		SkillSpeed:   -1,
		SkillPower:   -1,
		BaseYear:     2024,
	}

	// Verify key fields
	if player.PlayerID != 1000 {
		t.Errorf("Expected PlayerID 1000, got %d", player.PlayerID)
	}

	if player.PositionKey != 2 {
		t.Errorf("Expected PositionKey 2, got %d", player.PositionKey)
	}

	if player.OverallRating != 7 {
		t.Errorf("Expected OverallRating 7, got %d", player.OverallRating)
	}

	if player.GetDisplayName() != "Test Player" {
		t.Errorf("Expected name 'Test Player', got '%s'", player.GetDisplayName())
	}
}

func TestPlayerDefaultValues(t *testing.T) {
	// Test zero values for a new player
	player := &Player{}

	if player.PlayerID != 0 {
		t.Errorf("Expected default PlayerID 0, got %d", player.PlayerID)
	}

	if player.Team != 0 {
		t.Errorf("Expected default Team 0, got %d", player.Team)
	}

	if player.OverallRating != 0 {
		t.Errorf("Expected default OverallRating 0, got %d", player.OverallRating)
	}
}
