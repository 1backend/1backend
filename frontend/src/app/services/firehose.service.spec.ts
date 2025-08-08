import { TestBed } from '@angular/core/testing';

import { FirehoseService } from './firehose.service';

describe('FirehoseService', () => {
	let service: FirehoseService;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		service = TestBed.inject(FirehoseService);
	});

	it('should be created', () => {
		expect(service).toBeTruthy();
	});
});
