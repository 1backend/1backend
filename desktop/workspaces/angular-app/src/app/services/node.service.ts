/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import {
	RegistrySvcApi,
	RegistrySvcListNodesResponse,
	Configuration,
} from '@openorch/client';
import { UserService } from './user.service';
import { first } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class NodeService {
	private registryService!: RegistrySvcApi;

	constructor(
		private server: ServerService,
		private userService: UserService
	) {
		this.registryService = new RegistrySvcApi(
			new Configuration({
				basePath: this.server.addr(),
				apiKey: this.server.token(),
			})
		);

		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async nodesList(): Promise<RegistrySvcListNodesResponse> {
		return this.registryService.listNodes({});
	}

	async init() {}
}
