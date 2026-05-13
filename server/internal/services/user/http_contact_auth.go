/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/1backend/1backend/server/internal/universe"
)

const (
	authStateExpiration = 10 * time.Minute
	authProvidersSecret = "auth-providers"

	// Individual secret names are useful for simple CLI/bootstrap setup:
	// auth-provider-google-client-id, auth-provider-google-client-secret, etc.
	authProviderSecretPrefix = "auth-provider-"

	googleOIDCIssuer    = "https://accounts.google.com"
	slackOIDCIssuer     = "https://slack.com"
	linkedInOIDCIssuer  = "https://www.linkedin.com"
	gitLabOIDCIssuer    = "https://gitlab.com"
	gitHubAuthorizeURL  = "https://github.com/login/oauth/authorize"
	gitHubTokenURL      = "https://github.com/login/oauth/access_token"
	gitHubAPIURL        = "https://api.github.com"
	discordAuthorizeURL = "https://discord.com/oauth2/authorize"
	discordTokenURL     = "https://discord.com/api/oauth2/token"
	discordAPIURL       = "https://discord.com/api"
)

var contactSlugCleanupRegexp = regexp.MustCompile(`[^a-z0-9]+`)

var relayEmailDomainSuffixes = []string{
	"@privaterelay.appleid.com",
	"@proxymail.facebook.com",
	"@users.noreply.github.com",
}

var builtinContactAuthProviders = map[string]contactAuthProvider{
	"google": {
		id:        "google",
		name:      "Google",
		kind:      "oidc",
		issuerURL: googleOIDCIssuer,
		scopes:    []string{oidc.ScopeOpenID, "email", "profile"},
	},
	"facebook": {
		id:           "facebook",
		name:         "Facebook",
		kind:         "facebook",
		scopes:       []string{"public_profile", "email"},
		graphVersion: "v25.0",
	},
	"slack": {
		id:        "slack",
		name:      "Slack",
		kind:      "oidc",
		issuerURL: slackOIDCIssuer,
		scopes:    []string{oidc.ScopeOpenID, "email", "profile"},
	},
	"linkedin": {
		id:        "linkedin",
		name:      "LinkedIn",
		kind:      "oidc",
		issuerURL: linkedInOIDCIssuer,
		scopes:    []string{oidc.ScopeOpenID, "email", "profile"},
	},
	"gitlab": {
		id:        "gitlab",
		name:      "GitLab",
		kind:      "oidc",
		issuerURL: gitLabOIDCIssuer,
		scopes:    []string{oidc.ScopeOpenID, "email", "profile"},
	},
	"github": {
		id:       "github",
		name:     "GitHub",
		kind:     "github",
		scopes:   []string{"user:email"},
		authURL:  gitHubAuthorizeURL,
		tokenURL: gitHubTokenURL,
		apiURL:   gitHubAPIURL,
	},
	"discord": {
		id:       "discord",
		name:     "Discord",
		kind:     "discord",
		scopes:   []string{"identify", "email"},
		authURL:  discordAuthorizeURL,
		tokenURL: discordTokenURL,
		apiURL:   discordAPIURL,
	},
}

type contactAuthSecretField struct {
	suffix string
	apply  func(*universe.ContactAuthProviderConfig, string)
}

var contactAuthSecretFields = []contactAuthSecretField{
	{
		suffix: "-client-secret",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.ClientSecret = value
		},
	},
	{
		suffix: "-client-id",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.ClientID = value
		},
	},
	{
		suffix: "-issuer-url",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.IssuerURL = value
		},
	},
	{
		suffix: "-graph-version",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.GraphVersion = value
		},
	},
	{
		suffix: "-auth-url",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.AuthURL = value
		},
	},
	{
		suffix: "-token-url",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.TokenURL = value
		},
	},
	{
		suffix: "-api-url",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.APIURL = value
		},
	},
	{
		suffix: "-scopes",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.Scopes = parseContactAuthScopes(value)
		},
	},
	{
		suffix: "-kind",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.Kind = value
		},
	},
	{
		suffix: "-name",
		apply: func(config *universe.ContactAuthProviderConfig, value string) {
			config.Name = value
		},
	},
}

type contactAuthProvider struct {
	id           string
	name         string
	kind         string
	issuerURL    string
	clientID     string
	clientSecret string
	scopes       []string
	graphVersion string
	authURL      string
	tokenURL     string
	apiURL       string
}

// @ID listContactAuthProviders
// @Summary List Contact Auth Providers
// @Description Lists configured external login providers that can be used as verified contact proof.
// @Tags User Svc
// @Produce json
// @Param appHost query string true "1backend app host"
// @Success 200 {object} user.ListContactAuthProvidersResponse "Configured providers"
// @Router /user-svc/auth/providers [get]
func (s *UserService) ListContactAuthProviders(w http.ResponseWriter, r *http.Request) {
	s.listContactAuthProviders(w, r)
}

// @ID listOIDCContactAuthProviders
// @Summary List OIDC Contact Auth Providers
// @Description Lists configured external login providers that can be used as verified contact proof.
// @Description This endpoint is kept as a compatibility alias for /user-svc/auth/providers.
// @Tags User Svc
// @Produce json
// @Param appHost query string true "1backend app host"
// @Success 200 {object} user.ListContactAuthProvidersResponse "Configured providers"
// @Router /user-svc/oidc/providers [get]
func (s *UserService) ListOIDCContactAuthProviders(w http.ResponseWriter, r *http.Request) {
	s.listContactAuthProviders(w, r)
}

