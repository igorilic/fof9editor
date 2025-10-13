package models

import (
	"testing"
)

func TestCoachGetDisplayName(t *testing.T) {
	coach := &Coach{
		FirstName: "John",
		LastName:  "Smith",
	}

	expected := "John Smith"
	actual := coach.GetDisplayName()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestCoachGetPositionName(t *testing.T) {
	tests := []struct {
		position int
		expected string
	}{
		{PositionHeadCoach, "Head Coach"},
		{PositionOffensiveCoordinator, "Offensive Coordinator"},
		{PositionDefensiveCoordinator, "Defensive Coordinator"},
		{PositionSpecialTeamsCoordinator, "Special Teams Coordinator"},
		{PositionStrengthConditioning, "Strength & Conditioning"},
		{99, "Unknown"},
	}

	for _, tt := range tests {
		coach := &Coach{Position: tt.position}
		actual := coach.GetPositionName()
		if actual != tt.expected {
			t.Errorf("Position %d: expected %s, got %s", tt.position, tt.expected, actual)
		}
	}
}

func TestCoachStructFields(t *testing.T) {
	coach := &Coach{
		FirstName:      "Mike",
		LastName:       "Johnson",
		BirthMonth:     3,
		BirthDay:       15,
		BirthYear:      1975,
		BirthCity:      "Chicago_IL",
		BirthCityID:    100,
		College:        "Notre Dame",
		CollegeID:      50,
		Team:           5,
		Position:       PositionHeadCoach,
		PositionGroup:  10,
		OffensiveStyle: 3,
		DefensiveStyle: 2,
		PayScale:       120, // $1.2M
	}

	if coach.FirstName != "Mike" {
		t.Errorf("Expected FirstName Mike, got %s", coach.FirstName)
	}

	if coach.Position != PositionHeadCoach {
		t.Errorf("Expected Position %d, got %d", PositionHeadCoach, coach.Position)
	}

	if coach.PayScale != 120 {
		t.Errorf("Expected PayScale 120, got %d", coach.PayScale)
	}

	if coach.GetPositionName() != "Head Coach" {
		t.Errorf("Expected position name 'Head Coach', got '%s'", coach.GetPositionName())
	}
}

func TestCoachPositionConstants(t *testing.T) {
	// Verify position constants have expected values
	if PositionHeadCoach != 0 {
		t.Errorf("Expected PositionHeadCoach = 0, got %d", PositionHeadCoach)
	}

	if PositionOffensiveCoordinator != 1 {
		t.Errorf("Expected PositionOffensiveCoordinator = 1, got %d", PositionOffensiveCoordinator)
	}

	if PositionDefensiveCoordinator != 2 {
		t.Errorf("Expected PositionDefensiveCoordinator = 2, got %d", PositionDefensiveCoordinator)
	}
}
