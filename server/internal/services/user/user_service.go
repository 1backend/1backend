/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"crypto/rsa"
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/dgraph-io/ristretto"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"

	"github.com/google/uuid"
)

// Each auth token belongs to a device.
// When a device is not specified at login
// or registration the unknown device is used.
const unknownDevice = "unknown"

type UserService struct {
	options *universe.Options

	token string

	usersStore         datastore.DataStore
	credentialsStore   datastore.DataStore
	passwordsStore     datastore.DataStore
	authTokensStore    datastore.DataStore
	keyPairsStore      datastore.DataStore
	contactsStore      datastore.DataStore
	organizationsStore datastore.DataStore
	membershipsStore   datastore.DataStore
	permitsStore       datastore.DataStore
	enrollsStore       datastore.DataStore

	privateKey   *rsa.PrivateKey
	publicKeyPem string

	configCache map[string]any

	tokenReplacementCache *ristretto.Cache
}

func NewUserService(
	options *universe.Options,
) (*UserService, error) {
	usersStore, err := options.DataStoreFactory.Create("userSvcUsers", &usertypes.User{})
	if err != nil {
		return nil, err
	}

	authTokensStore, err := options.DataStoreFactory.Create(
		"userSvcAuthTokens",
		&usertypes.AuthToken{},
	)
	if err != nil {
		return nil, err
	}

	credentialsStore, err := options.DataStoreFactory.Create(
		"userSvcCredetentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}

	keyPairsStore, err := options.DataStoreFactory.Create(
		"userSvcKeyPairs",
		&usertypes.KeyPair{},
	)
	if err != nil {
		return nil, err
	}

	contactsStore, err := options.DataStoreFactory.Create(
		"userSvcContacts",
		&usertypes.Contact{},
	)
	if err != nil {
		return nil, err
	}

	enrollsStore, err := options.DataStoreFactory.Create(
		"userSvcEnrolls",
		&usertypes.Enroll{},
	)
	if err != nil {
		return nil, err
	}

	organizationsStore, err := options.DataStoreFactory.Create(
		"userSvcOrganizations",
		&usertypes.Organization{},
	)
	if err != nil {
		return nil, err
	}

	membershipsStore, err := options.DataStoreFactory.Create(
		"userSvcMemberships",
		&usertypes.Membership{},
	)
	if err != nil {
		return nil, err
	}

	permitsStore, err := options.DataStoreFactory.Create(
		"userSvcPermits",
		&usertypes.Permit{},
	)
	if err != nil {
		return nil, err
	}

	passwordsStore, err := options.DataStoreFactory.Create(
		"userSvcPasswords",
		&usertypes.Password{},
	)
	if err != nil {
		return nil, err
	}

	tokenReplacementCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e5,     // number of keys to track frequency (10x max items)
		MaxCost:     1 << 20, // max cost in bytes (~1 MiB)
		BufferItems: 64,      // recommended
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create token replacement cache")
	}

	service := &UserService{
		options:               options,
		usersStore:            usersStore,
		authTokensStore:       authTokensStore,
		credentialsStore:      credentialsStore,
		passwordsStore:        passwordsStore,
		keyPairsStore:         keyPairsStore,
		contactsStore:         contactsStore,
		organizationsStore:    organizationsStore,
		membershipsStore:      membershipsStore,
		permitsStore:          permitsStore,
		enrollsStore:          enrollsStore,
		tokenReplacementCache: tokenReplacementCache,
	}

	return service, nil
}

