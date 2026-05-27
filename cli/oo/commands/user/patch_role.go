package user

import (
	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// PatchRole assigns a role to a user. If fromRole is set, matching old
// user-targeted enrolls are removed after the new role is saved.
func PatchRole(cmd *cobra.Command, args []string, fromRole string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url and token")
	}

	userId := args[0]
	role := args[1]
	if fromRole == role {
		return errors.New("from-role must differ from the new role")
	}

	cf := client.NewApiClientFactory(url)
	apiClient := cf.Client(client.WithToken(token))

	_, hrsp, err := apiClient.
		UserSvcAPI.SaveEnrolls(ctx).
		Body(openapi.UserSvcSaveEnrollsRequest{
			Enrolls: []openapi.UserSvcEnrollInput{
				{
					UserId: openapi.PtrString(userId),
					Role:   role,
				},
			},
		}).
		Execute()
	if err != nil {
		return util.ErrorWithBody(err, hrsp, "failed to assign user role")
	}

	if fromRole == "" {
		return nil
	}

	listRsp, listHTTPRsp, err := apiClient.
		UserSvcAPI.ListEnrolls(ctx).
		Body(openapi.UserSvcListEnrollsRequest{
			Role: openapi.PtrString(fromRole),
		}).
		Execute()
	if err != nil {
		return util.ErrorWithBody(err, listHTTPRsp, "failed to find old user role enrolls")
	}

	enrollIds := enrollIdsForUserRole(listRsp.Enrolls, userId, fromRole)
	if len(enrollIds) == 0 {
		return nil
	}

	_, deleteHTTPRsp, err := apiClient.
		UserSvcAPI.DeleteEnrolls(ctx).
		Body(openapi.UserSvcDeleteEnrollsRequest{
			Ids: enrollIds,
		}).
		Execute()
	if err != nil {
		return util.ErrorWithBody(err, deleteHTTPRsp, "failed to remove old user role")
	}

	return nil
}

func enrollIdsForUserRole(enrolls []openapi.UserSvcEnroll, userId, role string) []string {
	ids := []string{}
	for _, enroll := range enrolls {
		if enroll.UserId == nil || *enroll.UserId != userId {
			continue
		}
		if enroll.Role != role {
			continue
		}
		ids = append(ids, enroll.Id)
	}
	return ids
}
