// ABOUTME: Tests for sidebar navigation component
// ABOUTME: Validates section selection and callback handling

package ui

import (
	"testing"
)

func TestNewSidebar(t *testing.T) {
	sb := NewSidebar(nil)

	if sb == nil {
		t.Fatal("NewSidebar returned nil")
	}

	if sb.container == nil {
		t.Fatal("Sidebar container is nil")
	}

	if sb.list == nil {
		t.Fatal("Sidebar list is nil")
	}

	if len(sb.sections) == 0 {
		t.Fatal("Sidebar sections is empty")
	}
}

func TestSidebar_GetSections(t *testing.T) {
	sb := NewSidebar(nil)

	sections := sb.GetSections()
	if len(sections) == 0 {
		t.Fatal("GetSections returned empty slice")
	}

	expectedSections := []string{"Players", "Coaches", "Teams", "League Info"}
	if len(sections) != len(expectedSections) {
		t.Fatalf("Expected %d sections, got %d", len(expectedSections), len(sections))
	}

	for i, expected := range expectedSections {
		if sections[i] != expected {
			t.Errorf("Section %d: expected '%s', got '%s'", i, expected, sections[i])
		}
	}
}

func TestSidebar_GetSelectedSection(t *testing.T) {
	sb := NewSidebar(nil)

	// Default selection should be first item
	selected := sb.GetSelectedSection()
	if selected != "Players" {
		t.Errorf("Expected default selection 'Players', got '%s'", selected)
	}
}

func TestSidebar_SetSelectedSection(t *testing.T) {
	sb := NewSidebar(nil)

	// Set to Coaches
	sb.SetSelectedSection("Coaches")

	selected := sb.GetSelectedSection()
	if selected != "Coaches" {
		t.Errorf("Expected selected section 'Coaches', got '%s'", selected)
	}

	if sb.selectedIndex != 1 {
		t.Errorf("Expected selectedIndex 1, got %d", sb.selectedIndex)
	}
}

func TestSidebar_SetSelectedSection_Invalid(t *testing.T) {
	sb := NewSidebar(nil)

	originalSelected := sb.GetSelectedSection()

	// Try to set invalid section
	sb.SetSelectedSection("NonExistent")

	// Should remain unchanged
	selected := sb.GetSelectedSection()
	if selected != originalSelected {
		t.Errorf("Selection should not change for invalid section, got '%s'", selected)
	}
}

func TestSidebar_GetContainer(t *testing.T) {
	sb := NewSidebar(nil)

	container := sb.GetContainer()
	if container == nil {
		t.Fatal("GetContainer returned nil")
	}

	if container != sb.container {
		t.Error("GetContainer did not return the correct container")
	}
}

func TestSidebar_OnSectionChange_Callback(t *testing.T) {
	callbackCalled := false
	var callbackSection string

	sb := NewSidebar(func(section string) {
		callbackCalled = true
		callbackSection = section
	})

	// Reset tracking
	callbackCalled = false
	callbackSection = ""

	// Trigger selection
	sb.SetSelectedSection("Teams")

	// Callback should have been called
	if !callbackCalled {
		t.Error("Expected callback to be called")
	}

	if callbackSection != "Teams" {
		t.Errorf("Expected callback section 'Teams', got '%s'", callbackSection)
	}
}

func TestSidebar_NilCallback(t *testing.T) {
	// Should not panic with nil callback
	sb := NewSidebar(nil)

	// Should not panic
	sb.SetSelectedSection("Coaches")

	selected := sb.GetSelectedSection()
	if selected != "Coaches" {
		t.Errorf("Expected 'Coaches', got '%s'", selected)
	}
}
