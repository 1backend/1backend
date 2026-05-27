package enroll

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/1backend/1backend/cli/oo/util"
	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// List
func List(cmd *cobra.Command, args []string, role, userId, contactId string) error {
	ctx := cmd.Context()

	url, token, err := util.GetSelectedUrlAndToken(cmd)
	if err != nil {
		return errors.Wrap(err, "cannot get env url")
	}

	cf := client.NewApiClientFactory(url)
	apiClient := cf.Client(client.WithToken(token))

	req := openapi.UserSvcListEnrollsRequest{}

	if role != "" {
		req.Role = &role
	}
	if userId != "" {
		req.UserId = &userId
	}
	if contactId != "" {
		req.ContactId = &contactId
	}

	rsp, hrsp, err := apiClient.
		UserSvcAPI.ListEnrolls(ctx).
		Body(req).
		Execute()
	if err != nil {
		return util.ErrorWithBody(err, hrsp, "failed to list enrolls")
	}

	contactIdsByUserId := map[string]string{}
	slugsByUserId := map[string]string{}
	totpEnabledByUserId := map[string]string{}
	userIds := userIdsFromEnrolls(rsp.Enrolls)
	contactIdFallbackUserIds := stringSet(userIdsMissingContactIdsFromEnrolls(rsp.Enrolls))
	if len(userIds) > 0 {
		usersRsp, usersHTTPRsp, err := apiClient.
			UserSvcAPI.ListUsers(ctx).
			Body(openapi.UserSvcListUsersRequest{
				Ids:   userIds,
				Limit: openapi.PtrInt32(int32(len(userIds))),
			}).
			Execute()
		if err != nil {
			fmt.Fprintln(
				os.Stderr,
				util.ErrorWithBody(err, usersHTTPRsp, "warning: failed to resolve enroll user details"),
			)
		} else {
			for _, user := range usersRsp.Users {
				slugsByUserId[user.Id] = user.Slug
				totpEnabledByUserId[user.Id] = boolPtrString(user.TotpEnabled)
				if _, ok := contactIdFallbackUserIds[user.Id]; ok {
					contactIdsByUserId[user.Id] = strings.Join(user.ContactIds, ", ")
				}
			}
		}
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(
		writer,
		"ENROLL ID\tROLE\tUSER ID\tUSER SLUG\tTOTP\tCONTACT IDS",
	)

	for _, enroll := range rsp.Enrolls {
		userId := ""
		if enroll.UserId != nil {
			userId = *enroll.UserId
		}

		userSlug := ""
		if userId != "" {
			userSlug = slugsByUserId[userId]
		}

		totpEnabled := ""
		if userId != "" {
			totpEnabled = totpEnabledByUserId[userId]
		}

		contactIds := contactIdsForEnroll(enroll, contactIdsByUserId)

		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%s\t%s\t%s\n",
			enroll.Id,
			enroll.Role,
			userId,
			userSlug,
			totpEnabled,
			contactIds,
		)
	}

	return nil
}

func userIdsFromEnrolls(enrolls []openapi.UserSvcEnroll) []string {
	seen := map[string]struct{}{}
	userIds := []string{}
	for _, enroll := range enrolls {
		if enroll.UserId == nil || *enroll.UserId == "" {
			continue
		}
		if _, ok := seen[*enroll.UserId]; ok {
			continue
		}

		seen[*enroll.UserId] = struct{}{}
		userIds = append(userIds, *enroll.UserId)
	}

	return userIds
}

func userIdsMissingContactIdsFromEnrolls(enrolls []openapi.UserSvcEnroll) []string {
	seen := map[string]struct{}{}
	userIds := []string{}
	for _, enroll := range enrolls {
		if enroll.UserId == nil || *enroll.UserId == "" {
			continue
		}
		if enroll.ContactIds != nil {
			continue
		}
		if enroll.ContactId != nil && *enroll.ContactId != "" {
			continue
		}
		if _, ok := seen[*enroll.UserId]; ok {
			continue
		}

		seen[*enroll.UserId] = struct{}{}
		userIds = append(userIds, *enroll.UserId)
	}

	return userIds
}

func contactIdsForEnroll(enroll openapi.UserSvcEnroll, contactIdsByUserId map[string]string) string {
	if enroll.ContactIds != nil {
		return strings.Join(enroll.ContactIds, ", ")
	}
	if enroll.ContactId != nil && *enroll.ContactId != "" {
		return *enroll.ContactId
	}
	if enroll.UserId != nil {
		return contactIdsByUserId[*enroll.UserId]
	}
	return ""
}

func stringSet(values []string) map[string]struct{} {
	set := map[string]struct{}{}
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func boolPtrString(value *bool) string {
	if value != nil && *value {
		return "true"
	}
	return "false"
}
