package dockerbackend

import (
	"testing"
)

func TestCalculateProgressPercentage(t *testing.T) {
	tests := []struct {
		name     string
		progress string
		expected float64
	}{
		{
			name:     "MB units",
			progress: "[======>]  500.0MB/1.0GB",
			expected: 50.0, // 500MB of 1GB
		},
		{
			name:     "KB units",
			progress: "[==>]  512KB/1MB",
			expected: 51.20, // 512KB of 1MB
		},
		{
			name:     "GB units",
			progress: "[=========>]  1.5GB/3GB",
			expected: 50.0, // 1.5GB of 3GB
		},
		{
			name:     "No progress bar",
			progress: "722.2MB/2.219GB",
			expected: 32.54, // 722.2MB of 2.219GB
		},
		{
			name:     "Invalid format",
			progress: "Invalid data",
			expected: 0.0, // Invalid input should return 0
		},
		{
			name:     "Zero total size",
			progress: "100MB/0B",
			expected: 0.0, // Division by zero should return 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateProgressPercentage(tt.progress)
			if (result < tt.expected-0.1) || (result > tt.expected+0.1) {
				t.Errorf("%s: expected %.2f, got %.2f", tt.name, tt.expected, result)
			}
		})
	}
}

func TestConvertToBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		unit     string
		expected float64
	}{
		{"Bytes", 100, "B", 100},
		{"Kilobytes", 1, "KB", base},
		{"Megabytes", 1, "MB", base * base},
		{"Gigabytes", 1, "GB", base * base * base},
		{"Unknown unit", 500, "TB", 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToBytes(tt.value, tt.unit)
			if result != tt.expected {
				t.Errorf("%s: expected %.0f, got %.0f", tt.name, tt.expected, result)
			}
		})
	}
}
