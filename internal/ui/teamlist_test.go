// ABOUTME: Tests for team list view component
// ABOUTME: Validates team display and table functionality

package ui

import (
	"testing"

	"github.com/igorilic/fof9editor/internal/models"
)

func TestNewTeamList(t *testing.T) {
	tl := NewTeamList()

	if tl == nil {
		t.Fatal("NewTeamList returned nil")
	}

	if tl.container == nil {
		t.Fatal("TeamList container is nil")
	}

	if tl.table == nil {
		t.Fatal("TeamList table is nil")
	}

	if len(tl.headers) == 0 {
		t.Fatal("TeamList headers is empty")
	}
}

func TestTeamList_SetTeams(t *testing.T) {
	tl := NewTeamList()

	teams := []models.Team{
		{TeamID: 1, TeamName: "New England", NickName: "Patriots", Abbreviation: "NE", Conference: 0, Division: 0},
		{TeamID: 2, TeamName: "Kansas City", NickName: "Chiefs", Abbreviation: "KC", Conference: 0, Division: 3},
	}

	tl.SetTeams(teams)

	if len(tl.teams) != 2 {
		t.Errorf("Expected 2 teams, got %d", len(tl.teams))
	}

	if tl.teams[0].TeamName != "New England" {
		t.Errorf("Expected first team 'New England', got '%s'", tl.teams[0].TeamName)
	}
}

func TestTeamList_GetTeams(t *testing.T) {
	tl := NewTeamList()

	teams := []models.Team{
		{TeamID: 1, TeamName: "New England", NickName: "Patriots", Abbreviation: "NE", Conference: 0, Division: 0},
	}

	tl.SetTeams(teams)

	retrieved := tl.GetTeams()
	if len(retrieved) != 1 {
		t.Errorf("Expected 1 team, got %d", len(retrieved))
	}

	if retrieved[0].TeamID != 1 {
		t.Errorf("Expected team ID 1, got %d", retrieved[0].TeamID)
	}
}

func TestTeamList_GetContainer(t *testing.T) {
	tl := NewTeamList()

	container := tl.GetContainer()
	if container == nil {
		t.Fatal("GetContainer returned nil")
	}

	if container != tl.container {
		t.Error("GetContainer did not return the correct container")
	}
}

func TestTeamList_Clear(t *testing.T) {
	tl := NewTeamList()

	teams := []models.Team{
		{TeamID: 1, TeamName: "New England", NickName: "Patriots", Abbreviation: "NE", Conference: 0, Division: 0},
		{TeamID: 2, TeamName: "Kansas City", NickName: "Chiefs", Abbreviation: "KC", Conference: 0, Division: 3},
	}

	tl.SetTeams(teams)

	if len(tl.teams) != 2 {
		t.Fatalf("Expected 2 teams before clear, got %d", len(tl.teams))
	}

	tl.Clear()

	if len(tl.teams) != 0 {
		t.Errorf("Expected 0 teams after clear, got %d", len(tl.teams))
	}
}

func TestTeamList_EmptyList(t *testing.T) {
	tl := NewTeamList()

	// Should handle empty list without panic
	tl.SetTeams([]models.Team{})

	if len(tl.teams) != 0 {
		t.Errorf("Expected 0 teams, got %d", len(tl.teams))
	}
}

func TestTeamList_Headers(t *testing.T) {
	tl := NewTeamList()

	expectedHeaders := []string{"ID", "Team Name", "Nickname", "Abbreviation", "Conference", "Division"}

	if len(tl.headers) != len(expectedHeaders) {
		t.Errorf("Expected %d headers, got %d", len(expectedHeaders), len(tl.headers))
	}

	for i, expected := range expectedHeaders {
		if tl.headers[i] != expected {
			t.Errorf("Header %d: expected '%s', got '%s'", i, expected, tl.headers[i])
		}
	}
}
