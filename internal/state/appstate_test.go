// ABOUTME: Tests for application state management
// ABOUTME: Validates singleton pattern, thread safety, and state operations

package state

import (
	"testing"

	"github.com/igorilic/fof9editor/internal/models"
)

func TestGetInstance_Singleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Error("GetInstance should return the same instance (singleton)")
	}
}

func TestGetInstance_InitialState(t *testing.T) {
	// Reset state for clean test
	state := GetInstance()
	state.Reset()

	if state.CurrentSection != "Players" {
		t.Errorf("Expected CurrentSection 'Players', got '%s'", state.CurrentSection)
	}
	if state.SelectedIndex != -1 {
		t.Errorf("Expected SelectedIndex -1, got %d", state.SelectedIndex)
	}
	if state.IsDirty {
		t.Error("Expected IsDirty to be false initially")
	}
	if state.Project != nil {
		t.Error("Expected Project to be nil initially")
	}
}

func TestSetGetProject(t *testing.T) {
	state := GetInstance()
	state.Reset()

	project := models.NewProject("TestLeague", "testleague", "/path/to/project", 2024)
	state.SetProject(project)

	retrieved := state.GetProject()
	if retrieved == nil {
		t.Fatal("GetProject returned nil")
	}
	if retrieved.LeagueName != "TestLeague" {
		t.Errorf("Expected LeagueName 'TestLeague', got '%s'", retrieved.LeagueName)
	}
}

func TestHasProject(t *testing.T) {
	state := GetInstance()
	state.Reset()

	if state.HasProject() {
		t.Error("HasProject should return false when no project is loaded")
	}

	project := models.NewProject("Test", "test", "/path", 2024)
	state.SetProject(project)

	if !state.HasProject() {
		t.Error("HasProject should return true when project is loaded")
	}
}

func TestSetGetPlayers(t *testing.T) {
	state := GetInstance()
	state.Reset()

	players := []models.Player{
		{PlayerID: 1, LastName: "Doe", FirstName: "John"},
		{PlayerID: 2, LastName: "Smith", FirstName: "Jane"},
	}

	state.SetPlayers(players)

	retrieved := state.GetPlayers()
	if len(retrieved) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(retrieved))
	}
	if retrieved[0].LastName != "Doe" {
		t.Errorf("Expected LastName 'Doe', got '%s'", retrieved[0].LastName)
	}
}

func TestSetPlayers_MarksDirty(t *testing.T) {
	state := GetInstance()
	state.Reset()
	state.MarkClean()

	players := []models.Player{
		{PlayerID: 1, LastName: "Test"},
	}

	state.SetPlayers(players)

	if !state.IsDirtyState() {
		t.Error("SetPlayers should mark state as dirty")
	}
}

func TestSetGetCoaches(t *testing.T) {
	state := GetInstance()
	state.Reset()

	coaches := []models.Coach{
		{LastName: "Belichick", FirstName: "Bill"},
		{LastName: "Reid", FirstName: "Andy"},
	}

	state.SetCoaches(coaches)

	retrieved := state.GetCoaches()
	if len(retrieved) != 2 {
		t.Fatalf("Expected 2 coaches, got %d", len(retrieved))
	}
	if retrieved[0].LastName != "Belichick" {
		t.Errorf("Expected LastName 'Belichick', got '%s'", retrieved[0].LastName)
	}
}

func TestSetCoaches_MarksDirty(t *testing.T) {
	state := GetInstance()
	state.Reset()
	state.MarkClean()

	coaches := []models.Coach{
		{LastName: "Test"},
	}

	state.SetCoaches(coaches)

	if !state.IsDirtyState() {
		t.Error("SetCoaches should mark state as dirty")
	}
}

func TestSetGetTeams(t *testing.T) {
	state := GetInstance()
	state.Reset()

	teams := []models.Team{
		{TeamID: 1, NickName: "Patriots"},
		{TeamID: 2, NickName: "Chiefs"},
	}

	state.SetTeams(teams)

	retrieved := state.GetTeams()
	if len(retrieved) != 2 {
		t.Fatalf("Expected 2 teams, got %d", len(retrieved))
	}
	if retrieved[0].NickName != "Patriots" {
		t.Errorf("Expected NickName 'Patriots', got '%s'", retrieved[0].NickName)
	}
}

