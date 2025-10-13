# Release Process

This document describes the release process for FOF9 Editor.

## Release Types

### Stable Releases (vX.Y.Z)
- Production-ready releases
- Full changelog required
- All tests must pass
- Released from `main` branch
- **Automatically released** when CHANGELOG.md is updated with new version

### Pre-Releases
- **Alpha** (vX.Y.Z-alpha.N): Early development, unstable, major bugs expected
- **Beta** (vX.Y.Z-beta.N): Feature complete, testing phase, minor bugs possible
- **RC** (vX.Y.Z-rc.N): Release candidate, final testing before stable release

---

## 🤖 Automatic Releases (Recommended)

**Every merge to `main` with a CHANGELOG.md update automatically creates a release!**

### How It Works

1. **Update CHANGELOG.md** in your PR with a new version section
2. **Merge PR to main**
3. **Auto-tag workflow runs**:
   - Detects new version in CHANGELOG.md
   - Validates version format (X.Y.Z)
   - Creates git tag automatically (e.g., v0.2.0)
4. **Release workflow triggers** automatically
5. **Release published** with Windows binaries

### Workflow

```markdown
## [Unreleased]

### Added
- Nothing yet

## [0.2.0] - 2025-10-15    ← Add this section in your PR

### Added
- New feature X
- New feature Y

### Fixed
- Bug Z
```

**That's it!** When the PR is merged:
- ✅ Tag `v0.2.0` is created automatically
- ✅ Release workflow builds Windows executable
- ✅ GitHub Release is published
- ✅ Changelog content is extracted to release notes

### Requirements for Auto-Release

1. ✅ Version follows semantic versioning: `X.Y.Z` (e.g., 0.2.0, 1.0.0)
2. ✅ Version has a date (not "TBD"): `## [0.2.0] - 2025-10-15`
3. ✅ Version is newer than the latest git tag
4. ✅ CHANGELOG.md is updated in the commit
5. ✅ **`RELEASE_TOKEN` secret configured** (Personal Access Token with `repo` scope)
   - See [PAT Token Setup Guide](.github/PAT_TOKEN_SETUP.md) for instructions

### Example PR Workflow

```bash
# 1. Create feature branch
git checkout -b feature/new-dashboard

# 2. Make your changes
# ... code changes ...

# 3. Update CHANGELOG.md
# Add new version section with today's date

# 4. Commit and push
git add .
git commit -m "feat: add new dashboard

- Added dashboard component
- Added metrics visualization
- Updated navigation"

git push origin feature/new-dashboard

# 5. Create PR and merge to main
# → Auto-tag workflow creates v0.2.0
# → Release workflow builds and publishes
```

---

## 📝 Manual Releases (Alternative)

If you prefer manual control, you can still use the scripts:

### Creating a Stable Release

### 1. Update CHANGELOG.md

Before starting the release process, ensure CHANGELOG.md is up to date:

```markdown
## [Unreleased]

### Added
- New features added since last release

### Changed
- Changes to existing functionality

### Fixed
- Bug fixes

## [0.2.0] - TBD

### Added
- Feature X
- Feature Y

### Fixed
- Bug Z
```

### 2. Prepare the Release

```bash
# Ensure you're on main branch with latest changes
git checkout main
git pull origin main

# Run the preparation script
./scripts/prepare-release.sh 0.2.0
```

This script will:
- ✓ Verify you're on the main branch
- ✓ Check for uncommitted changes
- ✓ Pull latest changes from remote
- ✓ Run all tests
- ✓ Verify CHANGELOG.md contains the version entry
- ✓ Update TBD date to today's date in CHANGELOG.md
- ✓ Create a commit with changelog updates (if needed)
- ✓ Create a git tag (e.g., v0.2.0)

### 3. Push the Tag

```bash
# Push the tag to trigger release workflow
git push origin v0.2.0

# Monitor the release workflow
gh run watch
```

The GitHub Actions workflow will:
- Run tests
- Build Windows executable with version info injected
- Create release archive (ZIP with exe, README, CHANGELOG)
- Generate checksums
- Extract changelog content for this version
- Create GitHub Release with:
  - Version-specific changelog entries
  - Download links
  - Installation instructions
  - Build information
  - Checksums file

### 4. Verify the Release

```bash
# Wait for workflow to complete, then verify
./scripts/verify-release.sh v0.2.0

# Check the release page
gh release view v0.2.0

# Open release page in browser
gh release view v0.2.0 --web
```

---

## Creating a Pre-Release

Pre-releases are useful for:
- Testing new features with early adopters
- Getting feedback before stable release
- Finding bugs in a near-release state

### Process

```bash
# Create a beta pre-release
./scripts/create-prerelease.sh 0.2.0-beta.1 beta

# Monitor the workflow
gh run watch
```

