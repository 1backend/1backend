/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package pubsub

import "context"

type Message struct {
	ID      string
	Topic   string
	Payload []byte
}

type Subscription interface {
	Chan() <-chan Message
	Unsubscribe() error
	Ack(ctx context.Context, messageID string) error
	Nack(ctx context.Context, messageID string) error
}

type PubSub interface {
	Publish(ctx context.Context, topic string, payload []byte) (string, error)
	Subscribe(ctx context.Context, subscriberId string, topic string) (Subscription, error)
	Close() error
}
