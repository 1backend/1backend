/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import { FirehoseService } from './firehose.service';
import { ReplaySubject } from 'rxjs';
import { UserService } from './user.service';
import { first } from 'rxjs';
import {
	ConfigSvcApi,
	ConfigSvcListConfigsResponse,
	Configuration,
	ConfigSvcConfig,
} from '@1backend/client';

@Injectable({
	providedIn: 'root',
})
export class ConfigService {
	private configService!: ConfigSvcApi;

	lastConfigs!: { [key: string]: ConfigSvcConfig };

	configsSubject = new ReplaySubject<{ [key: string]: ConfigSvcConfig }>(1);

	/** Config emitted whenever it's loaded (on startup) or saved */
	configs$ = this.configsSubject.asObservable();

	constructor(
		private server: ServerService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.init();
		this.userService.user$.pipe(first()).subscribe(() => {
			this.configService = new ConfigSvcApi(
				new Configuration({
					basePath: this.server.addr(),
					apiKey: this.server.token(),
				})
			);
			this.loggedInInit();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'configUpdate': {
					const rsp = await this.listConfigs('', []);
					this.configsSubject.next(rsp.configs);
					break;
				}
			}
		});
	}

	async loggedInInit() {
		try {
			const rsp = await this.listConfigs('', []);
			console.log('Config loaded', rsp);
			this.lastConfigs = rsp?.configs ? rsp?.configs : {};

			this.configsSubject.next(this.lastConfigs);
		} catch (error) {
			console.log(error);
			console.error('Error in pollConfig', {
				error: JSON.stringify(error),
			});
		}
	}

	async listConfigs(
		app: string,
		slugs: string[]
	): Promise<ConfigSvcListConfigsResponse> {
		return await this.configService.listConfigs({
			body: {
				app,
				slugs,
			},
		});
	}
}
