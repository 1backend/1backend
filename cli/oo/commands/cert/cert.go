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

	listCmd.Flags().
		StringArrayP("ids", "i", nil, "Ids to filter on.")

	certCmd.AddCommand(listCmd)

	rootCmd.AddCommand(certCmd)
}
