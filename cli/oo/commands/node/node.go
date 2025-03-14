package node

import "github.com/spf13/cobra"

func AddNodeCommands(rootCmd *cobra.Command) {
	var nodeCmd = &cobra.Command{
		Use:     "node",
		Aliases: []string{"n", "nodes"},
		Short:   "Manage nodes",
	}

	var deleteCmd = &cobra.Command{
		Use:     "remove [url]",
		Short:   "Remove a node by URL",
		Aliases: []string{"rm"},
		Args:    cobra.ExactArgs(1),
		RunE:    Delete,
	}

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List nodes",
		RunE:    List,
	}

	nodeCmd.AddCommand(deleteCmd)
	nodeCmd.AddCommand(listCmd)

	rootCmd.AddCommand(nodeCmd)
}
