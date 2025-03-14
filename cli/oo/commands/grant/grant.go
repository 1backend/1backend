package grant

import "github.com/spf13/cobra"

func AddGrantCommands(rootCmd *cobra.Command) {
	var grantCmd = &cobra.Command{
		Use:     "grant",
		Aliases: []string{"g", "grants"},
		Short:   "Manage grants",
	}

	var saveCmd = &cobra.Command{
		Use:   "save [filePath | dirPath]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Save grants from files",
		Long: `The 'save' command allows you to save grants in two ways:
	
	1. A single grant from a YAML file.
	2. Multiple grants from a YAML file.`,
		Example: `# Save a single grant from a file
save ./grantA.yaml

# Example contents of 'grantA.yaml':
id: "grant-id-1"
permissionId: "user-svc:role:read"
slugs:
  - "service1"
  - "user1"

# Save multiple grants from a file
save ./grants.yaml

# Example contents of 'grants.yaml':
- id: "grant-id-1"
  permissionId: "user-svc:role:read"
  slugs:
    - "service1"
    - "user1"
- id: "grant-id-2"
  permissionId: "user-svc:role:write"
  slugs:
    - "service2"
    - "user2"`,
		RunE: Save,
	}

	// var removeCmd = &cobra.Command{
	// 	Use:     "remove [id]",
	// 	Short:   "Delete a grant",
	// 	Aliases: []string{"rm"},
	// 	RunE:    Remove,
	// }

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List grants",
		Aliases: []string{"ls"},
		RunE:    List,
	}

	grantCmd.AddCommand(saveCmd)
	// grantCmd.AddCommand(removeCmd)
	grantCmd.AddCommand(listCmd)

	rootCmd.AddCommand(grantCmd)
}
