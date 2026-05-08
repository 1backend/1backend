package boot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/gorilla/mux"
)

const defaultReadinessCheckTimeout = 1500 * time.Millisecond

type HealthRoutesOptions struct {
	DataStoreFactory infra.DataStoreFactory
	Timeout          time.Duration
}

type healthResponse struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks"`
}

func RegisterHealthRoutes(router *mux.Router, options HealthRoutesOptions) {
	if router == nil {
		return
	}

	router.HandleFunc("/healthz", writeHealthy).Methods(http.MethodGet)
	router.HandleFunc("/livez", writeHealthy).Methods(http.MethodGet)

	timeout := options.Timeout
	if timeout == 0 {
		timeout = defaultReadinessCheckTimeout
	}

	router.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		checks := map[string]string{
			"startup": "ok",
		}

		if options.DataStoreFactory != nil {
			if err := infra.CheckDataStoreReady(ctx, options.DataStoreFactory); err != nil {
				checks["database"] = fmt.Sprintf("writer check failed: %v", err)
				writeHealthResponse(w, http.StatusServiceUnavailable, healthResponse{
					Status: "not-ready",
					Checks: checks,
				})
				return
			}
			checks["database"] = "ok"
		} else {
			checks["database"] = "disabled"
		}

		writeHealthResponse(w, http.StatusOK, healthResponse{
			Status: "ok",
			Checks: checks,
		})
	}).Methods(http.MethodGet)
}

func (o *Options) RegisterHealthRoutes(router *mux.Router, dataStoreFactory infra.DataStoreFactory) {
	RegisterHealthRoutes(router, HealthRoutesOptions{
		DataStoreFactory: dataStoreFactory,
	})
}

func writeHealthy(w http.ResponseWriter, r *http.Request) {
	writeHealthResponse(w, http.StatusOK, healthResponse{
		Status: "ok",
		Checks: map[string]string{
			"http": "ok",
		},
	})
}

func writeHealthResponse(w http.ResponseWriter, code int, response healthResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)
}
