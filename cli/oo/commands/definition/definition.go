package definition

import "github.com/spf13/cobra"

func AddDefinitionCommands(rootCmd *cobra.Command) {
	var envCmd = &cobra.Command{
		Use:     "definition",
		Aliases: []string{"def", "definitions"},
		Short:   "Manage definitions",
	}

	var envSaveCmd = &cobra.Command{
		Use:   "save [filePath]",
		Args:  cobra.ExactArgs(1),
		Short: "Save definition(s) found in a JSON or YAML file",
		RunE:  Save,
	}

	var envDeleteCmd = &cobra.Command{
		Use:     "remove [id]",
		Short:   "Remove a definition",
		Aliases: []string{"rm"},
		Args:    cobra.ExactArgs(1),
		RunE:    Delete,
	}

	var envListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List definitions",
		RunE:    List,
	}

	envCmd.AddCommand(envSaveCmd)
	envCmd.AddCommand(envDeleteCmd)
	envCmd.AddCommand(envListCmd)

	rootCmd.AddCommand(envCmd)
}
