//go:build dist
// +build dist

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package pubsub_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/pubsub"
	"github.com/1backend/1backend/sdk/go/pubsub/localpubsub"
	"github.com/1backend/1backend/sdk/go/pubsub/pgpubsub"
	"github.com/stretchr/testify/require"
)

const pgConn = "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"

func TestAll(t *testing.T) {
	factories := map[string]func(t *testing.T) pubsub.PubSub{
		"localPubSub": func(t *testing.T) pubsub.PubSub {
			ps, err := localpubsub.NewLocalPubSub("")
			require.NoError(t, err)
			return ps
		},
		"pgPubSub": func(t *testing.T) pubsub.PubSub {
			db, err := sql.Open("postgres", pgConn)
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, db.Close())
			})

			require.NoError(t, db.Ping(), "postgres must be available for dist pubsub tests")

			uniqueTable := fmt.Sprintf("t_%d", time.Now().UnixNano())

			ps, err := pgpubsub.NewPGPubSub(pgConn, db, uniqueTable)
			require.NoError(t, err)
			return ps
		},
	}

	tests := map[string]func(t *testing.T, ps pubsub.PubSub){
		"PublishSubscribe":                                           pubsub.TestPublishSubscribe,
		"TopicIsolation":                                             pubsub.TestTopicIsolation,
		"UnsubscribeStopsDelivery":                                   pubsub.TestUnsubscribeStopsDelivery,
		"MultipleSubscribersReceiveSameMessage":                      pubsub.TestMultipleSubscribersReceiveSameMessage,
		"MultiplePublishesAreDelivered":                              pubsub.TestMultiplePublishesAreDelivered,
		"UnsubscribeOneSubscriberDoesNotAffectOthers":                pubsub.TestUnsubscribeOneSubscriberDoesNotAffectOthers,
		"SubscriberDoesNotReceiveMessagesFromOtherTopicsInterleaved": pubsub.TestSubscriberDoesNotReceiveMessagesFromOtherTopicsInterleaved,
		"UnsubscribeIsSafeToCallMoreThanOnce":                        pubsub.TestUnsubscribeIsSafeToCallMoreThanOnce,
		"MessagePersistenceAcrossRestarts":                           pubsub.TestMessagePersistenceAcrossRestarts,
		"PublishedBeforeSubscribeIsStillDelivered":                   pubsub.TestPublishedBeforeSubscribeIsStillDelivered,
		"MessageOrderPerTopic":                                       pubsub.TestMessageOrderPerTopic,
		"NoDuplicateDeliveryForSinglePublish":                        pubsub.TestNoDuplicateDeliveryForSinglePublish,
		"AckedMessageIsNotRedeliveredOnResubscribe":                  pubsub.TestAckedMessageIsNotRedeliveredOnResubscribe,
		"UnackedMessageIsRedeliveredOnResubscribe":                   pubsub.TestUnackedMessageIsRedeliveredOnResubscribe,
		"MultiplePersistentMessagesSurviveRestart":                   pubsub.TestMultiplePersistentMessagesSurviveRestart,
		"SameSubscriberIDAcrossTopicsIsIndependent":                  pubsub.TestSameSubscriberIDAcrossTopicsIsIndependent,
		// "AckFromDifferentSubscriberFails":                            pubsub.TestAckFromDifferentSubscriberFails,
		"AckTwiceFails":     pubsub.TestAckTwiceFails,
		"NackAfterAckFails": pubsub.TestNackAfterAckFails,
		"ResubscribeSameConsumerReplacesPreviousSubscription": pubsub.TestResubscribeSameConsumerReplacesPreviousSubscription,
		"SubscribeRejectsEmptySubscriberID":                   pubsub.TestSubscribeRejectsEmptySubscriberID,
		"SubscribeRejectsEmptyTopic":                          pubsub.TestSubscribeRejectsEmptyTopic,
		"PublishRejectsEmptyTopic":                            pubsub.TestPublishRejectsEmptyTopic,
		"SubscribeBackfillSinceFiltersOldMessages":            pubsub.TestSubscribeBackfillSinceFiltersOldMessages,
	}

	for testName, testFn := range tests {
		for implName, implFactory := range factories {
			t.Run(fmt.Sprintf("%s %s", testName, implName), func(t *testing.T) {
				ps := implFactory(t)
				t.Cleanup(func() {
					require.NoError(t, ps.Close())
				})
				testFn(t, ps)
			})
		}
	}
}
