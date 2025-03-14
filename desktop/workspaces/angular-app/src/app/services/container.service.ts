/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ReplaySubject } from 'rxjs';
import { OnDockerInfo } from 'shared-lib/models/event-request-response';
import { ServerService } from './server.service';
import {
	ContainerSvcApi,
	Configuration,
	ContainerSvcDaemonInfoResponse,
} from '@openorch/client';

@Injectable({
	providedIn: 'root',
})
export class ContainerService {
	private daemonService!: ContainerSvcApi;

	onContainerInfoSubject = new ReplaySubject<OnDockerInfo>(1);
	onContainerInfo$ = this.onContainerInfoSubject.asObservable();

	constructor(private server: ServerService) {
		// @todo nothing to trigger docker info yet
		// so we fall back to pollings
		setInterval(() => {
			this.init();
		}, 2000);
	}

	initInProgress = false;
	async init() {
		try {
			if (this.initInProgress) {
				return;
			}
			this.initInProgress = true;
			if (!this.daemonService) {
				this.daemonService = new ContainerSvcApi(
					new Configuration({
						basePath: this.server.addr(),
						apiKey: this.server.token(),
					})
				);
			}

			const rsp = await this.daemonInfo();

			this.onContainerInfoSubject.next({
				hasDocker: rsp?.available || false,
			});
		} catch (error) {
			console.error('Error in docker.service init', {
				error: JSON.stringify(error),
			});
		} finally {
			this.initInProgress = true;
		}
	}

	async daemonInfo(): Promise<ContainerSvcDaemonInfoResponse> {
		return this.daemonService.containerDaemonInfo();
	}
}
