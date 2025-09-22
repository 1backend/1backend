/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
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
permission: "user-svc:role:read"
slugs:
  - "service1"
  - "user1"

# Save multiple permits from a file
save ./permits.yaml

# Example contents of 'permits.yaml':
- id: "permit-id-1"
  permission: "user-svc:role:read"
  slugs:
    - "service1"
    - "user1"
- id: "permit-id-2"
  permission: "user-svc:role:write"
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
