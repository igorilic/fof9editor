// ABOUTME: Main window implementation for FOF9 Editor
// ABOUTME: Manages the primary application window and its layout

package ui

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/data"
	"github.com/igorilic/fof9editor/internal/models"
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
	playerForm   *FormView
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
		playerForm:   NewFormView(),
	}

	mw.setupWindow()
	return mw
}

// setupWindow initializes the window with default settings
func (mw *MainWindow) setupWindow() {
	// Set default window size - larger for better content visibility
	mw.window.Resize(fyne.NewSize(1400, 900))

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

	// Use NewMax instead of NewCenter to fill available space
	mw.content = container.NewMax(container.NewCenter(welcomeContent))

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

	// Setup close intercept for unsaved changes prompt
	mw.window.SetCloseIntercept(func() {
		mw.handleWindowClose()
	})
}

// setupMenuBar creates and configures the application menu bar
func (mw *MainWindow) setupMenuBar() {
	// File menu
	newItem := fyne.NewMenuItem("New Project...", func() {
		mw.newProject()
	})
	openItem := fyne.NewMenuItem("Open Project...", func() {
		mw.openLeague()
	})

	// Load individual CSV files submenu
	loadPlayersItem := fyne.NewMenuItem("Players CSV...", func() {
		mw.loadPlayersCSV()
	})
	loadCoachesItem := fyne.NewMenuItem("Coaches CSV...", func() {
		mw.loadCoachesCSV()
	})
	loadTeamsItem := fyne.NewMenuItem("Teams CSV...", func() {
		mw.loadTeamsCSV()
	})
	loadCSVMenu := fyne.NewMenuItem("Load CSV", nil)
	loadCSVMenu.ChildMenu = fyne.NewMenu("", loadPlayersItem, loadCoachesItem, loadTeamsItem)

	saveItem := fyne.NewMenuItem("Save", func() {
		mw.saveLeague()
	})
	saveAsItem := fyne.NewMenuItem("Save As...", func() {
		mw.saveLeagueAs()
	})
	exitItem := fyne.NewMenuItem("Exit", func() {
		mw.app.Quit()
	})

	fileMenu := fyne.NewMenu("File", newItem, openItem, loadCSVMenu, fyne.NewMenuItemSeparator(), saveItem, saveAsItem, fyne.NewMenuItemSeparator(), exitItem)

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

		// Set callback to update form on row selection
		mw.playerList.SetOnSelectChange(func(index int) {
			mw.state.SetSelectedIndex(index)
			mw.updatePlayerForm()
		})

		// Create split view with list on top and form on bottom
		split := container.NewVSplit(
			mw.playerList.GetContainer(),
			mw.playerForm.GetContainer(),
		)
		split.SetOffset(0.4) // 40% list, 60% form

		// Wrap in NewMax to fill available space
		mw.content.Objects = []fyne.CanvasObject{container.NewMax(split)}
		mw.statusBar.SetRecordCount("Players", len(players))

	case "Coaches":
		// Load coaches from state and display in list
		coaches := mw.state.GetCoaches()
		mw.coachList.SetCoaches(coaches)
		// Wrap in NewMax to fill available space
		mw.content.Objects = []fyne.CanvasObject{container.NewMax(mw.coachList.GetContainer())}
		mw.statusBar.SetRecordCount("Coaches", len(coaches))

	case "Teams":
		// Load teams from state and display in list
		teams := mw.state.GetTeams()
		mw.teamList.SetTeams(teams)
		// Wrap in NewMax to fill available space
		mw.content.Objects = []fyne.CanvasObject{container.NewMax(mw.teamList.GetContainer())}
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

// updatePlayerForm updates the player form with the currently selected player
func (mw *MainWindow) updatePlayerForm() {
	// Get selected player index from state
	selectedIndex := mw.state.GetSelectedIndex()
	players := mw.state.GetPlayers()

	if selectedIndex < 0 || selectedIndex >= len(players) {
		// No valid selection - clear form
		mw.playerForm.Clear()
		return
	}

	player := players[selectedIndex]

	// Define form fields for player (key fields only for now)
	fields := []FieldDef{
		{Name: "firstName", Label: "First Name", Type: FieldTypeText, Value: player.FirstName},
		{Name: "lastName", Label: "Last Name", Type: FieldTypeText, Value: player.LastName},
		{Name: "team", Label: "Team", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Team)},
		{Name: "position", Label: "Position", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.PositionKey)},
		{Name: "uniform", Label: "Uniform", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Uniform)},
		{Name: "overall", Label: "Overall Rating", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.OverallRating)},
	}

	// Set fields in form
	mw.playerForm.SetFields(fields)

	// Add buttons if not already present
	if mw.playerForm.buttonBar == nil {
		mw.playerForm.AddButtons()
	}

	// Wire callbacks
	mw.playerForm.SetCallbacks(
		func() { // onSave
			mw.savePlayerForm()
		},
		func() { // onDelete
			mw.deletePlayer()
		},
		func() { // onNext
			mw.navigatePlayer(1)
		},
		func() { // onPrev
			mw.navigatePlayer(-1)
		},
	)

	// Form updates in-place within the split view - no need to refresh content
}

