/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/samber/lo"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	data "github.com/1backend/1backend/server/internal/services/data/types"
)

// @ID queryObjects
// @Summary Query Objects
// @Description Retrieves objects from a specified table based on search criteria.
// @Description Requires authorization and user authentication.
// @Description
// @Description
// @Description Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
// @Tags Data Svc
// @Accept json
// @Produce json
// @Param body body data.QueryRequest false "Query Request"
// @Success 200 {object} data.QueryResponse "Successful retrieval of objects"
// @Failure 400 {object} data.ErrorResponse "Invalid JSON"
// @Failure 401 {object} data.ErrorResponse "Unauthorized"
// @Failure 500 {object} data.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /data-svc/objects [post]
func (g *DataService) Query(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := g.options.PermissionChecker.HasPermission(
		r,
		data.PermissionObjectView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &data.QueryRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Error decoding request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	claims, err := g.options.Authorizer.ParseJWTFromRequest(g.publicKey, r)
	if err != nil {
		logger.Error(
			"Error parsing JWT", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	for i, v := range req.Readers {
		if v == "_self" {
			req.Readers[i] = isAuthRsp.User.Id
		}
	}

	identifiers := append(
		claims.Roles,
		[]string{isAuthRsp.User.Id, data.AnyIdentifier}...)

	allowedReaders := identifiers
	if req.Readers != nil {
		allowedReaders = lo.Intersect(identifiers, req.Readers)
	}

	objects, err := g.query(allowedReaders, data.QueryOptions{
		Table: req.Table,
		Query: req.Query,
	})
	if err != nil {
		logger.Error(
			"Error querying objects",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(data.QueryResponse{
		Objects: objects,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (g *DataService) query(
	readers []string,
	options data.QueryOptions,
) ([]*data.Object, error) {
	if options.Table == "" {
		return nil, errors.New("no table name")
	}

	prependDataInQuery(options)

	filters := []datastore.Filter{}
	if options.Query != nil {
		filters = append(filters, options.Query.Filters...)
	}

	filters = append(filters,
		datastore.Equals(datastore.Field("table"), options.Table),
	)

	readersAny := []any{}
	for _, reader := range readers {
		readersAny = append(readersAny, reader)
	}
	filters = append(filters,
		datastore.Intersects(datastore.Field("readers"), readersAny),
	)

	q := g.store.Query(
		filters...,
	)

	if options.Query != nil {
		q.OrderBy(options.Query.OrderBys...)

		if options.Query.Limit != 0 {
			q.Limit(options.Query.Limit)
		}

		if options.Query.JSONAfter != "" {
			v := []any{}
			err := json.Unmarshal([]byte(options.Query.JSONAfter), &v)
			if err != nil {
				return nil, err
			}
			q = q.After(v...)
		}
	}

	objectIs, err := q.Find()
	if err != nil {
		return nil, err
	}

	objects := []*data.Object{}
	for _, objectI := range objectIs {
		objects = append(objects, objectI.(*data.Object))
	}

	return objects, nil
}

func prependDataInQuery(options data.QueryOptions) {
	if options.Query == nil {
		return
	}

	for i := range options.Query.Filters {
		for k := range options.Query.Filters[i].Fields {
			options.Query.Filters[i].Fields[k] = prependField(options.Query.Filters[i].Fields[k])
		}
	}

	for i := range options.Query.OrderBys {
		options.Query.OrderBys[i].Field = prependField(options.Query.OrderBys[i].Field)
	}
}

func prependField(field string) string {
	switch field {
	case "id", "readers", "writers", "deleters", "table":
		return field
	}

	return "data." + field
}
