package api_key

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, opts listOptions) error {
	req := openapi.UserSvcListApiKeysRequest{
		Ids: opts.ids,
	}
	setString(&req.AppHost, opts.appHost)
	setString(&req.AppId, opts.appId)
	setString(&req.UserId, opts.userId)
	if opts.includeRevoked {
		req.IncludeRevoked = openapi.PtrBool(true)
	}

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)
	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.ListApiKeys(cmd.Context()).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list API keys")
	}

	if opts.verbose {
		out, err := yaml.Marshal(rsp.GetApiKeys())
		if err != nil {
			return errors.Wrap(err, "failed to marshal API keys")
		}
		fmt.Print(string(out))
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"API KEY ID\tUSER ID\tAPP ID\tNAME\tPREFIX\tACTIVE ORG\tEXPIRES\tLAST USED\tREVOKED",
	)

	for _, apiKey := range rsp.GetApiKeys() {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			apiKey.GetId(),
			apiKey.GetUserId(),
			apiKey.GetAppId(),
			apiKey.GetName(),
			apiKey.GetPrefix(),
			apiKey.GetActiveOrganizationId(),
			formatOptionalString(apiKey.GetExpiresAt()),
			formatOptionalString(apiKey.GetLastUsedAt()),
			formatOptionalString(apiKey.GetRevokedAt()),
		)
	}

	return nil
}

func formatOptionalString(v string) string {
	if v == "" {
		return "-"
	}

	return v
}
