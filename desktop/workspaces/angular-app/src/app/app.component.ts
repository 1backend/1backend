/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component } from '@angular/core';
import { ElectronIpcService } from './services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { Log } from '../../shared/backend-api/app';
import { IonRouterOutlet } from '@ionic/angular/standalone';

let loggingEnabled = true;

export function enableLogging() {
	loggingEnabled = true;
}

export function disableLogging() {
	loggingEnabled = false;
}

const originalConsole = {
	log: console.log,
	error: console.error,
	warn: console.warn,
	info: console.info,
	debug: console.debug,
	trace: console.trace,
};

function overrideConsole(ipcService: ElectronIpcService) {
	for (const mn of ['log', 'error', 'warn', 'info', 'debug', 'trace']) {
		const methodName: keyof Console = mn as any;

		console[methodName] = ((...arguments_: any[]) => {
			if (!loggingEnabled) {
				return;
			}
			(originalConsole as any)[methodName](...arguments_);
			try {
				const request: Log = {
					level: methodName,
					message: arguments_[0],
					source: 'frontend',
					fields: arguments_.length > 1 ? arguments_[1] : {},
				};
				ipcService.send(WindowApiConst.LOG_REQUEST, request);
			} catch (error) {
				originalConsole.error('Cannot send log to IPC', error);
			}
		}) as any;
	}
}

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss'],
	imports: [IonRouterOutlet],
})
export class AppComponent {
	title = 'singulatron-angular-app';

	constructor(private ipcService: ElectronIpcService) {
		overrideConsole(this.ipcService);
	}

	ngOnInit(): void {}
}
