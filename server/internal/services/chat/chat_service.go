/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package chatservice

import (
	"context"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/boot"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/lock"

	chattypes "github.com/1backend/1backend/server/internal/services/chat/types"
)

type ChatService struct {
	clientFactory client.ClientFactory
	token         string

	lock lock.DistributedLock

	messagesStore   datastore.DataStore
	threadsStore    datastore.DataStore
	credentialStore datastore.DataStore
}

func NewChatService(
	clientFactory client.ClientFactory,
	lock lock.DistributedLock,
	datastoreFactory func(tableName string, instance any) (datastore.DataStore, error),
) (*ChatService, error) {
	threadsStore, err := datastoreFactory("chatSvcThreads", &chattypes.Thread{})
	if err != nil {
		return nil, err
	}

	messagesStore, err := datastoreFactory(
		"chatSvcMessages",
		&chattypes.Message{},
	)
	if err != nil {
		return nil, err
	}

	credentialStore, err := datastoreFactory(
		"chatSvcCredentials",
		&auth.Credential{},
	)
	if err != nil {
		return nil, err
	}

	service := &ChatService{
		clientFactory: clientFactory,

		lock:            lock,
		messagesStore:   messagesStore,
		threadsStore:    threadsStore,
		credentialStore: credentialStore,
	}

	return service, nil
}

func (cs *ChatService) Start() error {
	ctx := context.Background()
	cs.lock.Acquire(ctx, "chat-svc-start")
	defer cs.lock.Release(ctx, "chat-svc-start")

	token, err := boot.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"chat-svc",
		"Chat Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token.Token

	return cs.registerPermissions()
}
