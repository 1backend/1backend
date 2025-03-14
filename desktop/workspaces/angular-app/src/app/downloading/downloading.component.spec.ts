import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DownloadingComponent } from './downloading.component';

describe('DownloadingComponent', () => {
	let component: DownloadingComponent;
	let fixture: ComponentFixture<DownloadingComponent>;

	beforeEach(async () => {
		await TestBed.configureTestingModule({
			imports: [DownloadingComponent],
		}).compileComponents();

		fixture = TestBed.createComponent(DownloadingComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});
