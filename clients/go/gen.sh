#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

SERVER_DIR="$SCRIPT_DIR/../../server"

cd "$SERVER_DIR"
swag init --parseDependency

cd "$SCRIPT_DIR"
rm -rf *.go
openapi-generator-cli generate \
-i "$SERVER_DIR/docs/swagger.yaml" \
-g go \
-o . \
--additional-properties generateInterfaces=true \

rm -rf api docs go.mod
cp go.mod.template go.mod

sed -i 's|github.com/GIT_USER_ID/GIT_REPO_ID|github.com/1backend/1backend/clients/go|g' test/*.go
go mod tidy



echo "Replacing []map[string]interface{} with []any in Go files after marker: $MARKER"

# Find all .go files and replace only lines after the marker
MARKER="@openapi-any-array"

find ./ -name "*.go" | while read -r file; do
  echo "Processing file: $file"
  awk -v marker="$MARKER" '
    {
      # Trim leading spaces/tabs for comparison
      line = $0
      gsub(/^[ \t]+/, "", line)

      if (last_line_marker && $0 ~ /\[\]map\[string\]interface\{\}/) {
        sub(/\[\]map\[string\]interface\{\}/, "[]any")
        last_line_marker = 0
      }

      if (index(line, marker) > 0) {
        last_line_marker = 1
      } else {
        last_line_marker = 0
      }

      print
    }
  ' "$file" > "$file.tmp"
  mv "$file.tmp" "$file" || true
done

# Delete getter methods because they are still using []map[string]interface{}
# What a cruft.
find ./ "model_*.go" | while read -r file; do
  echo "Cleaning methods in: $file"
  
  awk '
    BEGIN { skip=0 }
    # Match start of comment for Get*/Get*Ok/Set*
    /^\/\/ (Get|Set)[A-Z]/ { skip=1; next }
    # Match function definitions for Get*/Get*Ok/Set*
    /^func \(o \*/ && / (Get|Set)[A-Z]/ { skip=1; next }
    # End skipping at closing brace of the function
    skip && /^}/ { skip=0; next }
    # Skip function body lines
    skip { next }
    # Keep all other lines
    { print }
  ' "$file" > "$file.tmp"
  mv "$file.tmp" "$file" || true
done

bash "$SERVER_DIR/mock_go.sh"

# A hack to fix gopls because it doesn't like lots of client files being regenerated
if command -v gopls &> /dev/null; then
    killall gopls || true
else
    echo "gopls is not running, skipping kill command."
fi
