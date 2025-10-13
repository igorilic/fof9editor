// ABOUTME: Integration tests for CSV reading and writing
// ABOUTME: Tests round-trip operations (load -> modify -> save -> load)

package data

import (
	"path/filepath"
	"testing"

	"github.com/igorilic/fof9editor/internal/models"
)

// Player Save/Load Tests

func TestSavePlayers_RoundTrip(t *testing.T) {
	// Load players from fixture
	players, err := LoadPlayers("../../testdata/fixtures/csv/players_simple.csv")
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if len(players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(players))
	}

	// Save to temp file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "players_output.csv")

	err = SavePlayers(tmpFile, players)
	if err != nil {
		t.Fatalf("SavePlayers failed: %v", err)
	}

	// Load back from saved file
	loadedPlayers, err := LoadPlayers(tmpFile)
	if err != nil {
		t.Fatalf("LoadPlayers (second) failed: %v", err)
	}

	// Verify count
	if len(loadedPlayers) != len(players) {
		t.Fatalf("Expected %d players after round-trip, got %d", len(players), len(loadedPlayers))
	}

	// Verify first player data
	original := players[0]
	loaded := loadedPlayers[0]

	if loaded.PlayerID != original.PlayerID {
		t.Errorf("PlayerID mismatch: expected %d, got %d", original.PlayerID, loaded.PlayerID)
	}
	if loaded.LastName != original.LastName {
		t.Errorf("LastName mismatch: expected '%s', got '%s'", original.LastName, loaded.LastName)
	}
	if loaded.FirstName != original.FirstName {
		t.Errorf("FirstName mismatch: expected '%s', got '%s'", original.FirstName, loaded.FirstName)
	}
	if loaded.Team != original.Team {
		t.Errorf("Team mismatch: expected %d, got %d", original.Team, loaded.Team)
	}
	if loaded.Height != original.Height {
		t.Errorf("Height mismatch: expected %d, got %d", original.Height, loaded.Height)
	}
	if loaded.BirthYear != original.BirthYear {
		t.Errorf("BirthYear mismatch: expected %d, got %d", original.BirthYear, loaded.BirthYear)
	}
	if loaded.College != original.College {
		t.Errorf("College mismatch: expected '%s', got '%s'", original.College, loaded.College)
	}
}

func TestSavePlayers_EmptySlice(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty_players.csv")

	players := []models.Player{}
	err := SavePlayers(tmpFile, players)
	if err != nil {
		t.Fatalf("SavePlayers failed: %v", err)
	}

	// Load back and verify empty
	loadedPlayers, err := LoadPlayers(tmpFile)
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if len(loadedPlayers) != 0 {
		t.Errorf("Expected 0 players, got %d", len(loadedPlayers))
	}
}

func TestSavePlayers_ModifyAndSave(t *testing.T) {
	// Load players
	players, err := LoadPlayers("../../testdata/fixtures/csv/players_simple.csv")
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	// Modify first player
	players[0].LastName = "Modified"
	players[0].FirstName = "Test"
	players[0].Height = 99

	// Save
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "modified_players.csv")
	err = SavePlayers(tmpFile, players)
	if err != nil {
		t.Fatalf("SavePlayers failed: %v", err)
	}

	// Load back and verify modifications
	loadedPlayers, err := LoadPlayers(tmpFile)
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if loadedPlayers[0].LastName != "Modified" {
		t.Errorf("Expected modified LastName 'Modified', got '%s'", loadedPlayers[0].LastName)
	}
	if loadedPlayers[0].FirstName != "Test" {
		t.Errorf("Expected modified FirstName 'Test', got '%s'", loadedPlayers[0].FirstName)
	}
	if loadedPlayers[0].Height != 99 {
		t.Errorf("Expected modified Height 99, got %d", loadedPlayers[0].Height)
	}
}

// Coach Save/Load Tests

func TestSaveCoaches_RoundTrip(t *testing.T) {
	// Load coaches from fixture
	coaches, err := LoadCoaches("../../testdata/fixtures/csv/coaches_simple.csv")
	if err != nil {
		t.Fatalf("LoadCoaches failed: %v", err)
	}

	if len(coaches) != 2 {
		t.Fatalf("Expected 2 coaches, got %d", len(coaches))
	}

	// Save to temp file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "coaches_output.csv")

	err = SaveCoaches(tmpFile, coaches)
	if err != nil {
		t.Fatalf("SaveCoaches failed: %v", err)
	}

	// Load back
	loadedCoaches, err := LoadCoaches(tmpFile)
	if err != nil {
		t.Fatalf("LoadCoaches (second) failed: %v", err)
	}

	// Verify count
	if len(loadedCoaches) != len(coaches) {
		t.Fatalf("Expected %d coaches after round-trip, got %d", len(coaches), len(loadedCoaches))
	}

	// Verify first coach data
	original := coaches[0]
	loaded := loadedCoaches[0]

	if loaded.LastName != original.LastName {
		t.Errorf("LastName mismatch: expected '%s', got '%s'", original.LastName, loaded.LastName)
	}
	if loaded.FirstName != original.FirstName {
		t.Errorf("FirstName mismatch: expected '%s', got '%s'", original.FirstName, loaded.FirstName)
	}
	if loaded.Team != original.Team {
		t.Errorf("Team mismatch: expected %d, got %d", original.Team, loaded.Team)
	}
	if loaded.Position != original.Position {
		t.Errorf("Position mismatch: expected %d, got %d", original.Position, loaded.Position)
	}
	if loaded.BirthYear != original.BirthYear {
		t.Errorf("BirthYear mismatch: expected %d, got %d", original.BirthYear, loaded.BirthYear)
	}
}

