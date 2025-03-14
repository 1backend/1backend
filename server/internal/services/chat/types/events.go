/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package chat_svc

const EventMessageAddedName = "chatMessageAdded"

type EventMessageAdded struct {
	ThreadId string `json:"threadId"`
}

func (e EventMessageAdded) Name() string {
	return EventMessageAddedName
}

const EventThreadAddedName = "chatThreadAdded"

type EventThreadAdded struct {
	ThreadId string `json:"threadId"`
}

func (e EventThreadAdded) Name() string {
	return EventThreadAddedName
}

const EventThreadUpdateName = "chatThreadUpdate"

type EventThreadUpdate struct {
	ThreadId string `json:"threadId"`
}

func (e EventThreadUpdate) Name() string {
	return EventThreadUpdateName
}
