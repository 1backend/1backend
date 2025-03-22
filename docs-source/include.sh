#!/bin/bash

# This file idempotently replaces an INCLUDE directive
# in markdown files so we can include real code from files
# in our docusaurus doc site.

set -e

DOCS_DIR="docs"
INCLUDE_START="<!-- INCLUDE: "
INCLUDE_END="<!-- /INCLUDE -->"

# Process all .md files in the docs directory
find "$DOCS_DIR" -name "*.md" | while read -r md_file; do
    echo "Processing $md_file..."

    MD_DIR=$(dirname "$md_file")
    
    awk -v INCLUDE_START="$INCLUDE_START" -v INCLUDE_END="$INCLUDE_END" -v MD_DIR="$MD_DIR" '
    BEGIN { inside_include = 0 }
    {
        # Detect the start of an INCLUDE block
        if ($0 ~ INCLUDE_START) {
            inside_include = 1
            print $0  # Print the INCLUDE comment

            # Extract file path and strip potential comment markers (// or #)
            cmd = "echo \"" $0 "\" | sed -n \"s/.*<!-- INCLUDE: \\(.*\\) -->/\\1/p\" | sed \"s/^\\s*\\/\\///\""
            cmd | getline file
            close(cmd)

            # Resolve relative path based on Markdown files directory
            file = MD_DIR "/" file

            # Only include content if the file exists
            if (system("[ -f \"" file "\" ]") == 0) {
                while ((getline line < file) > 0) print line
                close(file)
            } else {
                print "❌ ERROR: File not found: " file > "/dev/stderr"
                exit 1
            }
        }
        # Detect the end of an INCLUDE block
        else if ($0 ~ INCLUDE_END) {
            inside_include = 0
        }

        # Print non-include lines
        if (!inside_include || $0 ~ INCLUDE_END) {
            print $0
        }
    }' "$md_file" > "$md_file.tmp"
    
    # Exit if `awk` failed (file not found)
    if [ $? -ne 0 ]; then
        echo "❌ Aborting due to missing file in $md_file"
        exit 1
    fi
    
    mv "$md_file.tmp" "$md_file"
done

echo "✅ Code snippets updated in Markdown files."
