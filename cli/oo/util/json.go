package util

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(rawJSON []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, rawJSON, "", "  ")
	if err != nil {
		return "", nil
	}

	return prettyJSON.String(), nil
}
