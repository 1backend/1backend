package basicservice

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"

	basic "github.com/1backend/1backend/examples/go/services/basic/internal/types"
)

// @ID savePet
// @Summary Save Pet
// @Description Save a pet for a user and an organization.
// @Tags Basic Svc
// @Accept json
// @Produce json
// @Param body body basic.SavePetRequest true "Registration Tracking Request"
// @Success 200 {object} basic.SavePetResponse "{}"
// @Failure 400 {string} string "Invalid JSON"
// @Router /basic-svc/pet [put]
func (s *BasicService) SavePet(w http.ResponseWriter, r *http.Request) {
	var request basic.SavePetRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	var authorizer auth.Authorizer = auth.AuthorizerImpl{}
	token, err := authorizer.ParseJWTFromRequest(s.userSvcPublicKey, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !lo.Contains(token.RoleIds, RolePetManager) {
		spew.Dump(token.RoleIds)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	pet := basic.Pet{
		Name: request.Name,
	}

	err = validatePet(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = s.savePet(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}

func (s *BasicService) savePet(pet basic.Pet) error {
	now := time.Now()

	if pet.Id == "" {
		pet.Id = sdk.Id("camp")
		pet.CreatedAt = now
	}

	pet.UpdatedAt = now

	return s.petsStore.Upsert(pet)
}

func validatePet(pet basic.Pet) error {
	if pet.Name == "" {
		return errors.New("no pet name")
	}

	return nil
}
