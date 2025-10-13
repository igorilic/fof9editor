package models

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestNewProject(t *testing.T) {
	name := "MyLeague"
	identifier := "myleague"
	basePath := "/test/path"
	baseYear := 2024

	project := NewProject(name, identifier, basePath, baseYear)

	if project.LeagueName != name {
		t.Errorf("Expected LeagueName %s, got %s", name, project.LeagueName)
	}

	if project.Identifier != identifier {
		t.Errorf("Expected Identifier %s, got %s", identifier, project.Identifier)
	}

	if project.BaseYear != baseYear {
		t.Errorf("Expected BaseYear %d, got %d", baseYear, project.BaseYear)
	}

	if project.Version != "1.0" {
		t.Errorf("Expected Version 1.0, got %s", project.Version)
	}

	if project.DataPath != "./data/" {
		t.Errorf("Expected DataPath ./data/, got %s", project.DataPath)
	}

	if project.ReferencePath != "./reference/" {
		t.Errorf("Expected ReferencePath ./reference/, got %s", project.ReferencePath)
	}

	// Check CSVFiles map is populated
	if len(project.CSVFiles) == 0 {
		t.Error("Expected CSVFiles to be populated")
	}

	expectedFiles := []string{"info", "players", "coaches", "teams", "teamColors"}
	for _, key := range expectedFiles {
		if _, ok := project.CSVFiles[key]; !ok {
			t.Errorf("Expected CSVFiles to contain key %s", key)
		}
	}
}

func TestGetFullPath(t *testing.T) {
	project := NewProject("Test", "test", "/base", 2024)

	// Test existing key
	playersPath := project.GetFullPath("players")
	expectedPath := filepath.Join("data", "test_players.csv")
	if playersPath != expectedPath {
		t.Errorf("Expected path %s, got %s", expectedPath, playersPath)
	}

	// Test non-existing key
	invalidPath := project.GetFullPath("nonexistent")
	if invalidPath != "" {
		t.Errorf("Expected empty string for invalid key, got %s", invalidPath)
	}
}

func TestProjectJSONMarshaling(t *testing.T) {
	project := NewProject("Test League", "testleague", "/test", 2024)

	// Marshal to JSON
	jsonData, err := json.Marshal(project)
	if err != nil {
		t.Fatalf("Failed to marshal project: %v", err)
	}

	// Unmarshal back
	var unmarshaled Project
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal project: %v", err)
	}

	// Verify fields
	if unmarshaled.LeagueName != project.LeagueName {
		t.Errorf("LeagueName mismatch after unmarshal")
	}

	if unmarshaled.BaseYear != project.BaseYear {
		t.Errorf("BaseYear mismatch after unmarshal")
	}

	if len(unmarshaled.CSVFiles) != len(project.CSVFiles) {
		t.Errorf("CSVFiles length mismatch after unmarshal")
	}
}
