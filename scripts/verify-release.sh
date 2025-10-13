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
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

gh release download "$TAG" -p "*.exe" -p "checksums.txt"

# 4. Verify checksums
echo -e "\n${YELLOW}Verifying checksums...${NC}"
if sha256sum -c checksums.txt 2>/dev/null; then
    echo -e "${GREEN}Checksums verified!${NC}"
else
    echo -e "${RED}Checksum verification failed!${NC}"
    cd -
    rm -rf "$TEMP_DIR"
    exit 1
fi

# 5. Test executable version (if wine is available)
echo -e "\n${YELLOW}Checking executable version...${NC}"
if command -v wine &> /dev/null; then
    WINE_VERSION=$(wine "fof9editor-${VERSION}-windows-amd64.exe" --version 2>/dev/null | head -n1) || true

    if [ -z "$WINE_VERSION" ]; then
        echo -e "${YELLOW}Warning: Cannot run executable via wine${NC}"
    else
        if echo "$WINE_VERSION" | grep -q "${VERSION}"; then
            echo -e "${GREEN}Version matches: ${WINE_VERSION}${NC}"
        else
            echo -e "${RED}Version mismatch!${NC}"
            echo "Expected: ${VERSION}"
            echo "Got: ${WINE_VERSION}"
            cd -
            rm -rf "$TEMP_DIR"
            exit 1
        fi
    fi
else
    echo -e "${YELLOW}Wine not available - skipping executable version check${NC}"
    echo "Manual verification needed on Windows"
fi

echo -e "\n${GREEN}Release verification complete!${NC}"

# Cleanup
cd -
rm -rf "$TEMP_DIR"
