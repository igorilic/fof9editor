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
