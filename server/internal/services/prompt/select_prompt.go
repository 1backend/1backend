/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice

import (
	"math"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"

	prompttypes "github.com/1backend/1backend/server/internal/services/prompt/types"
)

func SelectPrompt(promptsMem datastore.DataStore) (*prompttypes.Prompt, error) {
	promptIs, err := promptsMem.Query(
		datastore.IsInList(
			datastore.Field("status"),
			prompttypes.PromptStatusScheduled,
			prompttypes.PromptStatusErrored,
		),
	).
		OrderBy(datastore.OrderByField("createdAt", false)).
		Find()
	if err != nil {
		return nil, err
	}

	prompt, _, _ := selectPromptFromRows(promptIs)
	return prompt, nil
}

func selectPromptFromRows(promptIs []datastore.Row) (*prompttypes.Prompt, bool, time.Duration) {
	hasQueuedPrompt := false
	nextDue := time.Duration(0)
	now := TimeNow()

	for _, promptI := range promptIs {
		prompt := promptI.(*prompttypes.Prompt)

		if prompt.Status != prompttypes.PromptStatusScheduled &&
			prompt.Status != prompttypes.PromptStatusErrored {
			continue
		}
		hasQueuedPrompt = true

		backoff := promptBackoff(prompt.RunCount)
		wait := backoff - now.Sub(prompt.LastRun)
		if prompt.RunCount == 0 || wait <= 0 {
			return prompt, hasQueuedPrompt, 0
		}
		if nextDue == 0 || wait < nextDue {
			nextDue = wait
		}
	}

	return nil, hasQueuedPrompt, nextDue
}

func promptBackoff(runCount int) time.Duration {
	if runCount == 0 {
		// otherwise backoff is 0s
		runCount = 1
	}
	cappedRunCount := math.Min(float64(runCount), 10)
	return BaseDelay * time.Duration(math.Pow(2, cappedRunCount-1))
}
