package multiservice

import (
	"context"
	"encoding/json"
	"net/http"

	multi "github.com/1backend/1backend/examples/go/services/multi/internal/types"
)

// @ID countPets
// @Summary Count Pets
// @Description Count the pets living in the Multi Svc.
// @Tags Multi Svc
// @Accept json
// @Produce json
// @Param body body multi.CountPetsRequest false "Count Pets Request"
// @Success 200 {object} multi.CountPetsResponse "{}"
// @Failure 400 {string} string "Invalid JSON"
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
	w.Write(bs)
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
