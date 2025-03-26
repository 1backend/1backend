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
}

func addLoginCommand(rootCmd *cobra.Command) {
	var runCmd = &cobra.Command{
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
		RunE: Login,
	}

	rootCmd.AddCommand(runCmd)
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
		Short: "Display the user currently logged in",
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
		Use:   "use [slug]",
		Args:  cobra.ExactArgs(1),
		Short: "Switch user by specifying their slug.",
		RunE:  Use,
	}

	rootCmd.AddCommand(runCmd)
}
