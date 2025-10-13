// ABOUTME: Coach list view component for FOF9 Editor
// ABOUTME: Displays coaches in a sortable, filterable table

package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/models"
)

// CoachList represents a list view for coaches
type CoachList struct {
	container      *fyne.Container
	table          *widget.Table
	coaches        []models.Coach
	headers        []string
	selectedRow    int
	onSelectChange func(int)
	sortColumn     int
	sortAscending  bool
}

// NewCoachList creates a new coach list view
func NewCoachList() *CoachList {
	cl := &CoachList{
		coaches:       []models.Coach{},
		headers:       []string{"First Name", "Last Name", "Position", "Team", "Pay Scale"},
		selectedRow:   -1,
		sortColumn:    -1,
		sortAscending: true,
	}

	cl.setupTable()
	return cl
}

// setupTable creates and configures the table widget
func (cl *CoachList) setupTable() {
	cl.table = widget.NewTable(
		func() (int, int) {
			return len(cl.coaches) + 1, len(cl.headers) // +1 for header row
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			// Header row
			if id.Row == 0 {
				headerText := cl.headers[id.Col]
				if cl.sortColumn == id.Col {
					if cl.sortAscending {
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
			coachIdx := id.Row - 1
			if coachIdx >= len(cl.coaches) {
				label.SetText("")
				return
			}

			coach := cl.coaches[coachIdx]
			switch id.Col {
			case 0:
				label.SetText(coach.FirstName)
			case 1:
				label.SetText(coach.LastName)
			case 2:
				label.SetText(cl.getPositionName(coach.Position))
			case 3:
				label.SetText(fmt.Sprintf("%d", coach.Team))
			case 4:
				label.SetText(fmt.Sprintf("$%dK", coach.PayScale*10))
			default:
				label.SetText("")
			}
			label.TextStyle = fyne.TextStyle{}
		},
	)

	// Set column widths
	cl.table.SetColumnWidth(0, 120) // First Name
	cl.table.SetColumnWidth(1, 120) // Last Name
	cl.table.SetColumnWidth(2, 150) // Position
	cl.table.SetColumnWidth(3, 80)  // Team
	cl.table.SetColumnWidth(4, 100) // Pay Scale

	// Set up selection callback
	cl.table.OnSelected = func(id widget.TableCellID) {
		if id.Row == 0 {
			// Header row clicked - trigger sort
			cl.SortByColumn(id.Col)
		} else {
			// Data row clicked - select coach
			cl.selectedRow = id.Row - 1
			if cl.onSelectChange != nil {
				cl.onSelectChange(cl.selectedRow)
			}
		}
	}

	cl.container = container.NewMax(cl.table)
}

// getPositionName converts position code to display name
func (cl *CoachList) getPositionName(position int) string {
	switch position {
	case models.PositionHeadCoach:
		return "Head Coach"
	case models.PositionOffensiveCoordinator:
		return "Offensive Coordinator"
	case models.PositionDefensiveCoordinator:
		return "Defensive Coordinator"
	case models.PositionSpecialTeamsCoordinator:
		return "Special Teams Coordinator"
	case models.PositionStrengthConditioning:
		return "Strength & Conditioning"
	default:
		return fmt.Sprintf("Position %d", position)
	}
}

// SetCoaches updates the displayed coaches
func (cl *CoachList) SetCoaches(coaches []models.Coach) {
	cl.coaches = coaches
	cl.table.Refresh()
}

// GetCoaches returns the current list of coaches
func (cl *CoachList) GetCoaches() []models.Coach {
	return cl.coaches
}

// GetContainer returns the list container
func (cl *CoachList) GetContainer() *fyne.Container {
	return cl.container
}

// GetSelectedCoach returns the currently selected coach, or nil if none selected
func (cl *CoachList) GetSelectedCoach() *models.Coach {
	if cl.selectedRow < 0 || cl.selectedRow >= len(cl.coaches) {
		return nil
	}
	return &cl.coaches[cl.selectedRow]
}

// SetOnSelectChange sets the callback for when a coach is selected
func (cl *CoachList) SetOnSelectChange(callback func(int)) {
	cl.onSelectChange = callback
}

// Clear removes all coaches from the list
func (cl *CoachList) Clear() {
	cl.coaches = []models.Coach{}
	cl.selectedRow = -1
	cl.table.Refresh()
}

// SortByColumn sorts the coaches by the specified column
func (cl *CoachList) SortByColumn(column int) {
	if column < 0 || column >= len(cl.headers) {
		return
	}

	if cl.sortColumn == column {
		cl.sortAscending = !cl.sortAscending
	} else {
		cl.sortColumn = column
		cl.sortAscending = true
	}

	cl.sortCoaches()
	cl.table.Refresh()
}

// sortCoaches sorts the coaches based on current sort settings
func (cl *CoachList) sortCoaches() {
	if cl.sortColumn < 0 || len(cl.coaches) == 0 {
		return
	}

	for i := 0; i < len(cl.coaches)-1; i++ {
		for j := i + 1; j < len(cl.coaches); j++ {
			if cl.compareCoach(i, j) {
				cl.coaches[i], cl.coaches[j] = cl.coaches[j], cl.coaches[i]
			}
		}
	}
}

// compareCoach compares two coaches based on current sort column
func (cl *CoachList) compareCoach(i, j int) bool {
	c1, c2 := cl.coaches[i], cl.coaches[j]

	var result bool
	switch cl.sortColumn {
	case 0: // First Name
		result = c1.FirstName > c2.FirstName
	case 1: // Last Name
		result = c1.LastName > c2.LastName
	case 2: // Position
		result = c1.Position > c2.Position
	case 3: // Team
		result = c1.Team > c2.Team
	case 4: // Pay Scale
		result = c1.PayScale > c2.PayScale
	default:
		return false
	}

	if !cl.sortAscending {
		result = !result
	}
	return result
}
