package main

import (
	"os"

	"github.com/spf13/cobra"

	call "github.com/1backend/1backend/cli/oo/commands/call"
	definition "github.com/1backend/1backend/cli/oo/commands/definition"
	deployment "github.com/1backend/1backend/cli/oo/commands/deployment"
	"github.com/1backend/1backend/cli/oo/commands/env"
	"github.com/1backend/1backend/cli/oo/commands/grant"
	instance "github.com/1backend/1backend/cli/oo/commands/instance"
	"github.com/1backend/1backend/cli/oo/commands/invite"
	"github.com/1backend/1backend/cli/oo/commands/node"
	secret "github.com/1backend/1backend/cli/oo/commands/secret"
	"github.com/1backend/1backend/cli/oo/commands/user"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "oo",
		Short: "1Backend CLI",

		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceUsage: true,
	}

	user.AddUserCommands(rootCmd)
	env.AddEnvCommands(rootCmd)
	definition.AddDefinitionCommands(rootCmd)
	instance.AddInstanceCommands(rootCmd)
	deployment.AddDeploymentCommands(rootCmd)
	node.AddNodeCommands(rootCmd)
	secret.AddSecretCommands(rootCmd)
	grant.AddGrantCommands(rootCmd)
	invite.AddInviteCommands(rootCmd)

	addCallCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func addCallCommands(rootCmd *cobra.Command) {
	var postCmd = &cobra.Command{
		Use:                "post [service] [password] [key=value]...",
		Args:               cobra.ArbitraryArgs,
		Short:              "Make a POST request to a service",
		RunE:               call.Post,
		DisableFlagParsing: true,
	}
	var getCmd = &cobra.Command{
		Use:                "get [service] [password] [key=value]...",
		Args:               cobra.ArbitraryArgs,
		Short:              "Make a GET request to a service",
		RunE:               call.Get,
		DisableFlagParsing: true,
	}
	var putCmd = &cobra.Command{
		Use:                "put [service] [password] [key=value]...",
		Args:               cobra.ArbitraryArgs,
		Short:              "Make a PUT request to a service",
		RunE:               call.Put,
		DisableFlagParsing: true,
	}
	var deleteCmd = &cobra.Command{
		Use:                "delete [service] [password] [key=value]...",
		Args:               cobra.ArbitraryArgs,
		Short:              "Make a DELETE request to a service",
		RunE:               call.Delete,
		DisableFlagParsing: true,
	}

	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(deleteCmd)
}
