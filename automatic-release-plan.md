# Automatic Release Plan

## Overview

This plan implements automatic tag creation and release generation on every merge to the `main` branch.

---

## Design Decisions

### Versioning Strategy

We need to decide how to automatically determine version numbers. Here are the options:

#### Option A: Conventional Commits (Recommended)

Use commit message conventions to automatically determine version bumps:

- `feat:` → MINOR bump (0.1.0 → 0.2.0)
- `fix:` → PATCH bump (0.1.0 → 0.1.1)
- `BREAKING CHANGE:` in commit body → MAJOR bump (0.1.0 → 1.0.0)
- Other types → PATCH bump

**Pros:**
- Industry standard (used by Angular, React, etc.)
- Clear intent from commit messages
- Fully automated version determination

**Cons:**
- Requires discipline in commit messages
- May create releases for trivial changes

#### Option B: Version File

Store version in a file (e.g., `VERSION` or in code), manually update it:

**Pros:**
- Explicit control over version numbers
- Only release when version file changes

**Cons:**
- Requires manual version update
- Reduces automation

#### Option C: Changelog-Based

Parse CHANGELOG.md to determine if a new version exists:

**Pros:**
- Leverages existing changelog workflow
- Explicit version control

**Cons:**
- Complex parsing logic
- Requires CHANGELOG.md to be updated first

---

## Recommended Approach: Conventional Commits + Auto-Release

### Workflow

1. **On merge to main**:
   - Analyze commits since last tag
   - Determine version bump based on conventional commits
   - Create new tag automatically
   - Trigger release workflow

2. **Changelog**:
   - Auto-generate CHANGELOG.md entry from commits
   - Commit back to main (or include in release notes only)

3. **Skip conditions**:
   - Skip if commit message contains `[skip-release]`
   - Skip if only documentation changes
   - Skip if only CI/test file changes

---

## Implementation Steps

### Step 1: Add Conventional Commit Enforcement

Create `.github/workflows/commit-lint.yml`:

```yaml
name: Commit Lint

on:
  pull_request:
    types: [opened, synchronize, reopened]

jobs:
  commitlint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Validate PR title
        uses: amannn/action-semantic-pull-request@v5
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          types: |
            feat
            fix
            docs
            style
            refactor
            perf
            test
            build
            ci
            chore
```

### Step 2: Create Auto-Tag Workflow

Create `.github/workflows/auto-tag.yml`:

