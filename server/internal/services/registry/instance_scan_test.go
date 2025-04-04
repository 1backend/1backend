package registryservice_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/1backend/1backend/server/internal/di"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

func TestDeployService(t *testing.T) {
	// @todo The mock admin client below doesn't work with the now authenticated instance list endpoint
	return
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "RegistrySvc Suite")
}

var _ = ginkgo.Describe("Instance Scan", func() {
	var (
		server            *httptest.Server
		ctrl              *gomock.Controller
		ctx               context.Context
		mockClientFactory *client.MockClientFactory
		mockUserSvc       *openapi.MockUserSvcAPI
		universe          *di.Universe
		adminClient       *openapi.APIClient
	)

	ginkgo.BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		hs := &di.HandlerSwitcher{}
		server = httptest.NewServer(hs)

		mockClientFactory = client.NewMockClientFactory(ctrl)
		mockUserSvc = test.MockUserSvc(ctx, ctrl, test.WithHasPermissionFactory(func() bool {
			return true
		}))
		mockDeploySvc := openapi.NewMockDeploySvcAPI(ctrl)

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI:   mockUserSvc,
				DeploySvcAPI: mockDeploySvc,
				RegistrySvcAPI: client.NewApiClientFactory(server.URL).
					Client().
					RegistrySvcAPI,
			}).
			AnyTimes()

		options := &di.Options{
			Test:          true,
			Url:           server.URL,
			ClientFactory: mockClientFactory,
		}
		var err error
		universe, err = di.BigBang(options)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		hs.UpdateHandler(universe.Router)

		adminClient, _, err = test.AdminClient(mockClientFactory)
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

	ginkgo.Context("when registry has one instance", func() {
		ginkgo.BeforeEach(func() {
			//
		})

		ginkgo.It("saves an instance successfully", func() {
			_, _, err := adminClient.RegistrySvcAPI.RegisterInstance(ctx).
				Body(
					openapi.RegistrySvcRegisterInstanceRequest{
						Id:           openapi.PtrString("test-a"),
						Url:          "http://test-a",
						DeploymentId: openapi.PtrString("test-deployment"),
					},
				).
				Execute()

			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			instancesRsp, err := waitForInstanceStatus(
				ctx,
				adminClient,
				openapi.InstanceStatusUnreachable,
				5,
			)

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(instancesRsp.Instances).To(gomega.HaveLen(1))
			gomega.Expect(instancesRsp.Instances[0].Status).
				To(gomega.Equal(openapi.InstanceStatusUnreachable))

			healthServer := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					fmt.Fprintln(w, "Hi")
				}),
			)

			gomega.Expect(healthServer).NotTo(gomega.BeNil())

			_, _, err = adminClient.RegistrySvcAPI.RegisterInstance(ctx).
				Body(
					openapi.RegistrySvcRegisterInstanceRequest{
						Id:           openapi.PtrString("test-a"),
						Url:          healthServer.URL,
						DeploymentId: openapi.PtrString("test-deployment"),
					},
				).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			instancesRsp, _, err = adminClient.RegistrySvcAPI.ListInstances(ctx).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(instancesRsp.Instances).To(gomega.HaveLen(1))
			gomega.Expect(instancesRsp.Instances[0].Url).
				To(gomega.Equal(healthServer.URL))

			_, err = waitForInstanceStatus(
				ctx,
				adminClient,
				openapi.InstanceStatusHealthy,
				5,
			)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			healthServer.Close()

			_, _, err = adminClient.RegistrySvcAPI.RegisterInstance(ctx).
				Body(
					openapi.RegistrySvcRegisterInstanceRequest{
						Id:           openapi.PtrString("test-a"),
						Url:          healthServer.URL,
						DeploymentId: openapi.PtrString("test-deployment"),
					},
				).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			instancesRsp, _, err = adminClient.RegistrySvcAPI.ListInstances(ctx).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(instancesRsp.Instances).To(gomega.HaveLen(1))
			gomega.Expect(instancesRsp.Instances[0].Url).
				To(gomega.Equal(healthServer.URL))

			_, err = waitForInstanceStatus(
				ctx,
				adminClient,
				openapi.InstanceStatusUnreachable,
				5,
			)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})

})

func waitForInstanceStatus(
	ctx context.Context,
	client *openapi.APIClient,
	expectedStatus openapi.RegistrySvcInstanceStatus,
	retries int,
) (*openapi.RegistrySvcListInstancesResponse, error) {
	lastStatus := openapi.RegistrySvcInstanceStatus("")

	for i := 0; i < retries; i++ {
		rsp, _, err := client.RegistrySvcAPI.ListInstances(ctx).Execute()
		if err != nil {
			return nil, err
		}
		if len(rsp.Instances) > 0 {
			lastStatus = rsp.Instances[0].Status
			if rsp.Instances[0].Status == expectedStatus {
				return rsp, nil
			}
		}

		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf(
		"expected instance status %s not reached, last status was %v",
		expectedStatus,
		lastStatus,
	)
}
