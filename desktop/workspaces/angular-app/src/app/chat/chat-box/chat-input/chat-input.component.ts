/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	Output,
	EventEmitter,
	ViewChild,
	ViewContainerRef,
	ComponentFactoryResolver,
	AfterViewInit,
	OnInit,
	ViewEncapsulation,
} from '@angular/core';
import {
	IonIcon,
	IonTextarea,
	IonFabButton,
	IonInput,
} from '@ionic/angular/standalone';
import { Subscription } from 'rxjs';
import { TranslatePipe } from '../../../translate.pipe';
import { FormsModule } from '@angular/forms';
import {
	CharacterService,
	Character,
} from '../../../services/character.service';
import { CharacterComponent } from '../../character/character.component';
import { ModelService } from '../../../services/model.service';
import {
	ModelSvcModel as Model,
	PromptSvcEngineParameters,
} from '@openorch/client';
import { ConfigService } from '../../../services/config.service';
import { ChatService } from '../../../services/chat.service';
import { addIcons } from 'ionicons';
import { addOutline, arrowUpOutline, removeOutline } from 'ionicons/icons';
import { CommonModule } from '@angular/common';

export interface SendOutput {
	message: string;
	characterId: string;
	modelId: string;
	engineParameters: PromptSvcEngineParameters;
}

@Component({
	selector: 'app-chat-input',
	imports: [
		IonIcon,
		IonTextarea,
		IonFabButton,
		FormsModule,
		TranslatePipe,
		FormsModule,
		CommonModule,
		IonInput,
	],
	templateUrl: './chat-input.component.html',
	styleUrls: ['./chat-input.component.scss'],
	encapsulation: ViewEncapsulation.None,
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatInputComponent implements OnInit, AfterViewInit {
	public model: Model | undefined;
	private models: Model[] = [];
	public message: string = '';
	private subscriptions: Subscription[] = [];
	private characterModal!: CharacterComponent;

	@ViewChild('modalContainer', { read: ViewContainerRef })
	modalContainer!: ViewContainerRef;
	@Output() sends = new EventEmitter<SendOutput>();

	constructor(
		private modelService: ModelService,
		private characterService: CharacterService,
		private configService: ConfigService,
		private cd: ChangeDetectorRef,
		private chatService: ChatService,
		private resolver: ComponentFactoryResolver
	) {
		addIcons({
			'add-outline': addOutline,
			'arrow-up-outline': arrowUpOutline,
			'remove-outline': removeOutline,
		});
	}

	width: number = 200;
	height: number = 200;
	steps: number = 2;
	cfgScale: number = 7.5;
	hrScale: number = 2;

	jsonData: string = `{
  "stableDiffusion": {
    "txt2Img": {}
  }
}`
	parsedJson: any;

	async ngOnInit() {
		this.models = await this.modelService.getModels();
		this.cd.markForCheck();

		this.subscriptions.push(
			this.configService.config$.subscribe(async (config) => {
				this.model = this.models?.find(
					(m) => m.id == config?.data['model-svc']?.currentModelId
				);
			})
		);
	}

	isInvalidJson: boolean = false;

	expanded: boolean = false;
	isAdvancedConfigVisible: boolean = false;

	toggleExpand() {
		this.expanded = !this.expanded;
	}

	toggleAdvancedConfig() {
		this.isAdvancedConfigVisible = !this.isAdvancedConfigVisible
	}

	getPanelHeight() {
		if (this.isAdvancedConfigVisible) {
			return this.expanded ? '300px' : '0px';
		}
		return this.expanded ? '150px' : '0px';
	}

	ngAfterViewInit() {
		// Initial modal load if needed
		this.loadCharacterModal();
	}

	ionViewWillLeave() {
		for (const sub of this.subscriptions) {
			sub.unsubscribe();
		}
	}

	private loadCharacterModal() {
		try {
			const factory = this.resolver.resolveComponentFactory(CharacterComponent);
			this.modalContainer.clear();
			const componentReference = this.modalContainer.createComponent(factory);
			this.characterModal = componentReference.instance;
		} catch (error) {
			console.error('Error loading character modal:', error);
		}
	}

	public hasNonWhiteSpace(value: string): boolean {
		if (!value) {
			return false;
		}
		return /\S/.test(value);
	}

	handleKeydown(event: KeyboardEvent): void {
		if (event.key === 'Enter' && !event.shiftKey) {
			event.preventDefault();
			if (this.hasNonWhiteSpace(this.message)) {
				this.send();
			}
		} else if (event.key === 'Enter' && event.shiftKey) {
			event.preventDefault();
			this.message += '\n';
		}
	}

	async send() {
		if (/^demo\w*$/.test(this.message)) {
			this.demo(this.message);
			return;
		}
		const message = this.message;
		this.message = '';

		let engineParameters: PromptSvcEngineParameters = {};

		if (this.model?.platformId == 'stable-diffusion') {
			if (this.isAdvancedConfigVisible) {
				try {
					this.parsedJson = JSON.parse(this.jsonData);
					engineParameters = this.parsedJson;
				} catch (error) {
					console.error('Invalid JSON:', error);
					this.isInvalidJson = true;
					return;
				}
			} else {
				engineParameters = {
					stableDiffusion: {
						txt2Img: {
							steps: this.steps,
							width: this.width,
							height: this.height,
							cfg_scale: this.cfgScale,
							hr_scale: this.hrScale,
						},
					},
				};
			}
		}

		const character = this.getSelectedCharacter();
		this.sends.emit({
			characterId: character?.id || '',
			message: message,
			modelId: this.model?.id || '',
			engineParameters: engineParameters
		});
	}

	public getSelectedCharacter(): Character {
		return this.characterService.selectedCharacter;
	}

	public showCharacterModal() {
		this.loadCharacterModal();
		if (this.characterModal) {
			this.characterModal.show();
		} else {
			console.error('Character modal is not available.');
		}
	}

	public getSelectedCharacterText(): string {
		const selectedCharacter = this.getSelectedCharacter();
		return selectedCharacter.data.name || 'None';
	}

	async demo(message: string) {
		this.message = '';
		for (const question of demoQuestions) {
			await this.typeQuestion(question, 0);
			await this.send();
			if (message == 'demothreads') {
				this.chatService.onStartNewThreadSubject.next();
			}
			await new Promise((resolve) => setTimeout(resolve, 300));
		}
	}

	async typeQuestion(question: string, index: number): Promise<void> {
		if (index < question.length) {
			this.message = this.message + question.charAt(index);
			await new Promise((resolve) => setTimeout(resolve, 30));
			this.cd.markForCheck();
			return this.typeQuestion(question, index + 1);
		}
		return;
	}
}

const demoQuestions = [
	'When do experts predict the technological singularity will occur, and what might its implications be?',
	'How could quantum computing revolutionize industries like cryptography and pharmaceuticals?',
	'What are the main ethical concerns surrounding the deployment of AI in society?',
	'What are the goals and timelines of current missions to Mars and other planets?',
	'What are some of the most promising technologies being developed to combat climate change?',
	'How will the rollout of 5G technology impact the Internet of Things (IoT) and smart cities?',
	'What are the potential benefits and risks associated with the widespread adoption of blockchain technology?',
	'What advancements are being made in autonomous vehicle technology, and when might we see fully self-driving cars on the roads?',
	'How might CRISPR and other gene-editing technologies transform medicine and agriculture?',
	'What are the latest developments in virtual and augmented reality, and how might they change entertainment and education?',
];
