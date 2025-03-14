/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package dockerbackend

import (
	"testing"
)

func TestTransformModelDir(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no backslash",
			input:    "example",
			expected: "example",
		},
		{
			name:     "backslashes but no drive letter",
			input:    "folder\\subfolder\\file",
			expected: "folder/subfolder/file",
		},
		{
			name:     "drive letter with backslashes",
			input:    "C:\\folder\\subfolder\\file",
			expected: "/mnt/c/folder/subfolder/file",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "drive letter with no path",
			input:    "D:\\",
			expected: "/mnt/d/",
		},
		{
			name:     "drive letter with trailing backslash only",
			input:    "E:\\\\",
			expected: "/mnt/e//",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformWinPaths(tt.input)
			if result != tt.expected {
				t.Errorf(
					"transformModelDir(%q) = %q, want %q",
					tt.input,
					result,
					tt.expected,
				)
			}
		})
	}
}
