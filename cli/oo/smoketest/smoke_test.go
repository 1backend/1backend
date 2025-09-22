package smoketest

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/test"
)

func TestCLISmoke(t *testing.T) {
	serverBinary, cliBinary, cliRoot := buildBinaries(t)

	server, err := test.StartService(test.Options{
		Name: serverBinary,
		Test: true,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		server.Cleanup(t)
	})

	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)
	t.Setenv("NO_COLOR", "1")

	configDir := filepath.Join(homeDir, ".1backend")
	err = os.MkdirAll(configDir, 0o755)
	require.NoError(t, err)

	testEnvId := fmt.Sprintf("env-%d", rand.Intn(99999)+1)

	runCLI(t, cliBinary, "env", "add", testEnvId, server.Url)
	runCLI(t, cliBinary, "env", "select", testEnvId)
	runCLI(t, cliBinary, "login", "1backend", "changeme")

	defer func() {
		if t.Failed() {
			fmt.Println(runCLI(t, cliBinary, "env", "ls"))
			fmt.Println(runCLI(t, cliBinary, "whoami"))
		}
	}()

	dataDir := t.TempDir()

	fixturesDir := filepath.Join(cliRoot, "fixtures")

	t.Run("check if env is set up correctly", func(t *testing.T) {
		output := runCLI(t, cliBinary, "env", "list")
		require.Contains(t, output, "ENV NAME")
		require.Contains(t, output, testEnvId)
	})

	t.Run("test routes", func(t *testing.T) {
		routeAID := fmt.Sprintf("%s.example.com", sdk.Id("route-a"))
		routeBID := fmt.Sprintf("%s.example.com", sdk.Id("route-b"))

		routeAPath := filepath.Join(dataDir, "routeA.yaml")
		err = os.WriteFile(routeAPath, []byte(fmt.Sprintf(`id: "%s"
target: "http://127.0.0.1:8080"
`, routeAID)), 0o644)
		require.NoError(t, err)

		routeBPath := filepath.Join(dataDir, "routeB.yaml")
		err = os.WriteFile(routeBPath, []byte(fmt.Sprintf(`id: "%s"
target: "http://127.0.0.1:9090"
`, routeBID)), 0o644)
		require.NoError(t, err)

		runCLI(t, cliBinary, "route", "save", routeAPath)
		runCLI(t, cliBinary, "route", "save", routeBPath)

		output := runCLI(t, cliBinary, "route", "list")
		require.Contains(t, output, routeAID)
		require.Contains(t, output, routeBID)
	})

	t.Run("test configs", func(t *testing.T) {
		configID := fmt.Sprintf("cfg-%s", sdk.Id("cfg"))

		configPath := filepath.Join(dataDir, "config.yaml")
		err = os.WriteFile(configPath, []byte(fmt.Sprintf(`appHost: "%s"
id: "%s"
data:
  feature: "enabled"
  nested:
    deepkey: "deepvalue"
`, sdk.DefaultAppHost, configID)), 0o644)
		require.NoError(t, err)

		runCLI(t, cliBinary, "config", "save", configPath)

		output := runCLI(t, cliBinary, "config", "list")
		require.Contains(t, output, kebabToCamel(configID))

	})

	t.Run("test secret", func(t *testing.T) {
		secretPath := filepath.Join(fixturesDir, "secrets.yaml")
		runCLI(t, cliBinary, "secret", "save", secretPath)

		output := runCLI(t, cliBinary, "secret", "list", "--show")
		require.Contains(t, output, "API_KEY_1")
		require.Contains(t, output, "example")
	})

	t.Run("test permits", func(t *testing.T) {
		permitPath := filepath.Join(fixturesDir, "permitA.yaml")
		runCLI(t, cliBinary, "permit", "save", permitPath)

		output := runCLI(t, cliBinary, "permit", "list")
		require.Contains(t, output, "permit-id-1")
		require.Contains(t, output, "user-svc:role:read")
	})
}

func buildBinaries(t *testing.T) (string, string, string) {
	t.Helper()

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to determine test file path")
	}

	cliRoot := filepath.Clean(filepath.Join(filepath.Dir(thisFile), ".."))
	repoRoot := filepath.Clean(filepath.Join(cliRoot, "..", ".."))

	binDir := t.TempDir()

	serverBinary := filepath.Join(binDir, "server")
	buildCmd(t, repoRoot, "./server", serverBinary)

	cliBinary := filepath.Join(binDir, "oo")
	buildCmd(t, repoRoot, "./cli/oo", cliBinary)

	return serverBinary, cliBinary, cliRoot
}

func buildCmd(t *testing.T, repoRoot, pkg, output string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "build", "-o", output, pkg)
	cmd.Dir = repoRoot
	cmd.Env = os.Environ()

	out, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("building %s timed out", pkg)
	}
	require.NoErrorf(t, err, "failed to build %s: %s", pkg, string(out))
}

func runCLI(t *testing.T, cliBinary string, args ...string) string {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, cliBinary, args...)
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("command timed out: %s", strings.Join(args, " "))
	}
	require.NoErrorf(t, err, "command failed: %s\n%s", strings.Join(args, " "), string(output))

	return string(output)
}

func kebabToCamel(s string) string {
	parts := strings.Split(s, "-")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
