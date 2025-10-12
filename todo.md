# FOF9 Editor - Implementation TODO

## Status Legend
- [ ] Not started
- [🔄] In progress
- [✅] Completed
- [⏸️] Blocked/Paused

---

## Phase 0: Project Setup

- [ ] **Step 1**: Initialize Go Project
  - [ ] Create go.mod
  - [ ] Create directory structure
  - [ ] Create basic main.go
  - [ ] Create README.md
  - [ ] Test compilation

- [ ] **Step 2**: Add Fyne Dependency
  - [ ] Install Fyne
  - [ ] Update main.go with Fyne window
  - [ ] Test window opens
  - [ ] Update README with dependencies

- [ ] **Step 3**: Setup Build and CI
  - [ ] Create Makefile
  - [ ] Create GitHub Actions workflow
  - [ ] Create .gitignore
  - [ ] Test local build

---

## Phase 1: Core Data Models

- [ ] **Step 4**: Project Model
  - [ ] Define Project struct
  - [ ] Add helper methods
  - [ ] Add JSON tags
  - [ ] Write tests

- [ ] **Step 5**: Player Model
  - [ ] Define Player struct with all fields
  - [ ] Add CSV tags
  - [ ] Add helper methods
  - [ ] Test

- [ ] **Step 6**: Coach Model
  - [ ] Define Coach struct
  - [ ] Add CSV tags and constants
  - [ ] Add helper methods
  - [ ] Test

- [ ] **Step 7**: Team Model
  - [ ] Define Team struct
  - [ ] Add CSV tags and constants
  - [ ] Add helper methods for colors
  - [ ] Test

- [ ] **Step 8**: League Info Model
  - [ ] Define LeagueInfo struct
  - [ ] Add validation helpers
  - [ ] Add default constructor
  - [ ] Test

---

## Phase 2: CSV I/O

- [ ] **Step 9**: CSV Reader Foundation
  - [ ] Create CSVReader struct
  - [ ] Implement ReadAll
  - [ ] Write unit tests with fixtures
  - [ ] Test error handling

- [ ] **Step 10**: Player CSV Loading
  - [ ] Implement LoadPlayers
  - [ ] Add mapRowToPlayer helper
  - [ ] Write tests
  - [ ] Test error handling

- [ ] **Step 11**: Coach and Team CSV Loading
  - [ ] Implement LoadCoaches
  - [ ] Implement LoadTeams
  - [ ] Add helper functions
  - [ ] Write tests

- [ ] **Step 12**: CSV Writer Foundation
  - [ ] Create CSVWriter struct
  - [ ] Implement WriteAll with atomic writes
  - [ ] Write tests
  - [ ] Test error handling

- [ ] **Step 13**: Player/Coach/Team CSV Saving
  - [ ] Implement SavePlayers
  - [ ] Implement SaveCoaches
  - [ ] Implement SaveTeams
  - [ ] Write integration test

---

## Phase 3: Basic UI Framework

- [ ] **Step 14**: Application State
  - [ ] Define AppState struct
  - [ ] Implement singleton pattern
  - [ ] Add LoadProject/SaveProject stubs
  - [ ] Write tests

- [ ] **Step 15**: Main Window Foundation
  - [ ] Create MainWindow struct
  - [ ] Implement basic window
  - [ ] Update main.go
  - [ ] Test window opens

- [ ] **Step 16**: Status Bar
  - [ ] Create StatusBar struct
  - [ ] Implement with 4 sections
  - [ ] Add update methods
  - [ ] Integrate in MainWindow
  - [ ] Test display

- [ ] **Step 17**: Sidebar Navigation
  - [ ] Create Sidebar struct
  - [ ] Implement section list
  - [ ] Add callback handling
  - [ ] Test navigation

- [ ] **Step 18**: Layout Integration
  - [ ] Wire sidebar, content, status bar
  - [ ] Implement updateContent
  - [ ] Test layout
  - [ ] Test section switching

- [ ] **Step 19**: Menu Bar
  - [ ] Create File menu
  - [ ] Create Edit menu
  - [ ] Create Help menu with About
  - [ ] Test menu actions

- [ ] **Step 20**: Theme Support
  - [ ] Create Preferences struct
  - [ ] Implement save/load preferences
  - [ ] Add theme switching
  - [ ] Test theme changes

---

## Phase 4: List View

- [ ] **Step 21**: List View Foundation
  - [ ] Create ListView struct with table
  - [ ] Implement SetData
  - [ ] Test with sample data

- [ ] **Step 22**: Player List View
  - [ ] Implement showPlayersList
  - [ ] Add helper functions for names
  - [ ] Integrate in updateContent
  - [ ] Test display

- [ ] **Step 23**: Coach and Team List Views
  - [ ] Implement showCoachesList
  - [ ] Implement showTeamsList
  - [ ] Test both displays

