// ABOUTME: Tests for player list view component
// ABOUTME: Validates player display and table functionality

package ui

import (
	"testing"

	"github.com/igorilic/fof9editor/internal/models"
)

func TestNewPlayerList(t *testing.T) {
	pl := NewPlayerList()

	if pl == nil {
		t.Fatal("NewPlayerList returned nil")
	}

	if pl.container == nil {
		t.Fatal("PlayerList container is nil")
	}

	if pl.table == nil {
		t.Fatal("PlayerList table is nil")
	}

	if len(pl.headers) == 0 {
		t.Fatal("PlayerList headers is empty")
	}
}

func TestPlayerList_SetPlayers(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 1, FirstName: "Tom", LastName: "Brady", PositionKey: 1, Team: 1, OverallRating: 99},
		{PlayerID: 2, FirstName: "Aaron", LastName: "Rodgers", PositionKey: 1, Team: 2, OverallRating: 98},
	}

	pl.SetPlayers(players)

	if len(pl.players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(pl.players))
	}

	if pl.players[0].FirstName != "Tom" {
		t.Errorf("Expected first player 'Tom', got '%s'", pl.players[0].FirstName)
	}
}

func TestPlayerList_GetPlayers(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 1, FirstName: "Tom", LastName: "Brady", PositionKey: 1, Team: 1, OverallRating: 99},
	}

	pl.SetPlayers(players)

	retrieved := pl.GetPlayers()
	if len(retrieved) != 1 {
		t.Errorf("Expected 1 player, got %d", len(retrieved))
	}

	if retrieved[0].PlayerID != 1 {
		t.Errorf("Expected player ID 1, got %d", retrieved[0].PlayerID)
	}
}

func TestPlayerList_GetContainer(t *testing.T) {
	pl := NewPlayerList()

	container := pl.GetContainer()
	if container == nil {
		t.Fatal("GetContainer returned nil")
	}

	if container != pl.container {
		t.Error("GetContainer did not return the correct container")
	}
}

func TestPlayerList_Clear(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 1, FirstName: "Tom", LastName: "Brady", PositionKey: 1, Team: 1, OverallRating: 99},
		{PlayerID: 2, FirstName: "Aaron", LastName: "Rodgers", PositionKey: 1, Team: 2, OverallRating: 98},
	}

	pl.SetPlayers(players)

	if len(pl.players) != 2 {
		t.Fatalf("Expected 2 players before clear, got %d", len(pl.players))
	}

	pl.Clear()

	if len(pl.players) != 0 {
		t.Errorf("Expected 0 players after clear, got %d", len(pl.players))
	}
}

func TestPlayerList_EmptyList(t *testing.T) {
	pl := NewPlayerList()

	// Should handle empty list without panic
	pl.SetPlayers([]models.Player{})

	if len(pl.players) != 0 {
		t.Errorf("Expected 0 players, got %d", len(pl.players))
	}
}

func TestPlayerList_Headers(t *testing.T) {
	pl := NewPlayerList()

	expectedHeaders := []string{"ID", "First Name", "Last Name", "Position", "Team", "Overall"}

	if len(pl.headers) != len(expectedHeaders) {
		t.Errorf("Expected %d headers, got %d", len(expectedHeaders), len(pl.headers))
	}

	for i, expected := range expectedHeaders {
		if pl.headers[i] != expected {
			t.Errorf("Header %d: expected '%s', got '%s'", i, expected, pl.headers[i])
		}
	}
}

func TestPlayerList_SortByColumn(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 3, FirstName: "Charlie", LastName: "Wilson", PositionKey: 1, Team: 3, OverallRating: 85},
		{PlayerID: 1, FirstName: "Alice", LastName: "Smith", PositionKey: 2, Team: 1, OverallRating: 90},
		{PlayerID: 2, FirstName: "Bob", LastName: "Johnson", PositionKey: 1, Team: 2, OverallRating: 88},
	}

	pl.SetPlayers(players)

	// Sort by ID (column 0) ascending
	pl.SortByColumn(0)
	if pl.players[0].PlayerID != 1 {
		t.Errorf("After sorting by ID asc, expected first ID 1, got %d", pl.players[0].PlayerID)
	}
	if pl.sortColumn != 0 {
		t.Errorf("Expected sortColumn 0, got %d", pl.sortColumn)
	}
	if !pl.sortAscending {
		t.Error("Expected sortAscending true")
	}

	// Click same column to toggle descending
	pl.SortByColumn(0)
	if pl.players[0].PlayerID != 3 {
		t.Errorf("After sorting by ID desc, expected first ID 3, got %d", pl.players[0].PlayerID)
	}
	if pl.sortAscending {
		t.Error("Expected sortAscending false after toggle")
	}
}

func TestPlayerList_SortByFirstName(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 3, FirstName: "Charlie", LastName: "Wilson", PositionKey: 1, Team: 3, OverallRating: 85},
		{PlayerID: 1, FirstName: "Alice", LastName: "Smith", PositionKey: 2, Team: 1, OverallRating: 90},
		{PlayerID: 2, FirstName: "Bob", LastName: "Johnson", PositionKey: 1, Team: 2, OverallRating: 88},
	}

	pl.SetPlayers(players)

	// Sort by First Name (column 1)
	pl.SortByColumn(1)

	if pl.players[0].FirstName != "Alice" {
		t.Errorf("After sorting by first name, expected first 'Alice', got '%s'", pl.players[0].FirstName)
	}
	if pl.players[2].FirstName != "Charlie" {
		t.Errorf("After sorting by first name, expected last 'Charlie', got '%s'", pl.players[2].FirstName)
	}
}

func TestPlayerList_SortByOverallRating(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 3, FirstName: "Charlie", LastName: "Wilson", PositionKey: 1, Team: 3, OverallRating: 85},
		{PlayerID: 1, FirstName: "Alice", LastName: "Smith", PositionKey: 2, Team: 1, OverallRating: 90},
		{PlayerID: 2, FirstName: "Bob", LastName: "Johnson", PositionKey: 1, Team: 2, OverallRating: 88},
	}

	pl.SetPlayers(players)

	// Sort by Overall (column 5) ascending
	pl.SortByColumn(5)

	if pl.players[0].OverallRating != 85 {
		t.Errorf("After sorting by overall asc, expected first 85, got %d", pl.players[0].OverallRating)
	}
	if pl.players[2].OverallRating != 90 {
		t.Errorf("After sorting by overall asc, expected last 90, got %d", pl.players[2].OverallRating)
	}
}

func TestPlayerList_SortInvalidColumn(t *testing.T) {
	pl := NewPlayerList()

	players := []models.Player{
		{PlayerID: 1, FirstName: "Alice", LastName: "Smith", PositionKey: 2, Team: 1, OverallRating: 90},
	}

	pl.SetPlayers(players)

	// Try to sort by invalid column
	pl.SortByColumn(99)

	// Should not crash and sortColumn should remain -1
	if pl.sortColumn != -1 {
		t.Errorf("Expected sortColumn to remain -1, got %d", pl.sortColumn)
	}
}
