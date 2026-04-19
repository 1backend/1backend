/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/stretchr/testify/require"
)

func responsePermitPermissions(permit openapi.UserSvcPermit) []string {
	ret := make([]string, 0, len(permit.Permissions)+1)
	seen := map[string]struct{}{}

	add := func(value string) {
		if value == "" {
			return
		}
		if _, ok := seen[value]; ok {
			return
		}
		seen[value] = struct{}{}
		ret = append(ret, value)
	}

	if permit.Permission != nil {
		add(*permit.Permission)
	}
	for _, permission := range permit.Permissions {
		add(permission)
	}

	return ret
}

func countString(values []string, target string) int {
	count := 0
	for _, value := range values {
		if value == target {
			count++
		}
	}

	return count
}

func listPermissionsRequest(
	t *testing.T,
	apiClient *openapi.APIClient,
	roles []string,
) (*openapi.UserSvcListPermissionsResponse, *http.Response) {
	t.Helper()

	body, err := json.Marshal(map[string][]string{
		"roles": roles,
	})
	require.NoError(t, err)

	cfg := apiClient.GetConfig()
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		cfg.Servers[0].URL+"/user-svc/permissions",
		bytes.NewReader(body),
	)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	for key, value := range cfg.DefaultHeader {
		req.Header.Set(key, value)
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	rsp, err := httpClient.Do(req)
	require.NoError(t, err)
	defer rsp.Body.Close()

	decoded := &openapi.UserSvcListPermissionsResponse{}
	if rsp.StatusCode == http.StatusOK {
		require.NoError(t, json.NewDecoder(rsp.Body).Decode(decoded))
	}

	return decoded, rsp
}