// savePlayerForm saves changes from the player form
func (mw *MainWindow) savePlayerForm() {
	selectedIndex := mw.state.GetSelectedIndex()
	players := mw.state.GetPlayers()

	if selectedIndex < 0 || selectedIndex >= len(players) {
		return
	}

	// Get field values
	players[selectedIndex].FirstName = mw.playerForm.GetFieldValue("firstName")
	players[selectedIndex].LastName = mw.playerForm.GetFieldValue("lastName")

	// Parse numeric fields
	if team := mw.playerForm.GetFieldValue("team"); team != "" {
		if val, err := fmt.Sscan(team, &players[selectedIndex].Team); err == nil && val > 0 {
			// Successfully parsed
		}
	}
	if pos := mw.playerForm.GetFieldValue("position"); pos != "" {
		if val, err := fmt.Sscan(pos, &players[selectedIndex].PositionKey); err == nil && val > 0 {
			// Successfully parsed
		}
	}
	if uniform := mw.playerForm.GetFieldValue("uniform"); uniform != "" {
		if val, err := fmt.Sscan(uniform, &players[selectedIndex].Uniform); err == nil && val > 0 {
			// Successfully parsed
		}
	}
	if overall := mw.playerForm.GetFieldValue("overall"); overall != "" {
		if val, err := fmt.Sscan(overall, &players[selectedIndex].OverallRating); err == nil && val > 0 {
			// Successfully parsed
		}
	}

	// Mark as modified
	mw.state.MarkDirty()
	mw.statusBar.SetSavedStatus(true)

	// Refresh player list
	mw.playerList.SetPlayers(players)
}

// deletePlayer removes the currently selected player
func (mw *MainWindow) deletePlayer() {
	selectedIndex := mw.state.GetSelectedIndex()
	players := mw.state.GetPlayers()

	if selectedIndex < 0 || selectedIndex >= len(players) {
		return
	}

	// Remove player
	players = append(players[:selectedIndex], players[selectedIndex+1:]...)
	mw.state.SetPlayers(players)

	// Mark as modified
	mw.state.MarkDirty()
	mw.statusBar.SetSavedStatus(true)

	// Refresh list and go back to list view
	mw.playerList.SetPlayers(players)
	mw.content.Objects = []fyne.CanvasObject{container.NewMax(mw.playerList.GetContainer())}
	mw.content.Refresh()
	mw.statusBar.SetRecordCount("Players", len(players))
}

