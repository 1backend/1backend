package di

// This is some of the cruftiest files in the system.

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	pglock "github.com/1backend/1backend/sdk/go/lock/pg"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/1backend/1backend/sdk/go/lock"
	distlock "github.com/1backend/1backend/sdk/go/lock/local"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/middlewares"

	"github.com/1backend/1backend/server/internal/clients/llamacpp"
	"github.com/1backend/1backend/server/internal/router"
	chatservice "github.com/1backend/1backend/server/internal/services/chat"
	configservice "github.com/1backend/1backend/server/internal/services/config"
	containerservice "github.com/1backend/1backend/server/internal/services/container"
	dataservice "github.com/1backend/1backend/server/internal/services/data"
	deployservice "github.com/1backend/1backend/server/internal/services/deploy"
	emailservice "github.com/1backend/1backend/server/internal/services/email"
	fileservice "github.com/1backend/1backend/server/internal/services/file"
	firehoseservice "github.com/1backend/1backend/server/internal/services/firehose"
	modelservice "github.com/1backend/1backend/server/internal/services/model"
	policyservice "github.com/1backend/1backend/server/internal/services/policy"
	promptservice "github.com/1backend/1backend/server/internal/services/prompt"
	proxyservice "github.com/1backend/1backend/server/internal/services/proxy"
	registryservice "github.com/1backend/1backend/server/internal/services/registry"
	secretservice "github.com/1backend/1backend/server/internal/services/secret"
	sourceservice "github.com/1backend/1backend/server/internal/services/source"
	userservice "github.com/1backend/1backend/server/internal/services/user"
)

type Options struct {
	Port        int
	GpuPlatform string

	Az         string
	Region     string
	LLMHost    string
	VolumeName string
	ConfigPath string

	// eg. mysql, postgres
	Db string

	// Connection string eg.
	// "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	DbConnectionString string

	// Crucial for distributed features.
	// Please see the documentation for the envar OB_NODE_ID
	NodeId string

	// DbPrefix allows us to have isolated envs for different test cases
	// but still make multiple nodes in those test cases use the same
	// shard of the db.
	DbPrefix string

	SourceControlToken  string
	SecretEncryptionKey string

	// URL of the local 1Backend server instance
	Url string

	// Test mode if true will cause the localstore to
	// save data into random temporary folders.
	Test bool

	// Lock is a distributed lock. Use this when you want to synronize
	// across service instances/nodes.
	// eg: leader election
	Lock lock.DistributedLock

	LLamaCppClient llamacpp.ClientI

	// DataStoreFactory can create database tables
	DataStoreFactory infra.DataStoreFactory

	// HomeDir is the 1Backend config/data/uploads/downloads directory.
	// For tests it's something like /tmp/1backend-2698538720/
	// For live it's /home/youruser/.1backend
	HomeDir string

	// ClientFactory is used for service to service communication
	// ie. this is how services call each other
	ClientFactory client.ClientFactory

	// Authorizer is a helper interface that contains
	// auth related utility functions
	Authorizer auth.Authorizer
}

type Universe struct {
	Options       Options
	Router        *mux.Router
	StarterFunc   func() error
	ClientFactory client.ClientFactory
}

