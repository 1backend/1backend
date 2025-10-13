package secret

import (
	"fmt"
	"sort"
	"strings"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Read(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	var id string
	if len(args) > 0 {
		id = args[0]
	}

	cf := client.NewApiClientFactory(url)
	apiClient := cf.Client(client.WithToken(token))

	req := openapi.SecretSvcListSecretsRequest{}
	if id != "" {
		req.Id = openapi.PtrString(id)
	}

	resp, _, err := apiClient.SecretSvcAPI.
		ListSecrets(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list secrets")
	}

	if len(resp.Secrets) == 0 {
		if id != "" {
			return fmt.Errorf("secret '%s' not found", id)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "# No secrets found")
		return nil
	}

	secrets := append([]openapi.SecretSvcSecret(nil), resp.Secrets...)
	sort.Slice(secrets, func(i, j int) bool {
		return secrets[i].Id < secrets[j].Id
	})

	appIds := []string{}
	for _, secret := range secrets {
		if secret.AppId != "" {
			appIds = append(appIds, secret.AppId)
		}
	}

	appIdToHost, err := util.AppIdsToHosts(cmd.Context(), apiClient, appIds)
	if err != nil {
		return errors.Wrap(err, "failed to resolve app hosts")
	}

	exportSecrets := make([]openapi.SecretSvcSecretInput, 0, len(secrets))
	missingHosts := map[string]struct{}{}

	for _, secret := range secrets {
		export := openapi.SecretSvcSecretInput{
			Id:                secret.Id,
			Value:             openapi.PtrString(secret.Value),
			Readers:           secret.Readers,
			Writers:           secret.Writers,
			Deleters:          secret.Deleters,
			CanChangeReaders:  secret.CanChangeReaders,
			CanChangeWriters:  secret.CanChangeWriters,
			CanChangeDeleters: secret.CanChangeDeleters,
			Checksum:          secret.Checksum,
			ChecksumAlgorithm: secret.ChecksumAlgorithm,
		}

		if secret.Encrypted != nil {
			export.Encrypted = openapi.PtrBool(secret.GetEncrypted())
		}

		if host, ok := appIdToHost[secret.AppId]; ok && host != "" {
			export.AppHost = openapi.PtrString(host)
		} else if secret.AppId != "" {
			missingHosts[secret.AppId] = struct{}{}
		}

		exportSecrets = append(exportSecrets, export)
	}

	if len(missingHosts) > 0 {
		ids := make([]string, 0, len(missingHosts))
		for id := range missingHosts {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		fmt.Fprintf(
			cmd.ErrOrStderr(),
			"Warning: missing app host mapping for app ids: %s\n",
			strings.Join(ids, ", "),
		)
	}

	out, err := yaml.Marshal(exportSecrets)
	if err != nil {
		return errors.Wrap(err, "failed to encode secrets as yaml")
	}

	fmt.Fprint(cmd.OutOrStdout(), string(out))

	return nil
}
