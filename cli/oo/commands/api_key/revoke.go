package api_key

import (
	"fmt"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Revoke(cmd *cobra.Command, args []string, opts revokeOptions) error {
	ids := append([]string{}, opts.ids...)
	ids = append(ids, args...)
	if len(ids) == 0 {
		return errors.New("at least one API key id is required")
	}

	req := openapi.UserSvcRevokeApiKeysRequest{
		Ids: ids,
	}
	setString(&req.UserId, opts.userId)

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)
	_, _, err = cf.Client(client.WithToken(token)).
		UserSvcAPI.RevokeApiKeys(cmd.Context()).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to revoke API keys")
	}

	fmt.Printf("Revoked %d API key(s)\n", len(ids))
	return nil
}
