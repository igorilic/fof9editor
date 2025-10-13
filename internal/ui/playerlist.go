// ABOUTME: Player list view component for FOF9 Editor
// ABOUTME: Displays players in a sortable, filterable table

package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/models"
)

// PlayerList represents a list view for players
type PlayerList struct {
	container      *fyne.Container
	table          *widget.Table
	players        []models.Player
	headers        []string
	selectedRow    int
	onSelectChange func(int)
}

// NewPlayerList creates a new player list view
func NewPlayerList() *PlayerList {
	pl := &PlayerList{
		players:     []models.Player{},
		headers:     []string{"ID", "First Name", "Last Name", "Position", "Team", "Overall"},
		selectedRow: -1,
	}

	pl.setupTable()
	return pl
}

// setupTable creates and configures the table widget
func (pl *PlayerList) setupTable() {
	pl.table = widget.NewTable(
		func() (int, int) {
			return len(pl.players) + 1, len(pl.headers) // +1 for header row
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			// Header row
			if id.Row == 0 {
				label.SetText(pl.headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
				return
			}

			// Data rows
			playerIdx := id.Row - 1
			if playerIdx >= len(pl.players) {
				label.SetText("")
				return
			}

			player := pl.players[playerIdx]
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", player.PlayerID))
			case 1:
				label.SetText(player.FirstName)
			case 2:
				label.SetText(player.LastName)
			case 3:
				label.SetText(fmt.Sprintf("%d", player.PositionKey))
			case 4:
				label.SetText(fmt.Sprintf("%d", player.Team))
			case 5:
				label.SetText(fmt.Sprintf("%d", player.OverallRating))
			default:
				label.SetText("")
			}
			label.TextStyle = fyne.TextStyle{}
		},
	)

	// Set column widths
	pl.table.SetColumnWidth(0, 60)  // ID
	pl.table.SetColumnWidth(1, 120) // First Name
	pl.table.SetColumnWidth(2, 120) // Last Name
	pl.table.SetColumnWidth(3, 80)  // Position
	pl.table.SetColumnWidth(4, 80)  // Team
	pl.table.SetColumnWidth(5, 80)  // Overall

	// Set up selection callback
	pl.table.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 { // Skip header row
			pl.selectedRow = id.Row - 1
			if pl.onSelectChange != nil {
				pl.onSelectChange(pl.selectedRow)
			}
		}
	}

	pl.container = container.NewMax(pl.table)
}

// SetPlayers updates the displayed players
func (pl *PlayerList) SetPlayers(players []models.Player) {
	pl.players = players
	pl.table.Refresh()
}

// GetPlayers returns the current list of players
func (pl *PlayerList) GetPlayers() []models.Player {
	return pl.players
}

// GetContainer returns the list container
func (pl *PlayerList) GetContainer() *fyne.Container {
	return pl.container
}

// GetSelectedPlayer returns the currently selected player, or nil if none selected
func (pl *PlayerList) GetSelectedPlayer() *models.Player {
	if pl.selectedRow < 0 || pl.selectedRow >= len(pl.players) {
		return nil
	}
	return &pl.players[pl.selectedRow]
}

// SetOnSelectChange sets the callback for when a player is selected
func (pl *PlayerList) SetOnSelectChange(callback func(int)) {
	pl.onSelectChange = callback
}

// Clear removes all players from the list
func (pl *PlayerList) Clear() {
	pl.players = []models.Player{}
	pl.table.Refresh()
}
