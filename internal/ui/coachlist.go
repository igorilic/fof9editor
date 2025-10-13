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
}

// NewCoachList creates a new coach list view
func NewCoachList() *CoachList {
	cl := &CoachList{
		coaches:     []models.Coach{},
		headers:     []string{"First Name", "Last Name", "Position", "Team", "Pay Scale"},
		selectedRow: -1,
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
				label.SetText(cl.headers[id.Col])
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
		if id.Row > 0 { // Skip header row
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
	cl.table.Refresh()
}
