package secret

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string, show, allApps, verbose bool) error {
	ctx := cmd.Context()

	var id string
	if len(args) > 0 {
		id = args[0]
	}

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.SecretSvcListSecretsRequest{
		AllApps: &allApps,
		Id:      openapi.PtrString(id),
	}

	rsp, _, err := cf.Client(client.WithToken(token)).
		SecretSvcAPI.ListSecrets(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list secrets")
	}

	appIds := []string{}
	for _, secret := range rsp.Secrets {
		if secret.AppId != "" {
			appIds = append(appIds, secret.AppId)
		}
	}

	appIdToHost, err := util.AppIdsToHosts(ctx, cf.Client(client.WithToken(token)), appIds)
	if err != nil {
		return errors.Wrap(err, "failed to resolve app hosts")
	}

	if verbose {
		out, err := yaml.Marshal(rsp.Secrets)
		if err != nil {
			return errors.Wrap(err, "failed to marshal secrets to yaml")
		}
		fmt.Print(string(out))
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"SECRET ID\tAPP HOST\tLENGTH\tVALUE",
	)

	for _, secret := range rsp.Secrets {
		length := len(secret.Value)

		value := strings.Repeat("*", length)
		if show {
			value = secret.Value
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%v\t%s\n",
			secret.Id,
			appIdToHost[secret.AppId],
			length,
			value,
		)
	}

	return nil
}
