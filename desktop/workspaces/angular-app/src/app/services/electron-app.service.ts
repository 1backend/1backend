/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
//import { ElectronIpcService } from './electron-ipc.service';
import { ReplaySubject } from 'rxjs';

import {
	OnFolderSelect,
	OnOSInfo,
	//OnSystemLanguage,
} from 'shared-lib/models/event-request-response';
//import { WindowApiConst } from 'shared-lib';
import { ApiService } from '../api.service';

@Injectable({
	providedIn: 'root',
})
export class ElectronAppService {
	onOsInfoSubject = new ReplaySubject<OnOSInfo>(1);
	onOsInfo$ = this.onOsInfoSubject.asObservable();

	onFolderSelectSubject = new ReplaySubject<OnFolderSelect>(1);
	/** Emitted when a folder selection finished  */
	onFolderSelect$ = this.onFolderSelectSubject.asObservable();

	onRuntimeInstallLogSubject = new ReplaySubject<string>(1);
	onRuntimeInstallLog$ = this.onRuntimeInstallLogSubject.asObservable();

	constructor(
		//private ipcService: ElectronIpcService,
		public apiService: ApiService
	) {
		this.listenToIpcEvents();
		//this.ipcService.send(WindowApiConst.FRONTEND_READY_REQUEST, {});
	}

	private listenToIpcEvents(): void {
		//this.ipcService.receive<OnFolderSelect>(
		//	WindowApiConst.ON_FOLDER_SELECT,
		//	(data) => {
		//		this.onFolderSelectSubject.next(data);
		//	}
		//);
		//
		//this.ipcService.receive<OnSystemLanguage>(
		//	WindowApiConst.ON_SYSTEM_LANGUAGE,
		//	(data) => {
		//		this.apiService.setLocale(data.systemLanguage);
		//	}
		//);
		//
		//this.ipcService.receive<string>(
		//	WindowApiConst.ON_RUNTIME_INSTALL_LOG,
		//	(data) => {
		//		this.onRuntimeInstallLogSubject.next(data);
		//	}
		//);
	}
}
