// ABOUTME: Tests for coach and team CSV loading functionality
// ABOUTME: Validates CSV parsing and field mapping for Coach and Team structs

package data

import (
	"os"
	"path/filepath"
	"testing"
)

// Coach Loading Tests

func TestLoadCoaches_SimpleFile(t *testing.T) {
	coaches, err := LoadCoaches("../../testdata/fixtures/csv/coaches_simple.csv")
	if err != nil {
		t.Fatalf("LoadCoaches failed: %v", err)
	}

	if len(coaches) != 2 {
		t.Fatalf("Expected 2 coaches, got %d", len(coaches))
	}

	// Check first coach (Bill Belichick)
	belichick := coaches[0]
	if belichick.LastName != "Belichick" {
		t.Errorf("Expected LastName 'Belichick', got '%s'", belichick.LastName)
	}
	if belichick.FirstName != "Bill" {
		t.Errorf("Expected FirstName 'Bill', got '%s'", belichick.FirstName)
	}
	if belichick.Team != 1 {
		t.Errorf("Expected Team 1, got %d", belichick.Team)
	}
	if belichick.Position != 0 {
		t.Errorf("Expected Position 0 (Head Coach), got %d", belichick.Position)
	}
	if belichick.BirthYear != 1952 {
		t.Errorf("Expected BirthYear 1952, got %d", belichick.BirthYear)
	}
	if belichick.College != "Wesleyan" {
		t.Errorf("Expected College 'Wesleyan', got '%s'", belichick.College)
	}

	// Check second coach (Andy Reid)
	reid := coaches[1]
	if reid.LastName != "Reid" {
		t.Errorf("Expected LastName 'Reid', got '%s'", reid.LastName)
	}
	if reid.FirstName != "Andy" {
		t.Errorf("Expected FirstName 'Andy', got '%s'", reid.FirstName)
	}
	if reid.Team != 2 {
		t.Errorf("Expected Team 2, got %d", reid.Team)
	}
}

func TestLoadCoaches_EmptyFile(t *testing.T) {
	coaches, err := LoadCoaches("../../testdata/fixtures/csv/empty.csv")
	if err != nil {
		t.Fatalf("LoadCoaches failed: %v", err)
	}

	if len(coaches) != 0 {
		t.Errorf("Expected 0 coaches for empty file, got %d", len(coaches))
	}
}

func TestLoadCoaches_NonExistentFile(t *testing.T) {
	_, err := LoadCoaches("nonexistent.csv")
	if err == nil {
		t.Fatal("Expected error for non-existent file, got nil")
	}
}

func TestLoadCoaches_MissingColumns(t *testing.T) {
	// Create temporary file with missing columns
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "missing_cols.csv")

	content := `LASTNAME,FIRSTNAME,TEAM
Belichick,Bill,1
Reid,Andy,2
`
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	coaches, err := LoadCoaches(tmpFile)
	if err != nil {
		t.Fatalf("LoadCoaches failed: %v", err)
	}

	if len(coaches) != 2 {
		t.Fatalf("Expected 2 coaches, got %d", len(coaches))
	}

	// Check that missing columns default to zero values
	if coaches[0].LastName != "Belichick" {
		t.Errorf("Expected LastName 'Belichick', got '%s'", coaches[0].LastName)
	}
	if coaches[0].Position != 0 {
		t.Errorf("Expected Position 0 (zero value), got %d", coaches[0].Position)
	}
}

func TestMapRowToCoach_ValidData(t *testing.T) {
	row := map[string]string{
		"LASTNAME":  "Test",
		"FIRSTNAME": "Coach",
		"TEAM":      "1",
		"POSITION":  "0",
		"BIRTHYEAR": "1960",
	}

	coach, err := mapRowToCoach(row)
	if err != nil {
		t.Fatalf("mapRowToCoach failed: %v", err)
	}

	if coach.LastName != "Test" {
		t.Errorf("Expected LastName 'Test', got '%s'", coach.LastName)
	}
	if coach.FirstName != "Coach" {
		t.Errorf("Expected FirstName 'Coach', got '%s'", coach.FirstName)
	}
	if coach.Team != 1 {
		t.Errorf("Expected Team 1, got %d", coach.Team)
	}
	if coach.Position != 0 {
		t.Errorf("Expected Position 0, got %d", coach.Position)
	}
	if coach.BirthYear != 1960 {
		t.Errorf("Expected BirthYear 1960, got %d", coach.BirthYear)
	}
}

// Team Loading Tests

