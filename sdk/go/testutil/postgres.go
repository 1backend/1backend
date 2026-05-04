package testutil

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultDBName   = "mydatabase"
	defaultUser     = "postgres"
	defaultPassword = "mysecretpassword"
)

func StartPostgres(t *testing.T) string {
	t.Helper()

	ctx := context.Background()
	if !IsDockerAvailable(ctx) {
		t.Skip("docker is required for postgres integration tests")
	}

	lockFile, err := os.OpenFile(
		filepath.Join(os.TempDir(), "1backend-postgres-test.lock"),
		os.O_CREATE|os.O_RDWR,
		0600,
	)
	require.NoError(t, err)
	require.NoError(t, syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX))
	t.Cleanup(func() {
		require.NoError(t, syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN))
		require.NoError(t, lockFile.Close())
	})

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(defaultDBName),
		postgres.WithUsername(defaultUser),
		postgres.WithPassword(defaultPassword),
		tc.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(90*time.Second),
		),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, container.Terminate(ctx))
	})

	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "5432/tcp")
	require.NoError(t, err)

	ready := false
	for i := 0; i < 30; i++ {
		exitCode, _, err := container.Exec(ctx, []string{"pg_isready", "-U", defaultUser, "-d", defaultDBName})
		if err == nil && exitCode == 0 {
			ready = true
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	require.True(t, ready, "postgres container did not become ready")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", defaultUser, defaultPassword, host, port.Port(), defaultDBName)
}

func IsDockerAvailable(ctx context.Context) bool {
	_, err := tc.NewDockerClientWithOpts(ctx)
	return err == nil
}
