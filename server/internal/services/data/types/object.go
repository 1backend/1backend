package data_svc

import (
	"time"

	"github.com/openorch/openorch/sdk/go/datastore"
)

const AnyIdentifier string = "any"

type ErrorResponse struct {
	Error string `json:"error"`
}

// Object holds any kind of data, so
// we don't have to implement simple CRUD for
// any new simple entity.
type Object struct {
	Id    string `json:"id"`
	Table string `json:"table" binding:"required"`

	// Authors is a list of user ID and organization ID who created the object.
	// The authors field tracks which users or organizations created an entry, helping to prevent spam.
	// If an organization ID is not provided, the currently active organization will
	// be queried from the User Svc.
	Authors []string `json:"authors" example:"[\"usr_12345\", \"org_67890\"]"`

	// Readers is a list of user IDs and role IDs that can read the object.
	// `_self` can be used to refer to the caller user's userId and
	// `_org` can be used to refer to the user's currently active organization (if exists).
	Readers []string `json:"readers,omitempty" example:"[\"usr_12345\", \"org_67890\"]"`

	// Writers is a list of user IDs and role IDs that can write the object.
	// `_self` can be used to refer to the caller user's userId and
	// `_org` can be used to refer to the user's currently active organization (if exists).
	Writers []string `json:"writers,omitempty" example:"[\"usr_12345\", \"org_67890\"]"`

	// Deleters is a list of user IDs and role IDs that can delete the object.
	// `_self` can be used to refer to the caller user's userId and
	// `_org` can be used to refer to the user's currently active organization (if exists).
	Deleters []string `json:"deleters,omitempty" example:"[\"usr_12345\", \"org_67890\"]"`

	Data map[string]interface{} `json:"data,omitempty" binding:"required"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (g Object) GetId() string {
	return g.Id
}

type QueryRequest struct {
	Table   string           `json:"table"`
	Query   *datastore.Query `json:"query"`
	Readers []string         `json:"readers,omitempty"`
}

type QueryOptions struct {
	Table string
	Query *datastore.Query
}

type QueryResponse struct {
	Objects []*Object `json:"objects,omitempty"`
}

type CreateObjectFields struct {
	Id    string `json:"id"`
	Table string `json:"table" binding:"required"`

	// Authors is a list of user ID and organization ID who created the object.
	// If an organization ID is not provided, the currently active organization will
	// be queried from the User Svc.
	Authors []string `json:"authors" example:"[\"usr_12345\", \"org_67890\"]"`

	// Readers is a list of user IDs and role IDs that can read the object.
	// `_self` can be used to refer to the caller user's userId and
	// `_org` can be used to refer to the user's currently active organization (if exists).
	Readers []string `json:"readers,omitempty"`

	// Writers is a list of user IDs and role IDs that can write the object.
	// `_self` can be used to refer to the caller user's userId and
	// `_org` can be used to refer to the user's currently active organization (if exists).
	Writers []string `json:"writers,omitempty"`

	// Deleters is a list of user IDs and role IDs that can delete the object.
	// `_self` can be used to refer to the caller user's userId and
	// `_org` can be used to refer to the user's currently active organization (if exists).
	Deleters []string `json:"deleters,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" binding:"required"`
}

func (g CreateObjectFields) GetId() string {
	return g.Id
}

type CreateObjectRequest struct {
	Object *CreateObjectFields `json:"object,omitempty"`
}

type CreateObjectResponse struct {
	Object *Object `json:"object,omitempty"`
}

type CreateManyRequest struct {
	Objects []*CreateObjectFields `json:"objects,omitempty"`
}

type UpsertObjectRequest struct {
	Object *CreateObjectFields `json:"object,omitempty"`
}

type UpsertObjectResponse struct {
	Object *Object `json:"object,omitempty"`
}

type UpsertManyRequest struct {
	Objects []*Object `json:"objects,omitempty"`
}

type UpsertManyResponse struct {
	Objects []*Object `json:"objects,omitempty"`
}

type DeleteObjectRequest struct {
	Table   string             `json:"table"`
	Filters []datastore.Filter `json:"filters"`
}

type DeleteObjectResponse struct {
}

type UpdateObjectsRequest struct {
	Table string `json:"table,omitempty"`

	// Filters to determine which objects will be updated.
	// Only objects matching all filters will be modified.
	Filters []datastore.Filter `json:"filters,omitempty"`

	// The object containing the fields to update in matching objects.
	Object *Object `json:"object,omitempty"`
}

type UpdateObjectsResponse struct {
}

type UpsertObjectsRequest struct {
	Table string `json:"table,omitempty"`

	Objects []Object `json:"objects,omitempty"`
}

type UpsertObjectsResponse struct{}
