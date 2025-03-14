/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
// @todo get this from some config
export var apiURL = 'http://127.0.0.1:58231';

import { post } from './api';

import { logEndpoint, LogRequest } from './app';

const MAX_REQUESTS = 100;
const TIME_WINDOW = 60000;

let requestCount = 0;
let timeWindowStart = Date.now();

export async function sendLogToServer(req: LogRequest): Promise<any> {
	try {
		if (Date.now() > timeWindowStart + TIME_WINDOW) {
			timeWindowStart = Date.now();
			requestCount = 0;
		}

		if (requestCount >= MAX_REQUESTS) {
			return;
		}

		const response = await post(apiURL + logEndpoint, req);
		return response;
	} catch (error) {
		// we don't log anything here as that will be also tried to be sent
		// so it ends up in an infinite loop
	}
}
