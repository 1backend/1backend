package policyservice

import (
	"context"

	policy "github.com/1backend/1backend/server/internal/services/policy/types"
	"github.com/pkg/errors"
)

func (s *PolicyService) BootstrapSaveInstances(
	ctx context.Context,
	instances []*policy.Instance,
) error {
	if len(instances) == 0 {
		return nil
	}

	s.options.Lock.Acquire(ctx, "policy-svc-save-instances")
	defer s.options.Lock.Release(ctx, "policy-svc-save-instances")

	for _, instance := range instances {
		if instance == nil {
			continue
		}

		next := *instance
		if err := s.upsertInstance(&next); err != nil {
			return errors.Wrapf(err, "bootstrap policy instance %s", next.Id)
		}
	}

	return nil
}
