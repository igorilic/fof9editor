// ABOUTME: This file defines the project structure for FOF9 custom leagues
// ABOUTME: It manages the .fof9proj file format with league metadata and file references
package models

import (
	"path/filepath"
	"time"
)

// Project represents a FOF9 custom league project
type Project struct {
	Version         string                 `json:"version"`
	LeagueName      string                 `json:"leagueName"`
	Identifier      string                 `json:"identifier"`
	Created         time.Time              `json:"created"`
	LastModified    time.Time              `json:"lastModified"`
	BaseYear        int                    `json:"baseYear"`
	DataPath        string                 `json:"dataPath"`
	ReferencePath   string                 `json:"referencePath"`
	CSVFiles        map[string]string      `json:"csvFiles"`
	UserPreferences map[string]interface{} `json:"userPreferences"`
}

// NewProject creates a new project with default settings
func NewProject(name, identifier, basePath string, baseYear int) *Project {
	now := time.Now()
	return &Project{
		Version:      "1.0",
		LeagueName:   name,
		Identifier:   identifier,
		Created:      now,
		LastModified: now,
		BaseYear:     baseYear,
		DataPath:     "./data/",
		ReferencePath: "./reference/",
		CSVFiles: map[string]string{
			"info":       filepath.Join("data", identifier+"_info.csv"),
			"players":    filepath.Join("data", identifier+"_players.csv"),
			"coaches":    filepath.Join("data", identifier+"_coaches.csv"),
			"teams":      filepath.Join("data", "team_info.csv"),
			"teamColors": filepath.Join("data", "team_colors.csv"),
		},
		UserPreferences: make(map[string]interface{}),
	}
}

// GetFullPath returns the absolute path to a CSV file
func (p *Project) GetFullPath(csvKey string) string {
	if relPath, ok := p.CSVFiles[csvKey]; ok {
		return relPath
	}
	return ""
}
