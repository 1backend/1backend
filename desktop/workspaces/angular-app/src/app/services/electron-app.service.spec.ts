import { TestBed } from '@angular/core/testing';

import { ElectronAppService } from './electron-app.service';

describe('ElectronAppServiceService', () => {
	let service: ElectronAppService;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		service = TestBed.inject(ElectronAppService);
	});

	it('should be created', () => {
		expect(service).toBeTruthy();
	});
});
