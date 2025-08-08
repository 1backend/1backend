import { ComponentFixture, TestBed } from '@angular/core/testing';
import { UsersComponent } from './users.component';
import { UserService, User } from '../services/user.service';
import { ToastController } from '@ionic/angular';
import { FormBuilder } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';
import { of, ReplaySubject } from 'rxjs';
import { TranslatePipe } from '../translate.pipe';
import { ChangeDetectorRef } from '@angular/core';
import { RouterTestingModule } from '@angular/router/testing';

describe('UsersComponent', () => {
	let component: UsersComponent;
	let fixture: ComponentFixture<UsersComponent>;
	let userServiceMock: any;
	let toastControllerMock: any;
	let routerMock: any;
	let activatedRouteMock: any;
	let changeDetectorReferenceMock: any;

	beforeEach(async () => {
		userServiceMock = {
			user$: new ReplaySubject<User>(1).asObservable(),
			getUsers: jasmine.createSpy('getUsers').and.returnValue(
				Promise.resolve({
					users: [
						{
							id: '1',
							name: 'John Doe',
							email: 'john.doe@example.com',
							createdAt: '2023-01-01',
							updatedAt: '2023-01-01',
						},
					],
				})
			),
			saveProfile: jasmine
				.createSpy('saveProfile')
				.and.returnValue(Promise.resolve()),
			changePasswordAdmin: jasmine
				.createSpy('changePasswordAdmin')
				.and.returnValue(Promise.resolve()),
			deleteUser: jasmine
				.createSpy('deleteUser')
				.and.returnValue(Promise.resolve()),
		};

		toastControllerMock = jasmine.createSpyObj('ToastController', ['create']);
		toastControllerMock.create.and.returnValue(
			Promise.resolve({ present: jasmine.createSpy('present') })
		);

		routerMock = {
			events: of({}),
			navigate: jasmine.createSpy('navigate'),
		};

		activatedRouteMock = {
			queryParams: of({ search: '' }),
		};

		changeDetectorReferenceMock = jasmine.createSpyObj('ChangeDetectorRef', [
			'markForCheck',
		]);

		await TestBed.configureTestingModule({
			imports: [
				UsersComponent,
				RouterTestingModule, // Add RouterTestingModule
			],
			providers: [
				FormBuilder,
				{ provide: UserService, useValue: userServiceMock },
				{ provide: ToastController, useValue: toastControllerMock },
				{ provide: Router, useValue: routerMock },
				{ provide: ActivatedRoute, useValue: activatedRouteMock },
				{ provide: ChangeDetectorRef, useValue: changeDetectorReferenceMock },
				TranslatePipe,
			],
		}).compileComponents();

		fixture = TestBed.createComponent(UsersComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});

	it('should call getUsers on ngOnInit', async () => {
		expect(userServiceMock.getUsers).toHaveBeenCalled();
		await fixture.whenStable();
		expect(component.users.length).toBe(1);
		expect(component.users[0].name).toBe('John Doe');
	});

	it('should filter users based on search text', async () => {
		component.filterUsers('john');
		expect(component.filteredUsers.length).toBe(1);
		expect(component.filteredUsers[0].name).toBe('John Doe');

		component.filterUsers('jane');
		expect(component.filteredUsers.length).toBe(0);
	});

	it('should save user profile', async () => {
		await component.saveUser('1');
		expect(userServiceMock.saveProfile).toHaveBeenCalledWith(
			'john.doe@example.com',
			'John Doe'
		);
		expect(userServiceMock.changePasswordAdmin).not.toHaveBeenCalled();
	});

	it('should delete user', async () => {
		await component.deleteUser(new Event('click'), '1');
		expect(userServiceMock.deleteUser).toHaveBeenCalledWith('1');
	});
});
