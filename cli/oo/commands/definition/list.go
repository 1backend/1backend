package definition

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	rsp, _, err := cf.Client(client.WithToken(token)).
		RegistrySvcAPI.ListDefinitions(ctx).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list service definitions")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"ID\tIMAGE NAME\tREPO URL\tREPO BUILD CONTEXT\tREPO CONTAINER FILE\tINTERNAL PORT",
	)

	for _, definition := range rsp.Definitions {
		imageName := ""
		if definition.Image != nil {
			imageName = definition.Image.Name
		}
		repoUrl := ""
		if definition.Repository != nil {
			repoUrl = definition.Repository.Url
		}
		buildContext := ""
		if definition.Repository != nil &&
			definition.Repository.BuildContext != nil {
			buildContext = *definition.Repository.BuildContext
		}
		containerFile := ""
		if definition.Repository != nil &&
			definition.Repository.ContainerFile != nil {
			containerFile = *definition.Repository.ContainerFile
		}

		port := ""
		// @todo multiport issue
		if definition.Image != nil {
			port = fmt.Sprintf("%d", definition.Image.InternalPorts[0])
		}
		if definition.Repository != nil {
			port = fmt.Sprintf("%d", &definition.Repository.InternalPorts[0])
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%s\t%s\n",
			definition.Id,
			imageName,
			repoUrl,
			buildContext,
			containerFile,
			port,
		)
	}

	return nil
}
