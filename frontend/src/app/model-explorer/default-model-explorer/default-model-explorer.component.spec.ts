import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DefaultModelExplorerComponent } from './default-model-explorer.component';

describe('DefaultModelExplorerComponent', () => {
	let component: DefaultModelExplorerComponent;
	let fixture: ComponentFixture<DefaultModelExplorerComponent>;

	beforeEach(async () => {
		await TestBed.configureTestingModule({
			imports: [DefaultModelExplorerComponent],
		}).compileComponents();

		fixture = TestBed.createComponent(DefaultModelExplorerComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});
