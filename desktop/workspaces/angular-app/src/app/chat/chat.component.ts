/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	Component,
	OnInit,
} from '@angular/core';
import { ServerService } from '../services/server.service';

import { ElectronIpcService } from '../services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { Subscription } from 'rxjs';
import { ChatService } from '../services/chat.service';
import { PromptService } from '../services/prompt.service';
import { PromptSvcPrompt as Prompt } from '@openorch/client';
import { ModelService } from '../services/model.service';
import {
	ChatSvcThread as Thread,
	ModelSvcModel as Model,
} from '@openorch/client';
import { ConfigService } from '../services/config.service';
import { ChatBoxComponent } from './chat-box/chat-box.component';
import { NgFor, NgIf, AsyncPipe } from '@angular/common';
import { IonIcon } from '@ionic/angular/standalone';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { addIcons } from 'ionicons';
import { createOutline, playCircleOutline, trashOutline } from 'ionicons/icons';

@Component({
	selector: 'app-chat',
	templateUrl: './chat.component.html',
	styleUrl: './chat.component.scss',
	imports: [
		IonIcon,
		PageComponent,
		NgFor,
		NgIf,
		ChatBoxComponent,
		AsyncPipe,
		IconMenuComponent,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatComponent implements OnInit {
	public defaultPrompt = '[INST] {prompt} [/INST]';
	public chatThreads: Array<Thread> = [];
	public activeThread: Thread | undefined;

	public model: Model | undefined;
	private models: Model[] = [];

	private subscriptions: Subscription[] = [];

	constructor(
		private server: ServerService,
		private chatService: ChatService,
		private configService: ConfigService,
		public promptService: PromptService,
		private modelService: ModelService,
		private ipcService: ElectronIpcService,
		private cd: ChangeDetectorRef
	) {
		addIcons({
			'create-outline': createOutline,
			'trash-outline': trashOutline,
			'play-circle-outline': playCircleOutline,
		});
	}

	async ngOnInit() {
		await this.refreshThreadList();

		this.subscriptions.push(
			this.chatService.onThreadUpdate$.subscribe(() => {
				this.refreshThreadList();
			}),
			this.chatService.onThreadAdded$.subscribe(() => {
				this.refreshThreadList();
			}),
			this.chatService.onStartNewThread$.subscribe(() => {
				this.openNewThread();
			})
		);

		const activeThreadId = this.chatService.getActiveThreadId();
		if (activeThreadId) {
			const activeThread = this.chatThreads?.find(
				(v) => v.id === activeThreadId
			);
			if (activeThread) {
				this.activeThread = activeThread;
			}
		}
		if (!this.activeThread && this.chatThreads?.length) {
			this.activeThread = this.chatThreads[0];
		}
		if (!this.activeThread) {
			this.activeThread = {
				id: this.server.id('thr'),
			};
		}

		this.models = await this.modelService.getModels();
		this.subscriptions.push(
			this.configService.config$.subscribe((config) => {
				const model = this.models?.find(
					(m) => m.id == config?.data['model-svc'].currentModelId
				);
				if (model) {
					this.model = model;
				}
			})
		);
	}

	ionViewWillLeave() {
		for (const sub of this.subscriptions) {
			sub.unsubscribe();
		}
	}

	public async setThreadAsActive(thread: Thread) {
		this.activeThread = thread;
		console.debug('Loading thread', {
			threadId: thread.id,
		});
		if (!thread.id) {
			return;
		}

		this.chatService.setActiveThreadId(thread.id);
		this.cd.markForCheck();
	}

	public num(
		threadId: string | undefined,
		promptList: Prompt[] | null
	): number {
		if (!promptList) {
			return -1;
		}
		if (!threadId) {
			return -1;
		}
		let index = -1;
		for (const [index_, p] of promptList.entries()) {
			if (p.threadId == threadId) {
				index = index_;
			}
		}
		return index;
	}

	public async openNewThread() {
		this.activeThread = {
			id: this.server.id('thr'),
		};
		console.debug('Opened empty thread', {
			threadId: this.activeThread.id,
		});
	}

	public removeThread(thread: Thread) {
		if (!thread.id) {
			return;
		}
		this.chatService.chatThreadDelete(thread.id);
		this.refreshThreadList();
	}

	public async refreshThreadList() {
		const rsp = await this.chatService.chatThreads();
		this.chatThreads = rsp.threads!;
		if (!this.chatThreads?.length) {
			this.chatThreads = [];
		}
		this.cd.markForCheck();
	}

	public onCopyToClipboard(text: any) {
		this.ipcService.send(WindowApiConst.COPY_TO_CLIPBOARD_REQUEST, text);
	}

	trackById(_: number, message: { id?: string }): string {
		return message.id || '';
	}
}
