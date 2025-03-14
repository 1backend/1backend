/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dapper

import (
	"fmt"
	"os"
	"path"
	"runtime"

	features "github.com/openorch/openorch/dapper/features"
	dt "github.com/openorch/openorch/dapper/types"
)

// ConfigurationManager manages configurations and feature dependencies.
type ConfigurationManager struct {
	stream                 bool
	Features               map[string]dt.Feature
	CurrentPlatform        dt.Platform
	CacheFolder            string
	Printf                 func(format string, a ...any) (n int, err error)
	PreCheckCallback       func(*dt.Feature, *dt.RunContext)
	PostCheckCallback      func(*dt.Feature, *dt.RunContext)
	PreExeccutionCallback  func(*dt.Feature, *dt.RunContext)
	PostExecutionCallback  func(*dt.Feature, *dt.RunContext)
	PreAgainCheckCallback  func(*dt.Feature, *dt.RunContext)
	PostAgainCheckCallback func(*dt.Feature, *dt.RunContext)
}

func NewConfigurationManagerFromSource() *ConfigurationManager {
	fs := make(map[string]dt.Feature)
	for _, v := range features.Features {
		fs[v.ID] = v
	}

	return NewConfigurationManager(fs)
}

// NewConfigurationManager initializes a new configuration manager with predefined features.
func NewConfigurationManager(fs map[string]dt.Feature) *ConfigurationManager {
	cm := &ConfigurationManager{
		Features: fs,
		Printf:   fmt.Printf,
		stream:   true,
	}

	// check os
	switch runtime.GOOS {
	case "windows":
		cm.CurrentPlatform = dt.Windows
	case "linux":
		cm.CurrentPlatform = dt.Linux
	case "darwin":
		cm.CurrentPlatform = dt.MacOS
	default:
		//	cm.CurrentPlatform = dt.Unknown
	}

	home, err := os.UserHomeDir()
	if err == nil {
		cm.CacheFolder = path.Join(home, ".dapper")
		// create the cache folder if doesn't exist
		if _, err := os.Stat(cm.CacheFolder); os.IsNotExist(err) {
			os.Mkdir(cm.CacheFolder, 0755)
		}
	}

	return cm
}
