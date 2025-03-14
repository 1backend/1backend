/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable, NgZone } from '@angular/core';
import { WindowApi } from 'shared-lib';

@Injectable({
	providedIn: 'root',
})
export class ElectronIpcService {
	private _api!: WindowApi;

	constructor(private zone: NgZone) {
		if (window && (window as Window).api) {
			this._api = (window as Window).api;
			console.log('Preloader API has been loaded successfully');
		} else {
			console.warn('Preloader API is not loaded');
		}
	}

	public receive<Out>(channel: string, callback: (output: Out) => void): void {
		if (this._api) {
			this._api.receive<Out>(channel, (output) => {
				// console.log(`Received from main process channel [${channel}]`, output);

				// Change detection fix https://stackoverflow.com/a/49136353/11480016
				this.zone.run(() => {
					callback(output);
				});
			});
		}
	}

	public send<In>(channel: string, input: In): void {
		if (this._api) {
			// avoid infinite loop triggering here
			if (channel != 'logRequest') {
				console.debug(`Sending IPC message to main process`, {
					channel: channel,
				});
			}
			this._api.send<In>(channel, input);
		}
	}
}
