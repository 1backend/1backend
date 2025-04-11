package enroll

import "github.com/spf13/cobra"

func AddEnrollCommands(rootCmd *cobra.Command) {
	var enrollCmd = &cobra.Command{
		Use:     "enroll",
		Aliases: []string{"i", "enrolls"},
		Short:   "Manage enrolls",
	}

	var (
		userId    string
		contactId string
	)

	var saveCmd = &cobra.Command{
		Use:   "save [filePath | dirPath]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Save enrolls from files",
		Long: `The 'save' command allows you to save enrolls in two ways:
	
	1. Save enroll to role by user id or contact id.
	2. A single enroll from a YAML file.
	3. Multiple enrolls from a YAML file.`,
		Example: `# Save a single enroll from a file
save ./enrollA.yaml

# Example contents of 'enrollA.yaml':
id: "enroll-id-1"
role: "user-svc:admin"
contactId: "user1@user1.com"

# Save multiple enrolls from a file
save ./enrolls.yaml

# Example contents of 'enrolls.yaml':
- id: "enroll-id-1"
  role: "user-svc:admin"
  contactId: "user1@user1.com"
- id: "enroll-id-2"
  role: "user-svc:admin"
  contactId: "user2@user2.com"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Save(cmd, args, userId, contactId)
		},
	}

	saveCmd.Flags().StringVar(&userId, "userId", "", "User ID to associate with the enroll")
	saveCmd.Flags().StringVar(&contactId, "contactId", "", "Contact ID to associate with the enroll")

	// var removeCmd = &cobra.Command{
	// 	Use:     "remove [id]",
	// 	Short:   "Delete a enroll",
	// 	Aliases: []string{"rm"},
	// 	RunE:    Remove,
	// }

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List enrolls",
		Aliases: []string{"ls"},
		RunE:    List,
	}

	enrollCmd.AddCommand(saveCmd)
	// enrollCmd.AddCommand(removeCmd)
	enrollCmd.AddCommand(listCmd)

	rootCmd.AddCommand(enrollCmd)
}
