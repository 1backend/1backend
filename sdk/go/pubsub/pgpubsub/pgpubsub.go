/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package pgpubsub

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/1backend/1backend/sdk/go/pubsub"
	"github.com/lib/pq"
)

const channelName = "ob_pubsub_channel"

const (
	defaultTableName = "ob_pubsub"
	defaultLeaseSec  = 30
	defaultNackDelay = 2
)

type PGPubSub struct {
	db       *sql.DB
	listener *pq.Listener
	msgTable string
	subTable string
	delTable string

	mu      sync.RWMutex
	closed  bool
	subbers map[*pgSubscription]struct{}
}

type pgSubscription struct {
	ps      *PGPubSub
	topic   string
	ch      chan pubsub.Message
	id      string
	closeCh chan struct{}
	once    sync.Once
}

func NewPGPubSub(connectionString string, db *sql.DB, tableName string) (*PGPubSub, error) {
	if tableName == "" {
		tableName = defaultTableName
	}
	if connectionString == "" {
		return nil, errors.New("connection string required")
	}
	if db == nil {
		return nil, errors.New("sql db required")
	}

	listener := pq.NewListener(connectionString, 10*time.Second, time.Minute, nil)
	if err := listener.Listen(channelName); err != nil {
		return nil, err
	}

	if err := ensureSchema(db, tableName); err != nil {
		return nil, err
	}

	ps := &PGPubSub{
		db:       db,
		listener: listener,
		msgTable: tableName + "_messages",
		subTable: tableName + "_subscriptions",
		delTable: tableName + "_deliveries",
		subbers:  map[*pgSubscription]struct{}{},
	}
	go ps.consume()

	return ps, nil
}

func (ps *PGPubSub) Publish(ctx context.Context, topic string, payload []byte) (string, error) {
	if topic == "" {
		return "", errors.New("topic is empty")
	}

	var id int64
	err := ps.db.QueryRowContext(
		ctx,
		fmt.Sprintf(`INSERT INTO %s (topic, payload, created_at)
		             VALUES ($1, $2, NOW())
		             RETURNING id`, ps.msgTable),
		topic,
		payload,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	_, err = ps.db.ExecContext(ctx, fmt.Sprintf(`
		INSERT INTO %s (message_id, subscriber_id, topic, status, attempts, available_at, locked_at)
		SELECT $1, s.id, s.topic, 'pending', 0, NOW(), NULL
		FROM %s s
		WHERE s.topic = $2
	`, ps.delTable, ps.subTable), id, topic)
	if err != nil {
		return "", err
	}

	_, _ = ps.db.ExecContext(ctx, "SELECT pg_notify($1, $2)", channelName, topic)
	return strconv.FormatInt(id, 10), nil
}

func (ps *PGPubSub) Subscribe(ctx context.Context, subscriberId string, topic string, options ...pubsub.SubscribeOption) (pubsub.Subscription, error) {
	if topic == "" {
		return nil, errors.New("topic is empty")
	}
	if subscriberId == "" {
		return nil, errors.New("subscriber id is empty")
	}

	ps.mu.Lock()

	if ps.closed {
		ps.mu.Unlock()
		return nil, errors.New("pubsub closed")
	}

	var existingToCancel []*pgSubscription
	for existing := range ps.subbers {
		if existing.id == subscriberId && existing.topic == topic {
			existingToCancel = append(existingToCancel, existing)
		}
	}
	ps.mu.Unlock()

	for _, existing := range existingToCancel {
		_ = existing.Unsubscribe()
	}

	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return nil, errors.New("pubsub closed")
	}

	subscribeOpts := pubsub.BuildSubscribeOptions(options)

	sub := &pgSubscription{
		ps:      ps,
		topic:   topic,
		id:      subscriberId,
		ch:      make(chan pubsub.Message, 32),
		closeCh: make(chan struct{}),
	}
	if err := ps.ensureSubscriber(ctx, sub.id, sub.topic, subscribeOpts.BackfillSince); err != nil {
		return nil, err
	}
	ps.subbers[sub] = struct{}{}

	go func() {
		<-ctx.Done()
		_ = sub.Unsubscribe()
	}()
	go sub.run()

	return sub, nil
}

