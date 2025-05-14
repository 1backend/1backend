package basicservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	basic "github.com/1backend/1backend/examples/go/services/basic/internal/types"
)

// @ID listPets
// @Summary List Pets
// @Description List pets.
// @Tags Basic Svc
// @Accept json
// @Produce json
// @Param body body basic.ListPetsRequest false "List Pets Request"
// @Success 200 {object} basic.ListPetsResponse "{}"
// @Failure 400 {object} basic.ErrorResponse "Invalid JSON"
// @Failure 500 {object} basic.ErrorResponse "Error Listing Pets"
// @Router /basic-svc/pets [post]
func (s *BasicService) ListPets(w http.ResponseWriter, r *http.Request) {
	var request basic.ListPetsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	rsp, err := s.listPets(request)
	if err != nil {
		logger.Error(
			"Failed to list pets",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(rsp)
	w.Write(bs)
}

func (s *BasicService) listPets(request basic.ListPetsRequest) (*basic.ListPetsResponse, error) {
	petIs, err := s.petsStore.Query().OrderBy(
		datastore.OrderByField("updatedAt", true),
	).Find()
	if err != nil {
		return nil, err
	}

	pets := []basic.Pet{}
	for _, v := range petIs {
		pets = append(pets, *v.(*basic.Pet))
	}

	return &basic.ListPetsResponse{
		Pets: pets,
	}, nil
}
