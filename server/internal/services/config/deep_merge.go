package configservice

import (
	"fmt"
	"strings"

	types "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/pkg/errors"
)

var errInvalidConfigPatch = errors.New("invalid config patch")

// DeepMerge recursively merges src into dst.
// - Nested maps are merged deeply.
// - Arrays and non-map types are replaced.
// - Type mismatches result in overwrites.
func DeepMerge(dst, src map[string]interface{}) map[string]interface{} {
	for key, srcVal := range src {

		if dstVal, ok := dst[key]; ok {
			dstMap, dstIsMap := dstVal.(map[string]interface{})
			srcMap, srcIsMap := srcVal.(map[string]interface{})

			if dstIsMap && srcIsMap {
				dst[key] = DeepMerge(dstMap, srcMap)
				continue
			}
		}

		dst[key] = srcVal
	}
	return dst
}

// ExpandDotPath turns a dot path like "a.b.c" into {"a":{"b":{"c":val}}}.
func ExpandDotPath(path string, val any) map[string]any {
	parts := strings.Split(path, ".")
	root := map[string]any{}
	cur := root
	for i, p := range parts {
		if i == len(parts)-1 {
			cur[p] = val
		} else {
			next := map[string]any{}
			cur[p] = next
			cur = next
		}
	}
	return root
}

func applyConfigPatch(root map[string]any, patch []types.ConfigPatchOperation) error {
	for i, op := range patch {
		switch op.Op {
		case "remove":
			if err := removeJSONPointer(root, op.Path); err != nil {
				return fmt.Errorf("%w at index %d: %v", errInvalidConfigPatch, i, err)
			}
		default:
			return fmt.Errorf("%w at index %d: unsupported op %q", errInvalidConfigPatch, i, op.Op)
		}
	}

	return nil
}

func removeJSONPointer(root map[string]any, path string) error {
	parts, err := jsonPointerParts(path)
	if err != nil {
		return err
	}
	if len(parts) == 0 {
		return errors.New("removing the document root is not supported")
	}

	cur := root
	for _, part := range parts[:len(parts)-1] {
		next, ok := cur[part].(map[string]any)
		if !ok {
			return nil
		}
		cur = next
	}

	delete(cur, parts[len(parts)-1])
	return nil
}

func jsonPointerParts(path string) ([]string, error) {
	if path == "" {
		return nil, nil
	}
	if !strings.HasPrefix(path, "/") {
		return nil, errors.New("path must be empty or start with /")
	}

	rawParts := strings.Split(path[1:], "/")
	parts := make([]string, 0, len(rawParts))
	for _, part := range rawParts {
		decoded, err := decodeJSONPointerPart(part)
		if err != nil {
			return nil, err
		}
		parts = append(parts, decoded)
	}

	return parts, nil
}

func decodeJSONPointerPart(part string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(part); i++ {
		if part[i] != '~' {
			b.WriteByte(part[i])
			continue
		}

		if i+1 >= len(part) {
			return "", errors.New("invalid ~ escape in path")
		}

		switch part[i+1] {
		case '0':
			b.WriteByte('~')
		case '1':
			b.WriteByte('/')
		default:
			return "", errors.New("invalid ~ escape in path")
		}
		i++
	}

	return b.String(), nil
}
