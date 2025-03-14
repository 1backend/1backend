
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AdvancedModelExplorerComponent } from './advanced-model-explorer.component';
import { ModelService, Model } from '../../services/model.service';
import { DownloadService } from '../../services/download.service';
import { ConfigService } from '../../services/config.service';
import { of } from 'rxjs';
import { IonicModule } from '@ionic/angular';
import { DecimalPipe } from '@angular/common';
import { FormsModule } from '@angular/forms';

describe('AdvancedModelExplorerComponent', () => {
	let component: AdvancedModelExplorerComponent;
	let fixture: ComponentFixture<AdvancedModelExplorerComponent>;
	let modelService: jasmine.SpyObj<ModelService>;
	let downloadService: jasmine.SpyObj<DownloadService>;

	const mockModels: Model[] = [
		{
			id: 'huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf',
			assets: {
				MODEL:
					'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf',
			},
			platformId: 'PlatformLlamaCpp.Id', // Example placeholder
			name: 'Mistral',
			parameters: '7B',
			flavour: 'Instruct',
			version: 'v0.2',
			quality: 'Q2_K',
			extension: 'GGUF',
			fullName: 'Mistral 7B Instruct v0.2 Q2_K',
			size: 3.08,
			maxRam: 5.58,
			quantComment:
				'smallest, significant quality loss - not recommended for most purposes',
			description: 'mistral description', // Replace with actual description string
			promptTemplate: '[INST] {prompt} [/INST]',
		},
		{
			id: 'huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf',
			assets: {
				MODEL:
					'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf',
			},
			platformId: 'PlatformLlamaCpp.Id', // Example placeholder
			name: 'Mistral',
			parameters: '7B',
			flavour: 'Instruct',
			version: 'v0.2',
			quality: 'Q3_K_S',
			extension: 'GGUF',
			fullName: 'Mistral 7B Instruct v0.2 Q3_K_S',
			size: 3.16,
			maxRam: 5.66,
			quantComment: 'very small, high quality loss',
			description: 'mistral description', // Replace with actual description string
			promptTemplate: '[INST] {prompt} [/INST]',
		},
	];

	beforeEach(async () => {
		const modelServiceSpy = jasmine.createSpyObj('ModelService', [
			'getModels',
			'makeDefault',
		]);
		const downloadServiceSpy = jasmine.createSpyObj('DownloadService', [
			'downloadList',
			'downloadDo',
		]);
		const configServiceSpy = jasmine.createSpyObj('ConfigService', [
			'config$',
		]);

		(configServiceSpy.config$ as jasmine.Spy).and.returnValue(
			of({ model: { currentModelId: '1' } })
		);

		await TestBed.configureTestingModule({
			imports: [IonicModule.forRoot(), FormsModule, DecimalPipe],
			providers: [
				{ provide: ModelService, useValue: modelServiceSpy },
				{ provide: DownloadService, useValue: downloadServiceSpy },
				{ provide: ConfigService, useValue: configServiceSpy },
			],
		}).compileComponents();

		fixture = TestBed.createComponent(AdvancedModelExplorerComponent);
		component = fixture.componentInstance;
		modelService = TestBed.inject(ModelService) as jasmine.SpyObj<ModelService>;
		downloadService = TestBed.inject(
			DownloadService
		) as jasmine.SpyObj<DownloadService>;

		modelService.getModels.and.returnValue(Promise.resolve(mockModels));
		downloadService.downloadList.and.returnValue(
			Promise.resolve({ downloads: [] })
		);

		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});

	it('should initialize models on ngOnInit', async () => {
		await component.ngOnInit();
		expect(component.allModels).toEqual(mockModels);
		expect(component.allFilteredModels).toEqual(mockModels);
		expect(component.totalItems).toBe(mockModels.length);
	});

	it('should toggle grid and list views', () => {
		expect(component.gridView).toBe(true);
		component.switchGridListView();
		expect(component.gridView).toBe(false);
	});

	it('should handle download action', async () => {
		const model = mockModels[0];
		await component.download(model);
		expect(downloadService.downloadDo).toHaveBeenCalledWith(
			'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf'
		);
	});

	it('should detect large screen correctly', () => {
		spyOnProperty(window, 'innerWidth').and.returnValue(2000);
		component.detectLargeScreen();
		expect(component.veryLargeScreen).toBe(true);
	});

	it('should handle model category click', async () => {
		const option = component.modelCategoryOptions[0]; // 'Instruct' category
		component.modelCategoryClicked(option);
		expect(option.active).toBe(true);
		expect(modelService.getModels).toHaveBeenCalled();
	});
});
