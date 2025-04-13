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

	// Path of the config folder, configurable via the "OB_FOLDER" environment variable.
	// If Test is true, this value is ignored and a random temporary folder is used instead.
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
			logger.Error("Panic in BigBang",
				slog.String("error", fmt.Sprintf("%v", r)),
				slog.String("trace", string(debug.Stack())),
			)
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

	if options.ConfigPath == "" {
		options.ConfigPath = os.Getenv("OB_FOLDER")
	}

	homeDir, err := infra.HomeDir(infra.HomeDirOptions{
		Test:         options.Test,
		ConfigFolder: options.ConfigPath,
	})
	if err != nil {
		return nil, err
	}

	options.HomeDir = homeDir

	logger.Info("Using folder", slog.String("folder", options.HomeDir))

	if options.Lock == nil {
		options.Lock = distlock.NewLocalDistributedLock()
	}

	if options.Authorizer == nil {
		options.Authorizer = auth.AuthorizerImpl{}
	}

	router := mux.NewRouter().SkipClean(true).UseEncodedPath()

	configService, err := configservice.NewConfigService(
		options.Lock,
		options.Authorizer,
		options.HomeDir,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create config service")
	}
	configService.RegisterRoutes(router)

	if options.DataStoreFactory == nil {
		dc, err := infra.NewDataStoreFactory(infra.DataStoreConfig{
			HomeDir:            options.HomeDir,
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
				return nil, errors.Wrap(err, "failed to get db handle")
			}
			conn, err := dbHandle.(*sql.DB).Conn(context.Background())
			if err != nil {
				return nil, errors.Wrap(err, "failed to get db connection")
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
		return nil, errors.Wrap(err, "failed to create user service")
	}
	userService.RegisterRoutes(router)

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
		return nil, errors.Wrap(err, "failed to create firehose service")
	}
	firehoseService.RegisterRoutes(router)

	err = os.MkdirAll(options.HomeDir, 0755)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create config folder")
	}

	fileService, err := fileservice.NewFileService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
		options.HomeDir,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file service")
	}
	fileService.RegisterRoutes(router)

	containerService, err := containerservice.NewContainerService(
		options.VolumeName,
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create container service")
	}
	containerService.RegisterRoutes(router)

	modelService, err := modelservice.NewModelService(
		// @todo GPU platform maybe this could be autodetected
		options.GpuPlatform,
		options.LLMHost,
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create model service")
	}
	modelService.RegisterRoutes(router)

	chatService, err := chatservice.NewChatService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create chat service")
	}
	chatService.RegisterRoutes(router)

	promptService, err := promptservice.NewPromptService(
		options.ClientFactory,
		options.LLamaCppClient,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create prompt service")
	}
	promptService.RegisterRoutes(router)

	dataService, err := dataservice.NewDataService(
		options.ClientFactory,
		options.Lock,
		options.Authorizer,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create data service")
	}
	dataService.RegisterRoutes(router)

	policyService, err := policyservice.NewPolicyService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create policy service")
	}
	policyService.RegisterRoutes(router)

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
		return nil, errors.Wrap(err, "failed to create registry service")
	}
	registryService.RegisterRoutes(router)

	deployService, err := deployservice.NewDeployService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
		options.Test,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create deploy service")
	}
	deployService.RegisterRoutes(router)

	sourceService, err := sourceservice.NewSourceService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create source service")
	}
	sourceService.RegisterRoutes(router)

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
		return nil, errors.Wrap(err, "failed to create secret service")
	}
	secretService.RegisterRoutes(router)

	emailService, err := emailservice.NewEmailService(
		options.ClientFactory,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create email service")
	}
	emailService.RegisterRoutes(router)

	router.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found"))
		},
	)

	proxyService, err := proxyservice.NewProxyService(
		options.ClientFactory,
		options.Authorizer,
		options.Lock,
		options.DataStoreFactory.Create,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create proxy service")
	}
	proxyService.RegisterRoutes(router)

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return &Universe{
		Router: router,
		StarterFunc: func() error {
			err = promptService.Start()
			if err != nil {
				return errors.Wrap(err, "prompt service start failed")
			}

			err = registryService.Start()
			if err != nil {
				return errors.Wrap(err, "registry service start failed")
			}

			err = fileService.Start()
			if err != nil {
				return errors.Wrap(err, "file service start failed")
			}

			err = containerService.Start()
			if err != nil {
				return errors.Wrap(err, "container service start failed")
			}

			err = deployService.Start()
			if err != nil {
				return errors.Wrap(err, "deploy service start failed")
			}

			return nil
		},
		Options: *options,
	}, nil
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
