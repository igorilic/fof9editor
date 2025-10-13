# Release and Versioning Plan

## Overview

This sub-plan details the strategy for versioning, releasing, and maintaining the FOF9 Editor with automated changelog generation and artifact publishing.

---

## Versioning Strategy

### Semantic Versioning (SemVer)
We follow [Semantic Versioning 2.0.0](https://semver.org/):
- **MAJOR.MINOR.PATCH** (e.g., 1.2.3)
- **MAJOR**: Breaking changes to file formats or APIs
- **MINOR**: New features, backwards compatible
- **PATCH**: Bug fixes, backwards compatible

### Pre-release Tags
- **Alpha**: `v0.1.0-alpha.1` - Early development, unstable
- **Beta**: `v0.1.0-beta.1` - Feature complete, testing phase
- **Release Candidate**: `v0.1.0-rc.1` - Final testing before release
- **Stable**: `v0.1.0` - Production ready

### Version Management
- Version stored in `cmd/fof9editor/version.go`
- Build-time injection via ldflags: `-X main.Version=x.y.z`
- Displayed in: About dialog, window title, CLI `--version` flag

---

## Release Workflow Steps

### Step R1: Version File and Build Info

**Goal**: Create version management infrastructure

**Implementation**:

1. Create `cmd/fof9editor/version.go`:
```go
package main

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
		return fmt.Sprintf("%s (%s)", Version, CommitHash[:7])
	}
	return Version
}
```

2. Update `cmd/fof9editor/main.go` to add version flag:
```go
var showVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(GetVersionInfo())
		os.Exit(0)
	}

	// ... rest of main
}
```

3. Update Makefile to inject version info:
```makefile
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT_HASH := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags="-s -w \
	-X main.Version=$(VERSION) \
	-X main.CommitHash=$(COMMIT_HASH) \
	-X main.BuildDate=$(BUILD_DATE)"

build:
	go build $(LDFLAGS) -o bin/fof9editor.exe ./cmd/fof9editor
```

**Testing**:
- Run `./bin/fof9editor.exe --version` and verify output
- Test with dirty git state
- Test with tagged version

**Deliverables**:
- `cmd/fof9editor/version.go`
- Updated `cmd/fof9editor/main.go`
- Updated `Makefile`

---

### Step R2: Changelog Management

**Goal**: Implement automated changelog generation

**Implementation**:

1. Install changelog tool as dev dependency:
```bash
go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
```

2. Create `.chglog/config.yml`:
```yaml
style: github
template: CHANGELOG.tpl.md
info:
  title: CHANGELOG
  repository_url: https://github.com/igorilic/fof9editor

options:
  commits:
    filters:
      Type:
        - added
        - changed
        - deprecated
        - removed
        - fixed
        - security
  commit_groups:
    title_maps:
      added: Added
      changed: Changed
      deprecated: Deprecated
      removed: Removed
      fixed: Fixed
      security: Security
  header:
    pattern: "^(\\w+):\\s(.+)$"
    pattern_maps:
      - Type
      - Subject
  notes:
    keywords:
      - BREAKING CHANGE
```

3. Create `.chglog/CHANGELOG.tpl.md`:
```markdown
{{ range .Versions }}
## {{ if .Tag.Previous }}[{{ .Tag.Name }}]{{ else }}Unreleased{{ end }} - {{ datetime "2006-01-02" .Tag.Date }}
{{ range .CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
- {{ .Subject }}{{ if .Body }} ({{ .Body }}){{ end }}
{{ end }}
{{ end -}}
{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
### {{ .Title }}
{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}
```

4. Update CHANGELOG.md format:
- Move to Keep a Changelog format strictly
- Add version comparison links at bottom
- Automate with git-chglog

5. Create script `scripts/generate-changelog.sh`:
```bash
#!/bin/bash
set -e

# Generate full changelog
git-chglog -o CHANGELOG.md

# Generate changelog for latest tag only (for release notes)
if [ -n "$1" ]; then
    git-chglog --tag-filter-pattern "$1" -o RELEASE_NOTES.md "$1"
fi

echo "Changelog generated successfully"
```

**Alternative: Manual Changelog Enforcement**

If preferring manual changelog updates:

1. Create `.github/pull_request_template.md`:
```markdown
## Description
<!-- Describe your changes -->

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)

## Changelog Entry
<!-- Add your changelog entry here following Keep a Changelog format -->

### Added
-

### Changed
-

### Fixed
-

## Testing
<!-- Describe the tests you ran -->
```

2. Add CI check to verify CHANGELOG.md is updated:
```yaml
# .github/workflows/changelog-check.yml
name: Changelog Check

on:
  pull_request:
    branches: [main]

jobs:
  check-changelog:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Check if CHANGELOG.md was updated
        run: |
          git diff --name-only origin/main...HEAD | grep -q "CHANGELOG.md" || {
            echo "Error: CHANGELOG.md was not updated"
            echo "Please add an entry to CHANGELOG.md describing your changes"
            exit 1
          }
```

**Testing**:
- Generate changelog: `./scripts/generate-changelog.sh`
- Verify format matches Keep a Changelog
- Test release notes generation for specific tag

**Deliverables**:
- `.chglog/config.yml` (if using git-chglog)
- `.chglog/CHANGELOG.tpl.md` (if using git-chglog)
- `scripts/generate-changelog.sh`
- Updated `CHANGELOG.md`
- `.github/pull_request_template.md` (if manual)
- `.github/workflows/changelog-check.yml` (if manual)

---

### Step R3: Release Workflow Enhancement

**Goal**: Enhance GitHub Actions release workflow with changelog extraction

**Implementation**:

1. Update `.github/workflows/release.yml`:

```yaml
name: Release

on:
  push:
    tags:
      - 'v*.*.*'
  workflow_dispatch:
    inputs:
      version:
        description: 'Version number (e.g., 0.1.0)'
        required: true
        type: string
      prerelease:
        description: 'Mark as pre-release?'
        required: false
        type: boolean
        default: false

jobs:
  build-release:
    name: Build Release Binaries
    runs-on: windows-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0  # Need full history for changelog

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
          cache: true

      - name: Get version
        id: version
        shell: bash
        run: |
          if [ "${{ github.event_name }}" = "workflow_dispatch" ]; then
            echo "VERSION=${{ github.event.inputs.version }}" >> $GITHUB_OUTPUT
            echo "TAG=v${{ github.event.inputs.version }}" >> $GITHUB_OUTPUT
          else
            echo "VERSION=${GITHUB_REF#refs/tags/v}" >> $GITHUB_OUTPUT
            echo "TAG=${GITHUB_REF#refs/tags/}" >> $GITHUB_OUTPUT
          fi

          # Get commit hash and build date
          echo "COMMIT_HASH=$(git rev-parse HEAD)" >> $GITHUB_OUTPUT
          echo "BUILD_DATE=$(date -u '+%Y-%m-%d_%H:%M:%S')" >> $GITHUB_OUTPUT

      - name: Download dependencies
        run: go mod download

      - name: Run tests
        run: go test ./internal/... -v

      - name: Build Windows executable
        run: |
          go build -ldflags="-s -w -X main.Version=${{ steps.version.outputs.VERSION }} -X main.CommitHash=${{ steps.version.outputs.COMMIT_HASH }} -X main.BuildDate=${{ steps.version.outputs.BUILD_DATE }}" -o fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.exe ./cmd/fof9editor

      - name: Create release archive
        shell: bash
        run: |
          mkdir -p release
          cp fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.exe release/
          cp README.md release/
          cp CHANGELOG.md release/
          cp LICENSE release/ || echo "No LICENSE file"
          cd release
          7z a -tzip ../fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.zip *

      - name: Generate checksums
        shell: bash
        run: |
          sha256sum fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.exe > checksums.txt
          sha256sum fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.zip >> checksums.txt

      - name: Extract changelog for this version
        id: changelog
        shell: bash
        run: |
          VERSION="${{ steps.version.outputs.TAG }}"

          # Extract changelog section for this version
          CHANGELOG_CONTENT=$(awk "/## \[${VERSION#v}\]/{flag=1; next} /## \[/{flag=0} flag" CHANGELOG.md)

          if [ -z "$CHANGELOG_CONTENT" ]; then
            # Fallback to Unreleased section if version not found
            CHANGELOG_CONTENT=$(awk "/## \[Unreleased\]/{flag=1; next} /## \[/{flag=0} flag" CHANGELOG.md)
          fi

          if [ -z "$CHANGELOG_CONTENT" ]; then
            CHANGELOG_CONTENT="See [CHANGELOG.md](https://github.com/${{ github.repository }}/blob/main/CHANGELOG.md) for details."
          fi

          # Write to file to preserve formatting
          echo "$CHANGELOG_CONTENT" > release_notes.txt

      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          tag_name: ${{ steps.version.outputs.TAG }}
          name: Release ${{ steps.version.outputs.TAG }}
          draft: false
          prerelease: ${{ github.event.inputs.prerelease || false }}
          body_path: release_notes.txt
          files: |
            fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.exe
            fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.zip
            checksums.txt
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Upload build artifacts
        uses: actions/upload-artifact@v4
        with:
          name: release-artifacts-${{ steps.version.outputs.VERSION }}
          path: |
            fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.exe
            fof9editor-${{ steps.version.outputs.VERSION }}-windows-amd64.zip
            checksums.txt
          retention-days: 90
```

**Testing**:
- Test manual dispatch workflow with version input
- Test tag-based automatic release
- Verify changelog extraction works
- Test pre-release flag

**Deliverables**:
- Updated `.github/workflows/release.yml`

---

### Step R4: Pre-release Workflow

**Goal**: Create workflow for alpha/beta/rc releases

**Implementation**:

1. Create `.github/workflows/pre-release.yml`:

```yaml
name: Pre-Release

on:
  workflow_dispatch:
    inputs:
      version:
        description: 'Version number (e.g., 0.2.0-beta.1)'
        required: true
        type: string
      release_type:
        description: 'Pre-release type'
        required: true
        type: choice
        options:
          - alpha
          - beta
          - rc

jobs:
  pre-release:
    name: Create Pre-Release
    runs-on: windows-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
          cache: true

      - name: Validate version format
        shell: bash
        run: |
          VERSION="${{ github.event.inputs.version }}"
          if [[ ! "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+-(${{ github.event.inputs.release_type }})\.[0-9]+$ ]]; then
            echo "Error: Version must match format X.Y.Z-${{ github.event.inputs.release_type }}.N"
            exit 1
          fi

      - name: Download dependencies
        run: go mod download

      - name: Run tests
        run: go test ./internal/... -v

      - name: Build Windows executable
        shell: bash
        run: |
          VERSION="${{ github.event.inputs.version }}"
          COMMIT_HASH=$(git rev-parse HEAD)
          BUILD_DATE=$(date -u '+%Y-%m-%d_%H:%M:%S')

          go build -ldflags="-s -w -X main.Version=${VERSION} -X main.CommitHash=${COMMIT_HASH} -X main.BuildDate=${BUILD_DATE}" -o fof9editor-${VERSION}-windows-amd64.exe ./cmd/fof9editor

      - name: Create release archive
        shell: bash
        run: |
          VERSION="${{ github.event.inputs.version }}"
          mkdir -p release
          cp fof9editor-${VERSION}-windows-amd64.exe release/
          cp README.md release/
          cp CHANGELOG.md release/
          cd release
          7z a -tzip ../fof9editor-${VERSION}-windows-amd64.zip *

      - name: Generate checksums
        shell: bash
        run: |
          VERSION="${{ github.event.inputs.version }}"
          sha256sum fof9editor-${VERSION}-windows-amd64.exe > checksums.txt
          sha256sum fof9editor-${VERSION}-windows-amd64.zip >> checksums.txt

      - name: Create pre-release notes
        shell: bash
        run: |
          cat > pre_release_notes.md <<'EOF'
          ## ⚠️ Pre-Release - ${{ github.event.inputs.release_type | upper }}

          This is a **${{ github.event.inputs.release_type }}** pre-release version and may contain bugs or incomplete features.

          **Not recommended for production use!**

          ### What's New

          EOF

          # Extract unreleased changes from CHANGELOG.md
          awk "/## \[Unreleased\]/{flag=1; next} /## \[/{flag=0} flag" CHANGELOG.md >> pre_release_notes.md || echo "See CHANGELOG.md for details." >> pre_release_notes.md

          cat >> pre_release_notes.md <<'EOF'

          ### Installation
          1. Download the `.exe` file or extract the `.zip` file
          2. Run `fof9editor-${{ github.event.inputs.version }}-windows-amd64.exe`
          3. Report any issues at https://github.com/${{ github.repository }}/issues

          ### Testing Needed
          - [ ] Player data import/export
          - [ ] Coach data editing
          - [ ] Team configuration
          - [ ] League wizard
          - [ ] Data validation

          ### Known Issues
          <!-- Add known issues here -->

          EOF

      - name: Create GitHub Release
        uses: softprops/action-gh-release@v1
        with:
          tag_name: v${{ github.event.inputs.version }}
          name: Pre-Release v${{ github.event.inputs.version }}
          draft: false
          prerelease: true
          body_path: pre_release_notes.md
          files: |
            fof9editor-${{ github.event.inputs.version }}-windows-amd64.exe
            fof9editor-${{ github.event.inputs.version }}-windows-amd64.zip
            checksums.txt
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**Testing**:
- Trigger pre-release workflow with alpha version
- Verify pre-release flag is set
- Verify warning message appears
- Test version format validation

**Deliverables**:
- `.github/workflows/pre-release.yml`

---

### Step R5: Release Scripts and Documentation

**Goal**: Create helper scripts and document release process

**Implementation**:

1. Create `scripts/prepare-release.sh`:

```bash
#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}FOF9 Editor - Release Preparation Script${NC}"
echo "========================================"

# Check if version is provided
if [ -z "$1" ]; then
    echo -e "${RED}Error: Version number required${NC}"
    echo "Usage: ./scripts/prepare-release.sh X.Y.Z"
    exit 1
fi

VERSION=$1
TAG="v${VERSION}"

echo -e "\n${YELLOW}Preparing release for version ${VERSION}${NC}\n"

# 1. Check if on main branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo -e "${RED}Error: Must be on main branch to release${NC}"
    echo "Current branch: $CURRENT_BRANCH"
    exit 1
fi

# 2. Check for uncommitted changes
if [ -n "$(git status --porcelain)" ]; then
    echo -e "${RED}Error: Uncommitted changes detected${NC}"
    git status --short
    exit 1
fi

# 3. Pull latest changes
echo "Pulling latest changes..."
git pull origin main

# 4. Run tests
echo -e "\n${YELLOW}Running tests...${NC}"
go test ./internal/... -v
if [ $? -ne 0 ]; then
    echo -e "${RED}Tests failed! Fix issues before releasing.${NC}"
    exit 1
fi
echo -e "${GREEN}All tests passed!${NC}"

# 5. Check if CHANGELOG.md has been updated
if ! grep -q "## \[${VERSION}\]" CHANGELOG.md; then
    echo -e "${RED}Error: CHANGELOG.md doesn't contain entry for version ${VERSION}${NC}"
    echo "Please update CHANGELOG.md with release notes"
    exit 1
fi

# 6. Update CHANGELOG.md Unreleased section
echo -e "\n${YELLOW}Updating CHANGELOG.md...${NC}"
TODAY=$(date +%Y-%m-%d)
sed -i.bak "s/## \[Unreleased\]/## [Unreleased]\n\n## [${VERSION}] - ${TODAY}/" CHANGELOG.md
rm CHANGELOG.md.bak 2>/dev/null || true

# 7. Commit changelog changes if modified
if [ -n "$(git status --porcelain CHANGELOG.md)" ]; then
    git add CHANGELOG.md
    git commit -m "chore: prepare release ${VERSION}

Updated CHANGELOG.md for version ${VERSION}

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>"
fi

# 8. Create and push tag
echo -e "\n${YELLOW}Creating tag ${TAG}...${NC}"
git tag -a "${TAG}" -m "Release ${VERSION}"

echo -e "\n${GREEN}Release preparation complete!${NC}"
echo -e "\nNext steps:"
echo "1. Review the changes"
echo "2. Push the tag: ${YELLOW}git push origin ${TAG}${NC}"
echo "3. Monitor the release workflow: ${YELLOW}gh run watch${NC}"
echo "4. Verify release page: ${YELLOW}gh release view ${TAG}${NC}"
```

2. Create `scripts/create-prerelease.sh`:

```bash
#!/bin/bash
set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}FOF9 Editor - Pre-Release Creation Script${NC}"
echo "=========================================="

