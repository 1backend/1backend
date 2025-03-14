package dynamicservice_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/test"
	"github.com/openorch/openorch/server/internal/di"
	"github.com/stretchr/testify/require"

	client "github.com/openorch/openorch/clients/go"
	openapi "github.com/openorch/openorch/clients/go"
)

func TestQueryObjects(t *testing.T) {
	uniq := uuid.New().String()
	uniq = strings.Replace(uniq, "-", "", -1)[0:10]

	table1 := "test_table_" + uniq

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)

	err = starterFunc()
	require.NoError(t, err)

	manyClients, err := test.MakeClients(options.ClientFactory, 1)
	require.NoError(t, err)
	client1 := manyClients[0]

	uuids := []string{}

	for i := 0; i < 20; i++ {
		uuid1 := sdk.Id(table1)
		uuids = append(uuids, uuid1)

		obj := openapi.DataSvcCreateObjectFields{
			Id:       &uuid1,
			Table:    table1,
			Readers:  []string{"_self"},
			Writers:  []string{"_self"},
			Deleters: []string{"_self"},
			Data:     map[string]interface{}{"key": i},
		}

		_, _, err = client1.DataSvcAPI.CreateObject(context.Background()).
			Body(client.DataSvcCreateObjectRequest{
				Object: &obj,
			}).
			Execute()
		require.NoError(t, err)
	}

	t.Run("order by desc", func(t *testing.T) {
		req := openapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &openapi.DatastoreQuery{
				OrderBys: []openapi.DatastoreOrderBy{
					{
						Desc:        openapi.PtrBool(true),
						Field:       openapi.PtrString("key"),
						SortingType: openapi.SortingTypeNumeric.Ptr(),
					},
				},
			},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 20, len(rsp.Objects))

		require.Equal(t, float64(19), rsp.Objects[0].Data["key"])
	})

	t.Run("order by asc", func(t *testing.T) {
		req := openapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &openapi.DatastoreQuery{
				OrderBys: []openapi.DatastoreOrderBy{
					{
						Field:       openapi.PtrString("key"),
						SortingType: openapi.SortingTypeNumeric.Ptr(),
					},
				},
			},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 20, len(rsp.Objects))
		require.Equal(t, float64(0), rsp.Objects[0].Data["key"])
	})

	t.Run("limit", func(t *testing.T) {
		req := openapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &openapi.DatastoreQuery{
				OrderBys: []openapi.DatastoreOrderBy{
					{
						Field:       openapi.PtrString("key"),
						SortingType: openapi.SortingTypeNumeric.Ptr(),
					},
				},
				Limit: openapi.PtrInt32(5),
			},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 5, len(rsp.Objects))
		require.Equal(t, float64(0), rsp.Objects[0].Data["key"])
	})

	t.Run("limit, after", func(t *testing.T) {
		req := openapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &openapi.DatastoreQuery{
				OrderBys: []openapi.DatastoreOrderBy{
					{
						Field:       openapi.PtrString("key"),
						SortingType: openapi.SortingTypeNumeric.Ptr(),
					},
				},
				Limit:     openapi.PtrInt32(5),
				JsonAfter: openapi.PtrString("[4]"),
			},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 5, len(rsp.Objects))
		require.Equal(t, float64(5), rsp.Objects[0].Data["key"])
	})

	t.Run("limit, after desc", func(t *testing.T) {
		req := openapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &openapi.DatastoreQuery{
				OrderBys: []openapi.DatastoreOrderBy{
					{
						Field:       openapi.PtrString("key"),
						Desc:        openapi.PtrBool(true),
						SortingType: openapi.SortingTypeNumeric.Ptr(),
					},
				},
				Limit:     openapi.PtrInt32(5),
				JsonAfter: openapi.PtrString("[15]"),
			},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 5, len(rsp.Objects))
		require.Equal(t, float64(14), rsp.Objects[0].Data["key"])
	})
}
