// ABOUTME: Application state management for FOF9 Editor
// ABOUTME: Maintains current project, loaded data, and UI state using singleton pattern

package state

import (
	"fmt"
	"sync"
	"time"

	"github.com/igorilic/fof9editor/internal/data"
	"github.com/igorilic/fof9editor/internal/models"
)

// AppState represents the global application state
type AppState struct {
	// Current project
	Project     *models.Project
	ProjectPath string // Path to the .fof9proj file

	// Loaded data
	Players []models.Player
	Coaches []models.Coach
	Teams   []models.Team

	// UI state
	CurrentSection string // e.g., "Players", "Coaches", "Teams"
	SelectedIndex  int    // Currently selected item in list

	// Dirty flag
	IsDirty bool // True if there are unsaved changes

	// Mutex for thread safety
	mu sync.RWMutex
}

var (
	instance *AppState
	once     sync.Once
)

// GetInstance returns the singleton instance of AppState
func GetInstance() *AppState {
	once.Do(func() {
		instance = &AppState{
			CurrentSection: "Players",
			SelectedIndex:  -1,
			IsDirty:        false,
		}
	})
	return instance
}

// LoadProject loads a project from the given file path
func (s *AppState) LoadProject(filepath string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Load project file using data package
	project, err := data.LoadProject(filepath)
	if err != nil {
		return fmt.Errorf("failed to load project: %w", err)
	}

	// Store project
	s.Project = project
	s.ProjectPath = filepath

	// Load CSV data files
	if project.DataPath != "" {
		// Load players
		playersPath := project.GetFullPath("players")
		if players, err := data.LoadPlayers(playersPath); err == nil {
			s.Players = players
		}

		// Load coaches
		coachesPath := project.GetFullPath("coaches")
		if coaches, err := data.LoadCoaches(coachesPath); err == nil {
			s.Coaches = coaches
		}

		// Load teams
		teamsPath := project.GetFullPath("teams")
		if teams, err := data.LoadTeams(teamsPath); err == nil {
			s.Teams = teams
		}
	}

	// Mark as clean (no unsaved changes)
	s.IsDirty = false

	return nil
}

// SaveProject saves the current project to disk
func (s *AppState) SaveProject() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Project == nil {
		return fmt.Errorf("no project loaded")
	}

	if s.ProjectPath == "" {
		return fmt.Errorf("no project path set")
	}

	// Update LastModified timestamp
	s.Project.LastModified = time.Now()

	// Save project file
	if err := data.SaveProject(s.Project, s.ProjectPath); err != nil {
		return fmt.Errorf("failed to save project file: %w", err)
	}

	// Save CSV data files
	if s.Project.DataPath != "" {
		// Save players
		playersPath := s.Project.GetFullPath("players")
		if err := data.SavePlayers(playersPath, s.Players); err != nil {
			return fmt.Errorf("failed to save players: %w", err)
		}

		// Save coaches
		coachesPath := s.Project.GetFullPath("coaches")
		if err := data.SaveCoaches(coachesPath, s.Coaches); err != nil {
			return fmt.Errorf("failed to save coaches: %w", err)
		}

		// Save teams
		teamsPath := s.Project.GetFullPath("teams")
		if err := data.SaveTeams(teamsPath, s.Teams); err != nil {
			return fmt.Errorf("failed to save teams: %w", err)
		}
	}

	// Mark as clean (no unsaved changes)
	s.IsDirty = false

	return nil
}

// SetProject sets the current project
func (s *AppState) SetProject(project *models.Project) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Project = project
}

// GetProject returns the current project (thread-safe)
func (s *AppState) GetProject() *models.Project {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Project
}

// SetPlayers sets the players data
func (s *AppState) SetPlayers(players []models.Player) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Players = players
	s.IsDirty = true
}

// GetPlayers returns the players data (thread-safe)
func (s *AppState) GetPlayers() []models.Player {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Players
}

// SetCoaches sets the coaches data
func (s *AppState) SetCoaches(coaches []models.Coach) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Coaches = coaches
	s.IsDirty = true
}

// GetCoaches returns the coaches data (thread-safe)
func (s *AppState) GetCoaches() []models.Coach {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Coaches
}

// SetTeams sets the teams data
func (s *AppState) SetTeams(teams []models.Team) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Teams = teams
	s.IsDirty = true
}

// GetTeams returns the teams data (thread-safe)
func (s *AppState) GetTeams() []models.Team {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Teams
}

// SetCurrentSection sets the currently active section
func (s *AppState) SetCurrentSection(section string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CurrentSection = section
}

// GetCurrentSection returns the currently active section (thread-safe)
func (s *AppState) GetCurrentSection() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.CurrentSection
}

// SetSelectedIndex sets the currently selected item index
func (s *AppState) SetSelectedIndex(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SelectedIndex = index
}

// GetSelectedIndex returns the currently selected item index (thread-safe)
func (s *AppState) GetSelectedIndex() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.SelectedIndex
}

// MarkClean marks the state as having no unsaved changes
func (s *AppState) MarkClean() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IsDirty = false
}

// MarkDirty marks the state as having unsaved changes
func (s *AppState) MarkDirty() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IsDirty = true
}

// IsDirtyState returns whether there are unsaved changes (thread-safe)
func (s *AppState) IsDirtyState() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.IsDirty
}

// HasProject returns true if a project is currently loaded
func (s *AppState) HasProject() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Project != nil
}

// Reset resets the application state to initial values
func (s *AppState) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Project = nil
	s.Players = nil
	s.Coaches = nil
	s.Teams = nil
	s.CurrentSection = "Players"
	s.SelectedIndex = -1
	s.IsDirty = false
}
