/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';

@Injectable({
	providedIn: 'root',
})
export class LogService {
	constructor() {}

	async logDisable(): Promise<void> {}

	async logEnable(): Promise<void> {}

	async logStatus(): Promise<LoggingStatus> {
		return {} as any;
	}
}

interface LoggingStatus {
	enabled: boolean;
}
