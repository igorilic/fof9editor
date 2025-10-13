# Changelog

All notable changes to FOF9 Editor will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Nothing yet

## [0.3.0] - 2025-10-13

### Added
- Phase 5: Form View (Steps 26-28, 32)
  - FormView component for editing records with text, number, and select fields
  - Player form with key fields (name, team, position, uniform, overall rating)
  - Split view layout showing list (40%) and form (60%) simultaneously
  - Draggable divider between list and form
  - Previous/Next navigation buttons to move between records
  - Save button to persist changes to player data
  - Delete button to remove players from the list
  - Form updates in real-time when selecting players from list
- Phase 8: File Operations (Steps 43-47)
  - Project file I/O with atomic writes using temp files
  - SaveProject/LoadProject functions with comprehensive validation
  - File > Open Project menu item loads .fof9proj files and all CSV data
  - File > Save menu item saves project with LastModified timestamp
  - File > Save As menu item saves to new location with extension validation
  - Window close intercept prompts to save unsaved changes
  - Unsaved changes prompt when opening another project
  - Status bar updates to reflect saved/unsaved state
- Cross-compilation support
  - Makefile targets: build-windows (MinGW), build-linux (native)
  - MinGW cross-compiler support for building Windows executables from WSL2
- 10 tests for FormView component
- 8 tests for project I/O operations
- Total: 192 passing tests (all internal packages)

### Changed
- Content area now uses split view for simultaneous list and form display
- Window size increased to 1400x900 for better content visibility
- Sidebar width reduced to 180px for more content space
- AppState now includes ProjectPath field for tracking current project file
- SaveProject updates LastModified timestamp automatically

### Fixed
- Content area sizing issue - forms and lists now expand to fill available space

## [0.2.0] - 2025-10-13

### Added
- Phase 3: Basic UI Framework (Steps 14-20)
  - Application state management with singleton pattern
  - Main window with Fyne UI integration
  - Status bar with 4 sections (project, validation, records, saved status)
  - Sidebar navigation with section selection
  - Layout integration with BorderLayout
  - Menu bar with File, Edit, View, Help menus
  - Theme support with light/dark mode switching
- Phase 4: List View (Steps 21-25)
  - Player list view with table display (6 columns)
  - Coach list view with position names (5 columns)
  - Team list view with league structure (6 columns)
  - Column-based sorting (click headers to sort, toggle ascending/descending)
  - Visual sort indicators (▲/▼) on column headers
  - Search/filter functionality for player list
  - Real-time filtering as user types
- 135 tests for Phase 3 (UI framework)
- 29 tests for Phase 4 (list views)
- Total: 164 passing tests

### Changed
- Window now uses BorderLayout with sidebar, status bar, and content area
- Content area switches based on sidebar selection
- Tables support interactive sorting and filtering

## [0.1.2] - 2025-10-13

### Added
- PAT token support for workflow triggering (RELEASE_TOKEN)
- Comprehensive PAT_TOKEN_SETUP.md guide for maintainers

### Changed
- Auto-tag workflow now uses RELEASE_TOKEN instead of GITHUB_TOKEN
- README includes maintainer note about RELEASE_TOKEN requirement

### Fixed
- Release workflow now triggers properly when tags are created by auto-tag workflow
- Resolved GitHub Actions security limitation preventing workflow-to-workflow triggering

## [0.1.1] - 2025-10-13

### Added
- Automatic release system based on CHANGELOG.md updates
- Auto-tag workflow that creates git tags on version detection
- Comprehensive automatic-release-plan.md documentation
- Release badges in README (Latest Release, Build Status, License)

### Changed
- README now shows release badges and auto-release note
- PR template reminds developers about CHANGELOG updates
- RELEASING.md updated with automatic release workflow

### Fixed
- YAML syntax error in auto-tag workflow (multiline string)
- Removed accidental "auto_activate_base" line from README

### Infrastructure
- Auto-tag.yml workflow for automatic tag creation
- Release process now fully automated from CHANGELOG update

## [0.1.0] - 2025-10-13

### Added
- Initial project setup with Go 1.21 and Fyne v2.6.3 UI framework
- Core data models for Project, Player, Coach, Team, and LeagueInfo
- Complete data structures matching FOF9 game format (90+ player fields)
- CSV struct tags for all models
- Helper methods for data manipulation
- Basic application window with "Hello from FOF9 Editor" message
- Version management infrastructure with build-time injection
- `--version` CLI flag to display version information
- Version display in window title
- Pull request template with changelog checklist
- Changelog verification in CI pipeline for PRs
- Pre-release workflow for alpha/beta/rc releases
- Release helper scripts (prepare-release.sh, create-prerelease.sh, verify-release.sh)
- Comprehensive RELEASING.md documentation
- Automatic changelog extraction in release notes
- Build information in release artifacts (commit hash, build date)
- Comprehensive test suite (23 tests, all passing)
- GitHub Actions CI/CD pipeline for automated builds and releases
- Project specification and implementation plan documentation
- Release and versioning plan

### Infrastructure
- Makefile for build automation with version injection
- GitHub Actions workflows for build, test, and release
- Automated Windows executable generation for releases
- Enhanced release.yml with prerelease support and changelog extraction
- Added pre-release.yml workflow for alpha/beta/rc releases
- Added changelog-check.yml to enforce changelog updates in PRs
- Version information displayed in builds and window title
- Release archives include CHANGELOG.md

### Notes
- This is a development release
- Core features are in progress
- Only basic UI window is functional
- Not yet ready for production use - this release is for infrastructure testing

---

[Unreleased]: https://github.com/igorilic/fof9editor/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/igorilic/fof9editor/releases/tag/v0.1.0
