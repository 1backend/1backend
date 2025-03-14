/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import './logger';

import * as remoteMain from '@electron/remote/main';
import { app, BrowserWindow, ipcMain, nativeImage } from 'electron';
import * as path from 'node:path';
import express from 'express';
import { AppConfig } from 'shared-lib';
import fs from 'fs';
import { getAngularPath, getWebpackPath } from './util';
import cp from 'child_process';
import { Log } from '../../../angular-app/shared/backend-api/app';

declare const global: Global;

declare global {
	interface Global {
		appConfig: AppConfig;
	}
}

const PORT = 59517;
let serverProcess: cp.ChildProcessWithoutNullStreams;

installWindows();
launchFrontendServer();
launchServer();

declare const MAIN_WINDOW_PRELOAD_WEBPACK_ENTRY: string;

export class Window {
	private _electronWindow: BrowserWindow | undefined;

	constructor() {
		this.createWindow();
		this.loadRenderer();
	}

	private createWindow(): void {
		this._electronWindow = new BrowserWindow({
			width: 1280,
			height: 720,
			minWidth: 500,
			minHeight: 700,
			backgroundColor: '#FFFFFF',
			icon: this.loadIcon(),
			webPreferences: {
				nodeIntegration: global.appConfig.isNodeIntegration,
				contextIsolation: global.appConfig.isContextIsolation,
				sandbox: global.appConfig.isSandbox,
				preload: MAIN_WINDOW_PRELOAD_WEBPACK_ENTRY,
			},
			autoHideMenuBar: true,
		});

		if (global.appConfig.isEnableRemoteModule) {
			remoteMain.enable(this._electronWindow.webContents);
		}
	}

	private loadIcon(): Electron.NativeImage | undefined {
		let iconObject;
		if (global.appConfig.isIconAvailable) {
			const iconPath = path.join(__dirname, 'icons/icon.png');

			iconObject = nativeImage.createFromPath(iconPath);
			if (iconObject && process.platform === 'darwin') {
				app.dock.setIcon(iconObject);
			}
		}
		return iconObject;
	}

	private loadRenderer(): void {
		let env = global.appConfig.configId;

		if (env === 'development') {
			this.electronWindow?.loadURL(`http://localhost:4200`);
		} else {
			this.electronWindow?.loadURL(`http://127.0.0.1:${PORT}`);
		}

		if (global.appConfig.isOpenDevTools) {
			this.openDevTools();
		}

		this._electronWindow?.on('closed', () => {
			ipcMain.removeAllListeners();
			delete this._electronWindow;
		});
	}

	private openDevTools(): void {
		this._electronWindow?.webContents.openDevTools();
		this._electronWindow?.webContents.on('devtools-opened', () => {
			this._electronWindow?.focus();
			setImmediate(() => {
				this._electronWindow?.focus();
			});
		});
	}

	public get electronWindow(): BrowserWindow | undefined {
		return this._electronWindow;
	}
}

function launchFrontendServer() {
	const server = express();

	const angularAssetsParentFolder = getAngularPath();

	server.use(express.static(angularAssetsParentFolder));

	server.get('*', (req, res) => {
		const requestedPath = path.join(angularAssetsParentFolder, req.path);

		fs.stat(requestedPath, (err, stats) => {
			if (err || !stats.isFile()) {
				res.sendFile(path.join(angularAssetsParentFolder, 'index.html'));
			} else {
				res.sendFile(requestedPath);
			}
		});
	});

	const s = server.listen(PORT, () =>
		console.log(`App is serving`, {
			address: `http://localhost:${PORT}`,
		})
	);

	app.on('before-quit', (event) => {
		console.log('Application is quitting - closing HTTP server');
		s.close(() => {
			console.log('HTTP server closed');
		});
	});

	process.on('SIGINT', () => {
		console.log('SIGINT signal received: closing HTTP server');
		s.close(() => {
			console.log('HTTP server closed');
		});
	});

	process.on('SIGTERM', () => {
		console.log('SIGTERM signal received: closing HTTP server');
		s.close(() => {
			console.log('HTTP server closed');
		});
	});
}

