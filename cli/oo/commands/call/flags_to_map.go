package call

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseFlagToMap(payload map[string]interface{}, flag string) error {
	if !strings.HasPrefix(flag, "--") {
		return fmt.Errorf("invalid flag format: %s", flag)
	}
	flag = strings.TrimPrefix(flag, "--")

	parts := strings.SplitN(flag, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid flag format: %s", flag)
	}
	key := parts[0]
	valueStr := parts[1]

	value := parseValue(valueStr)

	keyParts := splitKey(key)
	if len(keyParts) == 0 {
		return errors.New("invalid key format")
	}

	if len(keyParts) > 1 &&
		strings.HasSuffix(keyParts[len(keyParts)-2], "-id") {
		parentKey := keyParts[len(keyParts)-2][:len(keyParts[len(keyParts)-2])-3]
		idKey := keyParts[len(keyParts)-1]

		var parentArray []map[string]interface{}
		if existingValue, exists := payload[parentKey]; exists {
			if arr, ok := existingValue.([]map[string]interface{}); ok {
				parentArray = arr
			}
		}

		parentArray = append(parentArray, map[string]interface{}{idKey: value})
		payload[parentKey] = parentArray

		return nil
	}

	currentMap := payload
	for i := 0; i < len(keyParts)-1; i++ {
		part := keyParts[i]
		if existingValue, exists := currentMap[part]; exists {
			if nestedMap, ok := existingValue.(map[string]interface{}); ok {
				currentMap = nestedMap
			} else {
				newMap := make(map[string]interface{})
				currentMap[part] = newMap
				currentMap = newMap
			}
		} else {
			newMap := make(map[string]interface{})
			currentMap[part] = newMap
			currentMap = newMap
		}
	}

	finalKey := keyParts[len(keyParts)-1]

	if len(keyParts) == 2 && finalKey == "id" && !strings.Contains(key, ".") {
		parentKey := keyParts[0]
		if existingParent, exists := payload[parentKey]; exists {
			if parentMap, ok := existingParent.(map[string]interface{}); ok {
				if existingID, exists := parentMap[finalKey]; exists && len(parentMap) == 1 {
					payload[parentKey] = []map[string]interface{}{
						{finalKey: existingID},
						{finalKey: value},
					}
					return nil
				}
			}
			if parentArray, ok := existingParent.([]map[string]interface{}); ok {
				payload[parentKey] = append(parentArray, map[string]interface{}{finalKey: value})
				return nil
			}
		}
	}

	if existingValue, exists := currentMap[finalKey]; exists {
		switch existingValue := existingValue.(type) {
		case []interface{}:

			currentMap[finalKey] = append(existingValue, value)
		default:

			currentMap[finalKey] = []interface{}{existingValue, value}
		}
	} else {
		currentMap[finalKey] = value
	}

	return nil
}

func splitFlags(flag string) []string {
	return strings.Fields(flag)
}

func splitKey(key string) []string {
	if strings.Contains(key, ".") {
		return strings.Split(key, ".")
	}
	return strings.Split(key, "-")
}

func parseValue(value string) interface{} {
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}

	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue
	}

	return value
}
