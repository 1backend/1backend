package secretservice

import (
	"context"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	secretssdk "github.com/1backend/1backend/sdk/go/secrets"
	secret "github.com/1backend/1backend/server/internal/services/secret/types"
	"github.com/pkg/errors"
)

func (cs *SecretService) BootstrapSaveSecrets(
	ctx context.Context,
	secrets []*secret.SecretInput,
	appIDFromHost func(string) (string, error),
) error {
	if len(secrets) == 0 {
		return nil
	}
	if appIDFromHost == nil {
		return errors.New("bootstrap app resolver is nil")
	}

	cs.options.Lock.Acquire(ctx, "secret-svc-save")
	defer cs.options.Lock.Release(ctx, "secret-svc-save")

	for _, input := range secrets {
		if input == nil {
			continue
		}

		next := *input
		appHost := next.AppHost
		if appHost == "" {
			appHost = defaultApp
		}

		appID, err := appIDFromHost(appHost)
		if err != nil {
			return errors.Wrapf(err, "bootstrap app %s", appHost)
		}

		existingI, found, err := cs.secretStore.Query(
			datastore.Equals([]string{"appId"}, appID),
			datastore.Equals([]string{"id"}, next.Id),
		).FindOne()
		if err != nil {
			return errors.Wrap(err, "query secret")
		}

		if next.Id == "" {
			next.Id = sdk.Id("secr")
		}

		if !next.Encrypted {
			next.Value, err = secretssdk.Encrypt(next.Value, cs.options.SecretEncryptionKey)
			if err != nil {
				return errors.Wrap(err, "failed to encrypt secret")
			}
		}

		entry := &secret.Secret{
			AppId:             appID,
			Id:                next.Id,
			Value:             next.Value,
			Readers:           next.Readers,
			Writers:           next.Writers,
			Deleters:          next.Deleters,
			CanChangeReaders:  next.CanChangeReaders,
			CanChangeWriters:  next.CanChangeWriters,
			CanChangeDeleters: next.CanChangeDeleters,
			Encrypted:         next.Encrypted,
			Checksum:          next.Checksum,
			ChecksumAlgorithm: next.ChecksumAlgorithm,
		}

		if found {
			existing := existingI.(*secret.Secret)
			entry.Id = existing.Id
			entry.Readers = existing.Readers
			entry.Writers = existing.Writers
			entry.Deleters = existing.Deleters
			entry.CanChangeReaders = existing.CanChangeReaders
			entry.CanChangeWriters = existing.CanChangeWriters
			entry.CanChangeDeleters = existing.CanChangeDeleters

			if next.Readers != nil {
				entry.Readers = next.Readers
			}
			if next.Writers != nil {
				entry.Writers = next.Writers
			}
			if next.Deleters != nil {
				entry.Deleters = next.Deleters
			}
			if next.CanChangeReaders != nil {
				entry.CanChangeReaders = next.CanChangeReaders
			}
			if next.CanChangeWriters != nil {
				entry.CanChangeWriters = next.CanChangeWriters
			}
			if next.CanChangeDeleters != nil {
				entry.CanChangeDeleters = next.CanChangeDeleters
			}
		}

		entry.InternalId, err = sdk.InternalId(appID, entry.Id)
		if err != nil {
			return errors.Wrap(err, "failed to create internal id")
		}

		if err := cs.checkSum(entry); err != nil {
			return errors.Wrap(err, "checksum failed")
		}

		if err := cs.secretStore.Upsert(entry); err != nil {
			return errors.Wrap(err, "save secret")
		}
	}

	return nil
}
