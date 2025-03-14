import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ChatComponent } from './chat.component';
import { IonicModule } from '@ionic/angular';
import { ServerService } from '../services/server.service';
import { ChatService } from '../services/chat.service';
import { ConfigService } from '../services/config.service';
import { ModelService } from '../services/model.service';
import { ElectronIpcService } from '../services/electron-ipc.service';
import { PromptService } from '../services/prompt.service';
import { ChangeDetectorRef } from '@angular/core';
import { of } from 'rxjs';

describe('ChatComponent', () => {
	let component: ChatComponent;
	let fixture: ComponentFixture<ChatComponent>;

	let chatServiceMock: any;

	beforeEach(async () => {
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


		const configServiceMock = {
			config$: of({ model: { currentModelId: '1' } }),
		};
		const modelServiceMock = jasmine.createSpyObj('ModelService', [
			'getModels',
		]);
		modelServiceMock.getModels.and.returnValue(Promise.resolve([]));

		await TestBed.configureTestingModule({
			imports: [IonicModule.forRoot()],
			providers: [
				{ provide: ChatService, useValue: chatServiceMock },
				{ provide: ConfigService, useValue: configServiceMock },
				{ provide: ModelService, useValue: modelServiceMock },
				ServerService,
				ElectronIpcService,
				PromptService,
				ChangeDetectorRef,
			],
		}).compileComponents();

		fixture = TestBed.createComponent(ChatComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});