func (us *UserService) Start() error {
	err := us.registerPermits()
	if err != nil {
		return err
	}

	err = us.bootstrap()
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) RegisterRoutes(router *mux.Router) {
	appl := us.options.Middlewares

	router.HandleFunc("/user-svc/login", appl(func(w http.ResponseWriter, r *http.Request) {
		us.Login(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/self", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ReadSelf(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/users", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ListUsers(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SaveUser(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/self", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SaveSelf(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/change-password", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ChangePassword(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/{userId}/reset-password", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ResetPassword(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/organization", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SaveOrganization(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/organizations", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ListOrganizations(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/organization/{organizationId}/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SaveMembership(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/organization/{organizationId}/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		us.DeleteMembership(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/user-svc/user", appl(func(w http.ResponseWriter, r *http.Request) {
		us.CreateUser(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		us.DeleteUser(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/user-svc/self/has/{permission}", appl(func(w http.ResponseWriter, r *http.Request) {
		us.HasPermission(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ListPermissions(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/register", appl(func(w http.ResponseWriter, r *http.Request) {
		us.Register(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/public-key", appl(func(w http.ResponseWriter, r *http.Request) {
		us.GetPublicKey(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/user-svc/permits", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SavePermits(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/permits", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ListPermits(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/enrolls", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SaveEnrolls(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/enrolls", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ListEnrolls(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/refresh-token", appl(func(w http.ResponseWriter, r *http.Request) {
		us.RefreshToken(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/tokens", appl(func(w http.ResponseWriter, r *http.Request) {
		us.RevokeTokens(w, r)
	})).
		Methods("OPTIONS", "DELETE")
}

func (s *UserService) bootstrap() error {
	// Bootstrapping key pairs

	keyPairs, err := s.keyPairsStore.Query().Find()
	if err != nil {
		return errors.Wrap(err, "failed to query key pairs")
	}

	if len(keyPairs) > 0 {
		kp := keyPairs[0].(*usertypes.KeyPair)
		privKey, err := privateKeyFromString(kp.PrivateKey)
		if err != nil {
			return errors.Wrap(err, "failed to parse private key")
		}
		s.privateKey = privKey
		s.publicKeyPem = kp.PublicKey
	} else {
		bits := 4096
		if s.options.Test {
			bits = 512
		}
		privKey, pubKey, err := generateRSAKeys(bits)

		if err != nil {
			return errors.Wrap(err, "failed to generate RSA keys")
		}
		now := time.Now()
		kp := &usertypes.KeyPair{
			Id:         sdk.Id("keyp"),
			CreatedAt:  now,
			UpdatedAt:  now,
			PublicKey:  pubKey,
			PrivateKey: privKey,
		}
		err = s.keyPairsStore.Upsert(kp)
		if err != nil {
			return errors.Wrap(err, "failed to upsert key pair")
		}

		privKeyTyped, err := privateKeyFromString(kp.PrivateKey)
		if err != nil {
			return errors.Wrap(err, "failed to parse private key")
		}
		s.privateKey = privKeyTyped
		s.publicKeyPem = kp.PublicKey

	}

	// Bootstrapping the admin user. Instead of inefficient role-based queries,
	// we enforce the existence of an admin account with the slug "1backend".
	// If absent, it's created with a default password, which should be updated for security.

	count, err := s.usersStore.Query(
		datastore.Equals([]string{"slug"}, "1backend"),
	).Count()
	if err != nil {
		return errors.Wrap(err, "failed to count users")
	}

	if count == 0 {
		_, err = s.register("1backend", "changeme", "Admin", []string{
			usertypes.RoleAdmin,
		})
		if err != nil {
			return errors.Wrap(err, "failed to register admin user")
		}
	}

	// Bootstrapping credentials

	credentials, err := s.credentialsStore.Query().Find()
	if err != nil {
		return errors.Wrap(err, "failed to query credentials")
	}

	if len(credentials) == 0 {
		slug := "user-svc"
		pw := uuid.NewString()

		err = s.credentialsStore.Upsert(&auth.Credential{
			Slug:     slug,
			Password: pw,
		})
		if err != nil {
			return errors.Wrap(err, "failed to upsert credential")
		}

		tok, err := s.register(slug, pw,
			"User Svc", []string{
				usertypes.RoleUser,
			})
		if err != nil {
			return errors.Wrap(err, "failed to register user-svc")
		}
		s.token = tok.Token
	} else {
		cred := credentials[0].(*auth.Credential)

		tok, err := s.login(&usertypes.LoginRequest{
			Slug:     cred.Slug,
			Password: cred.Password,
		})
		if err != nil {
			return errors.Wrap(err, "failed to login user-svc")
		}

		s.token = tok.Token
	}

	return nil
}
