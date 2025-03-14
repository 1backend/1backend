/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Component,
	Input,
	Output,
	EventEmitter,
	ChangeDetectionStrategy,
} from '@angular/core';
import { ChatService } from '../../../services/chat.service';
import { ChatSvcMessage as Message, ModelSvcPlatform } from '@openorch/client';
import { PromptService } from '../../../services/prompt.service';
import { ServerService } from '../../../services/server.service';
import { MarkdownComponent, provideMarkdown } from 'ngx-markdown';
import { IonIcon } from '@ionic/angular/standalone';
import { NgIf, DatePipe, AsyncPipe } from '@angular/common';
import { MobileService } from '../../../services/mobile.service';
import { UserService } from '../../../services/user.service';
import { addIcons } from 'ionicons';
import {
	personCircleOutline,
	copyOutline,
	reloadOutline,
	trashOutline,
} from 'ionicons/icons';

import { environment } from '../../../../environments/environment';
import { ModelService } from '../../../../app/services/model.service';
import {
	ModelSvcModel as Model,
} from '@openorch/client';

@Component({
	selector: 'app-message',
	templateUrl: './message.component.html',
	styleUrl: './message.component.scss',
	providers: [provideMarkdown()],
	imports: [
		MarkdownComponent,
		IonIcon,
		NgIf,
		MarkdownComponent,
		DatePipe,
		AsyncPipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MessageComponent {
	constructor(
		private chatService: ChatService,
		private promptService: PromptService,
		public userService: UserService,
		private server: ServerService,
		public mobile: MobileService,
		public modelService: ModelService,
	) {
		addIcons({
			'person-circle-outline': personCircleOutline,
			'copy-outline': copyOutline,
			'reload-outline': reloadOutline,
			'trash-outline': trashOutline,
		});
	}
	hasAsset = false;
	models: Model[] = [];
	currentModelName: string | undefined = "";

	@Input() message!: Message;
	@Input() streaming: boolean = false;
	@Input() modelId: string = '';

	@Output() onCopyToClipboard = new EventEmitter<string>();

	async ngOnInit() {

		if (this.message?.fileIds?.length) {
			this.hasAsset = true;
		}
	}


	formatPlatformName(platforms:ModelSvcPlatform[]| null):string{
		if(!platforms){
			return ""
		}
		 if(!this.message.meta){
			return ""
		 }
		 const pid = this.message.meta['platformId']
		 return platforms.find(p => {
			return p.id == pid
		 })?.name || ""
	} 

	async regenerateAnswer(message: Message) {
		if (message.userId) {
			return;
		}

		await this.promptService.promptAdd({
			id: this.server.id('msg'),
			prompt: message.text!,
			threadId: message.threadId as string,
			modelId: this.modelId as string,
		});
	}

	asset(message: Message): string {
		return `${environment.serverAddress}/file-svc/serve/upload/${message.fileIds![0]}`;
	}

	deleteMessage(messageId: string | undefined) {
		if (messageId === undefined) {
			return;
		}
		this.chatService.chatMessageDelete(messageId);
	}

	propagateCopyToClipboard(text: string | undefined) {
		if (text === undefined) {
			return;
		}
		this.onCopyToClipboard.emit(text);
	}
}
