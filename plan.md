# Front Office Football Nine CSV Editor - Implementation Plan

## Overview

This plan breaks down the implementation of the FOF9 CSV Editor (as specified in spec.md) into small, iterative steps. Each step builds on the previous one, with no orphaned code. The plan follows Go best practices and uses Fyne for the desktop UI.

---

## Architecture Overview

### Project Structure
```
fof9editor/
├── cmd/
│   └── fof9editor/
│       └── main.go                 # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go                  # Application lifecycle
│   ├── models/
│   │   ├── player.go               # Player data model
│   │   ├── coach.go                # Coach data model
│   │   ├── team.go                 # Team data model
│   │   ├── leagueinfo.go           # League info model
│   │   └── project.go              # Project file model
│   ├── data/
│   │   ├── csv_reader.go           # CSV parsing
│   │   ├── csv_writer.go           # CSV writing
│   │   ├── project_io.go           # Project file I/O
│   │   └── reference_loader.go     # Load cities, colleges, etc.
│   ├── validation/
│   │   ├── validator.go            # Validation engine
│   │   ├── player_rules.go         # Player validation rules
│   │   ├── coach_rules.go          # Coach validation rules
│   │   ├── team_rules.go           # Team validation rules
│   │   └── autofix.go              # Auto-fix suggestions
│   ├── ui/
│   │   ├── mainwindow.go           # Main window layout
│   │   ├── sidebar.go              # Navigation sidebar
│   │   ├── listview.go             # Record list view
│   │   ├── formview.go             # Edit form view
│   │   ├── statusbar.go            # Status bar
│   │   ├── dialogs/
│   │   │   ├── wizard.go           # New League Wizard
│   │   │   ├── settings.go         # Settings dialog
│   │   │   ├── columns.go          # Column config dialog
│   │   │   ├── filters.go          # Filter management dialog
│   │   │   └── validation.go       # Validation error dialog
│   │   └── widgets/
│   │       ├── searchable_select.go # Searchable dropdown
│   │       ├── color_picker.go      # Color picker widget
│   │       └── tooltip.go           # Enhanced tooltip
│   └── state/
│       ├── appstate.go             # Application state
│       ├── preferences.go          # User preferences
│       └── session.go              # Current editing session
├── pkg/
│   └── utils/
│       ├── idgen.go                # ID generation
│       ├── defaults.go             # Default value generation
│       └── stringutil.go           # String utilities
├── assets/
│   ├── icons/                      # Application icons
│   └── docs/                       # Documentation files
├── testdata/
│   └── fixtures/                   # Test data
├── go.mod
├── go.sum
├── README.md
├── spec.md                         # The specification
└── plan.md                         # This file
```

### Technology Stack
- **Language**: Go 1.21+
- **UI Framework**: Fyne v2.4+
- **CSV Parsing**: encoding/csv (standard library)
- **JSON**: encoding/json (standard library)
- **Testing**: testing package + testify for assertions

---

## Development Phases

The implementation is divided into manageable phases, each with iterative steps that build on each other.

### Phase 0: Project Setup (Steps 1-3)
Set up the project structure, dependencies, and build pipeline.

### Phase 1: Core Data Models (Steps 4-8)
Implement data structures for players, coaches, teams, and project files.

### Phase 2: CSV I/O (Steps 9-13)
Read and write CSV files with proper parsing and error handling.

### Phase 3: Basic UI Framework (Steps 14-20)
Create the main window, sidebar, and basic layout.

### Phase 4: List View (Steps 21-25)
Implement the record list display with sorting.

### Phase 5: Form View (Steps 26-32)
Create the edit form with field inputs and layout.

### Phase 6: Validation (Steps 33-38)
Add validation rules and error feedback.

### Phase 7: Reference Data (Steps 39-42)
Load and display cities, colleges, teams in dropdowns.

### Phase 8: File Operations (Steps 43-47)
Implement open/save/close league operations.

### Phase 9: New League Wizard (Steps 48-53)
Create the multi-step league creation wizard.

### Phase 10: Search and Filtering (Steps 54-59)
Add search, filter, and preset functionality.

### Phase 11: Settings and Preferences (Steps 60-64)
Implement settings dialog and user preferences.

### Phase 12: Auto-Save (Steps 65-68)
Add auto-save and crash recovery.

### Phase 13: Polish and Testing (Steps 69-75)
Final features, testing, and bug fixes.

---

## Detailed Step-by-Step Plan

---

### Phase 0: Project Setup

#### Step 1: Initialize Go Project
**Goal**: Create Go module and project structure

**Prompt**:
```
Initialize a new Go project for the FOF9 CSV Editor:

1. Create go.mod with module name "github.com/[username]/fof9editor"
2. Create the following directory structure:
   - cmd/fof9editor/
   - internal/app/
   - internal/models/
   - internal/data/
   - internal/validation/
   - internal/ui/
   - internal/state/
   - pkg/utils/
   - testdata/fixtures/

3. Create a basic main.go in cmd/fof9editor/ with:
   - Package main
   - A simple main() function that prints "FOF9 Editor Starting..."
   - Import statement for fmt

4. Create README.md with:
   - Project title: "Front Office Football Nine CSV Editor"
   - Brief description (1-2 sentences)
   - Build instructions: "go build ./cmd/fof9editor"
   - Run instructions: "go run ./cmd/fof9editor"

5. Test that the project compiles: go build ./cmd/fof9editor

Ensure all directories are created and the basic project structure is in place.
```

**Deliverable**: Compilable Go project with directory structure

---

#### Step 2: Add Fyne Dependency
**Goal**: Add Fyne UI framework to the project

**Prompt**:
```
Add Fyne UI framework to the FOF9 Editor project:

1. Run: go get fyne.io/fyne/v2@latest

2. Update main.go to create a minimal Fyne window:
   - Import fyne.io/fyne/v2/app
   - Import fyne.io/fyne/v2/widget
   - In main(), create a Fyne app: myApp := app.New()
   - Create a window: myWindow := myApp.NewWindow("FOF9 Editor")
   - Set window content to a simple label: widget.NewLabel("Hello from FOF9 Editor")
   - Set window size: myWindow.Resize(fyne.NewSize(800, 600))
   - Show and run: myWindow.ShowAndRun()

3. Test that the application launches and displays a window with the label

4. Update README.md to include required dependencies note for Fyne (e.g., on Windows, mention C compiler if needed for CGO)

Ensure the application window opens successfully.
```

**Deliverable**: Working Fyne application that opens a window

---

#### Step 3: Setup Build and CI
**Goal**: Configure local build and GitHub Actions

**Prompt**:
```
Set up build configuration and GitHub Actions for the FOF9 Editor:

1. Create Makefile with targets:
   - build: Compile the application to ./bin/fof9editor.exe (Windows)
   - run: Build and run the application
   - test: Run go test ./...
   - clean: Remove ./bin directory

2. Create .github/workflows/build.yml:
   - Name: "Build and Test"
   - Trigger: on push and pull_request
   - Jobs:
     - build-windows:
       - runs-on: windows-latest
       - steps:
         - Checkout code
         - Setup Go 1.21
         - Run: go mod download
         - Run: go build -o fof9editor.exe ./cmd/fof9editor
         - Upload artifact: fof9editor.exe
     - test:
       - runs-on: ubuntu-latest
       - steps:
         - Checkout code
         - Setup Go 1.21
         - Run: go mod download
         - Run: go test ./... -v

3. Create .gitignore:
   - /bin/
   - *.exe
   - *.dll
   - *.so
   - *.dylib
   - .DS_Store
   - *.log

4. Test locally: make build (should create bin/fof9editor.exe)

Ensure builds work locally and CI configuration is ready for GitHub.
```

**Deliverable**: Makefile, GitHub Actions workflow, .gitignore

---

### Phase 1: Core Data Models

#### Step 4: Project Model
**Goal**: Define the .fof9proj file structure

**Prompt**:
```
Create the project file data model in internal/models/project.go:

1. Define Project struct with fields matching spec.md section 3.2:
   - Version string
   - LeagueName string
   - Identifier string
   - Created time.Time
   - LastModified time.Time
   - BaseYear int
   - DataPath string
   - ReferencePath string
   - CSVFiles map[string]string (keys: "info", "players", "coaches", "teams", "teamColors")
   - UserPreferences map[string]interface{} (for future use, can be empty map)

2. Add helper methods:
   - NewProject(name, identifier, basePath string, baseYear int) *Project
     - Creates new project with current timestamp
     - Sets default paths (DataPath: "./data/", ReferencePath: "./reference/")
     - Initializes CSVFiles map with standard filenames
   - GetFullPath(csvKey string) string
     - Returns absolute path to a CSV file

3. Add JSON marshaling tags to all fields (e.g., `json:"version"`)

4. Add package-level comment explaining this file defines the project structure

Test: Create a sample Project in a test file, marshal to JSON, verify structure.
```

**Deliverable**: internal/models/project.go with Project struct and helpers

---

#### Step 5: Player Model
**Goal**: Define Player data structure

**Prompt**:
```
Create the player data model in internal/models/player.go:

1. Define Player struct with fields from spec.md section 3.3.1:
   - Basic fields: PlayerID, LastName, FirstName (all as appropriate types)
   - Team int
   - PositionKey int
   - Uniform int
   - Physical: Height, Weight, HandSize, ArmLength (all int)
   - Birth: BirthMonth, BirthDay, BirthYear (all int)
   - BirthCityID, CollegeID int
   - Draft: YearEntry, RoundDrafted, SelectionDrafted, Supplemental, OriginalTeam (all int)
   - Experience int
   - Contract: SalaryYears, SalaryYear1-5, BonusYear1-5 (all int)
   - OverallRating int
   - Skills: Map all 30+ skill attributes from spec (e.g., SkillSpeed, SkillPower, etc.) as int
   - BaseYear int

2. Add CSV struct tags for field mapping (e.g., `csv:"PLAYERID"`)

3. Add helper method:
   - GetDisplayName() string - returns "FirstName LastName"

4. Add package comment explaining this is the player data model

Do not worry about CSV parsing yet - just define the struct with proper tags.
```

**Deliverable**: internal/models/player.go with Player struct

---

#### Step 6: Coach Model
**Goal**: Define Coach data structure

