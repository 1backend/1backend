package cert

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/1backend/1backend/cli/oo/config"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	shortTimeFormat = "2006-01-02"
	timeFormat      = "2006-01-02 15:04:05"
)

// List
func List(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	cf := client.NewApiClientFactory(url)

	req := openapi.ProxySvcListCertsRequest{}
	rsp, _, err := cf.Client(client.WithToken(token)).
		ProxySvcAPI.ListCerts(ctx).
		Body(req).
		Execute()
	if err != nil {
		return errors.Wrap(err, "failed to list certs")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"CERT ID\tCREATED AT\tUPDATED AT\tCOMMON NAME\tISSUER\tNOT BEFORE\tNOT AFTER\tSERIAL",
	)

	for _, cert := range rsp.Certs {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			cert.Id,
			date(cert.CreatedAt, shortTimeFormat),
			date(cert.UpdatedAt, timeFormat),
			safeStr(cert.CommonName),
			safeStr(cert.Issuer),
			safeDate(cert.NotBefore, shortTimeFormat),
			safeDate(cert.NotAfter, shortTimeFormat),
			truncateSerial(safeStr(cert.SerialNumber), 8),
		)
	}

	return nil
}

func safeStr(s *string) string {
	if s == nil {
		return "-"
	}
	return *s
}

func date(s string, format string) string {
	if s == "" {
		return ""
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return s // fallback to raw string if parsing fails
	}
	d := t.Format(format) // or use any layout you prefer
	if d == "0001-01-01" {
		return ""
	}
	return d
}

func safeDate(s *string, format string) string {
	if s == nil {
		return ""
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return *s // fallback to raw string if parsing fails
	}
	d := t.Format(format) // or use any layout you prefer
	if d == "0001-01-01" {
		return ""
	}
	return d
}

func truncateSerial(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
