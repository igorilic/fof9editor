// ABOUTME: Main window implementation for FOF9 Editor
// ABOUTME: Manages the primary application window and its layout

package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/igorilic/fof9editor/internal/data"
	"github.com/igorilic/fof9editor/internal/models"
	"github.com/igorilic/fof9editor/internal/state"
	"github.com/igorilic/fof9editor/internal/validation"
	"github.com/igorilic/fof9editor/internal/version"
)

// getDefaultCSVPath returns the default folder location for CSV file dialogs
// Uses the FOF9 leagues folder if it exists, otherwise falls back to home directory
func getDefaultCSVPath() fyne.ListableURI {
	// Default FOF9 leagues folder path
	defaultPath := filepath.Join("C:", "Program Files (x86)", "Steam", "steamapps", "common", "Front Office Football Nine", "leagues")

	// Check if the path exists
	if _, err := os.Stat(defaultPath); err == nil {
		// Path exists, use it
		if uri := storage.NewFileURI(defaultPath); uri != nil {
			if listable, ok := uri.(fyne.ListableURI); ok {
				return listable
			}
		}
	}

	// Fallback to user's home directory
	homeDir, err := os.UserHomeDir()
	if err == nil {
		if uri := storage.NewFileURI(homeDir); uri != nil {
			if listable, ok := uri.(fyne.ListableURI); ok {
				return listable
			}
		}
	}

	// Last resort: return nil and let Fyne use its default
	return nil
}

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
	coachForm    *FormView
	teamForm     *FormView
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
		coachForm:    NewFormView(),
		teamForm:     NewFormView(),
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
	// File menu - Load CSV files
	loadPlayersItem := fyne.NewMenuItem("Load Players...", func() {
		mw.loadPlayersCSV()
	})
	loadCoachesItem := fyne.NewMenuItem("Load Coaches...", func() {
		mw.loadCoachesCSV()
	})
	loadTeamsItem := fyne.NewMenuItem("Load Teams...", func() {
		mw.loadTeamsCSV()
	})

	// Save CSV files
	savePlayersItem := fyne.NewMenuItem("Save Players...", func() {
		mw.savePlayersCSV()
	})
	saveCoachesItem := fyne.NewMenuItem("Save Coaches...", func() {
		mw.saveCoachesCSV()
	})
	saveTeamsItem := fyne.NewMenuItem("Save Teams...", func() {
		mw.saveTeamsCSV()
	})

	exitItem := fyne.NewMenuItem("Exit", func() {
		mw.app.Quit()
	})

	fileMenu := fyne.NewMenu("File",
		loadPlayersItem, loadCoachesItem, loadTeamsItem,
		fyne.NewMenuItemSeparator(),
		savePlayersItem, saveCoachesItem, saveTeamsItem,
		fyne.NewMenuItemSeparator(),
		exitItem)

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
		// Load coaches from state and display in split view
		coaches := mw.state.GetCoaches()
		mw.coachList.SetCoaches(coaches)

		// Set up coach selection callback
		mw.coachList.SetOnSelectChange(func(index int) {
			mw.state.SetSelectedIndex(index)
			mw.updateCoachForm()
		})

		// Create split view: 40% list, 60% form
		split := container.NewHSplit(
			mw.coachList.GetContainer(),
			mw.coachForm.GetContainer(),
		)
		split.SetOffset(0.4)

		// Wrap in NewMax to fill available space
		mw.content.Objects = []fyne.CanvasObject{container.NewMax(split)}
		mw.statusBar.SetRecordCount("Coaches", len(coaches))

	case "Teams":
		// Load teams from state and display in split view
		teams := mw.state.GetTeams()
		mw.teamList.SetTeams(teams)

		// Set up team selection callback
		mw.teamList.SetOnSelectChange(func(index int) {
			mw.state.SetSelectedIndex(index)
			mw.updateTeamForm()
		})

		// Create split view: 40% list, 60% form
		split := container.NewHSplit(
			mw.teamList.GetContainer(),
			mw.teamForm.GetContainer(),
		)
		split.SetOffset(0.4)

		// Wrap in NewMax to fill available space
		mw.content.Objects = []fyne.CanvasObject{container.NewMax(split)}
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

	// Get reference data for dropdowns
	refData := mw.state.ReferenceData
	teamOptions := refData.GetTeamOptions()
	positionOptions := refData.GetPositionOptions()

	// Get current team name and position name for the player
	currentTeamName := refData.GetTeamNameByID(player.Team)
	currentPositionName := models.GetPositionName(player.PositionKey)

	// Determine team field type based on whether teams are loaded
	var teamField FieldDef
	if len(teamOptions) > 0 {
		// Teams loaded - use dropdown
		teamField = FieldDef{Name: "team", Label: "Team", Type: FieldTypeSelect, Value: currentTeamName, Options: teamOptions}
	} else {
		// No teams loaded - show ID as text field
		teamField = FieldDef{Name: "team", Label: "Team (ID)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Team)}
	}

	// Define form fields for player - expanded with more useful fields
	fields := []FieldDef{
		// Basic Info
		{Name: "firstName", Label: "First Name", Type: FieldTypeText, Value: player.FirstName},
		{Name: "lastName", Label: "Last Name", Type: FieldTypeText, Value: player.LastName},
		teamField,
		{Name: "position", Label: "Position", Type: FieldTypeSelect, Value: currentPositionName, Options: positionOptions},
		{Name: "uniform", Label: "Uniform #", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Uniform)},
		{Name: "overall", Label: "Overall Rating", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.OverallRating)},

		// Physical Attributes
		{Name: "height", Label: "Height (inches)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Height)},
		{Name: "weight", Label: "Weight (lbs)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Weight)},
		{Name: "handSize", Label: "Hand Size", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.HandSize)},
		{Name: "armLength", Label: "Arm Length", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.ArmLength)},

		// Career Info
		{Name: "experience", Label: "Experience (years)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.Experience)},
		{Name: "college", Label: "College", Type: FieldTypeText, Value: player.College},
		{Name: "yearEntry", Label: "Year Entered League", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.YearEntry)},
		{Name: "roundDrafted", Label: "Draft Round", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.RoundDrafted)},
		{Name: "selectionDrafted", Label: "Draft Pick", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", player.SelectionDrafted)},
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

	// Clear previous validation errors
	mw.playerForm.ClearAllErrors()

	// Helper function to parse integer field
	parseIntField := func(fieldName string, target *int) {
		if value := mw.playerForm.GetFieldValue(fieldName); value != "" {
			if parsed, err := strconv.Atoi(value); err == nil {
				*target = parsed
			}
		}
	}

	// Get text field values
	players[selectedIndex].FirstName = mw.playerForm.GetFieldValue("firstName")
	players[selectedIndex].LastName = mw.playerForm.GetFieldValue("lastName")
	players[selectedIndex].College = mw.playerForm.GetFieldValue("college")

	// Handle team field (could be dropdown or number field)
	teamValue := mw.playerForm.GetFieldValue("team")
	if teamValue != "" {
		refData := mw.state.ReferenceData
		// Try to convert from team name first (if teams are loaded)
		if len(refData.Teams) > 0 {
			teamID := refData.GetTeamIDByName(teamValue)
			if teamID >= 0 {
				players[selectedIndex].Team = teamID
			}
		} else {
			// Teams not loaded - parse as integer ID
			if parsed, err := strconv.Atoi(teamValue); err == nil {
				players[selectedIndex].Team = parsed
			}
		}
	}

	// Convert position name to ID
	positionName := mw.playerForm.GetFieldValue("position")
	if positionName != "" {
		refData := mw.state.ReferenceData
		positionID := refData.GetPositionIDByName(positionName)
		if positionID >= 0 {
			players[selectedIndex].PositionKey = positionID
		}
	}

	// Parse all numeric fields
	parseIntField("uniform", &players[selectedIndex].Uniform)
	parseIntField("overall", &players[selectedIndex].OverallRating)
	parseIntField("height", &players[selectedIndex].Height)
	parseIntField("weight", &players[selectedIndex].Weight)
	parseIntField("handSize", &players[selectedIndex].HandSize)
	parseIntField("armLength", &players[selectedIndex].ArmLength)
	parseIntField("experience", &players[selectedIndex].Experience)
	parseIntField("yearEntry", &players[selectedIndex].YearEntry)
	parseIntField("roundDrafted", &players[selectedIndex].RoundDrafted)
	parseIntField("selectionDrafted", &players[selectedIndex].SelectionDrafted)

	// Validate player data
	validationResult := validation.ValidatePlayer(&players[selectedIndex])
	if !validationResult.Valid {
		// Display validation errors
		for _, err := range validationResult.Errors {
			// Map field names to form field names (convert to camelCase)
			formFieldName := fieldNameToFormField(err.Field)
			mw.playerForm.SetFieldError(formFieldName, err.Message)
		}
		return // Don't save if validation fails
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

// updateCoachForm populates the coach form with the currently selected coach's data
func (mw *MainWindow) updateCoachForm() {
	selectedIndex := mw.state.GetSelectedIndex()
	coaches := mw.state.GetCoaches()

	if selectedIndex < 0 || selectedIndex >= len(coaches) {
		mw.coachForm.Clear()
		return
	}

	coach := coaches[selectedIndex]

	// Get reference data for dropdowns
	refData := mw.state.ReferenceData
	teamOptions := refData.GetTeamOptions()
	coachPositionOptions := refData.GetCoachPositionOptions()

	// Get current team name and position name for the coach
	currentTeamName := refData.GetTeamNameByID(coach.Team)
	currentPositionName := refData.GetCoachPositionNameByID(coach.Position)

	// Determine team field type based on whether teams are loaded
	var teamField FieldDef
	if len(teamOptions) > 0 {
		// Teams loaded - use dropdown
		teamField = FieldDef{Name: "team", Label: "Team", Type: FieldTypeSelect, Value: currentTeamName, Options: teamOptions}
	} else {
		// No teams loaded - show ID as text field
		teamField = FieldDef{Name: "team", Label: "Team (ID)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.Team)}
	}

	// Define form fields for coach - comprehensive set
	fields := []FieldDef{
		// Basic Info
		{Name: "firstName", Label: "First Name", Type: FieldTypeText, Value: coach.FirstName},
		{Name: "lastName", Label: "Last Name", Type: FieldTypeText, Value: coach.LastName},
		teamField,
		{Name: "position", Label: "Position", Type: FieldTypeSelect, Value: currentPositionName, Options: coachPositionOptions},
		{Name: "positionGroup", Label: "Position Group", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.PositionGroup)},

		// Birth Info
		{Name: "birthMonth", Label: "Birth Month", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.BirthMonth)},
		{Name: "birthDay", Label: "Birth Day", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.BirthDay)},
		{Name: "birthYear", Label: "Birth Year", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.BirthYear)},
		{Name: "birthCity", Label: "Birth City", Type: FieldTypeText, Value: coach.BirthCity},
		{Name: "birthCityID", Label: "Birth City ID", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.BirthCityID)},

		// College
		{Name: "college", Label: "College", Type: FieldTypeText, Value: coach.College},
		{Name: "collegeID", Label: "College ID", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.CollegeID)},

		// Coaching Styles
		{Name: "offensiveStyle", Label: "Offensive Style (0-6)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.OffensiveStyle)},
		{Name: "defensiveStyle", Label: "Defensive Style (0-4)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.DefensiveStyle)},

		// Compensation
		{Name: "payScale", Label: "Pay Scale (x$10K)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", coach.PayScale)},
	}

	mw.coachForm.SetFields(fields)

	// Wire callbacks
	mw.coachForm.SetCallbacks(
		func() { // onSave
			mw.saveCoachForm()
		},
		func() { // onDelete
			mw.deleteCoach()
		},
		func() { // onNext
			mw.navigateCoach(1)
		},
		func() { // onPrev
			mw.navigateCoach(-1)
		},
	)

	// Add buttons for save, delete, prev/next
	mw.coachForm.AddButtons()
}

