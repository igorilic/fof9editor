// ABOUTME: Main window implementation for FOF9 Editor
// ABOUTME: Manages the primary application window and its layout

package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/state"
	"github.com/igorilic/fof9editor/internal/version"
)

// MainWindow represents the main application window
type MainWindow struct {
	window    fyne.Window
	app       fyne.App
	content   *fyne.Container
	state     *state.AppState
	statusBar *StatusBar
}

// NewMainWindow creates a new main window
func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow(fmt.Sprintf("FOF9 Editor v%s", version.GetShortVersion()))

	mw := &MainWindow{
		window: window,
		app:    app,
		state:  state.GetInstance(),
	}

	mw.setupWindow()
	return mw
}

// setupWindow initializes the window with default settings
func (mw *MainWindow) setupWindow() {
	// Set default window size
	mw.window.Resize(fyne.NewSize(1200, 800))

	// Set minimum window size
	mw.window.SetFixedSize(false)

	// Create status bar
	mw.statusBar = NewStatusBar()

	// Create placeholder content
	mw.content = container.NewCenter(
		widget.NewLabel("FOF9 Editor - Ready to load a project"),
	)

	// Create main layout with status bar at bottom
	mainLayout := container.NewBorder(
		nil,                        // top
		mw.statusBar.GetContainer(), // bottom
		nil,                        // left
		nil,                        // right
		mw.content,                 // center
	)

	mw.window.SetContent(mainLayout)
}

// Show displays the window
func (mw *MainWindow) Show() {
	mw.window.Show()
}

// ShowAndRun displays the window and runs the application
func (mw *MainWindow) ShowAndRun() {
	mw.window.ShowAndRun()
}

// GetWindow returns the underlying Fyne window
func (mw *MainWindow) GetWindow() fyne.Window {
	return mw.window
}

// GetStatusBar returns the status bar
func (mw *MainWindow) GetStatusBar() *StatusBar {
	return mw.statusBar
}

// SetContent sets the window content
func (mw *MainWindow) SetContent(content fyne.CanvasObject) {
	mw.window.SetContent(content)
}

// UpdateTitle updates the window title
func (mw *MainWindow) UpdateTitle(projectName string) {
	if projectName == "" {
		mw.window.SetTitle(fmt.Sprintf("FOF9 Editor v%s", version.GetShortVersion()))
	} else {
		mw.window.SetTitle(fmt.Sprintf("%s - FOF9 Editor v%s", projectName, version.GetShortVersion()))
	}
}
