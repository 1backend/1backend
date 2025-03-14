package policyservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/openorch/openorch/clients/go"
	"github.com/openorch/openorch/server/internal/di"
	policytypes "github.com/openorch/openorch/server/internal/services/policy/types"
)

func TestRateLimiting(t *testing.T) {
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

	client := openapi.NewAPIClient(&openapi.Configuration{
		Servers: openapi.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
	})

	adminLoginRsp, _, err := client.UserSvcAPI.Login(context.Background()).
		Body(openapi.UserSvcLoginRequest{
			Slug:     openapi.PtrString("openorch"),
			Password: openapi.PtrString("changeme"),
		}).
		Execute()
	require.NoError(t, err)

	client = openapi.NewAPIClient(&openapi.Configuration{
		Servers: openapi.ServerConfigurations{
			{
				URL:         server.URL,
				Description: "Default server",
			},
		},
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + *adminLoginRsp.Token.Token,
		},
	})

	policySvc := client.PolicySvcAPI
	instanceId := "instance-1"

	// Create a rate limit policy instance
	rateLimitReq := openapi.PolicySvcUpsertInstanceRequest{
		Instance: &openapi.PolicySvcInstance{
			Id:       &instanceId,
			Endpoint: openapi.PtrString("/test-endpoint"),
			TemplateId: openapi.PolicySvcTemplateId(
				policytypes.RateLimitPolicyTemplate.Id,
			),
			Parameters: openapi.PolicySvcParameters{
				RateLimit: &openapi.PolicySvcRateLimitParameters{
					MaxRequests: openapi.PtrInt32(5),
					TimeWindow:  openapi.PtrString("1m"),
					Entity:      openapi.EntityUserID.Ptr(),
					Scope:       openapi.ScopeEndpoint.Ptr(),
				},
			},
		},
	}
	_, _, err = policySvc.UpsertInstance(context.Background(), instanceId).
		Body(rateLimitReq).
		Execute()
	require.NoError(t, err)

	t.Run("allow up to the limit", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			checkRsp, _, err := policySvc.Check(context.Background()).
				Body(openapi.PolicySvcCheckRequest{
					Endpoint: openapi.PtrString("/test-endpoint"),
					Method:   openapi.PtrString("GET"),
					Ip:       openapi.PtrString("127.0.0.1"),
					UserId:   openapi.PtrString("user-1"),
				}).
				Execute()
			require.NoError(t, err)
			require.Equal(t, true, checkRsp.Allowed)
		}
	})

	t.Run("exceeding the limit", func(t *testing.T) {
		checkRsp, _, err := policySvc.Check(context.Background()).
			Body(openapi.PolicySvcCheckRequest{
				Endpoint: openapi.PtrString("/test-endpoint"),
				Method:   openapi.PtrString("GET"),
				Ip:       openapi.PtrString("127.0.0.1"),
				UserId:   openapi.PtrString("user-1"),
			}).
			Execute()
		require.NoError(t, err)
		require.Equal(t, false, checkRsp.Allowed)
	})
}
