/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import sudo from 'sudo-prompt';
import { EventService } from '../services/event-service';
import { promises as fs } from 'fs';
import path from 'path';
import os from 'os';
import { join } from 'path';
import { execSync, exec } from 'child_process';

const username = os.userInfo().username;

export class OSManager {
	logFilePath: string;

	constructor(
		private assetFolder: string,
		private eventService: EventService
	) {
		this.eventService.installRuntimeRequest$.subscribe(() => {
			console.info('Install runtime initiated');
			this.initializeEnvironment();
		});

		// @todo fix path, config service is now in server
		this.logFilePath = path.join(os.homedir(), 'openorch_install.log');

		this.initLogFile();

		new FileWatcher(this.logFilePath, (log) => {
			this.eventService.onRuntimeInstallLogSubject.next(log);
		});
	}

	async initLogFile() {
		try {
			let fd = await fs.open(this.logFilePath, 'w');
			fd.close();
		} catch (err) {
			console.error('Error opening log file:', err);
		}
	}

	async initializeEnvironment(): Promise<void> {
		const platform = os.platform();

		if (platform === 'win32') {
			await this.setupWindows();
		} else if (platform == 'linux' || platform == 'darwin') {
			await this.setupLinux();
		} else {
			console.log(`Running on ${platform}, no special setup required.`);
		}
	}

	async setupWindows(): Promise<void> {
		let exePath = join(this.assetFolder, 'dapper.exe');

		// chcp 65001 sets the code page to UTF-8 in the Command Prompt.
		const command = `chcp 65001 && "${exePath}" --var-username=${username} --var-assetfolder=${this.assetFolder} run --retry=3 "${join(this.assetFolder, 'app.json')}" > "${this.logFilePath}" 2>&1`;
		if (isAdminWindows()) {
			console.info('Already admin - running command normally');
			await this.executeCommand(command);
		} else {
			console.info('Not an admin - running command with sudo prompt');
			await this.executeCommandSudoPrompt(command);
		}
	}

	async setupLinux(): Promise<void> {
		let exePath = join(this.assetFolder, 'dapper');

		const command = `"${exePath}" --var-username=${username} --var-assetfolder=${this.assetFolder} run --retry=3 "${join(this.assetFolder, 'app.json')}" > "${this.logFilePath}" 2>&1`;
		await this.executeCommandSudoPrompt(command);
	}

	private async executeCommand(command: string): Promise<void> {
		return new Promise((resolve, reject) => {
			exec(command, (error, stdout, stderr) => {
				if (error) {
					reject(error);
				} else {
					resolve();
				}
			});
		});
	}

	private async executeCommandSudoPrompt(command: string): Promise<void> {
		return new Promise((resolve, reject) => {
			sudo.exec(
				command,
				{
					name: 'OpenOrch Environment Setup',
					icns: getIconPath(this.assetFolder),
				},
				(error, stdout, stderr) => {
					if (error) {
						console.error('Error:', error);
						reject(error);
					} else {
						console.log('stdout:', stdout);
						if (stderr) console.error('stderr:', stderr);
						resolve();
					}
				}
			);
		});
	}
}

function getIconPath(assetFolder: string): string {
	switch (os.platform()) {
		case 'darwin':
			return path.join(assetFolder, 'main', 'icons', 'icon.icns');
		case 'win32':
			return path.join(assetFolder, 'main', 'icons', 'icon.ico');
		case 'linux':
			return path.join(assetFolder, 'main', 'icons', 'icon.png');
		default:
			return '';
	}
}

class FileWatcher {
	filePath: string;
	lastContent: string;

	constructor(
		filePath: string,
		private callback: (s: string) => void
	) {
		this.filePath = filePath;
		this.lastContent = '';
		this.watchFile();
	}

	async readFile() {
		try {
			let data = await fs.readFile(this.filePath, 'utf8');

			if (data !== this.lastContent) {
				this.lastContent = data;
			}

			this.callback(this.lastContent);
		} catch (err) {
			console.error('Error reading file:', err);
		}
	}

	watchFile() {
		setInterval(() => this.readFile(), 500);
	}
}

function isAdminWindows(): boolean {
	try {
		// Attempt to execute a command that requires administrative privileges.
		execSync('net session', { stdio: 'ignore' });
		return true;
	} catch (error) {
		// If the command fails, assume the user does not have administrative privileges.
		return false;
	}
}
