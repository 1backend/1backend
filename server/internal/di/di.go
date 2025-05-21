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
	"time"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	pglock "github.com/1backend/1backend/sdk/go/lock/pg"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/1backend/1backend/sdk/go/infra"
	distlock "github.com/1backend/1backend/sdk/go/lock/local"
	"github.com/1backend/1backend/sdk/go/logger"

	"github.com/1backend/1backend/server/internal/router"
	chatservice "github.com/1backend/1backend/server/internal/services/chat"
	configservice "github.com/1backend/1backend/server/internal/services/config"
	containerservice "github.com/1backend/1backend/server/internal/services/container"
	dataservice "github.com/1backend/1backend/server/internal/services/data"
	deployservice "github.com/1backend/1backend/server/internal/services/deploy"
	emailservice "github.com/1backend/1backend/server/internal/services/email"
	fileservice "github.com/1backend/1backend/server/internal/services/file"
	firehoseservice "github.com/1backend/1backend/server/internal/services/firehose"
	imageservice "github.com/1backend/1backend/server/internal/services/image"
	modelservice "github.com/1backend/1backend/server/internal/services/model"
	policyservice "github.com/1backend/1backend/server/internal/services/policy"
	promptservice "github.com/1backend/1backend/server/internal/services/prompt"
	proxyservice "github.com/1backend/1backend/server/internal/services/proxy"
	registryservice "github.com/1backend/1backend/server/internal/services/registry"
	secretservice "github.com/1backend/1backend/server/internal/services/secret"
	sourceservice "github.com/1backend/1backend/server/internal/services/source"
	userservice "github.com/1backend/1backend/server/internal/services/user"
	"github.com/1backend/1backend/server/internal/universe"
)

type Universe struct {
	Options       universe.Options
	Router        *mux.Router
	StarterFunc   func() error
	ClientFactory client.ClientFactory
}

func BigBang(options *universe.Options) (*Universe, error) {
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
	if options.TokenExpiration == 0 {
		tokenExpiration := os.Getenv("OB_TOKEN_EXPIRATION")
		if tokenExpiration != "" {
			tokenExpiration, err := time.ParseDuration(tokenExpiration)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse token expiration")
			}
			options.TokenExpiration = tokenExpiration
		}
	}
	if options.TokenExpiration == 0 {
		options.TokenExpiration = 5 * time.Minute
	}

	if os.Getenv("OB_TOKEN_AUTO_REFRESH_OFF") == "true" {
		options.TokenAutoRefreshOff = true
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
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create config service")
	}

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

	if options.TokenRefresher == nil {
		tokenRefresher, err := endpoint.NewTokenRefresher(
			options.ClientFactory,
			options.Authorizer,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create token refresher")
		}
		options.TokenRefresher = tokenRefresher
	}

	if options.Middlewares == nil {
		mws := middlewares.Applicator(
			options.TokenRefresher,
			options.TokenAutoRefreshOff,
		)
		options.Middlewares = mws
	}

	if options.PermissionChecker == nil {
		options.PermissionChecker = endpoint.NewPermissionChecker(
			options.ClientFactory,
		)
	}

	configService.RegisterRoutes(router)

	userService, err := userservice.NewUserService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user service")
	}
	userService.RegisterRoutes(router)

	firehoseService, err := firehoseservice.NewFirehoseService(
		options,
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
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file service")
	}
	fileService.RegisterRoutes(router)

	imageService, err := imageservice.NewImageService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create image service")
	}
	imageService.RegisterRoutes(router)

	containerService, err := containerservice.NewContainerService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create container service")
	}
	containerService.RegisterRoutes(router)

	modelService, err := modelservice.NewModelService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create model service")
	}
	modelService.RegisterRoutes(router)

	chatService, err := chatservice.NewChatService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create chat service")
	}
	chatService.RegisterRoutes(router)

	promptService, err := promptservice.NewPromptService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create prompt service")
	}
	promptService.RegisterRoutes(router)

	dataService, err := dataservice.NewDataService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create data service")
	}
	dataService.RegisterRoutes(router)

	policyService, err := policyservice.NewPolicyService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create policy service")
	}
	policyService.RegisterRoutes(router)

	registryService, err := registryservice.NewRegistryService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create registry service")
	}
	registryService.RegisterRoutes(router)

	deployService, err := deployservice.NewDeployService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create deploy service")
	}
	deployService.RegisterRoutes(router)

	sourceService, err := sourceservice.NewSourceService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create source service")
	}
	sourceService.RegisterRoutes(router)

	if options.SecretEncryptionKey == "" {
		options.SecretEncryptionKey = "changeMeToSomethingSecureForReal"
	}
	secretService, err := secretservice.NewSecretService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create secret service")
	}
	secretService.RegisterRoutes(router)

	emailService, err := emailservice.NewEmailService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create email service")
	}
	emailService.RegisterRoutes(router)

	router.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("404 page not found"))
			if err != nil {
				logger.Error("Error writing response",
					slog.String("error", err.Error()),
				)
			}
		},
	)

	proxyService, err := proxyservice.NewProxyService(
		options,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create proxy service")
	}
	proxyService.RegisterRoutes(router)

	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return &Universe{
		Router: router,
		StarterFunc: func() error {
			err = userService.Start()
			if err != nil {
				return errors.Wrap(err, "user service start failed")
			}

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