// saveCoachForm saves changes from the form back to the coach data
func (mw *MainWindow) saveCoachForm() {
	selectedIndex := mw.state.GetSelectedIndex()
	coaches := mw.state.GetCoaches()

	if selectedIndex < 0 || selectedIndex >= len(coaches) {
		return
	}

	// Clear previous validation errors
	mw.coachForm.ClearAllErrors()

	// Helper function to parse integer field
	parseIntField := func(fieldName string, target *int) {
		if value := mw.coachForm.GetFieldValue(fieldName); value != "" {
			if parsed, err := strconv.Atoi(value); err == nil {
				*target = parsed
			}
		}
	}

	// Get text field values
	coaches[selectedIndex].FirstName = mw.coachForm.GetFieldValue("firstName")
	coaches[selectedIndex].LastName = mw.coachForm.GetFieldValue("lastName")
	coaches[selectedIndex].BirthCity = mw.coachForm.GetFieldValue("birthCity")
	coaches[selectedIndex].College = mw.coachForm.GetFieldValue("college")

	// Handle team field (could be dropdown or number field)
	teamValue := mw.coachForm.GetFieldValue("team")
	if teamValue != "" {
		refData := mw.state.ReferenceData
		// Try to convert from team name first (if teams are loaded)
		if len(refData.Teams) > 0 {
			teamID := refData.GetTeamIDByName(teamValue)
			if teamID >= 0 {
				coaches[selectedIndex].Team = teamID
			}
		} else {
			// Teams not loaded - parse as integer ID
			if parsed, err := strconv.Atoi(teamValue); err == nil {
				coaches[selectedIndex].Team = parsed
			}
		}
	}

	// Convert position name to ID
	positionName := mw.coachForm.GetFieldValue("position")
	if positionName != "" {
		refData := mw.state.ReferenceData
		positionID := refData.GetCoachPositionIDByName(positionName)
		if positionID >= 0 {
			coaches[selectedIndex].Position = positionID
		}
	}

	// Parse all numeric fields
	parseIntField("positionGroup", &coaches[selectedIndex].PositionGroup)
	parseIntField("birthMonth", &coaches[selectedIndex].BirthMonth)
	parseIntField("birthDay", &coaches[selectedIndex].BirthDay)
	parseIntField("birthYear", &coaches[selectedIndex].BirthYear)
	parseIntField("birthCityID", &coaches[selectedIndex].BirthCityID)
	parseIntField("collegeID", &coaches[selectedIndex].CollegeID)
	parseIntField("offensiveStyle", &coaches[selectedIndex].OffensiveStyle)
	parseIntField("defensiveStyle", &coaches[selectedIndex].DefensiveStyle)
	parseIntField("payScale", &coaches[selectedIndex].PayScale)

	// Validate coach data
	validationResult := validation.ValidateCoach(&coaches[selectedIndex])
	if !validationResult.Valid {
		// Display validation errors
		for _, err := range validationResult.Errors {
			formFieldName := fieldNameToFormField(err.Field)
			mw.coachForm.SetFieldError(formFieldName, err.Message)
		}
		return // Don't save if validation fails
	}

	// Mark as modified
	mw.state.MarkDirty()
	mw.statusBar.SetSavedStatus(true)

	// Refresh coach list
	mw.coachList.SetCoaches(coaches)
}

