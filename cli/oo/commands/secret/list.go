package secret

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/config"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string, show bool) error {
	ctx := cmd.Context()

	var key string
	if len(args) > 0 {
		key = args[0]
	}

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.SecretSvcListSecretsRequest{
		Key: openapi.PtrString(key),
	}
	namespaceFlag, _ := cmd.Flags().GetString("namespace")
	if namespaceFlag != "" {
		req.Namespace = openapi.PtrString(namespaceFlag)
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
		"SECRET ID\tKEY\tNAMESPACE\tLENGTH\tVALUE",
	)

	for _, secret := range rsp.Secrets {
		length := len(*secret.Value)

		value := strings.Repeat("*", length)
		if show {
			value = *secret.Value
		}

		namespace := ""
		if secret.Namespace != nil {
			namespace = *secret.Namespace
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%d\t%s\n",
			*secret.Id,
			*secret.Key,
			namespace,
			length,
			value,
		)
	}

	return nil
}
