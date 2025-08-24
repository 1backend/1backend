package configservice

import "strings"

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
