package userservice_test

import (
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestSendOtp__TemplateFallback(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	manyClients, _, err := test.MakeClients(clientFactory, sdk.DefaultTestAppHost, 1)
	require.NoError(t, err)
	client1 := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.SecretSvcAPI.SaveSecrets(t.Context()).Body(openapi.SecretSvcSaveSecretsRequest{
		Secrets: []openapi.SecretSvcSecretInput{
			{
				Id:      "otp-email-subject-hu",
				Value:   openapi.PtrString("Helló {{.Code}}"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "otp-email-body-hu",
				Value:   openapi.PtrString("This is the body {{.Code}}"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "otp-email-type",
				Value:   openapi.PtrString("text/html"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "sender-name",
				Value:   openapi.PtrString("Test Sender Name"),
				Readers: []string{"user-svc"},
			},
		},
	}).Execute()
	require.NoError(t, err)

	t.Run("Falls back to Global English when no App secrets exist", func(t *testing.T) {
		// Send OTP
		otpRsp, _, err := client1.UserSvcAPI.SendOtp(t.Context()).
			Body(openapi.UserSvcSendOtpRequest{
				AppHost:         sdk.DefaultTestAppHost,
				ContactId:       "fallback@test.com",
				ContactPlatform: "email",
			}).AcceptLanguage("hu").Execute()
		require.NoError(t, err)

		// In a real integration test, you'd check a mock EmailSvc or logs.
		// Since options.Test is true, we verify the response contains the generated code.
		require.NotEmpty(t, otpRsp.Code)
		require.NotEmpty(t, otpRsp.OtpId)

		require.Contains(t, *otpRsp.Subject, "Helló")
		require.Contains(t, *otpRsp.Subject, *otpRsp.Code)

		require.Contains(t, *otpRsp.Body, "This is the body")
		require.Contains(t, *otpRsp.Body, *otpRsp.Code)

		require.Equal(t, *otpRsp.ContentType, "text/html")
		require.Equal(t, *otpRsp.FromName, "Test Sender Name")
	})
}
