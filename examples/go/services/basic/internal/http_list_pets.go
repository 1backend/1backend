package basicservice

import (
	"encoding/json"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"

	basic "github.com/1backend/1backend/examples/go/services/basic/internal/types"
)

// @ID listPets
// @Summary List Pets
// @Description List pets.
// @Tags Basic Svc
// @Accept json
// @Produce json
// @Param body body basic.ListPetsRequest true "Registration Tracking Request"
// @Success 200 {object} basic.ListPetsResponse "{}"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 500 {string} string "Error Tracking Registration"
// @Router /basic-svc/pets [post]
func (s *BasicService) ListPets(w http.ResponseWriter, r *http.Request) {
	var request basic.ListPetsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	rsp, err := s.listPets(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