func TestLoadTeams_SimpleFile(t *testing.T) {
	teams, err := LoadTeams("../../testdata/fixtures/csv/teams_simple.csv")
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	if len(teams) != 2 {
		t.Fatalf("Expected 2 teams, got %d", len(teams))
	}

	// Check first team (Patriots)
	patriots := teams[0]
	if patriots.TeamID != 1 {
		t.Errorf("Expected TeamID 1, got %d", patriots.TeamID)
	}
	if patriots.Year != 2024 {
		t.Errorf("Expected Year 2024, got %d", patriots.Year)
	}
	if patriots.TeamName != "New England Patriots" {
		t.Errorf("Expected TeamName 'New England Patriots', got '%s'", patriots.TeamName)
	}
	if patriots.NickName != "Patriots" {
		t.Errorf("Expected NickName 'Patriots', got '%s'", patriots.NickName)
	}
	if patriots.Abbreviation != "NE" {
		t.Errorf("Expected Abbreviation 'NE', got '%s'", patriots.Abbreviation)
	}
	if patriots.City != 1234 {
		t.Errorf("Expected City 1234, got %d", patriots.City)
	}
	if patriots.PrimaryRed != 0 || patriots.PrimaryGreen != 0 || patriots.PrimaryBlue != 128 {
		t.Errorf("Expected Primary RGB(0,0,128), got RGB(%d,%d,%d)",
			patriots.PrimaryRed, patriots.PrimaryGreen, patriots.PrimaryBlue)
	}
	if patriots.Capacity != 65878 {
		t.Errorf("Expected Capacity 65878, got %d", patriots.Capacity)
	}

	// Check second team (Chiefs)
	chiefs := teams[1]
	if chiefs.TeamID != 2 {
		t.Errorf("Expected TeamID 2, got %d", chiefs.TeamID)
	}
	if chiefs.City != 5678 {
		t.Errorf("Expected City 5678, got %d", chiefs.City)
	}
	if chiefs.NickName != "Chiefs" {
		t.Errorf("Expected NickName 'Chiefs', got '%s'", chiefs.NickName)
	}
}

func TestLoadTeams_EmptyFile(t *testing.T) {
	teams, err := LoadTeams("../../testdata/fixtures/csv/empty.csv")
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	if len(teams) != 0 {
		t.Errorf("Expected 0 teams for empty file, got %d", len(teams))
	}
}

func TestLoadTeams_NonExistentFile(t *testing.T) {
	_, err := LoadTeams("nonexistent.csv")
	if err == nil {
		t.Fatal("Expected error for non-existent file, got nil")
	}
}

func TestLoadTeams_MissingColumns(t *testing.T) {
	// Create temporary file with missing columns
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "missing_cols.csv")

	content := `TEAMID,TEAMNAME,NICKNAME
1,New England Patriots,Patriots
2,Kansas City Chiefs,Chiefs
`
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	teams, err := LoadTeams(tmpFile)
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	if len(teams) != 2 {
		t.Fatalf("Expected 2 teams, got %d", len(teams))
	}

	// Check that missing columns default to zero values
	if teams[0].TeamID != 1 {
		t.Errorf("Expected TeamID 1, got %d", teams[0].TeamID)
	}
	if teams[0].Abbreviation != "" {
		t.Errorf("Expected empty Abbreviation (zero value), got '%s'", teams[0].Abbreviation)
	}
	if teams[0].Conference != 0 {
		t.Errorf("Expected Conference 0 (zero value), got %d", teams[0].Conference)
	}
}

func TestMapRowToTeam_ValidData(t *testing.T) {
	row := map[string]string{
		"TEAMID":       "1",
		"TEAMNAME":     "Test Team",
		"NICKNAME":     "Team",
		"ABBREVIATION": "TC",
		"CITY":         "1234",
		"CONFERENCE":   "0",
	}

	team, err := mapRowToTeam(row)
	if err != nil {
		t.Fatalf("mapRowToTeam failed: %v", err)
	}

	if team.TeamID != 1 {
		t.Errorf("Expected TeamID 1, got %d", team.TeamID)
	}
	if team.TeamName != "Test Team" {
		t.Errorf("Expected TeamName 'Test Team', got '%s'", team.TeamName)
	}
	if team.NickName != "Team" {
		t.Errorf("Expected NickName 'Team', got '%s'", team.NickName)
	}
	if team.Abbreviation != "TC" {
		t.Errorf("Expected Abbreviation 'TC', got '%s'", team.Abbreviation)
	}
	if team.City != 1234 {
		t.Errorf("Expected City 1234, got %d", team.City)
	}
	if team.Conference != 0 {
		t.Errorf("Expected Conference 0, got %d", team.Conference)
	}
}
