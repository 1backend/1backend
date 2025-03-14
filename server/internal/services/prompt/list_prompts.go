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
	"encoding/json"

	"github.com/openorch/openorch/sdk/go/datastore"
	prompttypes "github.com/openorch/openorch/server/internal/services/prompt/types"
)

func (p *PromptService) listPrompts(
	options *prompttypes.ListPromptOptions,
) ([]*prompttypes.Prompt, int64, error) {
	q := p.promptsStore.Query(
		options.Query.Filters...,
	).Limit(options.Query.Limit)

	if len(options.Query.OrderBys) > 0 {
		q = q.OrderBy(options.Query.OrderBys...)
	} else {
		q = q.OrderBy(datastore.OrderByField("createdAt", false))
	}

	if options.Query.JSONAfter != "" {
		v := []any{}
		err := json.Unmarshal([]byte(options.Query.JSONAfter), &v)
		if err != nil {
			return nil, 0, err
		}
		q = q.After(v...)
	}

	resI, err := q.Find()
	if err != nil {
		return nil, 0, err
	}

	var count int64
	if options.Query.Count {
		var err error
		count, err = q.Count()
		if err != nil {
			return nil, 0, err
		}
	}

	res := []*prompttypes.Prompt{}
	for _, v := range resI {
		res = append(res, v.(*prompttypes.Prompt))
	}

	return res, count, nil
}
