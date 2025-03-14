#!/bin/bash

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <old_version> <new_version>"
  exit 1
fi

OLD_VERSION=$1
NEW_VERSION=$2

FILES=(
  "./server/main.go"
  "./clients/js/client/package.json"
  "./clients/js/node/package.json"
  "./clients/js/example/package.json"
  "./desktop/package.json"
)

for FILE in "${FILES[@]}"; do
  if [ -f "$FILE" ]; then
    echo "Replacing version in $FILE"
    sed -i "s/$OLD_VERSION/$NEW_VERSION/g" "$FILE"
  else
    echo "File $FILE not found, skipping."
  fi
done

echo "Version bump completed."

