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
	sortColumn     int
	sortAscending  bool
}

// NewPlayerList creates a new player list view
func NewPlayerList() *PlayerList {
	pl := &PlayerList{
		players:       []models.Player{},
		headers:       []string{"ID", "First Name", "Last Name", "Position", "Team", "Overall"},
		selectedRow:   -1,
		sortColumn:    -1,
		sortAscending: true,
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
				headerText := pl.headers[id.Col]
				// Add sort indicator if this column is sorted
				if pl.sortColumn == id.Col {
					if pl.sortAscending {
						headerText += " ▲"
					} else {
						headerText += " ▼"
					}
				}
				label.SetText(headerText)
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
		if id.Row == 0 {
			// Header row clicked - trigger sort
			pl.SortByColumn(id.Col)
		} else {
			// Data row clicked - select player
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
	pl.selectedRow = -1
	pl.table.Refresh()
}

// SortByColumn sorts the players by the specified column
func (pl *PlayerList) SortByColumn(column int) {
	if column < 0 || column >= len(pl.headers) {
		return
	}

	// Toggle sort direction if clicking same column
	if pl.sortColumn == column {
		pl.sortAscending = !pl.sortAscending
	} else {
		pl.sortColumn = column
		pl.sortAscending = true
	}

	pl.sortPlayers()
	pl.table.Refresh()
}

// sortPlayers sorts the players based on current sort settings
func (pl *PlayerList) sortPlayers() {
	if pl.sortColumn < 0 || len(pl.players) == 0 {
		return
	}

	// Use a simple bubble sort for clarity (for large datasets, use sort.Slice)
	for i := 0; i < len(pl.players)-1; i++ {
		for j := i + 1; j < len(pl.players); j++ {
			if pl.comparePlayer(i, j) {
				pl.players[i], pl.players[j] = pl.players[j], pl.players[i]
			}
		}
	}
}

// comparePlayer compares two players based on current sort column
func (pl *PlayerList) comparePlayer(i, j int) bool {
	p1, p2 := pl.players[i], pl.players[j]

	var result bool
	switch pl.sortColumn {
	case 0: // ID
		result = p1.PlayerID > p2.PlayerID
	case 1: // First Name
		result = p1.FirstName > p2.FirstName
	case 2: // Last Name
		result = p1.LastName > p2.LastName
	case 3: // Position
		result = p1.PositionKey > p2.PositionKey
	case 4: // Team
		result = p1.Team > p2.Team
	case 5: // Overall
		result = p1.OverallRating > p2.OverallRating
	default:
		return false
	}

	if !pl.sortAscending {
		result = !result
	}
	return result
}
