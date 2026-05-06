package redirect

import "github.com/spf13/cobra"

func AddRedirectCommands(rootCmd *cobra.Command) {
	var redirectCmd = &cobra.Command{
		Use:     "redirect",
		Aliases: []string{"redirects"},
		Short:   "Manage redirects",
	}

	var (
		id         string
		target     string
		statusCode int
	)

	var saveCmd = &cobra.Command{
		Use:     "save [filePath | dirPath]",
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(2),
		Short:   "Save redirects from a file or directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Save(cmd, args, id, target, statusCode)
		},
	}

	saveCmd.Flags().StringVar(&id, "id", "", "ID of the redirect to save")
	saveCmd.Flags().StringVar(&target, "target", "", "Target URL of the redirect to save")
	saveCmd.Flags().IntVar(&statusCode, "status-code", 0, "HTTP redirect status code, defaults to 308")

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List redirects",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, args)
		},
	}

	listCmd.Flags().
		StringArrayP("ids", "i", nil, "Ids to filter on.")

	redirectCmd.AddCommand(saveCmd)
	redirectCmd.AddCommand(listCmd)

	rootCmd.AddCommand(redirectCmd)
}
