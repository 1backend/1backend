package cert

import "github.com/spf13/cobra"

func AddCertCommands(rootCmd *cobra.Command) {
	var certCmd = &cobra.Command{
		Use:     "cert",
		Aliases: []string{"ce", "certs"},
		Short:   "Manage certs",
	}

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List certs",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, args)
		},
	}

	var rawCmd = &cobra.Command{
		Use:     "raw [key]",
		Short:   "Get raw cert by key",
		Aliases: []string{"r"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Raw(cmd, args)
		},
	}

	listCmd.Flags().
		StringArrayP("ids", "i", nil, "Ids to filter on.")

	certCmd.AddCommand(listCmd)
	certCmd.AddCommand(rawCmd)

	rootCmd.AddCommand(certCmd)
}
