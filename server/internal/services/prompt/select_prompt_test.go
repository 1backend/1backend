/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice_test

import (
	"testing"
	"time"

	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/datastore/localstore"

	promptservice "github.com/openorch/openorch/server/internal/services/prompt"
	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
	"github.com/stretchr/testify/require"
)

func TestSelectPrompt(t *testing.T) {
	fixedTime := time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)
	promptservice.TimeNow = func() time.Time {
		return fixedTime
	}

	tests := []struct {
		name           string
		prompts        []datastore.Row
		expectedPrompt datastore.Row
	}{
		{
			name: "No prompts",

			expectedPrompt: nil,
		},
		{
			name: "Prompt with RunCount 0",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusScheduled,
					RunCount: 0,
				},
			},
			expectedPrompt: &prompttypes.Prompt{
				Status:   prompttypes.PromptStatusScheduled,
				RunCount: 0,
			},
		},
		{
			name: "Prompt not due yet",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusScheduled,
					RunCount: 1,
					LastRun:  fixedTime.Add(-promptservice.BaseDelay / 2),
				},
			},
			expectedPrompt: nil,
		},
		{
			name: "Prompt due",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusScheduled,
					RunCount: 1,
					LastRun:  fixedTime.Add(-promptservice.BaseDelay),
				},
			},
			expectedPrompt: &prompttypes.Prompt{
				Status:   prompttypes.PromptStatusScheduled,
				RunCount: 1,
				LastRun:  fixedTime.Add(-promptservice.BaseDelay),
			},
		},
		{
			name: "Abandoned prompt",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusAbandoned,
					RunCount: 0,
				},
			},
			expectedPrompt: nil,
		},
		{
			name: "Completed prompt",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusCompleted,
					RunCount: 0,
				},
			},
			expectedPrompt: nil,
		},
		{
			name: "Canceled prompt",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusCanceled,
					RunCount: 0,
				},
			},
			expectedPrompt: nil,
		},
		{
			name: "Prompt with RunCount greater than 1, not due yet",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusScheduled,
					RunCount: 2,
					LastRun:  fixedTime.Add(-promptservice.BaseDelay),
				},
			},
			expectedPrompt: nil,
		},
		{
			name: "Prompt with RunCount greater than 1, due",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusScheduled,
					RunCount: 2,
					LastRun:  fixedTime.Add(-promptservice.BaseDelay * 2),
				},
			},
			expectedPrompt: &prompttypes.Prompt{
				Status:   prompttypes.PromptStatusScheduled,
				RunCount: 2,
				LastRun:  fixedTime.Add(-promptservice.BaseDelay * 2),
			},
		},
		{
			name: "Prompt with RunCount greater than 1, off by one",
			prompts: []datastore.Row{
				&prompttypes.Prompt{
					Status:   prompttypes.PromptStatusScheduled,
					RunCount: 2,
					LastRun: fixedTime.Add(
						-promptservice.BaseDelay*2 + time.Second,
					),
				},
			},
			expectedPrompt: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memStore, err := localstore.NewLocalStore(&prompttypes.Prompt{}, "")
			require.NoError(t, err)

			err = memStore.UpsertMany(tt.prompts)
			require.NoError(t, err)
			actualPrompt, err := promptservice.SelectPrompt(memStore)
			require.NoError(t, err)

			if tt.expectedPrompt == nil {
				require.Nil(t, actualPrompt)
			} else {
				require.NotNil(t, actualPrompt)
				require.Equal(t, tt.expectedPrompt, actualPrompt)
			}
		})
	}
}
