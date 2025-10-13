// ABOUTME: Player list view component for FOF9 Editor
// ABOUTME: Displays players in a sortable, filterable table

package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/models"
)

// PlayerList represents a list view for players
type PlayerList struct {
	container       *fyne.Container
	table           *widget.Table
	players         []models.Player
	filteredPlayers []models.Player
	headers         []string
	selectedRow     int
	onSelectChange  func(int)
	sortColumn      int
	sortAscending   bool
	filterText      string
	searchEntry     *widget.Entry
}

// NewPlayerList creates a new player list view
func NewPlayerList() *PlayerList {
	pl := &PlayerList{
		players:         []models.Player{},
		filteredPlayers: []models.Player{},
		headers:         []string{"ID", "First Name", "Last Name", "Position", "Team", "Overall"},
		selectedRow:     -1,
		sortColumn:      -1,
		sortAscending:   true,
		filterText:      "",
	}

	pl.setupUI()
	return pl
}

// setupUI creates and configures the UI components
func (pl *PlayerList) setupUI() {
	// Create search entry
	pl.searchEntry = widget.NewEntry()
	pl.searchEntry.SetPlaceHolder("Search players...")
	pl.searchEntry.OnChanged = func(text string) {
		pl.filterText = text
		pl.applyFilter()
	}

	pl.setupTable()

	// Create container with search on top and table below
	pl.container = container.NewBorder(
		pl.searchEntry, // top
		nil,            // bottom
		nil,            // left
		nil,            // right
		container.NewMax(pl.table), // center
	)
}

// setupTable creates and configures the table widget
func (pl *PlayerList) setupTable() {
	pl.table = widget.NewTable(
		func() (int, int) {
			return len(pl.filteredPlayers) + 1, len(pl.headers) // +1 for header row
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
			if playerIdx >= len(pl.filteredPlayers) {
				label.SetText("")
				return
			}

			player := pl.filteredPlayers[playerIdx]
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
}

// SetPlayers updates the displayed players
func (pl *PlayerList) SetPlayers(players []models.Player) {
	pl.players = players
	pl.applyFilter()
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
	if pl.selectedRow < 0 || pl.selectedRow >= len(pl.filteredPlayers) {
		return nil
	}
	return &pl.filteredPlayers[pl.selectedRow]
}

// SetOnSelectChange sets the callback for when a player is selected
func (pl *PlayerList) SetOnSelectChange(callback func(int)) {
	pl.onSelectChange = callback
}

// Clear removes all players from the list
func (pl *PlayerList) Clear() {
	pl.players = []models.Player{}
	pl.filteredPlayers = []models.Player{}
	pl.selectedRow = -1
	pl.filterText = ""
	if pl.searchEntry != nil {
		pl.searchEntry.SetText("")
	}
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

// sortPlayers sorts the filtered players based on current sort settings
func (pl *PlayerList) sortPlayers() {
	if pl.sortColumn < 0 || len(pl.filteredPlayers) == 0 {
		return
	}

	// Use a simple bubble sort for clarity (for large datasets, use sort.Slice)
	for i := 0; i < len(pl.filteredPlayers)-1; i++ {
		for j := i + 1; j < len(pl.filteredPlayers); j++ {
			if pl.comparePlayer(i, j) {
				pl.filteredPlayers[i], pl.filteredPlayers[j] = pl.filteredPlayers[j], pl.filteredPlayers[i]
			}
		}
	}
}

// comparePlayer compares two players based on current sort column
func (pl *PlayerList) comparePlayer(i, j int) bool {
	p1, p2 := pl.filteredPlayers[i], pl.filteredPlayers[j]

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

// applyFilter applies the current filter text to the player list
func (pl *PlayerList) applyFilter() {
	if pl.filterText == "" {
		// No filter - show all players
		pl.filteredPlayers = pl.players
	} else {
		// Filter players based on search text
		pl.filteredPlayers = []models.Player{}
		for _, player := range pl.players {
			if pl.matchesFilter(player) {
				pl.filteredPlayers = append(pl.filteredPlayers, player)
			}
		}
	}

	// Re-apply sort if active
	if pl.sortColumn >= 0 {
		pl.sortPlayers()
	}

	pl.table.Refresh()
}

// matchesFilter checks if a player matches the current filter text
func (pl *PlayerList) matchesFilter(player models.Player) bool {
	if pl.filterText == "" {
		return true
	}

	filter := strings.ToLower(pl.filterText)

	// Search in ID, first name, last name
	if strings.Contains(strings.ToLower(fmt.Sprintf("%d", player.PlayerID)), filter) {
		return true
	}
	if strings.Contains(strings.ToLower(player.FirstName), filter) {
		return true
	}
	if strings.Contains(strings.ToLower(player.LastName), filter) {
		return true
	}

	return false
}
