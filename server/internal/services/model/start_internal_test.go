package modelservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectCudaVersionLength(t *testing.T) {
	tests := []struct {
		name        string
		precision   int
		cudaVersion string
		expected    string
	}{
		{"No change needed", 2, "12.2", "12.2"},
		{"Expand to match image tag format", 3, "12.2", "12.2.0"},
		{"Already correct format", 3, "12.2.0", "12.2.0"},
		{"Precision exceeds available segments", 4, "12.2.0", "12.2.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := correctCudaVersionLength(tt.precision, tt.cudaVersion)
			assert.Equal(t, tt.expected, result)
		})
	}
}
