# Front Office Football Nine CSV Editor - Specification

## 1. Overview

### 1.1 Project Description
A desktop application for Windows 10/11 that enables users to create and edit custom leagues for the video game "Front Office Football Nine". The application provides a comprehensive interface for managing players, coaches, teams, and league settings through CSV file editing with built-in validation and documentation.

### 1.2 Technology Stack
- **Language**: Go (Golang)
- **UI Framework**: Fyne (native desktop UI)
- **Target Platform**: Windows 10/11
- **Distribution**: Desktop application (.exe)

### 1.3 Key Features
- Create complete custom leagues with wizard-driven setup
- Edit players, coaches, teams, and league configuration
- Import data from default game files or other custom leagues
- Real-time and save-time validation with auto-fix suggestions
- Search, filter, and navigate records efficiently
- Auto-save with crash recovery
- Customizable UI (themes, columns, filters)

---

## 2. User Workflows

### 2.1 Primary Workflow
1. User launches application
2. User creates new league via wizard OR opens existing league project
3. User selects entity type (Players, Coaches, Teams, League Info, etc.) from sidebar
4. User views list of records with search/filter capabilities
5. User selects record to edit in split-view form
6. User makes changes with real-time validation feedback
7. User saves changes (with validation enforcement)
8. Repeat steps 3-7 as needed

### 2.2 League Creation Workflow
1. Launch New League Wizard
2. **Step 1**: League Identity - name league and choose save location
3. **Step 2**: League Settings - configure SCHEDULEID, BASE_YEAR, SALARYCAP, salary minimums
4. **Step 3**: Initial Data Source - choose to import from default 2024 data, start blank, or import from another custom league
5. **Step 4**: Select Entity Types - choose which data to import (Players, Coaches, Teams, etc.) with checkboxes
6. **Step 5**: Summary - review settings
7. Click "Finish" to create project structure with .fof9proj file and CSV files

---

## 3. Data Model

### 3.1 Project Structure
```
MyCustomLeague/
├── MyCustomLeague.fof9proj          # JSON project file
├── data/
│   ├── MyCustomLeague_info.csv      # League configuration
│   ├── MyCustomLeague_players.csv   # Player data
│   ├── MyCustomLeague_coaches.csv   # Coach data
│   ├── team_info.csv                # Team data
│   ├── team_colors.csv              # Team colors
│   └── ... (other CSV files)
└── reference/                        # Read-only reference data
    ├── cities.csv
    ├── colleges.csv
    └── area.csv
```

### 3.2 Project File Format (.fof9proj)
JSON format containing:
```json
{
  "version": "1.0",
  "leagueName": "MyCustomLeague",
  "identifier": "MyCustomLeague",
  "created": "2024-10-12T14:30:00Z",
  "lastModified": "2024-10-12T15:45:00Z",
  "baseYear": 2024,
  "dataPath": "./data/",
  "referencePath": "./reference/",
  "csvFiles": {
    "info": "data/MyCustomLeague_info.csv",
    "players": "data/MyCustomLeague_players.csv",
    "coaches": "data/MyCustomLeague_coaches.csv",
    "teams": "data/team_info.csv",
    "teamColors": "data/team_colors.csv"
  },
  "userPreferences": {
    "lastOpenedSection": "Players",
    "columnConfigs": {...},
    "filterPresets": {...}
  }
}
```

### 3.3 Entity Types

#### 3.3.1 Players (from xxxx_players.csv)
**Key Fields**:
- PLAYERID (auto-generated, unique, ≥1000)
- LASTNAME, FIRSTNAME
- TEAM (references team_info.csv)
- POSITION_KEY (2-28, see position enum)
- UNIFORM (0-99)
- HEIGHT, WEIGHT, HANDSIZE, ARMLENGTH
- BIRTHMONTH, BIRTHDAY, BIRTHYEAR
- CITYID (references cities.csv), COLLEGEID (references colleges.csv)
- YEARENTRY, ROUNDDRAFTED, SELECTIONDRAFTED, ORIGINALTEAM
- EXPERIENCE (0-23)
- OVERALLRATING (0-10)
- Contract fields: SALARYYEARS (0-5), SALARYYEAR1-5, BONUSYEAR1-5
- Skill attributes (all -1 or 0-250): SKILL_SPEED, SKILL_POWER, etc.
- BASE_YEAR (determines which draft class)

**Position Key Enum**:
```
2  = Running Back
3  = Fullback
4  = Tight End
5  = Flanker
6  = Split End
7  = Left Tackle
8  = Left Guard
9  = Center
10 = Right Guard
11 = Right Tackle
12 = Punter
13 = Kicker
14 = Defensive Left End
15 = Defensive Left Tackle
16 = Defensive Nose Tackle
17 = Defensive Right Tackle
18 = Defensive Right End
19 = Strong-Side Linebacker
20 = Strong Inside Linebacker
21 = Middle Linebacker
22 = Weak Inside Linebacker
23 = Weak-Side Linebacker
24 = Left Cornerback
25 = Right Cornerback
26 = Strong Safety
27 = Free Safety
28 = Long Snapper
```

**Skill Attributes** (grouped by category):
- **Ball Carrier**: HOLE_RECOGNITION, ELUSIVENESS, SECURE_HANDLING
- **Receiver**: CATCH_HANDS, ADJUST_TO_BALL, ROUTE_RUNNING, CATCH_IN_TRAFFIC, DEFEAT_BLOCKERS
- **Offensive Line**: RUN_BLOCK_TECHNIQUE, PASS_BLOCK_TECHNIQUE, BLOCKING_STRENGTH, SCHEME_ACQUISITION
- **Backfield**: BLITZ_PICKUP
- **Kicker**: FIELD_GOAL_ACCURACY, FIELD_GOAL_DISTANCE, KICKOFF_HANG_TIME
- **Punter**: PUNT_DISTANCE, PUNT_HANG_TIME, PUNT_DIRECTIONAL
- **Defensive**: RUN_DEFENSE, PASS_RUSH_TECHNIQUE, PASS_RUSH_STRENGTH, PASS_DEFENSE_MAN, PASS_DEFENSE_PHYSICAL, PASS_DEFENSE_ZONE, PASS_DEFENSE_HANDS, DEFENSIVE_DIAGNOSIS
- **Special Teams**: SPECIAL_TEAMS, PUNT_RETURNS, KICK_RETURNS, LONG_SNAPPING, KICK_HOLDING
- **General**: SKILL_SPEED, SKILL_POWER, ENDURANCE

