package proxyservice_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"

	openapi "github.com/1backend/1backend/clients/go"
)

// @todo Here we should test more things, like file uplods etc.
// The edge proxy should be fairly simple, transparent and reliable.

func TestProxyService_FrontendRoute(t *testing.T) {
	t.Parallel()

	mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/hello" {
			fmt.Fprintf(w, "Hello from backend!")
		} else if r.URL.Path == "/echo" {
			io.Copy(w, r.Body)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer mockBackend.Close()

	port, err := test.FindAvailablePort()
	require.NoError(t, err)

	server, err := test.StartService(test.Options{
		EdgeProxy:          true,
		EdgeProxyTestMode:  true,
		EdgeProxyHttpsPort: port,
		Test:               true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)
	adminClient, _, err := test.AdminClient(clientFactory)
	require.NoError(t, err)

	edgeProxyUrl := fmt.Sprintf("http://localhost:%d", port)

	routeReq := openapi.ProxySvcSaveRoutesRequest{
		Routes: []openapi.ProxySvcRouteInput{
			{
				Id:     openapi.PtrString("test.localhost"),
				Target: openapi.PtrString(mockBackend.URL),
			},
		},
	}

	t.Run("save route", func(t *testing.T) {
		_, hrsp, err := adminClient.ProxySvcAPI.SaveRoutes(
			context.Background(),
		).Body(routeReq).Execute()

		require.NoError(t, err, hrsp)
	})

	proxyClient := &http.Client{}

	t.Run("proxies GET request to /hello", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/hello", edgeProxyUrl), nil)
		require.NoError(t, err)
		req.Host = "test.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, "Hello from backend!", string(body))
	})

	t.Run("proxies POST request to /echo", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/echo", edgeProxyUrl), strings.NewReader("echo me"))
		require.NoError(t, err)
		req.Host = "test.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, "echo me", string(body))
	})

	t.Run("returns 404 if no matching route", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/hello", edgeProxyUrl), nil)
		require.NoError(t, err)
		req.Host = "unknown.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
