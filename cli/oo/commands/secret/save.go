package secret

import (
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Save [id] [value] | [filePath | dirPath]
func Save(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	url, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)

	// Determine the save type: single id-value pair, file, or directory
	if len(args) == 2 {
		// Case 1: Single id-value pair
		id := args[0]
		value := args[1]

		secret := openapi.SecretSvcSecretInput{
			Id:    id,
			Value: openapi.PtrString(value),
		}

		_, _, err = cf.Client(client.WithToken(token)).
			SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: []openapi.SecretSvcSecretInput{secret},
			}).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save single secret")
		}

		return nil
	} else if len(args) == 1 {
		// Case 2: File or directory-based secrets
		path := args[0]

		secrets, err := util.CollectFromPath[openapi.SecretSvcSecretInput](path, "secrets")
		if err != nil {
			return err
		}

		// Make a single API call to save all secrets
		_, _, err = cf.Client(client.WithToken(token)).
			SecretSvcAPI.SaveSecrets(ctx).
			Body(openapi.SecretSvcSaveSecretsRequest{
				Secrets: secrets,
			}).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to save secrets")
		}

		return nil
	}

	return errors.New("invalid arguments: use 'save [id] [value]' or 'save [filePath | dirPath]'")
}
