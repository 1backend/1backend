import { ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';
import { of } from 'rxjs';
import { Router } from '@angular/router';
import { ChangeDetectorRef } from '@angular/core';
import { ChatBoxComponent } from './chat-box.component';
import { ServerService } from '../../services/server.service';
import { ChatService } from '../../services/chat.service';
import { PromptService } from '../../services/prompt.service';
import { ElectronAppService } from '../../services/electron-app.service';
import { MobileService } from '../../services/mobile.service';
import { FooterService } from '../../services/footer.service';
import { UserService } from '../../services/user.service';

describe('ChatBoxComponent', () => {
	let component: ChatBoxComponent;
	let fixture: ComponentFixture<ChatBoxComponent>;

	let chatServiceMock: any;
	let promptServiceMock: any;
	let routerMock: any;
	let serverMock: any;
	let mobileServiceMock: any;
	let footerServiceMock: any;
	let userServiceMock: any;

	beforeEach(async () => {
		chatServiceMock = jasmine.createSpyObj('ChatService', [
			'chatMessages',
			'onMessageAdded$',
			'onThreadUpdate$',
			'onThreadAdded$',
		]);
		chatServiceMock.chatMessages.and.returnValue(
			Promise.resolve({ messages: [], assets: [] })
		);
		chatServiceMock.onMessageAdded$ = of();
		chatServiceMock.onThreadUpdate$ = of();
		chatServiceMock.onThreadAdded$ = of();

		promptServiceMock = jasmine.createSpyObj('PromptService', [
			'promptAdd',
			'promptSubscribe',
			'onPromptListUpdate$',
		]);
		promptServiceMock.onPromptListUpdate$ = of();

		routerMock = {
			events: of(),
			url: '/chat',
		};

		serverMock = jasmine.createSpyObj('ServerService', ['uuid']);
		serverMock.uuid.and.returnValue('mock-uuid');

		mobileServiceMock = jasmine.createSpyObj('MobileService', [
			'getMobileStatus',
		]);
		mobileServiceMock.getMobileStatus.and.returnValue(false);

		footerServiceMock = jasmine.createSpyObj('FooterService', [
			'updateFooterComponent',
		]);

		userServiceMock = jasmine.createSpyObj('UserService', ['init']);

		await TestBed.configureTestingModule({
			imports: [IonicModule.forRoot()],
			providers: [
				{ provide: ChatService, useValue: chatServiceMock },
				{ provide: PromptService, useValue: promptServiceMock },
				{ provide: Router, useValue: routerMock },
				{ provide: ServerService, useValue: serverMock },
				{ provide: MobileService, useValue: mobileServiceMock },
				{ provide: FooterService, useValue: footerServiceMock },
				{ provide: UserService, useValue: userServiceMock },
				ElectronAppService,
				ChangeDetectorRef,
			],
		}).compileComponents();

		fixture = TestBed.createComponent(ChatBoxComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});

	// Additional tests can be added here
});