- [ ] **Step 24**: Column Sorting
  - [ ] Add sorting to ListView
  - [ ] Implement header click handling
  - [ ] Add visual indicators
  - [ ] Test sorting

- [ ] **Step 25**: List Selection Highlighting
  - [ ] Add selection highlighting
  - [ ] Implement Select method
  - [ ] Test highlighting

---

## Phase 5: Form View

- [ ] **Step 26**: Form View Foundation
  - [ ] Create FormView struct
  - [ ] Implement SetFields
  - [ ] Add buttons
  - [ ] Test with sample fields

- [ ] **Step 27**: Player Form View
  - [ ] Implement showPlayerForm
  - [ ] Wire callbacks
  - [ ] Test editing

- [ ] **Step 28**: Split View Layout
  - [ ] Implement vertical split
  - [ ] Update showPlayersList
  - [ ] Test split view

- [ ] **Step 29**: Form Field Groups
  - [ ] Add accordion groups
  - [ ] Update showPlayerForm
  - [ ] Test collapsible groups

- [ ] **Step 30**: Coach and Team Forms
  - [ ] Implement showCoachForm
  - [ ] Implement showTeamForm
  - [ ] Test both forms

- [ ] **Step 31**: Form Validation UI
  - [ ] Add error display to FormView
  - [ ] Implement SetErrors
  - [ ] Add red borders for errors
  - [ ] Test error display

- [ ] **Step 32**: Previous/Next Navigation
  - [ ] Wire navigation buttons
  - [ ] Implement navigation logic
  - [ ] Add keyboard shortcuts
  - [ ] Test navigation

---

## Phase 6: Validation

- [ ] **Step 33**: Validation Engine
  - [ ] Define ValidationError struct
  - [ ] Create Validator interface
  - [ ] Implement ValidationEngine
  - [ ] Write tests

- [ ] **Step 34**: Player Field Validation
  - [ ] Create PlayerValidator
  - [ ] Implement field rules
  - [ ] Write tests

- [ ] **Step 35**: Cross-Field Validation
  - [ ] Add age vs experience checks
  - [ ] Add draft year validation
  - [ ] Add contract validation
  - [ ] Write tests

- [ ] **Step 36**: Coach and Team Validation
  - [ ] Create CoachValidator
  - [ ] Create TeamValidator
  - [ ] Write tests

- [ ] **Step 37**: Real-Time Validation
  - [ ] Integrate validation in FormView
  - [ ] Add on-blur validation
  - [ ] Test feedback

- [ ] **Step 38**: Save-Time Validation
  - [ ] Add validation enforcement on save
  - [ ] Show error dialog if blocked
  - [ ] Test save blocking

---

## Phase 7: Reference Data

- [ ] **Step 39**: Reference Data Loader
  - [ ] Create ReferenceData struct
  - [ ] Implement LoadReferenceData
  - [ ] Add to AppState
  - [ ] Write tests

- [ ] **Step 40**: Searchable Dropdown Widget
  - [ ] Create SearchableSelect struct
  - [ ] Implement filtering
  - [ ] Test with large lists

- [ ] **Step 41**: Integrate Searchable Dropdowns
  - [ ] Update FormView for searchable_select
  - [ ] Use in player form for city/college/team
  - [ ] Test selection

- [ ] **Step 42**: Invalid Reference Handling
  - [ ] Handle missing references in dropdown
  - [ ] Add validation warnings
  - [ ] Test with invalid IDs

---

## Phase 8: File Operations

- [ ] **Step 43**: Project File I/O
  - [ ] Implement SaveProject
  - [ ] Implement LoadProject
  - [ ] Write tests
  - [ ] Update AppState

- [ ] **Step 44**: Open League
  - [ ] Enable File > Open menu
  - [ ] Implement openLeague
  - [ ] Handle errors
  - [ ] Test opening

- [ ] **Step 45**: Save League
  - [ ] Enable File > Save menu
  - [ ] Implement saveLeague
  - [ ] Add validation check
  - [ ] Test saving

- [ ] **Step 46**: Save As
  - [ ] Enable File > Save As menu
  - [ ] Implement saveLeagueAs
  - [ ] Test saving to new location

- [ ] **Step 47**: Unsaved Changes Prompt
  - [ ] Add close intercept
  - [ ] Show save prompt
  - [ ] Test close scenarios

---

## Phase 9: New League Wizard

- [ ] **Step 48**: Wizard Dialog Foundation
  - [ ] Create NewLeagueWizard struct
  - [ ] Implement step navigation
  - [ ] Add buttons
  - [ ] Test navigation

- [ ] **Step 49**: Wizard Step 1: League Identity
  - [ ] Create WizardStepIdentity
  - [ ] Implement form
  - [ ] Add validation
  - [ ] Test step

- [ ] **Step 50**: Wizard Step 2: League Settings
  - [ ] Create WizardStepSettings
  - [ ] Add league info fields
  - [ ] Test with defaults

