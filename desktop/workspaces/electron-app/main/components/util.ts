/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { app, session } from 'electron';
import { join, sep } from 'path';

export function setupSession() {
	session.defaultSession.webRequest.onHeadersReceived((details, callback) => {
		const responseHeaders = { ...details.responseHeaders };

		delete responseHeaders['access-control-allow-origin'];

		responseHeaders['Access-Control-Allow-Origin'] = ['*'];

		responseHeaders['Access-Control-Allow-Methods'] = [
			'GET',
			'POST',
			'OPTIONS',
			'PUT',
			'DELETE',
		];

		callback({
			responseHeaders: {
				...responseHeaders,
			},
		});
	});
}

export function getWebpackPath(): string {
	return app.isPackaged
		? join(app.getAppPath(), '.webpack')
		: __dirname.replace(sep + 'main', '');
}

export function getAngularPath(): string {
	return app.isPackaged
		? join(app.getAppPath(), '.webpack', 'renderer', 'angular_window')
		: join(
				__dirname.replace('.webpack' + sep + 'main', ''),
				'workspaces',
				'angular-app',
				'src'
			);
}
