import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HomeComponent } from './home.component';
import { LogService } from '../services/log.service';
import { UserService } from '../services/user.service';

describe('HomeComponent', () => {
	let component: HomeComponent;
	let fixture: ComponentFixture<HomeComponent>;
	let logServiceMock: any;

	let userServiceMock: any;

	beforeEach(async () => {
		// Create mock implementations of the services
		logServiceMock = {
			logStatus: jasmine
				.createSpy('logStatus')
				.and.returnValue(Promise.resolve({ enabled: true })),
			logEnable: jasmine
				.createSpy('logEnable')
				.and.returnValue(Promise.resolve()),
			logDisable: jasmine
				.createSpy('logDisable')
				.and.returnValue(Promise.resolve()),
		};

		userServiceMock = {
			noop: jasmine.createSpy('noop'),
		};

		await TestBed.configureTestingModule({
			imports: [HomeComponent],
			providers: [
				{ provide: LogService, useValue: logServiceMock },

				{ provide: UserService, useValue: userServiceMock },
			],
		}).compileComponents();

		fixture = TestBed.createComponent(HomeComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});

	it('should initialize loggingEnabled based on logStatus', async () => {
		await component.ngOnInit();
		expect(component.loggingEnabled).toBe(true);
		expect(logServiceMock.logStatus).toHaveBeenCalled();
	});

	// it('should disable logging', async () => {
	// 	component.loggingEnabled = true;
	// 	await component.disableLog();
	// 	expect(component.loggingEnabled).toBe(false);
	// 	expect(logServiceMock.logDisable).toHaveBeenCalled();
	// 	expect(logServiceMock.logStatus).toHaveBeenCalled();
	// });

	it('should enable logging', async () => {
		component.loggingEnabled = false;
		await component.enableLog();
		expect(component.loggingEnabled).toBe(true);
		expect(logServiceMock.logEnable).toHaveBeenCalled();
		expect(logServiceMock.logStatus).toHaveBeenCalled();
	});
});
