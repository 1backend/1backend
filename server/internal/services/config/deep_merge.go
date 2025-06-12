package configservice

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
