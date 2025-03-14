/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { sendLogToServer } from '../../../angular-app/shared/backend-api/app_backend';
import { Log } from '../../../angular-app/shared/backend-api/app';

export function sendLogToLocalServer(log: Log) {
	try {
		sendLogToServer({
			logs: [log],
		});
	} catch (err) {
		console.error('Error saving logs', err);
	}
}
