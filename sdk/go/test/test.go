package test

import (
	"context"
	"fmt"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"go.uber.org/mock/gomock"
)

func AdminClient(clientFactory sdk.ClientFactory) (*openapi.APIClient, string, error) {
	userSvc := clientFactory.Client().UserSvcAPI

	adminLoginRsp, _, err := userSvc.Login(context.Background()).Body(openapi.UserSvcLoginRequest{
		Slug:     openapi.PtrString("openorch"),
		Password: openapi.PtrString("changeme"),
	}).Execute()
	if err != nil {
		return nil, "", err
	}

	return clientFactory.Client(sdk.WithToken(*adminLoginRsp.Token.Token)), *adminLoginRsp.Token.Token, nil
}

func MakeClients(clientFactory sdk.ClientFactory, num int) ([]*openapi.APIClient, error) {
	var ret []*openapi.APIClient

	for i := 0; i < num; i++ {
		slug := fmt.Sprintf("test-user-slug-%v", i)
		password := fmt.Sprintf("testUserPassword%v", i)
		username := fmt.Sprintf("Test User Name %v", i)

		token, err := sdk.RegisterUserAccount(clientFactory.Client().UserSvcAPI, slug, password, username)
		if err != nil {
			return nil, err
		}

		c := clientFactory.Client(sdk.WithToken(token))

		ret = append(ret, c)
	}

	return ret, nil
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
			Token: openapi.PtrString("HELLO"),
		},
	}
	mockLoginRequest := openapi.ApiLoginRequest{
		ApiService: mockUserSvc,
	}
	mockAddPermissionToRoleRequest := openapi.ApiAssignPermissionsRequest{
		ApiService: mockUserSvc,
	}
	expectedUserSvcAddPermissionToRoleResponse := map[string]any{}
	expectedUserSvcUpsertPermissionResponse := &openapi.UserSvcSavePermissionsResponse{}
	mockUpsertPermissionRequest := openapi.ApiSavePermissionsRequest{
		ApiService: mockUserSvc,
	}

	mockIsAuthorizedRequest := openapi.ApiIsAuthorizedRequest{
		ApiService: mockUserSvc,
	}

	mockUserSvc.EXPECT().GetPublicKey(ctx).Return(openapi.ApiGetPublicKeyRequest{
		ApiService: mockUserSvc,
	}).AnyTimes()
	mockUserSvc.EXPECT().GetPublicKeyExecute(gomock.Any()).Return(&openapi.UserSvcGetPublicKeyResponse{
		PublicKey: openapi.PtrString(""),
	}, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().Login(ctx).Return(mockLoginRequest).AnyTimes()
	mockUserSvc.EXPECT().LoginExecute(gomock.Any()).Return(expectedUserSvcLoginResponse, nil, nil).AnyTimes()
	mockUserSvc.EXPECT().SavePermissions(ctx).Return(mockUpsertPermissionRequest).AnyTimes()
	mockUserSvc.EXPECT().SavePermissionsExecute(gomock.Any()).Return(expectedUserSvcUpsertPermissionResponse, nil, nil).AnyTimes()
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
				Id:   openapi.PtrString(id),   // Dynamically evaluate
				Slug: openapi.PtrString(slug), // Dynamically evaluate
			},
		}, nil, nil
	}).
		AnyTimes()

	return mockUserSvc
}
