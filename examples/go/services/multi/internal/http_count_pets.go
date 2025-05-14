package multiservice

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	multi "github.com/1backend/1backend/examples/go/services/multi/internal/types"
	"github.com/1backend/1backend/sdk/go/logger"
)

// @ID countPets
// @Summary Count Pets
// @Description Count the pets living in the Multi Svc.
// @Tags Multi Svc
// @Accept json
// @Produce json
// @Param body body multi.CountPetsRequest false "Count Pets Request"
// @Success 200 {object} multi.CountPetsResponse "{}"
// @Failure 400 {object} multi.ErrorResponse "Invalid JSON"
// @Router /multi-svc/pets/count [get]
func (s *MultiService) CountPets(w http.ResponseWriter, r *http.Request) {
	count, err := s.countPets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(multi.CountPetsResponse{
		PetCount: count,
	})

	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Failed to write response", slog.Any("error", err))
		return
	}
}

func (s *MultiService) countPets() (int, error) {
	rsp, _, err := s.basicSvcClient.BasicSvcAPI.
		ListPets(context.Background()).
		Execute()

	if err != nil {
		return 0, err
	}

	return len(rsp.Pets), nil
}
