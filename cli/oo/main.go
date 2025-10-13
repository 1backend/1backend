package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"

	"github.com/1backend/1backend/cli/oo/commands/app"
	call "github.com/1backend/1backend/cli/oo/commands/call"
	"github.com/1backend/1backend/cli/oo/commands/cert"
	definition "github.com/1backend/1backend/cli/oo/commands/definition"
	deployment "github.com/1backend/1backend/cli/oo/commands/deployment"
	"github.com/1backend/1backend/cli/oo/commands/enroll"
	"github.com/1backend/1backend/cli/oo/commands/env"
	instance "github.com/1backend/1backend/cli/oo/commands/instance"
	"github.com/1backend/1backend/cli/oo/commands/node"
	"github.com/1backend/1backend/cli/oo/commands/permit"
	"github.com/1backend/1backend/cli/oo/commands/route"
	"github.com/1backend/1backend/cli/oo/commands/save"
	secret "github.com/1backend/1backend/cli/oo/commands/secret"
	"github.com/1backend/1backend/cli/oo/commands/user"
	"github.com/1backend/1backend/cli/oo/config"
	sdk "github.com/1backend/1backend/sdk/go"
)

func main() {
	var (
		globalEnv string
		globalApp string
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n%s", r, debug.Stack())
			os.Exit(2)
		}
	}()

	var rootCmd = &cobra.Command{
		Use:   "oo",
		Short: "1Backend CLI",

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.WithValue(cmd.Context(), "env", globalEnv)
			ctx = context.WithValue(ctx, "app-host", globalApp)
			cmd.SetContext(ctx)
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceErrors: false,
		SilenceUsage:  true,
	}

	rootCmd.PersistentFlags().
		StringVarP(&globalEnv, "env", "e", "", "Environment short name (overrides selected)")
	rootCmd.PersistentFlags().
		StringVarP(&globalApp, "app-host", "a", sdk.DefaultAppHost, "App host name")

	app.AddAppCommands(rootCmd)
	user.AddUserCommands(rootCmd)
	env.AddEnvCommands(rootCmd)
	definition.AddDefinitionCommands(rootCmd)
	instance.AddInstanceCommands(rootCmd)
	deployment.AddDeploymentCommands(rootCmd)
	node.AddNodeCommands(rootCmd)
	secret.AddSecretCommands(rootCmd)
	permit.AddPermitCommands(rootCmd)
	enroll.AddEnrollCommands(rootCmd)
	route.AddRouteCommands(rootCmd)
	cert.AddCertCommands(rootCmd)
	config.AddConfigCommands(rootCmd)
	save.AddSaveCommand(rootCmd)

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
