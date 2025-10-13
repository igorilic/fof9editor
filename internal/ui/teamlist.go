// ABOUTME: Team list view component for FOF9 Editor
// ABOUTME: Displays teams in a sortable, filterable table

package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/models"
)

// TeamList represents a list view for teams
type TeamList struct {
	container      *fyne.Container
	table          *widget.Table
	teams          []models.Team
	headers        []string
	selectedRow    int
	onSelectChange func(int)
}

// NewTeamList creates a new team list view
func NewTeamList() *TeamList {
	tl := &TeamList{
		teams:       []models.Team{},
		headers:     []string{"ID", "Team Name", "Nickname", "Abbreviation", "Conference", "Division"},
		selectedRow: -1,
	}

	tl.setupTable()
	return tl
}

// setupTable creates and configures the table widget
func (tl *TeamList) setupTable() {
	tl.table = widget.NewTable(
		func() (int, int) {
			return len(tl.teams) + 1, len(tl.headers) // +1 for header row
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			// Header row
			if id.Row == 0 {
				label.SetText(tl.headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
				return
			}

			// Data rows
			teamIdx := id.Row - 1
			if teamIdx >= len(tl.teams) {
				label.SetText("")
				return
			}

			team := tl.teams[teamIdx]
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", team.TeamID))
			case 1:
				label.SetText(team.TeamName)
			case 2:
				label.SetText(team.NickName)
			case 3:
				label.SetText(team.Abbreviation)
			case 4:
				label.SetText(fmt.Sprintf("%d", team.Conference))
			case 5:
				label.SetText(fmt.Sprintf("%d", team.Division))
			default:
				label.SetText("")
			}
			label.TextStyle = fyne.TextStyle{}
		},
	)

	// Set column widths
	tl.table.SetColumnWidth(0, 60)  // ID
	tl.table.SetColumnWidth(1, 150) // Team Name
	tl.table.SetColumnWidth(2, 150) // Nickname
	tl.table.SetColumnWidth(3, 100) // Abbreviation
	tl.table.SetColumnWidth(4, 100) // Conference
	tl.table.SetColumnWidth(5, 80)  // Division

	// Set up selection callback
	tl.table.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 { // Skip header row
			tl.selectedRow = id.Row - 1
			if tl.onSelectChange != nil {
				tl.onSelectChange(tl.selectedRow)
			}
		}
	}

	tl.container = container.NewMax(tl.table)
}

// SetTeams updates the displayed teams
func (tl *TeamList) SetTeams(teams []models.Team) {
	tl.teams = teams
	tl.table.Refresh()
}

// GetTeams returns the current list of teams
func (tl *TeamList) GetTeams() []models.Team {
	return tl.teams
}

// GetContainer returns the list container
func (tl *TeamList) GetContainer() *fyne.Container {
	return tl.container
}

// GetSelectedTeam returns the currently selected team, or nil if none selected
func (tl *TeamList) GetSelectedTeam() *models.Team {
	if tl.selectedRow < 0 || tl.selectedRow >= len(tl.teams) {
		return nil
	}
	return &tl.teams[tl.selectedRow]
}

// SetOnSelectChange sets the callback for when a team is selected
func (tl *TeamList) SetOnSelectChange(callback func(int)) {
	tl.onSelectChange = callback
}

// Clear removes all teams from the list
func (tl *TeamList) Clear() {
	tl.teams = []models.Team{}
	tl.table.Refresh()
}