#### 3.3.2 Coaches (from xxxx_coaches.csv)
**Key Fields**:
- LASTNAME, FIRSTNAME
- BIRTHMONTH, BIRTHDAY, BIRTHYEAR
- BIRTHCITY (text), CITYID (references cities.csv)
- COLLEGE (text), COLLEGEID (references colleges.csv)
- TEAM (references team_info.csv)
- POSITION (0=Head Coach, 1=Offensive Coordinator, 2=Defensive Coordinator, 3=Special Teams Coordinator, 4=Strength & Conditioning)
- POSITIONGROUP (position-specific role)
- OFFENSIVESTYLE (0-6)
- DEFENSIVESTYLE (0-4)
- PAYSCALE (salary in units of $10,000)

#### 3.3.3 Teams (from team_info.csv)
**Key Fields**:
- YEAR, TEAMID, TEAMNAME, NICKNAME, ABBREVIATION
- CONFERENCE, DIVISION
- CITY (references cities.csv)
- PRIMARYRED, PRIMARYGREEN, PRIMARYBLUE
- SECONDARYRED, SECONDARYGREEN, SECONDARYBLUE
- ROOF (0=outdoor, 1=dome, 2=retractable)
- TURF (0=grass, 1=artificial, 2=hybrid)
- BUILT (year), CAPACITY, LUXURY (luxury boxes), CONDITION
- ATTENDANCE, SUPPORT
- Future stadium: PLAN, COMPLETED, FUTURE, FUTURENAME, FUTUREABBR, FUTUREROOF, FUTURETURF, FUTURECAP, FUTURELUXURY, TEAMCONTRIBUTION

