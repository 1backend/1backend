package user

import (
	"github.com/spf13/cobra"
)

func AddUserCommands(rootCmd *cobra.Command) {
	addLoginCommand(rootCmd)
	addRegisterCommand(rootCmd)
	addWhoamiCommand(rootCmd)
	addTokenCommand(rootCmd)
	addUseCommand(rootCmd)

	var usersCmd = &cobra.Command{
		Use:     "user",
		Aliases: []string{"u", "users"},
		Short:   "Manage users",
	}
	rootCmd.AddCommand(usersCmd)

	addListUsersCommand(usersCmd)
}

func addLoginCommand(rootCmd *cobra.Command) {
	var loginCmd = &cobra.Command{
		Use:   "login [slug] [password]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Log in to the currently selected environment.",
		Long: `The 'login' command allows you to authenticate a user in the currently selected environment.

If no arguments are provided, the command will prompt for the username (slug) and password securely.
If only the username (slug) is provided, the command will securely prompt for the password.
If both the username and password are provided as arguments, the login will proceed, but keep in mind 
that the password will be visible in the terminal command history.`,
		Example: `  # Login with slug and prompt for password
  login johnny

  # Login by providing both slug and password (not secure, avoid this)
  login johnny myPass1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Login(cmd, args)
		},
	}

	rootCmd.AddCommand(loginCmd)
}

func addRegisterCommand(rootCmd *cobra.Command) {
	var (
		contactId       string
		contactPlatform string
	)

	var registerCmd = &cobra.Command{
		Use:   "register [slug] [password]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Register a new user in the currently selected environment.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Register(cmd, args)
		},
	}

	registerCmd.Flags().
		StringVarP(&contactId, "contact-id", "i", "", "Contact ID to register")

	registerCmd.Flags().
		StringVarP(&contactPlatform, "contact-platform", "p", "", "Contact platform to register")

	rootCmd.AddCommand(registerCmd)
}

func addWhoamiCommand(rootCmd *cobra.Command) {
	all := false

	var runCmd = &cobra.Command{
		Use:   "whoami",
		Args:  cobra.ExactArgs(0),
		Short: "Display the currently logged in user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Whoami(cmd, args, all)
		},
	}

	runCmd.Flags().BoolVarP(&all, "all", "a", false, "Display all logged in user information, not just the currently selected one")

	rootCmd.AddCommand(runCmd)
}

func addTokenCommand(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
		Use: "token",
		// Args:  cobra.ExactArgs(0),
		Short: "Display the token of the user currently logged in",
		RunE:  Token,
	}

	rootCmd.AddCommand(runCmd)
}

func addUseCommand(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
		Use:   "su [slug]",
		Args:  cobra.ExactArgs(1),
		Short: "Switch to user by specifying their slug.",
		RunE:  Su,
	}

	rootCmd.AddCommand(runCmd)
}

func addListUsersCommand(rootCmd *cobra.Command) {
	var (
		userId    string
		contactId string
		limit     int64
	)

	var runCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		Short:   "List users.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListUsers(cmd, args, userId, contactId, limit)
		},
	}

	runCmd.Flags().StringVarP(&userId, "userId", "u", "", "Filter by user Id")
	runCmd.Flags().StringVarP(&contactId, "contactId", "c", "", "Filter by contact Id")
	runCmd.Flags().Int64VarP(&limit, "limit", "l", 0, "Limit result set")

	rootCmd.AddCommand(runCmd)
}