func (s *UserService) listContactAuthProviders(w http.ResponseWriter, r *http.Request) {
	appHost := r.URL.Query().Get("appHost")
	if appHost == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "appHost missing")
		return
	}

	providers, err := s.configuredContactAuthProviderInfos(r.Context(), appHost)
	if err != nil {
		logger.Error("Failed to list contact auth providers", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ListContactAuthProvidersResponse{
		Providers: providers,
	})
}

// @ID contactAuthLogin
// @Summary Contact Auth Login
// @Description Verifies a configured external provider token as proof of contact ownership and returns a normal 1backend token.
// @Description OIDC providers must assert a real verified email contact. Relay and noreply addresses are rejected.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param provider path string true "Provider ID, such as google or facebook"
// @Param body body user.ContactAuthLoginRequest true "Contact Auth Login Request"
// @Success 200 {object} user.ContactAuthLoginResponse "Login successful"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/auth/{provider}/login [post]
func (s *UserService) ContactAuthLogin(w http.ResponseWriter, r *http.Request) {
	providerID := mux.Vars(r)["provider"]

	req := user.ContactAuthLoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode contact auth request", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.AppHost == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "AppHost missing")
		return
	}

	rawToken := firstNonEmpty(req.Token, req.IDToken, req.AccessToken)
	if rawToken == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "provider token missing")
		return
	}

	claims, err := s.verifyContactAuthToken(r.Context(), req.AppHost, providerID, rawToken)
	if err != nil {
		logger.Error(
			"Contact auth token verification failed",
			slog.String("provider", providerID),
			slog.Any("error", err),
		)
		endpoint.Unauthorized(w)
		return
	}

	app, err := s.getOrCreateApp(req.AppHost)
	if err != nil {
		logger.Error("Failed to get or create app", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	token, created, usr, err := s.loginWithVerifiedEmailContact(
		app.Id,
		claims.Email,
		claims.Name,
		req.Slug,
		req.Device,
	)
	if err != nil {
		logger.Error("Contact auth login failed", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	s.publishContactAuthLifecycleEvent(r.Context(), app, token, usr, created)
	endpoint.WriteJSON(w, http.StatusOK, user.ContactAuthLoginResponse{Token: token})
}

// @ID contactAuthStart
// @Summary Start Contact Auth
// @Description Redirects to a configured contact-auth provider. OIDC providers use authorization code flow; OAuth2 providers use their provider-specific verified email APIs.
// @Tags User Svc
// @Param provider path string true "Provider ID, such as google or facebook"
// @Param appHost query string true "1backend app host"
// @Param device query string false "Device"
// @Param slug query string false "Slug used only when creating a new user"
// @Param returnTo query string false "URL to redirect to after login. Token is returned in the URL fragment."
// @Success 302 "Redirect to provider"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Router /user-svc/auth/{provider}/start [get]
func (s *UserService) StartContactAuth(w http.ResponseWriter, r *http.Request) {
	providerID := mux.Vars(r)["provider"]
	appHost := r.URL.Query().Get("appHost")
	if appHost == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "appHost missing")
		return
	}
	returnTo := r.URL.Query().Get("returnTo")
	if err := validateContactAuthReturnTo(appHost, returnTo); err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, err.Error())
		return
	}

	provider, err := s.contactAuthProvider(r.Context(), appHost, providerID)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, err.Error())
		return
	}

	if provider.clientSecret == "" {
		endpoint.WriteString(w, http.StatusBadRequest, provider.name+" client secret missing")
		return
	}

	stateID, err := randomURLToken()
	if err != nil {
		logger.Error("Failed to create auth state", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}
	nonce, err := randomURLToken()
	if err != nil {
		logger.Error("Failed to create auth nonce", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	now := time.Now()
	authState := &user.AuthState{
		Id:        stateID,
		Provider:  provider.id,
		AppHost:   appHost,
		Device:    r.URL.Query().Get("device"),
		Slug:      r.URL.Query().Get("slug"),
		ReturnTo:  returnTo,
		Nonce:     nonce,
		CreatedAt: now,
		ExpiresAt: now.Add(authStateExpiration),
	}
	if err := s.authStateStore.Create(authState); err != nil {
		logger.Error("Failed to store auth state", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	redirectURL := s.contactAuthCallbackURL(provider.id)
	authURL, err := s.contactAuthCodeURL(r.Context(), provider, redirectURL, stateID, nonce)
	if err != nil {
		logger.Error("Failed to build auth URL", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	http.Redirect(w, r, authURL, http.StatusFound)
}

// @ID contactAuthCallback
// @Summary Contact Auth Callback
// @Description Handles a provider authorization-code callback and returns or redirects with a normal 1backend token.
// @Tags User Svc
// @Param provider path string true "Provider ID, such as google or facebook"
// @Success 200 {object} user.ContactAuthLoginResponse "Login successful"
// @Failure 400 {object} user.ErrorResponse "Invalid request"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Router /user-svc/auth/{provider}/callback [get]
func (s *UserService) ContactAuthCallback(w http.ResponseWriter, r *http.Request) {
	providerID := mux.Vars(r)["provider"]

	if providerErr := r.URL.Query().Get("error"); providerErr != "" {
		endpoint.WriteString(w, http.StatusBadRequest, providerErr)
		return
	}

	stateID := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	if stateID == "" || code == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "state and code are required")
		return
	}

	authState, err := s.consumeAuthState(stateID, providerID)
	if err != nil {
		logger.Error("Failed to consume auth state", slog.Any("error", err))
		endpoint.Unauthorized(w)
		return
	}

	provider, err := s.contactAuthProvider(r.Context(), authState.AppHost, providerID)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, err := s.exchangeContactAuthCode(
		r.Context(),
		provider,
		s.contactAuthCallbackURL(provider.id),
		code,
		authState.Nonce,
	)
	if err != nil {
		logger.Error("Failed to exchange auth code", slog.Any("error", err))
		endpoint.Unauthorized(w)
		return
	}

	app, err := s.getOrCreateApp(authState.AppHost)
	if err != nil {
		logger.Error("Failed to get or create app", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	token, created, usr, err := s.loginWithVerifiedEmailContact(
		app.Id,
		claims.Email,
		claims.Name,
		authState.Slug,
		authState.Device,
	)
	if err != nil {
		logger.Error("Contact auth callback login failed", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	s.publishContactAuthLifecycleEvent(r.Context(), app, token, usr, created)
	if authState.ReturnTo != "" {
		http.Redirect(w, r, authReturnURL(authState.ReturnTo, token), http.StatusFound)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.ContactAuthLoginResponse{Token: token})
}

func (s *UserService) verifyContactAuthToken(
	ctx context.Context,
	appHost string,
	providerID string,
	rawToken string,
) (*universe.ContactAuthClaims, error) {
	if isAppleProvider(providerID, "") {
		return nil, errors.New("Apple relay providers are not allowed")
	}

	if s.options.ContactAuthVerifier != nil {
		claims, err := s.options.ContactAuthVerifier.VerifyContactAuthToken(ctx, providerID, rawToken)
		if err != nil {
			return nil, err
		}
		return validateContactAuthClaims(providerID, claims)
	}

	provider, err := s.contactAuthProvider(ctx, appHost, providerID)
	if err != nil {
		return nil, err
	}

	switch provider.kind {
	case "oidc":
		return s.verifyOIDCToken(ctx, provider, rawToken)
	case "facebook":
		return s.verifyFacebookAccessToken(ctx, provider, rawToken)
	case "github":
		return s.verifyGitHubAccessToken(ctx, provider, rawToken)
	case "discord":
		return s.verifyDiscordAccessToken(ctx, provider, rawToken)
	default:
		return nil, fmt.Errorf("unsupported auth provider kind %q", provider.kind)
	}
}

func (s *UserService) contactAuthProvider(
	ctx context.Context,
	appHost string,
	providerID string,
) (*contactAuthProvider, error) {
	id := strings.ToLower(strings.TrimSpace(providerID))
	if id == "" {
		return nil, errors.New("provider missing")
	}
	if isAppleProvider(id, "") {
		return nil, errors.New("Apple relay providers are not allowed")
	}

	configs, err := s.contactAuthProviderConfigs(ctx, appHost)
	if err != nil {
		return nil, err
	}

	configured, ok := configs[id]
	if !ok {
		return nil, fmt.Errorf("auth provider %q is not configured", providerID)
	}

	provider, hasBuiltin := builtinContactAuthProviders[id]
	if !hasBuiltin {
		provider = contactAuthProvider{
			id:   id,
			name: configured.Name,
			kind: "oidc",
		}
	}

	applyContactAuthProviderConfig(&provider, configured, hasBuiltin)
	if isAppleProvider(provider.id, provider.issuerURL) {
		return nil, errors.New("Apple relay providers are not allowed")
	}
	if provider.name == "" {
		provider.name = provider.id
	}
	if provider.kind == "" {
		provider.kind = "oidc"
	}
	if provider.clientID == "" {
		return nil, fmt.Errorf("%s client id missing", provider.name)
	}
	if provider.kind == "oidc" {
		if provider.issuerURL == "" {
			return nil, fmt.Errorf("OIDC provider %q issuer URL missing", provider.id)
		}
		provider.issuerURL = strings.TrimRight(provider.issuerURL, "/")
		if len(provider.scopes) == 0 {
			provider.scopes = []string{oidc.ScopeOpenID, "email", "profile"}
		}
	}

	switch provider.kind {
	case "oidc", "facebook", "github", "discord":
		return &provider, nil
	default:
		return nil, fmt.Errorf("unsupported auth provider kind %q", provider.kind)
	}
}

func (s *UserService) configuredContactAuthProviderInfos(
	ctx context.Context,
	appHost string,
) ([]user.ContactAuthProviderInfo, error) {
	configs, err := s.contactAuthProviderConfigs(ctx, appHost)
	if err != nil {
		return nil, err
	}

	providerIDs := make([]string, 0, len(configs))
	for providerID := range configs {
		providerIDs = append(providerIDs, providerID)
	}
	sort.Strings(providerIDs)

	providers := []user.ContactAuthProviderInfo{}
	for _, providerID := range providerIDs {
		provider, err := s.contactAuthProvider(ctx, appHost, providerID)
		if err != nil || provider.clientSecret == "" {
			continue
		}
		providers = append(providers, contactAuthProviderInfo(provider))
	}

	return providers, nil
}

func (s *UserService) contactAuthProviderConfigs(
	ctx context.Context,
	appHost string,
) (map[string]universe.ContactAuthProviderConfig, error) {
	secretMap, err := s.contactAuthSecretMap(ctx, appHost)
	if err != nil {
		return nil, err
	}

	providers := map[string]universe.ContactAuthProviderConfig{}
	if raw := strings.TrimSpace(secretMap[authProvidersSecret]); raw != "" {
		if err := mergeContactAuthProvidersJSON(providers, raw); err != nil {
			return nil, errors.Wrapf(err, "failed to parse %s secret", authProvidersSecret)
		}
	}
	mergeIndividualContactAuthProviderSecrets(providers, secretMap)

	return providers, nil
}

func (s *UserService) contactAuthSecretMap(
	ctx context.Context,
	appHost string,
) (map[string]string, error) {
	backup, err := s.fetchSecretMap(ctx, s.token, nil)
	if err != nil {
		return nil, err
	}

	if appHost == "" || appHost == sdk.DefaultAppHost {
		return backup, nil
	}

	app, err := s.getOrCreateApp(appHost)
	if err != nil {
		return nil, err
	}

	exchangedToken, err := s.options.TokenExchanger.ExchangeToken(
		ctx,
		s.token,
		endpoint.ExchangeOptions{AppId: app.Id},
	)
	if err != nil {
		return nil, err
	}

	primary, err := s.fetchSecretMap(ctx, exchangedToken, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range primary {
		backup[key] = value
	}
	return backup, nil
}

func mergeContactAuthProvidersJSON(
	providers map[string]universe.ContactAuthProviderConfig,
	raw string,
) error {
	byID := map[string]universe.ContactAuthProviderConfig{}
	if err := json.Unmarshal([]byte(raw), &byID); err == nil {
		for providerID, provider := range byID {
			mergeContactAuthProvider(providers, providerID, provider)
		}
		return nil
	}

	list := []universe.ContactAuthProviderConfig{}
	if err := json.Unmarshal([]byte(raw), &list); err != nil {
		return err
	}
	for _, provider := range list {
		mergeContactAuthProvider(providers, provider.Id, provider)
	}
	return nil
}

func mergeIndividualContactAuthProviderSecrets(
	providers map[string]universe.ContactAuthProviderConfig,
	secrets map[string]string,
) {
	for secretID, value := range secrets {
		if !strings.HasPrefix(secretID, authProviderSecretPrefix) {
			continue
		}

		remainder := strings.TrimPrefix(secretID, authProviderSecretPrefix)
		for _, field := range contactAuthSecretFields {
			if !strings.HasSuffix(remainder, field.suffix) {
				continue
			}

			providerID := strings.TrimSuffix(remainder, field.suffix)
			if providerID == "" {
				continue
			}

			provider := universe.ContactAuthProviderConfig{Id: providerID}
			field.apply(&provider, value)
			mergeContactAuthProvider(providers, providerID, provider)
			break
		}
	}
}

func mergeContactAuthProvider(
	providers map[string]universe.ContactAuthProviderConfig,
	providerID string,
	provider universe.ContactAuthProviderConfig,
) {
	id := strings.ToLower(strings.TrimSpace(firstNonEmpty(provider.Id, providerID)))
	if id == "" {
		return
	}
	provider.Id = id
	provider.Kind = strings.ToLower(strings.TrimSpace(provider.Kind))

	existing := providers[id]
	existing.Id = id
	if provider.Name != "" {
		existing.Name = provider.Name
	}
	if provider.Kind != "" {
		existing.Kind = provider.Kind
	}
	if provider.IssuerURL != "" {
		existing.IssuerURL = provider.IssuerURL
	}
	if provider.ClientID != "" {
		existing.ClientID = provider.ClientID
	}
	if provider.ClientSecret != "" {
		existing.ClientSecret = provider.ClientSecret
	}
	if len(provider.Scopes) > 0 {
		existing.Scopes = provider.Scopes
	}
	if provider.GraphVersion != "" {
		existing.GraphVersion = provider.GraphVersion
	}
	if provider.AuthURL != "" {
		existing.AuthURL = provider.AuthURL
	}
	if provider.TokenURL != "" {
		existing.TokenURL = provider.TokenURL
	}
	if provider.APIURL != "" {
		existing.APIURL = provider.APIURL
	}
	providers[id] = existing
}

func parseContactAuthScopes(value string) []string {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}

	var scopes []string
	if err := json.Unmarshal([]byte(value), &scopes); err == nil {
		return scopes
	}

	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == ',' || r == ' ' || r == '\t' || r == '\n' || r == '\r'
	})
	scopes = scopes[:0]
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			scopes = append(scopes, part)
		}
	}
	return scopes
}

func applyContactAuthProviderConfig(
	provider *contactAuthProvider,
	config universe.ContactAuthProviderConfig,
	hasBuiltin bool,
) {
	if config.Name != "" {
		provider.name = config.Name
	}
	if !hasBuiltin && config.Kind != "" {
		provider.kind = strings.ToLower(strings.TrimSpace(config.Kind))
	}
	if config.IssuerURL != "" {
		provider.issuerURL = config.IssuerURL
	}
	if config.ClientID != "" {
		provider.clientID = config.ClientID
	}
	if config.ClientSecret != "" {
		provider.clientSecret = config.ClientSecret
	}
	if len(config.Scopes) > 0 {
		provider.scopes = config.Scopes
	}
	if config.GraphVersion != "" {
		provider.graphVersion = config.GraphVersion
	}
	if config.AuthURL != "" {
		provider.authURL = config.AuthURL
	}
	if config.TokenURL != "" {
		provider.tokenURL = config.TokenURL
	}
	if config.APIURL != "" {
		provider.apiURL = config.APIURL
	}
}

func contactAuthProviderInfo(provider *contactAuthProvider) user.ContactAuthProviderInfo {
	kind := provider.kind
	if kind != "oidc" {
		kind = "oauth2"
	}
	return user.ContactAuthProviderInfo{
		Id:       provider.id,
		Name:     provider.name,
		Kind:     kind,
		Scopes:   provider.scopes,
		StartUrl: "/user-svc/auth/" + url.PathEscape(provider.id) + "/start",
		LoginUrl: "/user-svc/auth/" + url.PathEscape(provider.id) + "/login",
	}
}

func (s *UserService) contactAuthCodeURL(
	ctx context.Context,
	provider *contactAuthProvider,
	redirectURL string,
	state string,
	nonce string,
) (string, error) {
	cfg, err := s.oauth2Config(ctx, provider, redirectURL)
	if err != nil {
		return "", err
	}

	opts := []oauth2.AuthCodeOption{}
	if provider.kind == "oidc" {
		opts = append(opts, oidc.Nonce(nonce))
	}

	return cfg.AuthCodeURL(state, opts...), nil
}

func (s *UserService) exchangeContactAuthCode(
	ctx context.Context,
	provider *contactAuthProvider,
	redirectURL string,
	code string,
	expectedNonce string,
) (*universe.ContactAuthClaims, error) {
	cfg, err := s.oauth2Config(ctx, provider, redirectURL)
	if err != nil {
		return nil, err
	}

	oauthToken, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	switch provider.kind {
	case "oidc":
		rawIDToken, _ := oauthToken.Extra("id_token").(string)
		if rawIDToken == "" {
			return nil, errors.New("OIDC id_token missing")
		}
		claims, err := s.verifyOIDCToken(ctx, provider, rawIDToken)
		if err != nil {
			return nil, err
		}
		if expectedNonce != "" && claims.Nonce != expectedNonce {
			return nil, errors.New("OIDC nonce mismatch")
		}
		return claims, nil
	case "facebook":
		return s.verifyFacebookAccessToken(ctx, provider, oauthToken.AccessToken)
	case "github":
		return s.verifyGitHubAccessToken(ctx, provider, oauthToken.AccessToken)
	case "discord":
		return s.verifyDiscordAccessToken(ctx, provider, oauthToken.AccessToken)
	default:
		return nil, fmt.Errorf("unsupported auth provider kind %q", provider.kind)
	}
}

func (s *UserService) oauth2Config(
	ctx context.Context,
	provider *contactAuthProvider,
	redirectURL string,
) (*oauth2.Config, error) {
	switch provider.kind {
	case "oidc":
		oidcProvider, err := oidc.NewProvider(ctx, provider.issuerURL)
		if err != nil {
			return nil, err
		}
		return &oauth2.Config{
			ClientID:     provider.clientID,
			ClientSecret: provider.clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       provider.scopes,
			Endpoint:     oidcProvider.Endpoint(),
		}, nil
	case "facebook":
		graphBase := facebookGraphBase(provider.graphVersion)
		return &oauth2.Config{
			ClientID:     provider.clientID,
			ClientSecret: provider.clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       provider.scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  facebookDialogURL(provider.graphVersion),
				TokenURL: strings.TrimRight(graphBase, "/") + "/oauth/access_token",
			},
		}, nil
	case "github", "discord":
		return &oauth2.Config{
			ClientID:     provider.clientID,
			ClientSecret: provider.clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       provider.scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  provider.authURL,
				TokenURL: provider.tokenURL,
			},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported auth provider kind %q", provider.kind)
	}
}

func (s *UserService) verifyOIDCToken(
	ctx context.Context,
	provider *contactAuthProvider,
	rawIDToken string,
) (*universe.ContactAuthClaims, error) {
	oidcProvider, err := oidc.NewProvider(ctx, provider.issuerURL)
	if err != nil {
		return nil, err
	}

	idToken, err := oidcProvider.Verifier(&oidc.Config{
		ClientID: provider.clientID,
	}).Verify(ctx, rawIDToken)
	if err != nil {
		return nil, err
	}

	rawClaims := map[string]any{}
	if err := idToken.Claims(&rawClaims); err != nil {
		return nil, err
	}

	claims := &universe.ContactAuthClaims{
		Provider:      provider.id,
		Email:         stringClaim(rawClaims, "email"),
		EmailVerified: boolClaim(rawClaims, "email_verified"),
		Name:          stringClaim(rawClaims, "name"),
		Nonce:         stringClaim(rawClaims, "nonce"),
	}

	return validateContactAuthClaims(provider.id, claims)
}

func (s *UserService) verifyFacebookAccessToken(
	ctx context.Context,
	provider *contactAuthProvider,
	accessToken string,
) (*universe.ContactAuthClaims, error) {
	if provider.clientSecret == "" {
		return nil, errors.New("Facebook app secret missing")
	}

	base := facebookGraphBase(provider.graphVersion)
	appAccessToken := provider.clientID + "|" + provider.clientSecret
	debugURL, err := url.Parse(strings.TrimRight(base, "/") + "/debug_token")
	if err != nil {
		return nil, err
	}
	debugValues := debugURL.Query()
	debugValues.Set("input_token", accessToken)
	debugValues.Set("access_token", appAccessToken)
	debugURL.RawQuery = debugValues.Encode()

	debugResp := struct {
		Data struct {
			AppID     string `json:"app_id"`
			IsValid   bool   `json:"is_valid"`
			UserID    string `json:"user_id"`
			ExpiresAt int64  `json:"expires_at"`
		} `json:"data"`
	}{}
	if err := getJSON(ctx, debugURL.String(), &debugResp); err != nil {
		return nil, err
	}
	if !debugResp.Data.IsValid {
		return nil, errors.New("Facebook token is not valid")
	}
	if debugResp.Data.AppID != provider.clientID {
		return nil, errors.New("Facebook token app id mismatch")
	}
	if debugResp.Data.ExpiresAt > 0 && time.Now().Unix() > debugResp.Data.ExpiresAt {
		return nil, errors.New("Facebook token expired")
	}

	meURL, err := url.Parse(strings.TrimRight(base, "/") + "/me")
	if err != nil {
		return nil, err
	}
	meValues := meURL.Query()
	meValues.Set("fields", "id,name,email")
	meValues.Set("access_token", accessToken)
	meValues.Set("appsecret_proof", appSecretProof(accessToken, provider.clientSecret))
	meURL.RawQuery = meValues.Encode()

	meResp := struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	if err := getJSON(ctx, meURL.String(), &meResp); err != nil {
		return nil, err
	}
	if meResp.ID == "" || meResp.ID != debugResp.Data.UserID {
		return nil, errors.New("Facebook user id mismatch")
	}

	return validateContactAuthClaims(provider.id, &universe.ContactAuthClaims{
		Provider:      provider.id,
		Email:         meResp.Email,
		EmailVerified: meResp.Email != "",
		Name:          meResp.Name,
	})
}

func (s *UserService) verifyGitHubAccessToken(
	ctx context.Context,
	provider *contactAuthProvider,
	accessToken string,
) (*universe.ContactAuthClaims, error) {
	if provider.clientSecret == "" {
		return nil, errors.New("GitHub client secret missing")
	}

	checkURL := strings.TrimRight(provider.apiURL, "/") +
		"/applications/" + url.PathEscape(provider.clientID) + "/token"
	checkResp := struct {
		Scopes []string `json:"scopes"`
		App    struct {
			ClientID string `json:"client_id"`
		} `json:"app"`
		User struct {
			Login string `json:"login"`
		} `json:"user"`
	}{}
	err := postJSONWithBasicAuth(
		ctx,
		checkURL,
		provider.clientID,
		provider.clientSecret,
		map[string]string{"access_token": accessToken},
		&checkResp,
		gitHubHeaders(),
	)
	if err != nil {
		return nil, err
	}
	if checkResp.App.ClientID != provider.clientID {
		return nil, errors.New("GitHub token app id mismatch")
	}
	if !hasAnyScope(checkResp.Scopes, "user:email", "user") {
		return nil, errors.New("GitHub token is missing email scope")
	}

	emails := []struct {
		Email    string `json:"email"`
		Verified bool   `json:"verified"`
		Primary  bool   `json:"primary"`
	}{}
	if err := getJSONWithBearer(
		ctx,
		strings.TrimRight(provider.apiURL, "/")+"/user/emails",
		accessToken,
		&emails,
		gitHubHeaders(),
	); err != nil {
		return nil, err
	}

	email := ""
	for _, next := range emails {
		if next.Primary && next.Verified && !isRelayEmail(next.Email) {
			email = next.Email
			break
		}
	}
	if email == "" {
		for _, next := range emails {
			if next.Verified && !isRelayEmail(next.Email) {
				email = next.Email
				break
			}
		}
	}

	return validateContactAuthClaims(provider.id, &universe.ContactAuthClaims{
		Provider:      provider.id,
		Email:         email,
		EmailVerified: email != "",
		Name:          checkResp.User.Login,
	})
}

func (s *UserService) verifyDiscordAccessToken(
	ctx context.Context,
	provider *contactAuthProvider,
	accessToken string,
) (*universe.ContactAuthClaims, error) {
	authResp := struct {
		Application struct {
			ID string `json:"id"`
		} `json:"application"`
		Scopes  []string `json:"scopes"`
		Expires string   `json:"expires"`
		User    struct {
			ID string `json:"id"`
		} `json:"user"`
	}{}
	if err := getJSONWithBearer(
		ctx,
		strings.TrimRight(provider.apiURL, "/")+"/oauth2/@me",
		accessToken,
		&authResp,
		nil,
	); err != nil {
		return nil, err
	}
	if authResp.Application.ID != provider.clientID {
		return nil, errors.New("Discord token app id mismatch")
	}
	if !hasScope(authResp.Scopes, "identify") || !hasScope(authResp.Scopes, "email") {
		return nil, errors.New("Discord token is missing identify or email scope")
	}
	if authResp.Expires != "" {
		expiresAt, err := time.Parse(time.RFC3339Nano, authResp.Expires)
		if err == nil && time.Now().After(expiresAt) {
			return nil, errors.New("Discord token expired")
		}
	}

	meResp := struct {
		ID         string `json:"id"`
		Username   string `json:"username"`
		GlobalName string `json:"global_name"`
		Email      string `json:"email"`
		Verified   bool   `json:"verified"`
	}{}
	if err := getJSONWithBearer(
		ctx,
		strings.TrimRight(provider.apiURL, "/")+"/users/@me",
		accessToken,
		&meResp,
		nil,
	); err != nil {
		return nil, err
	}
	if authResp.User.ID != "" && meResp.ID != "" && authResp.User.ID != meResp.ID {
		return nil, errors.New("Discord user id mismatch")
	}

	name := firstNonEmpty(meResp.GlobalName, meResp.Username)
	return validateContactAuthClaims(provider.id, &universe.ContactAuthClaims{
		Provider:      provider.id,
		Email:         meResp.Email,
		EmailVerified: meResp.Verified,
		Name:          name,
	})
}

func (s *UserService) loginWithVerifiedEmailContact(
	appId string,
	email string,
	name string,
	requestedSlug string,
	device string,
) (*user.Token, bool, *user.User, error) {
	email = normalizeEmail(strings.TrimSpace(email))
	if email == "" || !isEmail(email) {
		return nil, false, nil, errors.New("verified email contact is required")
	}
	if isRelayEmail(email) {
		return nil, false, nil, errors.New("relay email contacts are not allowed")
	}
	if device == "" {
		device = unknownDevice
	}

	contactI, found, err := s.contactStore.Query(
		datastore.Equals(datastore.Field("id"), email),
	).FindOne()
	if err != nil {
		return nil, false, nil, err
	}

	if found {
		contact := contactI.(*user.Contact)
		if !contact.Verified || contact.Platform == "" {
			contact.Verified = true
			contact.Platform = "email"
			contact.Handle = email
			contact.UpdatedAt = time.Now()
			if err := s.contactStore.Query(datastore.Id(contact.Id)).Update(contact); err != nil {
				return nil, false, nil, errors.Wrap(err, "failed to update verified contact")
			}
		}

		userI, userFound, err := s.userStore.Query(
			datastore.Equals(datastore.Field("id"), contact.UserId),
		).FindOne()
		if err != nil {
			return nil, false, nil, err
		}
		if !userFound {
			return nil, false, nil, fmt.Errorf("contact %q has missing user %q", contact.Id, contact.UserId)
		}
		usr := userI.(*user.User)
		token, err := s.issueToken(appId, usr, device)
		return token, false, usr, err
	}

	slug, err := s.slugForNewContactUser(email, requestedSlug)
	if err != nil {
		return nil, false, nil, err
	}

	now := time.Now()
	userInput := &user.UserInput{
		Id:   sdk.Id("usr"),
		Name: name,
		Slug: slug,
	}
	if userInput.Name == "" {
		userInput.Name = email
	}

	err = s.createUserWithoutVerification(
		appId,
		userInput,
		[]user.Contact{
			{
				Id:        email,
				CreatedAt: now,
				UpdatedAt: now,
				Platform:  "email",
				Handle:    email,
				Verified:  true,
				IsPrimary: true,
			},
		},
		"",
		nil,
	)
	if err != nil {
		return nil, false, nil, err
	}

	userI, found, err := s.userStore.Query(
		datastore.Equals(datastore.Field("id"), userInput.Id),
	).FindOne()
	if err != nil {
		return nil, false, nil, err
	}
	if !found {
		return nil, false, nil, errors.New("created user not found")
	}
	usr := userI.(*user.User)

	token, err := s.issueToken(appId, usr, device)
	return token, true, usr, err
}

func (s *UserService) slugForNewContactUser(email string, requestedSlug string) (string, error) {
	if requestedSlug != "" {
		if !SlugRegexp.MatchString(requestedSlug) {
			return "", errors.New("slug must be lowercase and can only contain letters, numbers, and dashes")
		}
		if isReservedUserSlug(requestedSlug) {
			return "", errors.New("slug is reserved")
		}
		_, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("slug"), requestedSlug),
		).FindOne()
		if err != nil {
			return "", err
		}
		if found {
			return "", errors.New("slug already exists")
		}
		return requestedSlug, nil
	}

	base := strings.Split(email, "@")[0]
	base = strings.ToLower(base)
	base = contactSlugCleanupRegexp.ReplaceAllString(base, "-")
	base = strings.Trim(base, "-")
	if base == "" || isReservedUserSlug(base) || !SlugRegexp.MatchString(base) {
		base = "contact"
	}

	for i := 0; i < 1000; i++ {
		candidate := base
		if i > 0 {
			candidate = fmt.Sprintf("%s-%d", base, i+1)
		}
		_, found, err := s.userStore.Query(
			datastore.Equals(datastore.Field("slug"), candidate),
		).FindOne()
		if err != nil {
			return "", err
		}
		if !found {
			return candidate, nil
		}
	}

	return "", errors.New("failed to allocate unique slug")
}

func (s *UserService) consumeAuthState(stateID string, providerID string) (*user.AuthState, error) {
	stateI, found, err := s.authStateStore.Query(datastore.Id(stateID)).FindOne()
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("auth state not found")
	}
	authState := stateI.(*user.AuthState)
	if !strings.EqualFold(authState.Provider, providerID) {
		return nil, errors.New("auth state provider mismatch")
	}
	if authState.Used {
		return nil, errors.New("auth state already used")
	}
	if time.Now().After(authState.ExpiresAt) {
		return nil, errors.New("auth state expired")
	}

	err = s.authStateStore.Query(datastore.Id(authState.Id)).UpdateFields(map[string]any{
		"used": true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to mark auth state used")
	}

	return authState, nil
}

func validateContactAuthClaims(
	providerID string,
	claims *universe.ContactAuthClaims,
) (*universe.ContactAuthClaims, error) {
	if claims == nil {
		return nil, errors.New("provider claims are missing")
	}
	claims.Provider = providerID
	claims.Email = normalizeEmail(strings.TrimSpace(claims.Email))
	if claims.Email == "" || !isEmail(claims.Email) {
		return nil, errors.New("provider did not return an email contact")
	}
	if isRelayEmail(claims.Email) {
		return nil, errors.New("relay email contacts are not allowed")
	}
	if !claims.EmailVerified {
		return nil, errors.New("provider did not verify email contact")
	}
	return claims, nil
}

func (s *UserService) publishContactAuthLifecycleEvent(
	ctx context.Context,
	app *user.App,
	token *user.Token,
	usr *user.User,
	created bool,
) {
	if s.options.PubSub == nil || token == nil {
		return
	}

	topic := "user.login"
	if created {
		topic = "user.register"
	}

	slug := ""
	if usr != nil {
		slug = usr.Slug
	}

	evt := userLifecycleEventPayload(app, token, slug, time.Now().UTC())
	payload, _ := json.Marshal(evt)
	if _, err := s.options.PubSub.Publish(ctx, topic, payload); err != nil {
		logger.Error("Failed to publish contact auth event", slog.Any("error", err))
	}
}

func (s *UserService) contactAuthCallbackURL(providerID string) string {
	return strings.TrimRight(s.options.Url, "/") + "/user-svc/auth/" + url.PathEscape(providerID) + "/callback"
}

func authReturnURL(rawReturnTo string, token *user.Token) string {
	u, err := url.Parse(rawReturnTo)
	if err != nil {
		return rawReturnTo
	}

	fragment := url.Values{}
	fragment.Set("token", token.Token)
	fragment.Set("tokenId", token.Id)
	fragment.Set("userId", token.UserId)
	fragment.Set("appId", token.AppId)
	fragment.Set("expiresAt", token.ExpiresAt.Format(time.RFC3339))
	u.Fragment = fragment.Encode()

	return u.String()
}

func validateContactAuthReturnTo(appHost string, rawReturnTo string) error {
	if rawReturnTo == "" {
		return nil
	}

	u, err := url.Parse(rawReturnTo)
	if err != nil {
		return errors.New("returnTo is invalid")
	}
	if u.Host == "" && u.Scheme == "" {
		return nil
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("returnTo must use http or https")
	}

	appURL, err := url.Parse(appHost)
	if err != nil {
		return errors.New("appHost is invalid")
	}
	expectedHost := appURL.Host
	if expectedHost == "" {
		expectedHost = appHost
	}
	expectedHost = strings.ToLower(strings.TrimSpace(expectedHost))
	actualHost := strings.ToLower(strings.TrimSpace(u.Host))
	if actualHost != expectedHost {
		return errors.New("returnTo host must match appHost")
	}

	return nil
}

func randomURLToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func facebookGraphBase(version string) string {
	version = strings.Trim(version, "/")
	if version == "" {
		version = "v25.0"
	}
	return "https://graph.facebook.com/" + version
}

func facebookDialogURL(version string) string {
	version = strings.Trim(version, "/")
	if version == "" {
		version = "v25.0"
	}
	return "https://www.facebook.com/" + version + "/dialog/oauth"
}

func appSecretProof(accessToken string, appSecret string) string {
	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write([]byte(accessToken))
	return hex.EncodeToString(mac.Sum(nil))
}

func getJSON(ctx context.Context, rawURL string, target any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("provider returned status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func getJSONWithBearer(
	ctx context.Context,
	rawURL string,
	bearerToken string,
	target any,
	headers map[string]string,
) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Accept", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("provider returned status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func postJSONWithBasicAuth(
	ctx context.Context,
	rawURL string,
	username string,
	password string,
	payload any,
	target any,
	headers map[string]string,
) error {
	bs, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("provider returned status %d", resp.StatusCode)
	}
	if target == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func gitHubHeaders() map[string]string {
	return map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}
}

func hasScope(scopes []string, required string) bool {
	for _, scope := range scopes {
		if scope == required {
			return true
		}
	}
	return false
}

func hasAnyScope(scopes []string, allowed ...string) bool {
	for _, scope := range scopes {
		for _, allowedScope := range allowed {
			if scope == allowedScope {
				return true
			}
		}
	}
	return false
}

func stringClaim(claims map[string]any, key string) string {
	value, _ := claims[key].(string)
	return value
}

func boolClaim(claims map[string]any, key string) bool {
	value, ok := claims[key]
	if !ok {
		return false
	}
	switch typed := value.(type) {
	case bool:
		return typed
	case string:
		return strings.EqualFold(typed, "true")
	default:
		return false
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func isAppleProvider(providerID string, issuerURL string) bool {
	providerID = strings.ToLower(providerID)
	issuerURL = strings.ToLower(issuerURL)
	return strings.Contains(providerID, "apple") ||
		strings.Contains(issuerURL, "appleid.apple.com")
}

func isRelayEmail(email string) bool {
	email = strings.ToLower(strings.TrimSpace(email))
	for _, suffix := range relayEmailDomainSuffixes {
		if strings.HasSuffix(email, suffix) {
			return true
		}
	}
	return false
}
