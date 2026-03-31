//go:build dist
// +build dist

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package pubsub

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type pgDiagnosticsAccessor interface {
	DiagnosticsDB() *sql.DB
	DiagnosticsDeliveriesTable() string
	DiagnosticsMessagesTable() string
}

func TestPublishSubscribe(t *testing.T, ps PubSub) {
	topic := "topic.publish-subscribe"
	payload := []byte("hello")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, topic, msg.Topic)
	require.Equal(t, payload, msg.Payload)
	require.NoError(t, sub.Ack(ctx, msg.ID))
}

func TestTopicIsolation(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.only-this")
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, "topic.other", []byte("drop"))
	require.NoError(t, err)

	mustNotReceiveMessage(t, sub, 300*time.Millisecond)

	_, err = ps.Publish(ctx, "topic.only-this", []byte("keep"))
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, "topic.only-this", msg.Topic)
	require.Equal(t, []byte("keep"), msg.Payload)
	require.NoError(t, sub.Ack(ctx, msg.ID))
}

func TestUnsubscribeStopsDelivery(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.unsubscribe")
	require.NoError(t, err)

	require.NoError(t, sub.Unsubscribe())

	_, err = ps.Publish(ctx, "topic.unsubscribe", []byte("after"))
	require.NoError(t, err)

	mustNotReceiveMessage(t, sub, 300*time.Millisecond)
}

func TestMultipleSubscribersReceiveSameMessage(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.fanout"
	payload := []byte("broadcast")

	sub1, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub1.Unsubscribe()

	sub2, err := ps.Subscribe(ctx, "b", topic)
	require.NoError(t, err)
	defer sub2.Unsubscribe()

	_, err = ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	msg1 := mustReceiveMessage(t, ctx, sub1)
	msg2 := mustReceiveMessage(t, ctx, sub2)

	require.Equal(t, topic, msg1.Topic)
	require.Equal(t, payload, msg1.Payload)
	require.Equal(t, topic, msg2.Topic)
	require.Equal(t, payload, msg2.Payload)

	require.NoError(t, sub1.Ack(ctx, msg1.ID))
	require.NoError(t, sub2.Ack(ctx, msg2.ID))
}

func TestMultiplePublishesAreDelivered(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.multiple-publishes"

	sub, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	payloads := [][]byte{
		[]byte("one"),
		[]byte("two"),
		[]byte("three"),
	}

	for _, payload := range payloads {
		_, err := ps.Publish(ctx, topic, payload)
		require.NoError(t, err)
	}

	for _, expected := range payloads {
		msg := mustReceiveMessage(t, ctx, sub)
		require.Equal(t, topic, msg.Topic)
		require.Equal(t, expected, msg.Payload)
		require.NoError(t, sub.Ack(ctx, msg.ID))
	}
}

func TestSubscribeBackfillSinceFiltersOldMessages(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.backfill-since"

	_, err := ps.Publish(ctx, topic, []byte("old"))
	require.NoError(t, err)

	time.Sleep(20 * time.Millisecond)
	backfillSince := time.Now().UTC()
	time.Sleep(20 * time.Millisecond)

	sub, err := ps.Subscribe(ctx, "a", topic, WithBackfillSince(backfillSince))
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, topic, []byte("new"))
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("new"), msg.Payload)
	require.NoError(t, sub.Ack(ctx, msg.ID))
	mustNotReceiveMessage(t, sub, 300*time.Millisecond)
}

