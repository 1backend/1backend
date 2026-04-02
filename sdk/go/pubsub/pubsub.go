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
	"time"
)

type Message struct {
	// ID is the backend-generated identifier for this message.
	ID string
	// Topic is the logical channel this message was published to.
	Topic string
	// Payload contains the opaque message bytes supplied by the publisher.
	Payload []byte
}

// DeliveryDiagnostics captures backend-level delivery state for a single
// subscriber+message pair. This is intended for tests and troubleshooting, not
// for production message processing flows.
type DeliveryDiagnostics struct {
	// Attempts is the number of delivery attempts recorded for this message.
	Attempts int
}

// DeliveryDiagnosticsReader is an optional capability for PubSub implementations
// that can expose internal delivery bookkeeping for diagnostics/testing.
//
// Consumers should always access this via type assertion and handle
// non-supporting implementations gracefully.
type DeliveryDiagnosticsReader interface {
	// ReadDeliveryDiagnostics returns internal per-subscriber delivery state for a
	// message, when supported by the backend.
	ReadDeliveryDiagnostics(ctx context.Context, subscriberID string, topic string, messageID string) (DeliveryDiagnostics, error)
}

// Subscription represents an active subscriber binding for a topic.
type Subscription interface {
	// Chan returns the stream of messages for this subscription.
	Chan() <-chan Message
	// Unsubscribe terminates the subscription and releases associated resources.
	Unsubscribe() error
	// Ack marks a message as successfully handled so it will not be redelivered.
	Ack(ctx context.Context, messageID string) error
	// Nack marks a message as not successfully handled and requests redelivery.
	// Options can be used to provide optional diagnostics such as a freeform
	// nack message for backend troubleshooting.
	// Compared to leaving a message inflight, Nack allows the backend to release
	// it immediately and schedule retry policy without waiting for lease timeout.
	Nack(ctx context.Context, messageID string, options ...NackOption) error
}

// NackOptions controls optional behavior for nack operations.
type NackOptions struct {
	// Message is an optional freeform diagnostic reason for the nack.
	Message string
}

// NackOption mutates NackOptions.
type NackOption func(*NackOptions)

// WithNackMessage sets a freeform diagnostic message for a nack.
func WithNackMessage(message string) NackOption {
	return func(opts *NackOptions) {
		opts.Message = message
	}
}

// BuildNackOptions materializes the final options from functional options.
func BuildNackOptions(options []NackOption) NackOptions {
	opts := NackOptions{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opts)
	}
	return opts
}

// SubscribeOptions controls optional behavior when creating a subscription.
type SubscribeOptions struct {
	// BackfillSince requests replay of messages published at or after this time.
	BackfillSince *time.Time
}

// SubscribeOption mutates SubscribeOptions for subscription creation.
type SubscribeOption func(*SubscribeOptions)

// WithBackfillSince configures subscription backfill to begin at the given time.
func WithBackfillSince(since time.Time) SubscribeOption {
	s := since.UTC()
	return func(opts *SubscribeOptions) {
		opts.BackfillSince = &s
	}
}

// BuildSubscribeOptions materializes the final options from functional options.
func BuildSubscribeOptions(options []SubscribeOption) SubscribeOptions {
	opts := SubscribeOptions{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option(&opts)
	}
	return opts
}

// PubSub defines publish/subscribe operations supported by a backend.
type PubSub interface {
	// Publish writes payload to topic and returns the created message ID.
	Publish(ctx context.Context, topic string, payload []byte) (string, error)
	// Subscribe creates or resumes a subscriber stream for a topic.
	Subscribe(ctx context.Context, subscriberId string, topic string, options ...SubscribeOption) (Subscription, error)
	// Close releases backend resources held by this PubSub instance.
	Close() error
}
