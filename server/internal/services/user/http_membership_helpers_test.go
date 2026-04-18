package userservice_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/stretchr/testify/require"
)

func saveMembership(
	t *testing.T,
	apiClient *openapi.APIClient,
	organizationId string,
	userId string,
	body *openapi.UserSvcSaveMembershipRequest,
) (*openapi.UserSvcSaveMembershipResponse, *http.Response) {
	t.Helper()

	req := apiClient.UserSvcAPI.SaveMembership(
		context.Background(),
		organizationId,
		userId,
	)
	if body != nil {
		req = req.Body(*body)
	}

	rsp, httpResp, err := req.Execute()
	require.NoError(t, err)
	t.Cleanup(func() {
		closeResponseBody(t, httpResp)
	})

	return rsp, httpResp
}

func acceptMembership(
	t *testing.T,
	apiClient *openapi.APIClient,
	organizationId string,
	body *openapi.UserSvcAcceptMembershipRequest,
) (*openapi.UserSvcAcceptMembershipResponse, *http.Response) {
	t.Helper()

	req := apiClient.UserSvcAPI.AcceptMembership(
		context.Background(),
		organizationId,
	)
	if body != nil {
		req = req.Body(*body)
	}

	rsp, httpResp, err := req.Execute()
	require.NoError(t, err)
	t.Cleanup(func() {
		closeResponseBody(t, httpResp)
	})

	return rsp, httpResp
}

func declineMembership(
	t *testing.T,
	apiClient *openapi.APIClient,
	organizationId string,
) (*openapi.UserSvcDeclineMembershipResponse, *http.Response) {
	t.Helper()

	rsp, httpResp, err := apiClient.UserSvcAPI.DeclineMembership(
		context.Background(),
		organizationId,
	).Execute()
	require.NoError(t, err)
	t.Cleanup(func() {
		closeResponseBody(t, httpResp)
	})

	return rsp, httpResp
}

func listMemberships(
	t *testing.T,
	apiClient *openapi.APIClient,
	body *openapi.UserSvcListMembershipsRequest,
) (*openapi.UserSvcListMembershipsResponse, *http.Response) {
	t.Helper()

	req := apiClient.UserSvcAPI.ListMemberships(context.Background())
	if body != nil {
		req = req.Body(*body)
	}

	rsp, httpResp, err := req.Execute()
	require.NoError(t, err)
	t.Cleanup(func() {
		closeResponseBody(t, httpResp)
	})

	return rsp, httpResp
}

func closeResponseBody(
	t *testing.T,
	httpResp *http.Response,
) {
	t.Helper()

	if httpResp == nil || httpResp.Body == nil {
		return
	}

	_, err := io.Copy(io.Discard, httpResp.Body)
	require.NoError(t, err)
	err = httpResp.Body.Close()
	require.NoError(t, err)
}
