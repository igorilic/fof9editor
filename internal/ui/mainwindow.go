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
	window       fyne.Window
	app          fyne.App
	content      *fyne.Container
	state        *state.AppState
	statusBar    *StatusBar
	sidebar      *Sidebar
	themeManager *ThemeManager
	playerList   *PlayerList
	coachList    *CoachList
	teamList     *TeamList
}

// NewMainWindow creates a new main window
func NewMainWindow(app fyne.App) *MainWindow {
	window := app.NewWindow(fmt.Sprintf("FOF9 Editor v%s", version.GetShortVersion()))

	mw := &MainWindow{
		window:       window,
		app:          app,
		state:        state.GetInstance(),
		themeManager: NewThemeManager(app),
		playerList:   NewPlayerList(),
		coachList:    NewCoachList(),
		teamList:     NewTeamList(),
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
	welcomeTitle := widget.NewLabel("FOF9 Editor")
	welcomeTitle.TextStyle = fyne.TextStyle{Bold: true}
	welcomeMessage := widget.NewLabel("Ready to load a project")
	welcomeMessage.Wrapping = fyne.TextWrapWord

	welcomeContent := container.NewVBox(
		welcomeTitle,
		widget.NewSeparator(),
		welcomeMessage,
	)

	mw.content = container.NewCenter(welcomeContent)

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

	// Setup menu bar
	mw.setupMenuBar()
}

// setupMenuBar creates and configures the application menu bar
func (mw *MainWindow) setupMenuBar() {
	// File menu
	newItem := fyne.NewMenuItem("New Project...", func() {
		// Placeholder for Phase 9 (New League Wizard)
	})
	openItem := fyne.NewMenuItem("Open Project...", func() {
		// Placeholder for Phase 8 (File Operations)
	})
	saveItem := fyne.NewMenuItem("Save", func() {
		// Placeholder for Phase 8 (File Operations)
	})
	saveAsItem := fyne.NewMenuItem("Save As...", func() {
		// Placeholder for Phase 8 (File Operations)
	})
	exitItem := fyne.NewMenuItem("Exit", func() {
		mw.app.Quit()
	})

	fileMenu := fyne.NewMenu("File", newItem, openItem, fyne.NewMenuItemSeparator(), saveItem, saveAsItem, fyne.NewMenuItemSeparator(), exitItem)

	// Edit menu
	undoItem := fyne.NewMenuItem("Undo", func() {
		// Placeholder for future implementation
	})
	undoItem.Disabled = true

	redoItem := fyne.NewMenuItem("Redo", func() {
		// Placeholder for future implementation
	})
	redoItem.Disabled = true

	editMenu := fyne.NewMenu("Edit", undoItem, redoItem)

	// View menu
	refreshItem := fyne.NewMenuItem("Refresh", func() {
		mw.RefreshLayout()
	})

	toggleThemeItem := fyne.NewMenuItem("Toggle Theme", func() {
		mw.themeManager.ToggleTheme()
	})

	viewMenu := fyne.NewMenu("View", refreshItem, fyne.NewMenuItemSeparator(), toggleThemeItem)

	// Help menu
	aboutItem := fyne.NewMenuItem("About", func() {
		mw.showAboutDialog()
	})

	helpMenu := fyne.NewMenu("Help", aboutItem)

	// Set main menu
	mainMenu := fyne.NewMainMenu(fileMenu, editMenu, viewMenu, helpMenu)
	mw.window.SetMainMenu(mainMenu)
}

// showAboutDialog displays the about dialog
func (mw *MainWindow) showAboutDialog() {
	// Import version package at the top of the file
	aboutText := fmt.Sprintf("FOF9 Editor\n\nVersion: %s\n\nA modern editor for Front Office Football Nine league files.", version.GetShortVersion())
	dialog := widget.NewLabel(aboutText)
	dialog.Wrapping = fyne.TextWrapWord

	// Create a simple about window
	aboutWindow := mw.app.NewWindow("About FOF9 Editor")
	aboutWindow.SetContent(container.NewVBox(
		dialog,
		widget.NewSeparator(),
		container.NewCenter(widget.NewButton("Close", func() {
			aboutWindow.Close()
		})),
	))
	aboutWindow.Resize(fyne.NewSize(400, 200))
	aboutWindow.Show()
}

// onSectionChange handles section navigation changes
func (mw *MainWindow) onSectionChange(section string) {
	// Update state
	mw.state.SetCurrentSection(section)

	// Update content area with section placeholder
	mw.updateContentArea(section)
}

// updateContentArea updates the main content area based on the current section
func (mw *MainWindow) updateContentArea(section string) {
	switch section {
	case "Players":
		// Load players from state and display in list
		players := mw.state.GetPlayers()
		mw.playerList.SetPlayers(players)
		mw.content.Objects = []fyne.CanvasObject{mw.playerList.GetContainer()}
		mw.statusBar.SetRecordCount("Players", len(players))

	case "Coaches":
		// Load coaches from state and display in list
		coaches := mw.state.GetCoaches()
		mw.coachList.SetCoaches(coaches)
		mw.content.Objects = []fyne.CanvasObject{mw.coachList.GetContainer()}
		mw.statusBar.SetRecordCount("Coaches", len(coaches))

	case "Teams":
		// Load teams from state and display in list
		teams := mw.state.GetTeams()
		mw.teamList.SetTeams(teams)
		mw.content.Objects = []fyne.CanvasObject{mw.teamList.GetContainer()}
		mw.statusBar.SetRecordCount("Teams", len(teams))

	default:
		// Create section-specific placeholder for other sections
		title := widget.NewLabel(fmt.Sprintf("%s", section))
		title.TextStyle = fyne.TextStyle{Bold: true}

		message := widget.NewLabel("Data loading and editing will be implemented in later steps")
		message.Wrapping = fyne.TextWrapWord

		content := container.NewVBox(
			title,
			widget.NewSeparator(),
			message,
		)

		mw.content.Objects = []fyne.CanvasObject{
			container.NewCenter(content),
		}
		mw.statusBar.SetRecordCount("", 0)
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

// GetThemeManager returns the theme manager
func (mw *MainWindow) GetThemeManager() *ThemeManager {
	return mw.themeManager
}

// RefreshLayout refreshes the entire window layout
func (mw *MainWindow) RefreshLayout() {
	if mw.content != nil {
		mw.content.Refresh()
	}
	if mw.statusBar != nil {
		mw.statusBar.GetContainer().Refresh()
	}
	if mw.sidebar != nil {
		mw.sidebar.GetContainer().Refresh()
	}
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
