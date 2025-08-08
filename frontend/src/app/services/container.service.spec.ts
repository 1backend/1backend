import { TestBed } from '@angular/core/testing';

import { ContainerService } from './docker.service';

describe('ContainerService', () => {
	let service: ContainerService;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		service = TestBed.inject(ContainerService);
	});

	it('should be created', () => {
		expect(service).toBeTruthy();
	});
});
