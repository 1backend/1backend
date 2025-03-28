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
)

type Options struct {
	Port        int
	GpuPlatform string

	Az         string
	Region     string
	LLMHost    string
	VolumeName string
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

	// URL of the local 1Backend server instance
	Url string

	// Test mode if true will cause the localstore to
	// save data into random temporary folders.
	Test bool

	// HomeDir is the 1Backend config/data/uploads/downloads directory.
	// For tests it's something like /tmp/1backend-2698538720/
	// For live it's /home/youruser/.1backend
	HomeDir string
}

type ServerProcess struct {
	Cmd        *exec.Cmd
	Url        string
	StdoutPipe io.ReadCloser
	StderrPipe io.ReadCloser
	Stdout     bytes.Buffer
	Stderr     bytes.Buffer
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	Port       string
}

func findAvailablePort() (string, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	defer listener.Close()
	return fmt.Sprintf("%v", listener.Addr().(*net.TCPAddr).Port), nil
}

func StartServer(options Options) (*ServerProcess, error) {
	var (
		port string
		err  error
	)

	if options.Url == "" {
		port, err = findAvailablePort()
		if err != nil {
			return nil, err
		}

		options.Url = fmt.Sprintf("http://127.0.0.1:%v", port)
	}

	if port == "" {
		port = "58231"
	}

	if options.Test {
		options.DbPrefix = sdk.Id("t")
	}

	envVars := map[string]string{
		"OB_GPU_PLATFORM":         options.GpuPlatform,
		"OB_SERVER_URL":           options.Url,
		"OB_NODE_ID":              options.NodeId,
		"OB_AZ":                   options.Az,
		"OB_REGION":               options.Region,
		"OB_LLM_HOST":             options.LLMHost,
		"OB_VOLUME_NAME":          options.VolumeName,
		"OB_DB_PREFIX":            options.DbPrefix,
		"OB_DB":                   options.Db,
		"OB_DB_CONNECTION_STRING": options.DbConnectionString,
		"OB_ENCRYPTION_KEY":       options.SecretEncryptionKey,
	}

	for key, value := range envVars {
		if value == "" {
			if envValue, exists := os.LookupEnv(key); exists {
				envVars[key] = envValue
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "server")

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

	server := &ServerProcess{
		Cmd:        cmd,
		StdoutPipe: stdoutPipe,
		StderrPipe: stderrPipe,
		cancel:     cancel,
		Port:       port,
		Url:        fmt.Sprintf("http://127.0.0.1:%v", port),
	}

	// **Wait until first line of output appears**
	waitChan := make(chan struct{})
	started := false

	readAndSignal := func(pipe io.ReadCloser, err bool) {
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			line := scanner.Bytes()
			if err {
				server.Stderr.Write(line)
			} else {
				server.Stdout.Write(line)
			}

			if started {
				continue
			}

			if strings.Contains(string(line), "Server started") {
				started = true
				close(waitChan) // Signal that the server is up
				return
			}
		}
	}

	// Start goroutines to read both stdout and stderr
	go readAndSignal(stdoutPipe, false)
	go readAndSignal(stderrPipe, true)

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "server failed to start")
	}

	timeout := 10 * time.Second

	select {
	case <-waitChan:
	case <-time.After(timeout): // Timeout in case the server fails to start
		server.Stop()
		return nil, errors.Errorf(
			"server did not produce output within %v", timeout,
		)
	}

	server.wg.Add(1)
	go func() {
		defer server.wg.Done()
		cmd.Wait()
	}()

	return server, nil
}

func (s *ServerProcess) Stop() {
	s.cancel()
	time.Sleep(100 * time.Millisecond) // Give process some time to exit
	_ = s.Cmd.Process.Kill()
	s.wg.Wait()
}

func (s *ServerProcess) Output() string {
	return s.Stdout.String() + s.Stderr.String()
}

func (s *ServerProcess) Cleanup(t *testing.T) {
	if t.Failed() {
		fmt.Println("=== SERVER OUTPUT ===")
		fmt.Print(s.Stdout.String() + s.Stderr.String())
		fmt.Println("=== END OF SERVER OUTPUT ===")
	}

	s.Stop()
}
