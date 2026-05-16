package api_key

import (
	"github.com/spf13/cobra"
)

type createOptions struct {
	userId               string
	appHost              string
	appId                string
	activeOrganizationId string
	expiresAt            string
	expiresIn            string
	keyOnly              bool
	verbose              bool
}

type listOptions struct {
	userId         string
	appHost        string
	appId          string
	ids            []string
	includeRevoked bool
	verbose        bool
}

type revokeOptions struct {
	userId string
	ids    []string
}

func AddAPIKeyCommands(rootCmd *cobra.Command) {
	var apiKeyCmd = &cobra.Command{
		Use:     "api-key",
		Aliases: []string{"api-keys", "apikey", "apikeys"},
		Short:   "Manage API keys",
	}

	createOpts := createOptions{}
	var createCmd = &cobra.Command{
		Use:     "create [name]",
		Aliases: []string{"new", "c"},
		Args:    cobra.MaximumNArgs(1),
		Short:   "Create an API key",
		RunE: func(cmd *cobra.Command, args []string) error {
			activeOrgChanged := cmd.Flags().Changed("active-org-id")
			return Create(cmd, args, createOpts, activeOrgChanged)
		},
	}
	createCmd.Flags().StringVarP(&createOpts.userId, "user-id", "u", "", "User ID to create the API key for")
	createCmd.Flags().StringVar(&createOpts.appHost, "app-host", "", "App host for the API key")
	createCmd.Flags().StringVar(&createOpts.appId, "app-id", "", "App ID for the API key")
	createCmd.Flags().StringVarP(&createOpts.activeOrganizationId, "active-org-id", "o", "", "Active organization ID for the API key")
	createCmd.Flags().StringVar(&createOpts.expiresAt, "expires-at", "", "Expiration timestamp in RFC3339 format")
	createCmd.Flags().StringVar(&createOpts.expiresIn, "expires-in", "", "Expiration duration, for example 24h or 720h")
	createCmd.Flags().BoolVar(&createOpts.keyOnly, "key-only", false, "Only print the raw API key")
	createCmd.Flags().BoolVarP(&createOpts.verbose, "verbose", "v", false, "Print YAML output")

	listOpts := listOptions{}
	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		Short:   "List API keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, listOpts)
		},
	}
	listCmd.Flags().StringVarP(&listOpts.userId, "user-id", "u", "", "User ID to list API keys for")
	listCmd.Flags().StringVar(&listOpts.appHost, "app-host", "", "Filter by app host")
	listCmd.Flags().StringVar(&listOpts.appId, "app-id", "", "Filter by app ID")
	listCmd.Flags().StringArrayVarP(&listOpts.ids, "id", "i", []string{}, "API key ID to list")
	listCmd.Flags().BoolVar(&listOpts.includeRevoked, "include-revoked", false, "Include revoked API keys")
	listCmd.Flags().BoolVarP(&listOpts.verbose, "verbose", "v", false, "Print YAML output")

	revokeOpts := revokeOptions{}
	var revokeCmd = &cobra.Command{
		Use:     "revoke [id...]",
		Aliases: []string{"rm", "remove", "delete"},
		Args:    cobra.ArbitraryArgs,
		Short:   "Revoke API keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Revoke(cmd, args, revokeOpts)
		},
	}
	revokeCmd.Flags().StringVarP(&revokeOpts.userId, "user-id", "u", "", "User ID that owns the API keys")
	revokeCmd.Flags().StringArrayVarP(&revokeOpts.ids, "id", "i", []string{}, "API key ID to revoke")

	apiKeyCmd.AddCommand(createCmd)
	apiKeyCmd.AddCommand(listCmd)
	apiKeyCmd.AddCommand(revokeCmd)

	rootCmd.AddCommand(apiKeyCmd)
}