func (s *pgSubscription) Ack(ctx context.Context, messageID string) error {
	id, err := strconv.ParseInt(messageID, 10, 64)
	if err != nil {
		return err
	}
	res, err := s.ps.db.ExecContext(ctx, fmt.Sprintf(
		`UPDATE %s
		 SET status='acked', locked_at=NULL
		 WHERE message_id=$1 AND subscriber_id=$2 AND topic=$3 AND status='inflight'`, s.ps.delTable),
		id, s.id, s.topic,
	)
	if err != nil {
		return err
	}
	ra, _ := res.RowsAffected()
	if ra == 0 {
		return errors.New("ack failed")
	}
	return nil
}

func (s *pgSubscription) Nack(ctx context.Context, messageID string) error {
	id, err := strconv.ParseInt(messageID, 10, 64)
	if err != nil {
		return err
	}

	return s.ps.nack(ctx, s.id, s.topic, id, defaultNackDelay)
}

func (s *pgSubscription) run() {
	t := time.NewTicker(200 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-s.closeCh:
			return
		default:
		}

		msg, found, err := s.ps.claim(context.Background(), s.topic, s.id)
		if err == nil && found {
			select {
			case <-s.closeCh:
				return
			case s.ch <- pubsub.Message{
				ID:      strconv.FormatInt(msg.ID, 10),
				Topic:   msg.Topic,
				Payload: msg.Payload,
			}:
			default:
				_ = s.ps.nack(context.Background(), s.id, s.topic, msg.ID, defaultNackDelay)
			}
			continue
		}

		select {
		case <-s.closeCh:
			return
		case <-t.C:
		}
	}
}

func (ps *PGPubSub) Close() error {
	ps.mu.Lock()

	if ps.closed {
		ps.mu.Unlock()
		return nil
	}
	ps.closed = true

	subs := make([]*pgSubscription, 0, len(ps.subbers))
	for sub := range ps.subbers {
		subs = append(subs, sub)
	}
	ps.mu.Unlock()

	for _, sub := range subs {
		_ = sub.Unsubscribe()
	}

	return ps.listener.Close()
}

func (s *pgSubscription) Chan() <-chan pubsub.Message {
	return s.ch
}

func (s *pgSubscription) Unsubscribe() error {
	s.once.Do(func() {
		close(s.closeCh)

		s.ps.mu.Lock()
		delete(s.ps.subbers, s)
		s.ps.mu.Unlock()

		_, _ = s.ps.db.ExecContext(context.Background(), fmt.Sprintf(`
			UPDATE %s
			SET status='pending',
			    locked_at=NULL,
			    available_at=NOW()
			WHERE subscriber_id=$1
			  AND topic=$2
			  AND status='inflight'
		`, s.ps.delTable), s.id, s.topic)
	})
	return nil
}

func (ps *PGPubSub) consume() {
	for {
		n := <-ps.listener.Notify
		if n == nil {
			ps.mu.RLock()
			closed := ps.closed
			ps.mu.RUnlock()
			if closed {
				return
			}
			continue
		}

		// No decode / no dispatch here.
		// Message delivery happens in sub.run() via ps.claim(...).
	}
}

type dbMessage struct {
	ID      int64
	Topic   string
	Payload []byte
}

