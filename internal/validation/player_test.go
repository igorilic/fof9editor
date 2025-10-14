// ABOUTME: Tests for player validation rules
// ABOUTME: Verifies player field validation according to game constraints

package validation

import (
	"testing"

	"github.com/igorilic/fof9editor/internal/models"
)

func TestValidatePlayer_ValidPlayer(t *testing.T) {
	player := &models.Player{
		FirstName:         "John",
		LastName:          "Doe",
		Team:              15,
		PositionKey:       10,
		Uniform:           12,
		OverallRating:     85,
		Height:            72,
		Weight:            210,
		HandSize:          9,
		ArmLength:         32,
		Experience:        5,
		College:           "State University",
		YearEntry:        2018,
		RoundDrafted:     3,
		SelectionDrafted: 85,
	}

	result := ValidatePlayer(player)

	if !result.Valid {
		t.Errorf("Expected valid player, got errors: %v", result.Errors)
	}
}

func TestValidatePlayer_EmptyName(t *testing.T) {
	player := &models.Player{
		FirstName: "",
		LastName:  "Doe",
	}

	result := ValidatePlayer(player)

	if result.Valid {
		t.Error("Expected validation to fail for empty first name")
	}

	if !result.HasError("FirstName") {
		t.Error("Expected error for FirstName field")
	}
}

func TestValidatePlayer_InvalidTeam(t *testing.T) {
	player := &models.Player{
		FirstName: "John",
		LastName:  "Doe",
		Team:      50, // Invalid: > 31
	}

	result := ValidatePlayer(player)

	if result.Valid {
		t.Error("Expected validation to fail for invalid team")
	}

	if !result.HasError("Team") {
		t.Error("Expected error for Team field")
	}
}

func TestValidatePlayer_InvalidUniform(t *testing.T) {
	player := &models.Player{
		FirstName: "John",
		LastName:  "Doe",
		Uniform:   100, // Invalid: > 99
	}

	result := ValidatePlayer(player)

	if result.Valid {
		t.Error("Expected validation to fail for invalid uniform number")
	}

	if !result.HasError("Uniform") {
		t.Error("Expected error for Uniform field")
	}
}

func TestValidatePlayer_InvalidHeight(t *testing.T) {
	player := &models.Player{
		FirstName: "John",
		LastName:  "Doe",
		Height:    50, // Invalid: < 60 inches
	}

	result := ValidatePlayer(player)

	if result.Valid {
		t.Error("Expected validation to fail for invalid height")
	}

	if !result.HasError("Height") {
		t.Error("Expected error for Height field")
	}
}

func TestValidatePlayer_InvalidWeight(t *testing.T) {
	player := &models.Player{
		FirstName: "John",
		LastName:  "Doe",
		Weight:    100, // Invalid: < 150 lbs
	}

	result := ValidatePlayer(player)

	if result.Valid {
		t.Error("Expected validation to fail for invalid weight")
	}

	if !result.HasError("Weight") {
		t.Error("Expected error for Weight field")
	}
}

func TestValidatePlayer_InvalidRating(t *testing.T) {
	player := &models.Player{
		FirstName:     "John",
		LastName:      "Doe",
		OverallRating: 100, // Invalid: > 99
	}

	result := ValidatePlayer(player)

	if result.Valid {
		t.Error("Expected validation to fail for invalid overall rating")
	}

	if !result.HasError("OverallRating") {
		t.Error("Expected error for OverallRating field")
	}
}

func TestValidatePlayerField_FirstName(t *testing.T) {
	// Valid name
	result := ValidatePlayerField("FirstName", "John")
	if !result.Valid {
		t.Errorf("Expected valid first name, got errors: %v", result.Errors)
	}

	// Empty name
	result = ValidatePlayerField("FirstName", "")
	if result.Valid {
		t.Error("Expected validation to fail for empty first name")
	}

	// Name too long
	longName := string(make([]byte, 51))
	result = ValidatePlayerField("FirstName", longName)
	if result.Valid {
		t.Error("Expected validation to fail for name too long")
	}
}

func TestValidatePlayerField_Uniform(t *testing.T) {
	// Valid uniform
	result := ValidatePlayerField("Uniform", 12)
	if !result.Valid {
		t.Errorf("Expected valid uniform, got errors: %v", result.Errors)
	}

	// Invalid uniform (too high)
	result = ValidatePlayerField("Uniform", 100)
	if result.Valid {
		t.Error("Expected validation to fail for uniform > 99")
	}

	// Invalid uniform (negative)
	result = ValidatePlayerField("Uniform", -1)
	if result.Valid {
		t.Error("Expected validation to fail for negative uniform")
	}
}

func TestValidatePlayerField_Height(t *testing.T) {
	// Valid height
	result := ValidatePlayerField("Height", 72)
	if !result.Valid {
		t.Errorf("Expected valid height, got errors: %v", result.Errors)
	}

	// Too short
	result = ValidatePlayerField("Height", 50)
	if result.Valid {
		t.Error("Expected validation to fail for height < 60")
	}

	// Too tall
	result = ValidatePlayerField("Height", 100)
	if result.Valid {
		t.Error("Expected validation to fail for height > 90")
	}
}

func TestValidatePlayerField_Experience(t *testing.T) {
	// Valid experience
	result := ValidatePlayerField("Experience", 5)
	if !result.Valid {
		t.Errorf("Expected valid experience, got errors: %v", result.Errors)
	}

	// Too much experience
	result = ValidatePlayerField("Experience", 30)
	if result.Valid {
		t.Error("Expected validation to fail for experience > 25")
	}

	// Negative experience
	result = ValidatePlayerField("Experience", -1)
	if result.Valid {
		t.Error("Expected validation to fail for negative experience")
	}
}
