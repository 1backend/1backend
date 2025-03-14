/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice

import (
	"math"
	"time"

	"github.com/openorch/openorch/sdk/go/datastore"

	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

func SelectPrompt(promptsMem datastore.DataStore) (*prompttypes.Prompt, error) {
	promptIs, err := promptsMem.Query().
		OrderBy(datastore.OrderByField("createdAt", false)).
		Find()
	if err != nil {
		return nil, err
	}

	for _, promptI := range promptIs {
		prompt := promptI.(*prompttypes.Prompt)

		if prompt.Status == prompttypes.PromptStatusAbandoned ||
			prompt.Status == prompttypes.PromptStatusCompleted ||
			prompt.Status == prompttypes.PromptStatusCanceled {
			continue
		}

		runCount := prompt.RunCount
		if prompt.RunCount == 0 {
			// otherwise backoff is 0s
			runCount = 1
		}
		cappedRunCount := math.Min(float64(runCount), 10)
		backoff := BaseDelay * time.Duration(math.Pow(2, cappedRunCount-1))

		if prompt.RunCount == 0 ||
			TimeNow().Sub(prompt.LastRun) >= backoff {
			return prompt, nil
		}
	}

	return nil, nil
}
