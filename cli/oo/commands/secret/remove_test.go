package secret

import (
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

// TestRemoveCommandOnlyAdvertisesIDs guards against a `--key` flag being
// offered again: /secret-svc/secrets removal accepts ids only, so a key
// flag would be accepted and then silently ignored.
func TestRemoveCommandOnlyAdvertisesIDs(t *testing.T) {
	root := &cobra.Command{Use: "oo"}
	AddSecretCommands(root)

	removeCmd, _, err := root.Find([]string{"secret", "remove"})
	require.NoError(t, err)
	require.NotNil(t, removeCmd)

	require.NotNil(t, removeCmd.Flags().Lookup("id"))
	require.Nil(t, removeCmd.Flags().Lookup("key"))
}

func TestRemoveWithoutIDsReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.SetContext(context.Background())

	err := Remove(cmd, nil)
	require.EqualError(t, err, "at least one --id must be specified")
}
