package app

import "github.com/spf13/cobra"

func AddAppCommands(rootCmd *cobra.Command) {
	var appCmd = &cobra.Command{
		Use:     "app",
		Aliases: []string{"apps"},
		Short:   "Manage apps",
	}

	var (
		listIds   []string
		listHosts []string
	)

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List apps",
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, args, listIds, listHosts)
		},
	}

	listCmd.Flags().StringArrayVar(&listIds, "id", nil, "Filter by app ID")
	listCmd.Flags().StringArrayVar(&listHosts, "host", nil, "Filter by app host")

	var (
		saveId   string
		saveHost string
	)

	var saveCmd = &cobra.Command{
		Use:   "save [filePath | dirPath]",
		Args:  cobra.MaximumNArgs(1),
		Short: "Create or save apps from a file, directory, or flags",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Save(cmd, args, saveId, saveHost)
		},
	}

	saveCmd.Flags().StringVar(&saveId, "id", "", "ID of the app to save")
	saveCmd.Flags().StringVar(&saveHost, "host", "", "Host of the app to save")

	appCmd.AddCommand(listCmd)
	appCmd.AddCommand(saveCmd)

	rootCmd.AddCommand(appCmd)
}
