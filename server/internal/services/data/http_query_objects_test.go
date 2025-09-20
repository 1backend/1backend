package dynamicservice_test

import (
	"context"
	"strings"
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	openapi "github.com/1backend/1backend/clients/go"
)

func TestQueryObjects(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := client.NewApiClientFactory(server.Url)

	uniq := uuid.New().String()
	uniq = strings.Replace(uniq, "-", "", -1)[0:10]

	table1 := "test_table_" + uniq

	manyClients, _, err := test.MakeClients(clientFactory, "test", 1)
	require.NoError(t, err)
	client1 := manyClients[0]

	for i := 0; i < 20; i++ {
		uuid1 := sdk.Id(table1)

		obj := openapi.DataSvcCreateObjectFields{
			Id:       &uuid1,
			Table:    table1,
			Readers:  []string{"_self"},
			Writers:  []string{"_self"},
			Deleters: []string{"_self"},
			Data:     map[string]interface{}{"key": i},
		}

		_, _, err = client1.DataSvcAPI.CreateObject(context.Background()).
			Body(openapi.DataSvcCreateObjectRequest{
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
				AfterJson: sdk.Marshal([]any{4}),
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
				AfterJson: sdk.Marshal([]any{15}),
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
