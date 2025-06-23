package config

import (
	"github.com/spf13/cobra"
)

func AddConfigCommands(rootCmd *cobra.Command) {
	addSaveCommand(rootCmd)

	var configCmd = &cobra.Command{
		Use:     "config",
		Aliases: []string{"co", "configs"},
		Short:   "Manage configs",
	}
	rootCmd.AddCommand(configCmd)
}

func addSaveCommand(rootCmd *cobra.Command) {
	var saveCmd = &cobra.Command{
		Use:     "save [filePath]",
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(2),
		Short:   "Save config from a file",
		Long:    `The 'save' command allows you to save configs under your own slug, or if you are an admin, under any slug. Can save into any app.`,
	}

	rootCmd.AddCommand(saveCmd)
}
