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

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	"github.com/openorch/openorch/sdk/go/lock"

	chattypes "github.com/openorch/openorch/server/internal/services/chat/types"
)

type ChatService struct {
	clientFactory sdk.ClientFactory
	token         string

	lock lock.DistributedLock

	messagesStore   datastore.DataStore
	threadsStore    datastore.DataStore
	credentialStore datastore.DataStore
}

func NewChatService(
	clientFactory sdk.ClientFactory,
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
		&sdk.Credential{},
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

	token, err := sdk.RegisterServiceAccount(
		cs.clientFactory.Client().UserSvcAPI,
		"chat-svc",
		"Chat Svc",
		cs.credentialStore,
	)
	if err != nil {
		return err
	}
	cs.token = token

	return cs.registerPermissions()
}
