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

# 6. Update CHANGELOG.md Unreleased section if needed
TODAY=$(date +%Y-%m-%d)
if grep -q "## \[${VERSION}\] - TBD" CHANGELOG.md; then
    echo -e "\n${YELLOW}Updating CHANGELOG.md date...${NC}"
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        sed -i '' "s/## \[${VERSION}\] - TBD/## [${VERSION}] - ${TODAY}/" CHANGELOG.md
    else
        # Linux
        sed -i "s/## \[${VERSION}\] - TBD/## [${VERSION}] - ${TODAY}/" CHANGELOG.md
    fi

    # Commit changelog changes if modified
    if [ -n "$(git status --porcelain CHANGELOG.md)" ]; then
        git add CHANGELOG.md
        git commit -m "chore: prepare release ${VERSION}

Updated CHANGELOG.md for version ${VERSION}

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>"
    fi
fi

# 7. Create and push tag
echo -e "\n${YELLOW}Creating tag ${TAG}...${NC}"
git tag -a "${TAG}" -m "Release ${VERSION}"

echo -e "\n${GREEN}Release preparation complete!${NC}"
echo -e "\nNext steps:"
echo "1. Review the changes"
echo "2. Push the tag: ${YELLOW}git push origin ${TAG}${NC}"
echo "3. Monitor the release workflow: ${YELLOW}gh run watch${NC}"
echo "4. Verify release page: ${YELLOW}gh release view ${TAG}${NC}"