// deleteCoach removes the currently selected coach
func (mw *MainWindow) deleteCoach() {
	selectedIndex := mw.state.GetSelectedIndex()
	coaches := mw.state.GetCoaches()

	if selectedIndex < 0 || selectedIndex >= len(coaches) {
		return
	}

	// Remove coach
	coaches = append(coaches[:selectedIndex], coaches[selectedIndex+1:]...)
	mw.state.SetCoaches(coaches)

	// Mark as modified
	mw.state.MarkDirty()
	mw.statusBar.SetSavedStatus(true)

	// Refresh list and go back to list view
	mw.coachList.SetCoaches(coaches)
	mw.content.Objects = []fyne.CanvasObject{container.NewMax(mw.coachList.GetContainer())}
	mw.content.Refresh()
	mw.statusBar.SetRecordCount("Coaches", len(coaches))
}

// navigateCoach moves to the next or previous coach
func (mw *MainWindow) navigateCoach(delta int) {
	selectedIndex := mw.state.GetSelectedIndex()
	coaches := mw.state.GetCoaches()

	newIndex := selectedIndex + delta
	if newIndex < 0 {
		newIndex = 0
	}
	if newIndex >= len(coaches) {
		newIndex = len(coaches) - 1
	}

	mw.state.SetSelectedIndex(newIndex)
	mw.updateCoachForm()
}