- [ ] **Step 51**: Wizard Step 3: Data Source
  - [ ] Create WizardStepDataSource
  - [ ] Add radio options
  - [ ] Test path selection

- [ ] **Step 52**: Wizard Step 4: Entity Selection
  - [ ] Create WizardStepEntitySelection
  - [ ] Add checkboxes with counts
  - [ ] Show dependency warnings
  - [ ] Test selection

- [ ] **Step 53**: Wizard Step 5: Summary and Creation
  - [ ] Create WizardStepSummary
  - [ ] Implement Finish logic
  - [ ] Create project structure
  - [ ] Test complete workflow

---

## Phase 10: Search and Filtering

- [ ] **Step 54**: Quick Search
  - [ ] Add search entry to ListView
  - [ ] Implement filtering
  - [ ] Test search

- [ ] **Step 55**: Advanced Filters UI
  - [ ] Create FilterBar struct
  - [ ] Implement filter rows
  - [ ] Test UI

- [ ] **Step 56**: Apply Filters
  - [ ] Implement filter logic
  - [ ] Add filter matching
  - [ ] Test filtering

- [ ] **Step 57**: Filter Presets - Save
  - [ ] Create FilterPreset struct
  - [ ] Implement save preset
  - [ ] Test saving

- [ ] **Step 58**: Filter Presets - Load
  - [ ] Implement load preset
  - [ ] Add default preset handling
  - [ ] Test loading

- [ ] **Step 59**: Filter Presets - Manage
  - [ ] Create PresetManagerDialog
  - [ ] Add edit/delete/export/import
  - [ ] Test management

---

## Phase 11: Settings and Preferences

- [ ] **Step 60**: Settings Dialog Foundation
  - [ ] Create SettingsDialog with tabs
  - [ ] Add buttons
  - [ ] Test dialog

- [ ] **Step 61**: General Settings Tab
  - [ ] Implement GeneralSettingsTab
  - [ ] Add path settings
  - [ ] Test

- [ ] **Step 62**: Appearance Settings Tab
  - [ ] Implement AppearanceSettingsTab
  - [ ] Add theme/font controls
  - [ ] Test theme switching

- [ ] **Step 63**: Editor Settings Tab
  - [ ] Implement EditorSettingsTab
  - [ ] Add auto-save controls
  - [ ] Test

- [ ] **Step 64**: Import/Export Settings Tab
  - [ ] Implement ImportExportSettingsTab
  - [ ] Add encoding controls
  - [ ] Test

---

## Phase 12: Auto-Save

- [ ] **Step 65**: Auto-Save Timer
  - [ ] Implement StartAutoSave
  - [ ] Implement SaveAutoSave
  - [ ] Test auto-save

- [ ] **Step 66**: Auto-Save Cleanup
  - [ ] Implement CleanupAutoSaves
  - [ ] Test cleanup

- [ ] **Step 67**: Crash Recovery
  - [ ] Add recovery dialog on startup
  - [ ] Implement LoadAutoSave
  - [ ] Test recovery

- [ ] **Step 68**: Auto-Save Status Indicator
  - [ ] Add to status bar
  - [ ] Update on auto-save
  - [ ] Test display

---

## Phase 13: Polish and Testing

- [ ] **Step 69**: Add New Record
  - [ ] Implement Add New button
  - [ ] Create defaults generation
  - [ ] Test adding records

- [ ] **Step 70**: Delete Record
  - [ ] Implement Delete button
  - [ ] Add team deletion protection
  - [ ] Test deletion

- [ ] **Step 71**: Column Configuration
  - [ ] Create ColumnConfigDialog
  - [ ] Implement show/hide/reorder
  - [ ] Test configuration

- [ ] **Step 72**: Color Picker Widget
  - [ ] Create ColorPicker widget
  - [ ] Integrate in team form
  - [ ] Test color selection

- [ ] **Step 73**: Keyboard Shortcuts
  - [ ] Add all shortcuts from spec
  - [ ] Create shortcuts help dialog
  - [ ] Test shortcuts

- [ ] **Step 74**: Enhanced Tooltips
  - [ ] Create EnhancedTooltip widget
  - [ ] Load documentation from files
  - [ ] Implement sticky behavior
  - [ ] Test tooltips

- [ ] **Step 75**: Final Testing and Bug Fixes
  - [ ] Create integration test suite
  - [ ] Fix identified bugs
  - [ ] Add error logging
  - [ ] Update documentation
  - [ ] Build release

---

## Current Status

**Phase**: Not Started
**Current Step**: Step 1 - Initialize Go Project
**Last Updated**: 2024-10-12

---

## Notes

- Each step should be completed and tested before moving to the next
- Update this file as you progress through the steps
- Mark items with [🔄] when starting work
- Mark items with [✅] when completed and tested
- Add notes about blockers or issues encountered
