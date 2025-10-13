package secret

import (
	"fmt"
	"sort"
	"strings"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Copy(cmd *cobra.Command, fromEnv string, toEnv string) error {
	if fromEnv == "" {
		return errors.New("flag --from is required")
	}
	if toEnv == "" {
		return errors.New("flag --to is required")
	}
	if fromEnv == toEnv {
		return errors.New("source and destination environments must differ")
	}

	ctx := cmd.Context()

	fromURL, fromToken, err := util.GetUrlAndTokenForEnv(cmd, fromEnv, "")
	if err != nil {
		return errors.Wrapf(err, "failed to load source environment '%s'", fromEnv)
	}

	toURL, toToken, err := util.GetUrlAndTokenForEnv(cmd, toEnv, "")
	if err != nil {
		return errors.Wrapf(err, "failed to load destination environment '%s'", toEnv)
	}

	fromFactory := client.NewApiClientFactory(fromURL)
	fromClient := fromFactory.Client(client.WithToken(fromToken))

	listRsp, _, err := fromClient.SecretSvcAPI.
		ListSecrets(ctx).
		Body(openapi.SecretSvcListSecretsRequest{
			AllApps: openapi.PtrBool(true),
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list secrets from source environment")
	}

	if len(listRsp.Secrets) == 0 {
		fmt.Fprintf(
			cmd.OutOrStdout(),
			"No secrets found in env '%s'\n",
			fromEnv,
		)
		return nil
	}

	appIDs := map[string]struct{}{}
	for _, secret := range listRsp.Secrets {
		if secret.AppId != "" {
			appIDs[secret.AppId] = struct{}{}
		}
	}

	appIdToHost := map[string]string{}
	if len(appIDs) > 0 {
		ids := make([]string, 0, len(appIDs))
		for id := range appIDs {
			ids = append(ids, id)
		}
		sort.Strings(ids)

		appsRsp, _, err := fromClient.UserSvcAPI.
			ListApps(ctx).
			Body(openapi.UserSvcListAppsRequest{
				Ids: ids,
			}).
			Execute()
		if err != nil {
			return errors.Wrap(err, "failed to resolve app hosts in source environment")
		}

		for _, app := range appsRsp.Apps {
			appIdToHost[app.Id] = app.Host
		}
	}

	toFactory := client.NewApiClientFactory(toURL)
	toSecretClient := toFactory.Client(client.WithToken(toToken)).SecretSvcAPI

	encryptedValues := make([]string, len(listRsp.Secrets))
	for idx, secret := range listRsp.Secrets {
		encryptRsp, _, err := toSecretClient.
			EncryptValue(ctx).
			Body(openapi.SecretSvcEncryptValueRequest{
				Value: openapi.PtrString(secret.Value),
			}).
			Execute()
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to encrypt secret '%s' for destination environment",
				secret.Id,
			)
		}

		if encryptRsp.Value == nil {
			return fmt.Errorf("encryption response missing value for secret '%s'", secret.Id)
		}

		encryptedValues[idx] = encryptRsp.GetValue()
	}

	inputs := make([]openapi.SecretSvcSecretInput, 0, len(listRsp.Secrets))
	missingHosts := map[string]struct{}{}
	for idx, secret := range listRsp.Secrets {
		input := openapi.SecretSvcSecretInput{
			Id:                secret.Id,
			Value:             openapi.PtrString(encryptedValues[idx]),
			Readers:           secret.Readers,
			Writers:           secret.Writers,
			Deleters:          secret.Deleters,
			CanChangeReaders:  secret.CanChangeReaders,
			CanChangeWriters:  secret.CanChangeWriters,
			CanChangeDeleters: secret.CanChangeDeleters,
			Encrypted:         openapi.PtrBool(true),
			Checksum:          secret.Checksum,
			ChecksumAlgorithm: secret.ChecksumAlgorithm,
		}

		if host, ok := appIdToHost[secret.AppId]; ok && host != "" {
			input.AppHost = openapi.PtrString(host)
		} else if secret.AppId != "" {
			missingHosts[secret.AppId] = struct{}{}
		}

		inputs = append(inputs, input)
	}

	_, _, err = toSecretClient.
		SaveSecrets(ctx).
		Body(openapi.SecretSvcSaveSecretsRequest{
			Secrets: inputs,
		}).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to save secrets in destination environment")
	}

	if len(missingHosts) > 0 {
		ids := make([]string, 0, len(missingHosts))
		for id := range missingHosts {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		fmt.Fprintf(
			cmd.ErrOrStderr(),
			"Warning: could not resolve app host for app ids: %s\n",
			strings.Join(ids, ", "),
		)
	}

	fmt.Fprintf(
		cmd.OutOrStdout(),
		"Copied %d secrets from '%s' to '%s'\n",
		len(inputs),
		fromEnv,
		toEnv,
	)

	return nil
}