// updateTeamForm populates the team form with the currently selected team's data
func (mw *MainWindow) updateTeamForm() {
	selectedIndex := mw.state.GetSelectedIndex()
	teams := mw.state.GetTeams()

	if selectedIndex < 0 || selectedIndex >= len(teams) {
		mw.teamForm.Clear()
		return
	}

	team := teams[selectedIndex]

	// Define form fields for team - most useful editing fields
	fields := []FieldDef{
		// Team Identity
		{Name: "teamName", Label: "Team Name", Type: FieldTypeText, Value: team.TeamName},
		{Name: "nickName", Label: "Nickname", Type: FieldTypeText, Value: team.NickName},
		{Name: "abbreviation", Label: "Abbreviation", Type: FieldTypeText, Value: team.Abbreviation},
		{Name: "year", Label: "Year", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Year)},
		{Name: "teamID", Label: "Team ID", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.TeamID)},

		// League Structure
		{Name: "conference", Label: "Conference", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Conference)},
		{Name: "division", Label: "Division", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Division)},
		{Name: "city", Label: "City ID", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.City)},

		// Team Colors
		{Name: "primaryRed", Label: "Primary Red (0-255)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.PrimaryRed)},
		{Name: "primaryGreen", Label: "Primary Green (0-255)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.PrimaryGreen)},
		{Name: "primaryBlue", Label: "Primary Blue (0-255)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.PrimaryBlue)},
		{Name: "secondaryRed", Label: "Secondary Red (0-255)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.SecondaryRed)},
		{Name: "secondaryGreen", Label: "Secondary Green (0-255)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.SecondaryGreen)},
		{Name: "secondaryBlue", Label: "Secondary Blue (0-255)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.SecondaryBlue)},

		// Stadium Info
		{Name: "roof", Label: "Roof Type (0=outdoor/1=dome/2=retract)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Roof)},
		{Name: "turf", Label: "Turf Type (0=grass/1=artificial/2=hybrid)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Turf)},
		{Name: "built", Label: "Year Built", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Built)},
		{Name: "capacity", Label: "Stadium Capacity", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Capacity)},
		{Name: "luxury", Label: "Luxury Boxes", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Luxury)},
		{Name: "condition", Label: "Condition (1-10)", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Condition)},

		// Financial Data
		{Name: "attendance", Label: "Average Attendance", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Attendance)},
		{Name: "support", Label: "Fan Support", Type: FieldTypeNumber, Value: fmt.Sprintf("%d", team.Support)},
	}

	mw.teamForm.SetFields(fields)

	// Wire callbacks
	mw.teamForm.SetCallbacks(
		func() { // onSave
			mw.saveTeamForm()
		},
		func() { // onDelete
			mw.deleteTeam()
		},
		func() { // onNext
			mw.navigateTeam(1)
		},
		func() { // onPrev
			mw.navigateTeam(-1)
		},
	)

	// Add buttons for save, delete, prev/next
	mw.teamForm.AddButtons()
}

// saveTeamForm saves changes from the form back to the team data
func (mw *MainWindow) saveTeamForm() {
	selectedIndex := mw.state.GetSelectedIndex()
	teams := mw.state.GetTeams()

	if selectedIndex < 0 || selectedIndex >= len(teams) {
		return
	}

	// Clear previous validation errors
	mw.teamForm.ClearAllErrors()

	// Helper function to parse integer field
	parseIntField := func(fieldName string, target *int) {
		if value := mw.teamForm.GetFieldValue(fieldName); value != "" {
			if parsed, err := strconv.Atoi(value); err == nil {
				*target = parsed
			}
		}
	}

	// Get text field values
	teams[selectedIndex].TeamName = mw.teamForm.GetFieldValue("teamName")
	teams[selectedIndex].NickName = mw.teamForm.GetFieldValue("nickName")
	teams[selectedIndex].Abbreviation = mw.teamForm.GetFieldValue("abbreviation")

	// Parse all numeric fields
	parseIntField("year", &teams[selectedIndex].Year)
	parseIntField("teamID", &teams[selectedIndex].TeamID)
	parseIntField("conference", &teams[selectedIndex].Conference)
	parseIntField("division", &teams[selectedIndex].Division)
	parseIntField("city", &teams[selectedIndex].City)
	parseIntField("primaryRed", &teams[selectedIndex].PrimaryRed)
	parseIntField("primaryGreen", &teams[selectedIndex].PrimaryGreen)
	parseIntField("primaryBlue", &teams[selectedIndex].PrimaryBlue)
	parseIntField("secondaryRed", &teams[selectedIndex].SecondaryRed)
	parseIntField("secondaryGreen", &teams[selectedIndex].SecondaryGreen)
	parseIntField("secondaryBlue", &teams[selectedIndex].SecondaryBlue)
	parseIntField("roof", &teams[selectedIndex].Roof)
	parseIntField("turf", &teams[selectedIndex].Turf)
	parseIntField("built", &teams[selectedIndex].Built)
	parseIntField("capacity", &teams[selectedIndex].Capacity)
	parseIntField("luxury", &teams[selectedIndex].Luxury)
	parseIntField("condition", &teams[selectedIndex].Condition)
	parseIntField("attendance", &teams[selectedIndex].Attendance)
	parseIntField("support", &teams[selectedIndex].Support)

	// Validate team data
	validationResult := validation.ValidateTeam(&teams[selectedIndex])
	if !validationResult.Valid {
		// Display validation errors
		for _, err := range validationResult.Errors {
			formFieldName := fieldNameToFormField(err.Field)
			mw.teamForm.SetFieldError(formFieldName, err.Message)
		}
		return // Don't save if validation fails
	}

	// Mark as modified
	mw.state.MarkDirty()
	mw.statusBar.SetSavedStatus(true)

	// Refresh team list
	mw.teamList.SetTeams(teams)
}

// deleteTeam removes the currently selected team
func (mw *MainWindow) deleteTeam() {
	selectedIndex := mw.state.GetSelectedIndex()
	teams := mw.state.GetTeams()

	if selectedIndex < 0 || selectedIndex >= len(teams) {
		return
	}

	// Remove team
	teams = append(teams[:selectedIndex], teams[selectedIndex+1:]...)
	mw.state.SetTeams(teams)

	// Mark as modified
	mw.state.MarkDirty()
	mw.statusBar.SetSavedStatus(true)

	// Refresh list and go back to list view
	mw.teamList.SetTeams(teams)
	mw.content.Objects = []fyne.CanvasObject{container.NewMax(mw.teamList.GetContainer())}
	mw.content.Refresh()
	mw.statusBar.SetRecordCount("Teams", len(teams))
}

// navigateTeam moves to the next or previous team
func (mw *MainWindow) navigateTeam(delta int) {
	selectedIndex := mw.state.GetSelectedIndex()
	teams := mw.state.GetTeams()

	newIndex := selectedIndex + delta
	if newIndex < 0 {
		newIndex = 0
	}
	if newIndex >= len(teams) {
		newIndex = len(teams) - 1
	}

	mw.state.SetSelectedIndex(newIndex)
	mw.updateTeamForm()
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

// fieldNameToFormField converts validation field names to form field names
// Examples: "FirstName" -> "firstName", "OverallRating" -> "overall"
func fieldNameToFormField(validationFieldName string) string {
	// Map of validation field names to form field names
	fieldMap := map[string]string{
		"FirstName":        "firstName",
		"LastName":         "lastName",
		"Team":             "team",
		"Position":         "position",
		"PositionKey":      "position",
		"Uniform":          "uniform",
		"OverallRating":    "overall",
		"Height":           "height",
		"Weight":           "weight",
		"HandSize":         "handSize",
		"ArmLength":        "armLength",
		"Experience":       "experience",
		"College":          "college",
		"YearEntry":        "yearEntry",
		"RoundDrafted":     "roundDrafted",
		"SelectionDrafted": "selectionDrafted",
		// Coach fields
		"BirthMonth":      "birthMonth",
		"BirthDay":        "birthDay",
		"BirthYear":       "birthYear",
		"BirthCity":       "birthCity",
		"BirthCityID":     "birthCityID",
		"CollegeID":       "collegeID",
		"PositionGroup":   "positionGroup",
		"OffensiveStyle":  "offensiveStyle",
		"DefensiveStyle":  "defensiveStyle",
		"PayScale":        "payScale",
		// Team fields
		"TeamName":        "teamName",
		"NickName":        "nickName",
		"Abbreviation":    "abbreviation",
		"Year":            "year",
		"TeamID":          "teamID",
		"Conference":      "conference",
		"Division":        "division",
		"City":            "city",
		"PrimaryRed":      "primaryRed",
		"PrimaryGreen":    "primaryGreen",
		"PrimaryBlue":     "primaryBlue",
		"SecondaryRed":    "secondaryRed",
		"SecondaryGreen":  "secondaryGreen",
		"SecondaryBlue":   "secondaryBlue",
		"Roof":            "roof",
		"Turf":            "turf",
		"Built":           "built",
		"Capacity":        "capacity",
		"Luxury":          "luxury",
		"Condition":       "condition",
		"Attendance":      "attendance",
		"Support":         "support",
		"FutureName":      "futureName",
		"FutureAbbr":      "futureAbbr",
		"FutureRoof":      "futureRoof",
		"FutureTurf":      "futureTurf",
		"FutureCap":       "futureCap",
		"FutureLuxury":    "futureLuxury",
	}

	if formFieldName, exists := fieldMap[validationFieldName]; exists {
		return formFieldName
	}

	// If not in map, return as-is (shouldn't happen)
	return validationFieldName
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
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
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

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		fileDialog.SetLocation(defaultLocation)
	}

	fileDialog.Show()
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
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
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

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		saveDialog.SetLocation(defaultLocation)
	}

	saveDialog.Show()
}

