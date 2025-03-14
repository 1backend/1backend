package userservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/openorch/openorch/clients/go"
	"github.com/openorch/openorch/server/internal/di"
)

func TestRegister(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)

	err = starterFunc()
	require.NoError(t, err)

	ctx := context.Background()

	tokenId := ""

	t.Run("anyone can register", func(t *testing.T) {
		rsp, _, err := options.ClientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				Slug:     openapi.PtrString("test-slug-1"),
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)

		tokenId = *rsp.Token.Id
	})

	t.Run("cannot re-register", func(t *testing.T) {
		_, _, err := options.ClientFactory.Client().UserSvcAPI.Register(ctx).Body(
			openapi.UserSvcRegisterRequest{
				Slug:     openapi.PtrString("test-slug-1"),
				Name:     openapi.PtrString("Test Name"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.Error(t, err)
	})

	t.Run("login doesn't produce a new token", func(t *testing.T) {
		rsp, _, err := options.ClientFactory.Client().UserSvcAPI.Login(ctx).Body(
			openapi.UserSvcLoginRequest{
				Slug:     openapi.PtrString("test-slug-1"),
				Password: openapi.PtrString("testPass123"),
			},
		).Execute()
		require.NoError(t, err)
		require.Equal(t, true, *rsp.Token.Id == tokenId)
	})
}