**Prompt**:
```
Create the coach data model in internal/models/coach.go:

1. Define Coach struct with fields from spec.md section 3.3.2:
   - LastName, FirstName string
   - BirthMonth, BirthDay, BirthYear int
   - BirthCityID, CollegeID int
   - Team int
   - Position int (0=Head Coach, 1=OC, 2=DC, 3=ST, 4=S&C)
   - PositionGroup int
   - OffensiveStyle int
   - DefensiveStyle int
   - PayScale int

2. Add CSV struct tags for field mapping (e.g., `csv:"LASTNAME"`)

3. Add helper methods:
   - GetDisplayName() string - returns "FirstName LastName"
   - GetPositionName() string - maps Position int to readable name

4. Add constants for position types:
   - const (
       PositionHeadCoach = 0
       PositionOffensiveCoordinator = 1
       PositionDefensiveCoordinator = 2
       PositionSpecialTeamsCoordinator = 3
       PositionStrengthConditioning = 4
     )

5. Add package comment

Test: Create a sample Coach, call GetPositionName(), verify output.
```

**Deliverable**: internal/models/coach.go with Coach struct

---

#### Step 7: Team Model
**Goal**: Define Team data structure

**Prompt**:
```
Create the team data model in internal/models/team.go:

1. Define Team struct with fields from spec.md section 3.3.3:
   - Year, TeamID int
   - TeamName, NickName, Abbreviation string
   - Conference, Division int
   - City int
   - Colors: PrimaryRed, PrimaryGreen, PrimaryBlue, SecondaryRed, SecondaryGreen, SecondaryBlue (all int, 0-255)
   - Stadium: Roof, Turf, Built, Capacity, Luxury, Condition int
   - Financial: Attendance, Support int
   - Future stadium: Plan, Completed, Future, FutureName, FutureAbbr, FutureRoof, FutureTurf, FutureCap, FutureLuxury, TeamContribution (all int or string as appropriate)

2. Add CSV struct tags for field mapping

3. Add helper methods:
   - GetDisplayName() string - returns "TeamName NickName"
   - GetPrimaryColor() color.Color - returns color.RGBA from RGB values
   - GetSecondaryColor() color.Color - returns color.RGBA from RGB values

4. Add constants for roof/turf types:
   - const (
       RoofOutdoor = 0
       RoofDome = 1
       RoofRetractable = 2
       TurfGrass = 0
       TurfArtificial = 1
       TurfHybrid = 2
     )

5. Add package comment

Test: Create a Team, call GetPrimaryColor(), verify returns valid color.
```

**Deliverable**: internal/models/team.go with Team struct

---

#### Step 8: League Info Model
**Goal**: Define League configuration structure

**Prompt**:
```
Create the league info model in internal/models/leagueinfo.go:

1. Define LeagueInfo struct with fields from spec.md section 3.3.4:
   - ScheduleID string
   - BaseYear int
   - SalaryCap int
   - Minimum int (rookie minimum)
   - Salary1, Salary2, Salary3, Salary45, Salary789, Salary10 int

2. Add CSV struct tags for field mapping

3. Add helper methods:
   - GetSalaryMinimum(experience int) int - returns appropriate minimum based on experience
   - ValidateScheduleID() bool - checks format "x_y_z"

4. Add default constructor:
   - NewDefaultLeagueInfo(baseYear int) *LeagueInfo
     - Returns LeagueInfo with defaults from spec.md section 12.4

5. Add package comment

Test: Call NewDefaultLeagueInfo(2024), verify all fields have expected default values.
```

**Deliverable**: internal/models/leagueinfo.go with LeagueInfo struct

---

### Phase 2: CSV I/O

#### Step 9: CSV Reader Foundation
**Goal**: Create CSV parsing utility

**Prompt**:
```
Create CSV reading utility in internal/data/csv_reader.go:

1. Create CSVReader struct:
   - filePath string
   - headers []string

2. Implement NewCSVReader(filePath string) (*CSVReader, error):
   - Open file
   - Read first line as headers
   - Return CSVReader with headers populated
   - Handle file not found, permission errors

3. Implement ReadAll() ([]map[string]string, error):
   - Read all rows
   - Return slice of maps where key=header, value=cell value
   - Handle CSV parsing errors
   - Skip empty lines

4. Implement Close() error:
   - Close underlying file

5. Add helper function ParseCSVFile(filePath string) ([]map[string]string, error):
   - Convenience function that opens, reads, closes

6. Write unit test in csv_reader_test.go:
   - Create testdata/fixtures/test_players.csv with 3 sample players
   - Test ParseCSVFile returns correct number of rows
   - Test headers are correctly parsed

Ensure proper error handling for missing files and malformed CSV.
```

**Deliverable**: internal/data/csv_reader.go with CSV parsing

---

#### Step 10: Player CSV Loading
**Goal**: Parse players.csv into Player structs

**Prompt**:
```
Add player-specific CSV loading to internal/data/csv_reader.go:

1. Implement LoadPlayers(filePath string) ([]*models.Player, error):
   - Use ParseCSVFile to get raw data
   - For each row:
     - Create models.Player
     - Map CSV fields to struct fields using csv tags
     - Convert string values to appropriate types (strconv.Atoi for ints)
     - Handle conversion errors gracefully
   - Return slice of Player pointers

2. Add helper function mapRowToPlayer(row map[string]string) (*models.Player, error):
   - Takes a single row map
   - Returns Player struct with all fields populated
   - Handle missing fields (use zero values)
   - Handle invalid values (return error)

3. Write unit test:
   - Create testdata/fixtures/players.csv with 5 sample players
   - Call LoadPlayers
   - Verify correct number of players loaded
   - Verify a few key fields (PlayerID, LastName, Team, PositionKey)

Handle cases where fields are missing or have invalid values.
```

**Deliverable**: Player CSV loading in csv_reader.go

---

#### Step 11: Coach and Team CSV Loading
**Goal**: Parse coaches.csv and team_info.csv

**Prompt**:
```
Add coach and team CSV loading to internal/data/csv_reader.go:

1. Implement LoadCoaches(filePath string) ([]*models.Coach, error):
   - Similar to LoadPlayers
   - Map CSV fields to Coach struct
   - Convert string values to ints

2. Implement LoadTeams(filePath string) ([]*models.Team, error):
   - Similar to LoadPlayers
   - Map CSV fields to Team struct
   - Handle color fields (convert strings to ints)

3. Add helper functions:
   - mapRowToCoach(row map[string]string) (*models.Coach, error)
   - mapRowToTeam(row map[string]string) (*models.Team, error)

4. Write unit tests:
   - Create testdata/fixtures/coaches.csv with 3 sample coaches
   - Create testdata/fixtures/team_info.csv with 2 sample teams
   - Test LoadCoaches and LoadTeams
   - Verify key fields are correctly populated

Ensure consistent error handling across all Load functions.
```

**Deliverable**: Coach and Team CSV loading

---

#### Step 12: CSV Writer Foundation
**Goal**: Create CSV writing utility

**Prompt**:
```
Create CSV writing utility in internal/data/csv_writer.go:

1. Create CSVWriter struct:
   - filePath string
   - headers []string

2. Implement NewCSVWriter(filePath string, headers []string) *CSVWriter:
   - Returns CSVWriter with file path and headers

3. Implement WriteAll(rows []map[string]string) error:
   - Create/overwrite file
   - Write headers as first line
   - Write each row in header order
   - Handle missing values (write empty string)
   - Use atomic write: write to temp file, then rename
   - Handle write errors

4. Add helper function SaveCSV(filePath string, headers []string, rows []map[string]string) error:
   - Convenience function for one-shot write

5. Write unit test:
   - Create sample data (3 rows)
   - Call SaveCSV to /tmp/test_output.csv
   - Read file back with CSVReader
   - Verify data matches

Ensure atomic writes so original file is never corrupted on error.
```

**Deliverable**: internal/data/csv_writer.go with CSV writing

---

#### Step 13: Player/Coach/Team CSV Saving
**Goal**: Save struct data back to CSV

**Prompt**:
```
Add struct-to-CSV conversion in internal/data/csv_writer.go:

1. Implement SavePlayers(filePath string, players []*models.Player) error:
   - Convert Player structs to []map[string]string
   - Use csv tags to get field names for headers
   - Convert int/other types to strings
   - Call SaveCSV with data

2. Implement SaveCoaches(filePath string, coaches []*models.Coach) error:
   - Similar to SavePlayers

3. Implement SaveTeams(filePath string, teams []*models.Team) error:
   - Similar to SavePlayers

4. Add helper functions:
   - playerToRow(player *models.Player) map[string]string
   - coachToRow(coach *models.Coach) map[string]string
   - teamToRow(team *models.Team) map[string]string

5. Write integration test:
   - Load testdata/fixtures/players.csv with LoadPlayers
   - Modify one player's name
   - Save with SavePlayers to /tmp/modified_players.csv
   - Load modified file
   - Verify change was saved

Ensure field ordering matches the original CSV column order.
```

**Deliverable**: SavePlayers, SaveCoaches, SaveTeams functions

---

### Phase 3: Basic UI Framework

#### Step 14: Application State
**Goal**: Create central application state manager

**Prompt**:
```
Create application state manager in internal/state/appstate.go:

1. Define AppState struct:
   - Project *models.Project
   - Players []*models.Player
   - Coaches []*models.Coach
   - Teams []*models.Team
   - LeagueInfo *models.LeagueInfo
   - IsModified bool (unsaved changes flag)
   - CurrentSection string (e.g., "Players", "Coaches")
   - SelectedRecordIndex int (which record is selected in list)

2. Implement singleton pattern:
   - var instance *AppState
   - func GetInstance() *AppState - returns singleton instance

3. Add methods:
   - LoadProject(projectPath string) error - loads all CSV files
   - SaveProject() error - saves all modified CSVs
   - SetModified(modified bool) - sets IsModified flag
   - GetCurrentRecords() interface{} - returns current section's records

4. Add stub implementations (will expand later):
   - LoadProject: just set Project field, load CSVs with data.Load* functions
   - SaveProject: just call data.Save* functions

5. Write unit test:
   - Create test project structure
   - Call LoadProject
   - Verify Players/Coaches/Teams are populated

This is the central data store for the application.
```

**Deliverable**: internal/state/appstate.go with AppState singleton

---

#### Step 15: Main Window Foundation
**Goal**: Create main application window structure

**Prompt**:
```
Create main window in internal/ui/mainwindow.go:

1. Define MainWindow struct:
   - window fyne.Window
   - appState *state.AppState

2. Implement NewMainWindow(app fyne.App) *MainWindow:
   - Create window with title "FOF9 Editor"
   - Set window size 1024x768
   - Get AppState instance
   - Return MainWindow

3. Implement Show() method:
   - window.ShowAndRun()

4. Implement SetContent(content fyne.CanvasObject):
   - window.SetContent(content)

5. Update cmd/fof9editor/main.go:
   - Remove simple hello world
   - Create MainWindow
   - Set content to a placeholder label "Main content area"
   - Call Show()

6. Test: Run application, verify window opens with placeholder text

This establishes the main window shell.
```