// let's try to install register the app icon into the Start Menu on Windows
// with squirrel
// https://github.com/electron/forge/issues/191
function installWindows() {
	if (process.platform != 'win32') {
		return;
	}

	function executeSquirrelCommand(args: string[], done: () => void) {
		const primaryPath = path.resolve(
			path.dirname(process.execPath),
			'..',
			'Update.exe'
		);
		const alternativePath = path.resolve(
			path.dirname(process.execPath),
			'Update.exe'
		);

		function executeUpdate(updatePath: string) {
			var child = cp.spawn(updatePath, args, { detached: true });

			child.on('close', function (code) {
				done();
			});

			child.on('error', function (err) {
				console.log('Could not install icon in start menu', err);
				done();
			});
		}

		fs.access(primaryPath, fs.constants.F_OK, (err) => {
			if (!err) {
				executeUpdate(primaryPath);
			} else {
				fs.access(alternativePath, fs.constants.F_OK, (err) => {
					if (!err) {
						executeUpdate(alternativePath);
					} else {
						console.error('Update.exe not found in expected locations.');
						done();
					}
				});
			}
		});
	}

	function install(done: () => void) {
		var target = process.execPath;
		executeSquirrelCommand(['--createShortcut', target], done);
	}

	// @todo fix this
	// function uninstall(done: () => void) {
	// 	var target = process.execPath;
	// 	executeSquirrelCommand(['--removeShortcut', target], done);
	// }

	install(() => {
		console.log('Shortcut created.');
	});
}

process.on('SIGTERM', () => {
	console.log('SIGTERM signal received: closing Localron');
	if (serverProcess) {
		serverProcess.kill();
	}
});

function launchServer() {
	console.log('Launching Server');

	app.on('before-quit', (event) => {
		console.log('Application is quitting - closing Server');
		if (serverProcess) {
			serverProcess.kill();
			console.log('Server closed');
		}
	});

	process.on('SIGINT', () => {
		console.log('SIGINT signal received: closing Server');
		if (serverProcess) {
			serverProcess.kill();
		}
	});

	let exeName = 'server';
	if (process.platform == 'win32') {
		exeName += '.exe';
	}
	let thePath = path.join(getWebpackPath(), exeName);

	if (!fs.existsSync(thePath)) {
		setInterval(() => {
			`Server at '${thePath}' does not exist`;
		}, 3000);

		return;
	}

	try {
		var child = cp.spawn(thePath, [], {
			detached: false,
			stdio: ['pipe', 'pipe', 'pipe'],
		});
		serverProcess = child;

		child.stdout.on('data', (data) => {
			consoleLogServer(data?.toString()?.replace(/\n+$/, ''));
		});

		child.stderr.on('data', (data) => {
			consoleLogServer(data?.toString()?.replace(/\n+$/, ''));
		});

		child.on('close', function (code) {
			if (code != null && code > 0) {
				console.error('Server closed with error', {
					code: code,
				});
				return;
			}
			console.log(`Server closed successfully`);
		});

		child.on('error', function (err) {
			console.error(`Failed to start Server`, {
				error: err,
			});
		});
	} catch (err) {
		console.error('Error spawning server', {
			error: JSON.stringify(err),
		});
		throw `Error spawning Server: ${err}`;
	}
}

function splitStringByNewline(inputString: string) {
	return inputString.split(/\r?\n/);
}

function consoleLogServer(jsonMessages: string) {
	// this split is needed because sometimes two or more logs
	// are coming back from server, like this:
	// {one json}\n{another json}
	// this happens when they are logged at exactly the same time
	const lines = splitStringByNewline(jsonMessages);
	lines.forEach((jsonMessage) => {
		let msg: { [key: string]: any };
		try {
			msg = JSON.parse(jsonMessage);
		} catch (err) {
			console.error('Error parsing Server line', {
				line: jsonMessage,
			});
			return;
		}

		let log: Log = {
			level: (msg.level as string)?.toLowerCase(),
			fields: {},
			time: msg.time as string,
			message: msg.msg as string,
			source: 'server',
		};

		for (let k in msg) {
			if (k !== 'level' && k !== 'time' && k !== 'msg') {
				log.fields[k] = msg[k];
			}
		}

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
}
