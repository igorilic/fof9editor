// ABOUTME: Project file I/O operations for FOF9 Editor
// ABOUTME: Handles saving and loading .fof9proj files

package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/igorilic/fof9editor/internal/models"
)

// SaveProject saves a project to a .fof9proj file
func SaveProject(project *models.Project, filePath string) error {
	if project == nil {
		return fmt.Errorf("project is nil")
	}

	// Marshal project to JSON with indentation
	data, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal project: %w", err)
	}

	// Write to temp file first for atomic write
	dir := filepath.Dir(filePath)
	tempFile := filepath.Join(dir, "."+filepath.Base(filePath)+".tmp")

	// Ensure directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write to temp file
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tempFile, filePath); err != nil {
		// Clean up temp file on error
		os.Remove(tempFile)
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

// LoadProject loads a project from a .fof9proj file
func LoadProject(filePath string) (*models.Project, error) {
	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read project file: %w", err)
	}

	// Unmarshal JSON
	var project models.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, fmt.Errorf("failed to unmarshal project: %w", err)
	}

	// Validate required fields
	if project.Version == "" {
		return nil, fmt.Errorf("project version is required")
	}
	if project.LeagueName == "" {
		return nil, fmt.Errorf("league name is required")
	}
	if project.Identifier == "" {
		return nil, fmt.Errorf("project identifier is required")
	}

	return &project, nil
}
