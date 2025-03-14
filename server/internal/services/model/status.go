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
	"log/slog"

	"github.com/flusflas/dipper"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"
	downloadtypes "github.com/openorch/openorch/server/internal/services/file/types"
	modeltypes "github.com/openorch/openorch/server/internal/services/model/types"
	"github.com/pkg/errors"
)

func (ms *ModelService) status(
	modelId string,
) (*modeltypes.ModelStatus, error) {
	hostRsp, _, err := ms.clientFactory.Client(sdk.WithToken(ms.token)).
		ContainerSvcAPI.GetHost(context.Background()).
		Execute()
	if err != nil {
		return nil, err
	}

	dockerHost := hostRsp.Host
	openorchLLMHost := ms.llmHost
	if openorchLLMHost != "" {
		dockerHost = openorchLLMHost
	}

	modelAddress := fmt.Sprintf("%v:%v", dockerHost, hostPortNum)

	if modelId == "" {
		rsp, _, err := ms.clientFactory.Client(sdk.WithToken(ms.token)).
			ConfigSvcAPI.GetConfig(context.Background()).
			Execute()
		if err != nil {
			return nil, err
		}

		modelIdI := dipper.Get(rsp.Config.Data, "model-svc.currentModelId")
		var ok bool
		modelId, ok = modelIdI.(string)
		if !ok {
			modelId = DefaultModelId
		}
	}

	modelI, found, err := ms.modelsStore.Query(
		datastore.Id(modelId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("model not found")
	}
	model := modelI.(*modeltypes.Model)

	for _, asset := range model.Assets {
		rsp, _, err := ms.clientFactory.Client(sdk.WithToken(ms.token)).
			FileSvcAPI.GetDownload(context.Background(), asset.Url).
			Execute()
		if err != nil {
			return nil, err
		}
		if !rsp.Exists ||
			*rsp.Download.Status != string(
				downloadtypes.DownloadStatusCompleted,
			) {
			return &modeltypes.ModelStatus{
				AssetsReady: false,
				Address:     modelAddress,
			}, nil
		}
	}

	platformI, found, err := ms.platformsStore.Query(
		datastore.Id(model.PlatformId),
	).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("cannot find platform")
	}
	platform := platformI.(*modeltypes.Platform)

	hash, err := modelToHash(model, platform)
	if err != nil {
		return nil, err
	}

	isRunning := false

	hashRsp, _, err := ms.clientFactory.Client(sdk.WithToken(ms.token)).
		ContainerSvcAPI.ContainerIsRunning(context.Background()).
		Hash(hash).
		Execute()
	if err != nil {
		logger.Warn("Checking if running error",
			slog.String("error", err.Error()),
		)
	}
	if err == nil && hashRsp != nil && hashRsp.IsRunning {
		isRunning = true
	}

	// @todo lock this
	if v, ok := ms.modelPortMap[hostPortNum]; ok {
		if !v.Answering {
			isRunning = false
		}
	}

	return &modeltypes.ModelStatus{
		Running:     isRunning,
		AssetsReady: true,
		Address:     modelAddress,
	}, nil
}
