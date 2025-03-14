/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
export interface SelectFolderRequest {}

export interface OnFolderSelect {
	location: string;
}

export interface OnModelCheck {
	assetsReady?: boolean;
}

export interface onModelReady {
	modelReady: boolean;
}

export interface OnOSInfo {
	platform: string;
	distro: string;
	release: string;
	arch: string;
	hostname: string;
}

export interface ModelLaunchRequest {}

export interface OnModelLaunch {
	error?: string;
}

export interface OnDockerInfo {
	/** Is the docker daemon running and dockerode can talk to it? */
	hasDocker: boolean;
	dockerDaemonAddress?: string;
	wsl?: boolean;
	error?: string;
}

export interface GraphicsInfoRequest {}

export interface OnGraphicsInfo {
	error?: string;
	controllers?: Array<{
		model: string;
		vendor: string;
		vram: number;
		vramDynamic: boolean;
	}>;
}

export interface OnSystemLanguage {
	systemLanguage: string;
}
