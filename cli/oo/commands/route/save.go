package route

import (
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [--id --target | filePath | dirPath]
func Save(
	cmd *cobra.Command,
	args []string,
	id string,
	target string,
) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	// Case 1: Flags-based route
	if id != "" && target != "" {
		_, _, err := cf.Client(client.WithToken(token)).
			ProxySvcAPI.SaveRoutes(ctx).
			Body(
				openapi.ProxySvcSaveRoutesRequest{
					Routes: []openapi.ProxySvcRouteInput{
						{
							Id:     id,
							Target: &target,
						},
					},
				},
			).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save route")
		}
		return nil
	}

	// Case 2: File or directory-based routes
	path := args[0]

	routes, err := util.CollectFromPath[openapi.ProxySvcRouteInput](path, "routes")
	if err != nil {
		return err
	}

	// Make a single API call to save all routes
	_, _, err = cf.Client(client.WithToken(token)).
		ProxySvcAPI.SaveRoutes(ctx).
		Body(openapi.ProxySvcSaveRoutesRequest{
			Routes: routes,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save routes")
	}

	return nil
}
