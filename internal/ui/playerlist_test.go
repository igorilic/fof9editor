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
