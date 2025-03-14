/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package promptservice

import (
	"net/http"

	prompt "github.com/openorch/openorch/server/internal/services/prompt/types"
)

// @ID promptTypes
// @Summary Prompt Types
// @Description The only purpose of this "endpoint" is to export types otherwise not appearing in the API docs.
// @Description This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
// @Tags Prompt Svc
// @Accept json
// @Produce json
// @Param body body prompt.TypesRequest true "Types Request"
// @Success 200 {object} prompt.TypesResponse
// @Failure 400 {object} prompt.ErrorResponse "Invalid JSON"
// @Failure 401 {object} prompt.ErrorResponse "Unauthorized"
// @Failure 500 {object} prompt.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /prompt-svc/types [post]
func (p *PromptService) Types(
	w http.ResponseWriter,
	r *http.Request,
) {

	// only here so the prompt package is not unimported by the tooling
	dummy := prompt.ErrorResponse{}

	w.Write([]byte(`{}` + dummy.Error))
}
