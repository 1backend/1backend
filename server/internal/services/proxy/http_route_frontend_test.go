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

	t.Run("preserves Authorization header", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			fmt.Fprintf(w, "auth: %s", auth)
		}))
		defer mockBackend.Close()

		// Save route again for new backend (or mock different one if needed)
		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/", nil)
		require.NoError(t, err)
		req.Host = "test.localhost"
		req.Header.Set("Authorization", "Bearer secret-token")

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		require.Contains(t, string(body), "Bearer secret-token")
	})

	t.Run("proxies file upload", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseMultipartForm(10 << 20) // 10MB
			if err != nil {
				http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
				return
			}
			file, _, err := r.FormFile("file")
			if err != nil {
				http.Error(w, "File not found", http.StatusBadRequest)
				return
			}
			defer file.Close()
			content, _ := io.ReadAll(file)
			w.Write(content)
		}))
		defer mockBackend.Close()

		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		var bodyBuf strings.Builder
		bodyBuf.WriteString("--boundary\r\n")
		bodyBuf.WriteString(`Content-Disposition: form-data; name="file"; filename="test.txt"` + "\r\n")
		bodyBuf.WriteString("Content-Type: text/plain\r\n\r\n")
		bodyBuf.WriteString("file contents here\r\n")
		bodyBuf.WriteString("--boundary--\r\n")

		req, err := http.NewRequest(http.MethodPost, edgeProxyUrl+"/upload", strings.NewReader(bodyBuf.String()))
		require.NoError(t, err)
		req.Host = "test.localhost"
		req.Header.Set("Content-Type", "multipart/form-data; boundary=boundary")

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		require.Equal(t, "file contents here", string(respBody))
	})

	t.Run("preserves query parameters", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "query: %s", r.URL.RawQuery)
		}))
		defer mockBackend.Close()

		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/search?q=test&sort=desc", nil)
		require.NoError(t, err)
		req.Host = "test.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		require.Contains(t, string(body), "q=test")
		require.Contains(t, string(body), "sort=desc")
	})

	t.Run("appends to X-Forwarded-For", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, r.Header.Get("X-Forwarded-For"))
		}))
		defer mockBackend.Close()

		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/", nil)
		require.NoError(t, err)
		req.Host = "test.localhost"
		req.Header.Set("X-Forwarded-For", "1.2.3.4")

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		require.Contains(t, string(body), "1.2.3.4")
		require.Contains(t, string(body), "127.0.0.1") // or localhost equivalent
	})

	t.Run("preserves X-Forwarded-Proto if set", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, r.Header.Get("X-Forwarded-Proto"))
		}))
		defer mockBackend.Close()

		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/", nil)
		require.NoError(t, err)
		req.Host = "test.localhost"
		req.Header.Set("X-Forwarded-Proto", "customproto")

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		require.Equal(t, "customproto", string(body))
	})

	t.Run("sets X-Real-IP correctly", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, r.Header.Get("X-Real-IP"))
		}))
		defer mockBackend.Close()

		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/", nil)
		require.NoError(t, err)
		req.Host = "test.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		require.NotEmpty(t, string(body))
	})

	t.Run("sets RFC 7239 Forwarded header", func(t *testing.T) {
		mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, r.Header.Get("Forwarded"))
		}))
		defer mockBackend.Close()

		routeReq.Routes[0].Target = openapi.PtrString(mockBackend.URL)
		_, _, err := adminClient.ProxySvcAPI.SaveRoutes(context.Background()).Body(routeReq).Execute()
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, edgeProxyUrl+"/", nil)
		require.NoError(t, err)
		req.Host = "test.localhost"

		resp, err := proxyClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		require.Contains(t, string(body), "for=")
		require.Contains(t, string(body), "proto=")
	})
}
