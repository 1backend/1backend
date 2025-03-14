/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { IpcMainEvent } from 'electron';
import * as si from 'systeminformation';

// import { OnOSInfo } from 'shared-lib/models/event-request-response';

export function requestOsInfo(event: IpcMainEvent) {
	console.log('Received request for OS info');

	si.osInfo()
		.then((data) => {
			// let ev: OnOSInfo = {
			// 	platform: data.platform,
			// 	distro: data.distro,
			// 	release: data.release,
			// 	arch: data.arch,
			// 	hostname: data.hostname,
			// };
			// event.sender.send(WindowApiConst.ON_OS_INFO, ev);
		})
		.catch((error) => {
			// event.sender.emit(WindowApiConst.ON_OS_INFO, {
			// 	error: error,
			// });
		});
}
