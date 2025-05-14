package basicservice

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
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
// @Failure 400 {object} basic.ErrorResponse "Invalid JSON"
// @Failure 400 {object} basic.ErrorResponse "Invalid Pet"
// @Failure 500 {object} basic.ErrorResponse "Internal Server Error"
// @Router /basic-svc/pet [put]
func (s *BasicService) SavePet(w http.ResponseWriter, r *http.Request) {
	var request basic.SavePetRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	var authorizer auth.Authorizer = auth.AuthorizerImpl{}
	token, err := authorizer.ParseJWTFromRequest(s.userSvcPublicKey, r)
	if err != nil {
		logger.Error(
			"Failed to parse JWT from request",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	if !lo.Contains(token.Roles, RolePetManager) {
		endpoint.Unauthorized(w)
		return
	}

	pet := basic.Pet{
		Name: request.Name,
	}

	err = validatePet(pet)
	if err != nil {
		logger.Error(
			"Invalid pet",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid Pet")
		return
	}

	err = s.savePet(pet)
	if err != nil {
		logger.Error(
			"Failed to save pet",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
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
