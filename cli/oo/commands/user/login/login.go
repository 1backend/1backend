package login

import (
	"fmt"
	"os"
	"strings"

	"github.com/openorch/openorch/cli/oo/config"
	"github.com/openorch/openorch/cli/oo/types"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// Login [slug] [password]
func Login(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	var slug, password string

	switch len(args) {
	case 0:
		fmt.Print("Enter slug: ")
		_, err := fmt.Scanln(&slug)
		if err != nil {
			return errors.Wrap(err, "failed to read slug")
		}
	case 1:
		// Use provided slug
		slug = args[0]
	default:
		// Both slug and password provided
		slug = args[0]
		password = args[1]
	}

	if password == "" {
		// Prompt for password securely if not already provided
		fmt.Print("Enter password: ")
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return errors.Wrap(err, "failed to read password")
		}
		password = strings.TrimSpace(string(bytePassword))
		fmt.Println() // Move to the next line after password input
	}

	if conf.Environments == nil {
		return fmt.Errorf("no environments found")
	}

	env, ok := conf.Environments[conf.SelectedEnvironment]
	if !ok {
		return fmt.Errorf(
			"failed to find selected env: %s",
			conf.SelectedEnvironment,
		)
	}

	cf := sdk.NewApiClientFactory(env.URL)

	rsp, _, err := cf.Client().
		UserSvcAPI.Login(cmd.Context()).
		Body(openapi.UserSvcLoginRequest{
			Slug:     &slug,
			Password: &password,
		}).
		Execute()
	if err != nil {
		return err
	}

	token := rsp.Token.GetToken()

	if env.Users == nil {
		env.Users = map[string]*types.User{}
	}

	env.Users[slug] = &types.User{
		Slug:  slug,
		Token: token,
	}
	env.SelectedUser = slug

	conf.Environments[env.ShortName] = env

	return config.SaveConfig(conf)
}
