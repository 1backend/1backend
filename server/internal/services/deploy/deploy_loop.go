/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package deployservice

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/logger"
	"github.com/openorch/openorch/server/internal/services/deploy/allocator"
	deploy "github.com/openorch/openorch/server/internal/services/deploy/types"
	"github.com/pkg/errors"
)

// image and container name prefix
const containerPrefix = "sup-"

func (ns *DeployService) loop(triggerOnly bool) {
	interval := 15 * time.Second
	if triggerOnly {
		interval = 100 * 365 * 24 * time.Hour
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	if !ns.triggerOnly {
		go func() {
			ns.triggerChan <- struct{}{}
		}()
	}

	for {
		select {
		case <-ticker.C:
			ns.runCycle()

		case <-ns.triggerChan:
			ns.runCycle()

		case <-sigChan:
			return
		}
	}
}

func (ns *DeployService) runCycle() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			logger.Error("Deploy cycle panic", slog.Any("panic", r))
		}
	}()

	err := ns.cycle()
	if err != nil {
		logger.Error("Deploy cycle error", slog.Any("error", err))
	}
}

func (ns *DeployService) cycle() error {
	ctx := context.Background()

	ns.lock.Acquire(ctx, "deploy-svc-deploy")
	defer ns.lock.Release(ctx, "deploy-svc-deploy")

	registry := ns.clientFactory.Client(sdk.WithToken(ns.token)).RegistrySvcAPI

	deploymentIs, err := ns.deploymentStore.Query().Find()
	if err != nil {
		return err
	}

	deployments := []*deploy.Deployment{}

	for _, deploymentI := range deploymentIs {
		deployment := deploymentI.(*deploy.Deployment)
		deployments = append(deployments, deployment)
	}

	listNodesRsp, _, err := registry.ListNodes(ctx).Execute()
	if err != nil {
		return errors.Wrap(err, "Error calling list nodes")
	}

	listInstancesRsp, _, err := registry.ListInstances(ctx).Execute()
	if err != nil {
		return errors.Wrap(err, "Error calling list service instances")
	}

	listDefinitionsRsp, _, err := registry.ListDefinitions(ctx).Execute()
	if err != nil {
		return errors.Wrap(err, "Error calling list service definitions")
	}

	commands := allocator.GenerateCommands(
		listNodesRsp.Nodes,
		listInstancesRsp.Instances,
		deployments,
	)

	for _, command := range commands {
		var node *openapi.RegistrySvcNode
		var definition *openapi.RegistrySvcDefinition
		var deployment *deploy.Deployment

		for _, v := range listNodesRsp.Nodes {
			if v.Url == command.NodeUrl {
				node = &v
			}
		}

		for _, v := range deployments {
			if command.DeploymentId != nil && v.Id == *command.DeploymentId {
				deployment = v
			}
		}

		if deployment == nil {
			deploymentId := ""
			if command.DeploymentId != nil {
				deploymentId = *command.DeploymentId
			}
			logger.Error("No deployment",
				slog.String("deploymentId", deploymentId),
				slog.String("commandAction", string(command.Action)),
			)
			continue
		}

		for _, v := range listDefinitionsRsp.Definitions {
			if v.Id == deployment.DefinitionId {
				definition = &v
			}
		}

		err := ns.processCommand(ctx, command, node, definition, deployment)
		if err != nil {
			logger.Error(
				"Error processing deploy command",
				slog.Any("error", err),
			)
		}
	}

	return nil
}

func (ns *DeployService) processCommand(
	ctx context.Context,
	command *deploy.Command,
	node *openapi.RegistrySvcNode,
	definition *openapi.RegistrySvcDefinition,
	deployment *deploy.Deployment,
) error {
	deployment, err := ns.getDeploymentById(*command.DeploymentId)
	if err != nil {
		return err
	}

	if deployment.Status == deploy.DeploymentStatusPending ||
		deployment.Status == deploy.DeploymentStatusError {
		deployment.Status = deploy.DeploymentStatusDeploying
		err = ns.deploymentStore.Upsert(deployment)
		if err != nil {
			return err
		}
	}

	switch command.Action {
	case deploy.CommandTypeStart:
		ns.executeStartCommand(ctx, command, node, definition, deployment)
	case deploy.CommandTypeScale:
	case deploy.CommandTypeKill:
		ns.executeKillCommand(ctx, command, node, definition, deployment)
	}

	return nil
}

func (ns *DeployService) executeKillCommand(
	ctx context.Context,
	command *deploy.Command,
	node *openapi.RegistrySvcNode,
	definition *openapi.RegistrySvcDefinition,
	deployment *deploy.Deployment,
) error {
	logger.Info(
		"Executing deploy kill command",
		slog.String("deploymentId", deployment.Id),
	)
	client := ns.clientFactory.Client(
		sdk.WithAddress(command.NodeUrl),
		sdk.WithToken(ns.token),
	)

	name := fmt.Sprintf("sup-%v", definition.Id)
	_, _, err := client.ContainerSvcAPI.StopContainer(ctx).Body(
		openapi.ContainerSvcStopContainerRequest{
			Name: &name,
		},
	).Execute()

	return err
}

