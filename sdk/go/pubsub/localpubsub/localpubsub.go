/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package localpubsub

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/pubsub"
)

type logEntry struct {
	ID        string    `json:"id"`
	Topic     string    `json:"topic"`
	Payload   string    `json:"payload"`
	Published time.Time `json:"published"`
}

type LocalPubSub struct {
	filePath string
	ackPath  string
	f        *os.File

	mu      sync.Mutex
	closed  bool
	subbers map[*localSubscription]struct{}

	acked map[string]map[string]bool
	meta  map[string]map[string]deliveryMeta
}

type ackState struct {
	// subscriberId -> messageId -> acked
	Acked map[string]map[string]bool         `json:"acked"`
	Meta  map[string]map[string]deliveryMeta `json:"meta"`
}

type deliveryMeta struct {
	Attempts    int       `json:"attempts"`
	AvailableAt time.Time `json:"available_at"`
}

type localSubscription struct {
	ps       *LocalPubSub
	topic    string
	id       string
	backfill *time.Time
	ch       chan pubsub.Message
	offset   int64
	closeCh  chan struct{}
	once     sync.Once
	inflight map[string]struct{}
	messages map[string]pubsub.Message
	order    []string
}

func NewLocalPubSub(filePath string) (*LocalPubSub, error) {
	if filePath == "" {
		filePath = filepath.Join(os.TempDir(), "1backend-pubsub-"+sdk.Id("")+".log")
	}

	ackPath := filePath + ".acks.json"
	acked := map[string]map[string]bool{}
	meta := map[string]map[string]deliveryMeta{}
	if b, err := os.ReadFile(ackPath); err == nil && len(b) > 0 {
		st := ackState{}
		if json.Unmarshal(b, &st) == nil && st.Acked != nil {
			acked = st.Acked
		}
		if st.Meta != nil {
			meta = st.Meta
		}
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return &LocalPubSub{
		filePath: filePath,
		ackPath:  ackPath,
		acked:    acked,
		meta:     meta,
		f:        f,
		subbers:  map[*localSubscription]struct{}{},
	}, nil
}

func (ps *LocalPubSub) Publish(ctx context.Context, topic string, payload []byte) (string, error) {
	if topic == "" {
		return "", errors.New("topic is empty")
	}

	messageID := strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + sdk.Id("msg")
	entry := logEntry{
		ID:        messageID,
		Topic:     topic,
		Payload:   base64.StdEncoding.EncodeToString(payload),
		Published: time.Now().UTC(),
	}
	line, err := json.Marshal(entry)
	if err != nil {
		return "", err
	}
	line = append(line, '\n')

	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return "", errors.New("pubsub closed")
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	_, err = ps.f.Write(line)
	if err != nil {
		return "", err
	}

	return messageID, ps.f.Sync()
}

func (ps *LocalPubSub) saveAckStateLocked() error {
	b, err := json.Marshal(ackState{Acked: ps.acked, Meta: ps.meta})
	if err != nil {
		return err
	}
	return os.WriteFile(ps.ackPath, b, 0644)
}

func (ps *LocalPubSub) Subscribe(ctx context.Context, subscriberId, topic string, options ...pubsub.SubscribeOption) (pubsub.Subscription, error) {
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

	var existingToCancel []*localSubscription
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

	sub := &localSubscription{
		ps:       ps,
		topic:    topic,
		id:       subscriberId,
		backfill: subscribeOpts.BackfillSince,
		ch:       make(chan pubsub.Message, 32),
		offset:   0,
		closeCh:  make(chan struct{}),
		inflight: map[string]struct{}{},
		messages: map[string]pubsub.Message{},
		order:    []string{},
	}
	ps.subbers[sub] = struct{}{}

	go sub.run(ctx)

	return sub, nil
}

func (s *localSubscription) Ack(ctx context.Context, messageID string) error {
	if messageID == "" {
		return errors.New("message id is empty")
	}
	s.ps.mu.Lock()
	defer s.ps.mu.Unlock()

	if _, ok := s.inflight[messageID]; !ok {
		return errors.New("message is not inflight for this subscription")
	}

	if s.ps.acked[s.id] == nil {
		s.ps.acked[s.id] = map[string]bool{}
	}
	s.ps.acked[s.id][messageID] = true
	if s.ps.meta[s.id] != nil {
		delete(s.ps.meta[s.id], messageID)
	}
	delete(s.inflight, messageID)
	delete(s.messages, messageID)

	return s.ps.saveAckStateLocked()
}

func (s *localSubscription) Nack(ctx context.Context, messageID string, options ...pubsub.NackOption) error {
	_ = pubsub.BuildNackOptions(options)
	if messageID == "" {
		return errors.New("message id is empty")
	}
	s.ps.mu.Lock()
	defer s.ps.mu.Unlock()

	if _, ok := s.inflight[messageID]; !ok {
		return errors.New("message is not inflight for this subscription")
	}

	delete(s.inflight, messageID)
	s.ps.setMessageBackoffLocked(s.id, messageID, time.Now().UTC())

	return s.ps.saveAckStateLocked()
}

func (ps *LocalPubSub) ReadDeliveryDiagnostics(ctx context.Context, subscriberID, topic, messageID string) (pubsub.DeliveryDiagnostics, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	_ = ctx
	_ = topic // Local backend stores diagnostics per subscriber/message; topic is included for API parity.

	meta := ps.meta[subscriberID][messageID]
	return pubsub.DeliveryDiagnostics{
		Attempts: meta.Attempts,
	}, nil
}

func (s *localSubscription) isAcked(messageID string) bool {
	s.ps.mu.Lock()
	defer s.ps.mu.Unlock()
	if s.ps.acked[s.id] == nil {
		return false
	}
	return s.ps.acked[s.id][messageID]
}

func nackBackoffDuration(attempts int) time.Duration {
	if attempts < 1 {
		attempts = 1
	}
	backoffMillis := 100 * (1 << (attempts - 1))
	if backoffMillis > 5000 {
		backoffMillis = 5000
	}
	return time.Duration(backoffMillis) * time.Millisecond
}

func (ps *LocalPubSub) setMessageBackoffLocked(subscriberID, messageID string, now time.Time) {
	if ps.meta[subscriberID] == nil {
		ps.meta[subscriberID] = map[string]deliveryMeta{}
	}
	meta := ps.meta[subscriberID][messageID]
	if meta.Attempts < 1 {
		meta.Attempts = 1
	}
	meta.AvailableAt = now.Add(nackBackoffDuration(meta.Attempts))
	ps.meta[subscriberID][messageID] = meta
}

func (ps *LocalPubSub) Close() error {
	ps.mu.Lock()

	if ps.closed {
		ps.mu.Unlock()
		return nil
	}

	ps.closed = true
	subs := make([]*localSubscription, 0, len(ps.subbers))
	for sub := range ps.subbers {
		subs = append(subs, sub)
	}
	ps.mu.Unlock()

	for _, sub := range subs {
		_ = sub.Unsubscribe()
	}

	return ps.f.Close()
}

func (s *localSubscription) Chan() <-chan pubsub.Message {
	return s.ch
}

func (s *localSubscription) Unsubscribe() error {
	s.once.Do(func() {
		close(s.closeCh)

		s.ps.mu.Lock()
		delete(s.ps.subbers, s)
		s.ps.mu.Unlock()
	})
	return nil
}

func (s *localSubscription) run(ctx context.Context) {
	ticker := time.NewTicker(25 * time.Millisecond)
	defer ticker.Stop()

	for {
		if err := s.forwardNewEntries(); err != nil {
			_ = s.Unsubscribe()
			return
		}
		if err := s.deliverAvailable(); err != nil {
			_ = s.Unsubscribe()
			return
		}

		select {
		case <-ctx.Done():
			_ = s.Unsubscribe()
			return
		case <-s.closeCh:
			return
		case <-ticker.C:
		}
	}
}

func (s *localSubscription) forwardNewEntries() error {
	f, err := os.Open(s.ps.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Seek(s.offset, io.SeekStart); err != nil {
		return err
	}

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				offset, seekErr := f.Seek(0, io.SeekCurrent)
				if seekErr != nil {
					return seekErr
				}
				s.offset = offset
				return nil
			}
			return err
		}

		entry := logEntry{}
		if err := json.Unmarshal(line, &entry); err != nil {
			continue
		}
		if entry.Topic != s.topic {
			continue
		}
		if entry.ID == "" {
			continue
		}
		if s.backfill != nil && entry.Published.Before(*s.backfill) {
			continue
		}
		payload, err := base64.StdEncoding.DecodeString(entry.Payload)
		if err != nil {
			continue
		}

		select {
		case <-s.closeCh:
			return nil
		default:
		}

		s.ps.mu.Lock()
		if s.ps.acked[s.id] != nil && s.ps.acked[s.id][entry.ID] {
			s.ps.mu.Unlock()
			continue
		}
		if _, exists := s.messages[entry.ID]; !exists {
			s.messages[entry.ID] = pubsub.Message{
				ID:      entry.ID,
				Topic:   entry.Topic,
				Payload: payload,
			}
			s.order = append(s.order, entry.ID)
		}
		s.ps.mu.Unlock()
	}
}

