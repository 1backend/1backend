package proxyservice_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

func TestRedirectsAPI(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	_, adminToken, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	userClients, userTokens, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	require.NotNil(t, userClients)

	redirects := proxy.SaveRedirectsRequest{
		Redirects: []proxy.RedirectInput{
			{
				Id:         "old.example.com",
				Target:     "https://new.example.com",
				StatusCode: http.StatusMovedPermanently,
			},
			{
				Id:     "old.example.com/docs",
				Target: "https://new.example.com/help",
			},
		},
	}

	statusCode, _, body := proxyJSON(t, http.MethodPut, server.Url+"/proxy-svc/redirects", adminToken, redirects)
	require.Equal(t, http.StatusOK, statusCode, string(body))

	var saveRsp proxy.SaveRedirectsResponse
	require.NoError(t, json.Unmarshal(body, &saveRsp))
	require.Len(t, saveRsp.Redirects, 2)
	require.Equal(t, http.StatusMovedPermanently, saveRsp.Redirects[0].StatusCode)
	require.Equal(t, http.StatusPermanentRedirect, saveRsp.Redirects[1].StatusCode)

	statusCode, _, _ = proxyJSON(t, http.MethodPost, server.Url+"/proxy-svc/redirects", userTokens[0].Token, nil)
	require.Equal(t, http.StatusUnauthorized, statusCode)

	statusCode, _, body = proxyJSON(t, http.MethodPost, server.Url+"/proxy-svc/redirects", adminToken, proxy.ListRedirectsRequest{})
	require.Equal(t, http.StatusOK, statusCode, string(body))

	var listRsp proxy.ListRedirectsResponse
	require.NoError(t, json.Unmarshal(body, &listRsp))
	require.Len(t, listRsp.Redirects, 2)

	statusCode, _, body = proxyJSON(t, http.MethodPost, server.Url+"/proxy-svc/redirects", adminToken, proxy.ListRedirectsRequest{
		Ids: []string{"old.example.com/docs"},
	})
	require.Equal(t, http.StatusOK, statusCode, string(body))
	require.NoError(t, json.Unmarshal(body, &listRsp))
	require.Len(t, listRsp.Redirects, 1)
	require.Equal(t, "old.example.com/docs", listRsp.Redirects[0].Id)

	statusCode, _, body = proxyJSON(t, http.MethodDelete, server.Url+"/proxy-svc/redirects", adminToken, proxy.DeleteRedirectsRequest{
		Ids: []string{"old.example.com"},
	})
	require.Equal(t, http.StatusOK, statusCode, string(body))

	statusCode, _, body = proxyJSON(t, http.MethodPost, server.Url+"/proxy-svc/redirects", adminToken, proxy.ListRedirectsRequest{})
	require.Equal(t, http.StatusOK, statusCode, string(body))
	require.NoError(t, json.Unmarshal(body, &listRsp))
	require.Len(t, listRsp.Redirects, 1)
	require.Equal(t, "old.example.com/docs", listRsp.Redirects[0].Id)
}

func TestProxyService_FrontendRedirect(t *testing.T) {
	t.Parallel()

	var backendHits atomic.Int32
	mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backendHits.Add(1)
		fmt.Fprint(w, "backend")
	}))
	defer mockBackend.Close()

	port, err := test.FindAvailablePort()
	require.NoError(t, err)

	server, err := test.StartService(test.Options{
		EdgeProxyTestMode: true,
		EdgeProxyHttpPort: port,
		Test:              true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	adminClient, adminToken, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(openapi.ProxySvcSaveRoutesRequest{
		Routes: []openapi.ProxySvcRouteInput{
			{
				Id:     "old.localhost",
				Target: openapi.PtrString(mockBackend.URL),
			},
		},
	}).Execute()
	require.NoError(t, err)

	statusCode, _, body := proxyJSON(t, http.MethodPut, server.Url+"/proxy-svc/redirects", adminToken, proxy.SaveRedirectsRequest{
		Redirects: []proxy.RedirectInput{
			{
				Id:         "old.localhost",
				Target:     "https://new.localhost",
				StatusCode: http.StatusMovedPermanently,
			},
			{
				Id:     "old.localhost/old",
				Target: "https://new.localhost/new",
			},
		},
	})
	require.Equal(t, http.StatusOK, statusCode, string(body))

	edgeProxyUrl := fmt.Sprintf("http://localhost:%d", port)
	proxyClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	t.Run("host redirect preserves path and query", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/docs/page?a=1", nil)
		require.NoError(t, err)
		req.Host = "old.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
		require.Equal(t, "https://new.localhost/docs/page?a=1", resp.Header.Get("Location"))
	})

	t.Run("path redirect replaces matched prefix", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/old/docs?q=1", nil)
		require.NoError(t, err)
		req.Host = "old.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusPermanentRedirect, resp.StatusCode)
		require.Equal(t, "https://new.localhost/new/docs?q=1", resp.Header.Get("Location"))
	})

	require.Equal(t, int32(0), backendHits.Load())
}

func proxyJSON(t *testing.T, method, url, token string, payload any) (int, http.Header, []byte) {
	t.Helper()

	var body io.Reader
	if payload != nil {
		bs, err := json.Marshal(payload)
		require.NoError(t, err)
		body = bytes.NewReader(bs)
	}

	req, err := http.NewRequest(method, url, body)
	require.NoError(t, err)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp.StatusCode, resp.Header.Clone(), respBody
}