**Deliverable**: internal/ui/mainwindow.go with basic window

---

#### Step 16: Status Bar
**Goal**: Create status bar at bottom of window

**Prompt**:
```
Create status bar in internal/ui/statusbar.go:

1. Define StatusBar struct:
   - container *fyne.Container
   - statusLabel *widget.Label (left section)
   - recordCountLabel *widget.Label (center-left section)
   - projectLabel *widget.Label (center-right section)
   - modifiedLabel *widget.Label (right section)

2. Implement NewStatusBar() *StatusBar:
   - Create 4 labels with default text
   - Use container.NewHBox to arrange horizontally
   - Add separators between sections (widget.NewSeparator())
   - Return StatusBar

3. Add update methods:
   - SetStatus(text string) - updates left section
   - SetRecordCount(count, total int) - updates record count
   - SetProject(name string) - updates project name
   - SetModified(modified bool) - updates modified indicator

4. Implement GetContainer() *fyne.Container:
   - Returns the container for embedding in main layout

5. Update MainWindow:
   - Add statusBar *StatusBar field
   - Create status bar in NewMainWindow
   - Use container.NewBorder to set status bar at bottom
   - Set initial status: "Ready"

6. Test: Run app, verify status bar appears at bottom with default text

Status bar provides user feedback.
```

**Deliverable**: internal/ui/statusbar.go with status bar widget

---

#### Step 17: Sidebar Navigation
**Goal**: Create left sidebar for section navigation

**Prompt**:
```
Create sidebar in internal/ui/sidebar.go:

1. Define Sidebar struct:
   - container *fyne.Container
   - sections []string (e.g., ["League Info", "Players", "Coaches", "Teams"])
   - selectedSection string
   - onSectionChange func(section string) (callback)

2. Implement NewSidebar(onSectionChange func(string)) *Sidebar:
   - Create list of section names
   - Use widget.NewList or container.NewVBox with buttons
   - Each button triggers onSectionChange callback
   - Return Sidebar

3. Implement GetContainer() *fyne.Container:
   - Returns container for embedding

4. Implement SetSelected(section string):
   - Highlights selected section
   - Updates selectedSection field

5. Update MainWindow:
   - Add sidebar *Sidebar field
   - Create sidebar in NewMainWindow with callback:
     - Callback: func(section string) { appState.CurrentSection = section; updateContent() }
   - Use container.NewBorder to place sidebar on left
   - Add placeholder for main content area (center)

6. Test: Run app, click sidebar sections, verify callback is triggered (log to console for now)

Sidebar enables navigation between entity types.
```

**Deliverable**: internal/ui/sidebar.go with navigation sidebar

---

#### Step 18: Layout Integration
**Goal**: Wire sidebar, content area, and status bar together

**Prompt**:
```
Integrate layout components in internal/ui/mainwindow.go:

1. Add fields to MainWindow:
   - sidebar *Sidebar
   - statusBar *StatusBar
   - contentArea *fyne.Container (center area for list/form)

2. Update NewMainWindow:
   - Create sidebar with onSectionChange callback
   - Create status bar
   - Create empty contentArea with widget.NewLabel("Select a section from the sidebar")
   - Use container.NewBorder layout:
     - Top: nil (menu bar later)
     - Bottom: statusBar.GetContainer()
     - Left: sidebar.GetContainer()
     - Right: nil
     - Center: contentArea

3. Implement updateContent() method:
   - Clears contentArea
   - Based on appState.CurrentSection, shows different content
   - For now, just show label: "Viewing: [section name]"

4. Implement onSectionChange callback:
   - Updates appState.CurrentSection
   - Calls updateContent()
   - Updates status bar: statusBar.SetStatus("Loaded " + section)

5. Test: Run app, click sidebar sections, verify center area updates with section name

This establishes the main layout structure.
```

**Deliverable**: Integrated main window layout

---

#### Step 19: Menu Bar
**Goal**: Add File/Edit/Help menu

**Prompt**:
```
Add menu bar to main window in internal/ui/mainwindow.go:

1. Create buildMenu() *fyne.MainMenu method in MainWindow:
   - File menu:
     - New League (disabled for now)
     - Open League (disabled for now)
     - Save (disabled for now)
     - Save As (disabled for now)
     - Separator
     - Exit (calls app.Quit())
   - Edit menu:
     - Settings (disabled for now)
   - Help menu:
     - About (shows simple dialog with app name and version)

2. Update NewMainWindow:
   - Call window.SetMainMenu(buildMenu())

3. Implement About dialog in buildMenu:
   - Use dialog.ShowInformation
   - Title: "About FOF9 Editor"
   - Message: "Front Office Football Nine CSV Editor\nVersion 0.1.0\n\n© 2024"

4. Test: Run app, verify menu bar appears, File > Exit works, Help > About shows dialog

Menu bar provides access to core commands.
```

**Deliverable**: Menu bar with File/Edit/Help menus

---

#### Step 20: Theme Support
**Goal**: Add light/dark theme switching

**Prompt**:
```
Add theme support in internal/state/preferences.go:

1. Create Preferences struct:
   - Theme string ("light", "dark", "system")
   - FontFamily string
   - FontSize int
   - AutoSaveEnabled bool
   - AutoSaveInterval int (minutes)

2. Implement:
   - NewDefaultPreferences() *Preferences - returns defaults
   - SavePreferences(prefs *Preferences, path string) error - save to JSON
   - LoadPreferences(path string) (*Preferences, error) - load from JSON

3. Add to AppState:
   - Preferences *Preferences field
   - LoadPreferences() method
   - SavePreferences() method

4. Update MainWindow:
   - In NewMainWindow, load preferences
   - Apply theme based on preferences.Theme:
     - Use app.Settings().SetTheme(theme.LightTheme()) or theme.DarkTheme()

5. Add to Help menu:
   - View > Light Theme
   - View > Dark Theme
   - On select, update preferences, save, apply theme

6. Test: Run app, switch themes via menu, verify UI updates

Theme support improves accessibility.
```

**Deliverable**: Theme switching in preferences

---

### Phase 4: List View

#### Step 21: List View Foundation
**Goal**: Create record list display component

**Prompt**:
```
Create list view in internal/ui/listview.go:

1. Define ListView struct:
   - container *fyne.Container
   - table *widget.Table
   - columns []string (column names)
   - data [][]string (2D array of cell values)
   - onRowSelect func(rowIndex int) (callback)

2. Implement NewListView(columns []string, onRowSelect func(int)) *ListView:
   - Create widget.Table with:
     - Length func: returns (len(columns), len(data))
     - CreateCell func: returns widget.NewLabel("")
     - UpdateCell func: sets label text to data[row][col]
   - Set table.OnSelected to trigger onRowSelect callback
   - Return ListView

3. Implement SetData(data [][]string):
   - Updates data field
   - Calls table.Refresh()

4. Implement GetContainer() *fyne.Container:
   - Returns container with table

5. Write simple test:
   - Create ListView with columns ["Name", "Team", "Position"]
   - Set data with 3 rows
   - Verify table displays correctly (manual test by rendering)

This is the foundation for displaying lists of records.
```

**Deliverable**: internal/ui/listview.go with table widget

---

#### Step 22: Player List View
**Goal**: Display players in list view

**Prompt**:
```
Add player list display to main window:

1. In internal/ui/mainwindow.go, add method:
   - showPlayersList()
     - Get players from appState
     - Define columns: ["Name", "Team", "Position", "Overall Rating", "Experience"]
     - Convert players to [][]string (2D array):
       - Row 0: [player[0].GetDisplayName(), getTeamName(player[0].Team), getPositionName(player[0].PositionKey), strconv.Itoa(player[0].OverallRating), strconv.Itoa(player[0].Experience)]
       - Repeat for all players
     - Create ListView with columns and data
     - Set onRowSelect callback: func(index int) { appState.SelectedRecordIndex = index; showPlayerForm() }
     - Set contentArea to listView.GetContainer()

2. Add helper functions:
   - getTeamName(teamID int) string - lookup team name from appState.Teams, return "Free Agent" if 0
   - getPositionName(posKey int) string - return readable position name

3. Update updateContent() in MainWindow:
   - If CurrentSection == "Players", call showPlayersList()

4. Test: Load sample project, click "Players" in sidebar, verify list of players appears

Players are now displayed in a table.
```

**Deliverable**: Player list view in main window

---

#### Step 23: Coach and Team List Views
**Goal**: Display coaches and teams in list view

**Prompt**:
```
Add coach and team list displays to main window:

1. In internal/ui/mainwindow.go, add methods:
   - showCoachesList()
     - Columns: ["Name", "Team", "Position", "Style", "Pay Scale"]
     - Convert coaches to [][]string
     - Create ListView
     - Set onRowSelect callback

   - showTeamsList()
     - Columns: ["Name", "City", "Stadium", "Conference", "Division"]
     - Convert teams to [][]string
     - Create ListView
     - Set onRowSelect callback

2. Update updateContent():
   - If CurrentSection == "Coaches", call showCoachesList()
   - If CurrentSection == "Teams", call showTeamsList()

3. Add helper:
   - getCityName(cityID int) string - lookup from reference data (stub for now: return "City " + strconv.Itoa(cityID))

4. Test: Click "Coaches" sidebar, verify coaches list appears. Click "Teams", verify teams list appears.

All entity types now have list views.
```

**Deliverable**: Coach and team list views

---

#### Step 24: Column Sorting
**Goal**: Allow sorting by clicking column headers

**Prompt**:
```
Add sorting to ListView in internal/ui/listview.go:

1. Add to ListView struct:
   - sortColumn int (currently sorted column, -1 = none)
   - sortAscending bool

2. Implement table header click handling:
   - In NewListView, add table.OnHeaderClick callback
   - On header click:
     - If clicked column == sortColumn, toggle sortAscending
     - Else, set sortColumn to clicked column, sortAscending = true
     - Call sortData()
     - Refresh table

3. Implement sortData() method:
   - Sort data [][]string by sortColumn
   - Use sort.Slice with custom comparison:
     - Compare data[i][sortColumn] vs data[j][sortColumn]
     - Handle ascending/descending based on sortAscending

4. Add visual indicator:
   - In UpdateCell for header row, append "↑" or "↓" to sorted column label

5. Test: Click "Name" column header in Players list, verify list sorts alphabetically. Click again, verify sorts descending.

Users can now sort by any column.
```