if [ -z "$1" ] || [ -z "$2" ]; then
    echo -e "${RED}Error: Version and type required${NC}"
    echo "Usage: ./scripts/create-prerelease.sh X.Y.Z-TYPE.N TYPE"
    echo "Example: ./scripts/create-prerelease.sh 0.2.0-beta.1 beta"
    exit 1
fi

VERSION=$1
TYPE=$2

# Validate type
if [[ ! "$TYPE" =~ ^(alpha|beta|rc)$ ]]; then
    echo -e "${RED}Error: Type must be alpha, beta, or rc${NC}"
    exit 1
fi

# Validate version format
if [[ ! "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+-${TYPE}\.[0-9]+$ ]]; then
    echo -e "${RED}Error: Version must match format X.Y.Z-${TYPE}.N${NC}"
    exit 1
fi

echo -e "\n${YELLOW}Triggering pre-release workflow for ${VERSION}${NC}\n"

gh workflow run pre-release.yml \
    -f version="${VERSION}" \
    -f release_type="${TYPE}"

echo -e "\n${GREEN}Pre-release workflow triggered!${NC}"
echo "Monitor progress: ${YELLOW}gh run watch${NC}"
```

3. Create `scripts/verify-release.sh`:

```bash
#!/bin/bash
set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}FOF9 Editor - Release Verification Script${NC}"
echo "=========================================="

if [ -z "$1" ]; then
    echo -e "${RED}Error: Version tag required${NC}"
    echo "Usage: ./scripts/verify-release.sh v0.1.0"
    exit 1
fi

TAG=$1
VERSION=${TAG#v}

echo -e "\n${YELLOW}Verifying release ${TAG}${NC}\n"

# 1. Check if tag exists
if ! git rev-parse "$TAG" >/dev/null 2>&1; then
    echo -e "${RED}Error: Tag ${TAG} does not exist${NC}"
    exit 1
fi

# 2. Check if release exists on GitHub
if ! gh release view "$TAG" >/dev/null 2>&1; then
    echo -e "${RED}Error: GitHub release ${TAG} not found${NC}"
    exit 1
fi

# 3. Download release assets
echo "Downloading release assets..."
mkdir -p /tmp/verify-release
cd /tmp/verify-release

gh release download "$TAG" -p "*.exe" -p "checksums.txt"

# 4. Verify checksums
echo -e "\n${YELLOW}Verifying checksums...${NC}"
if sha256sum -c checksums.txt; then
    echo -e "${GREEN}Checksums verified!${NC}"
else
    echo -e "${RED}Checksum verification failed!${NC}"
    exit 1
fi

# 5. Test executable version
echo -e "\n${YELLOW}Checking executable version...${NC}"
WINE_VERSION=$(wine "fof9editor-${VERSION}-windows-amd64.exe" --version 2>/dev/null | head -n1) || true

if [ -z "$WINE_VERSION" ]; then
    echo -e "${YELLOW}Warning: Cannot test Windows executable on this system${NC}"
    echo "Manual verification needed on Windows"
else
    if echo "$WINE_VERSION" | grep -q "${VERSION}"; then
        echo -e "${GREEN}Version matches: ${WINE_VERSION}${NC}"
    else
        echo -e "${RED}Version mismatch!${NC}"
        echo "Expected: ${VERSION}"
        echo "Got: ${WINE_VERSION}"
        exit 1
    fi
fi

echo -e "\n${GREEN}Release verification complete!${NC}"

# Cleanup
cd -
rm -rf /tmp/verify-release
```

4. Create `RELEASING.md` documentation:

```markdown
# Release Process

This document describes the release process for FOF9 Editor.

## Release Types

### Stable Releases (vX.Y.Z)
- Production-ready releases
- Full changelog required
- All tests must pass
- Released from `main` branch

### Pre-Releases
- **Alpha** (vX.Y.Z-alpha.N): Early development, unstable
- **Beta** (vX.Y.Z-beta.N): Feature complete, testing phase
- **RC** (vX.Y.Z-rc.N): Release candidate, final testing

## Creating a Stable Release

### 1. Prepare the Release

```bash
# Ensure you're on main branch with latest changes
git checkout main
git pull origin main

# Run the preparation script
./scripts/prepare-release.sh 0.2.0
```

This script will:
- Verify you're on the main branch
- Check for uncommitted changes
- Run all tests
- Update CHANGELOG.md
- Create a commit with changelog updates

### 2. Push the Tag

```bash
# Push the tag to trigger release workflow
git push origin v0.2.0

# Monitor the release workflow
gh run watch
```

### 3. Verify the Release

```bash
# Wait for workflow to complete, then verify
./scripts/verify-release.sh v0.2.0

# Check the release page
gh release view v0.2.0
```

## Creating a Pre-Release

```bash
# Create a beta pre-release
./scripts/create-prerelease.sh 0.2.0-beta.1 beta

# Monitor the workflow
gh run watch
```

## Updating CHANGELOG.md

We follow [Keep a Changelog](https://keepachangelog.com/) format.

### Format

```markdown
## [Unreleased]

### Added
- New features

### Changed
- Changes to existing functionality

### Deprecated
- Features that will be removed

### Removed
- Removed features

### Fixed
- Bug fixes

### Security
- Security fixes

## [0.2.0] - 2025-10-15

### Added
- Feature X
- Feature Y
```

### When to Update

- **Before PR merge**: Add entry to Unreleased section
- **During release**: Preparation script moves Unreleased to version section

## Versioning Guidelines

### When to Bump MAJOR (X.0.0)
- Breaking changes to .fof9proj file format
- Removal of CSV columns
- Changed CSV column meanings
- Major UI redesign requiring user relearning

### When to Bump MINOR (0.X.0)
- New features (new entity types, import/export formats)
- New CSV columns (backward compatible)
- New validation rules
- New UI sections

### When to Bump PATCH (0.0.X)
- Bug fixes
- UI polish
- Performance improvements
- Documentation updates
- Dependency updates

## Hotfix Process

For urgent fixes needed in production:

```bash
# 1. Create hotfix branch from tag
git checkout -b hotfix/0.1.1 v0.1.0

# 2. Make the fix and commit
git commit -m "fixed: critical bug in player import"

# 3. Update CHANGELOG.md
# Add ## [0.1.1] - YYYY-MM-DD with fix description

# 4. Commit changelog
git commit -am "chore: prepare hotfix release 0.1.1"

# 5. Tag and push
git tag -a v0.1.1 -m "Hotfix 0.1.1"
git push origin v0.1.1

# 6. Merge back to main
git checkout main
git merge hotfix/0.1.1
git push origin main
```

## Rollback Process

If a release has critical issues:

```bash
# 1. Mark release as draft (hides from users)
gh release edit v0.2.0 --draft

# 2. Add warning to release notes
gh release edit v0.2.0 --notes "⚠️ WARNING: This release has been pulled due to critical issue #123. Use v0.1.0 instead."

# 3. Create hotfix or new release

# 4. Delete the bad tag
git tag -d v0.2.0
git push origin :refs/tags/v0.2.0
```

## Release Checklist

Before creating a release:

- [ ] All tests passing
- [ ] CHANGELOG.md updated
- [ ] README.md reflects new version
- [ ] Documentation updated
- [ ] No TODO/FIXME in critical code
- [ ] Version constants updated
- [ ] Manual testing completed
- [ ] Known issues documented

After release:

- [ ] Release assets uploaded correctly
- [ ] Checksums verified
- [ ] Download links work
- [ ] Version info displays correctly
- [ ] GitHub release notes look correct
- [ ] Tweet/announce release (optional)

## Troubleshooting

### Release workflow failed
```bash
# Check workflow logs
gh run list --workflow=release.yml --limit 5
gh run view <run-id> --log-failed

# Fix issue and re-trigger
git tag -d v0.2.0  # Delete local tag
git push origin :refs/tags/v0.2.0  # Delete remote tag
# Fix the issue
git tag -a v0.2.0 -m "Release 0.2.0"
git push origin v0.2.0
```

### Wrong files in release
```bash
# Edit release and re-upload
gh release upload v0.2.0 <correct-file> --clobber
```

### Changelog extraction failed
```bash
# Manually edit release notes
gh release edit v0.2.0
```
```

5. Make scripts executable:
```bash
chmod +x scripts/prepare-release.sh
chmod +x scripts/create-prerelease.sh
chmod +x scripts/verify-release.sh
chmod +x scripts/generate-changelog.sh
```

**Testing**:
- Run prepare-release.sh in dry-run mode
- Test each script with error conditions
- Verify documentation is clear

**Deliverables**:
- `scripts/prepare-release.sh`
- `scripts/create-prerelease.sh`
- `scripts/verify-release.sh`
- `RELEASING.md`

---

## Summary

### Complete Implementation Order

1. **Step R1**: Version File and Build Info
   - Create version.go with build-time injection
   - Add --version flag
   - Update Makefile

2. **Step R2**: Changelog Management
   - Choose manual or automated approach
   - Set up changelog tooling
   - Create templates

3. **Step R3**: Release Workflow Enhancement
   - Update release.yml with changelog extraction
   - Add version injection
   - Improve release notes

4. **Step R4**: Pre-release Workflow
   - Create pre-release.yml
   - Add validation
   - Support alpha/beta/rc

5. **Step R5**: Release Scripts and Documentation
   - Create helper scripts
   - Write RELEASING.md
   - Document processes

### Files Created/Modified

**New Files**:
- `cmd/fof9editor/version.go`
- `.chglog/config.yml` (optional)
- `.chglog/CHANGELOG.tpl.md` (optional)
- `.github/workflows/pre-release.yml`
- `.github/pull_request_template.md` (optional)
- `.github/workflows/changelog-check.yml` (optional)
- `scripts/prepare-release.sh`
- `scripts/create-prerelease.sh`
- `scripts/verify-release.sh`
- `scripts/generate-changelog.sh`
- `RELEASING.md`

**Modified Files**:
- `cmd/fof9editor/main.go` (add --version flag)
- `Makefile` (add ldflags)
- `.github/workflows/release.yml` (enhance)
- `CHANGELOG.md` (format update)

### Benefits

1. **Automation**: Minimal manual steps, consistent process
2. **Traceability**: Every release has version, commit hash, build date
3. **Documentation**: Changelog automatically in release notes
4. **Safety**: Validation, testing, verification steps
5. **Flexibility**: Support for stable and pre-releases
6. **Transparency**: Clear process documented for contributors

### Next Steps After Implementation

1. Create first release following the new process
2. Gather feedback from team
3. Refine scripts based on real usage
4. Consider adding:
   - Automated release announcement (Twitter, Discord, etc.)
   - Download statistics tracking
   - User feedback collection
   - Crash reporting integration
