/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dappertypes

import "time"

// Feature, FeatureDefinition, or Dependency, or Resource, take your pick
type Feature struct {
	ID              string                `json:"id"`
	Name            string                `json:"name"`
	PlatformScripts map[Platform]*Scripts `json:"platformScript"`
	/*
		Feature can be either an string (`Feature` ID) or a `FeatureInvocation` object.
	*/
	Features []any `json:"features"`
	/*
		Same as `Features` but platform specific.
		Can be either an string (`Feature` ID) or a `FeatureInvocation`` object
	*/
	PlatformFeatures map[Platform][]any `json:"platformFeatures"`
	/*
		Is this something that can easily change, like a container running etc?
		If so, mark it as Temporary.
	*/
	Temporary bool       `json:"temporary"`
	Arguments []Argument `json:"arguments"`
}

type Argument struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        ArgumentType `json:"type"`
	Default     any          `json:"default"`
	Required    bool         `json:"required"`

	// Not used, ignore
	IsLast bool
	Index  int
}

type ArgumentType string

const (
	String ArgumentType = "string"
	Int    ArgumentType = "int"
	Bool   ArgumentType = "bool"
)

type App struct {
	Name string `json:"name"`
	/*
		Feature can be either an string (Feature ID) or a FeatureInvocation object
	*/
	Features         []interface{}              `json:"features"`
	PlatformFeatures map[Platform][]interface{} `json:"platformFeatures"`
}

type FeatureInvocation struct {
	FeatureID string `json:"featureId"`
	Args      []any  `json:"args"`
}

type Scripts struct {
	Check   *Script `json:"check"`
	Execute *Script `json:"execute"`
	/*
	 Uninstall scripts are optional, they are only used for testing purposes
	*/
	Uninstall *Script `json:"uninstall"`
}

type Script struct {
	Source   string  `json:"source"`
	Runtime  Runtime `json:"runtime"`
	Sudo     bool    `json:"sudo"`
	Reboot   bool    `json:"reboot"`
	Detach   bool    `json:"detach"`
	Cachable bool    `json:"cacheable"`
}

type Platform string

const (
	Windows Platform = "windows"
	Linux   Platform = "linux"
	MacOS   Platform = "macos"
)

type Runtime string

const (
	Powershell Runtime = "powershell"
	Bash       Runtime = "bash"
	Cmd        Runtime = "cmd"
)

type FeatureCheckResult struct {
	ID        string
	CheckedAt time.Time
	Success   bool
	Error     error
	Stdout    string
	Stderr    string
	Duration  time.Duration
}

type FeatureExecuteResult struct {
	ID         string
	ExecutedAt time.Time
	Success    bool
	Error      error
	Stdout     string
	Stderr     string
	Duration   time.Duration
}

type FeatureProcessed struct {
	Feature         *Feature
	Level           int
	CheckResult     *FeatureCheckResult
	CheckError      error
	ExecutionResult *FeatureExecuteResult
	ExecutionError  error
}

type RunContext struct {
	FeaturesProcessed []*FeatureProcessed
	Level             int
	// Passed down from the CLI like dapper run -var-username=joe -var-userage=25 myapp.json
	TopLevelVariables map[string]string
	// Were any Executes ran that require a reboot
	RebootRequired bool
	CommandCounter int
	Anon           bool
}