func TestSetTeams_MarksDirty(t *testing.T) {
	state := GetInstance()
	state.Reset()
	state.MarkClean()

	teams := []models.Team{
		{TeamID: 1},
	}

	state.SetTeams(teams)

	if !state.IsDirtyState() {
		t.Error("SetTeams should mark state as dirty")
	}
}

func TestSetGetCurrentSection(t *testing.T) {
	state := GetInstance()
	state.Reset()

	state.SetCurrentSection("Coaches")

	section := state.GetCurrentSection()
	if section != "Coaches" {
		t.Errorf("Expected CurrentSection 'Coaches', got '%s'", section)
	}
}

func TestSetGetSelectedIndex(t *testing.T) {
	state := GetInstance()
	state.Reset()

	state.SetSelectedIndex(5)

	index := state.GetSelectedIndex()
	if index != 5 {
		t.Errorf("Expected SelectedIndex 5, got %d", index)
	}
}

func TestMarkCleanDirty(t *testing.T) {
	state := GetInstance()
	state.Reset()

	// Initially clean
	state.MarkClean()
	if state.IsDirtyState() {
		t.Error("Expected IsDirty to be false after MarkClean")
	}

	// Mark dirty
	state.MarkDirty()
	if !state.IsDirtyState() {
		t.Error("Expected IsDirty to be true after MarkDirty")
	}

	// Mark clean again
	state.MarkClean()
	if state.IsDirtyState() {
		t.Error("Expected IsDirty to be false after second MarkClean")
	}
}

func TestReset(t *testing.T) {
	state := GetInstance()

	// Set up some state
	project := models.NewProject("Test", "test", "/path", 2024)
	state.SetProject(project)
	state.SetPlayers([]models.Player{{PlayerID: 1}})
	state.SetCoaches([]models.Coach{{LastName: "Test"}})
	state.SetTeams([]models.Team{{TeamID: 1}})
	state.SetCurrentSection("Teams")
	state.SetSelectedIndex(10)
	state.MarkDirty()

	// Reset
	state.Reset()

	// Verify everything is reset
	if state.GetProject() != nil {
		t.Error("Project should be nil after Reset")
	}
	if state.GetPlayers() != nil {
		t.Error("Players should be nil after Reset")
	}
	if state.GetCoaches() != nil {
		t.Error("Coaches should be nil after Reset")
	}
	if state.GetTeams() != nil {
		t.Error("Teams should be nil after Reset")
	}
	if state.GetCurrentSection() != "Players" {
		t.Errorf("CurrentSection should be 'Players' after Reset, got '%s'", state.GetCurrentSection())
	}
	if state.GetSelectedIndex() != -1 {
		t.Errorf("SelectedIndex should be -1 after Reset, got %d", state.GetSelectedIndex())
	}
	if state.IsDirtyState() {
		t.Error("IsDirty should be false after Reset")
	}
}

func TestLoadProject_Stub(t *testing.T) {
	state := GetInstance()
	state.Reset()

	err := state.LoadProject("/test/path")
	if err == nil {
		t.Error("LoadProject stub should return error")
	}
}

func TestSaveProject_Stub_NoProject(t *testing.T) {
	state := GetInstance()
	state.Reset()

	err := state.SaveProject()
	if err == nil {
		t.Error("SaveProject should return error when no project is loaded")
	}
}

func TestSaveProject_Stub_WithProject(t *testing.T) {
	state := GetInstance()
	state.Reset()

	project := models.NewProject("Test", "test", "/path", 2024)
	state.SetProject(project)

	err := state.SaveProject()
	if err == nil {
		t.Error("SaveProject stub should return error")
	}
}

// Thread safety test
func TestConcurrentAccess(t *testing.T) {
	state := GetInstance()
	state.Reset()

	done := make(chan bool)

	// Concurrent writers
	go func() {
		for i := 0; i < 100; i++ {
			state.SetCurrentSection("Test")
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			state.SetSelectedIndex(i)
		}
		done <- true
	}()

	// Concurrent readers
	go func() {
		for i := 0; i < 100; i++ {
			_ = state.GetCurrentSection()
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			_ = state.GetSelectedIndex()
		}
		done <- true
	}()

	// Wait for all goroutines
	for i := 0; i < 4; i++ {
		<-done
	}

	// If we get here without deadlock, the test passes
}
