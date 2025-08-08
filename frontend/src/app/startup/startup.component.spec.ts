import { ComponentFixture, TestBed } from '@angular/core/testing';
import { StartupComponent } from './startup.component';
import { ElectronIpcService } from '../services/electron-ipc.service';
import { ElectronAppService } from '../services/electron-app.service';
import { ConfigService } from '../services/config.service';
import { DownloadService } from '../services/download.service';
import { ModelService } from '../services/model.service';
import { ContainerService } from '../services/docker.service';
import { of } from 'rxjs';

describe('StartupComponent', () => {
	let component: StartupComponent;
	let fixture: ComponentFixture<StartupComponent>;

	let electronIpcServiceMock: any;
	let electronAppServiceMock: any;
	let configServiceMock: any;
	let downloadServiceMock: any;
	let modelServiceMock: any;
	let dockerServiceMock: any;

	beforeEach(async () => {
		// Create mock implementations of the services
		electronIpcServiceMock = {
			send: jasmine.createSpy('send'),
		};

		electronAppServiceMock = {
			onRuntimeInstallLog$: of('Installation log'),
			onFolderSelect$: of({ location: '/path/to/folder' }),
		};

		configServiceMock = {
			lastConfig: {
				model: { currentModelId: '1' },
			},
		};

		downloadServiceMock = {
			downloadDo: jasmine.createSpy('downloadDo'),
		};

		modelServiceMock = {
			getModels: jasmine.createSpy('getModels').and.returnValue(
				Promise.resolve([
					{
						id: '1',
						name: 'Model 1',
						flavour: 'flavour1',
						version: 'v1.0',
						assets: { asset1: 'url1' },
					},
				])
			),
			onModelCheck$: of({ assetsReady: true }),
			onModelLaunch$: of({}),
		};

		dockerServiceMock = {
			onDockerInfo$: of({ hasDocker: true }),
		};

		await TestBed.configureTestingModule({
			providers: [
				{ provide: ElectronIpcService, useValue: electronIpcServiceMock },
				{ provide: ElectronAppService, useValue: electronAppServiceMock },
				{ provide: ConfigService, useValue: configServiceMock },
				{ provide: DownloadService, useValue: downloadServiceMock },
				{ provide: ModelService, useValue: modelServiceMock },
				{ provide: ContainerService, useValue: dockerServiceMock },
			],
		}).compileComponents();

		fixture = TestBed.createComponent(StartupComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});

	it('should log installation logs and update the view accordingly', () => {
		// Test the logging functionality
		component.ngOnInit();
		fixture.detectChanges();
		expect(component.log).toContain('Installation log');
	});

	it('should handle model selection correctly', async () => {
		await component.ngOnInit();
		fixture.detectChanges();
		expect(component.models.length).toBe(1);
		expect(component.models[0].name).toBe('Model 1');
	});

	it('should toggle sections correctly', () => {
		component.toggleSection('model');
		expect(component.showSections.model).toBeTrue();

		component.toggleSection('model');
		expect(component.showSections.model).toBeFalse();
	});

	it('should start downloading assets correctly', async () => {
		await component.download();
		expect(downloadServiceMock.downloadDo).toHaveBeenCalledWith('url1');
	});

	it('should send IPC messages for folder select and runtime installation', () => {
		component.openFolderSelect();
		expect(electronIpcServiceMock.send).toHaveBeenCalledWith(
			'selectFolderRequest',
			{}
		);

		component.installRuntime();
		expect(electronIpcServiceMock.send).toHaveBeenCalledWith(
			'installRuntimeRequest',
			{}
		);
		expect(component.isRuntimeInstalling).toBeTrue();
	});
});
