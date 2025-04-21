package permit

import "github.com/spf13/cobra"

func AddPermitCommands(rootCmd *cobra.Command) {
	var permitCmd = &cobra.Command{
		Use:     "permit",
		Aliases: []string{"p", "permits"},
		Short:   "Manage permits",
	}

	var saveCmd = &cobra.Command{
		Use:     "save [filePath | dirPath]",
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(2),
		Short:   "Save permits from files",
		Long: `The 'save' command allows you to save permits in two ways:
	
	1. A single permit from a YAML file.
	2. Multiple permits from a YAML file.`,
		Example: `# Save a single permit from a file
save ./permitA.yaml

# Example contents of 'permitA.yaml':
id: "permit-id-1"
permissionId: "user-svc:role:read"
slugs:
  - "service1"
  - "user1"

# Save multiple permits from a file
save ./permits.yaml

# Example contents of 'permits.yaml':
- id: "permit-id-1"
  permissionId: "user-svc:role:read"
  slugs:
    - "service1"
    - "user1"
- id: "permit-id-2"
  permissionId: "user-svc:role:write"
  slugs:
    - "service2"
    - "user2"`,
		RunE: Save,
	}

	// var removeCmd = &cobra.Command{
	// 	Use:     "remove [id]",
	// 	Short:   "Delete a permit",
	// 	Aliases: []string{"rm"},
	// 	RunE:    Remove,
	// }

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List permits",
		Aliases: []string{"ls"},
		RunE:    List,
	}

	permitCmd.AddCommand(saveCmd)
	// permitCmd.AddCommand(removeCmd)
	permitCmd.AddCommand(listCmd)

	rootCmd.AddCommand(permitCmd)
}
