package env

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/openorch/openorch/cli/oo/config"
	"github.com/spf13/cobra"
)

// Current
func Current(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if len(conf.Environments) == 0 {
		fmt.Println("No environments found.")
		return nil
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "SELECTED\tUSERS\tURL\tDESCRIPTION")

	env := conf.Environments[conf.SelectedEnvironment]
	selected := "*"

	fmt.Fprintf(
		writer,
		"%s\t%s\t%s\t%s\t\n",
		selected,
		env.ShortName,
		env.URL,
		env.Description,
	)

	return nil
}
