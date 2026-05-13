package configservice

import (
	"reflect"
	"testing"

	types "github.com/1backend/1backend/server/internal/services/config/types"
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

func TestApplyConfigPatchRemove(t *testing.T) {
	tests := []struct {
		name    string
		root    map[string]any
		patch   []types.ConfigPatchOperation
		want    map[string]any
		wantErr bool
	}{
		{
			name: "removes top-level field",
			root: map[string]any{
				"removeMe": true,
				"keepMe":   true,
			},
			patch: []types.ConfigPatchOperation{
				{Op: "remove", Path: "/removeMe"},
			},
			want: map[string]any{
				"keepMe": true,
			},
		},
		{
			name: "removes nested field and preserves siblings",
			root: map[string]any{
				"contactAuth": map[string]any{
					"github": map[string]any{
						"clientId":     "id",
						"clientSecret": "secret",
					},
				},
			},
			patch: []types.ConfigPatchOperation{
				{Op: "remove", Path: "/contactAuth/github/clientSecret"},
			},
			want: map[string]any{
				"contactAuth": map[string]any{
					"github": map[string]any{
						"clientId": "id",
					},
				},
			},
		},
		{
			name: "missing path is no-op",
			root: map[string]any{
				"keepMe": true,
			},
			patch: []types.ConfigPatchOperation{
				{Op: "remove", Path: "/missing/path"},
			},
			want: map[string]any{
				"keepMe": true,
			},
		},
		{
			name: "decodes JSON Pointer escapes",
			root: map[string]any{
				"a/b": map[string]any{
					"c~d":  "value",
					"keep": "value",
				},
			},
			patch: []types.ConfigPatchOperation{
				{Op: "remove", Path: "/a~1b/c~0d"},
			},
			want: map[string]any{
				"a/b": map[string]any{
					"keep": "value",
				},
			},
		},
		{
			name: "rejects unsupported op",
			root: map[string]any{},
			patch: []types.ConfigPatchOperation{
				{Op: "replace", Path: "/a"},
			},
			wantErr: true,
		},
		{
			name: "rejects non-pointer path",
			root: map[string]any{},
			patch: []types.ConfigPatchOperation{
				{Op: "remove", Path: "a.b"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := applyConfigPatch(tt.root, tt.patch)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tt.root, tt.want) {
				t.Errorf("applyConfigPatch() = %v, want %v", tt.root, tt.want)
			}
		})
	}
}
