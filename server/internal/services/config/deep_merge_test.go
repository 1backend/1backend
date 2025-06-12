package configservice

import (
	"reflect"
	"testing"
)

func TestDeepMerge(t *testing.T) {
	tests := []struct {
		name string
		dst  map[string]interface{}
		src  map[string]interface{}
		want map[string]interface{}
	}{
		{
			name: "simple shallow merge",
			dst:  map[string]interface{}{"a": 1},
			src:  map[string]interface{}{"b": 2},
			want: map[string]interface{}{"a": 1, "b": 2},
		},
		{
			name: "deep merge nested maps",
			dst: map[string]interface{}{
				"logging": map[string]interface{}{
					"level":  "info",
					"output": "stdout",
				},
			},
			src: map[string]interface{}{
				"logging": map[string]interface{}{
					"level": "debug",
				},
			},
			want: map[string]interface{}{
				"logging": map[string]interface{}{
					"level":  "debug",
					"output": "stdout",
				},
			},
		},
		{
			name: "overwrite on type mismatch",
			dst:  map[string]interface{}{"key": map[string]interface{}{"a": 1}},
			src:  map[string]interface{}{"key": 42},
			want: map[string]interface{}{"key": 42},
		},
		{
			name: "replace arrays entirely",
			dst:  map[string]interface{}{"list": []interface{}{1, 2, 3}},
			src:  map[string]interface{}{"list": []interface{}{4, 5}},
			want: map[string]interface{}{"list": []interface{}{4, 5}},
		},
		{
			name: "add new nested key",
			dst:  map[string]interface{}{"a": map[string]interface{}{"b": 1}},
			src:  map[string]interface{}{"a": map[string]interface{}{"c": 2}},
			want: map[string]interface{}{"a": map[string]interface{}{"b": 1, "c": 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeepMerge(tt.dst, tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
