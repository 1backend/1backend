package deployservice_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"go.uber.org/mock/gomock"
)

func TestDeployService(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "DeployService Suite")
}

var _ = ginkgo.Describe("Deploy Loop", func() {
	var (
		server               *httptest.Server
		ctrl                 *gomock.Controller
		ctx                  context.Context
		mockClientFactory    *sdk.MockClientFactory
		mockUserSvc          *openapi.MockUserSvcAPI
		universe             *mux.Router
		mockRegistrySvc      *openapi.MockRegistrySvcAPI
		mockContainerSvc     *openapi.MockContainerSvcAPI
		starterFunc          func() error
		adminClient          *openapi.APIClient
		launchContainerError error

		nodes       []openapi.RegistrySvcNode
		instances   []openapi.RegistrySvcInstance
		definitions []openapi.RegistrySvcDefinition
	)

	ginkgo.BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(ginkgo.GinkgoT())
		hs := &di.HandlerSwitcher{}
		server = httptest.NewServer(hs)

		mockClientFactory = sdk.NewMockClientFactory(ctrl)
		mockUserSvc = test.MockUserSvc(ctx, ctrl, test.WithIsAuthorizedFactory(func() bool {
			return true
		}))
		mockRegistrySvc = openapi.NewMockRegistrySvcAPI(ctrl)
		mockContainerSvc = openapi.NewMockContainerSvcAPI(ctrl)

		mockClientFactory.EXPECT().
			Client(gomock.Any()).
			Return(&openapi.APIClient{
				UserSvcAPI:      mockUserSvc,
				RegistrySvcAPI:  mockRegistrySvc,
				ContainerSvcAPI: mockContainerSvc,
				DeploySvcAPI: sdk.NewApiClientFactory(server.URL).
					Client().
					DeploySvcAPI,
			}).
			AnyTimes()

		options := &di.Options{
			Test:          true,
			Url:           server.URL,
			ClientFactory: mockClientFactory,
		}
		var err error
		universe, starterFunc, err = di.BigBang(options)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		hs.UpdateHandler(universe)

		adminClient, _, err = test.AdminClient(mockClientFactory)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

	})

	ginkgo.JustBeforeEach(func() {
		mockListNodesRequest := openapi.ApiListNodesRequest{
			ApiService: mockRegistrySvc,
		}
		mockRegistrySvc.EXPECT().
			ListNodes(ctx).
			Return(mockListNodesRequest).
			AnyTimes()
		mockRegistrySvc.EXPECT().ListNodesExecute(gomock.Any()).Return(
			&openapi.RegistrySvcListNodesResponse{
				Nodes: nodes,
			}, nil, nil,
		).AnyTimes()

		mockListInstancesRequest := openapi.ApiListInstancesRequest{
			ApiService: mockRegistrySvc,
		}
		mockRegistrySvc.EXPECT().
			ListInstances(ctx).
			Return(mockListInstancesRequest).
			AnyTimes()
		mockRegistrySvc.EXPECT().ListInstancesExecute(gomock.Any()).Return(
			&openapi.RegistrySvcListInstancesResponse{
				Instances: instances,
			}, nil, nil,
		).AnyTimes()

		mockListDefinitionsRequest := openapi.ApiListDefinitionsRequest{
			ApiService: mockRegistrySvc,
		}
		mockRegistrySvc.EXPECT().
			ListDefinitions(ctx).
			Return(mockListDefinitionsRequest).
			AnyTimes()
		mockRegistrySvc.EXPECT().ListDefinitionsExecute(gomock.Any()).Return(
			&openapi.RegistrySvcListDefinitionsResponse{
				Definitions: definitions,
			}, nil, nil,
		).AnyTimes()

		mockLaunchContainerRequest := openapi.ApiRunContainerRequest{
			ApiService: mockContainerSvc,
		}
		mockContainerSvc.EXPECT().
			RunContainer(ctx).
			Return(mockLaunchContainerRequest).
			AnyTimes()
		mockContainerSvc.EXPECT().
			RunContainerExecute(gomock.Any()).
			Return(nil, nil, launchContainerError).
			AnyTimes()

		mockRegisterInstanceRequest := openapi.ApiRegisterInstanceRequest{
			ApiService: mockRegistrySvc,
		}
		mockRegistrySvc.EXPECT().
			RegisterInstance(ctx).
			Return(mockRegisterInstanceRequest).
			AnyTimes()
		mockRegistrySvc.EXPECT().
			RegisterInstanceExecute(gomock.Any()).
			Return(nil, nil, nil).
			AnyTimes()

		err := starterFunc()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	ginkgo.AfterEach(func() {
		server.Close()
		ctrl.Finish()
	})

	ginkgo.Context("when registry has one active node", func() {
		ginkgo.BeforeEach(func() {
			nodes = []openapi.RegistrySvcNode{
				{
					Url: server.URL,
				},
			}
			definitions = []openapi.RegistrySvcDefinition{
				{
					Id: "test-a",
					Image: &openapi.RegistrySvcImageSpec{
						Name:          "hashicorp/http-echo",
						InternalPorts: []int32{8080},
					},
					Ports: []openapi.RegistrySvcPortMapping{
						{
							Internal: 8080,
							Host:     8887,
						},
					},
				},
			}
		})

		ginkgo.It("saves a deployment successfully", func() {
			_, _, err := adminClient.DeploySvcAPI.SaveDeployment(ctx).
				Body(openapi.DeploySvcSaveDeploymentRequest{
					Deployment: &openapi.DeploySvcDeployment{
						DefinitionId: "test-a",
					},
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			rsp, err := waitForDeploymentStatus(
				ctx,
				adminClient,
				openapi.DeploymentStatusOK,
				5,
			)

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(rsp.Deployments).To(gomega.HaveLen(1))
			gomega.Expect(rsp.Deployments[0].DefinitionId).
				To(gomega.Equal("test-a"))
		})
	})

	ginkgo.When("docker launch container fails", func() {
		ginkgo.BeforeEach(func() {
			launchContainerError = fmt.Errorf("Internal Server Error")
			nodes = []openapi.RegistrySvcNode{
				{
					Url: server.URL,
				},
			}
			definitions = []openapi.RegistrySvcDefinition{
				{
					Id: "test-a",
					Image: &openapi.RegistrySvcImageSpec{
						Name:          "hashicorp/http-echo",
						InternalPorts: []int32{8080},
					},
					Ports: []openapi.RegistrySvcPortMapping{
						{
							Internal: 8080,
							Host:     8887,
						},
					},
				},
			}
		})

		ginkgo.It("status of the deployment is error", func() {
			_, _, err := adminClient.DeploySvcAPI.SaveDeployment(ctx).
				Body(openapi.DeploySvcSaveDeploymentRequest{
					Deployment: &openapi.DeploySvcDeployment{
						DefinitionId: "test-a",
					},
				}).
				Execute()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			rsp, err := waitForDeploymentStatus(
				ctx,
				adminClient,
				openapi.DeploymentStatusError,
				5,
			)

			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(rsp.Deployments).To(gomega.HaveLen(1))
			gomega.Expect(rsp.Deployments[0].DefinitionId).
				To(gomega.Equal("test-a"))
			gomega.Expect(*rsp.Deployments[0].Details).
				To(gomega.Equal("Internal Server Error"))
		})
	})
})

func waitForDeploymentStatus(
	ctx context.Context,
	client *openapi.APIClient,
	expectedStatus openapi.DeploySvcDeploymentStatus,
	retries int,
) (*openapi.DeploySvcListDeploymentsResponse, error) {
	for i := 0; i < retries; i++ {
		rsp, _, err := client.DeploySvcAPI.ListDeployments(ctx).Execute()
		if err != nil {
			return nil, err
		}
		if len(rsp.Deployments) > 0 &&
			*rsp.Deployments[0].Status == expectedStatus {
			return rsp, nil
		}

		time.Sleep(time.Second)
	}

	return nil, fmt.Errorf(
		"expected deployment status %s not reached",
		expectedStatus,
	)
}
