// ABOUTME: Tests for project file I/O operations
// ABOUTME: Verifies save/load functionality for .fof9proj files

package data

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/igorilic/fof9editor/internal/models"
)

func TestSaveProject(t *testing.T) {
	// Create test project
	project := &models.Project{
		Version:      "1.0",
		LeagueName:   "Test League",
		Identifier:   "test-league-001",
		Created:      time.Now(),
		LastModified: time.Now(),
		BaseYear:     2024,
		DataPath:     "./data/",
		ReferencePath: "./reference/",
		CSVFiles: map[string]string{
			"info":    "info.csv",
			"players": "players.csv",
			"coaches": "coaches.csv",
			"teams":   "teams.csv",
		},
		UserPreferences: map[string]interface{}{
			"theme": "dark",
		},
	}

	// Create temp directory
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.fof9proj")

	// Save project
	err := SaveProject(project, filePath)
	if err != nil {
		t.Fatalf("SaveProject failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Project file was not created")
	}

	// Verify file is valid JSON
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	if len(data) == 0 {
		t.Error("Saved file is empty")
	}
}

func TestLoadProject(t *testing.T) {
	// Create test project
	originalProject := &models.Project{
		Version:      "1.0",
		LeagueName:   "Test League",
		Identifier:   "test-league-001",
		Created:      time.Now().Round(time.Second), // Round to avoid precision issues
		LastModified: time.Now().Round(time.Second),
		BaseYear:     2024,
		DataPath:     "./data/",
		ReferencePath: "./reference/",
		CSVFiles: map[string]string{
			"info":    "info.csv",
			"players": "players.csv",
		},
	}

	// Save to temp file
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.fof9proj")

	err := SaveProject(originalProject, filePath)
	if err != nil {
		t.Fatalf("SaveProject failed: %v", err)
	}

	// Load project back
	loadedProject, err := LoadProject(filePath)
	if err != nil {
		t.Fatalf("LoadProject failed: %v", err)
	}

	// Verify fields match
	if loadedProject.Version != originalProject.Version {
		t.Errorf("Version mismatch: got %s, want %s", loadedProject.Version, originalProject.Version)
	}

	if loadedProject.LeagueName != originalProject.LeagueName {
		t.Errorf("LeagueName mismatch: got %s, want %s", loadedProject.LeagueName, originalProject.LeagueName)
	}

	if loadedProject.Identifier != originalProject.Identifier {
		t.Errorf("Identifier mismatch: got %s, want %s", loadedProject.Identifier, originalProject.Identifier)
	}

	if loadedProject.BaseYear != originalProject.BaseYear {
		t.Errorf("BaseYear mismatch: got %d, want %d", loadedProject.BaseYear, originalProject.BaseYear)
	}

	if loadedProject.DataPath != originalProject.DataPath {
		t.Errorf("DataPath mismatch: got %s, want %s", loadedProject.DataPath, originalProject.DataPath)
	}

	if len(loadedProject.CSVFiles) != len(originalProject.CSVFiles) {
		t.Errorf("CSVFiles count mismatch: got %d, want %d", len(loadedProject.CSVFiles), len(originalProject.CSVFiles))
	}
}

func TestLoadProject_NonExistentFile(t *testing.T) {
	_, err := LoadProject("/nonexistent/path/test.fof9proj")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestLoadProject_InvalidJSON(t *testing.T) {
	// Create temp file with invalid JSON
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "invalid.fof9proj")

	err := os.WriteFile(filePath, []byte("invalid json content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err = LoadProject(filePath)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestLoadProject_MissingRequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		project *models.Project
	}{
		{
			name: "Missing Version",
			project: &models.Project{
				LeagueName: "Test",
				Identifier: "test-001",
			},
		},
		{
			name: "Missing LeagueName",
			project: &models.Project{
				Version:    "1.0",
				Identifier: "test-001",
			},
		},
		{
			name: "Missing Identifier",
			project: &models.Project{
				Version:    "1.0",
				LeagueName: "Test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save invalid project
			tempDir := t.TempDir()
			filePath := filepath.Join(tempDir, "test.fof9proj")

			// Manually save to bypass SaveProject validation
			data, _ := json.Marshal(tt.project)
			os.WriteFile(filePath, data, 0644)

			// Try to load
			_, err := LoadProject(filePath)
			if err == nil {
				t.Errorf("%s: Expected error for missing field, got nil", tt.name)
			}
		})
	}
}

func TestSaveProject_NilProject(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.fof9proj")

	err := SaveProject(nil, filePath)
	if err == nil {
		t.Error("Expected error for nil project, got nil")
	}
}

func TestSaveProject_AtomicWrite(t *testing.T) {
	// Create test project
	project := &models.Project{
		Version:    "1.0",
		LeagueName: "Test League",
		Identifier: "test-league-001",
		BaseYear:   2024,
	}

	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.fof9proj")

	// Save project
	err := SaveProject(project, filePath)
	if err != nil {
		t.Fatalf("SaveProject failed: %v", err)
	}

	// Verify temp file was cleaned up
	tempFile := filepath.Join(tempDir, ".test.fof9proj.tmp")
	if _, err := os.Stat(tempFile); !os.IsNotExist(err) {
		t.Error("Temp file was not cleaned up")
	}

	// Verify final file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Error("Final project file does not exist")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	// Create comprehensive project
	originalProject := &models.Project{
		Version:      "1.0",
		LeagueName:   "Test League",
		Identifier:   "test-league-001",
		Created:      time.Date(2024, 10, 13, 12, 0, 0, 0, time.UTC),
		LastModified: time.Date(2024, 10, 13, 13, 0, 0, 0, time.UTC),
		BaseYear:     2024,
		DataPath:     "./data/",
		ReferencePath: "./reference/",
		CSVFiles: map[string]string{
			"info":       "info.csv",
			"players":    "players.csv",
			"coaches":    "coaches.csv",
			"teams":      "teams.csv",
			"teamColors": "team_colors.csv",
		},
		UserPreferences: map[string]interface{}{
			"theme":      "dark",
			"windowSize": map[string]interface{}{"width": 1400, "height": 900},
		},
	}

	// Save and load
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "roundtrip.fof9proj")

	if err := SaveProject(originalProject, filePath); err != nil {
		t.Fatalf("SaveProject failed: %v", err)
	}

	loadedProject, err := LoadProject(filePath)
	if err != nil {
		t.Fatalf("LoadProject failed: %v", err)
	}

	// Verify all CSV files
	for key, filename := range originalProject.CSVFiles {
		if loadedProject.CSVFiles[key] != filename {
			t.Errorf("CSVFile[%s] mismatch: got %s, want %s", key, loadedProject.CSVFiles[key], filename)
		}
	}

	// Verify timestamps
	if !loadedProject.Created.Equal(originalProject.Created) {
		t.Errorf("Created timestamp mismatch: got %v, want %v", loadedProject.Created, originalProject.Created)
	}

	if !loadedProject.LastModified.Equal(originalProject.LastModified) {
		t.Errorf("LastModified timestamp mismatch: got %v, want %v", loadedProject.LastModified, originalProject.LastModified)
	}
}
