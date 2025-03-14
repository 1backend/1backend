/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
module.exports = {
	packagerConfig: {
		name: 'OpenOrch',
		executableName: 'singulatron',
		icon: './workspaces/electron-app/main/assets/icons/icon.png',
	},
	hooks: {
		// afterPackaging: 'node ./workspaces/electron-app/main/after-copy-hook.ts',
	},
	makers: [
		{
			name: '@electron-forge/maker-dmg',
			config: {
				icon: './workspaces/electron-app/main/assets/icons/icon.icns',
			},
		},
		{
			name: '@electron-forge/maker-deb',
			config: {
				icon: './workspaces/electron-app/main/assets/icons/icon.png',
			},
		},
		{
			name: '@electron-forge/maker-squirrel',
			config: {
				icon: './workspaces/electron-app/main/assets/icons/icon.ico',
				certificateFile: './mycert.pfx',
				certificatePassword: process.env.SIGNING_PASS,
				signWithParams: `/fd sha256 /a /f mycert.pfx /p ${process.env.SIGNING_PASS}`,
				animationComments: [
					'https://github.com/electron-userland/electron-installer-windows/issues/34',
					'https://github.com/electron-userland/electron-installer-windows/blob/cd3ff22d5036ce9086e1f017e38ae7631a531988/src/installer.js#L124',
				],
				animation: './template/assets/loading.gif',
			},
		},
	],
	plugins: [
		{
			name: '@electron-forge/plugin-webpack',
			config: {
				mainConfig: './webpack.main.config.js',
				renderer: {
					config: './webpack.renderer.config.js',
					entryPoints: [
						{
							html: './workspaces/electron-app/renderer/index.html',
							js: './workspaces/electron-app/renderer/index.ts',
							name: 'main_window',
							preload: {
								js: './workspaces/electron-app/renderer/preload.ts',
							},
						},
					],
				},
			},
		},
	],
};