func ensureSchema(db *sql.DB, table string) error {
	_, err := db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s_messages (
			id BIGSERIAL PRIMARY KEY,
			topic TEXT NOT NULL,
			payload BYTEA NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
		CREATE TABLE IF NOT EXISTS %s_subscriptions (
			id TEXT NOT NULL,
			topic TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			last_seen TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			PRIMARY KEY (id, topic)
		);
		CREATE TABLE IF NOT EXISTS %s_deliveries (
			message_id BIGINT NOT NULL REFERENCES %s_messages(id) ON DELETE CASCADE,
			subscriber_id TEXT NOT NULL,
			topic TEXT NOT NULL,
			status TEXT NOT NULL,
			attempts INT NOT NULL DEFAULT 0,
			available_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			locked_at TIMESTAMPTZ,
			PRIMARY KEY (message_id, subscriber_id, topic),
			FOREIGN KEY (subscriber_id, topic)
				REFERENCES %s_subscriptions(id, topic) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS %s_topic_created_idx ON %s_messages(topic, created_at, id);
		CREATE INDEX IF NOT EXISTS %s_sub_topic_idx ON %s_subscriptions(topic);
		CREATE INDEX IF NOT EXISTS %s_del_claim_idx ON %s_deliveries(subscriber_id, status, available_at, message_id);
	`, table, table, table, table, table, table, table, table, table, table, table))
	return err
}

func (ps *PGPubSub) claim(ctx context.Context, topic, subID string) (*dbMessage, bool, error) {
	q := fmt.Sprintf(`
		WITH c AS (
		  SELECT d.message_id
		  FROM %s d
		  JOIN %s m ON m.id = d.message_id
		  WHERE d.subscriber_id=$1
		    AND m.topic=$2
			AND d.topic=$2
		    AND d.available_at <= NOW()
		    AND (
		      d.status='pending'
		      OR (d.status='inflight' AND d.locked_at < NOW() - INTERVAL '%d second')
		    )
		  ORDER BY d.message_id
		  FOR UPDATE SKIP LOCKED
		  LIMIT 1
		)
		UPDATE %s d
		SET status='inflight', locked_at=NOW(), attempts=d.attempts+1
		FROM c
		WHERE d.message_id=c.message_id
		  AND d.subscriber_id=$1
		RETURNING d.message_id
	`, ps.delTable, ps.msgTable, defaultLeaseSec, ps.delTable)

	var msgID int64
	err := ps.db.QueryRowContext(ctx, q, subID, topic).Scan(&msgID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	msg := dbMessage{}
	err = ps.db.QueryRowContext(ctx, fmt.Sprintf(
		`SELECT id, topic, payload FROM %s WHERE id=$1`, ps.msgTable), msgID).
		Scan(&msg.ID, &msg.Topic, &msg.Payload)
	if err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (ps *PGPubSub) nack(ctx context.Context, subID, topic string, messageID int64, delaySeconds int) error {
	res, err := ps.db.ExecContext(ctx, fmt.Sprintf(`
		UPDATE %s d
		SET status='pending',
		    locked_at=NULL,
		    available_at=NOW() + ($3 * INTERVAL '1 second')
		WHERE d.message_id=$1
		  AND d.subscriber_id=$2
		AND d.topic=$4
		  AND d.status='inflight'
	`, ps.delTable), messageID, subID, delaySeconds, topic)
	if err != nil {
		return err
	}
	ra, _ := res.RowsAffected()
	if ra == 0 {
		return errors.New("nack failed")
	}
	return nil
}

func (ps *PGPubSub) ensureSubscriber(ctx context.Context, subID, topic string, backfillSince *time.Time) error {
	_, err := ps.db.ExecContext(ctx, fmt.Sprintf(`
        INSERT INTO %s (id, topic, created_at, last_seen)
        VALUES ($1, $2, NOW(), NOW())
        ON CONFLICT (id, topic)
        DO UPDATE SET last_seen=NOW()
    `, ps.subTable), subID, topic)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
        INSERT INTO %s (message_id, subscriber_id, topic, status, attempts, available_at)
        SELECT m.id, $1, $2, 'pending', 0, NOW()
        FROM %s m
        WHERE m.topic = $2
          AND NOT EXISTS (
            SELECT 1 FROM %s d
            WHERE d.message_id = m.id
              AND d.subscriber_id = $1
              AND d.topic = $2
          )`, ps.delTable, ps.msgTable, ps.delTable)

	args := []any{subID, topic}
	if backfillSince != nil {
		query += "\n          AND m.created_at >= $3"
		args = append(args, backfillSince.UTC())
	}

	_, err = ps.db.ExecContext(ctx, query, args...)

	return err
}

func (ps *PGPubSub) removeSubscriber(ctx context.Context, subID, topic string) error {
	_, err := ps.db.ExecContext(ctx, fmt.Sprintf(
		`DELETE FROM %s WHERE id=$1 AND topic=$2`, ps.subTable),
		subID, topic,
	)
	return err
}