Version format must match: `X.Y.Z-TYPE.N` where:
- X.Y.Z is the target version
- TYPE is alpha, beta, or rc
- N is the iteration number (1, 2, 3, etc.)

Examples:
- `0.2.0-alpha.1` - First alpha of v0.2.0
- `0.2.0-beta.2` - Second beta of v0.2.0
- `0.2.0-rc.1` - First release candidate of v0.2.0

The pre-release workflow will:
- Validate version format
- Run tests
- Build executable
- Extract Unreleased section from CHANGELOG
- Create release with warning banner
- Mark as pre-release on GitHub

---

## Updating CHANGELOG.md

We follow [Keep a Changelog](https://keepachangelog.com/) format.

### Format

```markdown
## [Unreleased]

### Added
- New features not yet released

### Changed
- Changes to existing functionality

### Deprecated
- Features that will be removed in future

### Removed
- Features that were removed

### Fixed
- Bug fixes

### Security
- Security vulnerability fixes

## [0.2.0] - 2025-10-15

### Added
- Feature X that does Y
- Feature Z with capability W

### Fixed
- Bug #123: Issue with player import
```

### When to Update

1. **During Development**: Add entries to `[Unreleased]` section
2. **Before Release**: Create new version section (e.g., `## [0.2.0] - TBD`)
3. **During Release**: Script updates TBD to actual date

### Guidelines

- Use clear, user-focused language
- Reference issue numbers when applicable
- Group related changes together
- Don't include internal refactoring (unless it affects users)

---

## Versioning Guidelines

We follow [Semantic Versioning 2.0.0](https://semver.org/).

### When to Bump MAJOR (X.0.0)

Breaking changes that require user action:
- Breaking changes to .fof9proj file format
- Removal of CSV columns
- Changed CSV column meanings
- Incompatible changes to data models
- Major UI redesign requiring user relearning

### When to Bump MINOR (0.X.0)

New features (backward compatible):
- New entity types (e.g., adding schedules support)
- New CSV columns (optional, don't break existing files)
- New UI sections or dialogs
- New import/export formats
- New validation rules
- Enhanced functionality

### When to Bump PATCH (0.0.X)

Bug fixes and minor improvements:
- Bug fixes that don't change functionality
- UI polish and minor improvements
- Performance improvements
- Documentation updates
- Dependency updates (security or bug fixes)
- Internal refactoring

### Pre-release Numbering

- Start with `.1` for first pre-release
- Increment for each iteration
- After RC, next version is stable (no .0)

Example progression:
```
0.2.0-alpha.1 → 0.2.0-alpha.2 → 0.2.0-beta.1 → 0.2.0-rc.1 → 0.2.0
```

---

## Hotfix Process

For urgent fixes needed in production:

```bash
# 1. Create hotfix branch from tag
git checkout -b hotfix/0.1.1 v0.1.0

# 2. Make the fix and commit
# ... make changes ...
git commit -m "fixed: critical bug in player import

Fixed issue where player import would fail with non-ASCII characters

Fixes #456

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>"

# 3. Update CHANGELOG.md
# Add ## [0.1.1] - YYYY-MM-DD with fix description

# 4. Run tests
go test ./internal/... -v

# 5. Tag and push
git tag -a v0.1.1 -m "Hotfix 0.1.1"
git push origin hotfix/0.1.1
git push origin v0.1.1

# 6. Merge back to main
git checkout main
git merge hotfix/0.1.1
git push origin main

# 7. Delete hotfix branch
git branch -d hotfix/0.1.1
git push origin --delete hotfix/0.1.1
```

---

## Rollback Process

If a release has critical issues:

### Option 1: Mark as Draft (Hide from Users)

```bash
# Mark release as draft
gh release edit v0.2.0 --draft

# Add warning to release notes
gh release edit v0.2.0 --notes "⚠️ WARNING: This release has been pulled due to critical issue #123.

Please use v0.1.0 instead.

Issue: [Brief description of problem]
Status: Fixed in v0.2.1 (coming soon)

Do not download or use this version."
```

### Option 2: Delete Release (Nuclear Option)

```bash
# Delete the release (keeps tag)
gh release delete v0.2.0 --yes

# Delete the tag
git tag -d v0.2.0
git push origin :refs/tags/v0.2.0

# Create hotfix release
# ... follow hotfix process ...
```

---

## Release Checklist

Use this checklist before creating a release:

### Pre-Release Checks

- [ ] All tests passing (`go test ./internal/... -v`)
- [ ] CHANGELOG.md updated with all changes
- [ ] README.md reflects new version (if needed)
- [ ] Documentation updated (if new features added)
- [ ] No TODO/FIXME comments in critical code
- [ ] Manual testing completed on Windows
- [ ] Known issues documented in CHANGELOG or GitHub Issues
- [ ] Branch is up to date with remote main

### Release Execution

- [ ] Ran `./scripts/prepare-release.sh X.Y.Z`
- [ ] Pushed tag: `git push origin vX.Y.Z`
- [ ] Monitored workflow: `gh run watch`
- [ ] Workflow completed successfully

### Post-Release Verification

- [ ] Release page has correct version number
- [ ] Changelog content extracted correctly
- [ ] Download links work (exe and zip)
- [ ] Checksums match downloaded files
- [ ] Executable shows correct version (`--version` flag)
- [ ] Release marked correctly (stable vs pre-release)
- [ ] Ran `./scripts/verify-release.sh vX.Y.Z`

### Post-Release Tasks

- [ ] Verify downloads work for end users
- [ ] Update any external documentation (wiki, etc.)
- [ ] Announce release (optional: social media, Discord, etc.)
- [ ] Monitor for issue reports
- [ ] Close milestone (if using GitHub milestones)

---

## Troubleshooting

### Release workflow failed

```bash
# Check workflow logs
gh run list --workflow=release.yml --limit 5
gh run view <run-id> --log-failed

# Common issues:
# 1. Tests failed - Fix failing tests and re-tag
# 2. Changelog not found - Ensure version exists in CHANGELOG.md
# 3. Build failed - Check Go version and dependencies

# To retry:
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

# Or delete and re-upload
gh release delete-asset v0.2.0 <wrong-file>
gh release upload v0.2.0 <correct-file>
```

### Changelog extraction failed

```bash
# Manually edit release notes
gh release edit v0.2.0

# Or edit via web interface
gh release view v0.2.0 --web
```

### Version shows as "dev" in executable

This means ldflags didn't inject the version. Check:

1. Makefile has correct ldflags
2. GitHub Actions workflow has correct ldflags
3. Version format is correct

```bash
# Test locally
make build
./bin/fof9editor.exe --version
```

### Tag already exists

```bash
# List tags
git tag -l

# Delete local tag
git tag -d v0.2.0

# Delete remote tag
git push origin :refs/tags/v0.2.0

# Recreate
git tag -a v0.2.0 -m "Release 0.2.0"
```

---

## Manual Release (Fallback)

If GitHub Actions is unavailable, you can create a release manually:

```bash
# 1. Build locally on Windows
make build

# 2. Rename with version
mv bin/fof9editor.exe fof9editor-0.2.0-windows-amd64.exe

# 3. Create ZIP archive
mkdir release-0.2.0
cp fof9editor-0.2.0-windows-amd64.exe release-0.2.0/
cp README.md CHANGELOG.md release-0.2.0/
cd release-0.2.0
zip ../fof9editor-0.2.0-windows-amd64.zip *
cd ..

# 4. Generate checksums
sha256sum fof9editor-0.2.0-windows-amd64.exe > checksums.txt
sha256sum fof9editor-0.2.0-windows-amd64.zip >> checksums.txt

# 5. Create release manually
gh release create v0.2.0 \
  --title "Release v0.2.0" \
  --notes "See CHANGELOG.md" \
  fof9editor-0.2.0-windows-amd64.exe \
  fof9editor-0.2.0-windows-amd64.zip \
  checksums.txt
```

---

## Release Schedule

While we don't have a fixed schedule, here are general guidelines:

- **Major releases** (X.0.0): Yearly or when breaking changes accumulate
- **Minor releases** (0.X.0): Every 2-3 months or when significant features ready
- **Patch releases** (0.0.X): As needed for bug fixes, typically within days
- **Pre-releases**: As needed during development of major/minor releases

---

## Communication

### Release Announcements

Consider announcing releases through:

1. **GitHub Release Page**: Automatically created
2. **Repository README**: Update "Latest Release" badge/link
3. **Social Media**: Twitter, Reddit, etc. (optional)
4. **Community Forums**: FOF9 community forums (optional)
5. **Discord/Slack**: Project community channels (if any)

### Template for Announcements

```
🎉 FOF9 Editor v0.2.0 Released!

We're excited to announce the release of FOF9 Editor v0.2.0!

✨ Highlights:
- Feature X for easier Y
- Improved Z performance by 50%
- Fixed issues with W

📥 Download: https://github.com/igorilic/fof9editor/releases/tag/v0.2.0

📋 Full changelog: https://github.com/igorilic/fof9editor/blob/main/CHANGELOG.md

🐛 Issues? Report at: https://github.com/igorilic/fof9editor/issues
```

---

## Additional Resources

- [Semantic Versioning](https://semver.org/)
- [Keep a Changelog](https://keepachangelog.com/)
- [GitHub Releases Documentation](https://docs.github.com/en/repositories/releasing-projects-on-github)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
