// ABOUTME: Tests for status bar component
// ABOUTME: Validates status updates and display logic

package ui

import (
	"strings"
	"testing"
)

func TestNewStatusBar(t *testing.T) {
	sb := NewStatusBar()

	if sb == nil {
		t.Fatal("NewStatusBar returned nil")
	}

	if sb.container == nil {
		t.Fatal("StatusBar container is nil")
	}

	if sb.projectLabel == nil {
		t.Fatal("StatusBar projectLabel is nil")
	}

	if sb.validationLabel == nil {
		t.Fatal("StatusBar validationLabel is nil")
	}

	if sb.recordLabel == nil {
		t.Fatal("StatusBar recordLabel is nil")
	}

	if sb.savedLabel == nil {
		t.Fatal("StatusBar savedLabel is nil")
	}
}

func TestStatusBar_InitialState(t *testing.T) {
	sb := NewStatusBar()

	if sb.projectLabel.Text != "No project loaded" {
		t.Errorf("Expected initial project label 'No project loaded', got '%s'", sb.projectLabel.Text)
	}

	if sb.savedLabel.Text != "All changes saved" {
		t.Errorf("Expected initial saved label 'All changes saved', got '%s'", sb.savedLabel.Text)
	}
}

func TestStatusBar_SetProjectStatus(t *testing.T) {
	sb := NewStatusBar()

	// Set project name
	sb.SetProjectStatus("Test Project")
	if sb.projectLabel.Text != "Project: Test Project" {
		t.Errorf("Expected 'Project: Test Project', got '%s'", sb.projectLabel.Text)
	}

	// Clear project name
	sb.SetProjectStatus("")
	if sb.projectLabel.Text != "No project loaded" {
		t.Errorf("Expected 'No project loaded', got '%s'", sb.projectLabel.Text)
	}
}

func TestStatusBar_SetValidationStatus(t *testing.T) {
	sb := NewStatusBar()

	// No errors
	sb.SetValidationStatus(0)
	if !strings.Contains(sb.validationLabel.Text, "No validation errors") {
		t.Errorf("Expected 'No validation errors', got '%s'", sb.validationLabel.Text)
	}

	// One error
	sb.SetValidationStatus(1)
	if !strings.Contains(sb.validationLabel.Text, "1 validation error") {
		t.Errorf("Expected '1 validation error', got '%s'", sb.validationLabel.Text)
	}

	// Multiple errors
	sb.SetValidationStatus(5)
	if !strings.Contains(sb.validationLabel.Text, "5 validation errors") {
		t.Errorf("Expected '5 validation errors', got '%s'", sb.validationLabel.Text)
	}
}

func TestStatusBar_ClearValidationStatus(t *testing.T) {
	sb := NewStatusBar()

	sb.SetValidationStatus(3)
	if sb.validationLabel.Text == "" {
		t.Error("Validation label should not be empty after SetValidationStatus")
	}

	sb.ClearValidationStatus()
	if sb.validationLabel.Text != "" {
		t.Errorf("Expected empty validation label after clear, got '%s'", sb.validationLabel.Text)
	}
}

func TestStatusBar_SetRecordCount(t *testing.T) {
	sb := NewStatusBar()

	// Set record count
	sb.SetRecordCount("Players", 50)
	if sb.recordLabel.Text != "Players: 50 records" {
		t.Errorf("Expected 'Players: 50 records', got '%s'", sb.recordLabel.Text)
	}

	// Clear with empty section
	sb.SetRecordCount("", 50)
	if sb.recordLabel.Text != "" {
		t.Errorf("Expected empty record label with empty section, got '%s'", sb.recordLabel.Text)
	}

	// Clear with zero count
	sb.SetRecordCount("Teams", 0)
	if sb.recordLabel.Text != "" {
		t.Errorf("Expected empty record label with zero count, got '%s'", sb.recordLabel.Text)
	}
}

func TestStatusBar_SetSavedStatus(t *testing.T) {
	sb := NewStatusBar()

	// Dirty state
	sb.SetSavedStatus(true)
	if !strings.Contains(sb.savedLabel.Text, "Unsaved changes") {
		t.Errorf("Expected 'Unsaved changes', got '%s'", sb.savedLabel.Text)
	}

	// Clean state
	sb.SetSavedStatus(false)
	if !strings.Contains(sb.savedLabel.Text, "All changes saved") {
		t.Errorf("Expected 'All changes saved', got '%s'", sb.savedLabel.Text)
	}
}

func TestStatusBar_UpdateFromState(t *testing.T) {
	sb := NewStatusBar()

	// Update with full state
	sb.UpdateFromState("My Project", "Players", 100, true, 3)

	// Verify all sections updated
	if !strings.Contains(sb.projectLabel.Text, "My Project") {
		t.Errorf("Project label not updated correctly: '%s'", sb.projectLabel.Text)
	}

	if !strings.Contains(sb.recordLabel.Text, "Players") {
		t.Errorf("Record label not updated correctly: '%s'", sb.recordLabel.Text)
	}

	if !strings.Contains(sb.recordLabel.Text, "100") {
		t.Errorf("Record count not updated correctly: '%s'", sb.recordLabel.Text)
	}

	if !strings.Contains(sb.savedLabel.Text, "Unsaved") {
		t.Errorf("Saved label not updated correctly: '%s'", sb.savedLabel.Text)
	}

	if !strings.Contains(sb.validationLabel.Text, "3 validation errors") {
		t.Errorf("Validation label not updated correctly: '%s'", sb.validationLabel.Text)
	}
}

func TestStatusBar_UpdateFromState_NoErrors(t *testing.T) {
	sb := NewStatusBar()

	// Set some validation errors first
	sb.SetValidationStatus(5)

	// Update with no errors
	sb.UpdateFromState("Project", "Teams", 32, false, 0)

	// Validation label should be cleared
	if sb.validationLabel.Text != "" {
		t.Errorf("Expected empty validation label with 0 errors, got '%s'", sb.validationLabel.Text)
	}
}

func TestStatusBar_GetContainer(t *testing.T) {
	sb := NewStatusBar()

	container := sb.GetContainer()
	if container == nil {
		t.Fatal("GetContainer returned nil")
	}

	if container != sb.container {
		t.Error("GetContainer did not return the correct container")
	}
}
