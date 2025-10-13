// ABOUTME: This file manages version information for the FOF9 Editor application
// ABOUTME: Version, commit hash, and build date are injected at build time via ldflags
package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// CommitHash is set at build time via ldflags
	CommitHash = "unknown"
	// BuildDate is set at build time via ldflags
	BuildDate = "unknown"
)

// GetVersionInfo returns formatted version information
func GetVersionInfo() string {
	return fmt.Sprintf("FOF9 Editor v%s\nCommit: %s\nBuilt: %s\nGo: %s %s/%s",
		Version, CommitHash, BuildDate, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// GetShortVersion returns just the version string
func GetShortVersion() string {
	if Version == "dev" {
		if len(CommitHash) >= 7 {
			return fmt.Sprintf("%s (%s)", Version, CommitHash[:7])
		}
		return Version
	}
	return Version
}
