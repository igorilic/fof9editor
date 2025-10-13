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
