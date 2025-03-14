package containerservice

import (
	"testing"
)

func TestImageExistsInRegistry(t *testing.T) {
	tests := []struct {
		name         string
		image        string
		tag          string
		expectExists bool
		expectError  bool
	}{
		{
			name:         "Existing Short Image Tag",
			image:        "nginx",
			tag:          "latest",
			expectExists: true,
			expectError:  false,
		},
		{
			name:         "Existing Long Image Tag",
			image:        "crufter/llama-cpp-python",
			tag:          "cuda-12.2.0-latest",
			expectExists: true,
			expectError:  false,
		},
		{
			name:         "Non-Existing Short Tag",
			image:        "nginx",
			tag:          "nonexistenttag123",
			expectExists: false,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists, err := imageExistsInRegistry(tt.image, tt.tag)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error=%v, got error=%v", tt.expectError, err)
			}
			if exists != tt.expectExists {
				t.Errorf("expected exists=%v, got %v", tt.expectExists, exists)
			}
		})
	}
}