func (s *localSubscription) deliverAvailable() error {
	for {
		var message pubsub.Message
		found := false
		now := time.Now().UTC()

		s.ps.mu.Lock()
		for _, messageID := range s.order {
			msg, exists := s.messages[messageID]
			if !exists {
				continue
			}
			if s.ps.acked[s.id] != nil && s.ps.acked[s.id][messageID] {
				delete(s.messages, messageID)
				continue
			}
			if _, inflight := s.inflight[messageID]; inflight {
				continue
			}
			meta := s.ps.meta[s.id][messageID]
			if !meta.AvailableAt.IsZero() && now.Before(meta.AvailableAt) {
				continue
			}
			if s.ps.meta[s.id] == nil {
				s.ps.meta[s.id] = map[string]deliveryMeta{}
			}
			meta.Attempts++
			meta.AvailableAt = now.Add(nackBackoffDuration(meta.Attempts))
			s.ps.meta[s.id][messageID] = meta
			s.inflight[messageID] = struct{}{}
			if err := s.ps.saveAckStateLocked(); err != nil {
				s.ps.mu.Unlock()
				return err
			}
			message = msg
			found = true
			break
		}
		s.ps.mu.Unlock()

		if !found {
			return nil
		}

		select {
		case s.ch <- message:
		default:
			s.ps.mu.Lock()
			delete(s.inflight, message.ID)
			s.ps.mu.Unlock()
			return nil
		}
	}
}
