package secret

import "github.com/spf13/cobra"

func AddSecretCommands(rootCmd *cobra.Command) {
	var secretCmd = &cobra.Command{
		Use:     "secret",
		Aliases: []string{"s", "secrets"},
		Short:   "Manage secrets",
	}

	var fromEnv string
	var toEnv string

	var copyCmd = &cobra.Command{
		Use: "copy",
		Short: `Copy all secrets from one environment to another.
Uses the selected app context for both 'from' and 'to' environments.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Copy(cmd, fromEnv, toEnv)
		},
	}

	var readCmd = &cobra.Command{
		Use:   "read [id]",
		Short: "Read secrets and print them as YAML",
		Args:  cobra.MaximumNArgs(1),
		RunE:  Read,
	}

	var saveCmd = &cobra.Command{
		Use:     "save [key] [value] | [filePath]",
		Aliases: []string{"s"},
		Args:    cobra.MaximumNArgs(2),
		Short:   "Save secrets as key-value pairs or from a file",
		Long: `The 'save' command allows you to save secrets in two ways:
	
	1. A single key-value pair directly.
	2. A single secret from a yaml.
	3. Multiple secrets from a yaml.`,
		Example: `# Save a single key-value pair
save API_KEY 123456789

# Save a secret from a file
save ./secretA.yaml

# Example contents of 'secretA.yaml':
id: "example-id-1"
key: "API_KEY"
value: "a37/KUAr4SOYi6Xw9i9T8qo3QCk8WvnzONo47jHAkwk="
encrypted: true
readers:
  - "service1"
  - "user1"
writers:
  - "service2"
  - "user2"

# Save multiple secrets from a file
save ./secrets.yaml

# Example contents of 'secrets.yaml'
- id: "example-id-1"
  key: "API_KEY_1"
  value: "a37/KUAr4SOYi6Xw9i9T8qo3QCk8WvnzONo47jHAkwk="
  encrypted: true
  readers:
    - "service1"
    - "user1"
  writers:
    - "service2"
    - "user2"
- id: "example-id-2"
  key: "API_KEY_2"
  value: "example2"
  encrypted: false
  readers:
    - "service3"
    - "user3"
  writers:
    - "service4"
    - "user4"`,
		RunE: Save,
	}

	var encryptCmd = &cobra.Command{
		Use:     "encrypt [key] [value]",
		Aliases: []string{"e"},
		Args:    cobra.MinimumNArgs(1),
		Short:   "Encrypt a key-value secret",
		Long: `Encrypt a key-value secret and generate a YAML structure with the encrypted value.
The resulting YAML can be safely committed to source control.

If no value is provided, you will be securely prompted to enter it.
This prevents sensitive data from being exposed in terminal history.`,

		Example: `# Encrypt with key and value (for automation)
encrypt api-key my-api-key-1234

# Interactive encryption (value will not appear in terminal)
encrypt api-key
Enter secret value:`,

		RunE: Encrypt,
	}

	var removeCmd = &cobra.Command{
		Use:     "remove [id]",
		Short:   "Delete a secret",
		Aliases: []string{"rm"},
		RunE:    Remove,
	}

	removeCmd.Flags().StringArrayP("key", "k", []string{}, "Keys of secrets to remove")
	removeCmd.Flags().StringArrayP("id", "i", []string{}, "IDs of secrets to remove")

	var (
		show    bool
		allApps bool
		verbose bool
	)

	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List secrets",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return List(cmd, args, show, allApps, verbose)
		},
	}

	var isSecureCmd = &cobra.Command{
		Use:     "is-secure",
		Short:   "Tells if the secret service is secure (ie. encryption key is set, etc.)",
		Aliases: []string{"is"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return IsSecure(cmd, args)
		},
	}

	listCmd.Flags().
		BoolVar(&show, "show", false, "Show secrets unmasked")

	listCmd.Flags().
		BoolVarP(&allApps, "all-apps", "a", false, "List secrets across all apps. If false, the app from the authentication context is used.")

	listCmd.Flags().
		BoolVarP(&verbose, "verbose", "v", false, "Verbose output (yaml array)")

	secretCmd.AddCommand(readCmd)
	secretCmd.AddCommand(saveCmd)
	secretCmd.AddCommand(removeCmd)
	secretCmd.AddCommand(listCmd)
	secretCmd.AddCommand(encryptCmd)
	secretCmd.AddCommand(isSecureCmd)
	secretCmd.AddCommand(copyCmd)

	copyCmd.Flags().
		StringVarP(&fromEnv, "from", "f", "", "Source environment short name")

	copyCmd.Flags().
		StringVarP(&toEnv, "to", "t", "", "Destination environment short name")

	rootCmd.AddCommand(secretCmd)
}
