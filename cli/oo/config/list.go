package config

import (
	"encoding/json"
	"fmt"
	"os"
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

	var ids []string
	if len(args) > 0 {
		ids = args
	}

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.ConfigSvcListConfigsRequest{
		Ids: ids,
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
		"CONFIG ID\tAPP\tJSON",
	)

	for _, config := range rsp.Configs {
		jsonValue, err := json.Marshal(config.Data)
		if err != nil {
			return errors.Wrapf(err, "failed to marshal config data for ID '%s'", config)
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\n",
			config.Id,
			config.AppId,
			jsonValue,
		)
	}

	return nil
}
