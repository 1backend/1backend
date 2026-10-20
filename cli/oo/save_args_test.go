package main

import (
	"context"
	"io"
	"testing"

	"github.com/1backend/1backend/cli/oo/commands/enroll"
	"github.com/1backend/1backend/cli/oo/commands/permit"
	"github.com/1backend/1backend/cli/oo/commands/route"
	"github.com/1backend/1backend/cli/oo/config"
	types "github.com/1backend/1backend/cli/oo/types"
	"github.com/1backend/1backend/cli/oo/util"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

// selectTestUser writes a CLI config with a selected user and token so that
// commands get past the environment/token lookup and reach their argument
// handling. No network call is made by the tests below.
func selectTestUser(t *testing.T) {
	t.Helper()

	t.Setenv("HOME", t.TempDir())

	require.NoError(t, util.SaveConfig(types.Config{
		SelectedEnvironment: "test",
		Environments: map[string]*types.Environment{
			"test": {
				ShortName:    "test",
				URL:          "http://127.0.0.1:1",
				SelectedUser: "tester",
				Users: map[string]*types.User{
					"tester": {
						Slug:            "tester",
						SelectedAppHost: "test.app",
						TokensByAppHost: map[string]string{
							"test.app": "test-token",
						},
					},
				},
			},
		},
	}))
}

func executeRoot(t *testing.T, add func(*cobra.Command), args []string) error {
	t.Helper()

	root := &cobra.Command{Use: "oo"}
	add(root)
	root.SetArgs(args)
	root.SetOut(io.Discard)
	root.SetErr(io.Discard)
	root.SetContext(context.Background())

	var err error
	require.NotPanics(t, func() {
		err = root.Execute()
	})

	return err
}

// TestSaveCommandsRejectMissingPath verifies that the file based save
// commands report a usage error instead of panicking on args[0] when no
// path argument is provided.
func TestSaveCommandsRejectMissingPath(t *testing.T) {
	cases := []struct {
		name string
		add  func(*cobra.Command)
		args []string
	}{
		{
			name: "config",
			add:  config.AddConfigCommands,
			args: []string{"config", "save"},
		},
		{
			name: "permit",
			add:  permit.AddPermitCommands,
			args: []string{"permit", "save"},
		},
		{
			name: "enroll",
			add:  enroll.AddEnrollCommands,
			args: []string{"enroll", "save"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			selectTestUser(t)

			err := executeRoot(t, tc.add, tc.args)
			require.Error(t, err)
			require.Contains(t, err.Error(), "received 0")
		})
	}
}

// TestRouteSaveRejectsMissingPathAndFlags covers the flag based command:
// `route save` accepts no positional argument when --id and --target are
// set, so the missing path has to be handled in the command itself.
func TestRouteSaveRejectsMissingPathAndFlags(t *testing.T) {
	selectTestUser(t)

	err := executeRoot(t, route.AddRouteCommands, []string{"route", "save"})
	require.Error(t, err)
	require.Contains(t, err.Error(), "route path required")
}
