package userservice

import (
	"reflect"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

func (s *UserService) findMembership(
	appId string,
	userId string,
	organizationId string,
) (*user.Membership, bool, error) {
	rows, err := s.membershipStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.Equals(datastore.Field("organizationId"), organizationId),
	).Find()
	if err != nil {
		return nil, false, err
	}
	if len(rows) == 0 {
		return nil, false, nil
	}

	membership := rows[0].(*user.Membership)
	for _, extra := range rows[1:] {
		if err := s.membershipStore.Query(datastore.Id(extra.GetId())).Delete(); err != nil {
			return nil, false, err
		}
	}

	return membership, true, nil
}

func (s *UserService) createMembership(
	appId string,
	userId string,
	organizationId string,
	status user.MembershipStatus,
	Roles []string,
	invitedBy string,
	acceptedAt *time.Time,
) (*user.Membership, error) {
	now := time.Now()
	id := sdk.Id("memb")
	internalId, err := sdk.InternalId(appId, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create membership internal id")
	}

	membership := &user.Membership{
		InternalId:     internalId,
		Id:             id,
		AppId:          appId,
		UserId:         userId,
		OrganizationId: organizationId,
		Status:         status,
		Roles:          Roles,
		InvitedBy:      invitedBy,
		AcceptedAt:     acceptedAt,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.membershipStore.Upsert(membership); err != nil {
		return nil, err
	}

	return membership, nil
}

func (s *UserService) updateMembership(
	membership *user.Membership,
	status user.MembershipStatus,
	Roles []string,
	invitedBy string,
	acceptedAt *time.Time,
) (*user.Membership, bool, error) {
	next := *membership
	next.Status = status
	next.Roles = Roles
	next.InvitedBy = invitedBy
	next.AcceptedAt = acceptedAt
	next.UpdatedAt = time.Now()

	changed := membership.Status != next.Status ||
		!reflect.DeepEqual(membership.Roles, next.Roles) ||
		membership.InvitedBy != next.InvitedBy ||
		!reflect.DeepEqual(membership.AcceptedAt, next.AcceptedAt)
	if !changed {
		return membership, false, nil
	}

	if err := s.membershipStore.Upsert(&next); err != nil {
		return nil, false, err
	}

	return &next, true, nil
}
