// ABOUTME: Sidebar navigation component for FOF9 Editor
// ABOUTME: Provides section selection for Players, Coaches, Teams, etc.

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Sidebar represents the navigation sidebar
type Sidebar struct {
	container       *fyne.Container
	list            *widget.List
	sections        []string
	selectedIndex   int
	onSectionChange func(section string)
}

// NewSidebar creates a new sidebar with navigation options
func NewSidebar(onSectionChange func(section string)) *Sidebar {
	sb := &Sidebar{
		sections: []string{
			"Players",
			"Coaches",
			"Teams",
			"League Info",
		},
		selectedIndex:   0,
		onSectionChange: onSectionChange,
	}

	sb.setupList()
	return sb
}

// setupList creates and configures the list widget
func (sb *Sidebar) setupList() {
	sb.list = widget.NewList(
		func() int {
			return len(sb.sections)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.SetText(sb.sections[id])
		},
	)

	// Handle selection
	sb.list.OnSelected = func(id widget.ListItemID) {
		sb.selectedIndex = int(id)
		if sb.onSectionChange != nil {
			sb.onSectionChange(sb.sections[id])
		}
	}

	// Select first item by default
	sb.list.Select(0)

	// Set minimum width for sidebar (180 pixels for more content space)
	sb.list.Resize(fyne.NewSize(180, 0))

	sb.container = container.NewMax(sb.list)
}

// GetContainer returns the sidebar container
func (sb *Sidebar) GetContainer() *fyne.Container {
	return sb.container
}

// GetSelectedSection returns the currently selected section
func (sb *Sidebar) GetSelectedSection() string {
	if sb.selectedIndex >= 0 && sb.selectedIndex < len(sb.sections) {
		return sb.sections[sb.selectedIndex]
	}
	return ""
}

// SetSelectedSection sets the selected section by name
func (sb *Sidebar) SetSelectedSection(section string) {
	for i, s := range sb.sections {
		if s == section {
			sb.selectedIndex = i
			sb.list.Select(widget.ListItemID(i))
			return
		}
	}
}

// GetSections returns all available sections
func (sb *Sidebar) GetSections() []string {
	return sb.sections
}
