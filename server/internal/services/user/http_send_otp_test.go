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

	appHost := "some-app"

	manyClients, _, err := test.MakeClients(clientFactory, appHost, 1)
	require.NoError(t, err)
	client1 := manyClients[0]

	adminClient, _, err := test.AdminClient(clientFactory, sdk.DefaultTestAppHost)
	require.NoError(t, err)

	_, _, err = adminClient.SecretSvcAPI.SaveSecrets(t.Context()).Body(openapi.SecretSvcSaveSecretsRequest{
		Secrets: []openapi.SecretSvcSecretInput{
			{
				Id:      "otp-email-subject-hu",
				AppHost: &appHost,
				Value:   openapi.PtrString("Helló {{.Code}}"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "otp-email-subject",
				AppHost: &appHost,
				Value:   openapi.PtrString("Hello {{.Code}}"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "otp-email-body-hu",
				AppHost: &appHost,
				Value:   openapi.PtrString("Ez a test {{.Code}}"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "otp-email-body",
				AppHost: &appHost,
				Value:   openapi.PtrString("This is the body {{.Code}}"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "otp-email-type",
				AppHost: &appHost,
				Value:   openapi.PtrString("text/html"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "sender-name",
				AppHost: &appHost,
				Value:   openapi.PtrString("Test Sender Name"),
				Readers: []string{"user-svc"},
			},
			{
				Id:      "sender-email",
				AppHost: &appHost,
				Value:   openapi.PtrString("something@random.com"),
				Readers: []string{"user-svc"},
			},
		},
	}).Execute()
	require.NoError(t, err)

	t.Run("Uses language specific secrets", func(t *testing.T) {
		otpRsp, _, err := client1.UserSvcAPI.SendOtp(t.Context()).
			Body(openapi.UserSvcSendOtpRequest{
				AppHost:         appHost,
				ContactId:       "fallback@test.com",
				ContactPlatform: "email",
			}).AcceptLanguage("hu").Execute()
		require.NoError(t, err)

		require.NotEmpty(t, otpRsp.Code)
		require.NotEmpty(t, otpRsp.OtpId)

		require.Contains(t, *otpRsp.Subject, "Helló")
		require.Contains(t, *otpRsp.Subject, *otpRsp.Code)

		require.Contains(t, *otpRsp.Body, "Ez a test")
		require.Contains(t, *otpRsp.Body, *otpRsp.Code)

		require.Equal(t, *otpRsp.ContentType, "text/html")
		require.Equal(t, *otpRsp.FromName, "Test Sender Name")
	})

	t.Run("Falls back to global English", func(t *testing.T) {
		otpRsp, _, err := client1.UserSvcAPI.SendOtp(t.Context()).
			Body(openapi.UserSvcSendOtpRequest{
				AppHost:         appHost,
				ContactId:       "fallback@test.com",
				ContactPlatform: "email",
			}).Execute()
		require.NoError(t, err)

		require.NotEmpty(t, otpRsp.Code)
		require.NotEmpty(t, otpRsp.OtpId)

		require.Contains(t, *otpRsp.Subject, "Hello")
		require.Contains(t, *otpRsp.Subject, *otpRsp.Code)

		require.Contains(t, *otpRsp.Body, "This is the body")
		require.Contains(t, *otpRsp.Body, *otpRsp.Code)

		require.Equal(t, *otpRsp.ContentType, "text/html")
		require.Equal(t, *otpRsp.FromName, "Test Sender Name")
		require.Equal(t, *otpRsp.FromEmail, "something@random.com")
	})
}
