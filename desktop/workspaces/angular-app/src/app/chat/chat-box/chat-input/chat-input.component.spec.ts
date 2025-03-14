import { TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';
import { UserService } from '../../../services/user.service';
import { ConfigService } from '../../../services/config.service';
import { ChatInputComponent } from './chat-input.component';
import { ModelService } from '../../../services/model.service';
import { ChatService } from '../../../services/chat.service';
import { CharacterService } from '../../../services/character.service';
import { of } from 'rxjs';

describe('ChatInputComponent', () => {
	let component: ChatInputComponent;

	let userServiceMock: jasmine.SpyObj<UserService>;
	let chatServiceMock: jasmine.SpyObj<ChatService>;

	beforeEach(async () => {
		userServiceMock = jasmine.createSpyObj('UserService', ['init']);

		const modelServiceMock = jasmine.createSpyObj('ModelService', [
			'getModels',
		]);
		modelServiceMock.getModels.and.returnValue(Promise.resolve([]));

		const configServiceMock = {
			config$: of({ model: { currentModelId: '1' } }),
		};

		chatServiceMock = jasmine.createSpyObj('ChatService', [
			'chatThreads',
			'chatMessages',
			'getActiveThreadId',
			'onThreadUpdate$',
			'onThreadAdded$',
			'onMessageAdded$',
		]);
		chatServiceMock.chatThreads.and.returnValue(
			Promise.resolve({ threads: [] })
		);
		chatServiceMock.chatMessages.and.returnValue(
			Promise.resolve({ threads: [] })
		);

		chatServiceMock.onThreadUpdate$ = of();
		chatServiceMock.onThreadAdded$ = of();
		chatServiceMock.onMessageAdded$ = of();

		chatServiceMock.chatThreads.and.returnValue(
			Promise.resolve({ threads: [] })
		);

		const characterServiceMock = jasmine.createSpyObj('CharacterService', [
			'loadCharacters',
		]);
		characterServiceMock.loadCharacters.and.returnValue(Promise.resolve([]));

		await TestBed.configureTestingModule({
			imports: [IonicModule.forRoot()],
			providers: [
				{ provide: UserService, useValue: userServiceMock },
				{ provide: ModelService, useValue: modelServiceMock },
				{ provide: ConfigService, useValue: configServiceMock },
				{ provide: ChatService, useValue: chatServiceMock },
				{ provide: CharacterService, useValue: characterServiceMock },
			],
		}).compileComponents();

		const fixture = TestBed.createComponent(ChatInputComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});
