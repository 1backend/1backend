package cert

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/config"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Raw(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	var key string
	if len(args) > 0 {
		key = args[0]
	} else {
		return errors.New("cert key is required")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.ProxySvcListCertsRequest{
		Ids: []string{key},
	}
	rsp, _, err := cf.Client(client.WithToken(token)).
		ProxySvcAPI.ListCerts(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list certs")
	}

	fmt.Print(*rsp.Certs[0].Cert)

	return nil
}
