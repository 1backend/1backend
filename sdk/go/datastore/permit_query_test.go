package datastore_test

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/1backend/1backend/sdk/go/datastore"
	localstore "github.com/1backend/1backend/sdk/go/datastore/localstore"
	"github.com/1backend/1backend/sdk/go/datastore/sqlstore"
	"github.com/1backend/1backend/sdk/go/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type PermissionQueryTestObject struct {
	Id          string   `json:"id"`
	AppId       string   `json:"appId"`
	Permission  string   `json:"permission,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Slugs       []string `json:"slugs,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}

func (p PermissionQueryTestObject) GetId() string {
	return p.Id
}

func runPermissionQueryShapeTest(t *testing.T, store datastore.DataStore) {
	permLegacyRole := PermissionQueryTestObject{
		Id:         "legacy-role",
		AppId:      "app-a",
		Permission: "perm:edit",
		Roles:      []string{"role:a"},
	}
	permArraySlug := PermissionQueryTestObject{
		Id:          "array-slug",
		AppId:       "*",
		Permissions: []string{"perm:view", "perm:edit"},
		Slugs:       []string{"alice"},
	}
	wrongApp := PermissionQueryTestObject{
		Id:         "wrong-app",
		AppId:      "app-b",
		Permission: "perm:edit",
		Roles:      []string{"role:a"},
	}
	wrongSubject := PermissionQueryTestObject{
		Id:          "wrong-subject",
		AppId:       "app-a",
		Permissions: []string{"perm:edit"},
		Roles:       []string{"role:b"},
	}

	require.NoError(t, store.CreateMany([]datastore.Row{
		permLegacyRole,
		permArraySlug,
		wrongApp,
		wrongSubject,
	}))

	roleValues := []any{"role:a", "role:z"}
	results, err := store.Query(
		datastore.Or(
			datastore.Equals(datastore.Field("appId"), "app-a"),
			datastore.Equals(datastore.Field("appId"), "*"),
		),
		datastore.Or(
			datastore.Equals(datastore.Field("permission"), "perm:edit"),
			datastore.Equals(datastore.Field("permissions"), "perm:edit"),
		),
		datastore.Or(
			datastore.Equals(datastore.Field("slugs"), "alice"),
			datastore.Intersects(datastore.Field("roles"), roleValues),
		),
	).Find()
	require.NoError(t, err)
	require.Len(t, results, 2)
	require.Contains(t, results, permLegacyRole)
	require.Contains(t, results, permArraySlug)
	require.NotContains(t, results, wrongApp)
	require.NotContains(t, results, wrongSubject)
}

func TestPermissionQueryShapeLocalStore(t *testing.T) {
	store, err := localstore.NewLocalStore(PermissionQueryTestObject{}, "")
	require.NoError(t, err)

	runPermissionQueryShapeTest(t, store)
}

func TestPermissionQueryShapeSQLStore(t *testing.T) {
	defer func() {
		if recovered := recover(); recovered != nil {
			t.Skipf("docker-backed sqlstore unavailable: %v", recovered)
		}
	}()

	pgConn := testutil.StartPostgres(t)
	table := strings.ReplaceAll(uuid.New().String(), "-", "")[:10]

	db, err := sql.Open("postgres", pgConn)
	require.NoError(t, err)

	store, err := sqlstore.NewSQLStore(
		PermissionQueryTestObject{},
		sqlstore.DriverPostGRES,
		db,
		"table_"+table,
		true,
	)
	require.NoError(t, err)

	runPermissionQueryShapeTest(t, store)
}
