package deployment

import "github.com/spf13/cobra"

func AddDeploymentCommands(rootCmd *cobra.Command) {
	var deploymentCmd = &cobra.Command{
		Use:     "deployment",
		Aliases: []string{"depl", "deployments"},
		Short:   "Manage deployments",
	}

	var saveCmd = &cobra.Command{
		Use:   "save [filePath]",
		Args:  cobra.ExactArgs(1),
		Short: "Save deployment(s) found in a JSON or YAML file",
		RunE:  Save,
	}

	var deleteCmd = &cobra.Command{
		Use:     "remove [id]",
		Short:   "Remove a deployment",
		Aliases: []string{"rm"},
		Args:    cobra.ExactArgs(1),
		RunE:    Delete,
	}

	var full bool

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List deployments",
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, args, full)
		},
	}

	listCmd.Flags().
		BoolVar(&full, "full", false, "Show full deployment details without truncation")

	deploymentCmd.AddCommand(saveCmd)
	deploymentCmd.AddCommand(deleteCmd)
	deploymentCmd.AddCommand(listCmd)

	rootCmd.AddCommand(deploymentCmd)
}
