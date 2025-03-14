/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import * as fs from 'fs-extra';
import * as path from 'path';
import { exec } from 'child_process';
import { promisify } from 'util';

interface ProgramOptions {
	outputDirectory: string;
	inputDirectory: string;
	packageName: string;
	packageExecutable?: string;
	publisher: string;
	publisherDisplayName?: string;
	identityName?: string;
	packageVersion: string;
	packageDisplayName?: string;
	packageDescription?: string;
	packageBackgroundColor?: string;
	protocol?: string;
	windowsKit: string;
	assets?: string;
	manifest?: string;
	makeappxParams?: string[];
}

// Util function to replace variables in the manifest template
function replaceManifestVariables(
	template: string,
	options: ProgramOptions
): string {
	let result = template
		.replace(/\${publisherName}/g, options.publisher)
		.replace(
			/\${publisherDisplayName}/g,
			options.publisherDisplayName || 'Reserved'
		)
		.replace(/\${identityName}/g, options.identityName || options.packageName)
		.replace(/\${packageVersion}/g, options.packageVersion)
		.replace(/\${packageName}/g, options.packageName)
		.replace(
			/\${packageExecutable}/g,
			options.packageExecutable || `app\\${options.packageName}.exe`
		)
		.replace(
			/\${packageDisplayName}/g,
			options.packageDisplayName || options.packageName
		)
		.replace(
			/\${packageDescription}/g,
			options.packageDescription || options.packageName
		)
		.replace(
			/\${packageBackgroundColor}/g,
			options.packageBackgroundColor || '#464646'
		)
		.replace(
			/\${protocol}/g,
			options.protocol
				? `<uap:Extension Category="windows.protocol"><uap:Protocol Name="${options.protocol}"></uap:Protocol></uap:Extension>`
				: ''
		);
	return result;
}

const execAsync = promisify(exec);

async function packageApp(options: ProgramOptions) {
	const preAppx = path.join(options.outputDirectory, 'pre-appx');
	const app = path.join(preAppx, 'app');
	const manifestTemplatePath = path.join(
		__dirname,
		'..',
		'template',
		'AppXManifest.xml'
	);
	const assetsTemplatePath = path.join(__dirname, '..', 'template', 'assets');
	const manifestPath = path.join(preAppx, 'AppXManifest.xml');
	const assetsPath = path.join(preAppx, 'assets');

	await fs.emptyDir(preAppx);
	await fs.ensureDir(app);
	await fs.ensureDir(assetsPath);

	await fs.copy(options.inputDirectory, app);
	await fs.copy(assetsTemplatePath, assetsPath);

	const manifestTemplate = await fs.readFile(manifestTemplatePath, 'utf8');
	const manifestContent = replaceManifestVariables(manifestTemplate, options);
	await fs.writeFile(manifestPath, manifestContent, 'utf8');

	const makeappx = path.join(options.windowsKit, 'makeappx.exe');
	const destination = path.join(
		options.outputDirectory,
		`${options.packageName}.appx`
	);
	const params = ['pack', '/d', preAppx, '/p', destination, '/o'].concat(
		options.makeappxParams || []
	);

	if (options.assets) {
		params.push('/l');
	}

	const cmd = `"${makeappx}" ${params.join(' ')}`;
	await execAsync(cmd);
}

const packageJson = require('../package.json');
const packageNodeVersion = process.env.PACKAGE_VERSION || packageJson.version;
const packageWinVersion = packageNodeVersion + '.0';

const programOptions: ProgramOptions = {
	outputDirectory: process.env.OUTPUT_DIRECTORY || './out/appx/',
	inputDirectory:
		process.env.INPUT_DIRECTORY || './out/make/squirrel.windows/x64/',
	packageName: process.env.PACKAGE_NAME || packageJson.name || 'OpenOrch',
	packageExecutable:
		process.env.PACKAGE_EXECUTABLE ||
		`app\\OpenOrch-${packageNodeVersion} Setup.exe`,
	publisher: process.env.PUBLISHER || 'CN=A2452F69-42C3-494B-A516-500954C5BE4E',
	publisherDisplayName:
		process.env.PUBLISHER_DISPLAY_NAME || packageJson.author || 'OpenOrch',
	identityName:
		'OpenOrch.OpenOrch' || process.env.IDENTITY_NAME || packageJson.name,
	packageVersion: packageWinVersion,
	packageDisplayName:
		process.env.PACKAGE_DISPLAY_NAME ||
		packageJson.displayName ||
		packageJson.name,
	packageDescription:
		process.env.PACKAGE_DESCRIPTION ||
		packageJson.description ||
		'An application.',
	packageBackgroundColor: process.env.PACKAGE_BACKGROUND_COLOR || '#464646',
	protocol: process.env.PROTOCOL,
	windowsKit:
		process.env.WINDOWS_KIT ||
		'C:\\Program Files (x86)\\Windows Kits\\10\\bin\\10.0.22621.0\\x64',
	assets: process.env.ASSETS,
	manifest: process.env.MANIFEST,
	makeappxParams: process.env.MAKEAPPX_PARAMS?.split(' ') || [],
};

packageApp(programOptions)
	.then(() => console.log('Packaging complete.'))
	.catch(console.error);
