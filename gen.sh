#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Exit if any command fails
set -e

# Check if 'go' is installed
if ! command_exists go; then
    echo "Error: 'go' is not installed. Please install Go to proceed."
    exit 1
fi

# Check if 'npm' is installed
if ! command_exists npm; then
    echo "Error: 'npm' is not installed. Please install npm to proceed."
    exit 1
fi

TARGET_SWAG_VERSION="v2.0.0-rc4"
INSTALL_PATH="$(go env GOPATH)/bin/swag"

echo "Checking swag installation..."
ORIGINAL_DIR="$(pwd)"

# Get current swag version if installed
if command -v swag >/dev/null 2>&1; then
    INSTALLED_SWAG_VERSION=$(swag --version | sed -n '2p' | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' || true)
else
    INSTALLED_SWAG_VERSION=""
fi

echo "Current swag version: ${INSTALLED_SWAG_VERSION:-none}"

ROUGH_SWAG_VERSION=$(echo "v2.0.0-rc4" | sed 's/-.*//')

if [[ "$INSTALLED_SWAG_VERSION" != "$ROUGH_SWAG_VERSION" ]]; then
    echo "Swag version mismatch or not installed (installed: '${INSTALLED_SWAG_VERSION:-none}, expected: '$TARGET_SWAG_VERSION'). Installing..."

    TMP_DIR=$(mktemp -d)
    trap 'rm -rf "$TMP_DIR"' EXIT

    git clone --depth 1 --branch "$TARGET_SWAG_VERSION" https://github.com/swaggo/swag.git "$TMP_DIR/swag"
    cd "$TMP_DIR/swag/cmd/swag"
    go build -o "$INSTALL_PATH"

    echo "Installed swag version $TARGET_SWAG_VERSION to $INSTALL_PATH"
else
    echo "Swag is already at the correct version ($INSTALLED_SWAG_VERSION)."
fi

cd "$ORIGINAL_DIR"

# Check and install openapi-generator-cli if not installed or version doesn't match
echo "Checking openapi-generator-cli installation..."
INSTALLED_OPENAPI_VERSION=$(npm list -g @openapitools/openapi-generator-cli --depth=0 2>/dev/null | grep @openapitools/openapi-generator-cli | awk -F@ '{print $3}')
LATEST_OPENAPI_VERSION=$(npm show @openapitools/openapi-generator-cli version)

if [[ -z "$INSTALLED_OPENAPI_VERSION" ]]; then
    echo "openapi-generator-cli not installed. Installing..."
    npm install @openapitools/openapi-generator-cli -g
    elif [[ "$INSTALLED_OPENAPI_VERSION" != "$LATEST_OPENAPI_VERSION" ]]; then
    echo "openapi-generator-cli version is outdated. Updating..."
    npm install @openapitools/openapi-generator-cli -g
else
    echo "openapi-generator-cli is up to date."
fi

# Set the latest version of openapi-generator-cli
echo "Setting openapi-generator-cli to use the latest version..."
openapi-generator-cli version-manager set latest

# Run the generation scripts
echo "Running Go client generation script..."
bash clients/go/gen.sh

echo "Running JavaScript client generation script..."
bash clients/js/gen.sh

echo "Running documentation source generation script..."
bash docs-source/gen.sh

echo "All tasks completed successfully."
