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

# Check and install swag if not installed or version doesn't match
echo "Checking swag installation..."
if command -v swag >/dev/null 2>&1; then
    INSTALLED_SWAG_VERSION=$(swag --version | awk '{print $3}')
else
    INSTALLED_SWAG_VERSION=""
fi
LATEST_SWAG_VERSION=$(go list -m -u -versions -json github.com/swaggo/swag | jq -r '.Versions[-1]')

if [[ -z "$INSTALLED_SWAG_VERSION" ]]; then
    echo "Swag not installed. Installing the latest version..."
    go install github.com/swaggo/swag/cmd/swag@latest
    elif [[ "$INSTALLED_SWAG_VERSION" != "$LATEST_SWAG_VERSION" ]]; then
    echo "Swag version is outdated (installed: $INSTALLED_SWAG_VERSION, latest: $LATEST_SWAG_VERSION). Updating to the latest version..."
    go install github.com/swaggo/swag/cmd/swag@latest
else
    echo "Swag is up to date."
fi

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
