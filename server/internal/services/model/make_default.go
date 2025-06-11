/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package modelservice

import (
	"context"
	"fmt"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
)

func (ms *ModelService) makeDefault(ctx context.Context, modelId string) error {
	stat, err := ms.status(modelId)
	if err != nil {
		return err
	}
	if !stat.AssetsReady {
		return fmt.Errorf("cannot set model as it is not downloaded yet")
	}

	rsp, _, err := ms.options.ClientFactory.Client(client.WithToken(ms.token)).
		ConfigSvcAPI.ReadConfig(ctx).
		Execute()

	if err != nil {
		return err
	}

	_, ok := rsp.Config.Data["modelSvc"].(map[string]any)
	if !ok {
		rsp.Config.Data["modelSvc"] = map[string]any{}
	}

	m := rsp.Config.Data["modelSvc"].(map[string]any)

	m["currentModelId"] = modelId

	_, _, err = ms.options.ClientFactory.Client(client.WithToken(ms.token)).
		ConfigSvcAPI.SaveConfig(ctx).
		Body(openapi.ConfigSvcSaveConfigRequest{
			Data: m,
		}).
		Execute()

	return err
}