**Deliverable**: Sortable columns in ListView

---

#### Step 25: List Selection Highlighting
**Goal**: Highlight selected row in list

**Prompt**:
```
Add row selection highlighting to ListView:

1. Update ListView:
   - Add selectedRow int field (default -1)
   - In NewListView, set table.OnSelected callback:
     - Sets selectedRow = rowIndex
     - Calls onRowSelect(rowIndex)
     - Refreshes table

2. Update UpdateCell function:
   - If row == selectedRow, set background color (use theme accent color)
   - Else, use default background

3. Implement Select(rowIndex int) method:
   - Sets selectedRow
   - Calls table.Select(widget.TableCellID{Row: rowIndex, Col: 0})
   - Refreshes table

4. Test: Click a row in Players list, verify row is highlighted

Selected row is now visually distinct.
```

**Deliverable**: Row selection highlighting

---

### Phase 5: Form View

#### Step 26: Form View Foundation
**Goal**: Create record edit form component

**Prompt**:
```
Create form view in internal/ui/formview.go:

1. Define FormView struct:
   - container *fyne.Container
   - fields map[string]fyne.CanvasObject (field name -> widget)
   - onSave func() (callback)
   - onDelete func() (callback)
   - onNext func() (callback)
   - onPrev func() (callback)

2. Implement NewFormView() *FormView:
   - Create empty container
   - Return FormView

3. Implement SetFields(fieldDefs []FieldDef):
   - FieldDef struct: {Name, Label, Type, Value string}
   - For each field, create appropriate widget:
     - Type "text": widget.NewEntry()
     - Type "number": widget.NewEntry() (validate numeric)
     - Type "select": widget.NewSelect(options, onChange)
   - Arrange in form layout (container.NewVBox or container.NewGridWithColumns)
   - Add to fields map

4. Implement AddButtons():
   - Create button bar: [< Previous] [Next >] [Save] [Delete]
   - Wire to callbacks

5. Implement GetContainer() *fyne.Container:
   - Returns container

6. Test: Create FormView with 3 fields (name, team, position), render in window

This is the foundation for editing records.
```

**Deliverable**: internal/ui/formview.go with form widget

---

#### Step 27: Player Form View
**Goal**: Display and edit player fields in form

**Prompt**:
```
Add player form to main window:

1. In internal/ui/mainwindow.go, add method:
   - showPlayerForm()
     - Get selected player from appState (using SelectedRecordIndex)
     - Create FormView
     - Define fields for player (start with key fields):
       - First Name (text)
       - Last Name (text)
       - Team (select)
       - Position (select)
       - Uniform (number)
       - Overall Rating (number)
     - Set field values from player data
     - Wire callbacks:
       - onSave: update appState player, set modified flag, save
       - onDelete: remove player, refresh list
       - onNext: increment SelectedRecordIndex, reload form
       - onPrev: decrement SelectedRecordIndex, reload form
     - Set contentArea to formView.GetContainer()

2. Update showPlayersList:
   - onRowSelect callback should call showPlayerForm()

3. Test: Select a player from list, verify form displays with player data. Edit a field, click Save, verify data updates.

Player editing is now functional.
```

**Deliverable**: Player form view

---

#### Step 28: Split View Layout
**Goal**: Show list and form simultaneously (vertical split)

**Prompt**:
```
Implement split view layout for list and form:

1. In internal/ui/mainwindow.go, update showPlayersList:
   - Instead of replacing contentArea, create split layout:
     - Top: ListView (players list)
     - Bottom: FormView (player form)
     - Use container.NewVSplit(listView.GetContainer(), formView.GetContainer())
     - Set split ratio: 40% list, 60% form (using split.SetOffset(0.4))

2. Add draggable divider:
   - Fyne's VSplit already has draggable divider - just ensure it's enabled

3. Update onRowSelect:
   - Selecting row in list updates form in-place
   - No need to replace entire content area

4. Implement in showPlayersList:
   - Create listView
   - Create formView (initially empty)
   - On row select, call formView.SetFields with selected player data
   - Set contentArea to split container

5. Test: Click Players sidebar, verify list appears on top, form on bottom. Click a row, verify form updates. Drag divider, verify resizing works.

List and form are now visible simultaneously.
```

**Deliverable**: Split view layout

---

#### Step 29: Form Field Groups
**Goal**: Group player attributes into collapsible sections

**Prompt**:
```
Add collapsible field groups to FormView:

1. Update FormView to support groups:
   - Add method: AddGroup(groupName string, fields []FieldDef, collapsed bool)
   - Create widget.NewAccordion for each group
   - Add fields to accordion item
   - Collapsed state determines if group is open by default

2. Update showPlayerForm:
   - Group player fields:
     - Basic Info: First Name, Last Name, Team, Position, Uniform
     - Physical Attributes: Height, Weight, Hand Size, Arm Length
     - Birth Info: Birth Month, Birth Day, Birth Year, Birth City
     - Career: Experience, Overall Rating, Draft Info
     - Contract: Salary Years, Salary/Bonus fields
   - Each group is collapsible

3. Implement accordion in FormView:
   - Use widget.NewAccordion with widget.NewAccordionItem for each group

4. Test: Open player form, verify groups appear. Click group header, verify expand/collapse works.

Form is now organized into logical sections.
```

**Deliverable**: Collapsible field groups in form

---

#### Step 30: Coach and Team Forms
**Goal**: Create forms for coaches and teams

**Prompt**:
```
Add coach and team forms to main window:

1. In internal/ui/mainwindow.go, add methods:
   - showCoachForm()
     - Similar to showPlayerForm
     - Groups: Basic Info, Birth Info, Position & Role, Compensation
     - Fields: Name, Team, Position, Styles, Pay Scale

   - showTeamForm()
     - Groups: Team Identity, Stadium Info, Colors, Financial, Future Plans
     - Fields: Team Name, Nickname, City, Stadium details, RGB colors
     - For color fields, use custom color picker widget (stub for now: text entry)

2. Update showCoachesList and showTeamsList:
   - Use split view layout
   - Wire onRowSelect to show form

3. Test: Select coach, verify form appears with coach data. Select team, verify form with team data.

All entity types now have editable forms.
```

**Deliverable**: Coach and team forms

---

#### Step 31: Form Validation UI
**Goal**: Show validation errors in form

**Prompt**:
```
Add validation feedback to FormView:

1. Add to FormView struct:
   - errors map[string]string (field name -> error message)
   - errorContainer *fyne.Container (displays errors below form)

2. Implement SetErrors(errors map[string]string):
   - Clears errorContainer
   - For each error:
     - Create label with red text: "• [Field]: [Error message]"
     - Add to errorContainer
   - If no errors, hide errorContainer

3. Update SetFields:
   - For each field widget:
     - Add red border if field has error (use canvas.NewRectangle with red stroke)
     - Show error icon next to field (widget.NewIcon with theme.ErrorIcon())

4. Implement ClearErrors():
   - Clears errors map
   - Removes error indicators from fields
   - Hides errorContainer

5. Test: Manually add errors with SetErrors({"FirstName": "Required"}), verify red border appears on First Name field, error message displays below form.

Form now shows validation feedback.
```

**Deliverable**: Validation error display in FormView

---

#### Step 32: Previous/Next Navigation
**Goal**: Navigate through filtered records with buttons

**Prompt**:
```
Implement Previous/Next navigation in forms:

1. In FormView, wire button callbacks:
   - Previous button:
     - Calls onPrev()
     - Passed from parent (MainWindow)
   - Next button:
     - Calls onNext()

2. In MainWindow, implement navigation logic:
   - onPrev callback:
     - Decrements appState.SelectedRecordIndex
     - If < 0, wrap to last record
     - Reload form with new selected record
   - onNext callback:
     - Increments appState.SelectedRecordIndex
     - If > len(records), wrap to 0
     - Reload form

3. Update ListView:
   - When form calls onPrev/onNext, also update ListView selection highlight

4. Add keyboard shortcuts:
   - Alt+Left: Previous
   - Alt+Right: Next
   - Bind in MainWindow using window.Canvas().AddShortcut

5. Test: Select player, click Next button, verify next player loads. Click Previous, verify previous player loads. Test keyboard shortcuts.

Users can now navigate records without returning to list.
```

**Deliverable**: Previous/Next navigation

---

### Phase 6: Validation

#### Step 33: Validation Engine
**Goal**: Create validation rule framework

**Prompt**:
```
Create validation engine in internal/validation/validator.go:

1. Define ValidationError struct:
   - Field string
   - Message string
   - Severity string ("error" or "warning")

2. Define Validator interface:
   - Validate(record interface{}) []ValidationError

3. Create ValidationEngine struct:
   - validators []Validator

4. Implement methods:
   - AddValidator(v Validator)
   - Validate(record interface{}) []ValidationError
     - Calls all registered validators
     - Collects all errors
     - Returns combined list

5. Write unit test:
   - Create mock validator that returns error for test case
   - Add to engine
   - Validate test record
   - Verify error is returned

This is the framework for all validation rules.
```

**Deliverable**: internal/validation/validator.go with validation framework

---

#### Step 34: Player Field Validation
**Goal**: Implement player field validation rules

**Prompt**:
```
Create player validation rules in internal/validation/player_rules.go:

1. Define PlayerValidator struct (implements Validator interface)

2. Implement Validate(record interface{}) []ValidationError:
   - Cast record to *models.Player
   - Apply rules from spec.md section 6.1:
     - LastName: required, max 18 chars
     - FirstName: required, max 16 chars
     - Team: must exist in teams or = 0
     - PositionKey: 2-28
     - Uniform: 0-99
     - Height, Weight: > 0
     - BirthMonth: 1-12
     - BirthDay: 1-31
     - Experience: 0-23
     - OverallRating: 0-10
     - SalaryYears: 0-5
     - Skill attributes: -1 or 0-250
   - For each violation, create ValidationError
   - Return all errors

3. Write unit tests:
   - Test valid player: no errors
   - Test invalid position: error returned
   - Test missing name: error returned

Field validation catches basic errors.
```

**Deliverable**: internal/validation/player_rules.go with player validation

---

#### Step 35: Cross-Field Validation
**Goal**: Validate player age vs experience, draft year, etc.

