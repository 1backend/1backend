/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	config "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

// @Id listConfigVersions
// @Summary List Versions
// @Description Retrieves the current configurations for a specified app.
// @Description Since any user can save configurations, it is strongly advised that you supply a list of
// @Description owners to filter on.
// @Description If no app is specified, the default "unnamed" app is used.
// @Description This is a public endpoint and does not require authentication.
// @Description Configuration data is non-sensitive. For sensitive data, refer to the Secret Service.
// @Description
// @Description Configurations are used to control frontend behavior, A/B testing, feature flags, and other non-sensitive settings.
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param body body config.ListVersionsRequest true "List Configs Request"
// @Success 200 {object} config.ListVersionsResponse "Current configuration"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /config-svc/versions [post]
func (cs *ConfigService) ListVersions(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Version get should not be authorized because
	// it is public, nonsensitive information.

	req := &config.ListVersionsRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error("Failed to decode request body", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.AppHost == "" {
		logger.Error("AppHost missing in list configs request")
		endpoint.WriteErr(w, http.StatusBadRequest, errors.New("AppHost missing"))
		return
	}

	appId, err := cs.options.TokenExchanger.AppIdFromHost(
		r.Context(),
		req.AppHost,
	)
	if err != nil {
		logger.Error(
			"Failed to get app id from host",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	vers, err := cs.listVersions(appId, req)
	if err != nil {
		logger.Error("Failed to get version", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(config.ListVersionsResponse{
		Versions: vers,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *ConfigService) listVersions(
	appId string,
	req *config.ListVersionsRequest,
) ([]*config.Version, error) {
	ids := []any{}
	for _, slug := range req.Ids {
		slug = kebabToCamel(slug)
		ids = append(ids,
			slug,
		)
	}
	for slug := range req.Selector {
		slug = kebabToCamel(slug)
		ids = append(ids,
			slug,
		)
	}

	filters := []datastore.Filter{}

	if len(ids) > 0 {
		filters = append(filters, datastore.IsInList(
			datastore.Field("id"),
			ids...,
		))
	}

	if req.AppHost != "" {
		filters = append(filters, datastore.Equals(
			datastore.Field("appId"),
			appId,
		))
	}

	q := cs.versionStore.Query(
		filters...,
	).OrderBy(
		datastore.OrderByField("createdAt", true),
		datastore.OrderByField("id", true),
	)
	if req.AfterJson != "" {
		var afterCursor []any
		err := json.Unmarshal([]byte(req.AfterJson), &afterCursor)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal after cursor")
		}
		if len(afterCursor) != 2 {
			return nil, errors.New("after cursor must have exactly 2 elements")
		}

		q = q.After(afterCursor...)
	}
	if req.Limit != 0 {
		q = q.Limit(int64(req.Limit))
	} else {
		q = q.Limit(20)
	}

	versionIs, err := q.Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query configs")
	}

	ret := make([]*config.Version, 0, len(versionIs))

	for _, configI := range versionIs {
		conf := configI.(*config.Version)
		conf.InternalId = ""

		if conf.DataJSON != "" {
			if req.Selector != nil && req.Selector[conf.Id] != nil {
				conf.Data = map[string]any{}
				for _, path := range req.Selector[conf.Id] {
					res := gjson.Get(conf.DataJSON, path)
					if res.Exists() {
						m := ExpandDotPath(path, res.Value())
						conf.Data = DeepMerge(conf.Data, m)
					}
				}
			} else {
				err = json.Unmarshal([]byte(conf.DataJSON), &conf.Data)
				if err != nil {
					logger.Error("Failed to unmarshal version data",
						slog.Any("error", err),
						slog.String("dataJson", conf.DataJSON),
					)

					return nil, errors.Wrap(err, "failed to unmarshal version")
				}
			}
		}

		conf.DataJSON = ""
		ret = append(ret, conf)
	}

	return ret, nil
}
