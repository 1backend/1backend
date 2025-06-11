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
	"log/slog"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/logger"
	downloadtypes "github.com/1backend/1backend/server/internal/services/file/types"
	modeltypes "github.com/1backend/1backend/server/internal/services/model/types"
	"github.com/flusflas/dipper"
	"github.com/pkg/errors"
)

func (ms *ModelService) status(
	modelId string,
) (*modeltypes.ModelStatus, error) {
	hostRsp, _, err := ms.options.ClientFactory.Client(client.WithToken(ms.token)).
		ContainerSvcAPI.GetHost(context.Background()).
		Execute()
	if err != nil {
		return nil, err
	}

	dockerHost := hostRsp.Host

	if ms.options.LLMHost != "" {
		dockerHost = ms.options.LLMHost
	}

	modelAddress := fmt.Sprintf("%v:%v", dockerHost, hostPortNum)

	if modelId == "" {
		rsp, _, err := ms.options.ClientFactory.Client(client.WithToken(ms.token)).
			ConfigSvcAPI.ListConfigs(context.Background()).
			Body(openapi.ConfigSvcListConfigsRequest{
				Slugs: []string{"modelSvc"},
			}).
			Execute()
		if err != nil {
			return nil, err
		}

		modelIdI := dipper.Get(rsp.Configs["modelSvc"].Data, "currentModelId")
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
		rsp, _, err := ms.options.ClientFactory.Client(client.WithToken(ms.token)).
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

	hashRsp, _, err := ms.options.ClientFactory.Client(client.WithToken(ms.token)).
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