**Prompt**:
```
Add cross-field validation to PlayerValidator:

1. In Validate method, add cross-field checks:
   - Age vs experience:
     - Calculate age: baseYear - birthYear
     - If age < 18 + experience: warning "Player age inconsistent with experience"
   - Draft year vs birth year:
     - If yearEntry < birthYear + 21: warning "Draft year too early for birth year"
   - Contract validation:
     - If salaryYears > 0, check salaryYear1-N are populated
     - If salaryYears = 3, salaryYear1-3 must be > 0

2. For warnings (not errors), set Severity = "warning"

3. Write unit tests:
   - Test player with age 20, experience 10: warning returned
   - Test player with valid age/experience: no warning

4. Update ValidationError to distinguish errors vs warnings

Cross-field validation catches logical inconsistencies.
```

**Deliverable**: Cross-field validation in player_rules.go

---

#### Step 36: Coach and Team Validation
**Goal**: Validate coaches and teams

**Prompt**:
```
Create coach and team validation rules:

1. Create internal/validation/coach_rules.go:
   - CoachValidator struct
   - Validate method with rules from spec.md section 6.2:
     - LastName, FirstName: required
     - BirthMonth: 1-12
     - BirthDay: 1-31
     - Team: must exist or = 0
     - Position: 0-4
     - OffensiveStyle: 0-6
     - DefensiveStyle: 0-4
     - PayScale: > 0

2. Create internal/validation/team_rules.go:
   - TeamValidator struct
   - Validate method with rules from spec.md section 6.3:
     - TeamID: unique, > 0
     - TeamName, NickName, Abbreviation: required
     - City: must exist in cities
     - RGB values: 0-255
     - Roof: 0-2
     - Turf: 0-2
     - Capacity: > 0
     - Future stadium: if Plan=1, Completed >= current year

3. Write unit tests for both

All entity types now have validation.
```

**Deliverable**: coach_rules.go and team_rules.go

---

#### Step 37: Real-Time Validation
**Goal**: Validate fields as user types (on blur)

**Prompt**:
```
Integrate validation into FormView:

1. Update FormView:
   - Add validator *validation.ValidationEngine field
   - In SetFields, for each text/number entry widget:
     - Add OnChanged callback:
       - On blur (when focus leaves field), run validation
       - Call validator.Validate on current record
       - Filter errors to only current field
       - If error, show red border and error message
       - If no error, clear red border

2. Implement ValidateField(fieldName string) error:
   - Runs validation on current record
   - Returns error for specific field only

3. Update showPlayerForm:
   - Create validation engine
   - Add PlayerValidator
   - Pass to FormView

4. Test: Open player form, enter invalid position (e.g., 99), tab out of field, verify red border appears and error shows below form

Real-time validation provides immediate feedback.
```

**Deliverable**: Real-time validation in FormView

---

#### Step 38: Save-Time Validation
**Goal**: Enforce validation on save, block if errors

**Prompt**:
```
Add save-time validation enforcement:

1. Update FormView onSave callback:
   - Before saving, run full validation on record
   - If any errors (severity = "error"):
     - Show validation error dialog (use dialog.ShowError)
     - List all errors
     - Block save
     - Return without saving
   - If only warnings:
     - Show confirmation dialog: "There are warnings. Save anyway?"
     - If user confirms, proceed with save
     - If user cancels, return without saving
   - If no errors/warnings, save normally

2. Update MainWindow:
   - In onSave callback for player form:
     - Update player in appState with form values
     - Run ValidationEngine.Validate
     - If errors, show dialog, don't save
     - If no errors, set appState.IsModified = true, update status bar

3. Test: Edit player with invalid data, click Save, verify error dialog appears and save is blocked. Fix error, click Save, verify save succeeds.

Invalid data cannot be saved.
```

**Deliverable**: Save-time validation enforcement

---

### Phase 7: Reference Data

#### Step 39: Reference Data Loader
**Goal**: Load cities.csv and colleges.csv for lookups

**Prompt**:
```
Create reference data loader in internal/data/reference_loader.go:

1. Define ReferenceData struct:
   - Cities map[int]string (cityID -> city name)
   - Colleges map[int]string (collegeID -> college name)

2. Implement LoadReferenceData(basePath string) (*ReferenceData, error):
   - Load cities.csv from basePath + "/cities.csv"
   - Parse CSV, populate Cities map
   - Load colleges.csv
   - Parse CSV, populate Colleges map
   - Return ReferenceData

3. Add to AppState:
   - ReferenceData *data.ReferenceData field
   - LoadReferenceData() method

4. Update MainWindow:
   - In LoadProject, also call appState.LoadReferenceData()

5. Write unit test:
   - Create testdata/fixtures/cities.csv with 5 cities
   - Create testdata/fixtures/colleges.csv with 5 colleges
   - Call LoadReferenceData
   - Verify maps are populated

Reference data is now available for lookups.
```

**Deliverable**: internal/data/reference_loader.go

---

#### Step 40: Searchable Dropdown Widget
**Goal**: Create searchable select widget for references

**Prompt**:
```
Create searchable dropdown in internal/ui/widgets/searchable_select.go:

1. Define SearchableSelect struct:
   - entry *widget.Entry (text input for search)
   - list *widget.List (list of matching options)
   - options map[int]string (id -> label)
   - filteredIDs []int (currently visible IDs)
   - onSelected func(id int) (callback)
   - container *fyne.Container

2. Implement NewSearchableSelect(options map[int]string, onSelected func(int)) *SearchableSelect:
   - Create entry with OnChanged callback:
     - Filter options by search text
     - Update filteredIDs
     - Refresh list
   - Create list with filtered options
   - On list item select, call onSelected callback
   - Return SearchableSelect

3. Implement SetValue(id int):
   - Sets entry text to options[id]

4. Implement GetContainer() *fyne.Container:
   - Returns container with entry and list (in overlay or popup)

5. Test: Create widget with 100 options, type "New", verify list filters to only options containing "New"

Searchable dropdown enables fast selection from large lists.
```

**Deliverable**: internal/ui/widgets/searchable_select.go

---

#### Step 41: Integrate Searchable Dropdowns
**Goal**: Use searchable dropdowns for city, college, team fields

**Prompt**:
```
Update forms to use searchable dropdowns:

1. In FormView, update SetFields:
   - For field type "searchable_select":
     - Create SearchableSelect widget
     - Pass options from reference data
     - Add to form

2. Update showPlayerForm:
   - For BirthCity field:
     - Type: "searchable_select"
     - Options: appState.ReferenceData.Cities
   - For College field:
     - Type: "searchable_select"
     - Options: appState.ReferenceData.Colleges
   - For Team field:
     - Options: Convert appState.Teams to map[int]string

3. Test: Open player form, click Birth City field, type "New York", verify dropdown filters to cities containing "New York". Select a city, verify field updates.

Reference fields are now easy to search and select.
```

**Deliverable**: Searchable dropdowns in forms

---

