package api_key

import (
	"fmt"
	"time"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Create(
	cmd *cobra.Command,
	args []string,
	opts createOptions,
	activeOrganizationChanged bool,
) error {
	name := ""
	if len(args) > 0 {
		name = args[0]
	}

	expiresAt, err := parseExpiration(opts.expiresAt, opts.expiresIn)
	if err != nil {
		return err
	}

	req := openapi.UserSvcCreateApiKeyRequest{}
	setString(&req.AppHost, opts.appHost)
	setString(&req.AppId, opts.appId)
	setString(&req.UserId, opts.userId)
	setString(&req.Name, name)
	setString(&req.ExpiresAt, expiresAt)
	if activeOrganizationChanged {
		req.ActiveOrganizationId = &opts.activeOrganizationId
	}

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)
	rsp, _, err := cf.Client(client.WithToken(token)).
		UserSvcAPI.CreateApiKey(cmd.Context()).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to create API key")
	}

	if opts.keyOnly {
		fmt.Println(rsp.GetKey())
		return nil
	}

	if opts.verbose {
		out, err := yaml.Marshal(rsp)
		if err != nil {
			return errors.Wrap(err, "failed to marshal response")
		}
		fmt.Print(string(out))
		return nil
	}

	apiKey := rsp.GetApiKey()
	fmt.Printf("KEY\t%s\n", rsp.GetKey())
	fmt.Printf("ID\t%s\n", apiKey.GetId())
	fmt.Printf("PREFIX\t%s\n", apiKey.GetPrefix())
	if apiKey.GetUserId() != "" {
		fmt.Printf("USER ID\t%s\n", apiKey.GetUserId())
	}
	if apiKey.GetAppId() != "" {
		fmt.Printf("APP ID\t%s\n", apiKey.GetAppId())
	}
	if apiKey.GetActiveOrganizationId() != "" {
		fmt.Printf("ACTIVE ORG ID\t%s\n", apiKey.GetActiveOrganizationId())
	}
	if apiKey.GetExpiresAt() != "" {
		fmt.Printf("EXPIRES AT\t%s\n", apiKey.GetExpiresAt())
	}

	return nil
}

func parseExpiration(expiresAt string, expiresIn string) (string, error) {
	if expiresAt != "" && expiresIn != "" {
		return "", errors.New("--expires-at and --expires-in are mutually exclusive")
	}
	if expiresAt != "" {
		t, err := time.Parse(time.RFC3339, expiresAt)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse --expires-at")
		}
		return t.Format(time.RFC3339), nil
	}
	if expiresIn != "" {
		d, err := time.ParseDuration(expiresIn)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse --expires-in")
		}
		t := time.Now().Add(d).UTC()
		return t.Format(time.RFC3339), nil
	}

	return "", nil
}

func setString(target **string, value string) {
	if value == "" {
		return
	}
	*target = openapi.PtrString(value)
}
