#!/bin/bash

set -euo pipefail

SOURCE_FILE="./server/main.go"

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <patch|minor|major>"
  exit 1
fi

VERSION_PART="$1"

FILES=(
  "$SOURCE_FILE"
  "./clients/js/client/package.json"
  "./clients/js/node/package.json"
  "./clients/js/example/package.json"
  "./desktop/package.json"
)

if [ ! -f "$SOURCE_FILE" ]; then
  echo "Source file $SOURCE_FILE not found."
  exit 1
fi

CURRENT_VERSION=$(sed -n 's|.*@version[[:space:]]*\([0-9]\+\.[0-9]\+\.[0-9]\+\).*|\1|p' "$SOURCE_FILE" | head -n 1)

if [ -z "$CURRENT_VERSION" ]; then
  echo "Could not determine current version from $SOURCE_FILE"
  exit 1
fi

IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT_VERSION"

case "$VERSION_PART" in
  patch)
    NEW_VERSION="${MAJOR}.${MINOR}.$((PATCH + 1))"
    ;;
  minor)
    NEW_VERSION="${MAJOR}.$((MINOR + 1)).0"
    ;;
  major)
    NEW_VERSION="$((MAJOR + 1)).0.0"
    ;;
  *)
    echo "Usage: $0 <patch|minor|major>"
    exit 1
    ;;
esac

echo "Bumping version: $CURRENT_VERSION -> $NEW_VERSION"

for FILE in "${FILES[@]}"; do
  if [ -f "$FILE" ]; then
    echo "Replacing version in $FILE"
    sed -i "s/$CURRENT_VERSION/$NEW_VERSION/g" "$FILE"
  else
    echo "File $FILE not found, skipping."
  fi
done

echo "Version bump completed."
