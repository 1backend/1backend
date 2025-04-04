/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package configservice_test

import (
	"context"
	"net/http/httptest"
	"testing"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

func TestConfigService(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ConfigService Suite")
}

var _ = ginkgo.Describe("Config Tests", func() {
	var (
		server     *httptest.Server
		ctrl       *gomock.Controller
		ctx        context.Context
		userClient *openapi.APIClient

		mockClientFactory *client.MockClientFactory
		mockUserSvc       *openapi.MockUserSvcAPI
		mockFirehoseSvc   *openapi.MockFirehoseSvcAPI

		universe *di.Universe

		hasPermission bool
		isAdmin       bool
		userSlug      string
	)

	ginkgo.BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		hs := &di.HandlerSwitcher{}
		server = httptest.NewServer(hs)

		mockClientFactory = client.NewMockClientFactory(ctrl)

		mockUserSvc = test.MockUserSvc(
			ctx,
			ctrl,
			test.WithHasPermissionFactory(func() bool {
				return hasPermission
			}),
			test.WithSlugFactory(func() string {
				return userSlug
			}),
		)
		mockAuthorizer := auth.NewMockAuthorizer(ctrl)
		mockAuthorizer.EXPECT().
			IsAdminFromRequest(gomock.Any(), gomock.Any()).
			Return(isAdmin, nil).AnyTimes()

		mockFirehoseSvc = openapi.NewMockFirehoseSvcAPI(ctrl)

		mockFirehoseSvc.EXPECT().
			PublishEvent(gomock.Any()).
			Return(openapi.ApiPublishEventRequest{
				ApiService: mockFirehoseSvc,
			}).AnyTimes()
		mockFirehoseSvc.EXPECT().
			PublishEventExecute(gomock.Any()).
			Return(nil, nil).AnyTimes()

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI:     mockUserSvc,
				FirehoseSvcAPI: mockFirehoseSvc,
				ConfigSvcAPI: client.NewApiClientFactory(server.URL).
					Client().
					ConfigSvcAPI,
			}).
			AnyTimes()

		options := &di.Options{
			Test:          true,
			Url:           server.URL,
			Authorizer:    mockAuthorizer,
			ClientFactory: mockClientFactory,
		}

		var err error
		universe, err = di.BigBang(options)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		hs.UpdateHandler(universe.Router)

		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.JustBeforeEach(func() {
		err := universe.StarterFunc()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.AfterEach(func() {
		server.Close()
		ctrl.Finish()
	})

	ginkgo.Context("get", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()
		})

		ginkgo.It("publicly readable", func() {
			readRsp, _, err := userClient.ConfigSvcAPI.GetConfig(ctx).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(readRsp.Config.Data)).To(gomega.Equal(3))
		})
	})

	ginkgo.Context("save", func() {
		ginkgo.BeforeEach(func() {
			userClient = mockClientFactory.Client()

			hasPermission = true
			isAdmin = false
			userSlug = "test-user-1"
		})

		ginkgo.It("works", func() {
			_, _, err := userClient.ConfigSvcAPI.SaveConfig(ctx).
				Body(openapi.ConfigSvcSaveConfigRequest{
					Config: &openapi.ConfigSvcConfig{
						Data: map[string]interface{}{
							userSlug:       "test",
							"someOtherKey": "someOtherValue",
						},
					},
				}).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			readRsp, _, err := userClient.ConfigSvcAPI.GetConfig(ctx).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(readRsp.Config.Data[userSlug]).To(gomega.Equal("test"))
			gomega.Expect(readRsp.Config.Data["someOtherKey"]).To(gomega.BeNil())
		})
	})
})
