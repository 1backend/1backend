package call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/openorch/openorch/cli/oo/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	payload := make(map[string]interface{})
	for _, arg := range args {
		if err := parseFlagToMap(payload, arg); err != nil {
			return err
		}
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal payload to JSON")
	}

	url, token, err := config.GetSelectedUrlAndToken()
	if err != nil {
		return errors.Wrap(err, "cannot get environment URL and token")
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonData),
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

	if resp.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf(
			"HTTP request failed with status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	fmt.Println("Request successful")
	return nil
}
