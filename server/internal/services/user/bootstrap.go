package userservice

import (
	"context"

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

func (s *UserService) BootstrapSavePermits(ctx context.Context, permits []user.PermitInput) error {
	if len(permits) == 0 {
		return nil
	}

	permitsByAppHost := map[string][]*user.PermitInput{}
	for i := range permits {
		permit := permits[i]
		appHost := bootstrapAppHost(permit.AppHost)
		permit.AppHost = ""
		permitsByAppHost[appHost] = append(permitsByAppHost[appHost], &permit)
	}

	for appHost, reqPermits := range permitsByAppHost {
		appId, err := s.BootstrapAppID(appHost)
		if err != nil {
			return errors.Wrapf(err, "bootstrap app %s", appHost)
		}

		if err := s.savePermits(appId, ctx, &user.SavePermitsRequest{
			Permits: reqPermits,
		}); err != nil {
			return errors.Wrapf(err, "bootstrap permits for app %s", appHost)
		}
	}

	return nil
}

func (s *UserService) BootstrapSaveEnrolls(ctx context.Context, enrolls []user.EnrollInput) error {
	if len(enrolls) == 0 {
		return nil
	}

	app, err := s.getOrCreateApp(user.DefaultAppHost)
	if err != nil {
		return errors.Wrap(err, "bootstrap default app")
	}

	admin, err := s.bootstrapAdminUser()
	if err != nil {
		return err
	}

	_, err = s.saveEnrolls(app.Id, admin.Id, &user.SaveEnrollsRequest{
		Enrolls: enrolls,
	})
	return err
}

func (s *UserService) BootstrapAppID(host string) (string, error) {
	host = bootstrapAppHost(host)
	if host == "*" {
		return "*", nil
	}

	app, err := s.getOrCreateApp(host)
	if err != nil {
		return "", err
	}
	return app.Id, nil
}

func bootstrapAppHost(host string) string {
	if host == "" {
		return user.DefaultAppHost
	}
	if host == "_all" {
		return "*"
	}
	return host
}

func (s *UserService) bootstrapAdminUser() (*user.User, error) {
	row, found, err := s.userStore.Query(
		datastore.Equals([]string{"slug"}, "1backend"),
	).FindOne()
	if err != nil {
		return nil, errors.Wrap(err, "query bootstrap admin user")
	}
	if !found {
		return nil, errors.New("bootstrap admin user not found")
	}
	return row.(*user.User), nil
}