// handleWindowClose handles the window close event, prompting for unsaved changes
func (mw *MainWindow) handleWindowClose() {
	// Check for unsaved changes
	if mw.state.IsDirtyState() {
		dialog.ShowConfirm("Unsaved Changes",
			"You have unsaved changes. Close anyway?",
			func(close bool) {
				if close {
					mw.window.Close()
				}
			}, mw.window)
		return
	}

	// No unsaved changes, close immediately
	mw.window.Close()
}

// loadPlayersCSV loads players from a CSV file
func (mw *MainWindow) loadPlayersCSV() {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
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

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		fileDialog.SetLocation(defaultLocation)
	}

	fileDialog.Show()
}

// loadCoachesCSV loads coaches from a CSV file
func (mw *MainWindow) loadCoachesCSV() {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
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

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		fileDialog.SetLocation(defaultLocation)
	}

	fileDialog.Show()
}

// loadTeamsCSV loads teams from a CSV file
func (mw *MainWindow) loadTeamsCSV() {
	fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
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

		// Refresh player/coach forms if they're currently displayed (to update dropdowns)
		currentSection := mw.state.GetCurrentSection()
		if currentSection == "Players" && mw.state.GetSelectedIndex() >= 0 {
			mw.updatePlayerForm()
		} else if currentSection == "Coaches" && mw.state.GetSelectedIndex() >= 0 {
			mw.updateCoachForm()
		}

		// Update UI
		mw.sidebar.SetSelectedSection("Teams")
		mw.updateContentArea("Teams")
		mw.statusBar.SetProjectStatus("Teams CSV Loaded")

		dialog.ShowInformation("Success", fmt.Sprintf("Loaded %d teams", len(teams)), mw.window)
	}, mw.window)

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		fileDialog.SetLocation(defaultLocation)
	}

	fileDialog.Show()
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
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
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

		// Set default location to FOF9 installation folder
		if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
			saveDialog.SetLocation(defaultLocation)
		}

		saveDialog.Show()
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

