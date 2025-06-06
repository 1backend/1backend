#!/bin/bash

set -e

# Get the directory of the current script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

bash "$SCRIPT_DIR/include.sh"

# Define relevant paths
OB_DIR="$SCRIPT_DIR/../server"
DOCS_SOURCE_DIR="$SCRIPT_DIR/../docs-source"
DOCS_DIR="$SCRIPT_DIR/../docs"
BUILD_DIR="$DOCS_SOURCE_DIR/build"
SWAGGER_FILE="$OB_DIR/docs/swagger.yaml"
EXAMPLES_DIR="$DOCS_SOURCE_DIR/examples"
CNAME_FILE="$DOCS_SOURCE_DIR/CNAME"

mkdir -p $DOCS_DIR

# Error handler
trap 'echo "Error occurred in script at line $LINENO"; exit 1' ERR

# Initialize Swagger in server
echo "Initializing Swagger in $OB_DIR"
cd "$OB_DIR"
swag init --parseDependency

# Copy Swagger file to docs-source examples
echo "Copying Swagger file to $EXAMPLES_DIR"
cp "$SWAGGER_FILE" "$EXAMPLES_DIR/1backend.yaml"

# Clean and generate API documentation
echo "Cleaning and generating API documentation"
cd "$DOCS_SOURCE_DIR"
yarn clean-api-docs 1backend
yarn gen-api-docs 1backend

# Build the project
echo "Building the project"
npm run build

# Clean and update docs directory
echo "Cleaning up old docs in $DOCS_DIR"
rm -rf "$DOCS_DIR"/*

echo "Copying CNAME file to $DOCS_DIR"
cp "$CNAME_FILE" "$DOCS_DIR/CNAME"

echo "Copying new build files to $DOCS_DIR"
cp -r "$BUILD_DIR"/* "$DOCS_DIR/"

echo "Documentation generation complete."
