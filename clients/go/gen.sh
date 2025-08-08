#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

SERVER_DIR="$SCRIPT_DIR/../../server"

cd "$SERVER_DIR"
swag init --parseDependency --v3.1

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

find . -type f -name "*.go" ! -path "*/vendor/*" | while read -r file; do
  echo "Processing $file"

  # Replace method names like:
  # func (r Receiver) SomeSvcSomethingRequest(...) → func (r Receiver) Body(...)
  sed -i.bak -E \
    's/^(func\s*\([^)]+\)\s+)[A-Za-z0-9]+Svc[A-Za-z0-9]*Request(\s*\()/\1Body\2/' "$file"
done

bash "$SERVER_DIR/mock_go.sh"

# A hack to fix gopls because it doesn't like lots of client files being regenerated
if command -v gopls &> /dev/null; then
    killall gopls || true
else
    echo "gopls is not running, skipping kill command."
fi