func TestSaveCoaches_EmptySlice(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty_coaches.csv")

	coaches := []models.Coach{}
	err := SaveCoaches(tmpFile, coaches)
	if err != nil {
		t.Fatalf("SaveCoaches failed: %v", err)
	}

	// Load back and verify empty
	loadedCoaches, err := LoadCoaches(tmpFile)
	if err != nil {
		t.Fatalf("LoadCoaches failed: %v", err)
	}

	if len(loadedCoaches) != 0 {
		t.Errorf("Expected 0 coaches, got %d", len(loadedCoaches))
	}
}

// Team Save/Load Tests

func TestSaveTeams_RoundTrip(t *testing.T) {
	// Load teams from fixture
	teams, err := LoadTeams("../../testdata/fixtures/csv/teams_simple.csv")
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	if len(teams) != 2 {
		t.Fatalf("Expected 2 teams, got %d", len(teams))
	}

	// Save to temp file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "teams_output.csv")

	err = SaveTeams(tmpFile, teams)
	if err != nil {
		t.Fatalf("SaveTeams failed: %v", err)
	}

	// Load back
	loadedTeams, err := LoadTeams(tmpFile)
	if err != nil {
		t.Fatalf("LoadTeams (second) failed: %v", err)
	}

	// Verify count
	if len(loadedTeams) != len(teams) {
		t.Fatalf("Expected %d teams after round-trip, got %d", len(teams), len(loadedTeams))
	}

	// Verify first team data
	original := teams[0]
	loaded := loadedTeams[0]

	if loaded.TeamID != original.TeamID {
		t.Errorf("TeamID mismatch: expected %d, got %d", original.TeamID, loaded.TeamID)
	}
	if loaded.TeamName != original.TeamName {
		t.Errorf("TeamName mismatch: expected '%s', got '%s'", original.TeamName, loaded.TeamName)
	}
	if loaded.NickName != original.NickName {
		t.Errorf("NickName mismatch: expected '%s', got '%s'", original.NickName, loaded.NickName)
	}
	if loaded.Abbreviation != original.Abbreviation {
		t.Errorf("Abbreviation mismatch: expected '%s', got '%s'", original.Abbreviation, loaded.Abbreviation)
	}
	if loaded.City != original.City {
		t.Errorf("City mismatch: expected %d, got %d", original.City, loaded.City)
	}
	if loaded.PrimaryRed != original.PrimaryRed {
		t.Errorf("PrimaryRed mismatch: expected %d, got %d", original.PrimaryRed, loaded.PrimaryRed)
	}
	if loaded.Capacity != original.Capacity {
		t.Errorf("Capacity mismatch: expected %d, got %d", original.Capacity, loaded.Capacity)
	}
}

func TestSaveTeams_EmptySlice(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty_teams.csv")

	teams := []models.Team{}
	err := SaveTeams(tmpFile, teams)
	if err != nil {
		t.Fatalf("SaveTeams failed: %v", err)
	}

	// Load back and verify empty
	loadedTeams, err := LoadTeams(tmpFile)
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	if len(loadedTeams) != 0 {
		t.Errorf("Expected 0 teams, got %d", len(loadedTeams))
	}
}

func TestSaveTeams_ModifyAndSave(t *testing.T) {
	// Load teams
	teams, err := LoadTeams("../../testdata/fixtures/csv/teams_simple.csv")
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	// Modify first team
	teams[0].NickName = "Modified Team"
	teams[0].Capacity = 100000

	// Save
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "modified_teams.csv")
	err = SaveTeams(tmpFile, teams)
	if err != nil {
		t.Fatalf("SaveTeams failed: %v", err)
	}

	// Load back and verify modifications
	loadedTeams, err := LoadTeams(tmpFile)
	if err != nil {
		t.Fatalf("LoadTeams failed: %v", err)
	}

	if loadedTeams[0].NickName != "Modified Team" {
		t.Errorf("Expected modified NickName 'Modified Team', got '%s'", loadedTeams[0].NickName)
	}
	if loadedTeams[0].Capacity != 100000 {
		t.Errorf("Expected modified Capacity 100000, got %d", loadedTeams[0].Capacity)
	}
}
