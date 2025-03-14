/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { IpcMainEvent } from 'electron';
import * as si from 'systeminformation';

import { GraphicsInfoRequest } from 'shared-lib/models/event-request-response';

export function graphicsInfoRequest(
	event: IpcMainEvent,
	args: GraphicsInfoRequest
) {
	console.log('Received request for graphics info');

	si.graphics()
		.then((data) => {
			// The 'controllers' property contains information about the video cards
			if (data.controllers && data.controllers.length > 0) {
				// let ev: OnGraphicsInfo = {
				// 	controllers: data.controllers as any,
				// };
				// event.sender.send(WindowApiConst.ON_GRAPHICS_INFO, ev);
			}
		})
		.catch((error) => {
			// event.sender.send(WindowApiConst.ON_GRAPHICS_INFO, {
			// 	error: error,
			// });
		});
}
