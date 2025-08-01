package call

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/1backend/1backend/cli/oo/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Get(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) < 1 {
		return fmt.Errorf(
			"path is missing",
		)
	}

	uri, token, err := util.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get environment URL and token")
	}

	fullUrl := fmt.Sprintf(
		"%s%s",
		uri,
		strings.Join(args, "/"),
	)

	queryParams := make(map[string]string)

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// Convert the flag name and value to query parameters
		// For example, --key=value becomes ?key=value in the URL
		queryParams[flag.Name] = flag.Value.String()
	})

	var queryStrings []string
	for key, value := range queryParams {
		encodedValue := url.QueryEscape(value)
		queryStrings = append(
			queryStrings,
			fmt.Sprintf("%s=%s", key, encodedValue),
		)
	}

	if len(queryStrings) > 0 {
		fullUrl = fmt.Sprintf("%s?%s", fullUrl, strings.Join(queryStrings, "&"))
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fullUrl,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "failed to create HTTP request")
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to execute HTTP request")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf(
			"HTTP request failed with status %d: %s",
			resp.StatusCode,
			string(respBody),
		)
	}

	prettyJSON, err := util.PrettyJSON(respBody)
	if err != nil {
		return errors.Wrap(err, "failed to prettify JSON")
	}

	fmt.Println(string(prettyJSON))

	return nil
}
