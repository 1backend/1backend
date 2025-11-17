package config

import (
	"github.com/spf13/cobra"
)

func AddConfigCommands(rootCmd *cobra.Command) {
	var configCmd = &cobra.Command{
		Use:     "config",
		Aliases: []string{"co", "configs"},
		Short:   "Manage configs",
	}

	addSaveCommand(configCmd)
	addListCommand(configCmd)

	rootCmd.AddCommand(configCmd)
}

func addSaveCommand(rootCmd *cobra.Command) {
	var saveCmd = &cobra.Command{
		Use:     "save [filePath] | [folderPath]",
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(2),
		Short:   "Save config from a file",
		Long: `The 'save' command allows you to save configs under your own slug, or if you are an admin, under any slug. Can save into any app.

Example:
# Save a config from a file
save ./config.yaml

# Example contents of 'config.yaml':
appHost: "my-app"
data:
  # Limitation: only a single top level key can be saved,
  # even by admins.
  key1:
    subkey1: "subvalue1"
    subkey2:
      subsubkey1: "subsubvalue1"
      subsubkey2:
        deepkey: "deepvalue"
`,
		RunE: Save,
	}

	rootCmd.AddCommand(saveCmd)
}

func addListCommand(rootCmd *cobra.Command) {
	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Args:    cobra.MaximumNArgs(1),
		Short:   "List configs for an app",
		RunE:    List,
	}
	rootCmd.AddCommand(listCmd)
}
