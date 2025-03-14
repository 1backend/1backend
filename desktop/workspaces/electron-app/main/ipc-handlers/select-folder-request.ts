/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { IpcMainEvent, dialog } from 'electron';
import { WindowApiConst } from 'shared-lib';
import { SelectFolderRequest } from 'shared-lib/models/event-request-response';

export function selectFolderRequest(): (
	event: IpcMainEvent,
	args: SelectFolderRequest
) => void {
	return (event: IpcMainEvent, args: SelectFolderRequest) => {
		dialog
			.showOpenDialog({
				properties: ['openDirectory'],
				buttonLabel: 'Select Folder',
			})
			.then((result) => {
				if (result) {
					let location = result.filePaths[0];
					if (!location) {
						return;
					}
					let data = {
						location: result.filePaths[0],
					};
					// sender will receive this and save the selected folder to server
					event?.sender.send(WindowApiConst.ON_FOLDER_SELECT, data);
				}
			});
	};
}
