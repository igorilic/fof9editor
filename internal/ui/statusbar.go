// ABOUTME: Status bar component for FOF9 Editor
// ABOUTME: Displays project status, validation status, record count, and saved status

package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// StatusBar represents the application status bar with 4 sections
type StatusBar struct {
	container *fyne.Container

	// Status labels
	projectLabel    *widget.Label
	validationLabel *widget.Label
	recordLabel     *widget.Label
	savedLabel      *widget.Label
}

// NewStatusBar creates a new status bar
func NewStatusBar() *StatusBar {
	sb := &StatusBar{
		projectLabel:    widget.NewLabel("No project loaded"),
		validationLabel: widget.NewLabel(""),
		recordLabel:     widget.NewLabel(""),
		savedLabel:      widget.NewLabel("All changes saved"),
	}

	// Create horizontal container with 4 sections
	sb.container = container.NewBorder(
		nil, nil, nil, nil,
		container.NewHBox(
			sb.projectLabel,
			widget.NewSeparator(),
			sb.validationLabel,
			widget.NewSeparator(),
			sb.recordLabel,
			widget.NewSeparator(),
			sb.savedLabel,
		),
	)

	return sb
}

// GetContainer returns the status bar container
func (sb *StatusBar) GetContainer() *fyne.Container {
	return sb.container
}

// SetProjectStatus updates the project status section
func (sb *StatusBar) SetProjectStatus(projectName string) {
	if projectName == "" {
		sb.projectLabel.SetText("No project loaded")
	} else {
		sb.projectLabel.SetText(fmt.Sprintf("Project: %s", projectName))
	}
}

// SetValidationStatus updates the validation status section
func (sb *StatusBar) SetValidationStatus(errorCount int) {
	if errorCount == 0 {
		sb.validationLabel.SetText("✓ No validation errors")
	} else if errorCount == 1 {
		sb.validationLabel.SetText("⚠ 1 validation error")
	} else {
		sb.validationLabel.SetText(fmt.Sprintf("⚠ %d validation errors", errorCount))
	}
}

// ClearValidationStatus clears the validation status
func (sb *StatusBar) ClearValidationStatus() {
	sb.validationLabel.SetText("")
}

// SetRecordCount updates the record count section
func (sb *StatusBar) SetRecordCount(section string, count int) {
	if section == "" || count == 0 {
		sb.recordLabel.SetText("")
	} else {
		sb.recordLabel.SetText(fmt.Sprintf("%s: %d records", section, count))
	}
}

// SetSavedStatus updates the saved status section
func (sb *StatusBar) SetSavedStatus(isDirty bool) {
	if isDirty {
		sb.savedLabel.SetText("● Unsaved changes")
	} else {
		sb.savedLabel.SetText("All changes saved")
	}
}

// UpdateFromState updates all status sections based on application state
func (sb *StatusBar) UpdateFromState(projectName string, section string, recordCount int, isDirty bool, validationErrors int) {
	sb.SetProjectStatus(projectName)
	sb.SetRecordCount(section, recordCount)
	sb.SetSavedStatus(isDirty)

	if validationErrors > 0 {
		sb.SetValidationStatus(validationErrors)
	} else {
		sb.ClearValidationStatus()
	}
}
