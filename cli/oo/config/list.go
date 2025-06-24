package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	var keys []string
	if len(args) > 0 {
		keys = args
	}

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.ConfigSvcListConfigsRequest{
		Keys: keys,
	}

	rsp, _, err := cf.Client(client.WithToken(token)).
		ConfigSvcAPI.ListConfigs(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list configs")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"CONFIG KEY\tAPP\tJSON",
	)

	for _, config := range rsp.Configs {
		jsonValue, err := json.Marshal(config.Data)
		if err != nil {
			return errors.Wrapf(err, "failed to marshal config data for ID '%s'", config)
		}

		// @todo We should definitely not rely on the ID format here,
		// we only do it for backwards compatibility.
		parts := strings.Split(config.Id, "_")

		key := config.Key
		if key == nil {
			key = &parts[1]
		}

		app := config.App
		if app == nil {
			app = &parts[0]
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\n",
			*key,
			*app,
			jsonValue,
		)
	}

	return nil
}
