package proxyservice_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestProxyService_Route(t *testing.T) {
	t.Parallel()

	// Step 1: Spin up a mock target service that would be proxied to
	mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Mock-Backend", "true")
		if r.URL.Path == "/test-user-slug-0/error" {
			http.Error(w, "backend error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		if r.URL.Path == "/test-user-slug-0/endpoint" {
			w.Write([]byte(`{"success": true}`))
		} else if r.URL.Path == "/test-user-slug-0/echo-params" {
			query := r.URL.Query()
			response := fmt.Sprintf(`{"success": true, "query": "%s"}`, query.Encode())
			w.Write([]byte(response))
		} else if r.URL.Path == "/test-user-slug-0/echo-headers" {
			headers := r.Header
			response := fmt.Sprintf(`{"success": true, "headers": "%s"}`, headers)
			w.Write([]byte(response))
		} else {
			fmt.Fprintf(w, "proxied %s", r.URL.Path)
		}
	}))
	defer mockBackend.Close()

	// Step 2: Start the full service under test (including proxy)
	server, err := test.StartService(test.Options{
		TokenExpiration: 1 * time.Second,
		Test:            true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	client := client.NewApiClientFactory(server.Url)
	manyClients, _, err := test.MakeClients(client, 1)
	require.NoError(t, err)

	client1 := manyClients[0]

	_, _, err = client1.RegistrySvcAPI.RegisterInstance(
		context.Background(),
	).Body(
		openapi.RegistrySvcRegisterInstanceRequest{
			Url: mockBackend.URL,
		}).Execute()
	require.NoError(t, err)

	// Step 3: Create a client pointing to the proxy server
	proxyClient := &http.Client{}

	t.Run("should proxy GET request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/test-user-slug-0/some-endpoint", server.Url), nil)
		require.NoError(t, err)

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, "true", resp.Header.Get("X-Mock-Backend"))
	})

	t.Run("response can be read", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/test-user-slug-0/endpoint", server.Url), nil)
		require.NoError(t, err)
		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		responseBody, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, `{"success": true}`, string(responseBody))

		t.Run("proxying works after token expiration", func(t *testing.T) {
			time.Sleep(1 * time.Second)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/test-user-slug-0/endpoint", server.Url), nil)
			require.NoError(t, err)
			resp, err := proxyClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			require.Equal(t, http.StatusOK, resp.StatusCode)

			responseBody, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, `{"success": true}`, string(responseBody))
		})
	})

	t.Run("should return 500 on backend error", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/test-user-slug-0/error", server.Url), nil)
		require.NoError(t, err)

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("should return 404 if no instance found", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/unknown-svc/endpoint", server.Url), nil)
		require.NoError(t, err)

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("should handle OPTIONS requests gracefully", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodOptions, fmt.Sprintf("%s/test-user-slug-0/any", server.Url), nil)
		require.NoError(t, err)

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("should preserve query parameters", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/test-user-slug-0/echo-params?foo=bar&baz=qux", server.Url), nil)
		require.NoError(t, err)

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Contains(t, string(body), "foo=bar")
		require.Contains(t, string(body), "baz=qux")
	})

	t.Run("should preserve headers", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/test-user-slug-0/echo-headers", server.Url), nil)
		require.NoError(t, err)

		req.Header.Set("X-Custom-Header", "hello-world")
		req.Header.Set("X-Multi-Value", "one")
		req.Header.Add("X-Multi-Value", "two")

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), "X-Custom-Header")
		require.Contains(t, string(body), "hello-world")
		require.Contains(t, string(body), "X-Multi-Value")
		require.Contains(t, string(body), "one")
		require.Contains(t, string(body), "two")
	})
}