// navigatePlayer moves to the next or previous player
func (mw *MainWindow) navigatePlayer(delta int) {
	selectedIndex := mw.state.GetSelectedIndex()
	players := mw.state.GetPlayers()

	newIndex := selectedIndex + delta
	if newIndex < 0 {
		newIndex = 0
	}
	if newIndex >= len(players) {
		newIndex = len(players) - 1
	}

	mw.state.SetSelectedIndex(newIndex)
	mw.updatePlayerForm()
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

// openLeague opens an existing league project
func (mw *MainWindow) openLeague() {
	// Check for unsaved changes
	if mw.state.IsDirtyState() {
		dialog.ShowConfirm("Unsaved Changes", "You have unsaved changes. Do you want to save before opening another project?",
			func(save bool) {
				if save {
					mw.saveLeague()
				}
				mw.showOpenDialog()
			}, mw.window)
		return
	}

	mw.showOpenDialog()
}

// showOpenDialog shows the file open dialog
func (mw *MainWindow) showOpenDialog() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if reader == nil {
			// User cancelled
			return
		}
		defer reader.Close()

		filePath := reader.URI().Path()

		// Load project
		if err := mw.state.LoadProject(filePath); err != nil {
			dialog.ShowError(fmt.Errorf("failed to open project: %w", err), mw.window)
			return
		}

		// Update UI
		project := mw.state.GetProject()
		if project != nil {
			mw.UpdateTitle(project.LeagueName)
			mw.statusBar.SetProjectStatus(project.LeagueName)
		}

		// Navigate to Players section
		mw.sidebar.SetSelectedSection("Players")
		mw.updateContentArea("Players")

		// Show success message
		dialog.ShowInformation("Success", fmt.Sprintf("Opened league: %s", project.LeagueName), mw.window)
	}, mw.window)
}

// saveLeague saves the current league project
func (mw *MainWindow) saveLeague() {
	if !mw.state.HasProject() {
		dialog.ShowInformation("No Project", "No project is currently loaded.", mw.window)
		return
	}

	// Save project
	if err := mw.state.SaveProject(); err != nil {
		dialog.ShowError(fmt.Errorf("failed to save project: %w", err), mw.window)
		return
	}

	// Update status bar
	mw.statusBar.SetSavedStatus(false)

	// Update window title
	project := mw.state.GetProject()
	if project != nil {
		mw.UpdateTitle(project.LeagueName)
	}

	// Show success message
	dialog.ShowInformation("Success", "Project saved successfully.", mw.window)
}

// saveLeagueAs saves the league project to a new location
func (mw *MainWindow) saveLeagueAs() {
	if !mw.state.HasProject() {
		dialog.ShowInformation("No Project", "No project is currently loaded.", mw.window)
		return
	}

	project := mw.state.GetProject()
	if project == nil {
		return
	}

	// Show save dialog
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if writer == nil {
			// User cancelled
			return
		}
		defer writer.Close()

		newPath := writer.URI().Path()

		// Ensure .fof9proj extension
		if filepath.Ext(newPath) != ".fof9proj" {
			newPath += ".fof9proj"
		}

		// Update project path in state
		mw.state.ProjectPath = newPath

		// Save to new location
		if err := mw.state.SaveProject(); err != nil {
			dialog.ShowError(fmt.Errorf("failed to save project: %w", err), mw.window)
			return
		}

		// Update UI
		mw.UpdateTitle(project.LeagueName)
		mw.statusBar.SetSavedStatus(false)

		// Show success message
		dialog.ShowInformation("Success", fmt.Sprintf("Project saved as: %s", filepath.Base(newPath)), mw.window)
	}, mw.window)
}

// handleWindowClose handles the window close event, prompting for unsaved changes
func (mw *MainWindow) handleWindowClose() {
	// Check for unsaved changes
	if mw.state.IsDirtyState() {
		dialog.ShowConfirm("Unsaved Changes",
			"You have unsaved changes. Do you want to save before closing?",
			func(save bool) {
				if save {
					// Save project
					if err := mw.state.SaveProject(); err != nil {
						dialog.ShowError(fmt.Errorf("failed to save project: %w", err), mw.window)
						return
					}
				}
				// Close window
				mw.window.Close()
			}, mw.window)
		return
	}

	// No unsaved changes, close immediately
	mw.window.Close()
}

// loadPlayersCSV loads players from a CSV file
func (mw *MainWindow) loadPlayersCSV() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if reader == nil {
			return
		}
		defer reader.Close()

		filePath := reader.URI().Path()

		// Load players from CSV
		players, err := data.LoadPlayers(filePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to load players: %w", err), mw.window)
			return
		}

		// Update state
		mw.state.SetPlayers(players)

		// Update UI
		mw.sidebar.SetSelectedSection("Players")
		mw.updateContentArea("Players")
		mw.statusBar.SetProjectStatus("Players CSV Loaded")

		dialog.ShowInformation("Success", fmt.Sprintf("Loaded %d players", len(players)), mw.window)
	}, mw.window)
}

