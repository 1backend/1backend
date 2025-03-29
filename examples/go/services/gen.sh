#!/bin/bash

set -e

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

echo "Setting openapi-generator-cli to use the latest version..."
openapi-generator-cli version-manager set latest

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

for dir in */ ; do
    if [[ "$dir" == "scripts/" ]]; then
        continue
    fi
    
    pwd
    serviceName="${dir%/}"
    clientsDir="./$dir/client"
    echo "Processing $dir"
    
    mkdir -p "$clientsDir/spec"
    
    swag init --parseDependency --dir "$serviceName" --output "$clientsDir/spec"
    
    mkdir -p "$clientsDir/go"
    
    openapi-generator-cli generate \
    -i "$clientsDir/spec/swagger.yaml" \
    -g go \
    -o "$clientsDir/go" \
    --additional-properties generateInterfaces=true
    
    find "$clientsDir/go" -type f -name "*" -exec sed -i "s|github.com/GIT_USER_ID/GIT_REPO_ID|github.com/1backend/1backend/examples/go/services/$serviceName/client|g" {} +
    
    mkdir -p "$clientsDir/js/src"
    
    rm -r "$clientsDir/js/src/*" || true
    rm -r "$clientsDir/js/dist/*" || true
    
    echo "Generating into $clientsDir/js/src"
    
    openapi-generator-cli generate \
    -i "$clientsDir/spec/swagger.yaml" \
    -g typescript-fetch \
    -o "$clientsDir/js/src"
    
    cd "$clientsDir/js"
    
    echo "Building $serviceName client in $(pwd)..."
    npm install
    npm run build
    
    cd ../../..
done

# Kill gopls if running to avoid cache issues
if command -v gopls &> /dev/null; then
    killall gopls || true
else
    echo "gopls is not running, skipping kill command."
fi

echo "Client generation complete."