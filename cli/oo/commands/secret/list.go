package secret

import (
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
func List(cmd *cobra.Command, args []string, show bool) error {
	ctx := cmd.Context()

	var id string
	if len(args) > 0 {
		id = args[0]
	}

	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.SecretSvcListSecretsRequest{
		Id: openapi.PtrString(id),
	}

	rsp, _, err := cf.Client(client.WithToken(token)).
		SecretSvcAPI.ListSecrets(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list secrets")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"SECRET ID\tAPP\tLENGTH\tVALUE",
	)

	for _, secret := range rsp.Secrets {
		length := len(secret.Value)

		value := strings.Repeat("*", length)
		if show {
			value = secret.Value
		}

		app := "unnamed"
		if secret.AppId != "" {
			app = secret.AppId
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%d\n",
			secret.Id,
			app,
			length,
			value,
		)
	}

	return nil
}
