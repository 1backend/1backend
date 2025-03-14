package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

func PrettyJSON(rawJSON []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, rawJSON, "", "  ")
	if err != nil {
		return "", nil
	}

	return prettyJSON.String(), nil
}

// ExtractFromFile extracts one or more entities from a YAML file into a slice.
func ExtractFromFile(filePath string, entitySlice any) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read file at '%v'", filePath))
	}

	entityVal := reflect.ValueOf(entitySlice)
	if entityVal.Kind() != reflect.Ptr || entityVal.Elem().Kind() != reflect.Slice {
		return errors.New("entitySlice must be a pointer to a slice")
	}

	elemType := entityVal.Elem().Type().Elem()

	if err := yaml.Unmarshal(data, entitySlice); err == nil {
		return nil
	}

	singleEntity := reflect.New(elemType).Interface()
	if err := yaml.Unmarshal(data, singleEntity); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to parse YAML at '%v'", filePath))
	}

	entitySliceValue := entityVal.Elem()
	entitySliceValue.Set(reflect.Append(entitySliceValue, reflect.ValueOf(singleEntity).Elem()))

	return nil
}