#### 3.3.4 League Info (from xxxx_info.csv)
**Key Fields**:
- SCHEDULEID (format: x_y_z where x=teams, y=divisions, z=games, must match game's league_info.csv)
- BASE_YEAR (1900-2199)
- SALARYCAP (in units of $100,000)
- MINIMUM (rookie minimum salary, units of $10,000)
- SALARY1 (1 year experience minimum)
- SALARY2 (2 years experience minimum)
- SALARY3 (3 years experience minimum)
- SALARY45 (4-5 years experience minimum)
- SALARY789 (7-9 years experience minimum)
- SALARY10 (10+ years experience minimum)

#### 3.3.5 Reference Data (Read-Only)
- **Cities** (cities.csv): CITYID, city name, state/country
- **Colleges** (colleges.csv): COLLEGEID, college name
- **Area** (area.csv): geographical data

---

## 4. User Interface Design

### 4.1 Window Layout

```
┌─────────────────────────────────────────────────────────┐
│ Menu Bar: File | Edit | View | Tools | Help             │
├───────────┬─────────────────────────────────────────────┤
│           │                                             │
│ Sidebar   │  Main Content Area                          │
│           │  ┌─────────────────────────────────────┐    │
│ ☐ League  │  │ Search/Filter Bar                   │    │
│   Info    │  │ ┌─────────────┐ [Apply] [Clear]    │    │
│ ☐ Players │  │ │Quick search │ Advanced Filters    │    │
│ ☐ Coaches │  └─────────────────────────────────────┘    │
│ ☐ Teams   │  ┌─────────────────────────────────────┐    │
│ ☐ Cities* │  │ LIST VIEW (Top)                     │    │
│ ☐ Colleges│  │ ┌───────┬────────┬──────┬─────────┐ │    │
│           │  │ │Name   │Team    │Pos   │Rating  │ │    │
│           │  │ │────────────────────────────────│ │    │
│           │  │ │Smith  │Team 1  │RB    │7       │ │    │
│           │  │ │Jones  │Team 2  │QB    │8       │◄──   │
│           │  │ └───────┴────────┴──────┴─────────┘ │    │
│           │  └─────────────────────────────────────┘    │
│           │  ════════════════════════════════════════   │
│           │  ┌─────────────────────────────────────┐    │
│           │  │ FORM VIEW (Bottom - Resizable)      │    │
│           │  │ ┌─────────────────────────────────┐ │    │
│           │  │ │ John Smith                  [?] │ │    │
│           │  │ │ First Name: [John________]      │ │    │
│           │  │ │ Last Name:  [Smith_______]  [?] │ │    │
│           │  │ │ Team:       [Team 1 ▼]      [?] │ │    │
│           │  │ │ Position:   [RB ▼]          [?] │ │    │
│           │  │ │                                 │ │    │
│           │  │ │ ▼ Physical Attributes           │ │    │
│           │  │ │ ▼ Contract Details              │ │    │
│           │  │ │ ▼ Ball Carrier Skills           │ │    │
│           │  │ │                                 │ │    │
│           │  │ │ [< Prev] [Next >] [Save] [Del] │ │    │
│           │  │ └─────────────────────────────────┘ │    │
│           │  └─────────────────────────────────────┘    │
├───────────┴─────────────────────────────────────────────┤
│ Status: Saved | 1,247 Players | MyLeague2024 | Modified │
└─────────────────────────────────────────────────────────┘

* Read-only sections shown with different styling
```

### 4.2 Main Window Components

#### 4.2.1 Sidebar Navigation
- Vertical list of entity type sections
- Icons for each section
- Active section highlighted
- Read-only sections visually distinguished (grayed out or with lock icon)
- Sections:
  - League Info
  - Players
  - Coaches
  - Teams
  - Cities (read-only)
  - Colleges (read-only)
  - Area (read-only)

#### 4.2.2 Search/Filter Bar
- **Quick Search**: Single text input that searches across key fields (name, team, position)
- **Advanced Filters**: Button that expands filter options
  - Field-specific filters (dropdown for field selection)
  - Comparison operators (=, !=, >, <, >=, <=, contains)
  - Value input (text or dropdown depending on field)
  - Add/remove filter rows
- **Filter Preset Dropdown**: Load/save/manage filter presets
- **Apply Button**: Apply current filters
- **Clear Button**: Clear all filters

#### 4.2.3 List View (Top Panel)
- Table/grid showing records with columns
- Default columns per entity type:
  - **Players**: Name (First + Last), Team, Position, Overall Rating, Experience, Height, Weight, College, Draft Info (Round/Pick), Contract (Years/Total)
  - **Coaches**: Name (First + Last), Team, Position, Position Group, Offensive Style, Defensive Style, Payscale, Experience
  - **Teams**: Name, City, Stadium, Conference, Division, Colors (visual preview)
- Click column header to sort (ascending/descending)
- Click row to load in form view
- Inline action buttons per row: [Edit] [Delete]
- **Column Configuration Button**: Opens dialog to show/hide columns, drag to reorder
- Scrollable for large datasets
- Row highlighting for selected record

#### 4.2.4 Form View (Bottom Panel)
- Displays selected record in editable form
- Fields grouped by category with collapsible sections:
  - **Players**: Basic Info, Physical Attributes, Birth Info, College Info, Draft History, Experience & Career, Contract Details, Ball Carrier Skills, Receiver Skills, Offensive Line Skills, Kicker Skills, Punter Skills, Defensive Skills, Special Teams Skills, General Attributes
  - **Coaches**: Basic Info, Birth Info, College Info, Position & Role, Coaching Styles, Compensation
  - **Teams**: Team Identity, Stadium Info, Colors, Financial Data, Future Stadium Plans
  - **League Info**: Schedule & Structure, Financial Settings, Salary Minimums
- Each field has:
  - Label (human-readable from CSV header)
  - Input control (text field, searchable dropdown, color picker)
  - Help icon [?] for tooltip
  - Validation indicator (red border if invalid)
- Validation errors displayed below form
- Navigation buttons: [< Previous] [Next >]
- Action buttons: [Save] [Delete] [Add New]
- Adjustable split ratio with draggable divider between list and form

#### 4.2.5 Status Bar (Bottom)
Four sections (left to right):
1. **Operation Status**: Current action/message (e.g., "File saved successfully", "Validating...", "3 validation errors")
2. **Record Count**: Total and filtered count (e.g., "1,247 players" or "142 of 1,247 players (filtered)")
3. **Current File/League**: Project name (e.g., "League: MyCustomLeague2024")
4. **Unsaved Changes**: Indicator (e.g., "* Modified" or "Last saved: 2:35 PM")

### 4.3 Dialog Windows

#### 4.3.1 New League Wizard
Multi-step modal dialog with progress indicator (Step X of 6):

**Step 1: Welcome**
- Intro text explaining league creation
- [Next] button

**Step 2: League Identity**
- League Name: [text input]
- Identifier: [text input] (auto-filled from name, used for file naming)
- Save Location: [folder picker]
- [< Back] [Next >] [Cancel]

**Step 3: League Settings**
- Form with league info fields (SCHEDULEID, BASE_YEAR, SALARYCAP, salary minimums)
- Default values pre-filled
- Inline validation
- [< Back] [Next >] [Cancel]

**Step 4: Initial Data Source**
- Radio buttons:
  - ○ Import from default 2024 game data
    - Path: [C:\Program Files (x86)\Steam\steamapps\common\Front Office Football Nine\default_data] [Browse]
  - ○ Start with empty/minimal data
  - ○ Import from another custom league [Browse]
- [< Back] [Next >] [Cancel]

**Step 5: Select Entity Types**
- Checkboxes for each entity type with record counts:
  - ☑ Players (2,847 records)
  - ☑ Coaches (165 records)
  - ☑ Teams (32 records)
  - ☑ Team Colors
  - etc.
- [Select All] [Deselect All]
- Dependency warnings (e.g., "⚠ Players reference Teams. Consider importing Teams as well.")
- [< Back] [Next >] [Cancel]

**Step 6: Summary**
- Review all settings
- "The following league will be created:" summary text
- [< Back] [Finish] [Cancel]

On [Finish]:
- Create project folder structure
- Create .fof9proj file
- Copy/create CSV files
- Show success message
- Open newly created league

#### 4.3.2 Settings Dialog
Tabbed dialog:

**Tab: General**
- Default data folder: [path] [Browse]
- Default save location: [path] [Browse]
- Check for updates on startup: [☑]

**Tab: Appearance**
- Theme: [Light / Dark / System ▼]
- Font Family: [Segoe UI ▼]
- Font Size: [10 ▼] (8-16)

**Tab: Editor**
- New record defaults: [Random values / "John Doe" placeholders ▼]
- Auto-save enabled: [☑]
- Auto-save interval: [5 minutes ▼] (1, 5, 10, 30 minutes)

**Tab: Import/Export**
- CSV encoding: [UTF-8 ▼]
- Number format: [US (1,234.56) ▼]

[Save] [Cancel] [Apply]

#### 4.3.3 Column Configuration Dialog
Per entity type:

**Available Columns (left panel)**
- ☐ PLAYERID
- ☐ LASTNAME
- ☑ FIRSTNAME
- ☑ TEAM
- etc.

**Visible Columns (right panel)**
- Drag to reorder
- Remove button per column

[Reset to Defaults] [OK] [Cancel]

#### 4.3.4 Filter Preset Management
- **Load Preset**: Dropdown in filter bar
- **Save Preset**:
  - Name: [text input]
  - Set as default: [☐]
  - [Save] [Cancel]
- **Manage Presets**:
  - List of saved presets
  - [Edit] [Delete] [Export] [Import] buttons per preset
  - [Close]

#### 4.3.5 Validation Error Dialog
Modal showing detailed errors when save fails:
- Error count header: "Cannot save: 3 validation errors found"
- List of errors with record references:
  - "Player #1234 'John Smith': CITYID 99999 not found in cities.csv"
  - "Coach #45 'Mike Johnson': BIRTHYEAR 2010 inconsistent with EXPERIENCE 15 years"
  - etc.
- [Auto-fix All] [Close]

#### 4.3.6 Confirmation Dialogs
Standard Windows-style dialogs:
- **Unsaved Changes**: "Save changes to MyLeague before closing?" [Save] [Don't Save] [Cancel]
- **Delete Confirmation**: "Are you sure you want to delete John Smith?" [Yes] [No]
- **Team Delete Warning**: "Team 'Dallas Cowboys' has 53 assigned players and 12 coaches. Cannot delete." [OK]

#### 4.3.7 About Dialog
- Application name: "Front Office Football Nine CSV Editor"
- Version number
- Copyright/license info
- [Check for Updates] [Close]

### 4.4 Tooltips
Enhanced tooltip behavior:
- Trigger: Hover over help icon [?] next to field
- Delay: 500ms before showing
- Content:
  - Field name (human-readable)
  - Description from documentation
  - Valid range/values with examples
  - Example: "POSITION_KEY: Player's primary position. Valid range: 2-28. Examples: 2=Running Back, 3=Fullback, 4=Tight End, ..."
- Persistence: After 2 seconds of continuous hover, tooltip becomes sticky
- Dismiss: Close button [X] on sticky tooltip (does not dismiss on mouse out)

### 4.5 Validation Feedback

#### Real-time (on blur)
- Invalid field: Red border around input
- Icon indicator: ⚠ next to field
- Brief tooltip on hover showing error

#### On save
- All validation errors shown below form
- Grouped by category
- Each error shows:
  - Field name
  - Error message
  - [Auto-fix] button if applicable (e.g., "Calculate appropriate birth year")
- Save blocked until all errors resolved
- Status bar shows error count

#### Auto-fix Suggestions
For cross-field inconsistencies:
- "Player age 20 with 10 years experience" → [Suggest birth year based on experience]
- "Coach hired before birth year" → [Adjust hire year]
- "Future stadium completion year before current year" → [Set to current year + 1]

---

## 5. Features and Functionality

### 5.1 File Operations

#### 5.1.1 New League
- Launch New League Wizard (see 4.3.1)
- Create project folder with .fof9proj and CSV files
- Import data from selected source
- Load newly created league in main window

#### 5.1.2 Open League
- File → Open or Ctrl+O
- Browse for .fof9proj file
- If unsaved changes exist, prompt to save first
- Load all CSV files referenced in project file
- Validate file integrity on open
- If validation errors found, show detailed error dialog
- If corrupted/missing files, show error and abort open
- Update recent leagues list

#### 5.1.3 Save
- File → Save or Ctrl+S
- Run full validation
- If errors found, show validation dialog and block save
- Write all modified CSV files atomically:
  1. Write to temp file
  2. Verify write success
  3. Replace original file
  4. Update .fof9proj metadata (lastModified timestamp)
- Update status bar: "File saved successfully"
- Clear unsaved changes indicator

#### 5.1.4 Save As
- File → Save As or Ctrl+Shift+S
- Prompt for new location and league name
- Create new project folder
- Copy all CSV files to new location
- Create new .fof9proj file
- Keep new project open

#### 5.1.5 Recent Leagues
- File → Recent submenu
- Show last 5 opened leagues
- Click to open

#### 5.1.6 Close League
- File → Close
- If unsaved changes, prompt to save
- Clear main window
- Return to welcome state or prompt to open/create league

### 5.2 Editing Operations

#### 5.2.1 Add New Record
- Click [Add New] button in form view
- Generate new unique ID (PLAYERID, etc.)
- Pre-fill with defaults:
  - TEAM = 0 (free agent)
  - All skill attributes = -1
  - BIRTHYEAR = BASE_YEAR - 22 (for players)
  - Random name (or "John Doe" based on settings)
  - Random position
  - EXPERIENCE = 0
  - Other fields: appropriate defaults
- Load blank form for editing
- Mark as unsaved
- Add to list view (at top or bottom, depending on sort)

#### 5.2.2 Edit Record
- Click record in list view OR click [Edit] inline button
- Load record data into form view
- Enable editing
- Real-time validation on field blur
- Mark as unsaved when any field changes

#### 5.2.3 Delete Record
- Click [Delete] button in form view OR [Delete] inline button in list view
- Show confirmation dialog: "Are you sure you want to delete [name]?"
- If confirmed, remove from dataset
- Mark as unsaved
- If deleted from form view, navigate to next record
- If deleted from list view, remove row

#### 5.2.4 Navigate Records
- [< Previous] [Next >] buttons in form view
- Navigate only through filtered results (if filter active)
- Cycle through records (wrap around from last to first)
- If on unsaved record, prompt to save first
- Keyboard shortcuts:
  - Alt+Left: Previous
  - Alt+Right: Next

#### 5.2.5 Duplicate Record
- Feature for future consideration (not in initial spec)

### 5.3 Search and Filtering

#### 5.3.1 Quick Search
- Single text input in filter bar
- Searches across key fields (varies by entity type):
  - **Players**: FIRSTNAME, LASTNAME, TEAM (name), POSITION (name)
  - **Coaches**: FIRSTNAME, LASTNAME, TEAM (name), POSITION (name)
  - **Teams**: TEAMNAME, NICKNAME, CITY
- Case-insensitive
- Partial matching (contains)
- Updates list view in real-time as user types

#### 5.3.2 Advanced Filters
- Click [Advanced Filters] to expand
- Add filter rows with:
  - Field dropdown (all fields available)
  - Operator dropdown (=, !=, >, <, >=, <=, contains, not contains)
  - Value input (text field or dropdown based on field type)
- Multiple filters combined with AND logic
- [Add Filter] button to add rows
- [X] button per row to remove
- Click [Apply] to apply filters to list view
- Click [Clear] to remove all filters and show all records

#### 5.3.3 Filter Presets
- **Save Current Filter**:
  - Click [Save Filter] button in filter bar
  - Enter preset name
  - Option to set as default (applies automatically when opening this entity type)
  - Save to project file under userPreferences
- **Load Filter Preset**:
  - Dropdown in filter bar showing saved presets
  - Select preset to apply filters
- **Manage Presets**:
  - Click [Manage] button next to preset dropdown
  - Dialog shows all presets for current entity type
  - [Edit] to modify filter criteria
  - [Delete] to remove preset
  - [Export] to save preset to .json file
  - [Import] to load preset from .json file

### 5.4 Column Management

#### 5.4.1 Show/Hide Columns
- Click [Columns] button above list view
- Open Column Configuration Dialog (see 4.3.3)
- Check/uncheck columns to show/hide
- Changes apply immediately to list view
- Save to project file under userPreferences

#### 5.4.2 Reorder Columns
- In Column Configuration Dialog, drag visible columns to reorder
- Changes apply immediately to list view
- Save to project file

#### 5.4.3 Resize Columns
- Drag column borders in list view header
- Auto-fit: Double-click column border to fit content width

#### 5.4.4 Sort Columns
- Click column header to sort ascending
- Click again to sort descending
- Click third time to remove sort (return to default order)
- Visual indicator (arrow) showing sort direction

### 5.5 Validation System

#### 5.5.1 Field-Level Validation (Real-time)
Validate on field blur:
- **Range checks**: POSITION_KEY (2-28), OVERALLRATING (0-10), EXPERIENCE (0-23), etc.
- **Required fields**: LASTNAME, FIRSTNAME, POSITION must not be empty
- **Data type**: Numeric fields must be numbers
- **Reference checks**: CITYID must exist in cities.csv, COLLEGEID in colleges.csv, TEAM in team_info.csv
- **Date validation**: Valid month (1-12), day (1-31), year ranges
- **Format validation**: Height, weight within reasonable ranges for position

Show validation feedback:
- Red border on invalid field
- Error message in tooltip on hover
- Error listed below form

#### 5.5.2 Cross-Field Validation (Real-time)
Check for inconsistencies:
- Birth year vs. experience (warn if age < 18 + experience)
- Draft year vs. birth year (warn if drafted before age 18)
- Contract years vs. salary fields (if SALARYYEARS=3, must have SALARYYEAR1-3 populated)
- Future stadium COMPLETED year must be > current league year

Show warnings with [Auto-fix] button:
- "Player age 20 with 10 years experience. [Calculate birth year]"
- Click [Auto-fix] to apply suggestion

#### 5.5.3 Save-Time Validation (Enforced)
Before saving, validate entire dataset:
- All field-level validations must pass
- All cross-field validations must pass (warnings can be ignored, errors block save)
- Unique constraints (PLAYERID must be unique)
- Referential integrity (no dangling references to non-existent teams, cities, colleges)

If validation fails:
- Show Validation Error Dialog (see 4.3.5)
- List all errors with record references
- Offer [Auto-fix All] if auto-fixes available
- Block save until resolved

#### 5.5.4 Reference Data Validation
For fields referencing cities.csv, colleges.csv, team_info.csv:
- On load, build lookup tables
- On field edit, check if value exists in lookup
- If invalid reference found (e.g., opening old file with deleted city):
  - Show in dropdown as "99999 (Invalid - not found)"
  - Mark field with warning indicator
  - Allow editing to fix
  - Include in validation report

### 5.6 Auto-Save and Recovery

#### 5.6.1 Auto-Save Configuration
In Settings → Editor:
- Enable/disable auto-save
- Set interval (1, 5, 10, 30 minutes)
- Auto-save triggers only if changes made since last save

#### 5.6.2 Auto-Save Location
- Save to: `C:\Users\[username]\AppData\Local\fof9editor\autosave\`
- Filename: `[leagueName]_[timestamp].autosave`
- Keep last 5 auto-saves per league
- Clean up older auto-saves on successful manual save

#### 5.6.3 Auto-Save Format
- JSON format containing full in-memory dataset
- Includes all modified CSV data
- Includes metadata: original project path, auto-save timestamp

#### 5.6.4 Crash Recovery
On application startup:
- Check for auto-save files
- If found, show recovery dialog:
  - "Auto-saved data found for MyLeague from 2:35 PM. Recover?"
  - [Recover] [Ignore] [Delete]
- If [Recover]:
  - Load auto-save data
  - Mark as unsaved (user must explicitly save)
  - Delete auto-save file after successful manual save

### 5.7 Import/Export

#### 5.7.1 Import from Default Game Data
In New League Wizard (Step 4):
- Default path: `C:\Program Files (x86)\Steam\steamapps\common\Front Office Football Nine\default_data`
- [Browse] to change path
- On [Next], read selected CSV files from default_data folder
- Validate file format and required columns
- Import selected entity types (from Step 5)

#### 5.7.2 Import from Custom League
In New League Wizard (Step 4):
- [Browse] to select another .fof9proj file
- Read CSV files referenced in that project
- Validate compatibility:
  - Check BASE_YEAR difference (warn if > 10 years apart)
  - Check for missing reference data (cities, colleges)
- Import selected entity types (from Step 5)
- Preserve IDs from source (PLAYERID, TEAMID, etc.)

#### 5.7.3 Selective Import
In Step 5 of wizard:
- Checkboxes for each entity type
- Show record counts for each type
- [Select All] / [Deselect All] shortcuts
- Dependency warnings:
  - If Players selected but not Teams: "⚠ Players reference Teams. Consider importing Teams."
  - If Coaches selected but not Teams: "⚠ Coaches reference Teams. Consider importing Teams."

#### 5.7.4 Export (Future Feature)
- Export individual CSV files
- Export entire league as .zip
- Export to Excel format for external editing
- Not in initial specification

### 5.8 Team Management

#### 5.8.1 Edit Team Colors
- Color fields: PRIMARYRED/GREEN/BLUE, SECONDARYRED/GREEN/BLUE
- Input method: Color picker dialog
- Click color field to open picker
- Visual preview of selected colors
- Picker automatically converts to RGB values (0-255)

#### 5.8.2 Team ID Changes (Cascade Updates)
When a team's TEAMID is changed:
- Scan all players with TEAM = old TEAMID
- Update to new TEAMID
- Scan all coaches with TEAM = old TEAMID
- Update to new TEAMID
- Show confirmation: "Team ID changed. Updated 53 players and 12 coaches."

#### 5.8.3 Team Deletion Protection
When user attempts to delete a team:
- Scan for assigned players and coaches
- If any found, show error dialog:
  - "Cannot delete team 'Dallas Cowboys'. Team has 53 assigned players and 12 coaches."
  - "Reassign or delete these records first."
- Block deletion

#### 5.8.4 Future Stadium Validation
For future stadium fields (PLAN, COMPLETED, FUTURE, etc.):
- COMPLETED year must be >= current league BASE_YEAR
- If PLAN = 1, COMPLETED and FUTURE fields must be populated
- Warn if FUTURECAP < CAPACITY (downgrading stadium)

### 5.9 Theme and Appearance

#### 5.9.1 Theme Selection
In Settings → Appearance:
- Light theme: White backgrounds, dark text
- Dark theme: Dark backgrounds, light text
- System theme: Follow Windows system settings

Apply theme change immediately (no restart required)

#### 5.9.2 Font Customization
In Settings → Appearance:
- Font family: Dropdown of system-installed fonts
- Font size: 8-16 pt
- Apply immediately to all UI text (lists, forms, buttons)

### 5.10 Updates

#### 5.10.1 Check for Updates
On application startup:
- If "Check for updates on startup" enabled in Settings
- HTTP request to update server (URL TBD)
- Compare current version with latest version
- If newer version available:
  - Show notification dialog: "Version 1.2.0 is available. You are running 1.1.0. [Download] [Skip This Version] [Remind Later]"
  - [Download]: Open browser to download page
  - [Skip This Version]: Don't notify about this version again
  - [Remind Later]: Check again on next startup

#### 5.10.2 Manual Update Check
Help → Check for Updates:
- Same logic as startup check
- Show result even if up-to-date: "You are running the latest version (1.1.0)."

#### 5.10.3 Version Display
Help → About:
- Show current version number
- Show release date
- [Check for Updates] button

---

## 6. Data Validation Rules

### 6.1 Player Validation Rules

| Field | Rule | Error/Warning |
|-------|------|---------------|
| PLAYERID | Unique, integer ≥ 1000 | Error: "PLAYERID must be unique and ≥ 1000" |
| LASTNAME | Required, max 18 chars | Error: "Last name is required (max 18 chars)" |
| FIRSTNAME | Required, max 16 chars | Error: "First name is required (max 16 chars)" |
| TEAM | Must exist in team_info.csv or = 0 | Error: "Invalid team ID" |
| POSITION_KEY | 2-28 | Error: "Position must be 2-28" |
| UNIFORM | 0-99 | Error: "Uniform number must be 0-99" |
| HEIGHT | > 0 | Error: "Height must be > 0" |
| WEIGHT | > 0 | Error: "Weight must be > 0" |
| BIRTHMONTH | 1-12 | Error: "Month must be 1-12" |
| BIRTHDAY | 1-31 | Error: "Day must be 1-31" |
| BIRTHYEAR | Reasonable range (BASE_YEAR - 50 to BASE_YEAR - 18) | Warning: "Birth year seems unusual for league year" |
| CITYID | Must exist in cities.csv | Error: "Invalid city ID" |
| COLLEGEID | Must exist in colleges.csv or = 0 | Error: "Invalid college ID" |
| EXPERIENCE | 0-23 | Error: "Experience must be 0-23 years" |
| OVERALLRATING | 0-10 | Error: "Overall rating must be 0-10" |
| SALARYYEARS | 0-5 | Error: "Contract years must be 0-5" |
| Skill attributes | -1 or 0-250 | Error: "Skill must be -1 or 0-250" |
| Age vs. Experience | Age ≥ 18 + EXPERIENCE | Warning: "Player age inconsistent with experience [Auto-fix]" |
| Draft year vs. Birth year | YEARENTRY ≥ BIRTHYEAR + 21 | Warning: "Draft year too early for birth year [Auto-fix]" |

### 6.2 Coach Validation Rules

| Field | Rule | Error/Warning |
|-------|------|---------------|
| LASTNAME | Required, max 18 chars | Error: "Last name is required" |
| FIRSTNAME | Required, max 16 chars | Error: "First name is required" |
| BIRTHMONTH | 1-12 | Error: "Month must be 1-12" |
| BIRTHDAY | 1-31 | Error: "Day must be 1-31" |
| BIRTHYEAR | Reasonable range | Warning: "Birth year seems unusual" |
| CITYID | Must exist in cities.csv | Error: "Invalid city ID" |
| COLLEGEID | Must exist in colleges.csv or = 0 | Error: "Invalid college ID" |
| TEAM | Must exist in team_info.csv | Error: "Invalid team ID" |
| POSITION | 0-4 | Error: "Position must be 0-4" |
| OFFENSIVESTYLE | 0-6 | Error: "Offensive style must be 0-6" |
| DEFENSIVESTYLE | 0-4 | Error: "Defensive style must be 0-4" |
| PAYSCALE | > 0 | Error: "Payscale must be > 0" |

### 6.3 Team Validation Rules

| Field | Rule | Error/Warning |
|-------|------|---------------|
| TEAMID | Unique, integer > 0 | Error: "Team ID must be unique and > 0" |
| TEAMNAME | Required, max 50 chars | Error: "Team name is required" |
| NICKNAME | Required, max 50 chars | Error: "Nickname is required" |
| ABBREVIATION | Required, max 5 chars | Error: "Abbreviation is required (max 5 chars)" |
| CITY | Must exist in cities.csv | Error: "Invalid city ID" |
| PRIMARY/SECONDARY RGB | 0-255 for each R/G/B | Error: "RGB values must be 0-255" |
| ROOF | 0-2 | Error: "Roof type must be 0-2" |
| TURF | 0-2 | Error: "Turf type must be 0-2" |
| CAPACITY | > 0 | Error: "Capacity must be > 0" |
| COMPLETED | ≥ BUILT | Warning: "Future stadium completion before current year [Auto-fix]" |
| FUTURECAP | If PLAN=1, must be > 0 | Error: "Future capacity required when stadium plan active" |

### 6.4 League Info Validation Rules

| Field | Rule | Error/Warning |
|-------|------|---------------|
| SCHEDULEID | Must match entry in game's league_info.csv | Error: "Invalid schedule ID (must match game data)" |
| BASE_YEAR | 1900-2199 | Error: "Base year must be 1900-2199" |
| SALARYCAP | > 0 | Error: "Salary cap must be > 0" |
| MINIMUM | > 0 | Error: "Minimum salary must be > 0" |
| SALARY1-10 | > 0, ascending order | Warning: "Salary minimums should increase with experience" |

---

## 7. Technical Requirements

### 7.1 Performance

| Metric | Target |
|--------|--------|
| Application startup | < 2 seconds |
| Open league (5000 players) | < 3 seconds |
| Load entity type in list view | < 500 ms |
| Load record in form view | < 100 ms |
| Apply filter | < 500 ms |
| Save league | < 5 seconds |
| Search (quick search) | < 200 ms (feel instant) |

### 7.2 Data Capacity

| Entity | Expected Max | Performance Target |
|--------|-------------|-------------------|
| Players | 10,000 | No degradation |
| Coaches | 500 | No degradation |
| Teams | 100 | No degradation |
| Cities | 50,000 | No degradation |
| Colleges | 5,000 | No degradation |

### 7.3 Memory

- Maximum memory usage: 500 MB for typical league (2,000 players, 150 coaches, 32 teams)
- Graceful handling of low memory conditions

### 7.4 File Formats

#### CSV Format
- Encoding: UTF-8
- Line endings: Windows (CRLF)
- Delimiter: Comma (,)
- Quote character: Double quote (") for fields containing commas
- Header row: Required, exact column names as specified in game documentation

#### JSON Format (Project File)
- Encoding: UTF-8
- Format: Pretty-printed (indented) for human readability
- Schema version: Include version field for future compatibility

### 7.5 Error Handling

- All file I/O operations must handle errors gracefully
- Show detailed error messages with actionable information
- Never crash or lose data on error
- Log errors to application log file: `AppData\Local\fof9editor\logs\app.log`

### 7.6 Platform Requirements

- **OS**: Windows 10 (version 1809 or later), Windows 11
- **Architecture**: x64
- **.NET**: Not required (Go produces native binary)
- **Disk Space**: 50 MB application + user data
- **RAM**: 256 MB minimum, 512 MB recommended

---

## 8. User Experience Considerations

### 8.1 First-Time User Experience
On first launch:
1. Welcome screen explaining what the app does
2. [Create New League] button prominent
3. [Open Existing League] button
4. Link to documentation/help
5. "Show this on startup" checkbox

### 8.2 Keyboard Shortcuts

| Action | Shortcut |
|--------|----------|
| New League | Ctrl+N |
| Open League | Ctrl+O |
| Save | Ctrl+S |
| Save As | Ctrl+Shift+S |
| Close League | Ctrl+W |
| Exit Application | Alt+F4 |
| Find/Quick Search | Ctrl+F |
| Add New Record | Ctrl+Plus |
| Delete Record | Delete |
| Previous Record | Alt+Left |
| Next Record | Alt+Right |
| Settings | Ctrl+, |

### 8.3 Accessibility
- Keyboard navigation for all functions
- Tab order follows logical flow
- Focus indicators visible
- High contrast mode support (follows system theme)
- Screen reader compatibility (future enhancement)

### 8.4 Help and Documentation
- Help → View Documentation: Open user manual (PDF or web)
- Help → Keyboard Shortcuts: Show shortcut reference
- Help → About: Version info, credits, license
- Contextual help via tooltips (see 4.4)
- Documentation files (example_info.txt, example_players.txt) loaded and displayed in tooltips

---

## 9. Development Phases

### Phase 1: Core Application (MVP)
**Goal**: Basic functional editor

Features:
- Project file structure (.fof9proj, CSV files)
- Open/Save league
- Edit Players, Coaches, Teams, League Info
- Basic UI (sidebar, list view, form view, status bar)
- Field-level validation (ranges, required fields)
- Reference data (cities, colleges) read-only
- Searchable dropdowns for references
- Theme support (light/dark)
- Windows packaging

**Deliverable**: Working desktop app that can create and edit a basic custom league

### Phase 2: Advanced Editing
**Goal**: Enhanced productivity features

Features:
- New League Wizard
- Add/Delete records
- Column customization (show/hide, reorder)
- Collapsible form sections (group attributes)
- Cross-field validation with auto-fix
- Inline delete/edit buttons in list view
- Previous/Next navigation
- Unsaved changes prompts

**Deliverable**: Full-featured editor with complete workflow

### Phase 3: Search and Filtering
**Goal**: Powerful data navigation

Features:
- Quick search
- Advanced filters (multi-field, operators)
- Filter presets (save/load/manage)
- Sort by column
- Filter through filtered results with Previous/Next

**Deliverable**: Efficient data discovery and navigation

### Phase 4: Auto-Save and Recovery
**Goal**: Data safety

Features:
- Auto-save to temp location
- Configurable auto-save interval
- Crash recovery on startup
- Auto-save cleanup
- Status bar auto-save indicator

**Deliverable**: Robust data protection

### Phase 5: Import/Export
**Goal**: Flexible data management

Features:
- Import from default game data (with path detection)
- Import from other custom leagues
- Selective entity import with dependency warnings
- Compatibility validation

**Deliverable**: Easy league creation from existing data

### Phase 6: Polish and Settings
**Goal**: Professional finish

Features:
- Settings dialog (all preferences)
- Font customization
- Recent leagues list
- Update checker
- Enhanced tooltips (sticky, examples)
- Keyboard shortcuts
- About dialog
- Error logging

**Deliverable**: Production-ready application

### Phase 7: Advanced Features (Future)
- Export to Excel
- Batch edit (apply changes to multiple records)
- Duplicate record
- Undo/redo
- Team templates
- Custom validation rules
- Plugin system
- Localization

---

## 10. Non-Functional Requirements

### 10.1 Reliability
- No data loss under any circumstance
- Atomic file writes (never corrupt original file)
- Crash recovery via auto-save
- Validation prevents invalid data from being saved

### 10.2 Maintainability
- Clean code architecture (MVC or similar pattern)
- Modular design (separate concerns: UI, data, validation, file I/O)
- Unit tests for validation logic
- Integration tests for file operations
- Comments and documentation in code

### 10.3 Usability
- Intuitive UI following Windows conventions
- Consistent terminology with game documentation
- Clear error messages with actionable suggestions
- Responsive UI (no freezing during operations)
- Visual feedback for all actions

### 10.4 Security
- No network access except update checker (user can disable)
- No telemetry or data collection
- User data stays local
- No execution of arbitrary code from CSV files

### 10.5 Portability
- Single .exe file (no installer required for simple deployment)
- Optional installer for Start Menu integration, file associations
- Portable mode: Store settings in app directory instead of AppData

---

## 11. File Format Reference

### 11.1 CSV Encoding
- UTF-8 encoding
- Windows line endings (CRLF)
- Comma-separated values
- Quote fields containing commas or quotes
- Header row with exact column names (case-sensitive)

### 11.2 Project File Schema (.fof9proj)
```json
{
  "version": "1.0",
  "leagueName": "string",
  "identifier": "string",
  "created": "ISO8601 datetime",
  "lastModified": "ISO8601 datetime",
  "baseYear": "integer",
  "dataPath": "string (relative path)",
  "referencePath": "string (relative path)",
  "csvFiles": {
    "info": "string (file path)",
    "players": "string (file path)",
    "coaches": "string (file path)",
    "teams": "string (file path)",
    "teamColors": "string (file path)"
  },
  "userPreferences": {
    "lastOpenedSection": "string",
    "columnConfigs": {
      "Players": {
        "visible": ["FIRSTNAME", "LASTNAME", "TEAM", ...],
        "order": ["FIRSTNAME", "LASTNAME", ...]
      },
      "Coaches": { ... },
      ...
    },
    "filterPresets": {
      "Players": [
        {
          "name": "Starting QBs",
          "filters": [
            {"field": "POSITION_KEY", "operator": "=", "value": "1"}
          ],
          "isDefault": false
        },
        ...
      ],
      ...
    }
  }
}
```

### 11.3 Auto-Save File Schema
```json
{
  "version": "1.0",
  "originalProjectPath": "string (absolute path to .fof9proj)",
  "timestamp": "ISO8601 datetime",
  "data": {
    "players": [ {...}, {...}, ... ],
    "coaches": [ {...}, {...}, ... ],
    "teams": [ {...}, {...}, ... ],
    "leagueInfo": { ... }
  }
}
```

---

## 12. Default Values Reference

### 12.1 New Player Defaults
When creating new player (using "John Doe" mode):
```
PLAYERID: [Next available ID ≥ 1000]
LASTNAME: "Doe"
FIRSTNAME: "John"
TEAM: 0 (free agent)
POSITION_KEY: 2 (Running Back)
UNIFORM: 99
HEIGHT: 72 (6'0")
WEIGHT: 215
HANDSIZE: 0 (auto-generate)
ARMLENGTH: 0 (auto-generate)
BIRTHMONTH: 1
BIRTHDAY: 1
BIRTHYEAR: [BASE_YEAR - 22]
CITYID: 1 (first city in cities.csv)
COLLEGEID: 0 (No College)
YEARENTRY: [BASE_YEAR - 1]
ROUNDDRAFTED: 0 (undrafted)
SELECTIONDRAFTED: 0
SUPPLEMENTAL: 0
ORIGINALTEAM: 0
EXPERIENCE: 0
YEARSIGNED: [BASE_YEAR]
PLAYPERCENTAGE: 0
HALLOFFAMEPOINTS: 0
SALARYYEARS: 0
SALARYYEAR1-5: 0
BONUSYEAR1-5: 0
OVERALLRATING: 2 (starter quality)
All skill attributes: -1 (auto-generate)
BASE_YEAR: [Current league BASE_YEAR]
```

When using "Random values" mode:
- Names: Select randomly from pool of first/last names
- Position: Random from 2-28
- Birth date: Random within reasonable range
- Physical attributes: Random within position-appropriate ranges
- Other defaults: Same as "John Doe" mode

### 12.2 New Coach Defaults
```
LASTNAME: "Doe"
FIRSTNAME: "John"
BIRTHMONTH: 1
BIRTHDAY: 1
BIRTHYEAR: [BASE_YEAR - 45]
CITYID: 1
COLLEGEID: 0 (No College)
TEAM: 0 (free agent)
POSITION: 0 (Head Coach)
POSITIONGROUP: 10 (general)
OFFENSIVESTYLE: 0
DEFENSIVESTYLE: 0
PAYSCALE: 80 ($800,000)
```

### 12.3 New Team Defaults
```
YEAR: [Current league BASE_YEAR]
TEAMID: [Next available ID]
TEAMNAME: "New Team"
NICKNAME: "Expansion"
ABBREVIATION: "NEW"
CONFERENCE: 1
DIVISION: 1
CITY: 1 (first city)
PRIMARYRED: 0
PRIMARYGREEN: 0
PRIMARYBLUE: 0
SECONDARYRED: 255
SECONDARYGREEN: 255
SECONDARYBLUE: 255
ROOF: 0 (outdoor)
TURF: 0 (grass)
BUILT: [BASE_YEAR]
CAPACITY: 65000
LUXURY: 100
CONDITION: 10 (new)
ATTENDANCE: 0
SUPPORT: 0
PLAN: 0 (no future plans)
COMPLETED: 0
FUTURE: 0
... (future fields all 0/empty)
```

### 12.4 New League Info Defaults
```
SCHEDULEID: "32_8_17" (32 teams, 8 divisions, 17 games)
BASE_YEAR: 2024
SALARYCAP: 2000 ($200M)
MINIMUM: 70 ($700k)
SALARY1: 85 ($850k)
SALARY2: 100 ($1M)
SALARY3: 115 ($1.15M)
SALARY45: 130 ($1.3M)
SALARY789: 150 ($1.5M)
SALARY10: 180 ($1.8M)
```

---

## 13. Testing Requirements

### 13.1 Unit Tests
- Validation logic for all field types
- CSV parsing and writing
- JSON project file parsing and writing
- Reference data lookups
- ID generation (unique PLAYERID, etc.)
- Auto-fix suggestions

### 13.2 Integration Tests
- Open league workflow
- Save league workflow
- Create new league workflow
- Import from default data
- Import from custom league
- Auto-save and recovery
- Filter and search operations

### 13.3 UI Tests
- Navigation between sections
- Form editing and validation feedback
- List view sorting and filtering
- Column customization
- Theme switching

### 13.4 Manual Testing Scenarios
1. Create new league from default 2024 data
2. Edit 10 players, save, close, reopen, verify changes
3. Add new player with invalid data, verify validation blocks save
4. Delete player, verify cascade (not applicable for players, but test for teams)
5. Apply complex filters, navigate with Previous/Next
6. Crash app during editing, reopen, verify auto-save recovery
7. Change theme, font size, verify UI updates
8. Import league from another custom league, verify data integrity
9. Edit team colors with color picker, verify RGB values correct
10. Test all keyboard shortcuts

---

## 14. Future Enhancements

### 14.1 Potential Features for Later Versions
- **Export to Excel**: Edit in Excel, import back
- **Batch editing**: Select multiple records, apply same change
- **Undo/redo**: Track edit history
- **Compare leagues**: Side-by-side comparison of two league projects
- **Player progression simulator**: Simulate player development over years
- **Draft simulator**: Generate draft-eligible players
- **Team builder**: AI-assisted roster construction within salary cap
- **Custom validation rules**: User-defined validation logic
- **Plugin system**: Allow third-party extensions
- **Multi-language support**: Localization for non-English users
- **Cloud sync**: Save leagues to cloud storage (OneDrive, Google Drive)
- **Version control**: Track changes over time, revert to previous versions
- **Import from game saves**: Extract data from active game saves

### 14.2 Platform Expansion
- macOS version (using Fyne's cross-platform capabilities)
- Linux version
- Web version (lightweight online editor)

---

## 15. Constraints and Assumptions

### 15.1 Constraints
- Windows 10/11 only (initial version)
- No network features except update checker
- CSV format must match game's exact specification (can't deviate)
- Reference data (cities, colleges) read-only (can't modify game's data)
- Single-user application (no multi-user editing)

### 15.2 Assumptions
- Users have Front Office Football Nine installed (for default data import)
- Users understand basic football concepts (positions, roles)
- Users are comfortable with desktop applications (file management, etc.)
- Average league size: 2,000 players, 150 coaches, 32 teams
- CSV files are text-based and not too large (< 50 MB)
- Users have basic Windows proficiency

### 15.3 Known Limitations
- No real-time collaboration (one user per league at a time)
- No built-in backup (users should manually backup their leagues)
- No merge conflicts resolution (if editing same league in multiple locations)
- Performance degrades with very large datasets (> 10,000 players, untested)

---

## 16. Glossary

| Term | Definition |
|------|------------|
| BASE_YEAR | The starting year of the custom league |
| Draft Class | Group of players eligible for draft in a given year |
| Entity Type | Category of data (Players, Coaches, Teams, etc.) |
| Filter Preset | Saved search/filter configuration |
| Front Office Football Nine | The video game this editor supports |
| Fyne | Go UI framework for desktop applications |
| League Info | Configuration settings for a custom league |
| Overall Rating | Player's overall skill level (0-10) |
| Position Group | Subcategory of coaching position |
| Position Key | Numeric identifier for player position |
| Project File | .fof9proj file containing league metadata |
| Reference Data | Read-only lookup data (cities, colleges) |
| SCHEDULEID | Identifier for league structure (teams/divisions/games) |
| Skill Attributes | Individual player abilities (speed, power, etc.) |

---

## 17. Revision History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0 | 2024-10-12 | Initial | Complete specification based on 54-question brainstorming session |

---

## 18. Approval and Sign-Off

This specification document represents the complete requirements for the Front Office Football Nine CSV Editor application as of version 1.0.

**Prepared for**: Igor (User/Stakeholder)
**Prepared by**: Claude (AI Assistant)
**Date**: 2024-10-12

---

**End of Specification Document**
