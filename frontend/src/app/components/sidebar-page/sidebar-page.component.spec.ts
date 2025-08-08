import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SidebarPageComponent } from './sidebar-page.component';

describe('SidebarPageComponent', () => {
	let component: SidebarPageComponent;
	let fixture: ComponentFixture<SidebarPageComponent>;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		fixture = TestBed.createComponent(SidebarPageComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});
