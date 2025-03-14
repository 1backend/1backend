package secret

import (
	"fmt"

	"github.com/openorch/openorch/cli/oo/config"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// IsSecure
func IsSecure(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := sdk.NewApiClientFactory(url)

	rsp, _, err := cf.Client(sdk.WithToken(token)).
		SecretSvcAPI.IsSecure(ctx).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list secrets")
	}

	if !rsp.IsSecure {
		return errors.New(
			"secret svc is not secure: it is using the default encryption key from the open-source codebase",
		)
	}

	fmt.Println("Service is secure.")

	return nil
}