```yaml
name: Auto Tag and Release

on:
  push:
    branches:
      - main

permissions:
  contents: write

jobs:
  auto-tag:
    name: Create Tag and Release
    runs-on: ubuntu-latest
    # Skip if commit message contains [skip-release]
    if: "!contains(github.event.head_commit.message, '[skip-release]')"

    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
          token: ${{ secrets.GITHUB_TOKEN }}

      - name: Get latest tag
        id: latest_tag
        run: |
          # Get latest tag, or use 0.0.0 if no tags exist
          LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
          echo "tag=${LATEST_TAG}" >> $GITHUB_OUTPUT
          echo "version=${LATEST_TAG#v}" >> $GITHUB_OUTPUT
          echo "Latest tag: ${LATEST_TAG}"

      - name: Analyze commits
        id: analyze
        run: |
          LATEST_TAG="${{ steps.latest_tag.outputs.tag }}"

          # Get commits since last tag
          if [ "$LATEST_TAG" = "v0.0.0" ]; then
            COMMITS=$(git log --pretty=format:"%s")
          else
            COMMITS=$(git log ${LATEST_TAG}..HEAD --pretty=format:"%s")
          fi

          echo "Commits since ${LATEST_TAG}:"
          echo "$COMMITS"

          # Determine version bump
          MAJOR=0
          MINOR=0
          PATCH=0

          while IFS= read -r commit; do
            # Check for BREAKING CHANGE
            if echo "$commit" | grep -qi "BREAKING CHANGE"; then
              MAJOR=1
            # Check for feat:
            elif echo "$commit" | grep -qi "^feat"; then
              if [ $MAJOR -eq 0 ]; then
                MINOR=1
              fi
            # Check for fix:
            elif echo "$commit" | grep -qi "^fix"; then
              if [ $MAJOR -eq 0 ] && [ $MINOR -eq 0 ]; then
                PATCH=1
              fi
            fi
          done <<< "$COMMITS"

          # If no conventional commits found, do PATCH bump
          if [ $MAJOR -eq 0 ] && [ $MINOR -eq 0 ] && [ $PATCH -eq 0 ]; then
            PATCH=1
          fi

          echo "major=${MAJOR}" >> $GITHUB_OUTPUT
          echo "minor=${MINOR}" >> $GITHUB_OUTPUT
          echo "patch=${PATCH}" >> $GITHUB_OUTPUT

      - name: Calculate new version
        id: new_version
        run: |
          CURRENT="${{ steps.latest_tag.outputs.version }}"
          MAJOR="${{ steps.analyze.outputs.major }}"
          MINOR="${{ steps.analyze.outputs.minor }}"
          PATCH="${{ steps.analyze.outputs.patch }}"

          # Parse current version
          IFS='.' read -r -a VERSION_PARTS <<< "$CURRENT"
          V_MAJOR="${VERSION_PARTS[0]}"
          V_MINOR="${VERSION_PARTS[1]}"
          V_PATCH="${VERSION_PARTS[2]}"

          # Calculate new version
          if [ $MAJOR -eq 1 ]; then
            NEW_VERSION="$((V_MAJOR + 1)).0.0"
            BUMP_TYPE="major"
          elif [ $MINOR -eq 1 ]; then
            NEW_VERSION="${V_MAJOR}.$((V_MINOR + 1)).0"
            BUMP_TYPE="minor"
          else
            NEW_VERSION="${V_MAJOR}.${V_MINOR}.$((V_PATCH + 1))"
            BUMP_TYPE="patch"
          fi

          echo "version=${NEW_VERSION}" >> $GITHUB_OUTPUT
          echo "tag=v${NEW_VERSION}" >> $GITHUB_OUTPUT
          echo "bump_type=${BUMP_TYPE}" >> $GITHUB_OUTPUT

          echo "Bump type: ${BUMP_TYPE}"
          echo "New version: ${NEW_VERSION}"

      - name: Extract commits for changelog
        id: changelog
        run: |
          LATEST_TAG="${{ steps.latest_tag.outputs.tag }}"
          NEW_VERSION="${{ steps.new_version.outputs.version }}"

          # Get commits since last tag
          if [ "$LATEST_TAG" = "v0.0.0" ]; then
            COMMITS=$(git log --pretty=format:"- %s (%h)" --no-merges)
          else
            COMMITS=$(git log ${LATEST_TAG}..HEAD --pretty=format:"- %s (%h)" --no-merges)
          fi

          # Create changelog content
          cat > /tmp/release_notes.md <<EOF
          ## Changes in v${NEW_VERSION}

          ${COMMITS}

          ---

          **Auto-generated release** based on conventional commits.

          See [CHANGELOG.md](https://github.com/${{ github.repository }}/blob/main/CHANGELOG.md) for detailed changes.
          EOF

          echo "changelog_file=/tmp/release_notes.md" >> $GITHUB_OUTPUT

      - name: Update CHANGELOG.md
        run: |
          NEW_VERSION="${{ steps.new_version.outputs.version }}"
          NEW_TAG="${{ steps.new_version.outputs.tag }}"
          TODAY=$(date +%Y-%m-%d)
          LATEST_TAG="${{ steps.latest_tag.outputs.tag }}"

          # Get categorized commits
          if [ "$LATEST_TAG" = "v0.0.0" ]; then
            FEAT_COMMITS=$(git log --pretty=format:"- %s" --no-merges | grep "^- feat:" | sed 's/^- feat: /- /' || echo "")
            FIX_COMMITS=$(git log --pretty=format:"- %s" --no-merges | grep "^- fix:" | sed 's/^- fix: /- /' || echo "")
            OTHER_COMMITS=$(git log --pretty=format:"- %s" --no-merges | grep -v "^- feat:" | grep -v "^- fix:" || echo "")
          else
            FEAT_COMMITS=$(git log ${LATEST_TAG}..HEAD --pretty=format:"- %s" --no-merges | grep "^- feat:" | sed 's/^- feat: /- /' || echo "")
            FIX_COMMITS=$(git log ${LATEST_TAG}..HEAD --pretty=format:"- %s" --no-merges | grep "^- fix:" | sed 's/^- fix: /- /' || echo "")
            OTHER_COMMITS=$(git log ${LATEST_TAG}..HEAD --pretty=format:"- %s" --no-merges | grep -v "^- feat:" | grep -v "^- fix:" || echo "")
          fi

          # Create new version entry
          NEW_ENTRY="## [${NEW_VERSION}] - ${TODAY}\n\n"

          if [ -n "$FEAT_COMMITS" ]; then
            NEW_ENTRY="${NEW_ENTRY}### Added\n${FEAT_COMMITS}\n\n"
          fi

          if [ -n "$FIX_COMMITS" ]; then
            NEW_ENTRY="${NEW_ENTRY}### Fixed\n${FIX_COMMITS}\n\n"
          fi

          if [ -n "$OTHER_COMMITS" ]; then
            NEW_ENTRY="${NEW_ENTRY}### Changed\n${OTHER_COMMITS}\n\n"
          fi

          # Insert after [Unreleased] section
          awk -v new="$NEW_ENTRY" '
            /## \[Unreleased\]/ {
              print
              getline
              print
              print new
              next
            }
            {print}
          ' CHANGELOG.md > CHANGELOG.md.tmp

          mv CHANGELOG.md.tmp CHANGELOG.md

          # Update version comparison links
          sed -i "s|\[Unreleased\]:.*|[Unreleased]: https://github.com/${{ github.repository }}/compare/${NEW_TAG}...HEAD\n[${NEW_VERSION}]: https://github.com/${{ github.repository }}/releases/tag/${NEW_TAG}|" CHANGELOG.md

      - name: Commit CHANGELOG update
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"
          git add CHANGELOG.md
          git commit -m "chore: update CHANGELOG for v${{ steps.new_version.outputs.version }} [skip-release]

Auto-generated changelog entry for release v${{ steps.new_version.outputs.version }}

🤖 Generated by GitHub Actions" || echo "No changes to commit"
          git push origin main

      - name: Create and push tag
        run: |
          NEW_TAG="${{ steps.new_version.outputs.tag }}"
          NEW_VERSION="${{ steps.new_version.outputs.version }}"

          git tag -a "${NEW_TAG}" -m "Release ${NEW_VERSION}

Auto-generated release based on conventional commits

🤖 Generated by GitHub Actions"

          git push origin "${NEW_TAG}"

          echo "✅ Created and pushed tag: ${NEW_TAG}"

      - name: Wait for release workflow
        run: |
          echo "Tag pushed. Release workflow will be triggered automatically."
          echo "Monitor at: https://github.com/${{ github.repository }}/actions"
```

