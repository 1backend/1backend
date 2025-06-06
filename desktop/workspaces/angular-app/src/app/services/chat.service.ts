/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import { ReplaySubject, Subject } from 'rxjs';
import { FirehoseService } from './firehose.service';
import { UserService } from './user.service';
import { first } from 'rxjs';
import {
	ChatSvcSaveThreadResponse,
	ChatSvcApi,
	ChatSvcListMessagesResponse,
	ChatSvcListThreadsResponse,
	Configuration,
	ChatSvcThread as Thread,
	ChatSvcEventMessageAdded,
	ChatSvcEventThreadAdded,
	ChatSvcEventThreadUpdate,
} from '@1backend/client';

@Injectable({
	providedIn: 'root',
})
export class ChatService {
	private chatService!: ChatSvcApi;

	public activeThreadId: string = '';

	onMessageAddedSubject = new ReplaySubject<ChatSvcEventMessageAdded>(1);
	onMessageAdded$ = this.onMessageAddedSubject.asObservable();

	onThreadAddedSubject = new ReplaySubject<ChatSvcEventThreadAdded>(1);
	onThreadAdded$ = this.onMessageAddedSubject.asObservable();

	onThreadUpdateSubject = new ReplaySubject<ChatSvcEventThreadUpdate>(1);
	onThreadUpdate$ = this.onMessageAddedSubject.asObservable();

	onStartNewThreadSubject = new Subject<void>();
	// emitted when a new thread should be started
	onStartNewThread$ = this.onStartNewThreadSubject.asObservable();

	constructor(
		private server: ServerService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.chatService = new ChatSvcApi(
			new Configuration({
				basePath: this.server.addr(),
				apiKey: this.server.token(),
			})
		);

		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'chatMessageAdded': {
					this.onMessageAddedSubject.next(event.data);
					break;
				}
				case 'chatThreadAdded': {
					this.onMessageAddedSubject.next(event.data);
					break;
				}
				case 'chatThreadUpdate': {
					this.onThreadUpdateSubject.next(event.data);
					break;
				}
			}
		});
	}

	async chatMessageDelete(messageId: string): Promise<object> {
		return this.chatService.deleteMessage({
			messageId: messageId,
		});
	}

	async chatMessages(threadId: string): Promise<ChatSvcListMessagesResponse> {
		return this.chatService.listMessages({
			body: {
				threadId: threadId,
			},
		});
	}

	async chatThread(threadId: string): Promise<ChatSvcListThreadsResponse> {
		return this.chatService.listThreads({
			body: {
				ids: [threadId],
			},
		});
	}

	async chatThreadAdd(thread: Thread): Promise<ChatSvcSaveThreadResponse> {
		return this.chatService.saveThread({
			body: thread,
		});
	}

	async chatThreadUpdate(thread: Thread): Promise<object> {
		return this.chatService.saveThread({
			body: thread,
		});
	}

	async chatThreadDelete(threadId: string): Promise<void> {
		await this.chatService.deleteThread({
			threadId: threadId,
		});
		return;
	}

	async chatThreads(): Promise<ChatSvcListThreadsResponse> {
		return this.chatService.listThreads({
			body: {},
		});
	}

	setActiveThreadId(id: string) {
		localStorage.setItem(this.activeThreadId, id);
	}

	getActiveThreadId(): string {
		const activeThreadId = localStorage.getItem(this.activeThreadId);
		if (!activeThreadId) {
			return '';
		}
		return activeThreadId;
	}
}
