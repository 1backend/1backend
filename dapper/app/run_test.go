//go:build linux
// +build linux

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dapper

import (
	"testing"

	dt "github.com/openorch/openorch/dapper/types"
	"github.com/stretchr/testify/assert"
)

func newCrossPlatformFeature(id, name string) dt.Feature {
	return dt.Feature{
		ID:   id,
		Name: name,
		PlatformScripts: map[dt.Platform]*dt.Scripts{
			dt.Linux: {
				Execute: &dt.Script{
					Source:  "true",
					Runtime: "bash",
				},
				Check: &dt.Script{
					Source:  "true",
					Runtime: "bash",
				},
			},
		},
	}
}

func abcSet() []dt.Feature {
	a := newCrossPlatformFeature("test-a", "Test A")
	b := newCrossPlatformFeature("test-b", "Test B")
	c := newCrossPlatformFeature("test-c", "Test C")

	a.Features = []any{
		"test-b",
	}
	b.Features = []any{
		"test-c",
	}

	return []dt.Feature{
		a, b, c,
	}
}

func TestBasics(t *testing.T) {
	t.Run("'A' check succeeds -> no execute or again check", func(t *testing.T) {
		set := abcSet()
		setIndex := map[string]dt.Feature{}
		for _, v := range set {
			setIndex[v.ID] = v
		}
		cm := NewConfigurationManager(setIndex)
		ct, err := cm.Run(&dt.App{
			Features: []any{
				"test-a",
			},
		}, nil, false)
		assert.NoError(t, err)
		assert.Equals(t, len(ct.FeaturesProcessed), 1)
		// execute does not run because check succeeds
		assert.Equals(t, ct.FeaturesProcessed[0].ExecutionResult == nil, true)
	})

	t.Run("'A' check fails -> 'B' check passes -> 'A' execute succeeds -> 'A' check succeeds", func(t *testing.T) {
		set := abcSet()
		set[0].PlatformScripts[dt.Linux].Check.Source = "false"
		setIndex := map[string]dt.Feature{}
		for _, v := range set {
			setIndex[v.ID] = v
		}
		cm := NewConfigurationManager(setIndex)
		cm.PostCheckCallback = func(feature *dt.Feature, universe *dt.RunContext) {
			if feature.ID == "test-a" {
				set[0].PlatformScripts[dt.Linux].Check.Source = "true"
			}
		}
		ct, err := cm.Run(&dt.App{
			Features: []any{
				"test-a",
			},
		}, nil, false)
		assert.NoError(t, err)
		assert.Equals(t, 2, len(ct.FeaturesProcessed))
		// execute runs because first check fails
		assert.Equals(t, ct.FeaturesProcessed[0].ExecutionResult != nil, true)
	})

	t.Run("'A' check fails -> 'B' check passes -> 'A' execute succeeds -> 'A' check fails", func(t *testing.T) {
		set := abcSet()
		set[0].PlatformScripts[dt.Linux].Check.Source = "false"
		setIndex := map[string]dt.Feature{}
		for _, v := range set {
			setIndex[v.ID] = v
		}
		cm := NewConfigurationManager(setIndex)
		ct, err := cm.Run(&dt.App{
			Features: []any{
				"test-a",
			},
		}, nil, false)
		assert.Error(t, err)
		assert.Equals(t, 2, len(ct.FeaturesProcessed))
		// execute runs because first check fails
		assert.Equals(t, ct.FeaturesProcessed[0].ExecutionResult != nil, true)
		assert.Equals(t, ct.FeaturesProcessed[1].ExecutionResult == nil, true)
	})
}
