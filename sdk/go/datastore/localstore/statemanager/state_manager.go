/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package statemanager

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"os"
	"os/signal"
	"path"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/pkg/errors"
)

const periodicSaveAlertThreshold = 100 * 1024 // 10 KB

type StateFile struct {
	Rows []any `json:"rows"`
}

type StateManager struct {
	instance     any
	lock         sync.Mutex
	filePath     string
	hasChanged   bool
	stateGetter  func() []any
	lastChecksum [32]byte
}

func New(instance any, stateGetter func() []any, filePath string) *StateManager {
	sm := &StateManager{
		filePath:    filePath + ".zip",
		stateGetter: stateGetter,
		instance:    instance,
	}
	sm.setupSignalHandler()
	return sm
}

func (sm *StateManager) LoadState() (any, error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	_, err := os.Stat(sm.filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(sm.filePath), 0755)
		if err != nil {
			return nil, err
		}

		emptyData := []byte("{}")
		zippedEmptyData, err := zipData(sm.filePath, emptyData)
		if err != nil {
			return nil, errors.Wrap(err, "failed to zip data")
		}

		err = ioutil.WriteFile(sm.filePath, zippedEmptyData, 0644)
		if err != nil {
			return nil, err
		}

		//
		// Some hacks below to prevent sync issues with the file system
		// (fixes issue: Config service start failed: failed to create zip reader: zip: not a valid zip file)
		//

		// Open the file and sync it to disk (flush it)
		file, err := os.OpenFile(sm.filePath, os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Ensure the data is physically written to disk
		err = file.Sync()
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(sm.filePath)
	if err != nil {
		return nil, err
	}

	unzippedData, err := unzipData(data)
	if err != nil {
		return nil, err
	}

	elemType := reflect.TypeOf(sm.instance)
	slice, err := unmarshalIntoSlice([]byte(unzippedData), elemType)

	return slice, err
}

func unmarshalIntoSlice(unzippedData []byte, elemType reflect.Type) (interface{}, error) {
	var stateFile StateFile

	if strings.TrimSpace(string(unzippedData)) == "" {
		return nil, nil
	}

	err := json.Unmarshal(unzippedData, &stateFile)
	if err != nil {
		return nil, err
	}

	slice := createSliceOfType(elemType)
	sliceValue := reflect.ValueOf(slice)

	for _, row := range stateFile.Rows {
		rowData, err := json.Marshal(row)
		if err != nil {
			return nil, err
		}

		elem := reflect.New(elemType).Interface()
		err = json.Unmarshal(rowData, elem)
		if err != nil {
			return nil, err
		}

		sliceValue = reflect.Append(sliceValue, reflect.ValueOf(elem).Elem())
	}

	return sliceValue.Interface(), nil
}

func createSliceOfType(elemType reflect.Type) interface{} {
	sliceType := reflect.SliceOf(elemType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	return sliceValue.Interface()
}

func (sm *StateManager) SaveState(shallowCopy []any) (int, error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	data, err := json.Marshal(&StateFile{
		Rows: shallowCopy,
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to marshal data when saving state")
	}

	newChecksum := sha256.Sum256(data)
	if newChecksum == sm.lastChecksum {
		return len(shallowCopy), nil // no change, skip write
	}

	zippedData, err := zipData(sm.filePath, data)
	if err != nil {
		return 0, errors.Wrap(err, "failed to zip data when saving state")
	}

	tempFilePath := sm.filePath + ".tmp"
	err = ioutil.WriteFile(tempFilePath, zippedData, 0644)
	if err != nil {
		return 0, errors.Wrap(err, "failed to write temp file when saving state")
	}

	finalFilePath := sm.filePath
	err = os.Rename(tempFilePath, finalFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "failed to rename temp file when saving state")
	}

	sm.lastChecksum = newChecksum
	sm.hasChanged = false
	return len(shallowCopy), nil
}

func (sm *StateManager) MarkChanged() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.hasChanged = true
}

func (sm *StateManager) PeriodicSaveState(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if sm.hasChanged {
				s, err := sm.SaveState(sm.stateGetter())
				if err != nil {
					logger.Error("Error saving file state",
						slog.String("filePath", sm.filePath),
						slog.String("error", err.Error()),
					)
				}

				if s > periodicSaveAlertThreshold {
					// Warn if the periodic save file is unusually large.
					// Prevents users from accidentally writing huge files and wearing out SSDs.
					logger.Warn("Periodic save file is too large",
						slog.String("path", sm.filePath),
						slog.String("interval", interval.String()),
						slog.String("size", humanize.Bytes(uint64(s))),
					)
				}
			}
		}
	}
}

func (sm *StateManager) setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		_, err := sm.SaveState(sm.stateGetter())
		if err != nil {
			logger.Error("Error saving file state on shutdown",
				slog.String("filePath", sm.filePath),
				slog.String("error", err.Error()),
			)
		}
		os.Exit(0)
	}()
}

func (sm *StateManager) Close() error {
	_, err := sm.SaveState(sm.stateGetter())
	return err
}

func (sm *StateManager) Refresh() error {
	_, err := sm.SaveState(sm.stateGetter())
	return err
}

func zipData(filePath string, data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	writer, err := zipWriter.Create(
		fmt.Sprintf("%v.json", path.Base(filePath)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create zip writer")
	}

	_, err = writer.Write(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write zip data")
	}

	err = zipWriter.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close zip writer")
	}

	return buf.Bytes(), nil
}

func unzipData(data []byte) ([]byte, error) {
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create zip reader")
	}
	if len(reader.File) != 1 {
		return nil, errors.Wrap(err, "expected exactly one file in zip archive")
	}

	file := reader.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file in zip archive")
	}
	defer rc.Close()

	unzippedData, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file in zip archive")
	}

	return unzippedData, nil
}