// loadCoachesCSV loads coaches from a CSV file
func (mw *MainWindow) loadCoachesCSV() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if reader == nil {
			return
		}
		defer reader.Close()

		filePath := reader.URI().Path()

		// Load coaches from CSV
		coaches, err := data.LoadCoaches(filePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to load coaches: %w", err), mw.window)
			return
		}

		// Update state
		mw.state.SetCoaches(coaches)

		// Update UI
		mw.sidebar.SetSelectedSection("Coaches")
		mw.updateContentArea("Coaches")
		mw.statusBar.SetProjectStatus("Coaches CSV Loaded")

		dialog.ShowInformation("Success", fmt.Sprintf("Loaded %d coaches", len(coaches)), mw.window)
	}, mw.window)
}

// loadTeamsCSV loads teams from a CSV file
func (mw *MainWindow) loadTeamsCSV() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if reader == nil {
			return
		}
		defer reader.Close()

		filePath := reader.URI().Path()

		// Load teams from CSV
		teams, err := data.LoadTeams(filePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to load teams: %w", err), mw.window)
			return
		}

		// Update state
		mw.state.SetTeams(teams)

		// Update UI
		mw.sidebar.SetSelectedSection("Teams")
		mw.updateContentArea("Teams")
		mw.statusBar.SetProjectStatus("Teams CSV Loaded")

		dialog.ShowInformation("Success", fmt.Sprintf("Loaded %d teams", len(teams)), mw.window)
	}, mw.window)
}

// newProject creates a new project
func (mw *MainWindow) newProject() {
	// Check for unsaved changes
	if mw.state.IsDirtyState() {
		dialog.ShowConfirm("Unsaved Changes", "You have unsaved changes. Do you want to save before creating a new project?",
			func(save bool) {
				if save {
					mw.saveLeague()
				}
				mw.showNewProjectDialog()
			}, mw.window)
		return
	}

	mw.showNewProjectDialog()
}

// showNewProjectDialog shows the new project creation dialog
func (mw *MainWindow) showNewProjectDialog() {
	// Create form for new project
	leagueNameEntry := widget.NewEntry()
	leagueNameEntry.SetPlaceHolder("My League")

	baseYearEntry := widget.NewEntry()
	baseYearEntry.SetPlaceHolder("2024")
	baseYearEntry.SetText("2024")

	content := container.NewVBox(
		widget.NewLabel("Create New Project"),
		widget.NewSeparator(),
		widget.NewLabel("League Name:"),
		leagueNameEntry,
		widget.NewLabel("Base Year:"),
		baseYearEntry,
	)

	// Create dialog
	d := dialog.NewCustom("New Project", "Cancel", content, mw.window)

	// Add Create button
	createButton := widget.NewButton("Create", func() {
		leagueName := leagueNameEntry.Text
		if leagueName == "" {
			dialog.ShowError(fmt.Errorf("league name is required"), mw.window)
			return
		}

		baseYear := 2024
		if baseYearEntry.Text != "" {
			if _, err := fmt.Sscanf(baseYearEntry.Text, "%d", &baseYear); err != nil {
				dialog.ShowError(fmt.Errorf("invalid base year"), mw.window)
				return
			}
		}

		d.Hide()

		// Show file save dialog for project location
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, mw.window)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			projectPath := writer.URI().Path()
			if filepath.Ext(projectPath) != ".fof9proj" {
				projectPath += ".fof9proj"
			}

			// Create new project
			project := models.NewProject(leagueName, leagueName, filepath.Dir(projectPath), baseYear)

			// Save project
			mw.state.SetProject(project)
			mw.state.ProjectPath = projectPath

			if err := mw.state.SaveProject(); err != nil {
				dialog.ShowError(fmt.Errorf("failed to create project: %w", err), mw.window)
				return
			}

			// Update UI
			mw.UpdateTitle(leagueName)
			mw.statusBar.SetProjectStatus(leagueName)
			mw.state.MarkClean()

			dialog.ShowInformation("Success", fmt.Sprintf("Created project: %s", leagueName), mw.window)
		}, mw.window)
	})
	createButton.Importance = widget.HighImportance

	// Add button to dialog
	content.Add(widget.NewSeparator())
	content.Add(container.NewHBox(
		widget.NewButton("Cancel", func() { d.Hide() }),
		createButton,
	))

	d.Show()
}