func BigBang(options *Options) (*Universe, error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic in BinBang",
				slog.String("error", fmt.Sprintf("%v", r)),
				slog.String("trace", string(debug.Stack())),
			)
			os.Exit(1)
		}
	}()

	// @todo GPU platform maybe this could be autodetected
	if options.GpuPlatform == "" {
		options.GpuPlatform = os.Getenv("OB_GPU_PLATFORM")
	}
	if options.Url == "" {
		options.Url = os.Getenv("OB_SELF_URL")
	}
	if options.Url == "" {
		options.Url = router.SelfAddress()
	} else {
		if !strings.HasPrefix(options.Url, "http") {
			options.Url = fmt.Sprintf("http://%v", options.Url)
		}
		uri, err := url.Parse(options.Url)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse url")
		}
		p, err := strconv.ParseInt(uri.Port(), 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse port")
		}
		router.SetPort(int(p))
	}

	if options.NodeId == "" {
		options.NodeId = os.Getenv("OB_NODE_ID")
	}
	if options.Az == "" {
		options.Az = os.Getenv("OB_AZ")
	}
	if options.Region == "" {
		options.Region = os.Getenv("OB_AZ")
	}
	if options.LLMHost == "" {
		options.LLMHost = os.Getenv("OB_LLM_HOST")
	}
	if options.VolumeName == "" {
		options.VolumeName = os.Getenv("OB_VOLUME_NAME")
	}
	if options.DbPrefix == "" {
		options.DbPrefix = os.Getenv("OB_DB_PREFIX")
	}
	if options.Db == "" {
		options.Db = os.Getenv("OB_DB")
	}
	if options.DbConnectionString == "" {
		options.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}
	if options.SecretEncryptionKey == "" {
		options.SecretEncryptionKey = os.Getenv("OB_ENCRYPTION_KEY")
		if options.SecretEncryptionKey == "" {
			options.SecretEncryptionKey = "changeMeToSomethingSecureForReal"
		}
	}

	if !options.Test && os.Getenv("OB_TEST") == "true" {
		options.Test = true
	}

	homeDir, err := infra.HomeDir(infra.HomeDirOptions{
		Test:         options.Test,
		ConfigFolder: os.Getenv("OB_CONFIG_FOLDER"),
	})
	if err != nil {
		return nil, err
	}

	options.HomeDir = homeDir

	if options.Lock == nil {
		options.Lock = distlock.NewLocalDistributedLock()
	}

	if options.Authorizer == nil {
		options.Authorizer = auth.AuthorizerImpl{}
	}

	configService, err := configservice.NewConfigService(
		options.Lock,
		options.Authorizer,
		options.HomeDir,
	)
	if err != nil {
		logger.Error(
			"Config service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	if options.DataStoreFactory == nil {
		dc, err := infra.NewDataStoreFactory(infra.DataStoreConfig{
			Test:               options.Test,
			TablePrefix:        options.DbPrefix,
			Db:                 options.Db,
			DbConnectionString: options.DbConnectionString,
		})
		if err != nil {
			return nil, err
		}
		options.DataStoreFactory = dc

		if options.Db != "" {
			dbHandle, err := dc.Handle()
			if err != nil {
				logger.Error(
					"Failed to get DB handle",
					slog.String("error", err.Error()),
				)
				os.Exit(1)
			}
			conn, err := dbHandle.(*sql.DB).Conn(context.Background())
			if err != nil {
				logger.Error(
					"Failed to get DB connection",
					slog.String("error", err.Error()),
				)
				os.Exit(1)
			}
			options.Lock = pglock.NewPGDistributedLock(conn)
		}
	}

	if options.ClientFactory == nil {
		options.ClientFactory = client.NewApiClientFactory(options.Url)
	}

	configService.SetDataStoreFactory(options.DataStoreFactory.Create)

	configService.SetClientFactory(options.ClientFactory)

	userService, err := userservice.NewUserService(
		options.ClientFactory,
		options.Authorizer,
		options.DataStoreFactory.Create,
		options.Test,
	)
	if err != nil {
		logger.Error(
			"User service start failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	if err != nil {
		logger.Error(
			"Config service start failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	firehoseService, err := firehoseservice.NewFirehoseService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Firehose service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	err = os.MkdirAll(options.HomeDir, 0755)
	if err != nil {
		logger.Error(
			"Config folder creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	fileService, err := fileservice.NewFileService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
		options.HomeDir,
	)
	if err != nil {
		logger.Error(
			"Download service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	containerService, err := containerservice.NewContainerService(
		options.VolumeName,
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Container service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	modelService, err := modelservice.NewModelService(
		// @todo GPU platform maybe this could be autodetected
		options.GpuPlatform,
		options.LLMHost,
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Model service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	chatService, err := chatservice.NewChatService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Chat service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	promptService, err := promptservice.NewPromptService(
		options.ClientFactory,
		options.LLamaCppClient,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Prompt service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	dataService, err := dataservice.NewDataService(
		options.ClientFactory,
		options.Lock,
		options.Authorizer,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Data service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	policyService, err := policyservice.NewPolicyService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Policy service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	registryService, err := registryservice.NewRegistryService(
		options.Url,
		options.Az,
		options.Region,
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
		options.NodeId,
	)
	if err != nil {
		logger.Error(
			"Node service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	deployService, err := deployservice.NewDeployService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
		options.Test,
	)
	if err != nil {
		logger.Error(
			"Node service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	sourceService, err := sourceservice.NewSourceService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Source service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	if options.SecretEncryptionKey == "" {
		options.SecretEncryptionKey = "changeMeToSomethingSecureForReal"
	}
	secretService, err := secretservice.NewSecretService(
		options.ClientFactory,
		options.Authorizer,
		options.Lock,
		options.DataStoreFactory.Create,
		options.SecretEncryptionKey,
	)
	if err != nil {
		logger.Error(
			"Secret service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	proxyService, err := proxyservice.NewProxyService(
		options.ClientFactory,
		options.Authorizer,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Proxy service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	emailService, err := emailservice.NewEmailService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		logger.Error(
			"Email service creation failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	mws := []middlewares.Middleware{
		middlewares.ThrottledLogger,
		middlewares.Recover,
		middlewares.CORS,
		middlewares.GzipDecodeMiddleware,
	}
	appl := applicator(mws)

	router := mux.NewRouter().SkipClean(true).UseEncodedPath()

	router.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found"))
		},
	)

	router.HandleFunc("/firehose-svc/events/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseService.Subscribe(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/firehose-svc/event", appl(func(w http.ResponseWriter, r *http.Request) {
		firehoseService.Publish(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/file-svc/download", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.Download(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/download/{url}/pause", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.PauseDownload(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/file-svc/download/{url}", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.GetDownload(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/file-svc/downloads", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.ListDownloads(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/file-svc/upload", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.UploadFile(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/file-svc/uploads", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.ListUploads(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/file-svc/serve/upload/{fileId}", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.ServeUpload(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/file-svc/serve/download/{url}", appl(func(w http.ResponseWriter, r *http.Request) {
		fileService.ServeDownload(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/container-svc/daemon/info", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.DaemonInfo(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/container-svc/image/{imageName}/pullable", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.ImagePullable(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/container-svc/host", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.Host(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/container-svc/logs", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.ListLogs(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/container-svc/containers", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.ListContainers(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/container-svc/container", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.RunContainer(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/container-svc/image", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.BuildImage(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/container-svc/container/stop", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.StopContainer(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/container-svc/container/is-running", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.ContainerIsRunning(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/container-svc/container/summary", appl(func(w http.ResponseWriter, r *http.Request) {
		containerService.Summary(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/model-svc/default-model/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.DefaultStatus(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/model-svc/model/{modelId}/status", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.Status(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/model-svc/models", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.ListModels(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/model-svc/platforms", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.ListPlatforms(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/model-svc/model/{modelId}", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.Get(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/model-svc/default-model/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.StartDefault(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/model-svc/model/{modelId}/start", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.StartSpecific(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/model-svc/model/{modelId}/make-default", appl(func(w http.ResponseWriter, r *http.Request) {
		modelService.MakeDefault(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/config-svc/config", appl(func(w http.ResponseWriter, r *http.Request) {
		configService.Get(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/config-svc/config", appl(func(w http.ResponseWriter, r *http.Request) {
		configService.Save(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/chat-svc/thread/{threadId}/message", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.AddMessage(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/message/{messageId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetMessage(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/chat-svc/message/{messageId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.DeleteMessage(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/thread/{threadId}/messages", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetMessages(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.AddThread(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.DeleteThread(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/chat-svc/threads", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetThreads(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.GetThread(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/chat-svc/thread/{threadId}", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.UpdateThread(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/chat-svc/evens", appl(func(w http.ResponseWriter, r *http.Request) {
		chatService.Events(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompt", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.Prompt(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/prompt-svc/prompt/{promptId}", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.RemovePrompt(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/prompt-svc/prompts/{threadId}/responses/subscribe", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.SubscribeToPromptResponses(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/prompt-svc/prompts", appl(func(w http.ResponseWriter, r *http.Request) {
		promptService.ListPrompts(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/user-svc/login", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.Login(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user/by-token", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ReadUserByToken(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/users", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ListUsers(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SaveUser(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/self", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SaveSelf(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/change-password", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ChangePassword(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/{userId}/reset-password", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ResetPassword(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/organization", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SaveOrganization(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/organizations", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ListOrganizations(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/organization/{organizationId}/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.AddUserToOrganization(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/organization/{organizationId}/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.RemoveUserFromOrganization(w, r)
	})).
		Methods("OPTIONS", "DELETE")
	router.HandleFunc("/user-svc/user", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.CreateUser(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/user/{userId}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.DeleteUser(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/user-svc/self/has/{permission}", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.HasPermission(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/permissions", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ListPermissions(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/register", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.Register(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/public-key", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.GetPublicKey(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/user-svc/permits", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SavePermits(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/permits", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ListPermits(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/user-svc/enrolls", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.SaveEnrolls(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/user-svc/enrolls", appl(func(w http.ResponseWriter, r *http.Request) {
		userService.ListEnrolls(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/data-svc/object", appl(func(w http.ResponseWriter, r *http.Request) {
		dataService.Create(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/data-svc/objects/update", appl(func(w http.ResponseWriter, r *http.Request) {
		dataService.UpdateObjects(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/data-svc/objects/delete", appl(func(w http.ResponseWriter, r *http.Request) {
		dataService.DeleteObjects(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/data-svc/objects", appl(func(w http.ResponseWriter, r *http.Request) {
		dataService.Query(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/data-svc/object/{objectId}", appl(func(w http.ResponseWriter, r *http.Request) {
		dataService.Upsert(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/policy-svc/check", appl(func(w http.ResponseWriter, r *http.Request) {
		policyService.Check(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/policy-svc/instance/{instanceId}", appl(func(w http.ResponseWriter, r *http.Request) {
		policyService.UpsertInstance(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/registry-svc/node/self", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.NodeSelf(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/registry-svc/nodes", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.ListNodes(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/registry-svc/instances", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.ListInstances(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/registry-svc/definitions", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.ListDefinitions(w, r)
	})).
		Methods("OPTIONS", "GET")
	router.HandleFunc("/registry-svc/instance", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.RegisterInstance(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/registry-svc/definition", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.SaveDefinition(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/registry-svc/instance/{id}", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.RemoveInstance(w, r)
	})).
		Methods("OPTIONS", "DELETE")
	router.HandleFunc("/registry-svc/definition/{id}", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.DeleteDefinition(w, r)
	})).
		Methods("OPTIONS", "DELETE")
	router.HandleFunc("/registry-svc/node/{url}", appl(func(w http.ResponseWriter, r *http.Request) {
		registryService.DeleteNode(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/deploy-svc/deployment", appl(func(w http.ResponseWriter, r *http.Request) {
		deployService.SaveDeployment(w, r)
	})).
		Methods("OPTIONS", "PUT")
	router.HandleFunc("/deploy-svc/deployments", appl(func(w http.ResponseWriter, r *http.Request) {
		deployService.ListDeployments(w, r)
	})).
		Methods("OPTIONS", "POST")
	router.HandleFunc("/deploy-svc/deployment", appl(func(w http.ResponseWriter, r *http.Request) {
		deployService.DeleteDeployment(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/source-svc/repo/checkout", appl(func(w http.ResponseWriter, r *http.Request) {
		sourceService.CheckoutRepo(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", appl(func(w http.ResponseWriter, r *http.Request) {
		secretService.ListSecrets(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", appl(func(w http.ResponseWriter, r *http.Request) {
		secretService.SaveSecrets(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.HandleFunc("/secret-svc/encrypt", appl(func(w http.ResponseWriter, r *http.Request) {
		secretService.Encrypt(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/decrypt", appl(func(w http.ResponseWriter, r *http.Request) {
		secretService.Decrypt(w, r)
	})).
		Methods("OPTIONS", "POST")

	router.HandleFunc("/secret-svc/secrets", appl(func(w http.ResponseWriter, r *http.Request) {
		secretService.RemoveSecrets(w, r)
	})).
		Methods("OPTIONS", "DELETE")

	router.HandleFunc("/secret-svc/is-secure", appl(func(w http.ResponseWriter, r *http.Request) {
		secretService.Secure(w, r)
	})).
		Methods("OPTIONS", "GET")

	router.HandleFunc("/email-svc/email", appl(func(w http.ResponseWriter, r *http.Request) {
		emailService.SendEmail(w, r)
	})).
		Methods("OPTIONS", "PUT")

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyService.Route(w, r)
	})

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return &Universe{
		Router: router,
		StarterFunc: func() error {
			err = configService.Start()
			if err != nil {
				return errors.Wrap(err, "config service start failed")
			}
			err = firehoseService.Start()
			if err != nil {
				return errors.Wrap(err, "firehose service start failed")
			}
			err = containerService.Start()
			if err != nil {
				return errors.Wrap(err, "docker service start failed")
			}
			err = modelService.Start()
			if err != nil {
				return errors.Wrap(err, "model service start failed")
			}
			err = chatService.Start()
			if err != nil {
				return errors.Wrap(err, "chat service start failed")
			}
			err = promptService.Start()
			if err != nil {
				return errors.Wrap(err, "prompt service start failed")
			}
			err = dataService.Start()
			if err != nil {
				return errors.Wrap(err, "data service start failed")
			}
			err = policyService.Start()
			if err != nil {
				return errors.Wrap(err, "policy service start failed")
			}
			err = registryService.Start()
			if err != nil {
				return errors.Wrap(err, "registry service start failed")
			}
			err = fileService.Start()
			if err != nil {
				return errors.Wrap(err, "file service start failed")
			}
			err = deployService.Start()
			if err != nil {
				return errors.Wrap(err, "deploy service start failed")
			}
			err = sourceService.Start()
			if err != nil {
				return errors.Wrap(err, "source service start failed")
			}
			err = secretService.Start()
			if err != nil {
				return errors.Wrap(err, "secret service start failed")
			}
			err = proxyService.Start()
			if err != nil {
				return errors.Wrap(err, "proxy service start failed")
			}
			err = emailService.Start()
			if err != nil {
				return errors.Wrap(err, "email service start failed")
			}

			return nil
		},
		Options: *options,
	}, nil
}

func applicator(
	mws []middlewares.Middleware,
) func(http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range mws {
			h = middleware(h)
		}

		return h
	}
}

type HandlerSwitcher struct {
	mu      sync.RWMutex
	handler http.Handler
}

func (hs *HandlerSwitcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hs.mu.RLock()
	defer hs.mu.RUnlock()
	if hs.handler != nil {
		hs.handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (hs *HandlerSwitcher) UpdateHandler(handler http.Handler) {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	hs.handler = handler
}
