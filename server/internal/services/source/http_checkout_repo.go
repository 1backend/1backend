package sourceservice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	source "github.com/1backend/1backend/server/internal/services/source/types"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	ghttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	gssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

// @ID checkoutRepo
// @Summary Checkout a git repository
// @Description Checkout a git repository over https or ssh at a specific version into a temporary directory.
// @Description Performs a shallow clone with minimal history for faster checkout.
// @Tags Source Svc
// @Accept json
// @Produce json
// @Param body body source.CheckoutRepoRequest true "Checkout Repo Request"
// @Success 200 {object} source.CheckoutRepoResponse "Successfully checked out the repository"
// @Failure 400 {object} source.ErrorResponse "Invalid JSON"
// @Failure 401 {object} source.ErrorResponse "Unauthorized"
// @Failure 500 {object} source.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /source-svc/repo/checkout [post]
func (s *SourceService) CheckoutRepo(w http.ResponseWriter,
	r *http.Request) {

	isAuthRsp, statusCode, err := s.options.PermissionChecker.HasPermission(
		r,
		source.PermissionSourceRepoCheckout,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &source.CheckoutRepoRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error(
			"Failed to decode request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	dir, err := s.checkoutRepo(*req)
	if err != nil {
		logger.Error(
			"Failed to checkout repo",
			slog.String("repo", req.URL),
		)
		endpoint.InternalServerError(w)
		return
	}

	response := source.CheckoutRepoResponse{
		Dir: dir,
	}

	bs, _ := json.Marshal(response)
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

// checkoutRepo checks out a repository at a specified version using the given auth method.
func (s *SourceService) checkoutRepo(
	req source.CheckoutRepoRequest,
) (string, error) {
	if req.Version == "" {
		req.Version = "main"
	}

	tempDir := path.Join(
		os.TempDir(),
		fmt.Sprintf("repo-%s-%s", makeFilenameSafeURL(req.URL), req.Version),
	)

	if !isCommitHash(req.Version) {
		// delete branch checkouts as they might be outdated
		err := os.RemoveAll(tempDir)
		if err != nil {
			return "", fmt.Errorf("failed to remove temp dir: %w", err)
		}
	}

	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}

	refName := plumbing.NewBranchReferenceName(req.Version)
	if strings.HasPrefix(req.Version, "refs/") ||
		len(req.Version) == 40 { // For tag or commit SHA
		refName = plumbing.ReferenceName(req.Version)
	}

	var authMethod transport.AuthMethod
	if req.SSHKey != "" {
		authMethod, err = getSSHAuth(req.SSHKey, req.SSHKeyPwd, req.Username)
		if err != nil {
			return "", fmt.Errorf("failed to set up SSH auth: %w", err)
		}
	} else if req.Token != "" {
		authMethod = &ghttp.BasicAuth{
			Username: "x-access-token", // GitHub specific
			Password: req.Token,
		}
	} else {
		authMethod = &ghttp.BasicAuth{
			Username: req.Username,
			Password: req.Password,
		}
	}

	cloneOptions := &git.CloneOptions{
		URL:           req.URL,
		Depth:         1, // Shallow clone
		SingleBranch:  true,
		ReferenceName: refName,
		Auth:          authMethod,
		NoCheckout:    false,
	}

	_, err = git.PlainClone(tempDir, false, cloneOptions)
	if err != nil {
		return "", fmt.Errorf("failed to clone repo: %w", err)
	}

	return tempDir, nil
}

func getSSHAuth(
	sshKey string,
	sshKeyPwd string,
	username string,
) (transport.AuthMethod, error) {
	signer, err := ssh.ParsePrivateKey([]byte(sshKey))
	if err != nil {
		// If password is provided for private key
		if sshKeyPwd != "" {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(
				[]byte(sshKey),
				[]byte(sshKeyPwd),
			)
			if err != nil {
				return nil, err
			}
		}
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &gssh.PublicKeys{
		User:   username,
		Signer: signer,
	}, nil
}

func makeFilenameSafeURL(urlStr string) string {
	replacer := strings.NewReplacer(
		"/", "_",
		":", "_",
		"?", "_",
		"&", "_",
		"=", "_",
	)
	return replacer.Replace(urlStr)
}

func isCommitHash(s string) bool {
	// Check if the string is 40 characters long and contains only valid hex characters
	if len(s) == 40 {
		for _, c := range s {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				return false
			}
		}
		return true
	}
	return false
}
