package imageservice

import "strings"

const imageStorageSourceHeader = "X-1Backend-Storage-Source"

func imageCacheLabelFromStorageSource(storageSource string) string {
	switch strings.ToLower(strings.TrimSpace(storageSource)) {
	case "gcs":
		return "gcs"
	default:
		return "cold_miss"
	}
}