### Step 3: Update PR Template

Update `.github/pull_request_template.md` to include conventional commit guidance:

```markdown
# Description

<!-- Describe your changes in detail -->

## Type of Change

**IMPORTANT**: The PR title should follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat: description` - New feature (triggers MINOR version bump)
- `fix: description` - Bug fix (triggers PATCH version bump)
- `docs: description` - Documentation only changes
- `style: description` - Code style changes (formatting, etc.)
- `refactor: description` - Code refactoring
- `perf: description` - Performance improvements
- `test: description` - Adding or updating tests
- `build: description` - Build system changes
- `ci: description` - CI configuration changes
- `chore: description` - Other changes

For **BREAKING CHANGES**, add `BREAKING CHANGE:` in the commit body (triggers MAJOR version bump).

## Changelog Entry

<!-- Your changes will be automatically added to CHANGELOG.md based on your commit messages -->

## Testing

<!-- Describe the tests you ran to verify your changes -->

- [ ] All existing tests pass
- [ ] Added new tests for new functionality
- [ ] Manually tested the changes

## Checklist

- [ ] PR title follows Conventional Commits format
- [ ] Code follows the project's style guidelines
- [ ] Documentation has been updated
- [ ] All tests pass locally
```

### Step 4: Update RELEASING.md

Add section about automatic releases:

```markdown
## Automatic Releases

