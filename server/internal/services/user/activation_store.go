package userservice

import (
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

func (s *UserService) getActivation(
	appId string,
	userId string,
	device string,
) (*user.Activation, bool, error) {
	if device == "" {
		device = unknownDevice
	}

	rows, err := s.activationStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.Equals(datastore.Field("device"), device),
	).Find()
	if err != nil {
		return nil, false, err
	}
	if len(rows) == 0 {
		return nil, false, nil
	}

	activation := rows[0].(*user.Activation)
	for _, extra := range rows[1:] {
		if err := s.activationStore.Query(datastore.Id(extra.GetId())).Delete(); err != nil {
			return nil, false, err
		}
	}

	return activation, true, nil
}

func (s *UserService) getActiveOrganizationId(
	appId string,
	userId string,
	device string,
) (string, error) {
	activation, found, err := s.getActivation(appId, userId, device)
	if err != nil {
		return "", err
	}
	if !found {
		return "", nil
	}

	membership, membershipFound, err := s.findMembership(appId, userId, activation.OrganizationId)
	if err != nil {
		return "", err
	}
	if !membershipFound || membership.Status != user.MembershipStatusAccepted {
		if err := s.activationStore.Query(datastore.Id(activation.Id)).Delete(); err != nil {
			return "", err
		}
		return "", nil
	}

	return activation.OrganizationId, nil
}

func (s *UserService) setActivation(
	appId string,
	userId string,
	device string,
	organizationId string,
) error {
	if device == "" {
		device = unknownDevice
	}

	now := time.Now()
	activation, found, err := s.getActivation(appId, userId, device)
	if err != nil {
		return err
	}

	if organizationId == "" {
		if !found {
			return nil
		}
		return s.activationStore.Query(datastore.Id(activation.Id)).Delete()
	}

	if found {
		return s.activationStore.Query(datastore.Id(activation.Id)).UpdateFields(map[string]interface{}{
			"organizationId": organizationId,
			"updatedAt":      now,
		})
	}

	id := sdk.Id("act")
	internalId, err := sdk.InternalId(appId, id)
	if err != nil {
		return errors.Wrap(err, "failed to create activation internal id")
	}

	return s.activationStore.Upsert(&user.Activation{
		InternalId:     internalId,
		Id:             id,
		AppId:          appId,
		UserId:         userId,
		Device:         device,
		OrganizationId: organizationId,
		CreatedAt:      now,
		UpdatedAt:      now,
	})
}

func (s *UserService) deleteActivations(
	appId string,
	userId string,
	organizationId string,
) error {
	filters := []datastore.Filter{
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.Equals(datastore.Field("userId"), userId),
	}
	if organizationId != "" {
		filters = append(filters, datastore.Equals(datastore.Field("organizationId"), organizationId))
	}

	return s.activationStore.Query(filters...).Delete()
}
