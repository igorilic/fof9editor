// ABOUTME: Tests for theme management
// ABOUTME: Validates theme switching and state management

package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestNewThemeManager(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	tm := NewThemeManager(app)

	if tm == nil {
		t.Fatal("NewThemeManager returned nil")
	}

	if tm.app == nil {
		t.Fatal("ThemeManager app is nil")
	}

	// Should default to light theme
	if tm.currentTheme != ThemeLight {
		t.Errorf("Expected default theme ThemeLight, got %d", tm.currentTheme)
	}
}

func TestThemeManager_SetTheme(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	tm := NewThemeManager(app)

	// Set to dark theme
	tm.SetTheme(ThemeDark)
	if tm.currentTheme != ThemeDark {
		t.Errorf("Expected ThemeDark, got %d", tm.currentTheme)
	}

	// Set to light theme
	tm.SetTheme(ThemeLight)
	if tm.currentTheme != ThemeLight {
		t.Errorf("Expected ThemeLight, got %d", tm.currentTheme)
	}
}

func TestThemeManager_GetCurrentTheme(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	tm := NewThemeManager(app)

	// Default should be light
	current := tm.GetCurrentTheme()
	if current != ThemeLight {
		t.Errorf("Expected ThemeLight, got %d", current)
	}

	// Change to dark
	tm.SetTheme(ThemeDark)
	current = tm.GetCurrentTheme()
	if current != ThemeDark {
		t.Errorf("Expected ThemeDark, got %d", current)
	}
}

func TestThemeManager_ToggleTheme(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	tm := NewThemeManager(app)

	// Start with light
	if tm.currentTheme != ThemeLight {
		t.Fatalf("Expected initial theme ThemeLight, got %d", tm.currentTheme)
	}

	// Toggle to dark
	tm.ToggleTheme()
	if tm.currentTheme != ThemeDark {
		t.Errorf("Expected ThemeDark after toggle, got %d", tm.currentTheme)
	}

	// Toggle back to light
	tm.ToggleTheme()
	if tm.currentTheme != ThemeLight {
		t.Errorf("Expected ThemeLight after second toggle, got %d", tm.currentTheme)
	}
}

func TestThemeManager_GetThemeName(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	tm := NewThemeManager(app)

	// Light theme
	tm.SetTheme(ThemeLight)
	name := tm.GetThemeName()
	if name != "Light" {
		t.Errorf("Expected 'Light', got '%s'", name)
	}

	// Dark theme
	tm.SetTheme(ThemeDark)
	name = tm.GetThemeName()
	if name != "Dark" {
		t.Errorf("Expected 'Dark', got '%s'", name)
	}
}

func TestThemeType_Constants(t *testing.T) {
	// Verify theme constants have expected values
	if ThemeLight != 0 {
		t.Errorf("Expected ThemeLight to be 0, got %d", ThemeLight)
	}

	if ThemeDark != 1 {
		t.Errorf("Expected ThemeDark to be 1, got %d", ThemeDark)
	}
}
