/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { app, BrowserWindow, shell, globalShortcut } from 'electron';
import { Window } from './window';
import { ipcMain, clipboard } from 'electron';
import { WindowApiConst } from 'shared-lib';
import { selectFolderRequest } from '../ipc-handlers/select-folder-request';

import { getWebpackPath, setupSession } from './util';
import { Machine } from '../the-machine/the-machine';
import { enableLogging, disableLogging } from './logger';
import { Log } from '../../../angular-app/shared/backend-api/app';
import { registerShortcuts } from './shortcuts/shortcuts';

declare const global: Global;

app.on('will-quit', () => {
	globalShortcut.unregisterAll();
});

export class App {
	private static _wrapper: Window;
	private static machine: Machine = new Machine(getWebpackPath());
	private static subscriptionsUp = false;

	public static launch(): void {
		app.on('window-all-closed', App.quit);
		app.on('activate', App.start);
		app.whenReady().then(App.start);

		app.on('web-contents-created', App.openExternalLinksInDefaultBrowser);

		app.on('ready', async () => {
			// Startup.
			// THIS IS WHERE IT ALL STARTS.

			setupSession();

			ipcMain.on(WindowApiConst.FRONTEND_READY_REQUEST, (event, args) => {
				console.log('Frontend ready');

				registerShortcuts(App.electronWindow as BrowserWindow, globalShortcut);

				try {
					const systemLanguage: string = app.getLocale().split('-')[0];
					App.electronWindow?.webContents.send(
						WindowApiConst.ON_SYSTEM_LANGUAGE,
						{ systemLanguage: systemLanguage }
					);
				} catch (error) {
					console.error('Error while retrieving system language:', error);
				}

				if (App.subscriptionsUp) {
					return;
				}
				App.subscriptionsUp = true;

				App.machine.eventService.onRuntimeInstallLog$.subscribe(
					(info: string) => {
						App.electronWindow?.webContents.send(
							WindowApiConst.ON_RUNTIME_INSTALL_LOG,
							info
						);
					}
				);
			});

			ipcMain.on(WindowApiConst.LOG_REQUEST, (event, log: Log) => {
				if (log.level === 'info') {
					console.info(log);
				} else if (log.level === 'debug') {
					console.debug(log);
				} else if (log.level === 'warn') {
					console.warn(log);
				} else if (log.level === 'error') {
					console.error(log);
				} else {
					console.log(log);
				}
			});

			ipcMain.on(WindowApiConst.DISABLE_LOGGING_REQUEST, (event, args: any) => {
				disableLogging();
			});

			ipcMain.on(WindowApiConst.ENABLE_LOGGING_REQUEST, (event, args: any) => {
				enableLogging();
			});

			ipcMain.on(WindowApiConst.SELECT_FOLDER_REQUEST, selectFolderRequest());

			ipcMain.on(
				WindowApiConst.COPY_TO_CLIPBOARD_REQUEST,
				(event, args: string) => {
					App.copyToClipboard(args);
				}
			);

			ipcMain.on(
				WindowApiConst.INSTALL_RUNTIME_REQUEST,
				(event, args: string) => {
					App.machine.eventService.installRuntimeRequest.next();
				}
			);
		});
	}

	public static get electronWindow(): BrowserWindow | undefined {
		return this._wrapper ? this._wrapper.electronWindow : undefined;
	}

	private static start() {
		if (!App.electronWindow) {
			App._wrapper = new Window();
		}
	}

	private static quit() {
		if (
			process.platform !== 'darwin' ||
			(global as any).appConfig.configId === 'e2e-test'
		) {
			app.quit();
		}
	}

	private static openExternalLinksInDefaultBrowser = (
		event: Electron.Event,
		contents: Electron.WebContents
	) => {
		contents.setWindowOpenHandler((handler: Electron.HandlerDetails) => {
			shell.openExternal(handler.url);

			return { action: 'deny' };
		});

		contents.on(
			'will-navigate',
			(event: Electron.Event, navigationUrl: string) => {
				const parsedUrl = new URL(navigationUrl);
				if (parsedUrl.origin !== 'http://localhost:4200') {
					event.preventDefault();
				}
			}
		);
	};

	private static copyToClipboard = (text: string) => {
		clipboard.writeText(text);
	};
}
