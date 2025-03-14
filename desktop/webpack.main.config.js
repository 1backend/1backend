/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const os = require('os');

const fs = require('fs');

const child_process = require('child_process');

class OpenOrchAfterEmitPlugin {
	apply(compiler) {
		compiler.hooks.afterEmit.tapPromise(
			'OpenOrchAfterEmitPlugin',
			async (compilation) => {
				const outputPath = compilation.outputOptions.path;

				fs.chmodSync(path.join(outputPath, '..', 'server'), 0o755);
				fs.chmodSync(path.join(outputPath, '..', 'dapper'), 0o755);

				// Rename for Windows
				if (os.platform() === 'win32') {
					let serverOriginal = path.join(outputPath, '..', 'server');
					let serverFinal = path.join(outputPath, '..', 'server.exe');
					let dapperOriginal = path.join(outputPath, '..', 'dapper');
					let dapperFinal = path.join(outputPath, '..', 'dapper.exe');

					try {
						fs.renameSync(serverOriginal, serverFinal);
						fs.renameSync(dapperOriginal, dapperFinal);
					} catch (err) {
						console.log('Failed to rename', err);
					}
				}

				// Code signing
				try {
					const password = fs
						.readFileSync(path.join(__dirname, 'mycert.pass'), 'utf-8')
						.trim();

					if (os.platform() === 'win32') {
						child_process.execSync(
							`signtool sign /fd sha256 /f "${path.join(__dirname, 'mycert.pfx')}" /p ${password} "${path.join(outputPath, '..', 'server.exe')}"`
						);
						child_process.execSync(
							`signtool sign /fd sha256 /f "${path.join(__dirname, 'mycert.pfx')}" /p ${password} "${path.join(outputPath, '..', 'dapper.exe')}"`
						);
					} else if (os.platform() === 'darwin') {
						child_process.execSync(
							`codesign --sign "Developer ID Application: Your Developer ID" --keychain /path/to/keychain "${path.join(outputPath, '..', 'server')}"`
						);
						child_process.execSync(
							`codesign --sign "Developer ID Application: Your Developer ID" --keychain /path/to/keychain "${path.join(outputPath, '..', 'dapper')}"`
						);
					}
				} catch (error) {
					console.error('Code signing error:', {
						error: JSON.stringify(error),
					});
				}
			}
		);
	}
}

module.exports = {
	target: 'electron-main',
	cache: {
		type: 'filesystem',
	},
	entry: './workspaces/electron-app/main/index.ts',
	module: {
		rules: require('./webpack.rules'),
	},
	resolve: {
		extensions: ['.js', '.ts', '.jsx', '.tsx', '.css', '.json'],
		modules: [path.resolve(__dirname, 'node_modules'), 'node_modules'],
	},
	node: {
		__filename: false,
		__dirname: false,
	},
	plugins: [
		new CopyWebpackPlugin({
			patterns: [
				{ from: 'workspaces/electron-app/main/assets' },
				{
					from: 'workspaces/angular-app/.dist/angular-app',
					to: '../renderer/angular_window',
					noErrorOnMissing: true,
				},
				{
					from: 'dapper',
					to: '../',
				},
				{
					from: 'app.json',
					to: '../',
				},
				{
					from: 'server',
					to: '../',
				},
			],
		}),
		new OpenOrchAfterEmitPlugin(),
	],
};
