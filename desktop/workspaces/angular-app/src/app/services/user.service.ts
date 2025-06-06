/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import { CookieService } from 'ngx-cookie-service';
import {
	ReplaySubject,
	Observable,
	firstValueFrom,
	first,
	BehaviorSubject,
} from 'rxjs';
import { Router } from '@angular/router';
import {
	Configuration,
	UserSvcApi,
	UserSvcListUsersResponse,
	UserSvcLoginResponse,
	UserSvcReadSelfResponse,
	UserSvcUser,
	UserSvcListUsersRequest,
	UserSvcSaveUserRequest,
	UserSvcChangePasswordRequest,
	UserSvcResetPasswordRequest,
	UserSvcUserInput,
} from '@1backend/client';

@Injectable({
	providedIn: 'root',
})
export class UserService {
	private userService!: UserSvcApi;

	private token: string = '';

	private userSubject = new ReplaySubject<UserSvcUser>(1);
	/** Current logged in user */
	public user$ = this.userSubject.asObservable();

	private userCache: {
		[id: string]: BehaviorSubject<UserSvcUser | undefined>;
	} = {};

	constructor(
		private server: ServerService,
		private cookieService: CookieService,
		private router: Router
	) {
		this.userService = new UserSvcApi(
			new Configuration({
				basePath: this.server.addr(),
				apiKey: this.server.token(),
			})
		);
		this.init();
	}

	noop() {}

	async init() {
		this.getToken();

		if (!this.hasToken()) {
			try {
				const rsp = await this.login('singulatron', 'changeme');
				this.setToken(rsp.token!.token as string);
			} catch {
				console.error('Login with default credentials failed');
				this.router.navigateByUrl('/login');
				return;
			}

			if (!this.hasToken()) {
				console.error('Something is wrong with the setting of cookies');
				this.router.navigateByUrl('/login');
				return;
			}
		}

		try {
			const rsp = await this.readSelf();
			this.userSubject.next(rsp.user!);
		} catch (error) {
			console.error('Cannot read user even with a token', error);
			this.router.navigateByUrl('/login');
		}
	}

	getToken(): string {
		if (this.token) {
			return this.token;
		}
		this.token = this.cookieService.get('the_token');
		return this.token;
	}

	setToken(token: string) {
		this.token = token;
		this.cookieService.set(
			'the_token',
			this.token,
			3650,
			'/',
			'',
			this.server.config.env.production ? true : false
		);
	}

	removeToken() {
		this.cookieService.delete(
			'the_token',
			'/',
			'',
			this.server.config.env.production ? true : false
		);
	}

	hasToken(): boolean {
		const t = this.cookieService.get('the_token');
		return !!t;
	}

	login(slug: string, password: string): Promise<UserSvcLoginResponse> {
		return this.userService.login({
			body: { slug: slug, password: password },
		});
	}

	readSelf(): Promise<UserSvcReadSelfResponse> {
		return this.userService.readSelf({});
	}

	getUsers(
		request: UserSvcListUsersRequest
	): Promise<UserSvcListUsersResponse> {
		return this.userService.listUsers({
			body: request,
		});
	}

	/** Save user. For admins. */
	saveUser(id: string, name: string): Promise<object> {
		const request: UserSvcSaveUserRequest = {
			name: name,
		};
		return this.userService.saveUser({
			userId: id,
			body: request,
		});
	}

	changePassword(
		currentPassword: string,
		newPassword: string
	): Promise<object> {
		const request: UserSvcChangePasswordRequest = {
			currentPassword: currentPassword,
			newPassword: newPassword,
		};
		return this.userService.changePassword({
			body: request,
		});
	}

	resetPassword(userId: string, newPassword: string) {
		const request: UserSvcResetPasswordRequest = {
			newPassword: newPassword,
		};
		this.userService.resetPassword({
			userId: userId,
			body: request,
		});

		return;
	}

	/** Create a user - alternative to registration
	 */
	createUser(
		user: UserSvcUserInput,
		password: string,
		roleIds: string[]
	): Promise<object> {
		return this.userService.createUser({
			body: {
				user: user,
				password: password,
				roleIds: roleIds,
			},
		});
	}

	deleteUser(userId: string): Promise<object> {
		return this.userService.deleteUser({
			userId: userId,
		});
	}

	async getUserId(): Promise<string> {
		try {
			const user = await firstValueFrom(this.user$.pipe(first()));
			return user.id as string;
		} catch (error) {
			console.error('Error getting user ID:', error);
			throw error;
		}
	}

	getUser(id: string): Observable<UserSvcUser | undefined> {
		if (!this.userCache[id]) {
			this.userCache[id] = new BehaviorSubject<UserSvcUser | undefined>(
				undefined
			);
			this.getUsers({
				ids: [id],
			}).then((rsp) => {
				this.userCache[id].next(rsp.users![0]);
			});
		}

		return this.userCache[id].asObservable();
	}
}