// savePlayersCSV saves players to a CSV file
func (mw *MainWindow) savePlayersCSV() {
	players := mw.state.GetPlayers()
	if len(players) == 0 {
		dialog.ShowInformation("No Data", "No players to save.", mw.window)
		return
	}

	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if writer == nil {
			return
		}

		filePath := writer.URI().Path()

		// Close the writer immediately to release the file lock
		writer.Close()

		// Save players to CSV
		if err := data.SavePlayers(filePath, players); err != nil {
			dialog.ShowError(fmt.Errorf("failed to save players: %w", err), mw.window)
			return
		}

		// Mark as clean
		mw.state.MarkClean()
		mw.statusBar.SetSavedStatus(false)

		dialog.ShowInformation("Success", fmt.Sprintf("Saved %d players to %s", len(players), filepath.Base(filePath)), mw.window)
	}, mw.window)

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		saveDialog.SetLocation(defaultLocation)
	}

	saveDialog.Show()
}

// saveCoachesCSV saves coaches to a CSV file
func (mw *MainWindow) saveCoachesCSV() {
	coaches := mw.state.GetCoaches()
	if len(coaches) == 0 {
		dialog.ShowInformation("No Data", "No coaches to save.", mw.window)
		return
	}

	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if writer == nil {
			return
		}

		filePath := writer.URI().Path()

		// Close the writer immediately to release the file lock
		writer.Close()

		// Save coaches to CSV
		if err := data.SaveCoaches(filePath, coaches); err != nil {
			dialog.ShowError(fmt.Errorf("failed to save coaches: %w", err), mw.window)
			return
		}

		// Mark as clean
		mw.state.MarkClean()
		mw.statusBar.SetSavedStatus(false)

		dialog.ShowInformation("Success", fmt.Sprintf("Saved %d coaches to %s", len(coaches), filepath.Base(filePath)), mw.window)
	}, mw.window)

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		saveDialog.SetLocation(defaultLocation)
	}

	saveDialog.Show()
}

// saveTeamsCSV saves teams to a CSV file
func (mw *MainWindow) saveTeamsCSV() {
	teams := mw.state.GetTeams()
	if len(teams) == 0 {
		dialog.ShowInformation("No Data", "No teams to save.", mw.window)
		return
	}

	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, mw.window)
			return
		}
		if writer == nil {
			return
		}

		filePath := writer.URI().Path()

		// Close the writer immediately to release the file lock
		writer.Close()

		// Save teams to CSV
		if err := data.SaveTeams(filePath, teams); err != nil {
			dialog.ShowError(fmt.Errorf("failed to save teams: %w", err), mw.window)
			return
		}

		// Mark as clean
		mw.state.MarkClean()
		mw.statusBar.SetSavedStatus(false)

		dialog.ShowInformation("Success", fmt.Sprintf("Saved %d teams to %s", len(teams), filepath.Base(filePath)), mw.window)
	}, mw.window)

	// Set default location to FOF9 installation folder
	if defaultLocation := getDefaultCSVPath(); defaultLocation != nil {
		saveDialog.SetLocation(defaultLocation)
	}

	saveDialog.Show()
}
