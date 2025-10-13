// ABOUTME: Theme management for FOF9 Editor
// ABOUTME: Provides theme selection and customization

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// ThemeType represents available theme options
type ThemeType int

const (
	ThemeLight ThemeType = iota
	ThemeDark
)

// ThemeManager manages application themes
type ThemeManager struct {
	app         fyne.App
	currentTheme ThemeType
}

// NewThemeManager creates a new theme manager
func NewThemeManager(app fyne.App) *ThemeManager {
	return &ThemeManager{
		app:          app,
		currentTheme: ThemeLight,
	}
}

// SetTheme applies the specified theme
func (tm *ThemeManager) SetTheme(themeType ThemeType) {
	tm.currentTheme = themeType

	switch themeType {
	case ThemeDark:
		tm.app.Settings().SetTheme(theme.DarkTheme())
	case ThemeLight:
		tm.app.Settings().SetTheme(theme.LightTheme())
	default:
		tm.app.Settings().SetTheme(theme.LightTheme())
	}
}

// GetCurrentTheme returns the current theme type
func (tm *ThemeManager) GetCurrentTheme() ThemeType {
	return tm.currentTheme
}

// ToggleTheme switches between light and dark themes
func (tm *ThemeManager) ToggleTheme() {
	if tm.currentTheme == ThemeLight {
		tm.SetTheme(ThemeDark)
	} else {
		tm.SetTheme(ThemeLight)
	}
}

// GetThemeName returns the name of the current theme
func (tm *ThemeManager) GetThemeName() string {
	switch tm.currentTheme {
	case ThemeDark:
		return "Dark"
	case ThemeLight:
		return "Light"
	default:
		return "Light"
	}
}
