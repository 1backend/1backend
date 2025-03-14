/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package modelservice

import (
	"context"
	"fmt"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
)

func (ms *ModelService) makeDefault(ctx context.Context, modelId string) error {
	stat, err := ms.status(modelId)
	if err != nil {
		return err
	}
	if !stat.AssetsReady {
		return fmt.Errorf("cannot set model as it is not downloaded yet")
	}

	rsp, _, err := ms.clientFactory.Client(sdk.WithToken(ms.token)).
		ConfigSvcAPI.GetConfig(ctx).
		Execute()

	if err != nil {
		return err
	}

	_, ok := rsp.Config.Data["model-svc"].(map[string]any)
	if !ok {
		rsp.Config.Data["model-svc"] = map[string]any{}
	}

	m := rsp.Config.Data["model-svc"].(map[string]any)

	m["currentModelId"] = modelId

	_, _, err = ms.clientFactory.Client(sdk.WithToken(ms.token)).
		ConfigSvcAPI.SaveConfig(ctx).
		Body(openapi.ConfigSvcSaveConfigRequest{
			Config: rsp.Config,
		}).
		Execute()

	return err
}
