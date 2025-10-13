// ABOUTME: Tests for coach list view component
// ABOUTME: Validates coach display and table functionality

package ui

import (
	"testing"

	"github.com/igorilic/fof9editor/internal/models"
)

func TestNewCoachList(t *testing.T) {
	cl := NewCoachList()

	if cl == nil {
		t.Fatal("NewCoachList returned nil")
	}

	if cl.container == nil {
		t.Fatal("CoachList container is nil")
	}

	if cl.table == nil {
		t.Fatal("CoachList table is nil")
	}

	if len(cl.headers) == 0 {
		t.Fatal("CoachList headers is empty")
	}
}

func TestCoachList_SetCoaches(t *testing.T) {
	cl := NewCoachList()

	coaches := []models.Coach{
		{FirstName: "Bill", LastName: "Belichick", Position: models.PositionHeadCoach, Team: 1, PayScale: 500},
		{FirstName: "Andy", LastName: "Reid", Position: models.PositionHeadCoach, Team: 2, PayScale: 450},
	}

	cl.SetCoaches(coaches)

	if len(cl.coaches) != 2 {
		t.Errorf("Expected 2 coaches, got %d", len(cl.coaches))
	}

	if cl.coaches[0].FirstName != "Bill" {
		t.Errorf("Expected first coach 'Bill', got '%s'", cl.coaches[0].FirstName)
	}
}

func TestCoachList_GetCoaches(t *testing.T) {
	cl := NewCoachList()

	coaches := []models.Coach{
		{FirstName: "Bill", LastName: "Belichick", Position: models.PositionHeadCoach, Team: 1, PayScale: 500},
	}

	cl.SetCoaches(coaches)

	retrieved := cl.GetCoaches()
	if len(retrieved) != 1 {
		t.Errorf("Expected 1 coach, got %d", len(retrieved))
	}

	if retrieved[0].LastName != "Belichick" {
		t.Errorf("Expected coach 'Belichick', got '%s'", retrieved[0].LastName)
	}
}

func TestCoachList_GetContainer(t *testing.T) {
	cl := NewCoachList()

	container := cl.GetContainer()
	if container == nil {
		t.Fatal("GetContainer returned nil")
	}

	if container != cl.container {
		t.Error("GetContainer did not return the correct container")
	}
}

func TestCoachList_Clear(t *testing.T) {
	cl := NewCoachList()

	coaches := []models.Coach{
		{FirstName: "Bill", LastName: "Belichick", Position: models.PositionHeadCoach, Team: 1, PayScale: 500},
		{FirstName: "Andy", LastName: "Reid", Position: models.PositionHeadCoach, Team: 2, PayScale: 450},
	}

	cl.SetCoaches(coaches)

	if len(cl.coaches) != 2 {
		t.Fatalf("Expected 2 coaches before clear, got %d", len(cl.coaches))
	}

	cl.Clear()

	if len(cl.coaches) != 0 {
		t.Errorf("Expected 0 coaches after clear, got %d", len(cl.coaches))
	}
}

func TestCoachList_EmptyList(t *testing.T) {
	cl := NewCoachList()

	// Should handle empty list without panic
	cl.SetCoaches([]models.Coach{})

	if len(cl.coaches) != 0 {
		t.Errorf("Expected 0 coaches, got %d", len(cl.coaches))
	}
}

func TestCoachList_Headers(t *testing.T) {
	cl := NewCoachList()

	expectedHeaders := []string{"First Name", "Last Name", "Position", "Team", "Pay Scale"}

	if len(cl.headers) != len(expectedHeaders) {
		t.Errorf("Expected %d headers, got %d", len(expectedHeaders), len(cl.headers))
	}

	for i, expected := range expectedHeaders {
		if cl.headers[i] != expected {
			t.Errorf("Header %d: expected '%s', got '%s'", i, expected, cl.headers[i])
		}
	}
}

func TestCoachList_GetPositionName(t *testing.T) {
	cl := NewCoachList()

	tests := []struct {
		position int
		expected string
	}{
		{models.PositionHeadCoach, "Head Coach"},
		{models.PositionOffensiveCoordinator, "Offensive Coordinator"},
		{models.PositionDefensiveCoordinator, "Defensive Coordinator"},
		{models.PositionSpecialTeamsCoordinator, "Special Teams Coordinator"},
		{models.PositionStrengthConditioning, "Strength & Conditioning"},
		{99, "Position 99"},
	}

	for _, tt := range tests {
		got := cl.getPositionName(tt.position)
		if got != tt.expected {
			t.Errorf("getPositionName(%d): expected '%s', got '%s'", tt.position, tt.expected, got)
		}
	}
}