Every merge to `main` automatically creates a new release.

### How It Works

1. **Commit to main** (via merge)
2. **Analyze commits** since last tag using Conventional Commits
3. **Calculate version bump**:
   - `feat:` commits → MINOR bump (0.1.0 → 0.2.0)
   - `fix:` commits → PATCH bump (0.1.0 → 0.1.1)
   - `BREAKING CHANGE:` in body → MAJOR bump (0.1.0 → 1.0.0)
4. **Update CHANGELOG.md** automatically
5. **Create and push tag** (e.g., v0.2.0)
6. **Release workflow triggers** automatically
7. **GitHub Release created** with binaries

### Skip Auto-Release

To skip automatic release, include `[skip-release]` in commit message:

```bash
git commit -m "docs: update README [skip-release]"
```

### Manual Release Override

If you need to create a release manually (bypassing auto-release):

1. Skip the auto-release: include `[skip-release]` in the merge commit
2. Use the manual release process (see above)
```

---

## Alternative: Simpler Approach

If full conventional commits automation is too complex, here's a simpler approach:

### Simpler Auto-Release Workflow

Only create releases when a version bump commit is detected:

```yaml
name: Auto Release on Version Bump

on:
  push:
    branches:
      - main
    paths:
      - 'CHANGELOG.md'

permissions:
  contents: write

jobs:
  check-and-release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Check for new version in CHANGELOG
        id: check_version
        run: |
          # Extract latest version from CHANGELOG (after Unreleased)
          NEW_VERSION=$(awk '/## \[[0-9]/ {print $2; exit}' CHANGELOG.md | tr -d '[]')

          # Get latest git tag
          LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
          LATEST_VERSION=${LATEST_TAG#v}

          echo "Latest version from CHANGELOG: ${NEW_VERSION}"
          echo "Latest version from git: ${LATEST_VERSION}"

          if [ "$NEW_VERSION" != "$LATEST_VERSION" ]; then
            echo "new_version=${NEW_VERSION}" >> $GITHUB_OUTPUT
            echo "should_release=true" >> $GITHUB_OUTPUT
          else
            echo "should_release=false" >> $GITHUB_OUTPUT
          fi

      - name: Create tag and trigger release
        if: steps.check_version.outputs.should_release == 'true'
        run: |
          NEW_VERSION="${{ steps.check_version.outputs.new_version }}"
          NEW_TAG="v${NEW_VERSION}"

          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"

          git tag -a "${NEW_TAG}" -m "Release ${NEW_VERSION}"
          git push origin "${NEW_TAG}"

          echo "✅ Created release tag: ${NEW_TAG}"
```

This simpler approach:
- Only triggers when CHANGELOG.md changes
- Manually update CHANGELOG.md with new version
- Auto-tags and releases when version in CHANGELOG differs from latest tag

---

## Recommendation

**Start with the simpler approach**, then migrate to full Conventional Commits if needed.

**Why?**
- Less complexity to start
- Leverages existing CHANGELOG workflow
- Still automates the tag/release creation
- Can add Conventional Commits later if desired

**Workflow:**
1. Developer updates CHANGELOG.md with new version section
2. PR merged to main
3. Auto-tag workflow detects new version
4. Creates tag automatically
5. Release workflow triggers
6. Release published

This gives you automation without forcing strict commit message conventions immediately.
