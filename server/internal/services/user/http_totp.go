/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"image/png"
	"io"
	"log/slog"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

const (
	totpIssuer            = "1Backend"
	totpDisplayTextMaxLen = 128
	totpPeriod            = 30
	totpSecretSize        = 20
	totpQRSize            = 256
	missingTOTPUserID     = "__missing_totp_user__"
)

var (
	errTOTPRequired = errors.New("totp required")
	errTOTPInvalid  = errors.New("invalid totp")
)

// @ID beginTOTPSetup
// @Summary Begin TOTP Setup
// @Description Creates a pending time-based one-time password secret for the authenticated user's account.
// @Description If TOTP is already enabled, currentCode must verify the active authenticator before a reprovisioning setup is created with the requested issuer/accountName. Existing secrets are preserved unless rotateSecret is true.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.BeginTOTPSetupRequest false "Begin TOTP Setup Request"
// @Success 200 {object} user.BeginTOTPSetupResponse "TOTP setup started"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 409 {object} user.ErrorResponse "TOTP already enabled"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/totp/setup [post]
func (s *UserService) BeginTOTPSetup(w http.ResponseWriter, r *http.Request) {
	claims, usr, err := s.authenticatedTOTPUser(r)
	if err != nil {
		endpoint.Unauthorized(w)
		return
	}

	req := user.BeginTOTPSetupRequest{}
	if r.Body != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil && !errors.Is(err, io.EOF) {
			endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		defer r.Body.Close()
	}

	existingSecret := ""
	if current, enabled, err := s.enabledTOTPForUser(usr.Id); err != nil {
		logger.Error("Failed to check TOTP status", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	} else if enabled {
		if req.CurrentCode == "" {
			endpoint.WriteString(w, http.StatusConflict, "TOTP already enabled")
			return
		}

		if ok, err := validateTOTPCode(current.Secret, req.CurrentCode); err != nil || !ok {
			if err != nil {
				logger.Error("Failed to validate current TOTP", slog.Any("error", err))
			}
			endpoint.WriteString(w, http.StatusUnauthorized, "Invalid TOTP code")
			return
		}
		if !req.RotateSecret {
			existingSecret = current.Secret
		}
	}

	issuer, err := totpSetupIssuer(req.Issuer)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, err.Error())
		return
	}
	accountName, err := s.totpSetupAccountName(req.AccountName, usr)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, err.Error())
		return
	}

	key, err := generateTOTPKey(issuer, accountName, existingSecret)
	if err != nil {
		logger.Error("Failed to generate TOTP key", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	now := time.Now()
	totpId := sdk.Id("totp")
	internalId, err := sdk.InternalId(claims.AppId, totpId)
	if err != nil {
		logger.Error("Failed to create TOTP internal id", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	record := &user.TOTP{
		InternalId:      internalId,
		AppId:           claims.AppId,
		Id:              totpId,
		CreatedAt:       now,
		UpdatedAt:       now,
		UserId:          usr.Id,
		Issuer:          key.Issuer(),
		AccountName:     key.AccountName(),
		Secret:          key.Secret(),
		ProvisioningURI: key.URL(),
	}
	if err := s.totpStore.Create(record); err != nil {
		logger.Error("Failed to store TOTP setup", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.BeginTOTPSetupResponse{
		TOTPId:          record.Id,
		Secret:          record.Secret,
		ProvisioningURI: record.ProvisioningURI,
		QRImagePath:     "/user-svc/totp/setup/" + record.Id + "/qr.png",
	})
}

func totpSetupIssuer(raw string) (string, error) {
	return totpSetupDisplayText(raw, totpIssuer, "issuer", true)
}

func (s *UserService) totpSetupAccountName(raw string, usr *user.User) (string, error) {
	template := strings.TrimSpace(raw)
	if template == "" {
		template = "$slug"
	}

	if !strings.Contains(template, "$") {
		return totpSetupDisplayText(template, "", "accountName", false)
	}

	contacts, err := s.getContactsByUserId(usr.Id)
	if err != nil {
		return "", err
	}

	accountName, err := renderTOTPAccountNameTemplate(
		template,
		totpAccountNameFields(usr, contacts),
	)
	if err != nil {
		return "", err
	}

	return totpSetupDisplayText(accountName, "", "accountName", false)
}

func totpSetupDisplayText(
	raw string,
	fallback string,
	field string,
	rejectLabelDelimiters bool,
) (string, error) {
	value := strings.TrimSpace(raw)
	if value == "" {
		value = strings.TrimSpace(fallback)
	}
	if value == "" {
		return "", errors.New(field + " is required")
	}
	if len(value) > totpDisplayTextMaxLen {
		return "", errors.New(field + " is too long")
	}

	for _, r := range value {
		if r < 0x20 || r == 0x7f || r > 0x7e {
			return "", errors.New(field + " contains unsupported characters")
		}
		if rejectLabelDelimiters {
			switch r {
			case ':', '/', '\\', '?', '#':
				return "", errors.New(field + " contains unsupported characters")
			}
		}
	}

	return value, nil
}

func generateTOTPKey(
	issuer string,
	accountName string,
	existingSecret string,
) (*otp.Key, error) {
	var secret []byte
	if existingSecret != "" {
		var err error
		secret, err = base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(
			strings.ToUpper(strings.TrimSpace(existingSecret)),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to decode existing TOTP secret: %w", err)
		}
	}

	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		Period:      totpPeriod,
		SecretSize:  totpSecretSize,
		Secret:      secret,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
}

func renderTOTPAccountNameTemplate(
	template string,
	fields map[string]string,
) (string, error) {
	var out strings.Builder
	for i := 0; i < len(template); {
		if template[i] != '$' {
			out.WriteByte(template[i])
			i++
			continue
		}

		name := ""
		next := i + 1
		if next < len(template) && template[next] == '{' {
			end := strings.IndexByte(template[next+1:], '}')
			if end == -1 {
				return "", errors.New("accountName contains an unterminated placeholder")
			}
			name = template[next+1 : next+1+end]
			next = next + 1 + end + 1
		} else {
			for next < len(template) && isTOTPPlaceholderChar(template[next]) {
				next++
			}
			if next == i+1 {
				out.WriteByte(template[i])
				i++
				continue
			}
			name = template[i+1 : next]
		}

		value, ok := fields[name]
		if !ok {
			return "", fmt.Errorf("accountName contains unknown placeholder %q", "$"+name)
		}
		if value == "" {
			return "", fmt.Errorf("accountName placeholder %q is empty", "$"+name)
		}

		out.WriteString(value)
		i = next
	}

	return out.String(), nil
}

func isTOTPPlaceholderChar(c byte) bool {
	return c >= 'A' && c <= 'Z' ||
		c >= 'a' && c <= 'z' ||
		c >= '0' && c <= '9' ||
		c == '_'
}

func totpAccountNameFields(
	usr *user.User,
	contacts []user.Contact,
) map[string]string {
	fields := map[string]string{
		"name":       usr.Name,
		"slug":       usr.Slug,
		"contactId":  "",
		"contactIds": "",
		"email":      "",
		"phone":      "",
	}

	sort.SliceStable(contacts, func(i int, j int) bool {
		if contacts[i].IsPrimary != contacts[j].IsPrimary {
			return contacts[i].IsPrimary
		}
		if contacts[i].Platform != contacts[j].Platform {
			return contacts[i].Platform < contacts[j].Platform
		}
		return contacts[i].Id < contacts[j].Id
	})

	contactIds := []string{}
	for _, contact := range contacts {
		contactIds = append(contactIds, contact.Id)
		if fields["contactId"] == "" {
			fields["contactId"] = contact.Id
		}
		if contact.Platform != "" && fields[contact.Platform] == "" {
			fields[contact.Platform] = contact.Id
		}
	}
	fields["contactIds"] = strings.Join(contactIds, ",")

	return fields
}

// @ID readTOTPStatus
// @Summary Read TOTP Status
// @Description Returns whether an email contact belongs to an account with TOTP enabled.
// @Description Unknown emails return totpEnabled=false instead of 404 to avoid a registeredness signal for non-TOTP accounts.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ReadTOTPStatusRequest true "Read TOTP Status Request"
// @Success 200 {object} user.ReadTOTPStatusResponse "TOTP status"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/totp/status [post]
func (s *UserService) ReadTOTPStatus(w http.ResponseWriter, r *http.Request) {
	req := user.ReadTOTPStatusRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	email := normalizeEmail(req.Email)
	if email == "" || !isEmail(email) {
		endpoint.WriteString(w, http.StatusBadRequest, "email is required")
		return
	}

	userId := missingTOTPUserID
	contactI, found, err := s.contactStore.Query(
		datastore.Equals(datastore.Field("id"), email),
	).FindOne()
	if err != nil {
		logger.Error("Failed to read contact for TOTP status", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if found {
		if contact := contactI.(*user.Contact); contact.UserId != "" {
			userId = contact.UserId
		}
	}

	_, enabled, err := s.enabledTOTPForUser(userId)
	if err != nil {
		logger.Error("Failed to read TOTP status", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ReadTOTPStatusResponse{
		TOTPEnabled: enabled,
	})
}

// @ID serveTOTPQRCode
// @Summary Serve TOTP QR Code
// @Description Generates and serves a PNG QR code for a pending TOTP setup owned by the authenticated user.
// @Tags User Svc
// @Produce png
// @Param totpId path string true "TOTP setup ID"
// @Success 200 "PNG QR code"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 404 {object} user.ErrorResponse "TOTP setup not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/totp/setup/{totpId}/qr.png [get]
func (s *UserService) ServeTOTPQRCode(w http.ResponseWriter, r *http.Request) {
	claims, _, err := s.authenticatedTOTPUser(r)
	if err != nil {
		endpoint.Unauthorized(w)
		return
	}

	totpId := mux.Vars(r)["totpId"]
	record, found, err := s.totpForUser(totpId, claims.UserId)
	if err != nil {
		logger.Error("Failed to read TOTP setup", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if !found || record.Enabled || record.DisabledAt != nil {
		endpoint.WriteString(w, http.StatusNotFound, "TOTP setup not found")
		return
	}

	key, err := otp.NewKeyFromURL(record.ProvisioningURI)
	if err != nil {
		logger.Error("Failed to parse TOTP provisioning URI", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	image, err := key.Image(totpQRSize, totpQRSize)
	if err != nil {
		logger.Error("Failed to generate TOTP QR image", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-store")
	if err := png.Encode(w, image); err != nil {
		logger.Error("Failed to encode TOTP QR image", slog.Any("error", err))
	}
}

// @ID enableTOTP
// @Summary Enable TOTP
// @Description Verifies a pending TOTP setup code and enables TOTP for the authenticated user's account.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.EnableTOTPRequest true "Enable TOTP Request"
// @Success 200 {object} user.EnableTOTPResponse "TOTP enabled"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 404 {object} user.ErrorResponse "TOTP setup not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/totp/enable [post]
func (s *UserService) EnableTOTP(w http.ResponseWriter, r *http.Request) {
	claims, _, err := s.authenticatedTOTPUser(r)
	if err != nil {
		endpoint.Unauthorized(w)
		return
	}

	req := user.EnableTOTPRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.Code == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "code is required")
		return
	}

	record, found, err := s.pendingTOTPForUser(req.TOTPId, claims.UserId)
	if err != nil {
		logger.Error("Failed to read TOTP setup", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		endpoint.WriteString(w, http.StatusNotFound, "TOTP setup not found")
		return
	}

	if current, enabled, err := s.enabledTOTPForUser(claims.UserId); err != nil {
		logger.Error("Failed to read active TOTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	} else if enabled && !totpSetupCreatedAfterActive(record, current) {
		endpoint.WriteString(w, http.StatusConflict, "TOTP setup is no longer valid")
		return
	}

	if ok, err := validateTOTPCode(record.Secret, req.Code); err != nil || !ok {
		if err != nil {
			logger.Error("Failed to validate TOTP", slog.Any("error", err))
		}
		endpoint.WriteString(w, http.StatusUnauthorized, "Invalid TOTP code")
		return
	}

	if err := s.disableEnabledTOTPsForUser(claims.UserId); err != nil {
		logger.Error("Failed to disable previous TOTP factors", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	now := time.Now()
	record.Enabled = true
	record.EnabledAt = &now
	record.DisabledAt = nil
	record.UpdatedAt = now
	if err := s.totpStore.Upsert(record); err != nil {
		logger.Error("Failed to enable TOTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.EnableTOTPResponse{Enabled: true})
}

// @ID disableTOTP
// @Summary Disable TOTP
// @Description Verifies the current TOTP code and disables TOTP for the authenticated user's account.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.DisableTOTPRequest true "Disable TOTP Request"
// @Success 200 {object} user.DisableTOTPResponse "TOTP disabled"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 404 {object} user.ErrorResponse "TOTP not enabled"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/totp [delete]
func (s *UserService) DisableTOTP(w http.ResponseWriter, r *http.Request) {
	claims, _, err := s.authenticatedTOTPUser(r)
	if err != nil {
		endpoint.Unauthorized(w)
		return
	}

	req := user.DisableTOTPRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.Code == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "code is required")
		return
	}

	record, found, err := s.enabledTOTPForUser(claims.UserId)
	if err != nil {
		logger.Error("Failed to read TOTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	if !found {
		endpoint.WriteString(w, http.StatusNotFound, "TOTP not enabled")
		return
	}

	if ok, err := validateTOTPCode(record.Secret, req.Code); err != nil || !ok {
		if err != nil {
			logger.Error("Failed to validate TOTP", slog.Any("error", err))
		}
		endpoint.WriteString(w, http.StatusUnauthorized, "Invalid TOTP code")
		return
	}

	now := time.Now()
	record.Enabled = false
	record.DisabledAt = &now
	record.UpdatedAt = now
	if err := s.totpStore.Upsert(record); err != nil {
		logger.Error("Failed to disable TOTP", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.DisableTOTPResponse{})
}

func (s *UserService) authenticatedTOTPUser(r *http.Request) (*auth.Claims, *user.User, error) {
	rawToken, hasToken := s.options.Authorizer.TokenFromRequest(r)
	if !hasToken {
		return nil, nil, errors.New("token missing")
	}

	claims, err := s.options.Authorizer.ParseJWT(s.publicKeyPem, rawToken)
	if err != nil {
		return nil, nil, err
	}

	if _, found, err := s.tokenStore.Query(
		datastore.Equals(datastore.Field("token"), rawToken),
	).FindOne(); err != nil {
		return nil, nil, err
	} else if !found {
		return nil, nil, errors.New("token not found")
	}

	usr, err := s.readSelf(claims.UserId)
	if err != nil {
		return nil, nil, err
	}

	return claims, usr, nil
}

func (s *UserService) verifyLoginTOTP(userId string, code string) error {
	record, enabled, err := s.enabledTOTPForUser(userId)
	if err != nil {
		return err
	}
	if !enabled {
		return nil
	}
	if code == "" {
		return errTOTPRequired
	}

	ok, err := validateTOTPCode(record.Secret, code)
	if err != nil {
		return err
	}
	if !ok {
		return errTOTPInvalid
	}

	return nil
}

func (s *UserService) enabledTOTPForUser(userId string) (*user.TOTP, bool, error) {
	recordI, found, err := s.totpStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.Equals(datastore.Field("enabled"), true),
	).OrderBy(
		datastore.OrderByField("updatedAt", true),
	).Limit(1).FindOne()
	if err != nil || !found {
		return nil, found, err
	}
	return recordI.(*user.TOTP), true, nil
}

func (s *UserService) totpForUser(totpId string, userId string) (*user.TOTP, bool, error) {
	if totpId == "" {
		return nil, false, nil
	}

	recordI, found, err := s.totpStore.Query(
		datastore.Equals(datastore.Field("id"), totpId),
		datastore.Equals(datastore.Field("userId"), userId),
	).FindOne()
	if err != nil || !found {
		return nil, found, err
	}
	return recordI.(*user.TOTP), true, nil
}

func (s *UserService) pendingTOTPForUser(totpId string, userId string) (*user.TOTP, bool, error) {
	if totpId != "" {
		record, found, err := s.totpForUser(totpId, userId)
		if err != nil || !found {
			return nil, found, err
		}
		if record.Enabled || record.DisabledAt != nil {
			return nil, false, nil
		}
		return record, true, nil
	}

	recordI, found, err := s.totpStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.Equals(datastore.Field("enabled"), false),
	).OrderBy(
		datastore.OrderByField("createdAt", true),
	).Limit(1).FindOne()
	if err != nil || !found {
		return nil, found, err
	}
	record := recordI.(*user.TOTP)
	if record.DisabledAt != nil {
		return nil, false, nil
	}
	return record, true, nil
}

func (s *UserService) disableEnabledTOTPsForUser(userId string) error {
	records, err := s.totpStore.Query(
		datastore.Equals(datastore.Field("userId"), userId),
		datastore.Equals(datastore.Field("enabled"), true),
	).Find()
	if err != nil {
		return err
	}

	now := time.Now()
	for _, recordI := range records {
		record := recordI.(*user.TOTP)
		record.Enabled = false
		record.DisabledAt = &now
		record.UpdatedAt = now
		if err := s.totpStore.Upsert(record); err != nil {
			return err
		}
	}

	return nil
}

func totpSetupCreatedAfterActive(setup *user.TOTP, active *user.TOTP) bool {
	if active.EnabledAt != nil {
		return setup.CreatedAt.After(*active.EnabledAt)
	}
	return setup.CreatedAt.After(active.UpdatedAt)
}

func validateTOTPCode(secret string, code string) (bool, error) {
	return totp.ValidateCustom(code, secret, time.Now(), totp.ValidateOpts{
		Period:    totpPeriod,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}