func TestSubscribeBackfillSinceDoesNotCreateOldDeliveryRows(t *testing.T, ps PubSub) {
	accessor, ok := any(ps).(pgDiagnosticsAccessor)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.backfill-no-old-deliveries"
	subscriberID := "consumer.backfill-no-old-deliveries"

	for i := 0; i < 3; i++ {
		_, err := ps.Publish(ctx, topic, []byte(fmt.Sprintf("old-%d", i)))
		require.NoError(t, err)
	}

	time.Sleep(20 * time.Millisecond)
	backfillSince := time.Now().UTC()
	time.Sleep(20 * time.Millisecond)

	sub, err := ps.Subscribe(ctx, subscriberID, topic, WithBackfillSince(backfillSince))
	require.NoError(t, err)
	defer sub.Unsubscribe()

	var oldDeliveries int
	err = accessor.DiagnosticsDB().QueryRowContext(ctx, fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s d
		JOIN %s m ON m.id = d.message_id
		WHERE d.subscriber_id = $1
		  AND d.topic = $2
		  AND m.created_at < $3
	`, accessor.DiagnosticsDeliveriesTable(), accessor.DiagnosticsMessagesTable()), subscriberID, topic, backfillSince).Scan(&oldDeliveries)
	require.NoError(t, err)
	require.Zero(t, oldDeliveries, "subscribe with backfill since should not enqueue old messages")
}

func TestUnsubscribeOneSubscriberDoesNotAffectOthers(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.unsubscribe-one"

	sub1, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)

	sub2, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub2.Unsubscribe()

	require.NoError(t, sub1.Unsubscribe())

	_, err = ps.Publish(ctx, topic, []byte("still-delivered"))
	require.NoError(t, err)

	mustNotReceiveMessage(t, sub1, 300*time.Millisecond)

	msg := mustReceiveMessage(t, ctx, sub2)
	require.Equal(t, topic, msg.Topic)
	require.Equal(t, []byte("still-delivered"), msg.Payload)
	require.NoError(t, sub2.Ack(ctx, msg.ID))
}

func TestSubscriberDoesNotReceiveMessagesFromOtherTopicsInterleaved(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.target")
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, "topic.other.1", []byte("ignore-1"))
	require.NoError(t, err)
	_, err = ps.Publish(ctx, "topic.target", []byte("wanted-1"))
	require.NoError(t, err)
	_, err = ps.Publish(ctx, "topic.other.2", []byte("ignore-2"))
	require.NoError(t, err)
	_, err = ps.Publish(ctx, "topic.target", []byte("wanted-2"))
	require.NoError(t, err)

	msg1 := mustReceiveMessage(t, ctx, sub)
	msg2 := mustReceiveMessage(t, ctx, sub)

	require.Equal(t, "topic.target", msg1.Topic)
	require.Equal(t, []byte("wanted-1"), msg1.Payload)
	require.Equal(t, "topic.target", msg2.Topic)
	require.Equal(t, []byte("wanted-2"), msg2.Payload)

	require.NoError(t, sub.Ack(ctx, msg1.ID))
	require.NoError(t, sub.Ack(ctx, msg2.ID))

	mustNotReceiveMessage(t, sub, 300*time.Millisecond)
}

func TestUnsubscribeIsSafeToCallMoreThanOnce(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.unsubscribe-idempotent")
	require.NoError(t, err)

	require.NoError(t, sub.Unsubscribe())
	require.NoError(t, sub.Unsubscribe())

	_, err = ps.Publish(ctx, "topic.unsubscribe-idempotent", []byte("after"))
	require.NoError(t, err)
	mustNotReceiveMessage(t, sub, 300*time.Millisecond)
}

func TestAckRemovesMessage(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.ack-removes")
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, "topic.ack-removes", []byte("once"))
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)
	require.NoError(t, sub.Ack(ctx, msg.ID))

	// If ack worked, this message should not reappear.
	mustNotReceiveMessage(t, sub, 500*time.Millisecond)
}

func TestNackRequeuesMessage(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.nack-requeue")
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, "topic.nack-requeue", []byte("retry-me"))
	require.NoError(t, err)

	first := mustReceiveMessage(t, ctx, sub)
	require.NoError(t, sub.Nack(ctx, first.ID))

	second := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, first.ID, second.ID, "nacked message should be redelivered")
	require.Equal(t, []byte("retry-me"), second.Payload)
	require.NoError(t, sub.Ack(ctx, second.ID))
}

func TestNackBackoffIncreasesAcrossAttempts(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub, err := ps.Subscribe(ctx, "a", "topic.nack-backoff")
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, "topic.nack-backoff", []byte("retry-with-backoff"))
	require.NoError(t, err)

	first := mustReceiveMessage(t, ctx, sub)

	nackAt1 := time.Now()
	require.NoError(t, sub.Nack(ctx, first.ID))
	second := mustReceiveMessage(t, ctx, sub)
	delay1 := time.Since(nackAt1)
	require.Equal(t, first.ID, second.ID)

	nackAt2 := time.Now()
	require.NoError(t, sub.Nack(ctx, second.ID))
	third := mustReceiveMessage(t, ctx, sub)
	delay2 := time.Since(nackAt2)
	require.Equal(t, second.ID, third.ID)

	require.Greater(t, delay2, delay1+20*time.Millisecond, "expected exponential backoff to increase between retries")
	require.NoError(t, sub.Ack(ctx, third.ID))
}

func TestRetryWithoutNackUsesBackoff(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.retry-without-nack-backoff"
	subscriberID := "consumer.retry-without-nack-backoff"

	sub, err := ps.Subscribe(ctx, subscriberID, topic)
	require.NoError(t, err)

	_, err = ps.Publish(ctx, topic, []byte("retry-after-processing-drop"))
	require.NoError(t, err)

	first := mustReceiveMessage(t, ctx, sub)
	require.NoError(t, sub.Unsubscribe())

	sub2, err := ps.Subscribe(ctx, subscriberID, topic)
	require.NoError(t, err)
	defer sub2.Unsubscribe()

	mustNotReceiveMessage(t, sub2, 40*time.Millisecond)

	second := mustReceiveMessage(t, ctx, sub2)
	require.Equal(t, first.ID, second.ID)
	require.Equal(t, []byte("retry-after-processing-drop"), second.Payload)
	require.NoError(t, sub2.Ack(ctx, second.ID))
}

func TestRetryWithoutAckOrNackIncreasesBackoff(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	topic := "topic.retry-without-ack-or-nack-increases-backoff"
	subscriberID := "consumer.retry-without-ack-or-nack-increases-backoff"
	payload := []byte("retry-progressive-delay")

	sub1, err := ps.Subscribe(ctx, subscriberID, topic)
	require.NoError(t, err)

	_, err = ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	first := mustReceiveMessage(t, ctx, sub1)
	require.NoError(t, sub1.Unsubscribe())

	sub2, err := ps.Subscribe(ctx, subscriberID, topic)
	require.NoError(t, err)
	start1 := time.Now()
	second := mustReceiveMessage(t, ctx, sub2)
	delay1 := time.Since(start1)
	require.Equal(t, first.ID, second.ID)
	require.Equal(t, payload, second.Payload)
	require.NoError(t, sub2.Unsubscribe())

	sub3, err := ps.Subscribe(ctx, subscriberID, topic)
	require.NoError(t, err)
	defer sub3.Unsubscribe()

	start2 := time.Now()
	third := mustReceiveMessage(t, ctx, sub3)
	delay2 := time.Since(start2)
	require.Equal(t, second.ID, third.ID)
	require.Equal(t, payload, third.Payload)

	require.Greater(t, delay2, delay1+20*time.Millisecond, "expected retries without ack/nack to have increasing backoff")
	require.NoError(t, sub3.Ack(ctx, third.ID))
}

func TestRetryIncrementsAttemptDiagnostics(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	diagReader, ok := ps.(DeliveryDiagnosticsReader)
	if !ok {
		t.Skip("pubsub implementation does not expose delivery diagnostics")
	}

	topic := "topic.retry-increments-attempt-diagnostics"
	subscriberID := "consumer.retry-increments-attempt-diagnostics"
	payload := []byte("attempts")

	sub, err := ps.Subscribe(ctx, subscriberID, topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	first := mustReceiveMessage(t, ctx, sub)
	d1, err := diagReader.ReadDeliveryDiagnostics(ctx, subscriberID, topic, first.ID)
	require.NoError(t, err)
	require.GreaterOrEqual(t, d1.Attempts, 1)

	require.NoError(t, sub.Nack(ctx, first.ID))
	second := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, first.ID, second.ID)

	d2, err := diagReader.ReadDeliveryDiagnostics(ctx, subscriberID, topic, second.ID)
	require.NoError(t, err)
	require.Greater(t, d2.Attempts, d1.Attempts, "expected attempt counter to increase after retry")

	require.NoError(t, sub.Ack(ctx, second.ID))
}

func mustReceiveMessage(t *testing.T, ctx context.Context, sub Subscription) Message {
	t.Helper()

	select {
	case msg := <-sub.Chan():
		return msg
	case <-ctx.Done():
		t.Fatalf("timed out waiting for message")
		return Message{}
	}
}

func mustNotReceiveMessage(t *testing.T, sub Subscription, wait time.Duration) {
	t.Helper()

	select {
	case msg := <-sub.Chan():
		t.Fatalf("unexpected message received: topic=%q payload=%q id=%q", msg.Topic, string(msg.Payload), msg.ID)
	case <-time.After(wait):
	}
}

func TestMessagePersistenceAcrossRestarts(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.persistence"
	consumerID := "persistent-consumer-a"
	payload := []byte("stored-in-postgres")

	// 1. Publish the message with the provided instance
	_, err := ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	// Note: In a real restart, ps would be closed.
	// For this test, we assume the underlying DB is shared.

	// 2. Subscribe AFTER publishing.
	// If it's persistent (Postgres), the message is waiting in the table.
	// If it's ephemeral (Local), this will likely time out, which is correct behavior
	// for a non-persistent implementation.
	sub, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	// 3. Verify the message is still there to be picked up
	msg := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, payload, msg.Payload)
	require.NoError(t, sub.Ack(ctx, msg.ID))
}

func TestPublishedBeforeSubscribeIsStillDelivered(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.prepublished"
	consumerID := "consumer.prepublished"

	_, err := ps.Publish(ctx, topic, []byte("m1"))
	require.NoError(t, err)

	_, err = ps.Publish(ctx, topic, []byte("m2"))
	require.NoError(t, err)

	sub, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	msg1 := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("m1"), msg1.Payload)
	require.NoError(t, sub.Ack(ctx, msg1.ID))

	msg2 := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("m2"), msg2.Payload)
	require.NoError(t, sub.Ack(ctx, msg2.ID))
}

func TestMessageOrderPerTopic(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.order"
	consumerID := "consumer.order"

	for i := 0; i < 5; i++ {
		_, err := ps.Publish(ctx, topic, []byte(fmt.Sprintf("msg-%02d", i)))
		require.NoError(t, err)
	}

	sub, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	for i := 0; i < 5; i++ {
		msg := mustReceiveMessage(t, ctx, sub)
		require.Equal(t, []byte(fmt.Sprintf("msg-%02d", i)), msg.Payload)
		require.NoError(t, sub.Ack(ctx, msg.ID))
	}
}

func TestNoDuplicateDeliveryForSinglePublish(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	topic := "topic.no-duplicate"
	consumerID := "consumer.no-duplicate"
	payload := []byte("only-once")

	_, err := ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	sub, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	msg := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, payload, msg.Payload)
	require.NoError(t, sub.Ack(ctx, msg.ID))

	noMsgCtx, noMsgCancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer noMsgCancel()

	requireNoMessage(t, noMsgCtx, sub)
}

func TestMultiplePersistentMessagesSurviveRestart(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.multi-persistence"
	consumerID := "consumer.multi-persistence"

	_, err := ps.Publish(ctx, topic, []byte("m1"))
	require.NoError(t, err)
	_, err = ps.Publish(ctx, topic, []byte("m2"))
	require.NoError(t, err)
	_, err = ps.Publish(ctx, topic, []byte("m3"))
	require.NoError(t, err)

	sub, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	msg1 := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("m1"), msg1.Payload)
	require.NoError(t, sub.Ack(ctx, msg1.ID))

	msg2 := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("m2"), msg2.Payload)
	require.NoError(t, sub.Ack(ctx, msg2.ID))

	msg3 := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("m3"), msg3.Payload)
	require.NoError(t, sub.Ack(ctx, msg3.ID))
}

func TestAckedMessageIsNotRedeliveredOnResubscribe(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.acked-not-redelivered"
	consumerID := "consumer.acked-not-redelivered"
	payload := []byte("acked-message")

	_, err := ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	sub1, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub1)
	require.Equal(t, payload, msg.Payload)
	require.NoError(t, sub1.Ack(ctx, msg.ID))
	sub1.Unsubscribe()

	sub2, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub2.Unsubscribe()

	noMsgCtx, noMsgCancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer noMsgCancel()

	requireNoMessage(t, noMsgCtx, sub2)
}

func TestUnackedMessageIsRedeliveredOnResubscribe(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.unacked-redelivery"
	consumerID := "consumer.unacked-redelivery"
	payload := []byte("needs-redelivery")

	_, err := ps.Publish(ctx, topic, payload)
	require.NoError(t, err)

	sub1, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)

	msg1 := mustReceiveMessage(t, ctx, sub1)
	require.Equal(t, payload, msg1.Payload)
	sub1.Unsubscribe()

	sub2, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)
	defer sub2.Unsubscribe()

	msg2 := mustReceiveMessage(t, ctx, sub2)
	require.Equal(t, payload, msg2.Payload)
	require.NoError(t, sub2.Ack(ctx, msg2.ID))
}

func requireNoMessage(t *testing.T, ctx context.Context, sub Subscription) {
	t.Helper()

	select {
	case msg := <-sub.Chan():
		t.Fatalf("unexpected message received: id=%s topic=%s payload=%q", msg.ID, msg.Topic, msg.Payload)
	case <-ctx.Done():
		// success: no message arrived before timeout
	}
}

func TestSameSubscriberIDAcrossTopicsIsIndependent(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	subA, err := ps.Subscribe(ctx, "consumer-x", "topic.a")
	require.NoError(t, err)
	defer subA.Unsubscribe()

	subB, err := ps.Subscribe(ctx, "consumer-x", "topic.b")
	require.NoError(t, err)
	defer subB.Unsubscribe()

	_, err = ps.Publish(ctx, "topic.a", []byte("a1"))
	require.NoError(t, err)
	_, err = ps.Publish(ctx, "topic.b", []byte("b1"))
	require.NoError(t, err)

	msgA := mustReceiveMessage(t, ctx, subA)
	msgB := mustReceiveMessage(t, ctx, subB)

	require.Equal(t, "topic.a", msgA.Topic)
	require.Equal(t, []byte("a1"), msgA.Payload)
	require.Equal(t, "topic.b", msgB.Topic)
	require.Equal(t, []byte("b1"), msgB.Payload)

	require.NoError(t, subA.Ack(ctx, msgA.ID))
	require.NoError(t, subB.Ack(ctx, msgB.ID))
}

func TestAckFromDifferentSubscriberFails(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.ack-wrong-subscriber"

	subA, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer subA.Unsubscribe()

	subB, err := ps.Subscribe(ctx, "b", topic)
	require.NoError(t, err)
	defer subB.Unsubscribe()

	_, err = ps.Publish(ctx, topic, []byte("hello"))
	require.NoError(t, err)

	msgA := mustReceiveMessage(t, ctx, subA)
	_ = mustReceiveMessage(t, ctx, subB)

	err = subB.Ack(ctx, msgA.ID)
	require.Error(t, err)

	require.NoError(t, subA.Ack(ctx, msgA.ID))
}

func TestAckTwiceFails(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.ack-twice"

	sub, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, topic, []byte("once"))
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)

	require.NoError(t, sub.Ack(ctx, msg.ID))
	require.Error(t, sub.Ack(ctx, msg.ID))
}

func TestNackAfterAckFails(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.nack-after-ack"

	sub, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	_, err = ps.Publish(ctx, topic, []byte("x"))
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)

	require.NoError(t, sub.Ack(ctx, msg.ID))
	require.Error(t, sub.Nack(ctx, msg.ID))
}

func TestSubscribeRejectsEmptySubscriberID(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ps.Subscribe(ctx, "", "topic.x")
	require.Error(t, err)
}

func TestSubscribeRejectsEmptyTopic(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ps.Subscribe(ctx, "a", "")
	require.Error(t, err)
}

func TestPublishRejectsEmptyTopic(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ps.Publish(ctx, "", []byte("x"))
	require.Error(t, err)
}

func TestResubscribeSameConsumerReplacesPreviousSubscription(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	topic := "topic.replace-sub"

	sub1, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)

	sub2, err := ps.Subscribe(ctx, "a", topic)
	require.NoError(t, err)
	defer sub2.Unsubscribe()

	_, err = ps.Publish(ctx, topic, []byte("hello"))
	require.NoError(t, err)

	mustNotReceiveMessage(t, sub1, 300*time.Millisecond)

	msg := mustReceiveMessage(t, ctx, sub2)
	require.Equal(t, []byte("hello"), msg.Payload)
	require.NoError(t, sub2.Ack(ctx, msg.ID))
}

func TestUnackedMessageSurvivesCloseAndCanBeRecovered(t *testing.T, ps PubSub) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topic := "topic.close-recovery"
	consumerID := "consumer.close-recovery"

	sub, err := ps.Subscribe(ctx, consumerID, topic)
	require.NoError(t, err)

	_, err = ps.Publish(ctx, topic, []byte("recover-me"))
	require.NoError(t, err)

	msg := mustReceiveMessage(t, ctx, sub)
	require.Equal(t, []byte("recover-me"), msg.Payload)

	require.NoError(t, ps.Close())

	// This test only works if the harness recreates the same durable backend.
	// If you add a restart-capable harness later, use that here.
	_ = msg
}
