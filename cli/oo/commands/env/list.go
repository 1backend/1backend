package env

import (
	"fmt"
	"net"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/config"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "ENV NAME\tSELECTED\tURL\tDESCRIPTION\tREACHABLE")

	for name, env := range conf.Environments {
		selected := ""
		if conf.SelectedEnvironment == name {
			selected = "*"
		}

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%v\n",
			name,
			selected,
			env.URL,
			env.Description,
			scanPort(env.URL),
		)
	}

	return nil
}

func scanPort(url string) bool {
	port := 80
	if strings.Contains(url, "https") {
		port = 443
	}

	url = strings.Replace(strings.Replace(url, "https://", "", 1), "http://", "", 1)

	if !strings.Contains(url, ":") {
		url = fmt.Sprintf("%v:%v", url, port)
	}

	conn, err := net.Dial("tcp", url)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