func (ns *DeployService) executeStartCommand(
	ctx context.Context,
	command *deploy.Command,
	node *openapi.RegistrySvcNode,
	definition *openapi.RegistrySvcDefinition,
	deployment *deploy.Deployment,
) error {
	logger.Info(
		"Executing deploy start command",
		slog.String("deploymentId", deployment.Id),
	)
	client := ns.clientFactory.Client(
		sdk.WithAddress(command.NodeUrl),
		sdk.WithToken(ns.token),
	)

	err := ns.makeSureItRuns(client, ctx, definition, deployment)

	if err != nil {
		logger.Warn("Error executing deploy start command",
			slog.String("deploymentId", deployment.Id),
			slog.Any("error", err),
		)

		deployment.Status = deploy.DeploymentStatus(
			openapi.DeploymentStatusError,
		)
		deployment.Details = err.Error()

		writeErr := ns.deploymentStore.Query(datastore.Id(command.DeploymentId)).
			UpdateFields(map[string]any{
				"status":  deployment.Status,
				"details": deployment.Details,
			})

		if writeErr != nil {
			return writeErr
		}

		return err
	}

	logger.Debug("Successfully executed deploy start command",
		slog.String("deploymentId", deployment.Id),
	)

	deployment.Status = deploy.DeploymentStatus(openapi.DeploymentStatusOK)
	deployment.Details = ""

	writeErr := ns.deploymentStore.Query(datastore.Id(command.DeploymentId)).
		UpdateFields(map[string]any{
			"status":  deployment.Status,
			"details": deployment.Details,
		})
	if writeErr != nil {
		return writeErr
	}

	ur, err := url.Parse(node.Url)
	if err != nil {
		return errors.Wrap(err, "error parsing node url")
	}
	if definition.Ports != nil {
		ur.Host = strings.Replace(
			ur.Host,
			ur.Port(),
			// @todo multiport issue
			fmt.Sprintf("%v", definition.Ports[0].Host),
			1,
		)
	}

	// Services are expected to also call the RegisterInstance endpoint
	// to provide their slug if they want to be proxied by the slug.

	_, _, err = client.RegistrySvcAPI.RegisterInstance(ctx).Body(
		openapi.RegistrySvcRegisterInstanceRequest{
			Id:           openapi.PtrString(sdk.Id("inst")),
			DeploymentId: openapi.PtrString(deployment.Id),
			Url:          ur.String(),
			Host:         openapi.PtrString(ur.Hostname()),
			// @todo multiport
			Port:   &definition.Ports[0].Host,
			Scheme: openapi.PtrString(ur.Scheme),
			Path:   openapi.PtrString(ur.Path),
		},
	).Execute()
	if err != nil {
		return errors.Wrap(err, "error registering instance")
	}

	return nil
}

func (ns *DeployService) makeSureItRuns(
	client *openapi.APIClient,
	ctx context.Context,
	definition *openapi.RegistrySvcDefinition,
	deployment *deploy.Deployment,
) (err error) {
	defer func() {
		err = sdk.OpenAPIError(err)
	}()

	if definition == nil {
		return fmt.Errorf(
			"definition '%v' cannot be found",
			deployment.DefinitionId,
		)
	}

	if definition.Image != nil {
		runContainerReq := openapi.ContainerSvcRunContainerRequest{
			Image: definition.Image.Name,
			Ports: []openapi.ContainerSvcPortMapping{},
			Names: []string{
				fmt.Sprintf("openorch-%v", definition.Id),
			},
		}
		for _, port := range definition.Ports {
			runContainerReq.Ports = append(runContainerReq.Ports, openapi.ContainerSvcPortMapping{
				Host:     port.Host,
				Internal: port.Internal,
			})
		}

		_, _, err = client.ContainerSvcAPI.RunContainer(ctx).Body(
			runContainerReq,
		).Execute()
	} else {
		var checkoutRsp *openapi.SourceSvcCheckoutRepoResponse

		checkoutRsp, _, err = client.SourceSvcAPI.CheckoutRepo(ctx).Body(openapi.SourceSvcCheckoutRepoRequest{
			Url: &definition.Repository.Url,
		}).Execute()

		if err != nil {
			return err
		}

		buildContext := *checkoutRsp.Dir
		if definition.Repository.BuildContext != nil {
			buildContext = path.Join(buildContext, *definition.Repository.BuildContext)
		}

		_, _, err = client.ContainerSvcAPI.BuildImage(ctx).Body(
			openapi.ContainerSvcBuildImageRequest{
				ContextPath:    buildContext,
				DockerfilePath: definition.Repository.ContainerFile,
				Name:           fmt.Sprintf("%v-%v", containerPrefix, definition.Id),
			},
		).Execute()

		if err != nil {
			return err
		}

		runContainerReq := openapi.ContainerSvcRunContainerRequest{
			Image: fmt.Sprintf("openorch-%v", definition.Id),
			Ports: []openapi.ContainerSvcPortMapping{
				{
					// @todo multiport issues
					Internal: definition.Image.InternalPorts[0],
					Host:     definition.Repository.Ports[0],
				},
			},
			Names: []string{fmt.Sprintf("openorch-%v", definition.Id)},
		}

		_, _, err = client.ContainerSvcAPI.RunContainer(ctx).Body(
			runContainerReq,
		).Execute()
	}

	return err
}

func (ns *DeployService) getDeploymentById(
	deploymentId string,
) (*deploy.Deployment, error) {
	deploymentIs, err := ns.deploymentStore.Query(datastore.Id(deploymentId)).
		Find()
	if err != nil {
		return nil, err
	}

	deployment := deploymentIs[0].(*deploy.Deployment)

	return deployment, err
}

func rewritePort(inputURL string, newPort string) (string, error) {
	ur, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}

	// Split the host part and rewrite the port
	host := ur.Hostname() // Get the hostname without the port
	ur.Host = fmt.Sprintf("%s:%s", host, newPort)

	return ur.String(), nil
}