#### Step 42: Invalid Reference Handling
**Goal**: Show warning for invalid references (e.g., city ID doesn't exist)

**Prompt**:
```
Handle invalid references in searchable dropdowns:

1. Update SearchableSelect:
   - In NewSearchableSelect, check if current value exists in options
   - If not, add special entry: "[ID] (Invalid - not found)"
   - Mark with red text color

2. Update validation:
   - PlayerValidator checks if cityID exists in ReferenceData.Cities
   - If not, add warning: "Invalid city ID [ID] - not found in reference data"

3. Update FormView:
   - If field has validation warning, show warning icon next to dropdown

4. Test: Create test player with cityID = 99999 (invalid), open form, verify dropdown shows "99999 (Invalid - not found)" in red text. Verify validation warning appears.

Invalid references are clearly marked.
```

**Deliverable**: Invalid reference handling

---

### Phase 8: File Operations

#### Step 43: Project File I/O
**Goal**: Save and load .fof9proj files

**Prompt**:
```
Implement project file I/O in internal/data/project_io.go:

1. Implement SaveProject(project *models.Project, filePath string) error:
   - Marshal project to JSON with indentation
   - Write to filePath atomically (temp file + rename)
   - Handle errors

2. Implement LoadProject(filePath string) (*models.Project, error):
   - Read file
   - Unmarshal JSON to Project struct
   - Validate required fields are present
   - Return project

3. Write unit tests:
   - Create test project
   - Save to /tmp/test.fof9proj
   - Load back
   - Verify all fields match

4. Update AppState:
   - LoadProject method now also loads .fof9proj
   - SaveProject saves .fof9proj with updated LastModified timestamp

Project files can now be saved and loaded.
```

**Deliverable**: internal/data/project_io.go

---

#### Step 44: Open League
**Goal**: Implement File > Open League

**Prompt**:
```
Implement Open League in main window:

1. Update MainWindow:
   - Enable "File > Open League" menu item
   - On select, show file picker dialog (dialog.ShowFileOpen)
   - Filter to *.fof9proj files
   - On file selected, call openLeague(filePath)

2. Implement openLeague(filePath string) error:
   - If appState.IsModified, prompt to save first
   - Call appState.LoadProject(filePath)
   - Load all CSV files (players, coaches, teams)
   - Load reference data
   - Update sidebar to show "Players" section
   - Call showPlayersList
   - Update status bar: "Opened [league name]"
   - Add to recent files list

3. Handle errors:
   - If file not found, show error dialog
   - If corrupted project file, show error dialog with details

4. Test: Create sample .fof9proj, use File > Open, verify league loads and players list appears

Users can now open existing leagues.
```

**Deliverable**: Open League functionality

---

#### Step 45: Save League
**Goal**: Implement File > Save

**Prompt**:
```
Implement Save League in main window:

1. Update MainWindow:
   - Enable "File > Save" menu item
   - On select, call saveLeague()

2. Implement saveLeague() error:
   - Run full validation on all records (players, coaches, teams)
   - If validation errors, show validation error dialog, block save
   - Call appState.SaveProject()
   - Save all CSV files:
     - data.SavePlayers(path, appState.Players)
     - data.SaveCoaches(path, appState.Coaches)
     - data.SaveTeams(path, appState.Teams)
   - Update project LastModified timestamp
   - Save .fof9proj file
   - Set appState.IsModified = false
   - Update status bar: "Saved successfully"

3. Handle errors:
   - If write fails, show error dialog, don't clear IsModified flag

4. Test: Open league, edit a player, click File > Save, verify file is updated on disk

Users can now save changes.
```

**Deliverable**: Save League functionality

---

#### Step 46: Save As
**Goal**: Implement File > Save As

**Prompt**:
```
Implement Save As in main window:

1. Update MainWindow:
   - Enable "File > Save As" menu item
   - On select, show save dialog (dialog.ShowFileSave)
   - Filter to *.fof9proj files
   - Suggest default name: [leagueName].fof9proj
   - On file selected, call saveLeagueAs(filePath)

2. Implement saveLeagueAs(filePath string) error:
   - Create new project folder structure at new location
   - Copy all CSV files to new location
   - Update project file paths
   - Save .fof9proj to new location
   - Update appState.Project.Identifier and paths
   - Set IsModified = false
   - Update status bar

3. Test: Open league, click File > Save As, enter new path, verify new project folder is created with all files

Users can save copies of leagues.
```

**Deliverable**: Save As functionality

---

#### Step 47: Unsaved Changes Prompt
**Goal**: Prompt user before closing unsaved league

**Prompt**:
```
Implement unsaved changes handling:

1. Update MainWindow:
   - Override window close handler: window.SetCloseIntercept(func() { ... })
   - In close handler:
     - If appState.IsModified:
       - Show confirmation dialog: "Save changes to [league name]?"
       - Buttons: [Save] [Don't Save] [Cancel]
       - If Save: call saveLeague(), then close
       - If Don't Save: close without saving
       - If Cancel: don't close
     - Else: close immediately

2. Update openLeague:
   - Before loading new league, check IsModified
   - If true, show save prompt
   - If user cancels, abort open

3. Test: Open league, edit player, click X to close window, verify save prompt appears. Click Cancel, verify window stays open. Click Don't Save, verify window closes.

Unsaved changes are protected.
```

**Deliverable**: Unsaved changes prompt

---

### Phase 9: New League Wizard

#### Step 48: Wizard Dialog Foundation
**Goal**: Create multi-step wizard dialog

**Prompt**:
```
Create wizard dialog in internal/ui/dialogs/wizard.go:

1. Define NewLeagueWizard struct:
   - dialog *dialog.CustomDialog
   - steps []WizardStep
   - currentStep int
   - data map[string]interface{} (stores user input across steps)

2. Define WizardStep interface:
   - GetTitle() string
   - GetContent() fyne.CanvasObject
   - Validate() error
   - OnNext() error

3. Implement NewNewLeagueWizard(parent fyne.Window) *NewLeagueWizard:
   - Create dialog
   - Initialize steps (stubs for now)
   - Set size 600x400
   - Return wizard

4. Implement navigation:
   - ShowStep(step int) - displays current step content
   - NextStep() - validates current step, moves to next
   - PrevStep() - moves to previous step
   - Finish() - validates all, creates league

5. Add buttons: [< Back] [Next >] [Cancel] [Finish]
   - Finish button only enabled on last step

6. Test: Show wizard dialog, verify navigation works

Wizard framework is ready for steps.
```

**Deliverable**: internal/ui/dialogs/wizard.go with wizard framework

---

#### Step 49: Wizard Step 1: League Identity
**Goal**: First wizard step for naming league

**Prompt**:
```
Implement Step 1 of wizard:

1. Create WizardStepIdentity struct (implements WizardStep):
   - leagueNameEntry *widget.Entry
   - identifierEntry *widget.Entry
   - folderEntry *widget.Entry (path)
   - folderButton *widget.Button (browse button)

2. Implement GetTitle() string:
   - Returns "League Identity"

3. Implement GetContent() fyne.CanvasObject:
   - Form with:
     - Label: "League Name"
     - Entry: leagueNameEntry
     - Label: "Identifier" (auto-filled from name)
     - Entry: identifierEntry (read-only or editable)
     - Label: "Save Location"
     - Entry: folderEntry
     - Button: [Browse...] to select folder
   - OnChanged for leagueNameEntry:
     - Auto-fill identifierEntry with sanitized name (lowercase, no spaces)

4. Implement Validate() error:
   - Check leagueName is not empty
   - Check folder path is valid and writable
   - Return error if invalid

5. Implement OnNext() error:
   - Store values in wizard.data["leagueName"], wizard.data["identifier"], wizard.data["savePath"]

6. Add to wizard steps in NewNewLeagueWizard

7. Test: Open wizard, enter league name, verify identifier auto-fills. Click Next, verify validation works.

Step 1 is complete.
```

**Deliverable**: Wizard Step 1: League Identity

---

#### Step 50: Wizard Step 2: League Settings
**Goal**: Configure league settings (schedule, salary cap, etc.)

**Prompt**:
```
Implement Step 2 of wizard:

1. Create WizardStepSettings struct:
   - scheduleIDEntry *widget.Entry
   - baseYearEntry *widget.Entry
   - salaryCapEntry *widget.Entry
   - minimumEntry *widget.Entry
   - (other salary minimum entries)

2. Implement GetContent():
   - Form with all league info fields from spec.md section 3.3.4
   - Pre-filled with defaults from models.NewDefaultLeagueInfo(2024)

3. Implement Validate():
   - Check baseYear is 1900-2199
   - Check salaryCap > 0
   - Check salary minimums > 0

4. Implement OnNext():
   - Store all values in wizard.data["leagueInfo"]

5. Test: Navigate to Step 2, verify defaults are pre-filled. Click Next, verify validation.

Step 2 is complete.
```

**Deliverable**: Wizard Step 2: League Settings

---

#### Step 51: Wizard Step 3: Data Source
**Goal**: Choose where to import data from

**Prompt**:
```
Implement Step 3 of wizard:

1. Create WizardStepDataSource struct:
   - importOption *widget.RadioGroup (3 options)
   - defaultPathEntry *widget.Entry
   - defaultPathButton *widget.Button
   - customPathEntry *widget.Entry
   - customPathButton *widget.Button

2. Implement GetContent():
   - Radio group with options:
     - "Import from default 2024 game data"
     - "Start with empty/minimal data"
     - "Import from another custom league"
   - If option 1 selected, show path entry with default: C:\Program Files (x86)\Steam\steamapps\common\Front Office Football Nine\default_data
   - Browse button to change path
   - If option 3 selected, show path entry for .fof9proj file

3. Implement Validate():
   - If option 1: check default_data folder exists
   - If option 3: check .fof9proj file exists

4. Implement OnNext():
   - Store selected option and path in wizard.data["dataSource"]

5. Test: Select each option, verify path entry shows/hides. Click Browse, verify file picker works.

Step 3 is complete.
```

**Deliverable**: Wizard Step 3: Data Source

---

#### Step 52: Wizard Step 4: Entity Selection
**Goal**: Choose which entity types to import

**Prompt**:
```
Implement Step 4 of wizard:

1. Create WizardStepEntitySelection struct:
   - checkboxes map[string]*widget.Check (entity type -> checkbox)
   - recordCounts map[string]int (entity type -> count)
   - warningLabel *widget.Label

2. Implement GetContent():
   - If importing from data source (Step 3):
     - Load CSV files from selected source
     - Count records in each file
     - Show checkboxes:
       - ☑ Players (2,847 records)
       - ☑ Coaches (165 records)
       - ☑ Teams (32 records)
       - etc.
   - [Select All] [Deselect All] buttons
   - Warning label at bottom (initially hidden)

3. Implement dependency checking:
   - If Players selected but not Teams:
     - Show warning: "⚠ Players reference Teams. Consider importing Teams."
   - Similar for Coaches

4. Implement OnNext():
   - Store selected entity types in wizard.data["selectedEntities"]

5. Test: Select some entities, verify warnings appear. Click Select All, verify all checked.

Step 4 is complete.
```

**Deliverable**: Wizard Step 4: Entity Selection

---

#### Step 53: Wizard Step 5: Summary and Creation
**Goal**: Review settings and create league

**Prompt**:
```
Implement Step 5 of wizard and finish logic:

1. Create WizardStepSummary struct:
   - summaryLabel *widget.Label

2. Implement GetContent():
   - Display summary of all settings:
     - League Name: [name]
     - Base Year: [year]
     - Data Source: [source]
     - Entities to Import: [list]
   - Use formatted text (widget.NewRichTextFromMarkdown)

3. Implement Finish() in NewLeagueWizard:
   - Create project folder structure at wizard.data["savePath"]
   - Create data/ and reference/ subdirectories
   - Create .fof9proj file with settings
   - If importing, copy CSV files from source
   - If starting empty, create empty CSV files with headers
   - Copy reference data (cities.csv, colleges.csv)
   - Show success dialog
   - Load newly created project in main window

4. Wire "Finish" button in wizard:
   - On click, call Finish()
   - Close wizard dialog
   - Call mainWindow.openLeague(newProjectPath)

5. Test: Complete all wizard steps, click Finish, verify project folder is created, league loads in main window

New League Wizard is complete.
```

**Deliverable**: Wizard Step 5 and league creation logic

---

### Phase 10: Search and Filtering

#### Step 54: Quick Search
**Goal**: Add search bar above list view

**Prompt**:
```
Add quick search to list view:

1. Update ListView struct:
   - Add searchEntry *widget.Entry
   - Add originalData [][]string (unfiltered data)
   - Add filteredData [][]string (currently displayed)

2. Update NewListView:
   - Create search entry above table
   - Placeholder text: "Search..."
   - OnChanged callback:
     - Filter originalData by search text
     - Search across all columns
     - Case-insensitive
     - Update filteredData
     - Refresh table

3. Update SetData:
   - Stores data in originalData
   - If search text is empty, filteredData = originalData
   - Else, applies search filter

4. Implement search logic:
   - For each row, check if any cell contains search text
   - If yes, include in filteredData

5. Test: Open Players list, type "John" in search box, verify list filters to only players with "John" in any field

Quick search enables fast record finding.
```

**Deliverable**: Quick search in ListView

---

#### Step 55: Advanced Filters UI
**Goal**: Add advanced filter controls

**Prompt**:
```
Add advanced filters to list view:

1. Create FilterBar struct in internal/ui/widgets/filter_bar.go:
   - container *fyne.Container
   - filterRows []FilterRow
   - onApply func() (callback)

2. Define FilterRow struct:
   - fieldSelect *widget.Select (choose field)
   - operatorSelect *widget.Select (=, !=, >, <, contains)
   - valueEntry *widget.Entry
   - removeButton *widget.Button

3. Implement NewFilterBar(fields []string, onApply func()) *FilterBar:
   - Create [+] button to add filter row
   - Create [Apply] button to apply filters
   - Create [Clear] button to clear all filters

4. Implement AddFilterRow():
   - Creates new FilterRow
   - Adds to container
   - Wire remove button to remove row

5. Update ListView:
   - Add filterBar above search entry
   - Initially collapsed (show [Advanced Filters] button to expand)

6. Test: Open Players list, click Advanced Filters, add filter row, verify UI appears

Advanced filter UI is ready.
```

**Deliverable**: Filter bar UI

---

#### Step 56: Apply Filters
**Goal**: Apply advanced filters to data

**Prompt**:
```
Implement filter logic:

1. Update FilterBar:
   - Implement GetFilters() []Filter method
     - Returns slice of Filter structs with field, operator, value

2. Define Filter struct:
   - Field string
   - Operator string ("=", "!=", ">", "<", "contains")
   - Value string

3. Update ListView:
   - Add ApplyFilters(filters []Filter) method
     - For each row in originalData:
       - Check if row matches all filters (AND logic)
       - If yes, include in filteredData
     - Refresh table

4. Implement filter matching logic:
   - matchesFilter(row []string, filter Filter, columns []string) bool
     - Find column index for filter.Field
     - Get cell value
     - Compare based on operator:
       - "=": cell == value
       - "!=": cell != value
       - ">": convert to int, cell > value
       - "<": convert to int, cell < value
       - "contains": strings.Contains(cell, value)
     - Return true if matches

5. Wire FilterBar onApply callback:
   - Get filters from FilterBar
   - Call ListView.ApplyFilters(filters)

6. Test: Add filter "Position = RB", click Apply, verify only running backs shown

Advanced filters now work.
```

**Deliverable**: Filter application logic

---

#### Step 57: Filter Presets - Save
**Goal**: Allow saving filter configurations

**Prompt**:
```
Implement filter preset saving:

1. Create FilterPreset struct in internal/state/preferences.go:
   - Name string
   - EntityType string ("Players", "Coaches", "Teams")
   - Filters []Filter
   - IsDefault bool

2. Add to Preferences:
   - FilterPresets []FilterPreset

3. Update FilterBar:
   - Add [Save Preset] button
   - On click, show dialog:
     - Entry: "Preset name"
     - Checkbox: "Set as default"
     - [Save] [Cancel] buttons
   - On save:
     - Create FilterPreset with current filters
     - Add to appState.Preferences.FilterPresets
     - Save preferences to disk

4. Test: Apply filters, click Save Preset, enter name "Starting QBs", verify preset is saved

Users can save filter configurations.
```

**Deliverable**: Filter preset saving

---

#### Step 58: Filter Presets - Load
**Goal**: Load saved filter presets

**Prompt**:
```
Implement filter preset loading:

1. Update FilterBar:
   - Add preset dropdown above filter rows
   - Populate with saved presets for current entity type
   - On select:
     - Clear current filter rows
     - Load preset filters
     - Populate filter rows with preset values
     - Automatically apply filters

2. Implement LoadPreset(preset FilterPreset):
   - Clears current filters
   - For each filter in preset:
     - Add filter row
     - Set field, operator, value
   - Apply filters

3. Add default preset handling:
   - In showPlayersList/showCoachesList/showTeamsList:
     - Check if there's a default preset for this entity type
     - If yes, auto-load and apply

4. Test: Save a preset, close app, reopen, verify preset appears in dropdown. Select preset, verify filters are applied.

Users can quickly load saved filters.
```

**Deliverable**: Filter preset loading

---

#### Step 59: Filter Presets - Manage
**Goal**: Edit and delete presets

**Prompt**:
```
Add filter preset management:

1. Create PresetManagerDialog in internal/ui/dialogs/filters.go:
   - Show list of all presets for current entity type
   - For each preset:
     - [Edit] button - opens dialog to edit preset
     - [Delete] button - confirms and deletes
     - [Export] button - saves preset to .json file
   - [Import] button - loads preset from .json file
   - [Close] button

2. Implement Edit:
   - Show dialog with preset name and filters
   - Allow editing
   - Save changes to preferences

3. Implement Delete:
   - Show confirmation: "Delete preset '[name]'?"
   - If confirmed, remove from preferences
   - Save preferences

4. Implement Export/Import:
   - Export: Marshal preset to JSON, save to file
   - Import: Load JSON file, add to presets

5. Add "Manage Presets..." button to FilterBar

6. Test: Open preset manager, delete a preset, verify it's removed. Export preset, import in another league, verify it works.

Users can manage their filter presets.
```

**Deliverable**: Filter preset management

---

### Phase 11: Settings and Preferences

#### Step 60: Settings Dialog Foundation
**Goal**: Create settings dialog with tabs

**Prompt**:
```
Create settings dialog in internal/ui/dialogs/settings.go:

1. Define SettingsDialog struct:
   - dialog *dialog.CustomDialog
   - tabs *container.AppTabs
   - preferences *state.Preferences

2. Implement NewSettingsDialog(parent fyne.Window, prefs *state.Preferences) *SettingsDialog:
   - Create tabbed container with tabs:
     - General
     - Appearance
     - Editor
     - Import/Export
   - Add [Save] [Cancel] [Apply] buttons
   - Return dialog

3. Implement tabs (stubs for now):
   - Each tab returns fyne.CanvasObject with placeholder text

4. Wire buttons:
   - Save: apply changes, save preferences, close dialog
   - Cancel: discard changes, close dialog
   - Apply: apply changes, save preferences, keep dialog open

5. Add to Edit menu:
   - Settings (Ctrl+,)
   - On select, show SettingsDialog

6. Test: Open Edit > Settings, verify tabbed dialog appears

Settings dialog framework is ready.
```

**Deliverable**: internal/ui/dialogs/settings.go with tabbed dialog

---

#### Step 61: General Settings Tab
**Goal**: Configure data paths and updates

**Prompt**:
```
Implement General settings tab:

1. Create GeneralSettingsTab struct:
   - defaultDataFolderEntry *widget.Entry
   - defaultSaveLocationEntry *widget.Entry
   - checkUpdatesCheck *widget.Check

2. Implement GetContent() fyne.CanvasObject:
   - Form with:
     - Label: "Default data folder"
     - Entry + Browse button
     - Label: "Default save location"
     - Entry + Browse button
     - Checkbox: "Check for updates on startup"

3. Wire to preferences:
   - Load current values from prefs
   - On Save, update prefs with new values

4. Test: Open settings, change default data folder, click Save, verify preference is saved

General settings tab is complete.
```

**Deliverable**: General settings tab

---

#### Step 62: Appearance Settings Tab
**Goal**: Configure theme, font, etc.

**Prompt**:
```
Implement Appearance settings tab:

1. Create AppearanceSettingsTab struct:
   - themeSelect *widget.Select (Light, Dark, System)
   - fontFamilySelect *widget.Select
   - fontSizeSelect *widget.Select (8-16)

2. Implement GetContent():
   - Form with theme, font family, font size selectors
   - Load current values from preferences

3. Wire Apply button:
   - When theme changes, apply immediately:
     - app.Settings().SetTheme(selectedTheme)
   - When font changes, update app font settings

4. Test: Change theme to Dark, click Apply, verify UI switches to dark mode

Appearance settings tab is complete.
```

**Deliverable**: Appearance settings tab

---

#### Step 63: Editor Settings Tab
**Goal**: Configure auto-save and new record defaults

**Prompt**:
```
Implement Editor settings tab:

1. Create EditorSettingsTab struct:
   - newRecordDefaultsSelect *widget.Select (Random / John Doe)
   - autoSaveCheck *widget.Check
   - autoSaveIntervalSelect *widget.Select (1, 5, 10, 30 minutes)

2. Implement GetContent():
   - Form with:
     - Select: "New record defaults"
     - Checkbox: "Auto-save enabled"
     - Select: "Auto-save interval" (disabled if auto-save unchecked)

3. Wire to preferences:
   - Load/save auto-save settings
   - Load/save new record defaults

4. Test: Enable auto-save, set interval to 5 minutes, click Save, verify preference is saved

Editor settings tab is complete.
```

**Deliverable**: Editor settings tab

---

#### Step 64: Import/Export Settings Tab
**Goal**: Configure CSV encoding and format options

**Prompt**:
```
Implement Import/Export settings tab:

1. Create ImportExportSettingsTab struct:
   - encodingSelect *widget.Select (UTF-8, UTF-16, etc.)
   - numberFormatSelect *widget.Select (US, EU)

2. Implement GetContent():
   - Form with encoding and number format selectors
   - For this version, just UTF-8 and US format (others can be placeholders)

3. Wire to preferences

4. Test: Open settings, verify tab displays

Import/Export settings tab is complete.
```

**Deliverable**: Import/Export settings tab

---

### Phase 12: Auto-Save

#### Step 65: Auto-Save Timer
**Goal**: Periodically save to temp location

**Prompt**:
```
Implement auto-save in internal/state/appstate.go:

1. Add to AppState:
   - autoSaveTimer *time.Timer
   - autoSavePath string

2. Implement StartAutoSave():
   - Check if auto-save is enabled in preferences
   - Get interval from preferences (e.g., 5 minutes)
   - Create timer: time.NewTimer(interval)
   - In timer callback:
     - If IsModified:
       - Call SaveAutoSave()
     - Reset timer

3. Implement SaveAutoSave() error:
   - Create auto-save path: AppData\Local\fof9editor\autosave\[leagueName]_[timestamp].autosave
   - Marshal current state (Players, Coaches, Teams) to JSON
   - Write to auto-save file
   - Log success

4. Implement StopAutoSave():
   - Cancel timer

5. Update MainWindow:
   - In openLeague, call appState.StartAutoSave()
   - In closeLeague, call appState.StopAutoSave()

6. Test: Open league, edit player, wait for auto-save interval, verify auto-save file is created

Auto-save periodically saves progress.
```

**Deliverable**: Auto-save timer

---

#### Step 66: Auto-Save Cleanup
**Goal**: Remove old auto-save files

**Prompt**:
```
Implement auto-save cleanup:

1. In internal/state/appstate.go:
   - Implement CleanupAutoSaves()
     - List all files in auto-save directory
     - Parse timestamps from filenames
     - Sort by timestamp (newest first)
     - Keep last 5 files
     - Delete older files

2. Call CleanupAutoSaves():
   - After successful manual save (in saveLeague)
   - On app startup (in main.go)

3. Test: Create multiple auto-save files, call CleanupAutoSaves, verify only 5 most recent are kept

Old auto-saves are automatically cleaned up.
```

**Deliverable**: Auto-save cleanup

---

#### Step 67: Crash Recovery
**Goal**: Detect auto-save on startup and offer recovery

**Prompt**:
```
Implement crash recovery:

1. In cmd/fof9editor/main.go:
   - On startup, check for auto-save files in AppData\Local\fof9editor\autosave\
   - If found, show recovery dialog:
     - "Auto-saved data found for [league name] from [timestamp]. Recover?"
     - [Recover] [Ignore] [Delete]

2. Implement LoadAutoSave(filePath string) error in AppState:
   - Read auto-save JSON file
   - Unmarshal to Players, Coaches, Teams
   - Load into AppState
   - Set IsModified = true (user must explicitly save)
   - Open project referenced in auto-save

3. Wire dialog buttons:
   - Recover: call LoadAutoSave, open league, show success message
   - Ignore: delete auto-save file, continue normal startup
   - Delete: delete auto-save file

4. Test: Simulate crash (kill app during editing), restart app, verify recovery dialog appears. Click Recover, verify data is restored.

Crash recovery protects work.
```

**Deliverable**: Crash recovery on startup

---

#### Step 68: Auto-Save Status Indicator
**Goal**: Show last auto-save time in status bar

**Prompt**:
```
Add auto-save indicator to status bar:

1. Update StatusBar:
   - Add autoSaveLabel *widget.Label (far right section)
   - Implement SetAutoSaveStatus(time time.Time):
     - Formats time as "Auto-saved: 2:35 PM"
     - Updates label

2. Update AppState:
   - In SaveAutoSave, after successful save:
     - Call statusBar.SetAutoSaveStatus(time.Now())

3. Update MainWindow:
   - Pass statusBar reference to AppState

4. Test: Open league, wait for auto-save, verify status bar shows "Auto-saved: [time]"

Users can see when last auto-save occurred.
```

**Deliverable**: Auto-save status in status bar

---

### Phase 13: Polish and Testing

#### Step 69: Add New Record
**Goal**: Implement Add New button in forms

**Prompt**:
```
Implement Add New functionality:

1. Update FormView:
   - Add [Add New] button to button bar
   - Wire to onAddNew callback

2. In MainWindow, for each entity type:
   - showPlayersList: add onAddNew callback:
     - Generate new Player with defaults (from pkg/utils/defaults.go)
     - Add to appState.Players
     - Set SelectedRecordIndex to new player
     - Show form with new player
     - Set IsModified = true
   - Similar for coaches and teams

3. Create pkg/utils/defaults.go:
   - Implement GenerateDefaultPlayer(baseYear int, useRandom bool) *models.Player
     - Uses defaults from spec.md section 12.1
     - If useRandom, generates random name/position/etc.
   - Implement GenerateDefaultCoach(baseYear int, useRandom bool) *models.Coach
   - Implement GenerateDefaultTeam(baseYear int) *models.Team

4. Wire to preferences:
   - Check preferences.NewRecordDefaults ("random" or "johndoe")
   - Pass useRandom accordingly

5. Test: Click Add New in Players section, verify new player appears in form with defaults

Users can add new records.
```

**Deliverable**: Add New functionality

---

#### Step 70: Delete Record
**Goal**: Implement delete functionality

**Prompt**:
```
Implement Delete functionality:

1. Update FormView:
   - Wire [Delete] button to onDelete callback

2. In MainWindow:
   - onDelete callback for players:
     - Show confirmation dialog: "Delete [player name]?"
     - If confirmed:
       - Remove player from appState.Players
       - Set IsModified = true
       - Navigate to next player (or previous if last)
       - Refresh list view
   - Similar for coaches and teams

3. Add special handling for teams:
   - Before deleting team, check if any players/coaches assigned to it
   - If yes, show error dialog: "Cannot delete team. [X] players and [Y] coaches assigned."
   - Block deletion

4. Test: Select player, click Delete, confirm, verify player is removed. Try to delete team with players, verify error message.

Users can delete records with safety checks.
```

**Deliverable**: Delete functionality

---

#### Step 71: Column Configuration
**Goal**: Allow customizing visible columns

**Prompt**:
```
Implement column configuration:

1. Create ColumnConfigDialog in internal/ui/dialogs/columns.go:
   - Show all available columns with checkboxes
   - Show current visible columns in separate list
   - Drag to reorder visible columns
   - [Reset to Defaults] button

2. Update ListView:
   - Add [Columns] button above table
   - On click, show ColumnConfigDialog
   - On dialog close, update visible columns
   - Save to preferences

3. Implement in ColumnConfigDialog:
   - GetVisibleColumns() []string
   - SetVisibleColumns(columns []string)
   - OnSave: update preferences, refresh table

4. Update showPlayersList/showCoachesList/showTeamsList:
   - Load visible columns from preferences
   - Pass to ListView

5. Test: Click Columns button, uncheck "Height", click Save, verify Height column is hidden. Reopen, verify setting persists.

Users can customize column visibility.
```

**Deliverable**: Column configuration

---

#### Step 72: Color Picker Widget
**Goal**: Add color picker for team colors

**Prompt**:
```
Create color picker widget:

1. Create internal/ui/widgets/color_picker.go:
   - ColorPicker struct with:
     - colorRect *canvas.Rectangle (shows current color)
     - rgbLabels (show R/G/B values)
     - onChanged func(r, g, b int) callback

2. Implement NewColorPicker(r, g, b int, onChanged func(int, int, int)) *ColorPicker:
   - Create button that opens color picker dialog
   - Show current color in button (colored rectangle)
   - On click, open Fyne's color picker dialog
   - On color selected, update rectangle, call onChanged with RGB values

3. Update showTeamForm:
   - For PrimaryColor and SecondaryColor fields:
     - Use ColorPicker widget
     - On color change, update team's RGB fields

4. Test: Edit team, click Primary Color, select red, verify RGB values update to (255, 0, 0)

Team colors can be selected visually.
```

**Deliverable**: Color picker widget

---

#### Step 73: Keyboard Shortcuts
**Goal**: Implement keyboard shortcuts

**Prompt**:
```
Add keyboard shortcuts to MainWindow:

1. Define shortcuts in internal/ui/mainwindow.go:
   - Ctrl+N: New League (show wizard)
   - Ctrl+O: Open League
   - Ctrl+S: Save
   - Ctrl+Shift+S: Save As
   - Ctrl+F: Focus search box
   - Ctrl+Plus: Add New
   - Delete: Delete record
   - Alt+Left: Previous
   - Alt+Right: Next

2. Use window.Canvas().AddShortcut() for each:
   - Example: window.Canvas().AddShortcut(&desktop.CustomShortcut{
       KeyName: fyne.KeyS,
       Modifier: fyne.KeyModifierControl,
     }, func() { saveLeague() })

3. Update Help menu:
   - Add "Keyboard Shortcuts" menu item
   - Show dialog with list of shortcuts

4. Test: Press Ctrl+S, verify save is triggered. Press Alt+Right, verify next record loads.

Keyboard shortcuts improve productivity.
```

**Deliverable**: Keyboard shortcuts

---

#### Step 74: Enhanced Tooltips
**Goal**: Add sticky tooltips with documentation

**Prompt**:
```
Implement enhanced tooltips:

1. Create internal/ui/widgets/tooltip.go:
   - EnhancedTooltip struct with:
     - content *widget.RichText (documentation text)
     - popup *widget.PopUp
     - hoverTimer *time.Timer
     - persistTimer *time.Timer
     - closeButton *widget.Button

2. Implement NewEnhancedTooltip(content string) *EnhancedTooltip:
   - Create rich text with formatted documentation
   - Create popup (initially hidden)
   - Add [X] close button

3. Implement Show() and Hide():
   - Show displays popup near parent widget
   - Hide dismisses popup

4. Implement hover behavior:
   - On hover over help icon [?]:
     - Start 500ms timer
     - On timer expire, show tooltip
   - After 2 seconds of display:
     - Make tooltip "sticky" (doesn't hide on mouse out)
     - Show close button
   - On close button click, hide tooltip

5. Update FormView:
   - Add help icon [?] next to each field
   - On hover, show EnhancedTooltip with field documentation

6. Load documentation from example_players.txt:
   - Parse documentation file on startup
   - Build map: fieldName -> description
   - Use in tooltips

7. Test: Hover over help icon next to "Position" field, wait 500ms, verify tooltip appears. Wait 2s, verify tooltip becomes sticky and shows close button.

Enhanced tooltips provide in-context help.
```

**Deliverable**: Enhanced tooltip widget

---

#### Step 75: Final Testing and Bug Fixes
**Goal**: Integration testing and polish

**Prompt**:
```
Perform comprehensive testing and bug fixes:

1. Create integration test suite:
   - Test full workflow:
     - Open wizard
     - Create new league with default data
     - Open Players section
     - Edit a player
     - Save
     - Close and reopen
     - Verify changes persisted
   - Test validation:
     - Create player with invalid data
     - Verify save is blocked
     - Fix errors
     - Verify save succeeds
   - Test auto-save:
     - Edit data
     - Wait for auto-save
     - Kill app
     - Restart
     - Verify recovery works

2. Fix identified bugs:
   - Memory leaks
   - UI glitches
   - Data loss scenarios
   - Performance issues

3. Add error logging:
   - Log all errors to AppData\Local\fof9editor\logs\app.log
   - Include timestamps and stack traces

4. Update README.md:
   - Add user documentation
   - Installation instructions
   - Usage guide
   - Troubleshooting section

5. Build release:
   - Create Windows installer (optional)
   - Test on clean Windows 10/11 machines

Final application is tested and polished.
```

**Deliverable**: Tested, polished application ready for release

---

## Build and Deployment

### Local Build
```bash
# Build for Windows
make build

# Run
./bin/fof9editor.exe

# Test
make test
```

### GitHub Actions
Push to GitHub triggers:
- Automated build on Windows
- Run tests on Ubuntu
- Upload artifacts
- Create release (on tags)

### Packaging
- Single .exe (no installer required for portable use)
- Optional: Create Windows installer with Inno Setup or WiX

---

## Testing Strategy

### Unit Tests
- All validation rules
- CSV parsing/writing
- Data transformations
- ID generation
- Default value generation

### Integration Tests
- Open/save workflows
- Wizard completion
- Import from default data
- Search and filter
- Auto-save and recovery

### Manual Tests
- UI interactions
- Theme switching
- Keyboard shortcuts
- Large datasets (5000+ players)
- Error handling

---

## Summary

This plan breaks down the specification into 75 small, incremental steps across 13 phases. Each step:
- Builds on previous steps
- Has clear deliverables
- Can be implemented safely
- Integrates immediately (no orphaned code)
- Includes testing guidance

The plan follows best practices:
- Clean architecture (models, data, ui, validation separated)
- Test-driven development
- Incremental feature delivery
- Continuous integration

Total estimated time: 8-10 weeks for one developer working through all phases systematically.
