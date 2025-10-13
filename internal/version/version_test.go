package version

import (
	"strings"
	"testing"
)

func TestGetVersionInfo(t *testing.T) {
	info := GetVersionInfo()

	// Should contain version
	if !strings.Contains(info, "FOF9 Editor v") {
		t.Errorf("Version info should contain 'FOF9 Editor v', got: %s", info)
	}

	// Should contain commit hash
	if !strings.Contains(info, "Commit:") {
		t.Errorf("Version info should contain 'Commit:', got: %s", info)
	}

	// Should contain build date
	if !strings.Contains(info, "Built:") {
		t.Errorf("Version info should contain 'Built:', got: %s", info)
	}

	// Should contain Go version
	if !strings.Contains(info, "Go:") {
		t.Errorf("Version info should contain 'Go:', got: %s", info)
	}
}

func TestGetShortVersion(t *testing.T) {
	// Save original values
	origVersion := Version
	origCommitHash := CommitHash
	defer func() {
		Version = origVersion
		CommitHash = origCommitHash
	}()

	// Test with default "dev" version
	Version = "dev"
	CommitHash = "abc123def456"

	short := GetShortVersion()
	if !strings.Contains(short, "dev") {
		t.Errorf("Short version should contain 'dev', got: %s", short)
	}
	if !strings.Contains(short, "abc123d") {
		t.Errorf("Short version should contain short commit hash 'abc123d', got: %s", short)
	}

	// Test with release version
	Version = "0.1.0"
	short = GetShortVersion()
	if short != "0.1.0" {
		t.Errorf("Short version should be '0.1.0', got: %s", short)
	}

	// Test with short commit hash
	Version = "dev"
	CommitHash = "abc"
	short = GetShortVersion()
	if short != "dev" {
		t.Errorf("Short version with short hash should be 'dev', got: %s", short)
	}
}

func TestVersionDefaults(t *testing.T) {
	// Save original values
	origVersion := Version
	origCommitHash := CommitHash
	origBuildDate := BuildDate
	defer func() {
		Version = origVersion
		CommitHash = origCommitHash
		BuildDate = origBuildDate
	}()

	// Reset to defaults
	Version = "dev"
	CommitHash = "unknown"
	BuildDate = "unknown"

	info := GetVersionInfo()
	if !strings.Contains(info, "dev") {
		t.Errorf("Default version should be 'dev', got: %s", info)
	}
	if !strings.Contains(info, "unknown") {
		t.Errorf("Default commit hash should be 'unknown', got: %s", info)
	}
}
