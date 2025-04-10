package invite

import "github.com/spf13/cobra"

func AddInviteCommands(rootCmd *cobra.Command) {
	var inviteCmd = &cobra.Command{
		Use:     "invite",
		Aliases: []string{"i", "invites"},
		Short:   "Manage invites",
	}

	var (
		userId    string
		contactId string
	)

	var saveCmd = &cobra.Command{
		Use:   "save [filePath | dirPath]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Save invites from files",
		Long: `The 'save' command allows you to save invites in two ways:
	
	1. Save invite to role by user id or contact id.
	2. A single invite from a YAML file.
	3. Multiple invites from a YAML file.`,
		Example: `# Save a single invite from a file
save ./inviteA.yaml

# Example contents of 'inviteA.yaml':
id: "invite-id-1"
role: "user-svc:admin"
contactId: "user1@user1.com"

# Save multiple invites from a file
save ./invites.yaml

# Example contents of 'invites.yaml':
- id: "invite-id-1"
  role: "user-svc:admin"
  contactId: "user1@user1.com"
- id: "invite-id-2"
  role: "user-svc:admin"
  contactId: "user2@user2.com"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Save(cmd, args, userId, contactId)
		},
	}

	saveCmd.Flags().StringVar(&userId, "userId", "", "User ID to associate with the invite")
	saveCmd.Flags().StringVar(&contactId, "contactId", "", "Contact ID to associate with the invite")

	// var removeCmd = &cobra.Command{
	// 	Use:     "remove [id]",
	// 	Short:   "Delete a invite",
	// 	Aliases: []string{"rm"},
	// 	RunE:    Remove,
	// }

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List invites",
		Aliases: []string{"ls"},
		RunE:    List,
	}

	inviteCmd.AddCommand(saveCmd)
	// inviteCmd.AddCommand(removeCmd)
	inviteCmd.AddCommand(listCmd)

	rootCmd.AddCommand(inviteCmd)
}
