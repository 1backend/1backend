package di

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/1backend/1backend/server/internal/universe"
	"github.com/gorilla/mux"
)

const readinessCheckTimeout = 1500 * time.Millisecond

type readinessResponse struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks"`
}

func registerReadinessRoutes(router *mux.Router, univ *Universe, options *universe.Options) {
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		writeReadinessResponse(w, http.StatusOK, readinessResponse{
			Status: "ok",
			Checks: map[string]string{
				"http": "ok",
			},
		})
	}

	router.HandleFunc("/healthz", healthHandler).Methods(http.MethodGet)
	router.HandleFunc("/livez", healthHandler).Methods(http.MethodGet)

	router.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), readinessCheckTimeout)
		defer cancel()

		status := readiness(ctx, univ, options)
		code := http.StatusOK
		if status.Status != "ok" {
			code = http.StatusServiceUnavailable
		}
		writeReadinessResponse(w, code, status)
	}).Methods(http.MethodGet)
}

func readiness(ctx context.Context, univ *Universe, options *universe.Options) readinessResponse {
	checks := map[string]string{}
	if !univ.started.Load() {
		checks["startup"] = "not-started"
		return readinessResponse{
			Status: "not-ready",
			Checks: checks,
		}
	}
	checks["startup"] = "ok"

	if options.Db == "" {
		checks["database"] = "disabled"
		checks["lock"] = "local"
		return readinessResponse{
			Status: "ok",
			Checks: checks,
		}
	}

	if err := infra.CheckDataStoreReady(ctx, options.DataStoreFactory); err != nil {
		checks["database"] = fmt.Sprintf("writer check failed: %v", err)
		return readinessResponse{
			Status: "not-ready",
			Checks: checks,
		}
	}
	checks["database"] = "ok"

	if options.Lock == nil {
		checks["lock"] = "missing"
		return readinessResponse{
			Status: "not-ready",
			Checks: checks,
		}
	}

	lockKey := readinessLockKey()
	if err := options.Lock.Acquire(ctx, lockKey); err != nil {
		checks["lock"] = fmt.Sprintf("acquire failed: %v", err)
		return readinessResponse{
			Status: "not-ready",
			Checks: checks,
		}
	}
	if err := options.Lock.Release(context.Background(), lockKey); err != nil {
		checks["lock"] = fmt.Sprintf("release failed: %v", err)
		return readinessResponse{
			Status: "not-ready",
			Checks: checks,
		}
	}
	checks["lock"] = "ok"

	return readinessResponse{
		Status: "ok",
		Checks: checks,
	}
}

func readinessLockKey() string {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}
	return fmt.Sprintf("onebackend:readyz:%s:%d", hostname, os.Getpid())
}

func writeReadinessResponse(w http.ResponseWriter, code int, response readinessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}
