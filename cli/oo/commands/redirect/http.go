package redirect

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func doJSON(ctx context.Context, baseURL, token, method, path string, reqBody any, rspBody any) error {
	var body io.Reader
	if reqBody != nil {
		bs, err := json.Marshal(reqBody)
		if err != nil {
			return errors.Wrap(err, "marshal request")
		}
		body = bytes.NewReader(bs)
	}

	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(baseURL, "/")+path, body)
	if err != nil {
		return errors.Wrap(err, "build request")
	}
	req.Header.Set("Authorization", "Bearer "+token)
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request")
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.Errorf("%s: %s", resp.Status, string(respBytes))
	}

	if rspBody != nil && len(respBytes) > 0 {
		if err := json.Unmarshal(respBytes, rspBody); err != nil {
			return errors.Wrap(err, "decode response")
		}
	}

	return nil
}
