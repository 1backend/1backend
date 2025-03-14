/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component } from '@angular/core';
import {
	FormBuilder,
	FormGroup,
	Validators,
	FormsModule,
	ReactiveFormsModule,
} from '@angular/forms';
import { UserService } from '../services/user.service';
import { first } from 'rxjs';
import {
	ToastController,
	IonList,
	IonLabel,
	IonItem,
	IonIcon,
	IonCardContent,
	IonRow,
	IonCol,
	IonButton,
} from '@ionic/angular/standalone';
import { TranslatePipe } from '../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { NgFor, NgIf } from '@angular/common';
import { CenteredComponent } from '../components/centered/centered.component';
import { ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { Router, ActivatedRoute } from '@angular/router';
import { QueryParser } from '../services/query.service';
import {
	DatastoreFilter,
	UserSvcUser,
	DatastoreOp,
	UserSvcGetUsersRequest,
} from '@openorch/client';

interface UserVisible extends UserSvcUser {
	visible?: boolean;
}

@Component({
	selector: 'app-users',
	templateUrl: './users.component.html',
	styleUrls: ['./users.component.scss'],
	standalone: true,
	imports: [
		IonList,
		IonLabel,
		IonItem,
		IonIcon,
		IonCardContent,
		IonRow,
		IonCol,
		IonButton,
		PageComponent,
		IconMenuComponent,
		CenteredComponent,
		NgFor,
		FormsModule,
		ReactiveFormsModule,
		NgIf,
		TranslateModule,
		TranslatePipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UsersComponent {
	users: UserVisible[] = [];
	after: any;
	private userForms: Map<string, FormGroup> = new Map();

	count = 0;
	searchTerm = '';
	queryParser: QueryParser;

	constructor(
		private fb: FormBuilder,
		private userService: UserService,
		private toast: ToastController,
		private cd: ChangeDetectorRef,
		private router: Router,
		private activatedRoute: ActivatedRoute
	) {
		this.userForms = new Map();

		this.queryParser = new QueryParser();
		this.queryParser.defaultConditionFunc = (value: any): DatastoreFilter => {
			return {
				fields: ['name', 'slug'],
				jsonValues: JSON.stringify([value]),
				op: 'containsSubstring' as DatastoreOp,
			};
		};

		this.userService.user$.pipe(first()).subscribe(() => {
			this.initializeOnLogin();
		});
	}

	private async initializeOnLogin() {
		this.activatedRoute.queryParams.subscribe(async (parameters) => {
			this.searchTerm =
				this.queryParser.convertQueryParamsToSearchTerm(parameters);

			await this.fetchUsers();
			this.cd.markForCheck();

			this.userForms?.clear();

			for (const user of this.users) {
				this.userForms.set(user.id!, this.createUserForm(user));
			}
		});
	}

	public redirect() {
		const query = this.queryParser.parse(this.searchTerm);

		const kv = filtersToKeyValue(
			query.filters
				? query.filters.filter((v) => {
						return !v.fields?.includes('name') && v.fields?.includes('slug');
					})
				: []
		);

		if (Object.keys(kv)?.length) {
			this.router.navigate([], {
				queryParams: kv,
			});
			return;
		}

		if (this.searchTerm) {
			this.router.navigate([], {
				queryParams: { search: this.searchTerm },
			});
			return;
		}

		this.router.navigate([], {
			queryParams: {},
		});
	}

	public toggleVisible(userId: string) {
		for (const u of this.users) {
			if (u.id == userId) {
				u.visible = !u.visible;
			}
		}
	}

	public async fetchUsers() {
		const query = this.queryParser.parse(this.searchTerm);
		query.count = true;
		query.filters = query.filters || [];

		const request: UserSvcGetUsersRequest = {
			query: query,
		};

		if (this.after) {
			request.query!.jsonAfter = JSON.stringify([this.after]);
		}

		const response = await this.userService.getUsers(request);

		// eslint-disable-next-line
		if (response.users && this.after) {
			this.users = [...this.users, ...response.users];
		} else if (response.users) {
			this.users = response.users;
		} else {
			this.users = [];
		}

		this.count = response.count || 0;

		// eslint-disable-next-line
		if (response.after && response.after != `0001-01-01T00:00:00Z`) {
			this.after = response.after;
		} else {
			this.after = undefined;
		}

		this.cd.markForCheck();
	}

	createUserForm(user: UserVisible): FormGroup {
		return this.fb.group({
			name: [user.name, Validators.required],
			slug: [{ value: user.slug, disabled: true }, [Validators.required]],
			password: [''],
			passwordConfirmation: [''],
			createdAt: [{ value: user.createdAt, disabled: true }],
			updatedAt: [{ value: user.updatedAt, disabled: true }],
			visible: [user.visible || false],
		});
	}

	getUserForm(userId: string): FormGroup {
		return this.userForms.get(userId)!;
	}

	async saveUser(userId: string) {
		const userForm = this.getUserForm(userId);
		if (userForm.invalid) {
			return;
		}

		const { name, slug, password, passwordConfirmation } = userForm.value;

		if (password && password !== passwordConfirmation) {
			const toast = await this.toast.create({
				message: 'Passwords do not match',
				duration: 5000,
				color: 'danger',
				cssClass: 'white-text',
				position: 'middle',
			});
			toast.present();
			return;
		}

		try {
			let toastMessage = `User '${name}' saved`;
			await this.userService.saveProfile(userId, slug, name);

			if (password) {
				toastMessage += ' and password changed';
				await this.userService.resetPassword(userId, password);
			}

			const toast = await this.toast.create({
				color: 'secondary',
				message: toastMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();

			this.initializeOnLogin();
		} catch (error) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(error as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
				cssClass: 'white-text',
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}

	async deleteUser($event: any, userId: string) {
		$event.stopPropagation();

		try {
			await this.userService.deleteUser(userId);

			const toastMessage = `User ${name} deleted`;

			const toast = await this.toast.create({
				color: 'secondary',
				message: toastMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();

			this.initializeOnLogin();
		} catch (error) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(error as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
				cssClass: 'white-text',
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}

	trackById(_: number, message: { id?: string }): string {
		return message.id || '';
	}

	async loadMoreData() {
		if (!this.after) {
			console.log('No more users to load');
			return;
		}
		await this.fetchUsers();
	}
}

function filtersToKeyValue(filters: DatastoreFilter[]): {
	[key: string]: any;
} {
	if (!filters) {
		return {};
	}
	const object: { [key: string]: any } = {};

	for (const filter of filters) {
		object[filter.fields![0]] = JSON.parse(filter.jsonValues!)[0];
	}

	return object;
}
