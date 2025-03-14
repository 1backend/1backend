#!/bin/bash

HTML_DIR="/usr/share/nginx/html"

if [ -z "$BACKEND_ADDRESS" ]; then
    echo "BACKEND_ADDRESS environment variable is not set. Using default."
else
    echo "Replacing backend address in files with: $BACKEND_ADDRESS"
    find "$HTML_DIR" -type f -exec sed -i "s|http://127.0.0.1:58231|${BACKEND_ADDRESS}|g" {} +
fi

nginx -g "daemon off;"