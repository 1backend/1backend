package dynamicservice_test

import (
	"context"
	"strings"
	"testing"

	sdk "github.com/1backend/1backend/sdk/go"
	sdkclient "github.com/1backend/1backend/sdk/go/client"

	"github.com/1backend/1backend/sdk/go/test"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	obapi "github.com/1backend/1backend/clients/go"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	server, err := test.StartService(test.Options{
		Test: true,
	})
	require.NoError(t, err)
	defer server.Cleanup(t)

	clientFactory := sdkclient.NewApiClientFactory(server.Url)

	uniq := uuid.New().String()
	uniq = strings.Replace(uniq, "-", "", -1)[0:10]

	table1 := "test_table" + uniq
	table2 := "test_table2" + uniq

	manyClients, _, err := test.MakeClients(clientFactory, 2)
	require.NoError(t, err)
	client1 := manyClients[0]
	client2 := manyClients[1]

	tokenReadRsp1, _, err := client1.UserSvcAPI.ReadSelf(context.Background()).
		Execute()
	require.NoError(t, err)

	tokenReadRsp2, _, err := client2.UserSvcAPI.ReadSelf(context.Background()).
		Execute()
	require.NoError(t, err)

	uuid1 := sdk.Id(table1)
	uuid2 := sdk.Id(table2)

	obj := obapi.DataSvcCreateObjectFields{
		Id:       &uuid1,
		Table:    table1,
		Readers:  []string{"_self"},
		Writers:  []string{"_self"},
		Deleters: []string{"_self"},
		Data:     map[string]interface{}{"key": "value"},
	}

	_, _, err = client1.DataSvcAPI.CreateObject(context.Background()).
		Body(obapi.DataSvcCreateObjectRequest{
			Object: &obj,
		}).
		Execute()
	require.NoError(t, err)

	t.Run("user 1 can find its own private record", func(t *testing.T) {
		req := obapi.DataSvcQueryRequest{
			Table:   &table1,
			Readers: []string{tokenReadRsp1.User.Id},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	obj2 := obapi.DataSvcCreateObjectFields{
		Id:      &uuid2,
		Table:   table2,
		Readers: []string{tokenReadRsp2.User.Id},
		Data:    map[string]interface{}{"key": "value"},
	}

	_, _, err = client2.DataSvcAPI.CreateObject(context.Background()).
		Body(obapi.DataSvcCreateObjectRequest{
			Object: &obj2,
		}).
		Execute()
	require.NoError(t, err)

	t.Run("query user2 records", func(t *testing.T) {
		req := obapi.DataSvcQueryRequest{
			Table:   &table2,
			Readers: []string{tokenReadRsp2.User.Id},
		}

		rsp, _, err := client2.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid2, *rsp.Objects[0].Id)
	})

	t.Run("query user1 records", func(t *testing.T) {
		req := obapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &obapi.DatastoreQuery{Filters: []obapi.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         obapi.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{uuid1}),
				},
			}},

			Readers: []string{tokenReadRsp1.User.Id},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	t.Run("query user1 records with _self", func(t *testing.T) {
		req := obapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &obapi.DatastoreQuery{Filters: []obapi.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         obapi.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{uuid1}),
				},
			}},

			Readers: []string{"_self"},
		}

		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	t.Run("already exists", func(t *testing.T) {
		_, _, err = client1.DataSvcAPI.CreateObject(context.Background()).
			Body(obapi.DataSvcCreateObjectRequest{
				Object: &obj,
			}).
			Execute()

		require.Error(t, err)
	})

	t.Run("user 1 cannot see record of user 2", func(t *testing.T) {
		req := obapi.DataSvcQueryRequest{
			Table: &table1,
			Query: &obapi.DatastoreQuery{Filters: []obapi.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         obapi.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{uuid2}),
				},
			}},
			Readers: []string{tokenReadRsp2.User.Id},
		}
		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(req).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})

	t.Run("user 2 cannot update record of user 1", func(t *testing.T) {
		req := &obapi.DataSvcUpsertObjectRequest{
			Object: &obj,
		}
		_, _, err = client2.DataSvcAPI.UpsertObject(context.Background(), *obj.Id).
			Body(*req).
			Execute()

		// unauthorized
		require.Error(t, err)
	})

	t.Run("user 1 can upsert its own record", func(t *testing.T) {
		req := &obapi.DataSvcUpsertObjectRequest{
			Object: &obj,
		}
		_, _, err = client1.DataSvcAPI.UpsertObject(context.Background(), *obj.Id).
			Body(*req).
			Execute()

		require.NoError(t, err)
	})

	t.Run("user 1 can find its own record", func(t *testing.T) {
		req := &obapi.DataSvcQueryRequest{
			Table:   obapi.PtrString(table1),
			Readers: []string{tokenReadRsp1.User.Id},
		}
		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(*req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, uuid1, *rsp.Objects[0].Id)
	})

	t.Run("user 2 cannot delete user 1's record", func(t *testing.T) {
		req := &obapi.DataSvcDeleteObjectRequest{
			Table: obapi.PtrString(table1),
			Filters: []obapi.DatastoreFilter{
				{
					Fields:     []string{"id"},
					Op:         obapi.OpEquals.Ptr(),
					JsonValues: sdk.Marshal([]any{obj.Id}),
				},
			},
		}

		_, _, err = client2.DataSvcAPI.DeleteObjects(context.Background()).
			Body(*req).
			Execute()

		require.NoError(t, err)

		// Check if user 1 can still find it
		listReq := &obapi.DataSvcQueryRequest{
			Table:   obapi.PtrString(table1),
			Readers: []string{tokenReadRsp1.User.Id},
		}
		rsp, _, err := client1.DataSvcAPI.QueryObjects(context.Background()).
			Body(*listReq).
			Execute()
		require.NoError(t, err)
		require.Equal(t, 1, len(rsp.Objects))
		require.Equal(t, *obj.Id, *rsp.Objects[0].Id)
	})

	// ...item wont be deleted
	t.Run("user 2 will no see other tables", func(t *testing.T) {
		req := &obapi.DataSvcQueryRequest{
			Table: obapi.PtrString(table1),
		}
		rsp, _, err := client2.DataSvcAPI.QueryObjects(context.Background()).
			Body(*req).
			Execute()

		require.NoError(t, err)
		require.Equal(t, 0, len(rsp.Objects))
	})
}
