package test

import (
	"context"
	"fmt"
	"net/http"

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
	IdFactory           func() string
	SlugFactory         func() string
	IsAuthorizedFactory func() bool
}

type MockUserOption func(*MockUserOptions)

func WithIdFactory(idFactory func() string) MockUserOption {
	return func(o *MockUserOptions) {
		o.IdFactory = idFactory
	}
}

func WithSlugFactory(slugFactory func() string) MockUserOption {
	return func(o *MockUserOptions) {
		o.SlugFactory = slugFactory
	}
}

func WithIsAuthorizedFactory(isAuthorizedFactory func() bool) MockUserOption {
	return func(o *MockUserOptions) {
		o.IsAuthorizedFactory = isAuthorizedFactory
	}
}

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
	mockAddPermissionToRoleRequest := openapi.ApiAssignPermissionsRequest{
		ApiService: mockUserSvc,
	}
	expectedUserSvcAddPermissionToRoleResponse := map[string]any{}

	mockIsAuthorizedRequest := openapi.ApiIsAuthorizedRequest{
		ApiService: mockUserSvc,
	}

	mockUserSvc.EXPECT().GetPublicKey(ctx).Return(openapi.ApiGetPublicKeyRequest{
		ApiService: mockUserSvc,
	}).AnyTimes()
	mockUserSvc.EXPECT().GetPublicKeyExecute(gomock.Any()).Return(&openapi.UserSvcGetPublicKeyResponse{
		PublicKey: "",
	}, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().Login(ctx).Return(mockLoginRequest).AnyTimes()
	mockUserSvc.EXPECT().LoginExecute(gomock.Any()).Return(expectedUserSvcLoginResponse, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().AssignPermissions(ctx).Return(mockAddPermissionToRoleRequest).AnyTimes()
	mockUserSvc.EXPECT().AssignPermissionsExecute(gomock.Any()).Return(expectedUserSvcAddPermissionToRoleResponse, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().IsAuthorized(gomock.Any(), gomock.Any()).Return(mockIsAuthorizedRequest).AnyTimes()
	mockUserSvc.EXPECT().IsAuthorizedExecute(gomock.Any()).DoAndReturn(func(req openapi.ApiIsAuthorizedRequest) (*openapi.UserSvcIsAuthorizedResponse, *http.Response, error) {
		var isAuthorized bool
		if opts.IsAuthorizedFactory != nil {
			isAuthorized = opts.IsAuthorizedFactory()
		}

		var id string
		if opts.IdFactory != nil {
			id = opts.IdFactory()
		}

		var slug string
		if opts.SlugFactory != nil {
			slug = opts.SlugFactory()
		}

		return &openapi.UserSvcIsAuthorizedResponse{
			Authorized: openapi.PtrBool(isAuthorized), // Dynamically evaluate
			User: &openapi.UserSvcUser{
				Id:   id,   // Dynamically evaluate
				Slug: slug, // Dynamically evaluate
			},
		}, nil, nil
	}).
		AnyTimes()

	return mockUserSvc
}
