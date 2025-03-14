#!/bin/bash

# This is just a quick script intended for local use

set -e

cd client;
npm run build;
npm link;
cd ..

cd example;
npm link @1backend/client;
npm run build
cd ..

cd ../../desktop/workspaces/angular-app/
npm link @1backend/client
cd ../../../clients/js