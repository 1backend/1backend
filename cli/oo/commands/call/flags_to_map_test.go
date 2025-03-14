package call

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlagToJSON(t *testing.T) {
	tests := []struct {
		name     string
		flag     string
		expected map[string]interface{}
	}{
		{
			name: "Single Level Flag",
			flag: "--id=1",
			expected: map[string]interface{}{
				"id": 1,
			},
		},
		{
			name: "Repeated Flag as Array",
			flag: "--id=5 --id=5",
			expected: map[string]interface{}{
				"id": []interface{}{5, 5},
			},
		},
		{
			name: "Nested Object Array Flag",
			flag: "--user-id=5 --user-id=5",
			expected: map[string]interface{}{
				"user": []map[string]interface{}{
					{"id": 5},
					{"id": 5},
				},
			},
		},
		{
			name: "Nested Flag",
			flag: "--user-name=joe",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"name": "joe",
				},
			},
		},
		{
			name: "Nested Flag With Dot",
			flag: "--user.name=joe",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"name": "joe",
				},
			},
		},
		{
			name: "Nested Flag with Multiple Levels",
			flag: "--user-name=joe --user-age=30",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"name": "joe",
					"age":  30,
				},
			},
		},
		{
			name: "Nested Flag with Multiple Levels With Dot",
			flag: "--user.name=joe --user.age=30",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"name": "joe",
					"age":  30,
				},
			},
		},
		{
			name: "Flag with Boolean Value",
			flag: "--user-active=true",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"active": true,
				},
			},
		},
		{
			name: "Flag with Boolean Value With Dot",
			flag: "--user-active=true",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"active": true,
				},
			},
		},
		{
			name: "Flag with Integer Value",
			flag: "--user-age=30",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"age": 30,
				},
			},
		},
		{
			name: "Flag with Integer Value With Dot",
			flag: "--user.age=30",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"age": 30,
				},
			},
		},
		{
			name: "Flag with Multiple Nested Objects",
			flag: "--user-id=1 --order-id=100",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"id": 1,
				},
				"order": map[string]interface{}{
					"id": 100,
				},
			},
		},
		{
			name: "Flag with Multiple Nested Objects With Dot",
			flag: "--user.id=1 --order.id=100",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"id": 1,
				},
				"order": map[string]interface{}{
					"id": 100,
				},
			},
		},
		{
			name: "Flag with Complex Nested Structure",
			flag: "--user-info-name=alice --user-info-age=25",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"info": map[string]interface{}{
						"name": "alice",
						"age":  25,
					},
				},
			},
		},
		{
			name: "Flag with Complex Nested Structure With Dot",
			flag: "--user.info.name=alice --user.info.age=25",
			expected: map[string]interface{}{
				"user": map[string]interface{}{
					"info": map[string]interface{}{
						"name": "alice",
						"age":  25,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := make(map[string]interface{})
			args := []string{}
			for _, f := range splitFlags(tt.flag) {
				args = append(args, f)
			}

			for _, arg := range args {
				err := parseFlagToMap(payload, arg)
				assert.NoError(t, err, "expected no error for flag: "+tt.flag)
			}

			assert.Equal(t, tt.expected, payload, "expected payload to match")
		})
	}
}
