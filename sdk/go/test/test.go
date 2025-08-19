package test

import (
	"context"
	"fmt"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"go.uber.org/mock/gomock"
)

func AdminClient(clientFactory client.ClientFactory) (*openapi.APIClient, string, error) {
	userSvc := clientFactory.Client().UserSvcAPI

	adminLoginRsp, _, err := userSvc.Login(context.Background()).Body(openapi.UserSvcLoginRequest{
		Slug:     openapi.PtrString("1backend"),
		Password: openapi.PtrString("changeme"),
	}).Execute()
	if err != nil {
		return nil, "", err
	}

	return clientFactory.Client(client.WithToken(adminLoginRsp.Token.Token)), adminLoginRsp.Token.Token, nil
}

func MakeClients(clientFactory client.ClientFactory, num int) ([]*openapi.APIClient, []*openapi.UserSvcAuthToken, error) {
	var (
		clients []*openapi.APIClient
		tokens  []*openapi.UserSvcAuthToken
	)

	for i := 0; i < num; i++ {
		slug := fmt.Sprintf("test-user-slug-%v", i)
		password := fmt.Sprintf("testUserPassword%v", i)
		username := fmt.Sprintf("Test User Name %v", i)

		token, err := boot.RegisterUserAccount(clientFactory.Client().UserSvcAPI, slug, password, username)
		if err != nil {
			return nil, nil, err
		}

		c := clientFactory.Client(client.WithToken(token.Token))

		clients = append(clients, c)
		tokens = append(tokens, token)
	}

	return clients, tokens, nil
}

func LoggedInClient(
	clientFactory client.ClientFactory,
	slug,
	password string,
) (*openapi.APIClient, *openapi.UserSvcAuthToken, error) {
	loginReq := openapi.UserSvcLoginRequest{
		Slug:     openapi.PtrString(slug),
		Password: openapi.PtrString(password),
	}

	loginRsp, _, err := clientFactory.Client().UserSvcAPI.Login(context.Background()).
		Body(loginReq).
		Execute()
	if err != nil {
		return nil, nil, err
	}

	cl := clientFactory.Client(client.WithToken(loginRsp.Token.Token))
	return cl, loginRsp.Token, nil
}

type MockUserOptions struct {
}

type MockUserOption func(*MockUserOptions)

// Returns a mock User Svc with expects set up for calls that happen during the startup of the services.
func MockUserSvc(ctx context.Context, ctrl *gomock.Controller, options ...MockUserOption) *openapi.MockUserSvcAPI {
	opts := &MockUserOptions{}
	for _, o := range options {
		o(opts)
	}

	mockUserSvc := openapi.NewMockUserSvcAPI(ctrl)

	expectedUserSvcLoginResponse := &openapi.UserSvcLoginResponse{
		Token: &openapi.UserSvcAuthToken{
			Token: "HELLO",
		},
	}
	mockLoginRequest := openapi.ApiLoginRequest{
		ApiService: mockUserSvc,
	}
	mockRegisterRequest := openapi.ApiRegisterRequest{
		ApiService: mockUserSvc,
	}
	expectedUserSvcRegisterResponse := &openapi.UserSvcRegisterResponse{
		Token: &openapi.UserSvcAuthToken{
			Token: "HELLO",
		},
	}
	mockAddPermissionToRoleRequest := openapi.ApiSavePermitsRequest{
		ApiService: mockUserSvc,
	}
	expectedUserSvcAddPermissionToRoleResponse := map[string]any{}

	mockUserSvc.EXPECT().GetPublicKey(ctx).Return(openapi.ApiGetPublicKeyRequest{
		ApiService: mockUserSvc,
	}).AnyTimes()
	mockUserSvc.EXPECT().GetPublicKeyExecute(gomock.Any()).Return(&openapi.UserSvcGetPublicKeyResponse{
		PublicKey: "",
	}, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().Login(ctx).Return(mockLoginRequest).AnyTimes()
	mockUserSvc.EXPECT().Register(ctx).Return(mockRegisterRequest).AnyTimes()
	mockUserSvc.EXPECT().RegisterExecute(gomock.Any()).Return(expectedUserSvcRegisterResponse, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().LoginExecute(gomock.Any()).Return(expectedUserSvcLoginResponse, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().SavePermits(ctx).Return(mockAddPermissionToRoleRequest).AnyTimes()
	mockUserSvc.EXPECT().SavePermitsExecute(gomock.Any()).Return(expectedUserSvcAddPermissionToRoleResponse, nil, nil).AnyTimes()

	return mockUserSvc
}
