// ABOUTME: Tests for player CSV loading functionality
// ABOUTME: Validates CSV parsing and field mapping for Player structs

package data

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPlayers_SimpleFile(t *testing.T) {
	players, err := LoadPlayers("../../testdata/fixtures/csv/players_simple.csv")
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if len(players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(players))
	}

	// Check first player (Tom Brady)
	brady := players[0]
	if brady.PlayerID != 1000 {
		t.Errorf("Expected PlayerID 1000, got %d", brady.PlayerID)
	}
	if brady.LastName != "Brady" {
		t.Errorf("Expected LastName 'Brady', got '%s'", brady.LastName)
	}
	if brady.FirstName != "Tom" {
		t.Errorf("Expected FirstName 'Tom', got '%s'", brady.FirstName)
	}
	if brady.Team != 1 {
		t.Errorf("Expected Team 1, got %d", brady.Team)
	}
	if brady.PositionKey != 1 {
		t.Errorf("Expected PositionKey 1, got %d", brady.PositionKey)
	}
	if brady.Uniform != 12 {
		t.Errorf("Expected Uniform 12, got %d", brady.Uniform)
	}
	if brady.Height != 76 {
		t.Errorf("Expected Height 76, got %d", brady.Height)
	}
	if brady.Weight != 225 {
		t.Errorf("Expected Weight 225, got %d", brady.Weight)
	}
	if brady.BirthYear != 1977 {
		t.Errorf("Expected BirthYear 1977, got %d", brady.BirthYear)
	}
	if brady.College != "Michigan" {
		t.Errorf("Expected College 'Michigan', got '%s'", brady.College)
	}
	if brady.Experience != 25 {
		t.Errorf("Expected Experience 25, got %d", brady.Experience)
	}
	if brady.SalaryYear1 != 10000000 {
		t.Errorf("Expected SalaryYear1 10000000, got %d", brady.SalaryYear1)
	}

	// Check second player (Patrick Mahomes)
	mahomes := players[1]
	if mahomes.PlayerID != 1001 {
		t.Errorf("Expected PlayerID 1001, got %d", mahomes.PlayerID)
	}
	if mahomes.LastName != "Mahomes" {
		t.Errorf("Expected LastName 'Mahomes', got '%s'", mahomes.LastName)
	}
	if mahomes.FirstName != "Patrick" {
		t.Errorf("Expected FirstName 'Patrick', got '%s'", mahomes.FirstName)
	}
	if mahomes.SalaryYear1 != 45000000 {
		t.Errorf("Expected SalaryYear1 45000000, got %d", mahomes.SalaryYear1)
	}
}

func TestLoadPlayers_EmptyFile(t *testing.T) {
	players, err := LoadPlayers("../../testdata/fixtures/csv/empty.csv")
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if len(players) != 0 {
		t.Errorf("Expected 0 players for empty file, got %d", len(players))
	}
}

func TestLoadPlayers_NonExistentFile(t *testing.T) {
	_, err := LoadPlayers("nonexistent.csv")
	if err == nil {
		t.Fatal("Expected error for non-existent file, got nil")
	}
}

func TestLoadPlayers_MissingColumns(t *testing.T) {
	// Create temporary file with missing columns
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "missing_cols.csv")

	content := `PLAYERID,LASTNAME,FIRSTNAME
1000,Brady,Tom
1001,Mahomes,Patrick
`
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	players, err := LoadPlayers(tmpFile)
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if len(players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(players))
	}

	// Check that missing columns default to zero values
	if players[0].PlayerID != 1000 {
		t.Errorf("Expected PlayerID 1000, got %d", players[0].PlayerID)
	}
	if players[0].Team != 0 {
		t.Errorf("Expected Team 0 (zero value), got %d", players[0].Team)
	}
	if players[0].Height != 0 {
		t.Errorf("Expected Height 0 (zero value), got %d", players[0].Height)
	}
}

func TestMapRowToPlayer_ValidData(t *testing.T) {
	row := map[string]string{
		"PLAYERID":  "1000",
		"LASTNAME":  "Test",
		"FIRSTNAME": "Player",
		"TEAM":      "1",
		"HEIGHT":    "72",
		"WEIGHT":    "200",
	}

	player, err := mapRowToPlayer(row)
	if err != nil {
		t.Fatalf("mapRowToPlayer failed: %v", err)
	}

	if player.PlayerID != 1000 {
		t.Errorf("Expected PlayerID 1000, got %d", player.PlayerID)
	}
	if player.LastName != "Test" {
		t.Errorf("Expected LastName 'Test', got '%s'", player.LastName)
	}
	if player.FirstName != "Player" {
		t.Errorf("Expected FirstName 'Player', got '%s'", player.FirstName)
	}
	if player.Team != 1 {
		t.Errorf("Expected Team 1, got %d", player.Team)
	}
	if player.Height != 72 {
		t.Errorf("Expected Height 72, got %d", player.Height)
	}
	if player.Weight != 200 {
		t.Errorf("Expected Weight 200, got %d", player.Weight)
	}
}

func TestMapRowToPlayer_EmptyValues(t *testing.T) {
	row := map[string]string{
		"PLAYERID":  "1000",
		"LASTNAME":  "Test",
		"FIRSTNAME": "Player",
		"TEAM":      "",
		"HEIGHT":    "",
	}

	player, err := mapRowToPlayer(row)
	if err != nil {
		t.Fatalf("mapRowToPlayer failed: %v", err)
	}

	if player.PlayerID != 1000 {
		t.Errorf("Expected PlayerID 1000, got %d", player.PlayerID)
	}
	if player.Team != 0 {
		t.Errorf("Expected Team 0 (zero value for empty string), got %d", player.Team)
	}
	if player.Height != 0 {
		t.Errorf("Expected Height 0 (zero value for empty string), got %d", player.Height)
	}
}

func TestMapRowToPlayer_InvalidIntegerValue(t *testing.T) {
	row := map[string]string{
		"PLAYERID": "not_a_number",
		"LASTNAME": "Test",
	}

	_, err := mapRowToPlayer(row)
	if err == nil {
		t.Fatal("Expected error for invalid integer value, got nil")
	}
}

func TestLoadPlayers_IntegrationWithRealStructure(t *testing.T) {
	// This test ensures all fields in the Player struct can be mapped
	players, err := LoadPlayers("../../testdata/fixtures/csv/players_simple.csv")
	if err != nil {
		t.Fatalf("LoadPlayers failed: %v", err)
	}

	if len(players) < 1 {
		t.Fatal("Expected at least 1 player")
	}

	// Verify that complex fields are properly loaded
	p := players[0]
	if p.SalaryYear1 != 10000000 {
		t.Errorf("Expected SalaryYear1 10000000, got %d", p.SalaryYear1)
	}
	if p.SalaryYear2 != 11000000 {
		t.Errorf("Expected SalaryYear2 11000000, got %d", p.SalaryYear2)
	}
	if p.BonusYear1 != 0 {
		t.Errorf("Expected BonusYear1 0 (not in CSV), got %d", p.BonusYear1)
	}
}
