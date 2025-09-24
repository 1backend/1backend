package util

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// CollectFromPath collects entities of type T from a file or directory path.
// The entityName is used to provide context in error messages.
func CollectFromPath[T any](path string, entityName string) ([]T, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, errors.Wrap(err, fmt.Sprintf("%s path not found: '%v'", entityName, path))
	} else if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error checking %s path", entityName))
	}

	var entities []T

	if stat.IsDir() {
		err = filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error accessing file '%v'", filePath))
			}
			if d.IsDir() {
				return nil
			}

			fileEntities, err := extractEntitiesFromFile[T](filePath, entityName)
			if err != nil {
				return err
			}
			entities = append(entities, fileEntities...)
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		fileEntities, err := extractEntitiesFromFile[T](path, entityName)
		if err != nil {
			return nil, err
		}
		entities = append(entities, fileEntities...)
	}

	return entities, nil
}

func extractEntitiesFromFile[T any](filePath string, entityName string) ([]T, error) {
	var items []T
	if err := ExtractFromFile(filePath, &items); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to parse %s file at '%v'", entityName, filePath))
	}

	return items, nil
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

func ErrorWithBody(err error, resp *http.Response, msg string) error {
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return errors.Wrap(readErr, "failed to read response body")
	}

	return errors.Wrap(err, fmt.Sprintf("%s: '%s'", msg, string(body)))
}
