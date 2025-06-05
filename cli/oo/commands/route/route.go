package route

import "github.com/spf13/cobra"

func AddRouteCommands(rootCmd *cobra.Command) {
	var routeCmd = &cobra.Command{
		Use:     "route",
		Aliases: []string{"r", "routes"},
		Short:   "Manage routes",
	}

	var saveCmd = &cobra.Command{
		Use:     "save [filePath | dirPath]",
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(2),
		Short:   "Save routes from a file or directory",
		RunE:    Save,
	}

	// var removeCmd = &cobra.Command{
	// 	Use:     "remove [id]",
	// 	Short:   "Delete a route",
	// 	Aliases: []string{"rm"},
	// 	RunE:    Remove,
	// }

	// removeCmd.Flags().StringArrayP("key", "k", []string{}, "Keys of routes to remove")
	// removeCmd.Flags().StringArrayP("id", "i", []string{}, "IDs of routes to remove")

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List routes",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, args)
		},
	}

	listCmd.Flags().
		StringArrayP("ids", "i", nil, "Ids to filter on.")

	routeCmd.AddCommand(saveCmd)
	//routeCmd.AddCommand(removeCmd)
	routeCmd.AddCommand(listCmd)

	rootCmd.AddCommand(routeCmd)
}
