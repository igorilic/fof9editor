// ABOUTME: Tests for main window functionality
// ABOUTME: Validates window creation, configuration, and basic operations

package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestNewMainWindow(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)
	if mw == nil {
		t.Fatal("NewMainWindow returned nil")
	}

	if mw.window == nil {
		t.Fatal("MainWindow.window is nil")
	}

	if mw.app == nil {
		t.Fatal("MainWindow.app is nil")
	}

	if mw.state == nil {
		t.Fatal("MainWindow.state is nil")
	}
}

func TestMainWindow_GetWindow(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)
	window := mw.GetWindow()

	if window == nil {
		t.Fatal("GetWindow returned nil")
	}

	if window != mw.window {
		t.Error("GetWindow did not return the correct window")
	}
}

func TestMainWindow_UpdateTitle(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Test with project name
	mw.UpdateTitle("Test Project")
	title := mw.window.Title()
	if title == "" {
		t.Error("Window title should not be empty after UpdateTitle")
	}
	// Title should contain the project name
	// We can't do exact match because it includes version info

	// Test with empty project name (resets to default)
	mw.UpdateTitle("")
	title = mw.window.Title()
	if title == "" {
		t.Error("Window title should not be empty after UpdateTitle with empty string")
	}
}

func TestMainWindow_SetContent(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Get initial content
	initialContent := mw.window.Content()
	if initialContent == nil {
		t.Fatal("Initial window content is nil")
	}

	// Create new content (we can't easily verify UI structure, but we can test the API)
	// Just verify the method doesn't panic
	mw.SetContent(initialContent)
}

func TestMainWindow_DefaultSize(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Verify window has reasonable size set
	canvas := mw.window.Canvas()
	if canvas == nil {
		t.Fatal("Window canvas is nil")
	}

	// We can't easily check exact size in test environment,
	// but we can verify the window was created without panic
	if mw.window == nil {
		t.Fatal("Window should not be nil")
	}
}

func TestMainWindow_StatusBar(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Verify status bar is created
	statusBar := mw.GetStatusBar()
	if statusBar == nil {
		t.Fatal("Status bar should not be nil")
	}

	// Verify we can update status bar
	statusBar.SetProjectStatus("Test Project")
	// Just verify no panic - we can't easily check UI state in tests
}

func TestMainWindow_GetStatusBar(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	statusBar := mw.GetStatusBar()
	if statusBar == nil {
		t.Fatal("GetStatusBar returned nil")
	}

	if statusBar != mw.statusBar {
		t.Error("GetStatusBar did not return the correct status bar")
	}
}

func TestMainWindow_Sidebar(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Verify sidebar is created
	sidebar := mw.GetSidebar()
	if sidebar == nil {
		t.Fatal("Sidebar should not be nil")
	}

	// Verify default section is Players
	selected := sidebar.GetSelectedSection()
	if selected != "Players" {
		t.Errorf("Expected default section 'Players', got '%s'", selected)
	}
}

func TestMainWindow_GetSidebar(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	sidebar := mw.GetSidebar()
	if sidebar == nil {
		t.Fatal("GetSidebar returned nil")
	}

	if sidebar != mw.sidebar {
		t.Error("GetSidebar did not return the correct sidebar")
	}
}

func TestMainWindow_SectionChangeCallback(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Change section via sidebar
	mw.sidebar.SetSelectedSection("Coaches")

	// Verify state was updated
	currentSection := mw.state.GetCurrentSection()
	if currentSection != "Coaches" {
		t.Errorf("Expected state section 'Coaches', got '%s'", currentSection)
	}

	// Change to Teams
	mw.sidebar.SetSelectedSection("Teams")

	currentSection = mw.state.GetCurrentSection()
	if currentSection != "Teams" {
		t.Errorf("Expected state section 'Teams', got '%s'", currentSection)
	}
}

func TestMainWindow_UpdateContentArea(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Test updating content area for different sections
	sections := []string{"Players", "Coaches", "Teams", "League Info"}
	for _, section := range sections {
		mw.updateContentArea(section)
		// Verify content was updated (just check no panic)
		if mw.content == nil {
			t.Error("Content should not be nil after update")
		}
	}
}

func TestMainWindow_RefreshLayout(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Should not panic
	mw.RefreshLayout()

	// Verify components still accessible
	if mw.content == nil {
		t.Error("Content should not be nil after refresh")
	}
	if mw.statusBar == nil {
		t.Error("StatusBar should not be nil after refresh")
	}
	if mw.sidebar == nil {
		t.Error("Sidebar should not be nil after refresh")
	}
}

func TestMainWindow_MenuBar(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Verify menu bar is set
	mainMenu := mw.window.MainMenu()
	if mainMenu == nil {
		t.Fatal("Main menu should not be nil")
	}

	// Verify expected menus exist
	if len(mainMenu.Items) < 4 {
		t.Errorf("Expected at least 4 menus, got %d", len(mainMenu.Items))
	}

	// Verify menu names
	expectedMenus := []string{"File", "Edit", "View", "Help"}
	for i, expected := range expectedMenus {
		if i >= len(mainMenu.Items) {
			t.Errorf("Missing menu: %s", expected)
			continue
		}
		if mainMenu.Items[i].Label != expected {
			t.Errorf("Menu %d: expected '%s', got '%s'", i, expected, mainMenu.Items[i].Label)
		}
	}
}

func TestMainWindow_FileMenu(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	mainMenu := mw.window.MainMenu()
	if mainMenu == nil || len(mainMenu.Items) == 0 {
		t.Fatal("Main menu not initialized")
	}

	fileMenu := mainMenu.Items[0]
	if fileMenu.Label != "File" {
		t.Errorf("Expected 'File' menu, got '%s'", fileMenu.Label)
	}

	// Verify File menu has expected items
	if len(fileMenu.Items) < 5 {
		t.Errorf("Expected at least 5 items in File menu, got %d", len(fileMenu.Items))
	}
}

func TestMainWindow_ShowAboutDialog(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	mw := NewMainWindow(app)

	// Should not panic
	mw.showAboutDialog()
	// Can't easily verify dialog content in tests, but we can verify no crash
}
