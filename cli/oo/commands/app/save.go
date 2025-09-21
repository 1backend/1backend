package app

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save handles the `oo app save` command.
func Save(cmd *cobra.Command, args []string, id string, host string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	inputs := []openapi.UserSvcApp{}
	if host != "" {
		inputs = append(inputs, openapi.UserSvcApp{Id: id, Host: host})
	}

	if len(args) > 0 {
		path := args[0]
		fileInputs, err := util.CollectFromPath[openapi.UserSvcApp](path, "apps")
		if err != nil {
			return err
		}
		inputs = append(inputs, fileInputs...)
	}

	if len(inputs) == 0 {
		return errors.New("no app data provided")
	}

	for _, input := range inputs {
		if input.Host == "" {
			return errors.New("host is required for each app")
		}

		req := openapi.UserSvcReadAppRequest{}
		req.SetHost(input.Host)

		rsp, _, err := cf.Client(client.WithToken(token)).
			UserSvcAPI.ReadApp(ctx).
			Body(req).
			Execute()
		if err != nil {
			return errors.Wrapf(err, "failed to save app '%s'", input.Host)
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Saved app %s (ID: %s)\n", rsp.App.Host, rsp.App.Id)
	}

	return nil
}
