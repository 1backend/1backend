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
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/dgraph-io/ristretto"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

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

	userStore         datastore.DataStore
	credentialStore   datastore.DataStore
	passwordStore     datastore.DataStore
	tokenStore        datastore.DataStore
	keyPairStore      datastore.DataStore
	contactStore      datastore.DataStore
	organizationStore datastore.DataStore
	membershipStore   datastore.DataStore
	permitStore       datastore.DataStore
	enrollStore       datastore.DataStore
	appStore          datastore.DataStore
	otpStore          datastore.DataStore

	privateKey   *rsa.PrivateKey
	publicKeyPem string

	configCache map[string]any

	tokenReplacementCache *ristretto.Cache
	refreshGroup          singleflight.Group
}

func NewUserService(
	options *universe.Options,
) (*UserService, error) {
	usersStore, err := options.DataStoreFactory.Create("userSvcUsers", &usertypes.User{})
	if err != nil {
		return nil, err
	}

	authTokensStore, err := options.DataStoreFactory.Create(
		"userSvcTokens",
		&usertypes.Token{},
	)
	if err != nil {
		return nil, err
	}

	credentialsStore, err := options.DataStoreFactory.Create(
		"userSvcCredentials",
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

	appStore, err := options.DataStoreFactory.Create(
		"userSvcApps",
		&usertypes.App{},
	)
	if err != nil {
		return nil, err
	}

	otpStore, err := options.DataStoreFactory.Create(
		"userSvcOtps",
		&usertypes.OTP{},
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
		userStore:             usersStore,
		tokenStore:            authTokensStore,
		credentialStore:       credentialsStore,
		passwordStore:         passwordsStore,
		keyPairStore:          keyPairsStore,
		contactStore:          contactsStore,
		organizationStore:     organizationsStore,
		membershipStore:       membershipsStore,
		permitStore:           permitsStore,
		enrollStore:           enrollsStore,
		appStore:              appStore,
		otpStore:              otpStore,
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

	router.HandleFunc("/user-svc/token/exchange", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ExchangeToken(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/tokens", appl(func(w http.ResponseWriter, r *http.Request) {
		us.RevokeTokens(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/user-svc/apps", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ListApps(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/app", appl(func(w http.ResponseWriter, r *http.Request) {
		us.ReadApp(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/app", appl(func(w http.ResponseWriter, r *http.Request) {
		us.UpdateApp(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/user-svc/otp/send", appl(func(w http.ResponseWriter, r *http.Request) {
		us.SendOTP(w, r)
	})).
		Methods("OPTIONS", "POST")
}

func (s *UserService) bootstrap() error {
	// Bootstrapping key pairs

	keyPairs, err := s.keyPairStore.Query().Find()
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
		var (
			privKey string
			pubKey  string
		)

		if s.options.Test {
			privKey, pubKey, err = parseTestKey()
			if err != nil {
				return errors.Wrap(err, "failed to parse test RSA key")
			}
		} else {
			bits := 4096
			privKey, pubKey, err = generateRSAKeys(bits)

			if err != nil {
				return errors.Wrap(err, "failed to generate RSA keys")
			}
		}

		now := time.Now()
		kp := &usertypes.KeyPair{
			Id:         sdk.Id("keyp"),
			CreatedAt:  now,
			UpdatedAt:  now,
			PublicKey:  pubKey,
			PrivateKey: privKey,
		}
		err = s.keyPairStore.Upsert(kp)
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

	count, err := s.userStore.Query(
		datastore.Equals([]string{"slug"}, "1backend"),
	).Count()
	if err != nil {
		return errors.Wrap(err, "failed to count users")
	}

	if count == 0 {
		_, err = s.register(
			usertypes.DefaultAppHost,
			"1backend",
			"changeme",
			"Admin", []string{
				usertypes.RoleAdmin,
			})
		if err != nil {
			return errors.Wrap(err, "failed to register admin user")
		}
	}

	// Bootstrapping credentials

	credentials, err := s.credentialStore.Query().Find()
	if err != nil {
		return errors.Wrap(err, "failed to query credentials")
	}

	if len(credentials) == 0 {
		slug := "user-svc"
		pw := uuid.NewString()

		err = s.credentialStore.Upsert(&auth.Credential{
			Slug:     slug,
			Password: pw,
		})
		if err != nil {
			return errors.Wrap(err, "failed to upsert credential")
		}

		tok, err := s.register(
			usertypes.DefaultAppHost,
			slug,
			pw,
			"User Svc", []string{
				usertypes.RoleUser,
			})
		if err != nil {
			return errors.Wrap(err, "failed to register user-svc")
		}
		s.token = tok.Token
	} else {
		cred := credentials[0].(*auth.Credential)

		app, err := s.getOrCreateApp(usertypes.DefaultAppHost)
		if err != nil {
			return errors.Wrap(err, "failed to get or create app")
		}

		tok, err := s.login(
			app.Id,
			&usertypes.LoginRequest{
				Slug:     cred.Slug,
				Password: cred.Password,
			},
			false)
		if err != nil {
			return errors.Wrap(err, "failed to login user-svc")
		}

		s.token = tok.Token
	}

	return nil
}

// RFC 9500 pregenerated test-only 2048-bit key
const testRSA2048PEM = `-----BEGIN RSA TESTING KEY-----
MIIEowIBAAKCAQEAsPnoGUOnrpiSqt4XynxA+HRP7S+BSObI6qJ7fQAVSPtRkqso
tWxQYLEYzNEx5ZSHTGypibVsJylvCfuToDTfMul8b/CZjP2Ob0LdpYrNH6l5hvFE
89FU1nZQF15oVLOpUgA7wGiHuEVawrGfey92UE68mOyUVXGweJIVDdxqdMoPvNNU
l86BU02vlBiESxOuox+dWmuVV7vfYZ79Toh/LUK43YvJh+rhv4nKuF7iHjVjBd9s
B6iDjj70HFldzOQ9r8SRI+9NirupPTkF5AKNe6kUhKJ1luB7S27ZkvB3tSTT3P59
3VVJvnzOjaA1z6Cz+4+eRvcysqhrRgFlwI9TEwIDAQABAoIBAEEYiyDP29vCzx/+
dS3LqnI5BjUuJhXUnc6AWX/PCgVAO+8A+gZRgvct7PtZb0sM6P9ZcLrweomlGezI
FrL0/6xQaa8bBr/ve/a8155OgcjFo6fZEw3Dz7ra5fbSiPmu4/b/kvrg+Br1l77J
aun6uUAs1f5B9wW+vbR7tzbT/mxaUeDiBzKpe15GwcvbJtdIVMa2YErtRjc1/5B2
BGVXyvlJv0SIlcIEMsHgnAFOp1ZgQ08aDzvilLq8XVMOahAhP1O2A3X8hKdXPyrx
IVWE9bS9ptTo+eF6eNl+d7htpKGEZHUxinoQpWEBTv+iOoHsVunkEJ3vjLP3lyI/
fY0NQ1ECgYEA3RBXAjgvIys2gfU3keImF8e/TprLge1I2vbWmV2j6rZCg5r/AS0u
pii5CvJ5/T5vfJPNgPBy8B/yRDs+6PJO1GmnlhOkG9JAIPkv0RBZvR0PMBtbp6nT
Y3yo1lwamBVBfY6rc0sLTzosZh2aGoLzrHNMQFMGaauORzBFpY5lU50CgYEAzPHl
u5DI6Xgep1vr8QvCUuEesCOgJg8Yh1UqVoY/SmQh6MYAv1I9bLGwrb3WW/7kqIoD
fj0aQV5buVZI2loMomtU9KY5SFIsPV+JuUpy7/+VE01ZQM5FdY8wiYCQiVZYju9X
Wz5LxMNoz+gT7pwlLCsC4N+R8aoBk404aF1gum8CgYAJ7VTq7Zj4TFV7Soa/T1eE
k9y8a+kdoYk3BASpCHJ29M5R2KEA7YV9wrBklHTz8VzSTFTbKHEQ5W5csAhoL5Fo
qoHzFFi3Qx7MHESQb9qHyolHEMNx6QdsHUn7rlEnaTTyrXh3ifQtD6C0yTmFXUIS
CW9wKApOrnyKJ9nI0HcuZQKBgQCMtoV6e9VGX4AEfpuHvAAnMYQFgeBiYTkBKltQ
XwozhH63uMMomUmtSG87Sz1TmrXadjAhy8gsG6I0pWaN7QgBuFnzQ/HOkwTm+qKw
AsrZt4zeXNwsH7QXHEJCFnCmqw9QzEoZTrNtHJHpNboBuVnYcoueZEJrP8OnUG3r
UjmopwKBgAqB2KYYMUqAOvYcBnEfLDmyZv9BTVNHbR2lKkMYqv5LlvDaBxVfilE0
2riO4p6BaAdvzXjKeRrGNEKoHNBpOSfYCOM16NjL8hIZB1CaV3WbT5oY+jp7Mzd5
7d56RZOE+ERK2uz/7JX9VSsM/LbH9pJibd4e8mikDS9ntciqOH/3
-----END RSA TESTING KEY-----`

func parseTestKey() (privateKeyPem, publicKeyPem string, err error) {
	block, _ := pem.Decode([]byte(strings.ReplaceAll(
		testRSA2048PEM,
		"TESTING KEY", "PRIVATE KEY",
	)))
	if block == nil {
		return "", "", fmt.Errorf("failed to decode test RSA key")
	}

	// Private key PEM is literally the block we decoded
	privateKeyPem = string(pem.EncodeToMemory(block))

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", "", err
	}

	// Marshal public key into PEM
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyPem = string(pem.EncodeToMemory(publicKeyBlock))

	return privateKeyPem, publicKeyPem, nil
}
