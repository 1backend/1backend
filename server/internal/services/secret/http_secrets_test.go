/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package secretservice_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"go.uber.org/mock/gomock"
)

func TestSecretService(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SecretService Suite")
}

var _ = ginkgo.Describe("Secret Tests", func() {
	var (
		server     *httptest.Server
		ctrl       *gomock.Controller
		ctx        context.Context
		userClient *openapi.APIClient

		mockClientFactory *sdk.MockClientFactory
		mockUserSvc       *openapi.MockUserSvcAPI

		universe    *mux.Router
		starterFunc func() error

		isAdmin      bool
		isAuthorized bool
		userSlug     string
	)

	ginkgo.BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		hs := &di.HandlerSwitcher{}
		server = httptest.NewServer(hs)

		mockClientFactory = sdk.NewMockClientFactory(ctrl)

		mockUserSvc = test.MockUserSvc(
			ctx,
			ctrl,
			test.WithIsAuthorizedFactory(func() bool {
				return isAuthorized
			}),
			test.WithSlugFactory(func() string {
				return userSlug
			}),
		)
		mockAuthorizer := sdk.NewMockAuthorizer(ctrl)
		mockAuthorizer.EXPECT().
			IsAdminFromRequest(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ string, _ *http.Request) (bool, error) {
				return isAdmin, nil
			}).AnyTimes()

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI: mockUserSvc,
				SecretSvcAPI: sdk.NewApiClientFactory(server.URL).
					Client().
					SecretSvcAPI,
			}).
			AnyTimes()

		options := &di.Options{
			Test:          true,
			Url:           server.URL,
			Authorizer:    mockAuthorizer,
			ClientFactory: mockClientFactory,
		}

		var err error
		universe, starterFunc, err = di.BigBang(options)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		hs.UpdateHandler(universe)

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.JustBeforeEach(func() {
		err := starterFunc()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.AfterEach(func() {
		server.Close()
		ctrl.Finish()
	})

	ginkgo.Context("read non-existent", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			isAuthorized = true
			isAdmin = false
			userSlug = "test-user-1"
		})

		ginkgo.It("works", func() {
			readRsp, _, err := userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Key: openapi.PtrString("nonexistent"),
				}).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(0))
		})
	})

	ginkgo.Context("as a user trying to claim non-prefixed", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			isAuthorized = true
			isAdmin = false
			userSlug = "test-user-1"
		})

		ginkgo.It("fails", func() {
			_, _, err := userClient.SecretSvcAPI.SaveSecrets(ctx).
				Body(openapi.SecretSvcSaveSecretsRequest{
					Secrets: []openapi.SecretSvcSecret{
						{
							Key:   openapi.PtrString("non-prefixed"),
							Value: openapi.PtrString("value"),
						}},
				}).
				Execute()

			gomega.Expect(err).To(gomega.HaveOccurred())
		})
	})

	ginkgo.Context("as an admin trying to claim non-prefixed", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			isAuthorized = true
			isAdmin = true
			userSlug = "test-admin-user-1"
		})

		ginkgo.It("passes", func() {
			_, _, err := userClient.SecretSvcAPI.SaveSecrets(ctx).
				Body(openapi.SecretSvcSaveSecretsRequest{
					Secrets: []openapi.SecretSvcSecret{
						{
							Key:   openapi.PtrString("non-prefixed"),
							Value: openapi.PtrString("value"),
						}},
				}).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			readRsp, _, err := userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Key: openapi.PtrString("non-prefixed"),
				}).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(1))

			gomega.Expect(readRsp.Secrets[0].Readers).To(gomega.BeEmpty())
			gomega.Expect(readRsp.Secrets[0].Writers).To(gomega.BeEmpty())
			gomega.Expect(readRsp.Secrets[0].Deleters).To(gomega.BeEmpty())
		})
	})

	ginkgo.Context("user 1 can write and read, but user 2 will not have access", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			isAuthorized = true
			isAdmin = false
			userSlug = "test-user-1"
		})

		ginkgo.It("works", func() {
			readRsp, _, err := userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Key: openapi.PtrString("test-user-1/non-existent"),
				}).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(0))

			_, _, err = userClient.SecretSvcAPI.SaveSecrets(ctx).
				Body(openapi.SecretSvcSaveSecretsRequest{
					Secrets: []openapi.SecretSvcSecret{
						{
							Key:   openapi.PtrString("test-user-1/non-existent"),
							Value: openapi.PtrString("value"),
						}},
				}).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			readRsp, _, err = userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Key: openapi.PtrString("test-user-1/non-existent"),
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(1))

			userSlug = "test-user-2"

			readRsp, _, err = userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Key: openapi.PtrString("test-user-1/non-existent"),
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(0))
		})
	})

	ginkgo.Context("as a user writing already encrypted values", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			isAuthorized = true
			isAdmin = false
			userSlug = "test-user-1"
		})

		ginkgo.It("works", func() {
			rsp, _, err := userClient.SecretSvcAPI.EncryptValue(ctx).
				Body(openapi.SecretSvcEncryptValueRequest{
					Value: openapi.PtrString("value"),
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			_, _, err = userClient.SecretSvcAPI.SaveSecrets(ctx).
				Body(openapi.SecretSvcSaveSecretsRequest{
					Secrets: []openapi.SecretSvcSecret{
						{
							Key:       openapi.PtrString("test-user-1/enc1"),
							Value:     rsp.Value,
							Encrypted: openapi.PtrBool(true),
						}},
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			readRsp, _, err := userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Key: openapi.PtrString("test-user-1/enc1"),
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(1))
			gomega.Expect(*readRsp.Secrets[0].Value).To(gomega.Equal("value"))
		})
	})

	ginkgo.Context("list multiple", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			isAuthorized = true
			isAdmin = true
			userSlug = "test-admin-user-1"
		})

		ginkgo.It("works", func() {
			_, _, err := userClient.SecretSvcAPI.SaveSecrets(ctx).
				Body(openapi.SecretSvcSaveSecretsRequest{
					Secrets: []openapi.SecretSvcSecret{
						{
							Key:   openapi.PtrString("a"),
							Value: openapi.PtrString("value"),
						}},
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			_, _, err = userClient.SecretSvcAPI.SaveSecrets(ctx).
				Body(openapi.SecretSvcSaveSecretsRequest{
					Secrets: []openapi.SecretSvcSecret{
						{
							Key:   openapi.PtrString("b"),
							Value: openapi.PtrString("value"),
						}},
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			readRsp, _, err := userClient.SecretSvcAPI.ListSecrets(ctx).
				Body(openapi.SecretSvcListSecretsRequest{
					Keys: []string{"a", "b"},
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Secrets)).To(gomega.Equal(2))
			gomega.Expect(*readRsp.Secrets[0].Value).To(gomega.Equal("value"))
		})
	})
})
