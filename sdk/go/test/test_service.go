package test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type Options struct {
	Name string

	Port        int
	GpuPlatform string

	Az         string
	Region     string
	LLMHost    string
	VolumeName string

	// Path of the config folder, configurable via the "OB_FOLDER" environment variable.
	// If Test is true, this value is ignored and a random temporary folder is used instead.
	ConfigPath string

	// eg. mysql, postgres
	Db string

	// Connection string eg.
	// "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	DbConnectionString string

	// Crucial for distributed features.
	// Please see the documentation for the envar OB_NODE_ID
	NodeId string

	// DbPrefix allows us to have isolated envs for different test cases
	// but still make multiple nodes in those test cases use the same
	// shard of the db.
	DbPrefix string

	SourceControlToken  string
	SecretEncryptionKey string

	// Url of the 1Backend server
	ServerUrl string

	// OB_EDGE_PROXY is used to enable the edge proxy.
	EdgeProxy bool

	EdgeProxyTestMode bool

	// Only used in tests
	EdgeProxyHttpPort int
	// Only used in tests
	EdgeProxyHttpsPort int

	VerifyContacts bool

	// Self url
	Url string

	// Test mode if true will cause the localstore to
	// save data into random temporary folders.
	Test bool

	// HomeDir is the 1Backend config/data/uploads/downloads directory.
	// For tests it's something like /tmp/1backend-2698538720/
	// For live it's /home/youruser/.1backend
	HomeDir string

	// Defaults to 5m
	TokenExpiration     time.Duration
	TokenAutoRefreshOff bool
}

type ServiceProcess struct {
	Options Options
	// Url of the service process
	Url        string
	Cmd        *exec.Cmd
	StdoutPipe io.ReadCloser
	StderrPipe io.ReadCloser
	Stdout     bytes.Buffer
	Stderr     bytes.Buffer
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	Port       int
}

func FindAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

// Start either the 1Backend server (if no `Name` is specified)
// or a microservice by executable name.
//
// It will wait for "Server started" or "Service started" outputs.
// Will time out if it won't see those outputs.
func StartService(options Options) (*ServiceProcess, error) {
	if options.Name == "" {
		// By default this launches the executable called "server"
		options.Name = "server"
	}

	var (
		port int
		err  error
	)

	if options.Url == "" {
		port, err = FindAvailablePort()
		if err != nil {
			return nil, err
		}

		options.Url = fmt.Sprintf("http://127.0.0.1:%v", port)
	}

	if port == 0 {
		port = 11337
	}

	if options.Test {
		options.DbPrefix = "t_" + sdk.Id("")
	}

	envVars := map[string]string{
		"OB_TEST":                   fmt.Sprintf("%v", options.Test),
		"OB_SELF_URL":               options.Url,
		"OB_FOLDER":                 options.ConfigPath,
		"OB_SERVER_URL":             options.ServerUrl,
		"OB_GPU_PLATFORM":           options.GpuPlatform,
		"OB_NODE_ID":                options.NodeId,
		"OB_AZ":                     options.Az,
		"OB_REGION":                 options.Region,
		"OB_LLM_HOST":               options.LLMHost,
		"OB_VOLUME_NAME":            options.VolumeName,
		"OB_DB_PREFIX":              options.DbPrefix,
		"OB_DB":                     options.Db,
		"OB_DB_CONNECTION_STRING":   options.DbConnectionString,
		"OB_ENCRYPTION_KEY":         options.SecretEncryptionKey,
		"OB_TOKEN_EXPIRATION":       fmt.Sprintf("%v", options.TokenExpiration),
		"OB_TOKEN_AUTO_REFRESH_OFF": fmt.Sprintf("%v", options.TokenAutoRefreshOff),
		"OB_EDGE_PROXY":             fmt.Sprintf("%v", options.EdgeProxy),
		"OB_EDGE_PROXY_TEST_MODE":   fmt.Sprintf("%v", options.EdgeProxyTestMode),
		"OB_EDGE_PROXY_HTTP_PORT":   fmt.Sprintf("%v", options.EdgeProxyHttpPort),
		"OB_EDGE_PROXY_HTTPS_PORT":  fmt.Sprintf("%v", options.EdgeProxyHttpsPort),
		"OB_VERIFY_CONTACTS":        fmt.Sprintf("%v", options.VerifyContacts),
	}

	for key, value := range envVars {
		if value == "" {
			if envValue, exists := os.LookupEnv(key); exists {
				envVars[key] = envValue
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, options.Name)

	cmd.Env = append(cmd.Env, os.Environ()...)
	for key, value := range envVars {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	service := &ServiceProcess{
		Options:    options,
		Cmd:        cmd,
		StdoutPipe: stdoutPipe,
		StderrPipe: stderrPipe,
		cancel:     cancel,
		Port:       port,
		Url:        options.Url,
	}

	// **Wait until first line of output appears**
	waitChan := make(chan struct{})
	started := false

	readAndSignal := func(pipe io.ReadCloser, isErr bool) {
		reader := io.TeeReader(pipe, funcWriter(func(p []byte) (n int, err error) {
			if isErr {
				return service.Stderr.Write(p)
			}
			return service.Stdout.Write(p)
		}))

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()

			if started {
				continue
			}

			if strings.Contains(line, "Server started") ||
				strings.Contains(line, "Service started") {
				started = true
				close(waitChan)
			}
		}
	}

	// Start goroutines to read both stdout and stderr
	go readAndSignal(stdoutPipe, false)
	go readAndSignal(stderrPipe, true)

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "service failed to start")
	}

	timeout := 8 * time.Second

	select {
	case <-waitChan:
	case <-time.After(timeout): // Timeout in case the service fails to start
		service.Stop()
		return nil, errors.Errorf(
			"process '%v' did not produce desired output within %v: %v",
			options.Name,
			timeout,
			service.Output(),
		)
	}

	service.wg.Add(1)
	go func() {
		defer service.wg.Done()
		cmd.Wait()
	}()

	return service, nil
}

type funcWriter func([]byte) (int, error)

func (f funcWriter) Write(p []byte) (int, error) {
	return f(p)
}

func (s *ServiceProcess) Stop() {
	s.cancel()
	time.Sleep(100 * time.Millisecond) // Give process some time to exit
	_ = s.Cmd.Process.Kill()
	s.wg.Wait()
}

func (s *ServiceProcess) Output() string {
	return s.Stdout.String() + s.Stderr.String()
}

func (s *ServiceProcess) Cleanup(t *testing.T) {
	processName := strings.ToUpper(s.Options.Name)

	if t.Failed() {
		fmt.Printf(
			"=== %v OUTPUT ===\n",
			processName,
		)
		fmt.Print(s.Stdout.String() + s.Stderr.String())
		fmt.Printf(
			"=== END OF %v OUTPUT ===\n",
			processName,
		)
	}

	s.Stop()
}

func NewSelfUrl(t *testing.T) string {
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	selfURL := fmt.Sprintf("http://localhost:%d", port)

	return selfURL
}
