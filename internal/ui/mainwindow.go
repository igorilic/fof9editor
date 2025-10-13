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
	sidebar   *Sidebar
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

	// Create placeholder content BEFORE sidebar to avoid nil pointer
	mw.content = container.NewCenter(
		widget.NewLabel("FOF9 Editor - Ready to load a project"),
	)

	// Create sidebar with callback
	mw.sidebar = NewSidebar(func(section string) {
		mw.onSectionChange(section)
	})

	// Create main layout with sidebar on left and status bar at bottom
	mainLayout := container.NewBorder(
		nil,                        // top
		mw.statusBar.GetContainer(), // bottom
		mw.sidebar.GetContainer(),  // left
		nil,                        // right
		mw.content,                 // center
	)

	mw.window.SetContent(mainLayout)
}

// onSectionChange handles section navigation changes
func (mw *MainWindow) onSectionChange(section string) {
	// Update state
	mw.state.SetCurrentSection(section)

	// Update status bar
	// For now, just show the section name
	// In later phases, this will load and display actual data
	mw.content.Objects = []fyne.CanvasObject{
		container.NewCenter(
			widget.NewLabel(fmt.Sprintf("Section: %s (data loading not yet implemented)", section)),
		),
	}
	mw.content.Refresh()
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

// GetSidebar returns the sidebar
func (mw *MainWindow) GetSidebar() *Sidebar {
	return mw.sidebar
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
